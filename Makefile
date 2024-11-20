all: test

test:
	go test -run=1000 -bench=.
