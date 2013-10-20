PACKETSTRUCTS = $(wildcard pkt_*.go)
CLIENTFILES = clientConfig.go \
			  client.go \
			  interfaces.go \
			  packetUtils.go \
			  handler_client.go \
			  general_types.go

SERVERFILES = server.go \
			  serverFuncs.go \
			  packetUtils.go \
			  interfaces.go \
			  general_types.go \
			  serverTree.go \
			  handler_server.go

all: client server

client:
	go build -o ../client $(PACKETSTRUCTS) $(CLIENTFILES) 	

server: $(PACKETSTRUCTS) $(SERVERFILES)
	go build -o ../server $(PACKETSTRUCTS) $(SERVERFILES)

clean:
	rm -f ../client ../server
