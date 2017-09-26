
.PHONY: lint
lint:
	gometalinter ./...

.PHONY: clean
clean:
	@rm -f prof base62.test bench.txt

.PHONY: benchmark
benchmark: clean
	go test -c
	./base62.test -test.bench=. -test.count=5 >bench.txt
	@cat bench.txt

.PHONY: mem-profile
mem-profile: clean
	go test -run=XXX -bench=. -memprofile=prof kkn.fi/base62
	go tool pprof base62.test prof

.PHONY: cpu-profile
cpu-profile: clean
	go test -run=XXX -bench=. -cpuprofile=prof kkn.fi/base62
	go tool pprof base62.test prof

.PHONY: build
build:
	go build kkn.fi/base62

.PHONY: test
test:
	go test kkn.fi/base62
