PACKETSTRUCTS = $(wildcard pkt_*.go)
CLIENTFILES = clientConfig.go \
			  client.go \
			  interfaces.go \
			  packetUtils.go

SERVERFILES = server.go \
			  serverFuncs.go \
			  packetUtils.go \
			  interfaces.go 

all: client server

client:
	go build $(PACKETSTRUCTS) $(CLIENTFILES) -o ../client
	
server:
	go build $(PACKETSTRUCTS) $(SERVERFILES) -o ../server
