# go-echo
Simple echo web server


## USE
```
curl -v -X GET "http://localhost:8080/health"    
===> 200

curl -v -X GET "http://localhost:8080/v1/address"
===> 200 {"host_address":"192.168.0.5"}
```

## REFACTORING
```
goimports -w .
```

## TESTS
```
go test -v -race -p 1 ./...
```