module github.com/netbsder/uranus

require (
	github.com/StackExchange/wmi v0.0.0-20181212234831-e0a55b97c705 // indirect
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/golang/protobuf v1.2.0
	github.com/json-iterator/go v1.1.5 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.2.0+incompatible
	github.com/lestrrat-go/strftime v0.0.0-20180821113735-8b31f9c59b0f // indirect
	github.com/mattn/go-isatty v0.0.6 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/pkg/errors v0.8.1 // indirect
	github.com/shirou/gopsutil v2.18.12+incompatible
	github.com/sirupsen/logrus v1.3.0
	github.com/uber-go/zap v1.9.1
	github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43 // indirect
	go.uber.org/atomic v1.3.2 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.9.1
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/yaml.v2 v2.2.2
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190228161510-8dd112bcdc25
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190303192550-c2f5717e611c
)

go 1.12
