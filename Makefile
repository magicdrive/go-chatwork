CURRENTDIR:=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))


.PHONY: test
test:
	cd $(CURRENTDIR); \
	go test -v -coverpkg=. $(CURRENTDIR)/test/...
.PHONY: coverage
coverage:
	cd $(CURRENTDIR); \
	go test -coverpkg=. $(CURRENTDIR)/test/... -coverprofile=$(CURRENTDIR)/coverage.txt -covermode=count
