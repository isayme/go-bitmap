.PHONY: test bench cover

test:
	go test .

bench:
	go test -bench .

cover:
	go test -v -covermode=count -coverprofile=coverage.out