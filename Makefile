.PHONY: build
build:
	go build -o ./.built/opts

.PHONY: test
test:
	$(MAKE) build
	for test in tests/*.caller.sh; do
		"$test"
	done
