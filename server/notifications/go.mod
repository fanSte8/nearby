module nearby/notifications

go 1.21.2

replace nearby/common => ../common

require (
	github.com/caarlos0/env/v9 v9.0.0
	github.com/golang-migrate/migrate/v4 v4.16.2
	github.com/gorilla/mux v1.7.4
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.2
	nearby/common v0.0.0-00010101000000-000000000000
)

require (
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/pascaldekloe/jwt v1.12.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
)
