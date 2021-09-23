project: proto server client
proto:
	protoc -I=taskproto --go_out=plugins=grpc:. taskproto/task.proto
server:
	cd server && go build -o ../server .
client:
	cd client && go build -o ../client .