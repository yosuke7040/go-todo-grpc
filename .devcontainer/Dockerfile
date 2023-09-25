FROM mcr.microsoft.com/devcontainers/go:1-1.21-bullseye

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1



# https://github.com/7oh2020/connect-tasklist/tree/main/backend

# モック生成のためにvektra/mockeryをインストール
RUN curl -L https://github.com/vektra/mockery/releases/download/v2.25.0/mockery_2.25.0_Linux_arm64.tar.gz | tar xvz && mv ./mockery /usr/bin/mockery

# データベースのマイグレーションのためにgolang-migrate/migrateをインストール
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-arm64.tar.gz | tar xvz && mv ./migrate /usr/bin/migrate

# サーバーのライブリロードのためにcosmtrek/airをインストール
RUN go install github.com/cosmtrek/air@latest

# データベースアクセスのためにkyleconroy/sqlcをインストール
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

# gRPCとREST APIのコード生成のためにconnect-goとbufをインストール
RUN go install github.com/bufbuild/buf/cmd/buf@latest \
  && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
  && go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

# モジュールへ実行権限を与える(vscodeユーザーでもアクセス可能にする)
RUN sudo chmod -R a+rwX /go/pkg
