package repository

import (
	"database/sql"
)

type StoreRepository struct {
	db *sql.DB
}

type Row struct {
	Key   string
	Value string
}

func NewStoreRepository(db *sql.DB) *StoreRepository {
	return &StoreRepository{db: db}
}

func (r *StoreRepository) Insert(key, value string) error {
	_, err := r.db.Exec(
		"INSERT INTO store (`key`, `value`) VALUES (?, ?)",
		key, value,
	)
	return err
}

func (r *StoreRepository) Update(key, value string) error {
	_, err := r.db.Exec(
		"UPDATE store SET value = ? WHERE `key` = ?",
		value, key,
	)
	return err
}

func (r *StoreRepository) SelectValue(key string) (string, bool, error) {
	var val string
	err := r.db.QueryRow(
		"SELECT value FROM store WHERE `key` = ?",
		key,
	).Scan(&val)

	if err == sql.ErrNoRows {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}

	return val, true, nil
}

func (r *StoreRepository) Delete(key string) error {
	_, err := r.db.Exec(
		"DELETE FROM store WHERE `key` = ?",
		key,
	)
	return err
}

func (r *StoreRepository) RandomKey() (string, bool, error) {
	var k string
	err := r.db.QueryRow(
		"SELECT `key` FROM store ORDER BY RAND() LIMIT 1",
	).Scan(&k)

	if err == sql.ErrNoRows {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}

	return k, true, nil
}

func (r *StoreRepository) Dump() ([]Row, error) {
	rows, err := r.db.Query(
		"SELECT `key`, value FROM store",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Row

	for rows.Next() {
		var row Row
		if err := rows.Scan(&row.Key, &row.Value); err != nil {
			return nil, err
		}
		result = append(result, row)
	}

	return result, nil
}

func (r *StoreRepository) Mutate(key string) error {
	_, err := r.db.Exec(
		"UPDATE store SET value = CONCAT(value, '*') WHERE `key` = ? AND RAND() < 0.1",
		key,
	)
	return err
}
