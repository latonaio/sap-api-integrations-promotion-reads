package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-promotion-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToPromotionCollection(raw []byte, l *logger.Logger) ([]PromotionCollection, error) {
	pm := &responses.PromotionCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PromotionCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	promotionCollection := make([]PromotionCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		promotionCollection = append(promotionCollection, PromotionCollection{
			ObjectID:                           data.ObjectID,
			ETag:                               data.ETag,
			AccountType:                        data.AccountType,
			AccountTypeText:                    data.AccountTypeText,
			Currency:                           data.Currency,
			CurrencyText:                       data.CurrencyText,
			ID:                                 data.ID,
			Name:                               data.Name,
			Objective:                          data.Objective,
			ObjectiveText:                      data.ObjectiveText,
			Priority:                           data.Priority,
			PriorityText:                       data.PriorityText,
			ProductPlanningBasisCode:           data.ProductPlanningBasisCode,
			ProductPlanningBasisCodeText:       data.ProductPlanningBasisCodeText,
			PromotionType:                      data.PromotionType,
			PromotionTypeText:                  data.PromotionTypeText,
			ConsistencyStatusCode:              data.ConsistencyStatusCode,
			ConsistencyStatusCodeText:          data.ConsistencyStatusCodeText,
			ExternalStatusCode:                 data.ExternalStatusCode,
			ExternalStatusCodeText:             data.ExternalStatusCodeText,
			LifeCycleStatusCode:                data.LifeCycleStatusCode,
			LifeCycleStatusCodeText:            data.LifeCycleStatusCodeText,
			CreationDate:                       data.CreationDate,
			LastChangeDate:                     data.LastChangeDate,
			Tactic:                             data.Tactic,
			TacticText:                         data.TacticText,
			PlannigAccountName:                 data.PlannigAccountName,
			PlanningAccountID:                  data.PlanningAccountID,
			PlanPeriodStartDate:                data.PlanPeriodStartDate,
			PlanPeriodEndDate:                  data.PlanPeriodEndDate,
			BuyingPeriodStartDate:              data.BuyingPeriodStartDate,
			BuyingPeriodEndDate:                data.BuyingPeriodEndDate,
			EmployeeResponsible:                data.EmployeeResponsible,
			EmployeeResposibleID:               data.EmployeeResposibleID,
			SalesUnitOrganisationCentreID:      data.SalesUnitOrganisationCentreID,
			SalesUnitOrganisationCentreName:    data.SalesUnitOrganisationCentreName,
			SalesOrganisationID:                data.SalesOrganisationID,
			SalesOrganisationName:              data.SalesOrganisationName,
			DistributionChannelCode:            data.DistributionChannelCode,
			DistributionChannelCodeText:        data.DistributionChannelCodeText,
			DivisionCode:                       data.DivisionCode,
			DivisionCodeText:                   data.DivisionCodeText,
			SalesTerritoryID:                   data.SalesTerritoryID,
			SalesTerritoryName:                 data.SalesTerritoryName,
			ActualPeriodStartDate:              data.ActualPeriodStartDate,
			ActualPeriodEndDate:                data.ActualPeriodEndDate,
			Campaign:                           data.Campaign,
			CampaignDescription:                data.CampaignDescription,
			TargetGroupID:                      data.TargetGroupID,
			TargetGroupDescription:             data.TargetGroupDescription,
			ExternalID:                         data.ExternalID,
			InformationLifeCycleStatusCode:     data.InformationLifeCycleStatusCode,
			InformationLifeCycleStatusCodeText: data.InformationLifeCycleStatusCodeText,
			Cancel:                             data.Cancel,
			RevokeCancellation:                 data.RevokeCancellation,
			EntityLastChangedOn:                data.EntityLastChangedOn,
			ToPromotionParty:                   data.PromotionParty.Deferred.URI,
		})
	}

	return promotionCollection, nil
}

func ConvertToPromotionParty(raw []byte, l *logger.Logger) ([]PromotionParty, error) {
	pm := &responses.ToPromotionParty{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PromotionParty. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	promotionParty := make([]PromotionParty, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		promotionParty = append(promotionParty, PromotionParty{
			ObjectID:       data.ObjectID,
			ParentObjectID: data.ParentObjectID,
			ETag:           data.ETag,
			ID:             data.ID,
			RoleCode:       data.RoleCode,
			RoleCodeText:   data.RoleCodeText,
			Name:           data.Name,
			MainIndicator:  data.MainIndicator,
		})
	}

	return promotionParty, nil
}
