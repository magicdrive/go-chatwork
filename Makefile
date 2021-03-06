CURRENTDIR:=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))
SHELL:=/bin/bash


.PHONY: test
test:
	cd $(CURRENTDIR); \
	go test -v $(CURRENTDIR)/test/...


.PHONY: coverage
coverage:
	cd $(CURRENTDIR); \
	for x in $$(find ./test/ -name \*_test.go | xargs dirname | sort | uniq | sed 's/^\.\/test\///'); \
	do \
		go test ./test/$${x} -coverpkg=./$${x} -covermode=count -coverprofile=profile.out ; \
		if [ -f ./profile.out ]; \
		then \
			cat ./profile.out >> ./coverage.txt; \
			rm -f ./profile.out; \
		fi \
	done;
