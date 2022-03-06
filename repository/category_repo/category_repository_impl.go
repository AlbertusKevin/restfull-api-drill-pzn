package category_repo

import (
	"context"
	"database/sql"
	"pzn-restful-api/exception"
	"pzn-restful-api/helper"
	"pzn-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository{
	return &CategoryRepositoryImpl{}
}

func (categoryRepository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "INSERT INTO category(name) VALUES(?)"
	result, err := tx.ExecContext(ctx,sql,category.Name)
	helper.PanicError(err)

	id, err := result.LastInsertId()
	helper.PanicError(err)

	category.Id = int(id)
	return category
}

func (categoryRepository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error) {
	sql := "UPDATE category SET name = ? WHERE id = ?"
	
	result, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helper.PanicError(err)
	
	rowsAffected, err := result.RowsAffected()
	helper.PanicError(err)

	if rowsAffected == 0{
		panic(exception.NewNotFoundError("Data tidak ditemukan"))
	}

	return category, nil
}

func (categoryRepository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) (bool,error){
	sql := "DELETE FROM category WHERE id = ?"
	
	result, err := tx.ExecContext(ctx, sql, category.Id)
	helper.PanicError(err)
	
	rowsAffected, err := result.RowsAffected()
	helper.PanicError(err)

	if rowsAffected == 0{
		panic(exception.NewNotFoundError("Data tidak ditemukan"))
	}

	return true, nil
}

func (categoryRepository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error) {
	sql := "SELECT id, name FROM category WHERE id=?"
	rows, err := tx.QueryContext(ctx, sql, id)
	helper.PanicError(err)
	defer rows.Close()

	var category domain.Category
	if rows.Next(){
		if err := rows.Scan(&category.Id, &category.Name); err != nil{
			helper.PanicError(err)
		}
		return category, nil
	}

	panic(exception.NewNotFoundError("Data tidak ditemukan"))
}

func (categoryRepository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	sql := "SELECT id, name FROM category"

	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicError(err)
	defer rows.Close()

	var categories []domain.Category
	
	for rows.Next(){
		var category domain.Category
		if err := rows.Scan(&category.Id, &category.Name); err != nil{
			helper.PanicError(err)
		}
		categories = append(categories, category)
	}

	return categories
}