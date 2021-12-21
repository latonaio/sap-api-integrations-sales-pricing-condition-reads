package main

import (
	sap_api_caller "sap-api-integrations-sales-pricing-condition-reads/SAP_API_Caller"
	"sap-api-integrations-sales-pricing-condition-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Sales_Pricing_Condition_Material_DistChannel_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/", l,
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"MaterialDistChannel", "MaterialDistChannelCustomer", "MaterialSalesOrgDistChannel", "MaterialSalesOrgDistChannelCustomer",
		}
	}

	caller.AsyncGetSalesPricingCondition(
		inoutSDC.SalesPricingConditionValidity.Material,
		inoutSDC.SalesPricingConditionValidity.DistributionChannel,
		inoutSDC.SalesPricingConditionValidity.Customer,
		inoutSDC.SalesPricingConditionValidity.SalesOrganization,
		accepter,
	)
}