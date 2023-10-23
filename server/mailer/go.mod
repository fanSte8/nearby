module nearby/mailer

go 1.21.2

replace nearby/common => ../common

require (
	github.com/caarlos0/env/v9 v9.0.0
	github.com/go-mail/mail/v2 v2.3.0
	github.com/joho/godotenv v1.5.1
	github.com/julienschmidt/httprouter v1.3.0
	nearby/common v0.0.0-00010101000000-000000000000
)

require (
	github.com/pascaldekloe/jwt v1.12.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/mail.v2 v2.3.1 // indirect
)
