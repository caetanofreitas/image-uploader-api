package repository

import (
	"database/sql"
	"errors"
	"time"
	"uploader/entity"
	"uploader/utils/errors_messages"
)

type ImageRepositoryDb struct {
	db *sql.DB
}

func NewImageRepository(db *sql.DB) *ImageRepositoryDb {
	return &ImageRepositoryDb{db: db}
}

func (i *ImageRepositoryDb) Insert(id string, name string, size float64, extension string, status string, error_message string) error {
	stmt, err := i.db.Prepare(`
		Insert into images (id, name, size, extension, status, error_message, created_at, updated_at)
		values(?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		id,
		name,
		size,
		extension,
		status,
		error_message,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (i *ImageRepositoryDb) Update(id string, name string, size float64, extension string, status string, error_message string) error {
	stmt, err := i.db.Prepare(`
		Update images
		set name = ?, size = ?, extension = ?, status = ?, error_message = ?, updated_at = ?
		where id = ?
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		name,
		size,
		extension,
		status,
		error_message,
		time.Now(),
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (i *ImageRepositoryDb) Get() ([]entity.Image, error) {
	var result = []entity.Image{}
	stmt, err := i.db.Prepare(`
		select * from images;
	`)
	if err != nil {
		return result, errors.New(errors_messages.DATABASE_ERROR)
	}

	rows, err := stmt.Query()

	if err != nil {
		return result, errors.New(errors_messages.DATABASE_ERROR)
	}

	if err = rows.Err(); err != nil {
		return result, errors.New(errors_messages.DATABASE_ERROR)
	}

	for rows.Next() {
		var img entity.Image
		created_at := ""
		updated_at := ""
		if err := rows.Scan(&img.ID, &img.Name, &img.Size, &img.Extension, &img.Status, &img.ErrorMessage, &created_at, &updated_at); err != nil {
			return result, err
		}
		result = append(result, img)
	}

	return result, nil
}

func (i *ImageRepositoryDb) GetDetail(id string) (entity.Image, error) {
	var result entity.Image
	var emptyImage entity.Image
	stmt, err := i.db.Prepare(`
		Select * from images where id = ?;
	`)
	if err != nil {
		return result, errors.New(errors_messages.DATABASE_ERROR)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return result, errors.New(errors_messages.DATABASE_ERROR)
	}

	if err = rows.Err(); err != nil {
		return result, errors.New(errors_messages.DATABASE_ERROR)
	}
	created_at := ""
	updated_at := ""

	for rows.Next() {
		if err := rows.Scan(&result.ID, &result.Name, &result.Size, &result.Extension, &result.Status, &result.ErrorMessage, &created_at, &updated_at); err != nil {
			return result, err
		}
	}

	if result == emptyImage {
		return result, errors.New(errors_messages.IMAGE_NOT_FOUND)
	}

	return result, nil
}
