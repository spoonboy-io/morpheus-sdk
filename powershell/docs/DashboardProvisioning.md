# DashboardProvisioning
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceCount** | **Int64** |  | [optional] 
**FavoriteInstances** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$DashboardProvisioning = Initialize-PSOpenAPIToolsDashboardProvisioning  -InstanceCount null `
 -FavoriteInstances null
```

- Convert the resource to JSON
```powershell
$DashboardProvisioning | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

