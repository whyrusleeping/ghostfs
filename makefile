PACKETSTRUCTS = $(wildcard src/pkt_*.go)
CLIENTFILES = src/clientConfig.go \
			  src/client.go \
			  src/interfaces.go \
			  src/packetUtils.go \
			  src/general_types.go

SERVERFILES = src/server.go \
			  src/serverFuncs.go \
			  src/packetUtils.go \
			  src/interfaces.go \
			  src/general_types.go

all: client server

client: $(PACKETSTRUCTS) $(CLIENTFILES) 
	go build -o client $(PACKETSTRUCTS) $(CLIENTFILES) 	

server: $(PACKETSTRUCTS) $(SERVERFILES)
	go build -o server $(PACKETSTRUCTS) $(SERVERFILES)

clean:
	rm -f client server
