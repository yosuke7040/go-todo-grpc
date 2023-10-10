# go-todo-grpc

参考

<https://github.com/connectrpc/examples-go>
<https://connectrpc.com/docs/go/getting-started/>
<https://zenn.dev/7oh/articles/3b944f9b744932>

ディレクトリ構成
<https://go.dev/doc/modules/layout>

buf

```bash
buf lint
buf generate
```

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

grpcurl -protoset <(buf build -o -) -plaintext -H  "Authorization: Bearer <token>" localhost:8080 todo.v1.TodoService/List

```

evans <https://github.com/ktr0731/evans>

TODO: reflection 導入

```bash
evans --proto protos/todo/v1/todo.proto repl -p 8080
call Create
call --enrich Create

show service
```

シナリオテスト <https://github.com/zoncoen/scenarigo>

TODO: .devcontainer 便利そう。忘れないようにとりあえず入れてるだけ

migration

```bash
$ docker exec -it migrate sh

--- in container ----
$ migrate -path /app/sql/ -database "mysql://gogo:gogo@tcp(mysql:3306)/todo" up 1
$ migrate -path /app/sql/ -database "mysql://gogo:gogo@tcp(mysql:3306)/todo" down 1
$ exit
```

TODO: validation

[https://zenn.dev/mattn/articles/893f28eff96129]

TODO: error
status code をエラーに合わせて変更できるようにする
error.As,Is とか

db

```bash
docker exec -it todo-db sh
mysql -h 127.0.0.1 -u gogo todo -p
```

auth

sample user

- <email@example.com>
- p@ssw0rd

memo

> gRPC-Web では今のところ Client streaming と Bidirectional streaming ができません。
> Bidirectional streaming を行おうとすると、サーバからの Stream は問題ありませんが、クライアントからの二回目以降の送信ではブラウザのコンソールにエラーが出力されます。
> Server streaming 相当の動作にしかならないということだと思われます。

<https://qiita.com/kabochapo/items/6848457ea7a966baf957#bidirectional-streaming-rpc-%E3%81%AE%E4%BE%8B>
