# BillingServersVolumes
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | **Int64** |  | [optional] 
**TypeCode** | **String** |  | [optional] 
**Datastore** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingServersVolumes = Initialize-PSOpenAPIToolsBillingServersVolumes  -Size null `
 -TypeCode null `
 -Datastore null
```

- Convert the resource to JSON
```powershell
$BillingServersVolumes | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

