package auth

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "embed"

	"github.com/traefik/traefik/v2/pkg/config/dynamic"
	"github.com/traefik/traefik/v2/pkg/log"
	"github.com/traefik/traefik/v2/pkg/middlewares"
	"github.com/traefik/traefik/v2/pkg/tracing"
)

const (
	powTypeName = "proofofworkauth"

	// Each hex character takes 4 bits
	bitsPerHexChar int = 4
	// ASCII code for number zero
	zero rune = 48

	timeExpFormat = "020106150405"
)

//go:embed pow.html
var challengePage []byte

type proofOfWorkAuth struct {
	next         http.Handler
	name         string
	removeHeader bool
	difficulty   int
	problemBits  int
	secret       []byte

	algProblem func() hash.Hash
	algVerify  func() hash.Hash
}

func NewProofOfWork(ctx context.Context, next http.Handler, config dynamic.ProofOfWorkAuth, name string) (http.Handler, error) {
	if config.Secret == "" {
		return nil, fmt.Errorf("error secret must be set", config.Secret)
	}

	pow := &proofOfWorkAuth{
		name: name,
		next: next,
		// TODO: Add check must be more then 0
		difficulty: config.Difficulty,
		// TODO: Add check must be more then 0
		problemBits:  config.ProblemBits,
		secret:       []byte(config.Secret),
		removeHeader: config.RemoveHeader,

		algProblem: sha1.New,
		algVerify:  sha256.New,
	}

	return pow, nil
}

func acceptableHeader(hash string, char rune, bits int) bool {
	wantZeros := bits / bitsPerHexChar
	for _, val := range hash[:wantZeros] {
		if val != char {
			return false
		}
	}
	return true
}

func (p *proofOfWorkAuth) verifyChallengeCookie(req *http.Request) bool {
	logger := log.FromContext(middlewares.GetLoggerCtx(req.Context(), p.name, powTypeName))

	challengeCookie, err := req.Cookie("challenge")
	if err != nil {
		return false
	}

	// "problem:hmac:exp:nonce:count"
	challengeParts := strings.Split(challengeCookie.Value, ":")
	if len(challengeParts) != 5 {
		logger.Debug("Incorrect challenge")
		tracing.SetErrorWithEvent(req, "Incorrect challenge")

		return false
	}

	hashesProblem := p.algProblem()
	io.WriteString(hashesProblem, challengeCookie.Value)
	sumProblem := hashesProblem.Sum(nil)

	// TODO: optimize performance
	if !acceptableHeader(hex.EncodeToString(sumProblem), zero, p.difficulty) {
		logger.Debug("Challenge verify failed")
		tracing.SetErrorWithEvent(req, "Challenge verify failed")

		return false
	}

	expValue := challengeParts[1]

	problemValue, err := hex.DecodeString(challengeParts[0])
	if err != nil {
		logger.Debug("Decode problem failed")
		tracing.SetErrorWithEvent(req, "Decode problem failed")

		return false
	}

	verifyValue, err := hex.DecodeString(challengeParts[2])
	if err != nil {
		logger.Debug("Decode verify failed")
		tracing.SetErrorWithEvent(req, "Decode verify failed")

		return false
	}

	verifyHashes := hmac.New(p.algVerify, p.secret)
	verifyHashes.Write(problemValue)
	io.WriteString(verifyHashes, expValue)

	if !hmac.Equal(verifyHashes.Sum(nil), verifyValue) {
		logger.Debug("Verify challenge signature failed")
		tracing.SetErrorWithEvent(req, "Verify challenge signature failed")

		return false
	}

	expTime, err := time.Parse(timeExpFormat, expValue)
	if err != nil {
		logger.Debug("Decode exp time failed")
		tracing.SetErrorWithEvent(req, "Decode exp time failed")

		return false
	}

	if time.Now().After(expTime) {
		logger.Debug("Auth token expired")
		tracing.SetErrorWithEvent(req, "Decode exp time expired")

		return false
	}

	return true
}

func (p *proofOfWorkAuth) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	logger := log.FromContext(middlewares.GetLoggerCtx(req.Context(), p.name, powTypeName))

	if !p.verifyChallengeCookie(req) {
		logger.Debug("Start authentication challenge page")

		expTime := time.Now().Add(time.Hour).Format(timeExpFormat)

		problem := make([]byte, p.problemBits)
		rand.Read(problem)

		verifyHashes := hmac.New(p.algVerify, p.secret)
		verifyHashes.Write(problem)

		io.WriteString(verifyHashes, expTime)

		rw.WriteHeader(http.StatusOK)

		challengePage = bytes.Replace(
			challengePage,
			[]byte("{problem}"),
			[]byte(hex.EncodeToString(problem)),
			-1,
		)
		challengePage = bytes.Replace(
			challengePage, []byte("{exp}"),
			[]byte(expTime),
			-1,
		)
		challengePage = bytes.Replace(
			challengePage, []byte("{sig}"),
			[]byte(hex.EncodeToString(verifyHashes.Sum(nil))),
			-1,
		)
		challengePage = bytes.Replace(
			challengePage, []byte("{bits}"),
			[]byte(strconv.Itoa(p.difficulty)),
			-1,
		)

		// show challenge page
		rw.Write(challengePage)
		return
	}

	logger.Debug("Authentication succeeded")

	if p.removeHeader {
		logger.Debug("Removing authorization header")
		req.Header.Del("Authorization")
	}
	p.next.ServeHTTP(rw, req)
}
