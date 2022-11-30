package postgres

import (
	"errors"

	"github.com/Aziz0310/bootcamp/article/models"
)

// AddAuthor ...
func (stg Postgres) AddAuthor(id string, entity models.CreateAuthorModel) error {
	_, err := stg.db.Exec(`INSERT INTO 
	author (
		id,
		fullname
		) 
	VALUES (
		$1, 
		$2
		)`,
		id,
		entity.Fullname,
	)
	if err != nil {
		return err
	}
	return nil

}

// GetAuthorByID ...
func (stg Postgres) GetAuthorByID(id string) (models.GetAuthorByIdResp, error) {
	var res models.GetAuthorByIdResp
	err := stg.db.QueryRow(`SELECT 
		au.id,
		au.fullname,
		au.created_at,
		au.updated_at,
		au.deleted_at
    FROM author AS au  WHERE au.id = $1`, id).Scan(
		&res.ID,
		&res.Fullname,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.DeletedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetAuthorList ...
func (stg Postgres) GetAuthorList(offset, limit int, search string) (resp []models.Author, err error) {
	rows, err := stg.db.Queryx(`SELECT
	id,
	fullname,
	created_at,
	updated_at,
	deleted_at 
	FROM author WHERE deleted_at IS NULL AND ((title ILIKE '%' || $1 || '%') OR (body ILIKE '%' || $1 || '%'))
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var a models.Author

		err := rows.Scan(
			&a.ID,
			&a.Fullname,
			&a.CreatedAt,
			&a.UpdatedAt,
			&a.DeletedAt,
		)
		if err != nil {
			return resp, err
		}
		resp = append(resp, a)
	}

	return resp, err
}

// UpdateAuthor
func (stg Postgres) UpdateAuthor(entity models.UpdateAuthorModel) error {
	res, err := stg.db.NamedExec("UPDATE author  SET fullname=:f, updated_at=now() WHERE deleted_at IS NULL AND id=:id", map[string]interface{}{
		"id": entity.ID,
		"f":  entity.Fullname,
	})
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("author not found")
}

// DeleteAuthor
func (stg Postgres) DeleteAuthor(id string) error {
	res, err := stg.db.Exec("UPDATE author  SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL", id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("author not found")
}
