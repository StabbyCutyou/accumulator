.PHONY: test test_full bench benchmem deps check

test_full:
	go test -v -race -benchmem -bench=. -timeout=99m

bench:
	go test -v -bench=. -timeout=99m

benchmem:
	go test -v -benchmem -bench=. -timeout=99m

test:
	go test -v

deps:
	go get -u honnef.co/go/tools/cmd/staticcheck
	go get -u honnef.co/go/tools/cmd/gosimple
	go get -u honnef.co/go/tools/cmd/unused

check:
	staticcheck $$(go list ./... | grep -v /vendor/)
	gosimple $$(go list ./... | grep -v /vendor/)
	unused $$(go list ./... | grep -v /vendor/)