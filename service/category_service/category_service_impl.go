package category_service

import (
	"context"
	"database/sql"
	"pzn-restful-api/helper"
	"pzn-restful-api/model/domain"
	"pzn-restful-api/model/web"
	"pzn-restful-api/repository/category_repo"

	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	CategoryRepository category_repo.CategoryRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewCategoryService(categoryRepository category_repo.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService{
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB: DB,
		Validate: validate,
	}
}

func (categoryService CategoryServiceImpl)Create(ctx context.Context, createRequest web.CategoryCreateRequest) web.CategoryResponse{
	err := categoryService.Validate.Struct(createRequest)
	helper.PanicError(err)
	
	tx, err := categoryService.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: createRequest.Name,
	}

	category = categoryService.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (categoryService CategoryServiceImpl)Update(ctx context.Context, updateRequest web.CategoryUpdateRequest) web.CategoryResponse{
	err := categoryService.Validate.Struct(updateRequest)
	helper.PanicError(err)
	
	tx, err := categoryService.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Id: updateRequest.Id,
		Name: updateRequest.Name,
	}

	category, err = categoryService.CategoryRepository.Update(ctx,tx,category)
	helper.PanicError(err)

	return helper.ToCategoryResponse(category)
}

func (categoryService CategoryServiceImpl)Delete(ctx context.Context, categoryId int) (bool, error){
	tx, err := categoryService.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Id: categoryId,
	}

	status, err := categoryService.CategoryRepository.Delete(ctx,tx,category)
	return status, err
}

func (categoryService CategoryServiceImpl)FindById(ctx context.Context, categoryId int) web.CategoryResponse{
	tx, err := categoryService.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)
	
	category, err := categoryService.CategoryRepository.FindById(ctx,tx,categoryId)
	helper.PanicError(err)

	return helper.ToCategoryResponse(category)
}

func (categoryService CategoryServiceImpl)FindAll(ctx context.Context) []web.CategoryResponse{
	tx, err := categoryService.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	categories := categoryService.CategoryRepository.FindAll(ctx,tx)

	return helper.ToCategoryResponses(categories)
}