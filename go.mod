module traefik

go 1.22

toolchain go1.22.1

require (
	github.com/containous/alice v0.0.0-20181107144136-d83ebdd94cbd
	github.com/traefik/paerser v0.2.0
	github.com/traefik/traefik/v2 v2.11.0
)

replace github.com/abbot/go-http-auth => github.com/containous/go-http-auth v0.4.1-0.20200324110947-a37a7636d23e

replace github.com/go-check/check => github.com/containous/check v0.0.0-20170915194414-ca0bf163426a

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/abbot/go-http-auth v0.0.0-00010101000000-000000000000 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/go-acme/lego/v4 v4.15.0 // indirect
	github.com/go-jose/go-jose/v3 v3.0.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gravitational/trace v1.1.16-0.20220114165159-14a9a7dd6aaf // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/mailgun/minheap v0.0.0-20170619185613-3dbe6c6bf55f // indirect
	github.com/mailgun/multibuf v0.1.2 // indirect
	github.com/mailgun/timetools v0.0.0-20141028012446-7e6055773c51 // indirect
	github.com/mailgun/ttlmap v0.0.0-20170619185759-c1c17f74874f // indirect
	github.com/miekg/dns v1.1.58 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/unrolled/secure v1.0.9 // indirect
	github.com/vulcand/oxy/v2 v2.0.0-20230427132221-be5cf38f3c1c // indirect
	github.com/vulcand/predicate v1.2.0 // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/term v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
)
