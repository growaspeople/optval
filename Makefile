.PHONY: build
build:
	go build -o ./.built/optval

.PHONY: test
test:
	$(MAKE) build
	bats tests.bats
