package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

var (
	accountID  = os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	databaseID = os.Getenv("CLOUDFLARE_D1_DATABASE_ID")
	tableName  = os.Getenv("CLOUDFLARE_D1_TABLE_NAME")
)

var accountRC = cloudflare.AccountIdentifier(accountID)

func main() {
	api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// Most API calls require a Context
	ctx := context.Background()

	// Fetch user details on the account
	actual, err := api.QueryD1Database(ctx, accountRC, cloudflare.QueryD1DatabaseParams{
		DatabaseID: databaseID,
		SQL:        fmt.Sprintf("SELECT * FROM %s WHERE id > ?", tableName),
		Parameters: []string{"1"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", actual)
}
