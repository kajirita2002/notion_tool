package client

import (
	"context"
	"fmt"

	"github.com/dstotijn/go-notion"
)

func NewClient(apiKey string) *notion.Client {
	client := notion.NewClient(apiKey)
	return client
}

func FindDatabaseByID(c *notion.Client) {
	ctx := context.Background()
	db_id := "f38a21302e234460b6e7b0682bad2428"
	db, err := c.FindDatabaseByID(ctx, db_id)
	if err != nil {
		fmt.Println("can't find db:", err)
	}
	fmt.Println(db.ID)
}

func AppendBlockChildren(c *notion.Client) {
	text := []notion.RichText{
		{
			Type:      notion.RichTextTypeText,
			Text:      &notion.Text{Content: "こんにちは"},
		},
	}
	blocks := []notion.Block{
		{
			Object: "block",
			Type:   notion.BlockTypeHeading1,
			Heading1: &notion.Heading{Text: text},
		},
	}
	_, err := c.AppendBlockChildren(context.Background(), "de51c1f0-f9bd-4415-a3b7-b5b90243cdc7", blocks)
	if err != nil {
		fmt.Println("cannnot append Block:", err)
	}
}
func CreatePage(c *notion.Client) {
	prams := notion.CreatePageParams{
		ParentType: notion.ParentTypePage,
		ParentID:   "de51c1f0-f9bd-4415-a3b7-b5b90243cdc7",
		Title: []notion.RichText{
			{
				Type:      notion.RichTextTypeText,
				PlainText: "Hello World",
			},
		}}
	_, err := c.CreatePage(context.Background(), prams)
	if err != nil {
		fmt.Println("cannot create pages:", err)
	}
}
