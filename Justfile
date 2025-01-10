default:
    just --list
 
run *ARGS:
    go run cmd/gameserver/main.go {{ARGS}}

client *ARGS:
    go run cmd/testclient/main.go {{ARGS}}
