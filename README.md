# command

## add schema

``` shell
go run -mod=mod entgo.io/ent/cmd/ent new Dept
```

## generate resource

``` shell
go generate ./ent
go run entgo.io/ent/cmd/ent generate ./schema --target ent
go run -mod=mod entgo.io/ent/cmd/ent generate --feature intercept,schema/snapshot ./schema  --target ./entcore
```

## run

``` shell
CGO_ENABLED=1 go run main.go
```

# test

``` shell
CGO_ENABLED=1 go test -v ./crud
```
