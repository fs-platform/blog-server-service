package handler

import (
	"blog/app/models/article"
	"blog/app/models/category"
	blog "blog/proto/blog"
	"context"
)

type Blog struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Blog) Category(ctx context.Context, req *blog.CategoryIndexRequest, rsp *blog.CategoryIndexResponse) error {
	data, err := category.GetAll()
	if err != nil {
		return err
	}
	for _, val := range data {
		item := blog.CategoryIndexResponse_CategoryItem{
			Name:   val.Name,
			Id:     val.ID,
			Status: uint32(val.Status),
		}
		rsp.CategoryList = append(rsp.CategoryList, &item)
	}
	return nil
}

// Article 文章列表
func (e *Blog) Article(ctx context.Context, req *blog.ArticleRequest, rsp *blog.ArticleResponse) error {
	data, err := article.GetByCid(req.Id)
	if err != nil {
		return err
	}
	for _, val := range data {
		item := blog.ArticleResponse_ArticleItem{
			Name:   val.Name,
			Id:     val.ID,
			Status: uint32(val.Status),
			Like: val.Like,
			Author: val.Author,
		}
		rsp.ArticleList = append(rsp.ArticleList, &item)
	}
	return nil
}