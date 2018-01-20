package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("crypto", func() {
	Title("Get price of a cryptocurrency")
	Description("A simple cryptocurrency price service")
	Scheme("http")
	Host("localhost")
	Version("0.1")
	Origin("*", func() { // Define CORS policy, may be prefixed with "*" wildcard
		Headers("*")   // One or more authorized headers, use "*" to authorize all
		Methods("GET") // One or more authorized HTTP methods
	})

})

var _ = Resource("currency", func() {
	BasePath("/currencies")
	DefaultMedia(CurrencyMedia)

	Action("show", func() {
		Description("Get currency price by currency ID")
		Routing(GET("/:currencyID"))
		Params(func() {
			Param("currencyID", String, "Currency ID")
		})
		Response(OK)
		Response(NotFound)
	})
})

// CurrencyMedia defines the media type used to render cryptocurrencies
var CurrencyMedia = MediaType("application/vnd.goa.example.currency+json", func() {
	Description("A cryptocurrency")
	Attributes(func() {
		Attribute("id", String, "Unique currency ID")
		Attribute("name", String, "Name of the cryptocurrency")
		Attribute("BTCprice", Number, "Bitcoin price of the cryptocurrency")
		Attribute("USDprice", Number, "USD price of the cryptocurrency")
		Required("id", "name", "BTCprice", "USDprice")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("BTCprice")
		Attribute("USDprice")
	})
})
