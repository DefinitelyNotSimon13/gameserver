default:
    just --list
 
run *ARGS:
    go run cmd/gameserver/main.go {{ARGS}}
