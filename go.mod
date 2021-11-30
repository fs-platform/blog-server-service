module blog

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1 // indirect
	github.com/micro/micro/v2 v2.9.3 // indirect
	github.com/spf13/cast v1.3.0
	github.com/spf13/viper v1.6.3
	google.golang.org/protobuf v1.27.1
	gorm.io/driver/mysql v1.2.0
	gorm.io/gorm v1.22.3
)
