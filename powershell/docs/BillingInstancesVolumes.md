# BillingInstancesVolumes
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **Int64** |  | [optional] 
**TypeCode** | **String** |  | [optional] 
**Datastore** | [**BillingInstancesDatastore**](BillingInstancesDatastore.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingInstancesVolumes = Initialize-PSOpenAPIToolsBillingInstancesVolumes  -Size null `
 -TypeCode null `
 -Datastore null
```

- Convert the resource to JSON
```powershell
$BillingInstancesVolumes | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

