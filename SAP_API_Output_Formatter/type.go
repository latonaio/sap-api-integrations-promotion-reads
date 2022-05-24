package sap_api_output_formatter

type Promotion struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	PromotionCode string `json:"promotion_code"`
	Deleted       bool   `json:"deleted"`
}

type PromotionCollection struct {
	ObjectID                           string `json:"ObjectID"`
	ETag                               string `json:"ETag"`
	AccountType                        string `json:"AccountType"`
	AccountTypeText                    string `json:"AccountTypeText"`
	Currency                           string `json:"Currency"`
	CurrencyText                       string `json:"CurrencyText"`
	ID                                 string `json:"ID"`
	Name                               string `json:"Name"`
	Objective                          string `json:"Objective"`
	ObjectiveText                      string `json:"ObjectiveText"`
	Priority                           string `json:"Priority"`
	PriorityText                       string `json:"PriorityText"`
	ProductPlanningBasisCode           string `json:"ProductPlanningBasisCode"`
	ProductPlanningBasisCodeText       string `json:"ProductPlanningBasisCodeText"`
	PromotionType                      string `json:"PromotionType"`
	PromotionTypeText                  string `json:"PromotionTypeText"`
	ConsistencyStatusCode              string `json:"ConsistencyStatusCode"`
	ConsistencyStatusCodeText          string `json:"ConsistencyStatusCodeText"`
	ExternalStatusCode                 string `json:"ExternalStatusCode"`
	ExternalStatusCodeText             string `json:"ExternalStatusCodeText"`
	LifeCycleStatusCode                string `json:"LifeCycleStatusCode"`
	LifeCycleStatusCodeText            string `json:"LifeCycleStatusCodeText"`
	CreationDate                       string `json:"CreationDate"`
	LastChangeDate                     string `json:"LastChangeDate"`
	Tactic                             string `json:"Tactic"`
	TacticText                         string `json:"TacticText"`
	PlannigAccountName                 string `json:"PlannigAccountName"`
	PlanningAccountID                  string `json:"PlanningAccountID"`
	PlanPeriodStartDate                string `json:"PlanPeriodStartDate"`
	PlanPeriodEndDate                  string `json:"PlanPeriodEndDate"`
	BuyingPeriodStartDate              string `json:"BuyingPeriodStartDate"`
	BuyingPeriodEndDate                string `json:"BuyingPeriodEndDate"`
	EmployeeResponsible                string `json:"EmployeeResponsible"`
	EmployeeResposibleID               string `json:"EmployeeResposibleID"`
	SalesUnitOrganisationCentreID      string `json:"SalesUnitOrganisationCentreID"`
	SalesUnitOrganisationCentreName    string `json:"SalesUnitOrganisationCentreName"`
	SalesOrganisationID                string `json:"SalesOrganisationID"`
	SalesOrganisationName              string `json:"SalesOrganisationName"`
	DistributionChannelCode            string `json:"DistributionChannelCode"`
	DistributionChannelCodeText        string `json:"DistributionChannelCodeText"`
	DivisionCode                       string `json:"DivisionCode"`
	DivisionCodeText                   string `json:"DivisionCodeText"`
	SalesTerritoryID                   string `json:"SalesTerritoryID"`
	SalesTerritoryName                 string `json:"SalesTerritoryName"`
	ActualPeriodStartDate              string `json:"ActualPeriodStartDate"`
	ActualPeriodEndDate                string `json:"ActualPeriodEndDate"`
	Campaign                           string `json:"Campaign"`
	CampaignDescription                string `json:"CampaignDescription"`
	TargetGroupID                      string `json:"TargetGroupID"`
	TargetGroupDescription             string `json:"TargetGroupDescription"`
	ExternalID                         string `json:"ExternalID"`
	InformationLifeCycleStatusCode     string `json:"InformationLifeCycleStatusCode"`
	InformationLifeCycleStatusCodeText string `json:"InformationLifeCycleStatusCodeText"`
	Cancel                             bool   `json:"Cancel"`
	RevokeCancellation                 bool   `json:"RevokeCancellation"`
	EntityLastChangedOn                string `json:"EntityLastChangedOn"`
	ToPromotionParty                   string `json:"PromotionParty"`
}

type PromotionParty struct {
	ObjectID       string `json:"ObjectID"`
	ParentObjectID string `json:"ParentObjectID"`
	ETag           string `json:"ETag"`
	ID             string `json:"ID"`
	RoleCode       string `json:"RoleCode"`
	RoleCodeText   string `json:"RoleCodeText"`
	Name           string `json:"Name"`
	MainIndicator  bool   `json:"MainIndicator"`
}
