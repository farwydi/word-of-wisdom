module traefik

go 1.22

toolchain go1.22.1

require github.com/traefik/traefik/v2 v2.11.0

replace github.com/abbot/go-http-auth => github.com/containous/go-http-auth v0.4.1-0.20200324110947-a37a7636d23e

replace github.com/go-check/check => github.com/containous/check v0.0.0-20170915194414-ca0bf163426a

require (
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/go-acme/lego/v4 v4.15.0 // indirect
	github.com/go-jose/go-jose/v3 v3.0.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/miekg/dns v1.1.58 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/traefik/paerser v0.2.0 // indirect
	github.com/umahmood/hashcash v0.0.0-20180415142835-f4a6a1a056f9 // indirect
	github.com/vulcand/oxy/v2 v2.0.0-20230427132221-be5cf38f3c1c // indirect
	golang.org/x/crypto v0.18.0 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/tools v0.17.0 // indirect
)
