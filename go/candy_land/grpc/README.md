## Generate GRPC Server

```
protoc --go_out=plugins=grpc:. *.proto
```

## Generate mock files

```
mockgen -source=<path>.go -destination=<path>.go -package=<package-name>
```