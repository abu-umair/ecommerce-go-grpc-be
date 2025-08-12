package database

import "database/sql"

func ConnectDB(connStr string) *sql.DB { //?menerima sebua string dan mengembalikan sql.DB
	{
		db, err := sql.Open("postgres", connStr) //?membuka koneksi ke database, mereturn db, err
		if err != nil {
			panic(err)
		}

		return db
	}
}
