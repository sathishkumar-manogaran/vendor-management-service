package services

import (
	"database/sql"
	"fmt"
	"github.com/sathishkumar-manogaran/vendor-management-service/database"
	. "github.com/sathishkumar-manogaran/vendor-management-service/models"
	"net/url"
)

var typeOfBusiness string
var validation url.Values

const DoorToDoor = "DoorToDoor"
const DoorToPort = "DoorToPort"
const PortToPort = "PortToPort"
const PortToDoor = "PortToDoor"

func GetBookingVendorDetails(booking *Booking) (url.Values, Vendors) {
	// TODO
	// 1. JSON Deep Validation (same address for custom clearance)
	// 2. Identifying type of booking
	//1. Door to door - Will have all fields.
	//2. Door to port - Will have fields 1 to 4.
	//3. Port to port - Will have only fields 3 and 4.
	//4. Port to door - Will have fields 3 to 6.
	// 3. Assign vendors
	// 4. Save Booking Details and Assigned vendor details (optional)
	// 5. ORM Implementation
	// 6. Vendor Table Creation
	//	(3 table required. 1. Vendor Name 2. Vendor Country (master table) 3. Vendor Services (master table))
	// 7. Type of booking Table Creation

	typeOfBusinessIdentifier(booking)
	vendors := Vendors{}
	// based on type of business and country look for services
	// if it is PortToPort, then need to look for customs in source and destination port and Ocean freight or Air freight
	// if it is PortToDoor, then need to look for customs in source and destination port and destination transportation and Ocean freight or Air freight
	// if it is DoorToPort, then need to look for Source Transportation customs in source and destination port and Ocean freight or Air freight
	// if it is DoorToDoor, then need to look for Source and destination Transportation customs in source and destination port and Ocean freight or Air freight
	switch typeOfBusiness {
	case DoorToDoor:
		validation = doorToDoorValidation(booking)
	case DoorToPort:
		validation = doorToPortValidation(booking)
	case PortToPort:
		validation = portToPortValidation(booking)
		// customs and Ocean freight or Air freight
		//importCountry := booking.ImportCustoms.Country
		//exportCountry := booking.ExportCustoms.Country
		// for this country 1st find out customs and transportation
		vendors = getPortToPortServices(booking)
	case PortToDoor:
		validation = portToDoorValidation(booking)
	default:
		fmt.Println("No error occurred while validating")
	}

	//s1 := Service{
	//	Name:    "Ocean freight",
	//	Country: "India - Singapore",
	//}
	//s2 := Service{
	//	Name:    "Customs",
	//	Country: "Singapore",
	//}
	//s3 := Service{
	//	Name:    "Transportation",
	//	Country: "India",
	//}
	//
	//ser1 := []Service{s1, s2}
	//ser2 := []Service{s3}
	//
	//v := Vendor{
	//	Name:     "V1",
	//	Services: ser1,
	//}
	//v1 := Vendor{
	//	Name:     "V2",
	//	Services: ser2,
	//}
	//
	//ve := []Vendor{v, v1}

	//return validation, Vendors{Vendors: []Vendor{{
	//	Name: "V1",
	//	Services: []Service{{
	//		Name:    "Ocean freight",
	//		Country: "India - Singapore",
	//	}},
	//}}}
	return validation, vendors
}

func typeOfBusinessIdentifier(booking *Booking) {
	if booking.Source.Address == "" {
		if booking.ImportCustoms.Country == "" {
			typeOfBusiness = PortToPort
		} else {
			typeOfBusiness = PortToDoor
		}

	} else if booking.ImportCustoms.Country == "" {
		typeOfBusiness = DoorToPort
	} else {
		typeOfBusiness = DoorToDoor
	}
	fmt.Println(typeOfBusiness)
}

