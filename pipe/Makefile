all:build
	./w | ./r

1k:build
	./w 1024 | ./r

4k:build
	./w 4096 | ./r

build:
	@go build -o w writer.go
	@go build -o r reader.go

clean:
	@rm -f w r
