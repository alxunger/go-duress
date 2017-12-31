package db

import "database/sql"

var err error

func YieldSingleRow(rows *sql.Rows, ret *int64) error {
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(ret)
		if err != nil {
			return err
		}
	}
	return err
}
