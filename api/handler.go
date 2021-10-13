//go:generate go run github.com/prisma/prisma-client-go generate

package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gdp-v-sls/db"
)

var client = db.NewClient()

func init() {
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	user, err := client.User.FindFirst().Exec(context.Background())
	if err != nil {
		panic(err)
	}

	v, _ := json.MarshalIndent(user, "", "  ")
	log.Printf("data proxy response: %s", v)

	fmt.Fprintf(w, string(v))
}
