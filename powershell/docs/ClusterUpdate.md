# ClusterUpdate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Cluster name | [optional] 
**Description** | **String** | Cluster description | [optional] 
**Labels** | **String[]** | Array of label strings, can be used for filtering. | [optional] 
**Enabled** | **Boolean** | Cluster enabled | [optional] 
**ServiceUrl** | **String** | Cluster API Url | [optional] 
**ServiceToken** | **String** | Cluster API token | [optional] 
**Refresh** | **Boolean** | Queue cluster refresh | [optional] 
**Managed** | **Boolean** | Cluster managed | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterUpdate = Initialize-PSOpenAPIToolsClusterUpdate  -Name null `
 -Description null `
 -Labels null `
 -Enabled null `
 -ServiceUrl null `
 -ServiceToken null `
 -Refresh null `
 -Managed null
```

- Convert the resource to JSON
```powershell
$ClusterUpdate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

