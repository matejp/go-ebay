package ebay

import (
	"fmt"
	"encoding/xml"
	"strconv"
	"net/url"
	"errors"
	"github.com/heatxsink/go-httprequest"
)


func New(application_id string) *EBay {
	e := EBay {}
	e.ApplicationId = application_id
	e.HttpRequest = httprequest.NewWithDefaults()
	return &e
} 

// https://developer.ebay.com/Devzone/shopping/docs/CallRef/FindProducts.html#UsageDetails
func (e *EBay) build_search_url(global_id string, operation_call string, keywords string, entries_per_page int) (string, error) {
	var u *url.URL
	u, err := url.Parse("http://svcs.ebay.com/services/search/FindingService/v1")
	// u, err := url.Parse("http://open.api.ebay.com/shopping?")
	if err != nil {
		return "", err
	}
	params := url.Values{}
	params.Add("OPERATION-NAME", operation_call)
	params.Add("SERVICE-VERSION", "1.0.0")
	params.Add("SECURITY-APPNAME", e.ApplicationId)
	params.Add("GLOBAL-ID", global_id)
	params.Add("RESPONSE-DATA-FORMAT", "XML")
	params.Add("REST-PAYLOAD", "")
	params.Add("keywords", keywords)
	params.Add("paginationInput.entriesPerPage", strconv.Itoa(entries_per_page))
	// params.Add("outputSelector(0)", "AspectHistogram")
	// params.Add("outputSelector(1)", "CategoryHistogram")
	params.Add("itemFilter(0).name", "ListingType")
	params.Add("itemFilter(0).value(0)", "FixedPrice")
	params.Add("itemFilter(0).value(1)", "AuctionWithBIN")
	u.RawQuery = params.Encode()
	return u.String(), err
}

func (e *EBay) FindItemsByKeywords(global_id string, keywords string, entries_per_page int) (FindItemsByKeywordResponse, error) {
	var response FindItemsByKeywordResponse
	url, err := e.build_search_url(global_id, "findItemsByKeywords", keywords, entries_per_page)
	if err != nil {
		return response, err
	}
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11"
	body, status_code, err := e.HttpRequest.Get(url, headers)
	if err != nil {
		return response, err
	}
	if status_code != 200 {
		var em ErrorMessage
		err = xml.Unmarshal([]byte(body), &em)
		if err != nil {
			return response, err
		}
		return response, errors.New(em.Error.Message)
	} else {
		err = xml.Unmarshal([]byte(body), &response)
		if err != nil {
			return response, err
		}
	}
	return response, err
}

func (r *FindItemsByKeywordResponse) Dump() {
	fmt.Println("FindItemsByKeywordResponse")
	fmt.Println("--------------------------")
	fmt.Println("Timestamp: ", r.Timestamp)
	fmt.Println("Items:")
	fmt.Println("------")
	for _, i := range r.Items {
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


func (e *EBay) FindCompletedItemsByKeywords(global_id string, keywords string, entries_per_page int) (FindItemsByKeywordResponse, error) {
	var response FindItemsByKeywordResponse
	url, err := e.build_search_url(global_id, "findCompletedItems", keywords, entries_per_page)
	if err != nil {
		return response, err
	}
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11"
	body, status_code, err := e.HttpRequest.Get(url, headers)
	if err != nil {
		return response, err
	}
	if status_code != 200 {
		var em ErrorMessage
		err = xml.Unmarshal([]byte(body), &em)
		if err != nil {
			return response, err
		}
		return response, errors.New(em.Error.Message)
	} else {
		err = xml.Unmarshal([]byte(body), &response)
		if err != nil {
			return response, err
		}
	}
	return response, err
}