SRCDIR=src

PACKETSTRUCTS = $(wildcard $(SRCDIR)/pkt_*.go)
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

CLIENT = $(addprefix $(SRCDIR)/, $(CLIENTFILES)) 
SERVER = $(addprefix $(SRCDIR)/, $(SERVERFILES))


all: client server

client: $(CLIENT) $(PACKETSTRUCTS)
	go build -o client $(CLIENT)	$(PACKETSTRUCTS)


server: $(SERVER) $(PACKETSTRUCTS)
	go build -o server $(SERVER) $(PACKETSTRUCTS)


clean:
	rm -f client server
