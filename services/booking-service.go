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
var validationFlag bool

const DoorToDoor = "DoorToDoor"
const DoorToPort = "DoorToPort"
const PortToPort = "PortToPort"
const PortToDoor = "PortToDoor"
const OceanFreight = "Ocean Freight"
const AirFreight = "Air Freight"

func GetBookingVendorDetails(booking *Booking) (url.Values, Vendors) {
	typeOfBusinessIdentifier(booking)

	vendors := Vendors{}
	switch typeOfBusiness {

	case DoorToDoor:
		validation, validationFlag = doorToDoorValidation(booking)
		if !validationFlag {
			vendors = getDoorToDoorServices(booking)
		}

	case DoorToPort:
		validation, validationFlag = doorToPortValidation(booking)
		if !validationFlag {
			vendors = getDoorToPortServices(booking)
		}

	case PortToPort:
		validation, validationFlag = portToPortValidation(booking)
		if !validationFlag {
			vendor := getPortToPortServices(booking)
			vendorPortToPort := []Vendor{vendor}
			vendors = Vendors{Vendors: vendorPortToPort}
		}
	case PortToDoor:
		validation, validationFlag = portToDoorValidation(booking)
		if !validationFlag {
			vendors = getPortToDoorServices(booking)
		}

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

func doorToDoorValidation(booking *Booking) (errs url.Values, validationFlag bool) {
	errs = url.Values{}
	errs = sourceValidation(booking, errs)
	errs = exportCustomsValidation(booking, errs)
	errs = sourcePortValidation(booking, errs)
	errs = destinationPortValidation(booking, errs)
	errs = importCustomsValidation(booking, errs)
	errs = destinationValidation(booking, errs)
	return errs, validationFlag

}
func doorToPortValidation(booking *Booking) (errs url.Values, validationFlag bool) {
	errs = url.Values{}
	errs = sourceValidation(booking, errs)
	errs = exportCustomsValidation(booking, errs)
	errs = sourcePortValidation(booking, errs)
	errs = destinationPortValidation(booking, errs)
	if len(errs) > 0 {
		validationFlag = true
	}
	return errs, validationFlag
}
func portToPortValidation(booking *Booking) (errs url.Values, validationFlag bool) {
	errs = url.Values{}
	errs = sourcePortValidation(booking, errs)
	errs = destinationPortValidation(booking, errs)
	if len(errs) > 0 {
		validationFlag = true
	}
	return errs, validationFlag
}
func portToDoorValidation(booking *Booking) (errs url.Values, validationFlag bool) {
	errs = url.Values{}
	errs = sourcePortValidation(booking, errs)
	errs = destinationPortValidation(booking, errs)
	errs = importCustomsValidation(booking, errs)
	errs = destinationValidation(booking, errs)
	if len(errs) > 0 {
		validationFlag = true
	}
	return errs, validationFlag
}

func sourceValidation(booking *Booking, errs url.Values) url.Values {
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

func exportCustomsValidation(booking *Booking, errs url.Values) url.Values {
	// ExportCustoms: Country Validation
	if booking.ExportCustoms.Country == "" {
		errs.Add("Country", "The ExportCustoms Country field is required!")
	}
	return errs
}

func sourcePortValidation(booking *Booking, errs url.Values) url.Values {
	// SourcePort: Address Validation
	if booking.SourcePort.City == "" {
		errs.Add("City", "The SourcePort City field is required!")
	}
	if booking.SourcePort.Country == "" {
		errs.Add("Country", "The SourcePort Country field is required!")
	}
	return errs
}

func destinationPortValidation(booking *Booking, errs url.Values) url.Values {
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

func importCustomsValidation(booking *Booking, errs url.Values) url.Values {
	// ImportCustoms: Country Validation
	if booking.ImportCustoms.Country == "" {
		errs.Add("Country", "The ImportCustoms Country field is required!")
	}
	return errs
}

func destinationValidation(booking *Booking, errs url.Values) url.Values {
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
	var importCustomsService, transportationService Service
	freightVendor := getPortToPortServices(booking)

	// Get Import Customs
	importCustomsVendorName := getVendorNameByService(db, booking.ImportCustoms.Country, "Customs")
	importCustomsService = Service{
		Name:    "Customs",
		Country: booking.ImportCustoms.Country,
	}

	// Get Transportation
	transportationVendorName := getVendorNameByService(db, booking.DestinationPort.Country, "Transportation")
	transportationService = Service{
		Name:    "Transportation",
		Country: booking.DestinationPort.Country,
	}

	// Combine Import Customs Vendor and Transportation Vendor
	if importCustomsVendorName == transportationVendorName {
		servicesList := []Service{importCustomsService, transportationService}
		importCustomsVendor = Vendor{
			Name:     importCustomsVendorName,
			Services: servicesList,
		}
		vendors = Vendors{Vendors: []Vendor{freightVendor, importCustomsVendor}}
	} else {
		importCustomServices := []Service{importCustomsService}
		importCustomsVendor = Vendor{
			Name:     importCustomsVendorName,
			Services: importCustomServices,
		}
		transportationServices := []Service{transportationService}
		transportationVendor = Vendor{
			Name:     transportationVendorName,
			Services: transportationServices,
		}
		vendors = Vendors{Vendors: []Vendor{freightVendor, importCustomsVendor, transportationVendor}}
	}

	return vendors
}
func getDoorToPortServices(booking *Booking) (vendors Vendors) {
	db, _ := database.DbConnection()
	var exportCustomsVendor, transportationVendor Vendor
	var exportCustomsService, transportationService Service
	freightVendor := getPortToPortServices(booking)

	// Get Export Customs
	exportCustomsVendorName := getVendorNameByService(db, booking.ExportCustoms.Country, "Customs")
	exportCustomsService = Service{
		Name:    "Customs",
		Country: booking.ExportCustoms.Country,
	}

	// Get Transportation
	transportationVendorName := getVendorNameByService(db, booking.SourcePort.Country, "Transportation")
	transportationService = Service{
		Name:    "Transportation",
		Country: booking.SourcePort.Country,
	}

	// Combine Export Customs Vendor and Transportation Vendor
	if exportCustomsVendorName == transportationVendorName {
		servicesList := []Service{exportCustomsService, transportationService}
		exportCustomsVendor = Vendor{
			Name:     exportCustomsVendorName,
			Services: servicesList,
		}
		vendors = Vendors{Vendors: []Vendor{freightVendor, exportCustomsVendor}}
	} else {
		importCustomServices := []Service{exportCustomsService}
		exportCustomsVendor = Vendor{
			Name:     exportCustomsVendorName,
			Services: importCustomServices,
		}
		transportationServices := []Service{transportationService}
		transportationVendor = Vendor{
			Name:     transportationVendorName,
			Services: transportationServices,
		}
		vendors = Vendors{Vendors: []Vendor{freightVendor, exportCustomsVendor, transportationVendor}}
	}

	return vendors
}
func getDoorToDoorServices(booking *Booking) (vendors Vendors) {
	doorToPortVendor := getDoorToPortServices(booking)
	portToDoorVendor := getPortToDoorServices(booking)
	freightVendor := getPortToPortServices(booking)

	// Combining all the services based on Vendor Name
	var finalVendor Vendor
	for _, vendor := range doorToPortVendor.Vendors {
		vendorName := vendor.Name
		for _, vendor2 := range portToDoorVendor.Vendors {
			if vendorName == vendor2.Name {
				vendor2.Services = append(vendor2.Services, vendor.Services...)
				finalVendor.Name = vendorName
				finalVendor.Services = vendor2.Services
			}
		}
	}
	vendors = Vendors{Vendors: []Vendor{freightVendor, finalVendor}}
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
