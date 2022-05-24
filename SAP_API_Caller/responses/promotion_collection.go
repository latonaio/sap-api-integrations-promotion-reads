package responses

type PromotionCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                           string      `json:"ObjectID"`
			ETag                               string      `json:"ETag"`
			AccountType                        string      `json:"AccountType"`
			AccountTypeText                    string      `json:"AccountTypeText"`
			Currency                           string      `json:"Currency"`
			CurrencyText                       string      `json:"CurrencyText"`
			ID                                 string      `json:"ID"`
			Name                               string      `json:"Name"`
			Objective                          string      `json:"Objective"`
			ObjectiveText                      string      `json:"ObjectiveText"`
			Priority                           string      `json:"Priority"`
			PriorityText                       string      `json:"PriorityText"`
			ProductPlanningBasisCode           string      `json:"ProductPlanningBasisCode"`
			ProductPlanningBasisCodeText       string      `json:"ProductPlanningBasisCodeText"`
			PromotionType                      string      `json:"PromotionType"`
			PromotionTypeText                  string      `json:"PromotionTypeText"`
			ConsistencyStatusCode              string      `json:"ConsistencyStatusCode"`
			ConsistencyStatusCodeText          string      `json:"ConsistencyStatusCodeText"`
			ExternalStatusCode                 string      `json:"ExternalStatusCode"`
			ExternalStatusCodeText             string      `json:"ExternalStatusCodeText"`
			LifeCycleStatusCode                string      `json:"LifeCycleStatusCode"`
			LifeCycleStatusCodeText            string      `json:"LifeCycleStatusCodeText"`
			CreationDate                       string      `json:"CreationDate"`
			LastChangeDate                     string      `json:"LastChangeDate"`
			Tactic                             string      `json:"Tactic"`
			TacticText                         string      `json:"TacticText"`
			PlannigAccountName                 string      `json:"PlannigAccountName"`
			PlanningAccountID                  string      `json:"PlanningAccountID"`
			PlanPeriodStartDate                string      `json:"PlanPeriodStartDate"`
			PlanPeriodEndDate                  string      `json:"PlanPeriodEndDate"`
			BuyingPeriodStartDate              string      `json:"BuyingPeriodStartDate"`
			BuyingPeriodEndDate                string      `json:"BuyingPeriodEndDate"`
			EmployeeResponsible                string      `json:"EmployeeResponsible"`
			EmployeeResposibleID               string      `json:"EmployeeResposibleID"`
			SalesUnitOrganisationCentreID      string      `json:"SalesUnitOrganisationCentreID"`
			SalesUnitOrganisationCentreName    string      `json:"SalesUnitOrganisationCentreName"`
			SalesOrganisationID                string      `json:"SalesOrganisationID"`
			SalesOrganisationName              string      `json:"SalesOrganisationName"`
			DistributionChannelCode            string      `json:"DistributionChannelCode"`
			DistributionChannelCodeText        string      `json:"DistributionChannelCodeText"`
			DivisionCode                       string      `json:"DivisionCode"`
			DivisionCodeText                   string      `json:"DivisionCodeText"`
			SalesTerritoryID                   string      `json:"SalesTerritoryID"`
			SalesTerritoryName                 string      `json:"SalesTerritoryName"`
			ActualPeriodStartDate              string      `json:"ActualPeriodStartDate"`
			ActualPeriodEndDate                string      `json:"ActualPeriodEndDate"`
			Campaign                           string      `json:"Campaign"`
			CampaignDescription                string      `json:"CampaignDescription"`
			TargetGroupID                      string      `json:"TargetGroupID"`
			TargetGroupDescription             string      `json:"TargetGroupDescription"`
			ExternalID                         string      `json:"ExternalID"`
			InformationLifeCycleStatusCode     string      `json:"InformationLifeCycleStatusCode"`
			InformationLifeCycleStatusCodeText string      `json:"InformationLifeCycleStatusCodeText"`
			Cancel                             bool        `json:"Cancel"`
			RevokeCancellation                 bool        `json:"RevokeCancellation"`
			EntityLastChangedOn                string      `json:"EntityLastChangedOn"`
			PromotionParty struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"PromotionParty"`
		} `json:"results"`
	} `json:"d"`
}