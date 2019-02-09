build:
	go build -v -o bin/web_api cmd/web/web.go

clean:
	find . -name "*.o" -o -name "*.test" -o -name "debug" | xargs rm -f
	-rm bin/*

start:
	bin/web_api

format:
	gofmt -w pkg cmd test
