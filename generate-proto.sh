# If it fails:
#  - `go install` the packages from `.default-golang-pkgs`
#  - Run `asdf reshim golang`

cd "`dirname \"$0\"`"

GOPATH=`go env GOPATH`

mkdir -p $GOPATH/src/github.com/nodeum-io
ln -svf $PWD $GOPATH/src/github.com/nodeum-io/nodeum-plugins

PROTO_PATH="--proto_path=$GOPATH/src:$GOPATH/pkg/mod/github.com/googleapis/googleapis@v0.0.0-20220408000435-23a2d2961f24:."

find . -name '*.type.proto' -exec protoc $PROTO_PATH --go_out=$GOPATH/src {} \;
find . -name '*.micro.proto' -exec protoc $PROTO_PATH --go_out=$GOPATH/src --micro_out=$GOPATH/src {} \;
find . -name '*.plugin.proto' -exec protoc $PROTO_PATH --go_out=$GOPATH/src --go-grpc_out=$GOPATH/src {} \;
