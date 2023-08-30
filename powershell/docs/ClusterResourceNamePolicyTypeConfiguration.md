# ClusterResourceNamePolicyTypeConfiguration
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerNamingType** | **String** |  | [optional] 
**ServerNamingPattern** | **String** |  | [optional] 
**ServerNamingConflict** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterResourceNamePolicyTypeConfiguration = Initialize-PSOpenAPIToolsClusterResourceNamePolicyTypeConfiguration  -ServerNamingType null `
 -ServerNamingPattern null `
 -ServerNamingConflict null
```

- Convert the resource to JSON
```powershell
$ClusterResourceNamePolicyTypeConfiguration | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

