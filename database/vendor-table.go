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
					vendor_name varchar(500) unique,
					vendor_rating int,
					service_id int,
					created_at    datetime default CURRENT_TIMESTAMP,
					updated_at    datetime default CURRENT_TIMESTAMP,
					foreign key (service_id) references service(id)
				)`
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating vendor table", err)
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

func CreateMasterCountryTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS country
				(
					id int primary key auto_increment,
					name varchar(150) unique not null,
					is_active varchar(1) not null
				)`
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating country table", err)
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

func CreateMasterServiceTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS service
				(
					id int primary key auto_increment,
					name varchar(150) unique not null,
					country_id int,
					is_active varchar(1) not null,
					foreign key (country_id) references country(id)
				)`
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating services table", err)
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
