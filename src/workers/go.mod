module github.com/quanndh/go-app/worker

go 1.17

require (
	github.com/gookit/config/v2 v2.1.6
	github.com/hibiken/asynq v0.23.0
	github.com/quanndh/go-app/adapter v0.0.0-00010101000000-000000000000
	go.uber.org/fx v1.18.2
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-redis/redis/v8 v8.11.2 // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/gookit/goutil v0.5.12 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/dig v1.15.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/sys v0.0.0-20220829200755-d48e67d00261 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace github.com/quanndh/go-app/adapter => ./../adapter
