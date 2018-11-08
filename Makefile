# Makefile for neco

SUDO = sudo
FAKEROOT = fakeroot
ETCD_DIR = /tmp/neco-etcd

### for debian package
PACKAGES := fakeroot
WORKDIR := $(CURDIR)/work
CONTROL := $(WORKDIR)/DEBIAN/control
DOCDIR := $(WORKDIR)/usr/share/doc/neco
BINDIR := $(WORKDIR)/usr/bin
SBINDIR := $(WORKDIR)/usr/sbin
VERSION = 0.0.1-master
DEB = neco_$(VERSION)_amd64.deb

all:
	@echo "Specify one of these targets:"
	@echo
	@echo "    start-etcd  - run etcd on localhost."
	@echo "    stop-etcd   - stop etcd."
	@echo "    test        - run single host tests."
	@echo "    mod         - update and vendor Go modules."
	@echo "    deb         - build Debian package."
	@echo "    setup       - install dependencies."

start-etcd:
	systemd-run --user --unit neco-etcd.service etcd --data-dir $(ETCD_DIR)

stop-etcd:
	systemctl --user stop neco-etcd.service

test:
	test -z "$(gofmt -s -l . | grep -v '^vendor' | tee /dev/stderr)"
	golint -set_exit_status $(go list -mod=vendor ./... | grep -v /vendor/)
	go build -mod=vendor ./...
	go test -mod=vendor -race -v ./...
	go vet -mod=vendor ./...

mod:
	go mod tidy
	go mod vendor
	git add -f vendor
	git add go.mod go.sum

$(CONTROL): debian/DEBIAN/control
	rm -rf $(WORKDIR)
	cp -r debian $(WORKDIR)
	sed 's/@VERSION@/$(patsubst v%,%,$(VERSION))/' $< > $@

deb: $(DEB)

$(DEB): $(CONTROL)
	mkdir -p $(BINDIR)
	GOBIN=$(BINDIR) go install -mod=vendor ./pkg/neco
	mkdir -p $(SBINDIR)
	GOBIN=$(SBINDIR) go install -mod=vendor ./pkg/neco-updater ./pkg/neco-worker
	mkdir -p $(DOCDIR)
	cp README.md LICENSE $(DOCDIR)
	chmod -R g-w $(WORKDIR)
	$(FAKEROOT) dpkg-deb --build $(WORKDIR) .

setup:
	GO111MODULE=off go get -u golang.org/x/lint/golint
	$(SUDO) apt-get -y install --no-install-recommends $(PACKAGES)

clean:
	rm -rf $(ETCD_DIR) $(WORKDIR) $(DEB)

.PHONY:	all test mod deb setup clean