package ebay

import (
	"fmt"
	"testing"
)

var (
	test_application_id = "your_application_id_here"
)

func TestFindItemsByKeywords(t *testing.T) {
	fmt.Println("ebay.FindItemsByKeywords")
	e := New(test_application_id)
	response, err := e.FindItemsByKeywords(GLOBAL_ID_EBAY_DE, "ORIS DIVERS SIXTY-FIVE", 2)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	} else {
		fmt.Println("Timestamp: ", response.Timestamp)
		fmt.Println("Items:")
		fmt.Println("------")
		for _, i := range response.Items {
			fmt.Println("Title: ", i.Title)
			fmt.Println("------")
			fmt.Println("\tListing Url:     ", i.ListingUrl)
			fmt.Println("\tBin Price:       ", i.BinPrice)
			fmt.Println("\tCurrent Price:   ", i.CurrentPrice)
			fmt.Println("\tShipping Price:  ", i.ShippingPrice)
			fmt.Println("\tShips To:        ", i.ShipsTo)
			fmt.Println("\tSeller Location: ", i.Location)
			fmt.Println()

		}
	}
}

func TestFindCompletedItemsByKeywords(t *testing.T) {
	fmt.Println("ebay.FindCompletedItemsByKeywords")
	e := New(test_application_id)
	response, err := e.FindCompletedItemsByKeywords(GLOBAL_ID_EBAY_DE, "ORIS DIVERS SIXTY-FIVE", 2)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	} else {
		fmt.Println("Timestamp: ", response.Timestamp)
		fmt.Println("Items:")
		fmt.Println("------")
		for _, i := range response.Items {
			fmt.Println("Title: ", i.Title)
			fmt.Println("------")
			fmt.Println("\tListing Url:     ", i.ListingUrl)
			fmt.Println("\tBin Price:       ", i.BinPrice)
			fmt.Println("\tCurrent Price:   ", i.CurrentPrice)
			fmt.Println("\tShipping Price:  ", i.ShippingPrice)
			fmt.Println("\tShips To:        ", i.ShipsTo)
			fmt.Println("\tSeller Location: ", i.Location)
			fmt.Println()
		}
	}
}