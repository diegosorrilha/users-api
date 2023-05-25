run: 
	go run .

build: clean
	go build .

clean:
	rm -f users-api

test:
	go test ./...

docs:
	godoc -http :9000
	$(info Running in http://localhost:9000/pkg/github.com/diegosorrilha/users-api/)
