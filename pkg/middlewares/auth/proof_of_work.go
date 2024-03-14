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
	"net/http"
	"strconv"

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
)

//go:embed pow.html
var challengePage []byte

type proofOfWorkAuth struct {
	next         http.Handler
	name         string
	removeHeader bool
	difficulty   int
	secret       []byte
}

func NewProofOfWork(ctx context.Context, next http.Handler, config dynamic.ProofOfWorkAuth, name string) (http.Handler, error) {
	if config.Secret == "" {
		return nil, fmt.Errorf("error secret must be set", config.Secret)
	}

	pow := &proofOfWorkAuth{
		name:         name,
		next:         next,
		difficulty:   config.Difficulty,
		secret:       []byte(config.Secret),
		removeHeader: config.RemoveHeader,
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

func (p *proofOfWorkAuth) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	logger := log.FromContext(middlewares.GetLoggerCtx(req.Context(), p.name, powTypeName))

	challenge, err := req.Cookie("challenge")
	if err == nil {
		alg := sha1.New()
		alg.Write([]byte(challenge.Value))
		h := alg.Sum(nil)

		if !acceptableHeader(hex.EncodeToString(h[:]), zero, p.difficulty) {
			logger.Debug("Challenge verify failed")
			tracing.SetErrorWithEvent(req, "Challenge verify failed")
			err = fmt.Errorf("challenge verify failed")
		}
	}
	if err != nil {
		logger.Debug("Start authentication challenge page")

		problem := make([]byte, 10)
		rand.Read(problem)

		hasher := hmac.New(sha256.New, p.secret)
		hasher.Write(problem)

		rw.WriteHeader(http.StatusOK)
		challengePage = bytes.Replace(challengePage, []byte("{problem}"), []byte(hex.EncodeToString(problem)), -1)
		challengePage = bytes.Replace(challengePage, []byte("{sig}"), []byte(hex.EncodeToString(hasher.Sum(nil))), -1)
		challengePage = bytes.Replace(challengePage, []byte("{bits}"), []byte(strconv.Itoa(p.difficulty)), -1)

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
