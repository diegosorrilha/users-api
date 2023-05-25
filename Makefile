run: 
	go run .

build: clean
	go build .

clean:
	rm -f users-api

test:
	go test ./...
