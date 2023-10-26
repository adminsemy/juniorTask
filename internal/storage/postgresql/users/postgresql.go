package postgresqlUsers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/adminsemy/juniorTask/internal/model"
	"github.com/adminsemy/juniorTask/internal/storage/postgresql"
)

const (
	tableName  = "users"
	name       = "name"
	surname    = "surname"
	patronymic = "patronymic"
	age        = "age"
	gender     = "gender"
)

type Repository struct {
	Db *sql.DB
}

// NewRepository — создаёт новый экземпляр репозитория
func NewRepository() *Repository {
	return &Repository{
		Db: postgresql.DbConnect(),
	}
}

// Add — добавляет запись в базу данных
func (r *Repository) Add(model *model.User, ctx context.Context) error {
	query := fmt.Sprintf(`
	INSERT INTO %s (%s, %s, %s, %s, %s)
	VALUES ($1, $2, $3, $4, $5)`, tableName, name, surname, patronymic, age, gender)
	sqlPrepare, err := r.Db.PrepareContext(ctx, query)
	defer sqlPrepare.Close()
	if err != nil {
		return err
	}
	sqlPrepare.QueryContext(ctx,
		model.Name,
		model.Surname,
		model.Patronymic,
		model.Age,
		model.Gender)
	return nil
}

// GetById — получает запись из базы данных по идентификатору
func (r *Repository) GetById(id int, ctx context.Context) (*model.User, error) {
	entity := &model.User{}
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, tableName)
	sql, err := r.Db.PrepareContext(ctx, query)
	if err != nil {
		return entity, err
	}
	err = sql.QueryRowContext(ctx, id).Scan(&entity.ID, &entity.Name, &entity.Surname, &entity.Patronymic, &entity.Age, &entity.Gender)
	if err != nil {
		return entity, err
	}

	return entity, nil
}

// Get — выполняет поиск записи в базе данных по параметрам
func (r *Repository) Get(params map[string]interface{}, ctx context.Context) (*model.User, error) {
	entity := &model.User{}
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, tableName)
	sql, err := r.Db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	err = sql.QueryRowContext(ctx, params["id"]).Scan(&entity.ID, &entity.Name, &entity.Surname, &entity.Patronymic, &entity.Age, &entity.Gender)
	return entity, err
}

// GetAll — получает все записи из базы данных
func (r *Repository) GetAll(ctx context.Context) ([]*model.User, error) {
	var entities []*model.User
	query := "SELECT * FROM " + tableName
	sql, err := r.Db.PrepareContext(ctx, query)
	if err != nil {
		return entities, err
	}
	rows, err := sql.QueryContext(ctx)
	if err != nil {
		return entities, err
	}
	for rows.Next() {
		entity := &model.User{}
		err := rows.Scan(&entity.ID, &entity.Name, &entity.Surname, &entity.Patronymic, &entity.Age, &entity.Gender)
		if err != nil {
			return entities, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}

// Update — обновляет запись в базе данных
func (r *Repository) Update(entity *model.User, ctx context.Context) error {
	query := "UPDATE " + tableName + " SET name = $1, surname = $2, patronymic = $3, age = $4, gender = $5 WHERE id = $6"
	sql, err := r.Db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	sql.QueryContext(ctx, entity.Name, entity.Surname, entity.Patronymic, entity.Age, entity.Gender, entity.ID)
	return nil
}
