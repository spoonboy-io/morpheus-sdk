# BillingInstancesInstances
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **Int64** |  | [optional] 
**InstanceUUID** | **String** |  | [optional] 
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**Name** | **String** |  | [optional] 
**Price** | **Decimal** |  | [optional] 
**Cost** | **Decimal** |  | [optional] 
**Currency** | **String** |  | [optional] 
**Containers** | [**BillingInstancesContainers[]**](BillingInstancesContainers.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingInstancesInstances = Initialize-PSOpenAPIToolsBillingInstancesInstances  -InstanceId null `
 -InstanceUUID null `
 -StartDate null `
 -EndDate null `
 -Name null `
 -Price null `
 -Cost null `
 -Currency null `
 -Containers null
```

- Convert the resource to JSON
```powershell
$BillingInstancesInstances | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

