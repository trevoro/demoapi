clean:
	rm -f gin-bin

.PHONY: clean

watch:
	gin -laddr 127.0.0.1 --immediate

.PHONY: watch

test:
	go test -v ./...
