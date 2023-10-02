module plugins/todo

go 1.20

require (
	github.com/yosuke7040/go-todo-grpc v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.58.2
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230711160842-782d3b101e98 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)


// 上位ディレクトリのモジュールを参照するためにreplaceを追加
// replace github.com/yosuke7040/go-todo-grpc => ../..
// これでtidyすると以下のようになる
replace github.com/yosuke7040/go-todo-grpc v0.0.0-00010101000000-000000000000 => ../..

replace golang.org/x/net v0.15.0 => golang.org/x/net v0.12.0

replace golang.org/x/sys v0.12.0 => golang.org/x/sys v0.10.0

replace google.golang.org/grpc v1.58.2 => google.golang.org/grpc v1.58.1
