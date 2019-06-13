package ebay

import (
	"encoding/xml"
	"github.com/heatxsink/go-httprequest"
)

const (
	GLOBAL_ID_EBAY_US = "EBAY-US"
	GLOBAL_ID_EBAY_FR = "EBAY-FR"
	GLOBAL_ID_EBAY_DE = "EBAY-DE"
	GLOBAL_ID_EBAY_IT = "EBAY-IT"
	GLOBAL_ID_EBAY_ES = "EBAY-ES"
)

type Item struct {
	ItemId string `xml:"itemId"`
	Title string `xml:"title"`
	Location string `xml:"location"`
	Country string `xml:"country"`
	CurrentPrice float64 `xml:"sellingStatus>currentPrice"`
	ShippingPrice float64 `xml:"shippingInfo>shippingServiceCost"`
	BinPrice float64 `xml:"listingInfo>buyItNowPrice"`
	ShipsTo []string  `xml:"shippingInfo>shipToLocations"`
	ListingUrl string `xml:"viewItemURL"`
	ImageUrl string `xml:"galleryURL"`
	Site string `xml:"globalId"`
}

type FindItemsByKeywordResponse struct {
	XmlName xml.Name `xml:"findItemsByKeywordsResponse"`
	Items []Item `xml:"searchResult>item"`
	Timestamp string `xml:"timestamp"`
}

type ErrorMessage struct {
	XmlName xml.Name `xml:"errorMessage"`
	Error Error `xml:"error"`
}

type Error struct {
	ErrorId string `xml:"errorId"`
	Domain string `xml:"domain"`
	Severity string `xml:"severity"`
	Category string `xml:"category"`
	Message string `xml:"message"`
	SubDomain string `xml:"subdomain"`
}

type EBay struct {
	ApplicationId string
	HttpRequest *httprequest.HttpRequest
}
