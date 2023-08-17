# Budgets
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Account** | [**ApplianceSettingsEnabledZoneTypesInner**](ApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**RefScope** | **String** |  | [optional] 
**RefType** | **String** |  | [optional] 
**RefId** | **Int64** |  | [optional] 
**RefName** | **String** |  | [optional] 
**Period** | **String** |  | [optional] 
**Year** | **String** |  | [optional] 
**ResourceType** | **String** |  | [optional] 
**Timezone** | **String** |  | [optional] 
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**Interval** | **String** |  | [optional] 
**Costs** | **Int64[]** |  | [optional] 
**IsFiscal** | **Boolean** |  | [optional] 
**AverageCost** | **Int64** |  | [optional] 
**TotalCost** | **Int64** |  | [optional] 
**Currency** | **String** |  | [optional] 
**Rollover** | **Boolean** |  | [optional] 
**WarningLimit** | **String** |  | [optional] 
**OverLimit** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**InternalId** | **String** |  | [optional] 
**CreatedById** | **Int64** |  | [optional] 
**CreatedByName** | **String** |  | [optional] 
**UpdatedById** | **String** |  | [optional] 
**UpdatedByName** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Budgets = Initialize-PSOpenAPIToolsBudgets  -Id null `
 -Name null `
 -Description null `
 -Account null `
 -Enabled null `
 -RefScope null `
 -RefType null `
 -RefId null `
 -RefName null `
 -Period null `
 -Year null `
 -ResourceType null `
 -Timezone null `
 -StartDate null `
 -EndDate null `
 -Interval null `
 -Costs null `
 -IsFiscal null `
 -AverageCost null `
 -TotalCost null `
 -Currency null `
 -Rollover null `
 -WarningLimit null `
 -OverLimit null `
 -ExternalId null `
 -InternalId null `
 -CreatedById null `
 -CreatedByName null `
 -UpdatedById null `
 -UpdatedByName null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$Budgets | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

