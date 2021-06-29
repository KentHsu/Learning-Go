# Notes

## go mod

Turn off go mod if it shows ```package motd/message is not in GOROOT```

```
go env -w GO111MODULE=off
```

## build cross-platform go binary

```
GOOS=linux GOARCH=amd64 go build -o app.linux
```

```
GOOS=darwin GOARCH=amd64 go build -o app.darwin
```