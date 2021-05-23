package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kaji2002/notion_tool/client"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("cannnot Read .env file: ", err)
	}
	fmt.Println("success read .env")
}

func main() {
	loadEnv()
	sercret := os.Getenv("SECRET")
	did := os.Getenv("DID")
	pid := os.Getenv("PID")

	c := client.NewClient(sercret)
	client.CreatePage(c, pid)
	client.FindDatabaseByID(c, did)
}
