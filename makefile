SRCDIR=src

all: types client server

types:
	(cd sfs_types && go install .)

.PHONY: client
client:
	(cd sfs_client && go build -o ../client)

.PHONY: server
server:
	(cd sfs_server && go build -o ../server)

clean:
	rm -f client server
