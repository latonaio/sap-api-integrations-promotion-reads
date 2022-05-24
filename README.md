# sap-api-integrations-promotion-reads  
sap-api-integrations-promotion-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP プロモーションデータを取得するマイクロサービスです。  
sap-api-integrations-promotion-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-promotion-reads は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/promotion/overview  

## 動作環境
sap-api-integrations-promotion-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）  

## クラウド環境での利用  
sap-api-integrations-promotion-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-promotion-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/promotion/overview 
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-promotion-reads には、次の API をコールするためのリソースが含まれています。  

* PromotionCollection（プロモーション - プロモーション）※プロモーションの詳細データを取得するために、ToPromotionParty、と合わせて利用されます。
* ToPromotionParty（プロモーション - プロモーション先 ※To）

## API への 値入力条件 の 初期値
sap-api-integrations-promotion-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.PromotionCollection.ID（ID） 


## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"PromotionCollection" が指定されています。    
  
```
	"api_schema": "PromotionCollection",
	"accepter": ["PromotionCollection"],
	"promotion_code": "22",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "PromotionCollection",
	"accepter": ["All"],
	"promotion_code": "22",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetPromotion(iD string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "PromotionCollection":
			func() {
				c.PromotionCollection(iD)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP プロモーション  の プロモーションデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "PromotionParty" は、/SAP_API_Output_Formatter/type.go 内 の Type PromotionCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-promotion-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-promotion-reads/SAP_API_Caller.(*SAPAPICaller).PromotionCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E0C6CDB1ED59EC348C905A296ED",
			"ETag": "2020-04-24T14:24:58+09:00",
			"AccountType": "01",
			"AccountTypeText": "Account",
			"Currency": "",
			"CurrencyText": "",
			"ID": "22",
			"Name": "MKT-Region",
			"Objective": "0002",
			"ObjectiveText": "Brand Awareness",
			"Priority": "3",
			"PriorityText": "Normal",
			"ProductPlanningBasisCode": "005",
			"ProductPlanningBasisCodeText": "Product",
			"PromotionType": "P001",
			"PromotionTypeText": "Price Promotion",
			"ConsistencyStatusCode": "3",
			"ConsistencyStatusCodeText": "Consistent",
			"ExternalStatusCode": "I1004",
			"ExternalStatusCodeText": "Released",
			"LifeCycleStatusCode": "2",
			"LifeCycleStatusCodeText": "Released",
			"CreationDate": "2015-10-24T15:42:39+09:00",
			"LastChangeDate": "2015-10-26T18:54:40+09:00",
			"Tactic": "",
			"TacticText": "",
			"PlannigAccountName": "",
			"PlanningAccountID": "",
			"PlanPeriodStartDate": "",
			"PlanPeriodEndDate": "",
			"BuyingPeriodStartDate": "2014-01-30T09:00:00+09:00",
			"BuyingPeriodEndDate": "2014-01-31T09:00:00+09:00",
			"EmployeeResponsible": "Eddie Smoke",
			"EmployeeResposibleID": "8000000009",
			"SalesUnitOrganisationCentreID": "US1100",
			"SalesUnitOrganisationCentreName": "Sales Unit US",
			"SalesOrganisationID": "US1100",
			"SalesOrganisationName": "Sales Unit US",
			"DistributionChannelCode": "01",
			"DistributionChannelCodeText": "Direct sales",
			"DivisionCode": "00",
			"DivisionCodeText": "Cross Division",
			"SalesTerritoryID": "",
			"SalesTerritoryName": "",
			"ActualPeriodStartDate": "",
			"ActualPeriodEndDate": "",
			"Campaign": "",
			"CampaignDescription": "",
			"TargetGroupID": "",
			"TargetGroupDescription": "",
			"ExternalID": "",
			"InformationLifeCycleStatusCode": "AC",
			"InformationLifeCycleStatusCodeText": "Active",
			"Cancel": false,
			"RevokeCancellation": false,
			"EntityLastChangedOn": "2015-10-26T18:54:40+09:00",
			"PromotionParty": "https://sandbox.api.sap.com/sap/c4c/odata/v1/c4codataapi/PromotionCollection('00163E0C6CDB1ED59EC348C905A296ED')/PromotionParty"
		}
	],
	"time": "2022-05-24T17:37:50+09:00"
}

```