func doorToDoorValidation(booking *Booking) url.Values {
	errs := url.Values{}
	errs = sourceValidation(booking)
	errs = exportCustomsValidation(booking)
	errs = sourcePortValidation(booking)
	errs = destinationPortValidation(booking)
	errs = importCustomsValidation(booking)
	errs = destinationValidation(booking)
	return errs
}
func doorToPortValidation(booking *Booking) url.Values {
	errs := url.Values{}
	errs = sourceValidation(booking)
	errs = exportCustomsValidation(booking)
	errs = sourcePortValidation(booking)
	errs = destinationPortValidation(booking)
	return errs
}
func portToPortValidation(booking *Booking) url.Values {
	errs := url.Values{}
	errs = sourcePortValidation(booking)
	errs = destinationPortValidation(booking)
	return errs
}
func portToDoorValidation(booking *Booking) url.Values {
	errs := url.Values{}
	errs = sourcePortValidation(booking)
	errs = destinationPortValidation(booking)
	errs = importCustomsValidation(booking)
	errs = destinationValidation(booking)
	return errs
}

func sourceValidation(booking *Booking) url.Values {
	errs := url.Values{}
	// Source: Address Validation
	if booking.Source.Address == "" {
		errs.Add("Address", "The Source Address field is required!")
	}
	if booking.Source.City == "" {
		errs.Add("City", "The Source City field is required!")
	}
	if booking.Source.Country == "" {
		errs.Add("Country", "The Source Country field is required!")
	}
	return errs
}

func exportCustomsValidation(booking *Booking) url.Values {
	errs := url.Values{}
	// ExportCustoms: Country Validation
	if booking.ExportCustoms.Country == "" {
		errs.Add("Country", "The ExportCustoms Country field is required!")
	}
	return errs
}

func sourcePortValidation(booking *Booking) url.Values {
	errs := url.Values{}
	// SourcePort: Address Validation
	if booking.SourcePort.City == "" {
		errs.Add("City", "The SourcePort City field is required!")
	}
	if booking.SourcePort.Country == "" {
		errs.Add("Country", "The SourcePort Country field is required!")
	}
	return errs
}

func destinationPortValidation(booking *Booking) url.Values {
	errs := url.Values{}
	// DestinationPort: City Validation
	if booking.DestinationPort.City == "" {
		errs.Add("City", "The DestinationPort City field is required!")
	}
	// DestinationPort: Country Validation
	if booking.DestinationPort.Country == "" {
		errs.Add("Country", "The DestinationPort Country field is required!")
	}
	return errs
}

func importCustomsValidation(booking *Booking) url.Values {
	errs := url.Values{}
	// ImportCustoms: Country Validation
	if booking.ImportCustoms.Country == "" {
		errs.Add("Country", "The ImportCustoms Country field is required!")
	}
	return errs
}

func destinationValidation(booking *Booking) url.Values {
	errs := url.Values{}
	// Destination: Address Validation
	if booking.Destination.Address == "" {
		errs.Add("Address", "The Destination Address field is required!")
	}
	if booking.Destination.City == "" {
		errs.Add("City", "The Destination City field is required!")
	}
	if booking.Destination.Country == "" {
		errs.Add("Country", "The Destination Country field is required!")
	}
	return errs
}

func getPortToPortServices(booking *Booking) (vendors Vendors) {
	db, _ := database.DbConnection()
	//// Get Import Customs
	//importCustomsVendor := getCustoms(db, booking.SourcePort.Country)
	//
	//service1 := Service{
	//	Name:    "Customs",
	//	Country: booking.SourcePort.Country,
	//}
	//
	//// Get Export Customs
	//exportCustomsVendor := getCustoms(db, booking.DestinationPort.Country)
	//service2 := Service{
	//	Name:    "Customs",
	//	Country: booking.DestinationPort.Country,
	//}

	// Get Freight
	freightVendor := getFreight(db, booking.SourcePort.Country, booking.DestinationPort.Country)
	service3 := Service{
		Name:    "Ocean freight",
		Country: booking.SourcePort.Country + " - " + booking.DestinationPort.Country,
	}

	ser2 := []Service{service3}

	v1 := Vendor{
		Name:     freightVendor,
		Services: ser2,
	}

	ve := []Vendor{v1}
	vendors = Vendors{Vendors: ve}
	return vendors
}

func getCustoms(db *sql.DB, country Country) (vendor Name) {
	vendor = database.GetVendorByService(db, country, "Customs")
	return vendor
}

func getFreight(db *sql.DB, sourceCountry Country, destinationCountry Country) (vendor Name) {
	vendor = database.GetVendorByService(db, sourceCountry, "Ocean Freight")
	return vendor
}
