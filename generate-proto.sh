mkdir -p $GOPATH/src/github.com/nodeum-io
ln -svfT /workspaces/nodeum-plugins $GOPATH/src/github.com/nodeum-io/nodeum-plugins

PROTO_PATH="--proto_path=$GOPATH/src:$GOPATH/src/github.com/googleapis/googleapis:."

find . -name '*.type.proto' -exec protoc $PROTO_PATH --go_out=$GOPATH/src {} \;
find . -name '*.micro.proto' -exec protoc $PROTO_PATH --go_out=$GOPATH/src --micro_out=$GOPATH/src {} \;
find . -name '*.plugin.proto' -exec protoc $PROTO_PATH --go_out=$GOPATH/src --go-grpc_out=$GOPATH/src {} \;
