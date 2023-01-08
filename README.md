# echo-template

## Note

DB 認証情報の管理についてはどうにかしたい

## Overview

golang の echo を使って、とても雑な Web API 開発用テンプレートを用意するもの

## Requirements

すべて DockerImage で指定しているが列挙

- golang 1.18.4
- nginx 1.23.2
- mysql 8.0.31

## Usage

リポジトリのルートにて、以下コマンドを実行することでコンテナ環境の立ち上げが可能

```
docker-compose up -d
```

コンテナ立ち上げ後、localhost の 5000 番でリクエストを受け付けます

## References

- https://echo.labstack.com/
