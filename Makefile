GC=go build
GFILES=gconsts.go server.go sdatabase.go sconfig.go
.PHONY: default build clean
default: build
build: $(GFILES)
	$(GC) server.go gconsts.go sdatabase.go sconfig.go
clean:
	rm -f client.db server.db server client
