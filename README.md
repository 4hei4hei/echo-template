# echo-template

## Note

DB 認証情報の管理についてはどうにかしたい

## Overview

golang の echo ライブラリを使い、とても雑な Web API 開発用テンプレートを用意するもの

## Requirements

すべて DockerImage で指定していますが列挙

- golang 1.18.4
- Nginx 1.23.2
- MySQL 8.0

## Usage

リポジトリのルートにて、以下コマンドを実行することでコンテナ環境の立ち上げが可能

```
docker-compose up -d
```

コンテナ立ち上げ後、localhost の 5000 番でリクエストを受け付けます

## References

- https://echo.labstack.com/
