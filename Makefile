.PHONY: build build-prod build-test run-prod run-test clean all

all: build-prod build-test

build-prod:
	go build -o bin/stripe-pan-sub-prod ./cmd/prod

build-test:
	go build -o bin/stripe-pan-sub-test ./cmd/test

build-customer:
	go build -o bin/stripe-pan-sub-customer ./cmd/customer

run-prod: build-prod
	./bin/stripe-pan-sub-prod

run-test: build-test
	./bin/stripe-pan-sub-test

run-customer: build-customer
	./bin/stripe-pan-sub-customer

clean:
	rm -rf bin/
