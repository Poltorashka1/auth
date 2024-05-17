package utils

import (
	"github.com/jackc/pgx/v4"
)

func ScanAll(dest interface{}, rows pgx.Rows) error {

	for rows.Next() {

		err := rows.Scan()
		if err != nil {
			return err
		}
	}

	return nil
}

func ScanOne(dest interface{}, rows pgx.Rows) error {
	for rows.Next() {
		err := rows.Scan(&dest)
		if err != nil {
			return err
		}
	}

	return nil
}
