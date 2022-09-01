# OctaneServer

Octaneのサーバーサイドです。  
開発環境の人は完全デフォルト設定でローカル動作（mongo, s3無し）で動作しますが、本番運用する場合はクレデンシャル等設定してから公開してください
。

Built on [Fiber](https://gofiber.io/).

## Build

### Local

```sh
go build .
```

### Docker

```sh
docker build .
```

```sh
sudo docker buildx build --platform linux/386,linux/amd64,linux/arm64,linux/arm/v6,linux/arm/v7 -t ghcr.io/team-kamo/server:latest --push .
```

## Run

### Dev

```sh
go run .
```

### Binary

```sh
./OctaneServer
```

### Docker

```sh
docker run --rm -it -p 3000:3000 ghcr.io/team-kamo/server:latest
```

### Docker-compose

```sh
docker-compose up
```

## Other libs

Check out [go.mod](./go.mod).
