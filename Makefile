.PHONY: test bench cover

test:
	go test .

bench:
	go test -bench .

cover:
	go test -v -covermode=count -coverprofile=coverage.out

cover_html: cover
	go tool cover -html=coverage.out