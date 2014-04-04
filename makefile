SRCDIR=src

all: types client server

types:
	(cd gfs_types && go install .)

.PHONY: client
client:
	(cd gfs_client && go install .)

.PHONY: server
server:
	(cd gfs_server && go install .)

clean:
	rm -f client server
