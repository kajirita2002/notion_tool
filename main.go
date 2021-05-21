package main

import (
	"github.com/kaji2002/notion_tool/client"
)

func main() {
	c := client.NewClient("secret_klan5z0XsZTvjKNJAdy74TtvO8SlRX5uWmFNgUCopkr")
	client.FindDatabaseByID(c)
	client.AppendBlockChildren(c)
	client.CreatePage(c)
}
