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

release:	## build and create 'hn' tool executable at './install/'.
	GOOS=linux GOARCH=amd64 go build -o install/linux/hn
	GOOS=darwin GOARCH=amd64 go build -o install/macos/hn
	GOOS=windows GOARCH=amd64 go build -o install/windows/hn.exe
	@echo "Release Done, check 'install' dir."
# 	tar -czvf release/hn-tool-release.tar.gz -C install .