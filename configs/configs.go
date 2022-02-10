package configs

import (
	"github.com/duel80003/my-url-shorter/tools"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("./configs/.env")
	if err != nil {
		tools.Logger.Fatalf("Error loading .env file %s", err)
	}
}
