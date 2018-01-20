package main

import (
	"crypto/app"

	"github.com/goadesign/goa"

	coinApi "github.com/miguelmota/go-coinmarketcap"
)

// CurrencyController implements the currency resource.
type CurrencyController struct {
	*goa.Controller
}

// NewCurrencyController creates a currency controller.
func NewCurrencyController(service *goa.Service) *CurrencyController {
	return &CurrencyController{Controller: service.NewController("CurrencyController")}
}

// Show runs the show action.
func (c *CurrencyController) Show(ctx *app.ShowCurrencyContext) error {
	// CurrencyController_Show: start_implement
	if ctx.CurrencyID == "" {
		return ctx.NotFound()
	}

	// call coinapi
	coinInfo, err := coinApi.GetCoinData(ctx.CurrencyID)
	if err != nil {
		return ctx.NotFound()
	} else {
		res := &app.GoaExampleCurrency{
			ID:       ctx.CurrencyID,
			Name:     coinInfo.Name,
			BTCprice: coinInfo.PriceBtc,
			USDprice: coinInfo.PriceUsd,
		}
		return ctx.OK(res)
	}

	// CurrencyController_Show: end_implement
}
