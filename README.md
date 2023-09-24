# go-todo-grpc

参考

<https://github.com/connectrpc/examples-go>
<https://connectrpc.com/docs/go/getting-started/>

ディレクトリ構成
<https://go.dev/doc/modules/layout>

リクエスト

```bash
grpcurl \
    -protoset <(buf build -o -) -plaintext \
    -d '{"title": "test"}' \
    localhost:8080 protos.todo.v1.TodoService/Create

grpcurl \
    -protoset <(buf build -o -) -plaintext \
    -d '{"id": 1}' \
    localhost:8080 protos.todo.v1.TodoService/Read

```

evans <https://github.com/ktr0731/evans>

TODO: reflection導入

```evans
evans --proto protos/todo/v1/todo.proto repl -p 8080
call Create

show service
```

シナリオテスト <https://github.com/zoncoen/scenarigo>
