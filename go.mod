module github.com/kordar/goframework-cron

go 1.18

replace github.com/kordar/gocron => ../gocron

require (
	github.com/kordar/gocron v0.0.0-00010101000000-000000000000
	github.com/kordar/godb v0.0.7
	github.com/kordar/gorabbitmq v1.0.1
	github.com/spf13/cast v1.7.0

)

require (
	github.com/kordar/gologger v0.0.8 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/streadway/amqp v1.0.0 // indirect
)
