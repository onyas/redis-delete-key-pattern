

### Usage
```
cp config.yaml.tpl config.yaml

#update your redis host and deleted keys in the config.yaml
vim config.yaml

#execute search and delete keys
go run *.go
```


### How to create a golang project
``` shell
mkdir redis-cluster-clean
cd redis-cluster-clean
go mod init github.com/onyas/redis-cluster-clean
go get github.com/go-redis/redis/v8
```