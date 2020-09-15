module github.com/hourglasshoro/reviery.api

go 1.14

require (
	github.com/go-redis/redis/v8 v8.0.0
	github.com/google/uuid v1.1.2
	github.com/joho/godotenv v1.3.0
	github.com/stretchr/testify v1.6.1
	google.golang.org/grpc v1.32.0
)

replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.2.0
