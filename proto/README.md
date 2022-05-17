
## Generate from proto

Install [asdf](https://asdf-vm.com/).

Install required plugins:

```bash
asdf plugin add protoc
asdf plugin add golang
asdf plugin add buf
```

Install required version with `asdf install`

Install these required golang modules:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install go-micro.dev/v4/cmd/protoc-gen-micro@v4
```

Generate with `buf generate`. Publish with `buf push`.

