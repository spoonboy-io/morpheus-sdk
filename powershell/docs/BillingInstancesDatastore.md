# BillingInstancesDatastore
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Type** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**InternalId** | **String** |  | [optional] 
**ExternalPath** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingInstancesDatastore = Initialize-PSOpenAPIToolsBillingInstancesDatastore  -Id null `
 -Name null `
 -Type null `
 -ExternalId null `
 -InternalId null `
 -ExternalPath null
```

- Convert the resource to JSON
```powershell
$BillingInstancesDatastore | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

