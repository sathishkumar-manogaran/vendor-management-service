package services

import (
	"fmt"
	"github.com/sathishkumar-manogaran/vendor-management-service/models"
)

func GetBookingVendorDetails(booking *models.Booking) models.Vendors {
	if len(booking.Destination.Address) > 0 {
		fmt.Printf("Booking %s\n", booking)
	}
	return models.Vendors{
		Name:     "Test",
		Services: nil,
	}
}
