# BillingInstances
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **Decimal** |  | [optional] 
**Cost** | **Decimal** |  | [optional] 
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**Instances** | [**BillingInstancesInstancesInner[]**](BillingInstancesInstancesInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingInstances = Initialize-PSOpenAPIToolsBillingInstances  -Price null `
 -Cost null `
 -StartDate null `
 -EndDate null `
 -Instances null
```

- Convert the resource to JSON
```powershell
$BillingInstances | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

