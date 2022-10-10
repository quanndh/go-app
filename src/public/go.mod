module github.com/quanndh/go-app/public

go 1.16

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/gookit/config/v2 v2.1.6
	github.com/hibiken/asynq v0.23.0
	go.uber.org/fx v1.18.2
	golang.org/x/crypto v0.0.0-20221005025214-4161e89ecf1b
	gopkg.in/yaml.v3 v3.0.1

)

require (
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/quanndh/go-app/adapter v0.0.0-00010101000000-000000000000
)

replace github.com/quanndh/go-app/adapter => ./../adapter
