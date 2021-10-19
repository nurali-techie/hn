help:		## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build:		## build 'hn' tool executable and output at './out/'.
	go build -o out/hn

install:	## install 'hn' tool using go.
	go install

clean:		## remove 'hn' tool executable from './out/'.
	rm -f out/hn

go-mod:		## run go module specific commands, 'tidy' and 'verify'.
	go mod tidy
	go mod verify
