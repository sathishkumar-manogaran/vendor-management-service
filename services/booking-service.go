package services

import (
	"fmt"
	. "github.com/sathishkumar-manogaran/vendor-management-service/models"
	"net/url"
	"reflect"
)

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

	//if len(booking.Destination.Address) > 0 {
	//	fmt.Printf("Booking %s\n", booking)
	//}

	//typeOfBusinessIdentifier(booking)

	validation := bookingRequestBasicValidation(booking)

	s1 := Service{
		Name:    "Ocean freight",
		Country: "India - Singapore",
	}
	s2 := Service{
		Name:    "Customs",
		Country: "Singapore",
	}
	s3 := Service{
		Name:    "Transportation",
		Country: "India",
	}

	ser1 := []Service{s1, s2}
	ser2 := []Service{s3}

	v := Vendor{
		Name:     "V1",
		Services: ser1,
	}
	v1 := Vendor{
		Name:     "V2",
		Services: ser2,
	}

	ve := []Vendor{v, v1}

	//return validation, Vendors{Vendors: []Vendor{{
	//	Name: "V1",
	//	Services: []Service{{
	//		Name:    "Ocean freight",
	//		Country: "India - Singapore",
	//	}},
	//}}}
	return validation, Vendors{Vendors: ve}
}

func typeOfBusinessIdentifier(booking *Booking) {
	//1. Door to door - Will have all fields.
	//2. Door to port - Will have fields 1 to 4.
	//3. Port to port - Will have only fields 3 and 4.
	//4. Port to door - Will have fields 3 to 6.

	// if we dont have 5 (ImportCustoms) then it will go to 2,3
	// if we don't have 1 (Source) then it will go to 3
	// else 2
	// if we dont have 1 (Source) then it will go to 3,4 and check for 5 (ImportCustoms)
	// if 5 (ImportCustoms) present then it wil go to 4
	// else 3
	//

	// Check Source and Destination first
	// if source not present then 3,4
	// if destination not present then 2,3

	if reflect.DeepEqual(booking.Source, Source{}) {
		fmt.Println("Source Present")
	} else {
		fmt.Println("Source Not Present")
	}
}

func bookingRequestBasicValidation(booking *Booking) url.Values {
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

//func (x Person) IsStructureEmpty() bool {
//	return reflect.DeepEqual(x, Person{})
//}
