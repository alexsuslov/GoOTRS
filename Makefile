All: bin/gotrs

bin/gotrs:
	go build -ldflags "-X main.version=`git describe --abbrev=0`" -o bin/gotrs main.go
