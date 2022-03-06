package category_repo

import (
	"context"
	"database/sql"
	"pzn-restful-api/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) (domain.Category, error)
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category) (bool,error)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.Category,error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}