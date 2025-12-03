package trading212

import (
	"iter"
	"testing"

	models "github.com/cyrbil/go-trading212/api/pkg/trading212"
)

func Test_Pies_Operations(t *testing.T) {
	t.Parallel()

	t.Run("Test Pies FetchAllPies", func(t *testing.T) {
		validateOperationItems[models.PieSummary](t,
			func(api *API) (iter.Seq[*models.PieSummary], error) {
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
		`)
	})

	t.Run("Test Pies CreatePie", func(t *testing.T) {
		validateOperationObject[models.PieDetails](t,
			func(api *API) (*models.PieDetails, error) {
				return api.Pies.CreatePie(models.PieRequest{})
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
		`)
	})

	t.Run("Test Pies DeletePie", func(t *testing.T) {
		validateOperationObject[models.Empty](t,
			func(api *API) (*models.Empty, error) {
				return &models.Empty{}, api.Pies.DeletePie(0)
			}, "")
	})

	t.Run("Test Pies FetchPie", func(t *testing.T) {
		validateOperationObject[models.PieDetails](t,
			func(api *API) (*models.PieDetails, error) {
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
		`)
	})

	t.Run("Test Pies UpdatePie", func(t *testing.T) {
		validateOperationObject[models.PieDetails](t,
			func(api *API) (*models.PieDetails, error) {
				return api.Pies.UpdatePie(0, models.PieRequest{})
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
		`)
	})

	t.Run("Test Pies DuplicatePies", func(t *testing.T) {
		validateOperationObject[models.PieDetails](t,
			func(api *API) (*models.PieDetails, error) {
				return api.Pies.DuplicatePies(0, models.PieMetaRequest{})
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
		`)
	})
}
