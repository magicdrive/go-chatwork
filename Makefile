CURRENTDIR:=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))


.PHONY: test
test:
	go test -v $(CURRENTDIR)/test/...
