package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/dstotijn/go-notion"
)

func NewClient(apiKey string) *notion.Client {
	client := notion.NewClient(apiKey)
	return client
}

func FindDatabaseByID(c *notion.Client, did string) *notion.Database{
	ctx := context.Background()
	db, err := c.FindDatabaseByID(ctx, did)
	if err != nil {
		fmt.Println("can't find db:", err)
	}
	return &db
}

func FindBlockChildrenByID(c *notion.Client, bid string) *notion.BlockChildrenResponse {
	ctx := context.Background()
	query := &notion.PaginationQuery{}
	result, err := c.FindBlockChildrenByID(ctx, bid, query)
	if err != nil {
		fmt.Println("cannot find block child by id: ", err)
	}
	return &result
}

func FindPageByID(c *notion.Client, id string) *notion.Page {
	ctx := context.Background()
	page, err := c.FindPageByID(ctx, id)
	if err != nil {
		fmt.Println("cannnot find page by id:", err)
	}
	return &page
}

func FindUserByID(c *notion.Client) *notion.ListUsersResponse {
	ctx := context.Background()
	query := &notion.PaginationQuery{}
	result, err := c.ListUsers(ctx, query)
	if err != nil {
		fmt.Println("cannot list users:", err)
	}
	return &result
}

func AppendBlockChildren(c *notion.Client, id string) *notion.Block {
	text := []notion.RichText{
		{
			Type: notion.RichTextTypeText,
			Text: &notion.Text{Content: "こんにちは"},
		},
	}
	blocks := []notion.Block{
		{
			Object:   "block",
			Type:     notion.BlockTypeHeading1,
			Heading1: &notion.Heading{Text: text},
		},
	}
	block, err := c.AppendBlockChildren(context.Background(), id, blocks)
	if err != nil {
		fmt.Println("cannnot append Block:", err)
	}
	return &block

}
func CreatePage(c *notion.Client, pid string) *notion.Page {
	prams := notion.CreatePageParams{
		ParentType: notion.ParentTypePage,
		ParentID:   pid,
		Title: []notion.RichText{
			{
				Type: notion.RichTextTypeText,
				Text: &notion.Text{
					Content: "新しいpageです",
				},
			},
		},
	}
	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(prams)
	if err != nil {
		fmt.Println("cannot encode:", err)
	}
	fmt.Println("this is encode:", body)
	page, err := c.CreatePage(context.Background(), prams)
	if err != nil {
		fmt.Println("cannot create pages:", err)
	}
	return &page
}

func QueryDatabase(c *notion.Client, did string) notion.DatabaseQueryResponse {
	ctx := context.Background()
	query := notion.DatabaseQuery{}
	result, err := c.QueryDatabase(ctx, did, &query)
	if err != nil {
		fmt.Println("cannot query databse:", err)
	}
	return result
}

func Search(c *notion.Client) notion.SearchResponse {
	ctx := context.Background()
	opts := &notion.SearchOpts{}
	result, err := c.Search(ctx, opts)
	if err != nil {
		fmt.Println("cannot search it")
	}
	return result
}
func UpdatePageProps(c *notion.Client, pid string) *notion.Page {
	ctx := context.Background()
	params := notion.UpdatePageParams{Title: []notion.RichText{{PlainText: "更新したもの"}}}
	page, err := c.UpdatePageProps(ctx, pid, params)
	if err != nil {
		fmt.Println("cannnot update page props:", err)
	}
	return &page
}
