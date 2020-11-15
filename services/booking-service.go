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
const OceanFreight = "Ocean Freight"

func GetBookingVendorDetails(booking *Booking) (url.Values, Vendors) {
	typeOfBusinessIdentifier(booking)
	vendors := Vendors{}
	switch typeOfBusiness {
	case DoorToDoor:
		validation = doorToDoorValidation(booking)
	case DoorToPort:
		validation = doorToPortValidation(booking)
	case PortToPort:
		validation = portToPortValidation(booking)
		vendor := getPortToPortServices(booking)
		vendorPortToPort := []Vendor{vendor}
		vendors = Vendors{Vendors: vendorPortToPort}
	case PortToDoor:
		validation = portToDoorValidation(booking)
		vendors = getPortToDoorServices(booking)
	default:
		fmt.Println("No error occurred while validating")
	}
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

func getPortToPortServices(booking *Booking) (vendor Vendor) {
	db, _ := database.DbConnection()

	// Get Freight
	freightVendor := getFreightVendorName(db, booking.SourcePort.Country, booking.DestinationPort.Country)
	freightService := Service{
		Name:    "Ocean freight",
		Country: booking.SourcePort.Country + " - " + booking.DestinationPort.Country,
	}

	freightServices := []Service{freightService}

	vendor = Vendor{
		Name:     freightVendor,
		Services: freightServices,
	}
	return vendor
}

func getPortToDoorServices(booking *Booking) (vendors Vendors) {
	db, _ := database.DbConnection()
	var importCustomsVendor, transportationVendor Vendor
	freightVendor := getPortToPortServices(booking)

	// Get Export Customs
	importCustomsVendorName := getVendorNameByService(db, booking.ImportCustoms.Country, "Customs")
	importCustomsService := Service{
		Name:    "Customs",
		Country: booking.ImportCustoms.Country,
	}

	// Get Transportation
	transportationVendorName := getVendorNameByService(db, booking.DestinationPort.Country, "Transportation")
	transportationService := Service{
		Name:    "Transportation",
		Country: booking.DestinationPort.Country,
	}

	// Combine Import Customs Vendor and Transportation Vendor
	if importCustomsVendorName == transportationVendorName {
		services := []Service{importCustomsService, transportationService}
		importCustomsVendor = Vendor{
			Name:     importCustomsVendorName,
			Services: services,
		}
	} else {
		services1 := []Service{importCustomsService}
		importCustomsVendor = Vendor{
			Name:     importCustomsVendorName,
			Services: services1,
		}
		services3 := []Service{transportationService}
		transportationVendor = Vendor{
			Name:     transportationVendorName,
			Services: services3,
		}
	}
	vendors = Vendors{Vendors: []Vendor{freightVendor, importCustomsVendor, transportationVendor}}
	return vendors
}

func getVendorNameByService(db *sql.DB, country Country, service Name) (vendor Name) {
	vendor = database.GetVendorByService(db, country, service)
	return vendor
}

func getFreightVendorName(db *sql.DB, sourceCountry Country, destinationCountry Country) (vendor Name) {
	vendor = database.GetFreightVendorByService(db, sourceCountry, destinationCountry, OceanFreight)
	return vendor
}
