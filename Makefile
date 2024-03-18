all: build demo

build:
	./scripts/build

demo:
	docker compose up --build
