package ebay

import (
	"os"
	"fmt"
	"time"
	"testing"
	"strings"
	"encoding/csv"
)

var (
	test_application_id = "<Put your id here>"
)

func TestFindItemsByKeywords(t *testing.T) {
	fmt.Println("ebay.FindItemsByKeywords")
	e := New(test_application_id)
	response, err := e.FindItemsByKeywords(GLOBAL_ID_EBAY_DE, "ORIS DIVERS SIXTY-FIVE", 2)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	} else {
		current_time := time.Now().Local()

		file, err := os.Create(strings.Join([]string{"result-items-for-sale_",current_time.Format("2006-01-02"),".csv"},""))
		if err != nil {
			fmt.Println("Cannot create file", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		fmt.Println("Timestamp: ", response.Timestamp)
		fmt.Println("Items:")
		fmt.Println("------")

		err = writer.Write([]string{
			"ListingUrl",
			"BinPrice",
			"CurrentPrice",
			"ShippingPrice",
			"ShipsTo",
			"Location",
			"Country",
		})
		if err != nil {
			fmt.Println("Cannot write to file", err)
		}
		for _, i := range response.Items {
			err := writer.Write([]string{
				i.ListingUrl, 
				fmt.Sprintf("%f", i.BinPrice), 
				fmt.Sprintf("%f", i.CurrentPrice),
				fmt.Sprintf("%f", i.ShippingPrice),
				strings.Join(i.ShipsTo, " "),
				i.Location,
				i.Country,
			})
			if err != nil {
				fmt.Println("Cannot write to file", err)
			}

			fmt.Println("Title: ", i.Title)
			fmt.Println("------")
			fmt.Println("\tListing Url:     ", i.ListingUrl)
			fmt.Println("\tBin Price:       ", i.BinPrice)
			fmt.Println("\tCurrent Price:   ", i.CurrentPrice)
			fmt.Println("\tShipping Price:  ", i.ShippingPrice)
			fmt.Println("\tShips To:        ", i.ShipsTo)
			fmt.Println("\tSeller Location: ", i.Location)
			fmt.Println("\tSeller Country: ", i.Country)
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
			fmt.Println("\tSeller Country: ", i.Country)
			fmt.Println()
		}
	}
}