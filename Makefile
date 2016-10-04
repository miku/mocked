mocked: cmd/mocked/main.go mock.go
	go build -o mocked cmd/mocked/main.go

clean:
	rm -f mocked
