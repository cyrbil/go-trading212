package trading212

import (
	trading213 "github.com/cyrbil/go-trading212/pkg/trading212/models"
	"iter"
	"testing"
)

func Test_Pies_Operations(t *testing.T) {
	t.Parallel()

	t.Run(
		"Test Pies FetchAllPies", func(t *testing.T) {
			t.Parallel()
			validateOperationItems[trading213.PieSummary](
				t,
				func(api *API) (iter.Seq[*trading213.PieSummary], error) {
					return api.Pies.FetchAllPies()
				}, `
			[{
				"cash": 0,
				"dividendDetails": {
					"gained": 0,
					"inCash": 0,
					"reinvested": 0
				},
				"id": 0,
				"progress": 0.5,
				"result": {
					"priceAvgInvestedValue": 0,
					"priceAvgResult": 0,
					"priceAvgResultCoef": 0,
					"priceAvgValue": 0
				},
				"status": "AHEAD"
			}]
		`,
			)
		},
	)

	t.Run(
		"Test Pies CreatePie", func(t *testing.T) {
			t.Parallel()
			validateOperationObject[trading213.PieDetails](
				t,
				func(api *API) (*trading213.PieDetails, error) {
					return api.Pies.CreatePie(trading213.PieRequest{})
				}, `
			{
				"instruments": [{
					"currentShare": 0,
					"expectedShare": 0,
					"issues": [{
						"name": "DELISTED",
						"severity": "IRREVERSIBLE"
					}],
					"ownedQuantity": 0,
					"result": {
						"priceAvgInvestedValue": 0,
						"priceAvgResult": 0,
						"priceAvgResultCoef": 0,
						"priceAvgValue": 0
					},
					"ticker": "string"
					}
				],
				"settings": {
					"creationDate": "2019-08-24T14:15:22Z",
					"dividendCashAction": "REINVEST",
					"endDate": "2019-08-24T14:15:22Z",
					"goal": 0,
					"icon": "string",
					"id": 0,
					"initialInvestment": 0,
					"instrumentShares": {
						"property1": 0,
						"property2": 0
					},
					"name": "string",
					"publicUrl": "string"
				}
			}
		`,
			)
		},
	)

	t.Run(
		"Test Pies DeletePie", func(t *testing.T) {
			t.Parallel()
			validateOperationObject[trading213.Empty](
				t,
				func(api *API) (*trading213.Empty, error) {
					return &trading213.Empty{}, api.Pies.DeletePie(0)
				}, "",
			)
		},
	)

	t.Run(
		"Test Pies FetchPie", func(t *testing.T) {
			t.Parallel()
			validateOperationObject[trading213.PieDetails](
				t,
				func(api *API) (*trading213.PieDetails, error) {
					return api.Pies.FetchPie(0)
				}, `
			{
				"instruments": [{
					"currentShare": 0,
					"expectedShare": 0,
					"issues": [{
						"name": "DELISTED",
						"severity": "IRREVERSIBLE"
					}],
					"ownedQuantity": 0,
					"result": {
						"priceAvgInvestedValue": 0,
						"priceAvgResult": 0,
						"priceAvgResultCoef": 0,
						"priceAvgValue": 0
					},
					"ticker": "string"
				}],
				"settings": {
					"creationDate": "2019-08-24T14:15:22Z",
					"dividendCashAction": "REINVEST",
					"endDate": "2019-08-24T14:15:22Z",
					"goal": 0,
					"icon": "string",
					"id": 0,
					"initialInvestment": 0,
					"instrumentShares": {
						"property1": 0,
						"property2": 0
					},
					"name": "string",
					"publicUrl": "string"
				}
			}
		`,
			)
		},
	)

	t.Run(
		"Test Pies UpdatePie", func(t *testing.T) {
			t.Parallel()
			validateOperationObject[trading213.PieDetails](
				t,
				func(api *API) (*trading213.PieDetails, error) {
					return api.Pies.UpdatePie(0, trading213.PieRequest{})
				}, `
			{
				"instruments": [{
					"currentShare": 0,
					"expectedShare": 0,
					"issues": [{
						"name": "DELISTED",
						"severity": "IRREVERSIBLE"
					}],
					"ownedQuantity": 0,
					"result": {
						"priceAvgInvestedValue": 0,
						"priceAvgResult": 0,
						"priceAvgResultCoef": 0,
						"priceAvgValue": 0
					},
					"ticker": "string"
				}],
				"settings": {
					"creationDate": "2019-08-24T14:15:22Z",
					"dividendCashAction": "REINVEST",
					"endDate": "2019-08-24T14:15:22Z",
					"goal": 0,
					"icon": "string",
					"id": 0,
					"initialInvestment": 0,
					"instrumentShares": {
						"property1": 0,
						"property2": 0
					},
					"name": "string",
					"publicUrl": "string"
				}
			}
		`,
			)
		},
	)

	t.Run(
		"Test Pies DuplicatePies", func(t *testing.T) {
			t.Parallel()
			validateOperationObject[trading213.PieDetails](
				t,
				func(api *API) (*trading213.PieDetails, error) {
					return api.Pies.DuplicatePies(0, trading213.PieMetaRequest{})
				}, `
			{
				"instruments": [{
					"currentShare": 0,
					"expectedShare": 0,
					"issues": [{
						"name": "DELISTED",
						"severity": "IRREVERSIBLE"
					}],
					"ownedQuantity": 0,
					"result": {
						"priceAvgInvestedValue": 0,
						"priceAvgResult": 0,
						"priceAvgResultCoef": 0,
						"priceAvgValue": 0
					},
					"ticker": "string"
				}],
				"settings": {
					"creationDate": "2019-08-24T14:15:22Z",
					"dividendCashAction": "REINVEST",
					"endDate": "2019-08-24T14:15:22Z",
					"goal": 0,
					"icon": "string",
					"id": 0,
					"initialInvestment": 0,
					"instrumentShares": {
						"property1": 0,
						"property2": 0
					},
					"name": "string",
					"publicUrl": "string"
				}
			}
		`,
			)
		},
	)
}
