run-cilint:
	# make sure you have setup GOPATH and install golint
	GO111MODULE=on
	CGO_ENABLED=0
	go list -e -compiled -test=true -export=false -deps=true -find=false -tags= -- ./... > /dev/null
	golangci-lint cache clean
	golangci-lint run -v --timeout=240s