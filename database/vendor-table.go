package database

import (
	"context"
	"database/sql"
	"fmt"
	. "github.com/sathishkumar-manogaran/vendor-management-service/models"
	"log"
	"time"
)

func CreateMasterCountryTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS country
				(
					id int primary key auto_increment,
					name varchar(150) not null,
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
					name varchar(150) not null,
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

func CreateVendorTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS vendor
				(
					id    int primary key auto_increment,
					vendor_name varchar(500),
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

func GetVendorByService(db *sql.DB, country Country, service Name) (vendor Name) {
	row := db.QueryRow("select v.vendor_name from vendor v "+
		"left join service s on s.id = v.service_id "+
		"left join country c on c.id = s.country_id where c.name=? and s.name=?", country, service)
	switch err := row.Scan(&vendor); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(vendor)
	default:
		panic(err)
	}
	return vendor
}

func GetFreightVendorByService(db *sql.DB, sourceCountry Country, destinationCountry Country, service Name) (vendor Name) {
	var sourceCountryVendor Name
	rows, err := db.Query("select v.vendor_name from vendor v left join service s on s.id = v.service_id left join country c on c.id = s.country_id where c.name = ? and s.name = ?", sourceCountry, service)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&sourceCountryVendor)
		if err != nil {
			panic(err)
			db.Close()
		}
		fmt.Println(sourceCountryVendor)
		row := db.QueryRow("select v.vendor_name from vendor v left join service s on s.id = v.service_id left join country c on c.id = s.country_id where c.name = ? and s.name = ? and v.vendor_name = ?", destinationCountry, service, sourceCountryVendor)
		switch err := row.Scan(&vendor); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			fmt.Println(vendor)
			break
		default:
			panic(err)
			db.Close()
		}
	}
	return vendor
}
