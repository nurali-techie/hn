build:
	go build -o out/hn

install:
	go install

clean:
	rm -f out/hn

go-mod:
	go mod tidy
	go mod verify
