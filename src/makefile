PACKETSTRUCTS = $(wildcard pkt_*.go)
CLIENTFILES = clientConfig.go \
			  client.go \
			  interfaces.go \
			  packetUtils.go \
			  general_types.go

SERVERFILES = server.go \
			  serverFuncs.go \
			  packetUtils.go \
			  interfaces.go \
			  general_types.go

all: server

client:
	go build -o ../client $(PACKETSTRUCTS) $(CLIENTFILES) 	

server: $(PACKETSTRUCTS) $(SERVERFILES)
	go build -o ../server $(PACKETSTRUCTS) $(SERVERFILES)

clean:
	rm -f ../client ../server
