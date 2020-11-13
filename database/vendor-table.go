package database

import (
	"context"
	"database/sql"
	"log"
	"time"
)

func CreateVendorTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS vendor
				(
					id    int primary key auto_increment,
					vendor_name  varchar(500),
					vendor_rating int,
					created_at    datetime default CURRENT_TIMESTAMP,
					updated_at    datetime default CURRENT_TIMESTAMP
				)`
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return err
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return nil
}
