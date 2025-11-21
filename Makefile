help:		## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build:		## build 'hn' tool executable at './out/'.
	go build -o out/hn
	@echo "Build Done, check 'out' dir."

install:	## install 'hn' tool using go.
	go install
	@echo "Install Done."

clean:		## remove 'hn' tool executable from './out/'.
	rm -rf out/
	@echo "Clean Done."

go-mod:		## run go module specific commands, 'tidy' and 'verify'.
	go mod tidy
	go mod verify

release:	## build and create 'hn' tool executable at './install/' for multiple platforms.
	mkdir -p install
	GOOS=linux GOARCH=amd64 go build -o install/hn-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o install/hn-darwin-amd64
	GOOS=windows GOARCH=amd64 go build -o install/hn-windows-amd64.exe
	@echo "Release Done, check 'install' dir."