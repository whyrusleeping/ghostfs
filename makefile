SRCDIR=src

all: types client server

types:
	(cd sfs_types && go install .)

.PHONY: client
client:
	(cd sfs_client && go install .)

.PHONY: server
server:
	(cd sfs_server && go install .)

clean:
	rm -f client server
