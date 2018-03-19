# generate golang struct in according to ssh_config `man documantation`
go run ./gen/ssh_config_generate.go

# formatting generated files
gofmt -s -w *.go
