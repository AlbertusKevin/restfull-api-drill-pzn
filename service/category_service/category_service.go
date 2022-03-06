package category_service

import (
	"context"
	"pzn-restful-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, createRequest web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, updateRequest web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int) (bool,error)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
