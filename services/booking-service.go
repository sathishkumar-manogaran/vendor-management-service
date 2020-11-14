package services

import (
	"github.com/sathishkumar-manogaran/vendor-management-service/models"
	"net/url"
)

func GetBookingVendorDetails(booking *models.Booking) (url.Values, models.Vendors) {
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

	//if len(booking.Destination.Address) > 0 {
	//	fmt.Printf("Booking %s\n", booking)
	//}

	validation := bookingRequestBasicValidation(booking)

	return validation, models.Vendors{
		Name:     "Test",
		Services: nil,
	}
}

func bookingRequestBasicValidation(booking *models.Booking) url.Values {
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
