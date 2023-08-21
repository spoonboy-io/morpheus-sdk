# BillingInstanceContainers
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefType** | **String** |  | [optional] 
**RefUUID** | **String** |  | [optional] 
**RefId** | **Int64** |  | [optional] 
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**Cost** | **Decimal** |  | [optional] 
**Price** | **Decimal** |  | [optional] 
**NumUnits** | **Decimal** |  | [optional] 
**Unit** | **String** |  | [optional] 
**Currency** | **String** |  | [optional] 
**Usages** | [**BillingInstanceUsages[]**](BillingInstanceUsages.md) |  | [optional] 
**NumUsages** | **Int64** |  | [optional] 
**TotalUsages** | **Int64** |  | [optional] 
**HasMoreUsages** | **Boolean** |  | [optional] 
**FoundPricing** | **Boolean** |  | [optional] 
**Name** | **String** |  | [optional] 
**ServerId** | **Int64** |  | [optional] 
**ServerUUID** | **String** |  | [optional] 
**ServerUniqueId** | **String** |  | [optional] 
**ServerExternalId** | **String** |  | [optional] 
**ServerInternalId** | **String** |  | [optional] 
**ResourcePoolId** | **Int64** |  | [optional] 
**ResourcePoolName** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingInstanceContainers = Initialize-PSOpenAPIToolsBillingInstanceContainers  -RefType null `
 -RefUUID null `
 -RefId null `
 -StartDate null `
 -EndDate null `
 -Cost null `
 -Price null `
 -NumUnits null `
 -Unit null `
 -Currency null `
 -Usages null `
 -NumUsages null `
 -TotalUsages null `
 -HasMoreUsages null `
 -FoundPricing null `
 -Name null `
 -ServerId null `
 -ServerUUID null `
 -ServerUniqueId null `
 -ServerExternalId null `
 -ServerInternalId null `
 -ResourcePoolId null `
 -ResourcePoolName null
```

- Convert the resource to JSON
```powershell
$BillingInstanceContainers | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

