.PHONY: build
build:
	go build -o ./.built/optval

.PHONY: deb
deb:
	go build -o ./.built/deb/optval_0.1-1/usr/local/bin/optval
	cp --recursive ./DEBIAN ./.built/deb/optval_0.1-1/
	sleep 1s # To avoid tar's "file changed as we read it" error
	(cd ./.built/deb && dpkg-deb --build optval_0.1-1)

.PHONY: test
test:
	$(MAKE) build
	bats tests.bats
