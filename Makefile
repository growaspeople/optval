.PHONY: build
build:
	go build -o ./.built/optval

.PHONY: test
test:
	$(MAKE) build
	for test in tests/*.caller.sh; do
		"$test"
	done
