# ecommerce
ecommerce services

![Go](https://github.com/PauloKeller/ecommerce/workflows/Go/badge.svg) ![Node.js CI](https://github.com/PauloKeller/ecommerce/workflows/Node.js%20CI/badge.svg)

# Development

```
docker-compose -f docker-compose.dev.yml up --build
```

# Generate proto files

```
protoc --go_out=plugins=grpc:. *.proto
```


