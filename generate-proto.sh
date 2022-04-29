# If it fails:
#  - Run `asdf reshim golang`

cd "`dirname \"$0\"`"

GOOGLEAPIS=github.com/googleapis/googleapis@v0.0.0-20220408000435-23a2d2961f24

go install $GOOGLEAPIS
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install go-micro.dev/v4/cmd/protoc-gen-micro@v4

GOPATH=`go env GOPATH`

NODEUM_PLUGINS=$GOPATH/src/github.com/nodeum-io/nodeum-plugins
GOOGLEAPIS_PATH=$GOPATH/pkg/mod/$GOOGLEAPIS

mkdir -p `dirname $NODEUM_PLUGINS`
unlink $NODEUM_PLUGINS
ln -svf $PWD $NODEUM_PLUGINS

PROTO_PATH="--proto_path=$GOPATH/src:$GOOGLEAPIS_PATH:."

find . -name '*.type.proto' -exec protoc $PROTO_PATH --go_out=$GOPATH/src {} \;
find . -name '*.micro.proto' -exec protoc $PROTO_PATH --go_out=$GOPATH/src --micro_out=$GOPATH/src {} \;
find . -name '*.plugin.proto' -exec protoc $PROTO_PATH --go_out=$GOPATH/src --go-grpc_out=$GOPATH/src {} \;
