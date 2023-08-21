# ClusterCreate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**OneOfstringobject**](OneOfstringobject.md) |  | 
**Name** | **String** | Name of the cluster to be created | 
**Description** | **String** | Description of the cluster to be created | [optional] 
**Labels** | **String[]** | Array of strings (keywords). This will override labels passed under the &#x60;server&#x60; object. | [optional] 
**Group** | [**ClusterCreateGroup**](ClusterCreateGroup.md) |  | 
**Cloud** | [**ClusterCreateCloud**](ClusterCreateCloud.md) |  | 
**Layout** | [**ClusterCreateLayout**](ClusterCreateLayout.md) |  | 
**Server** | [**ClusterServerCreate**](ClusterServerCreate.md) |  | 

## Examples

- Prepare the resource
```powershell
$ClusterCreate = Initialize-PSOpenAPIToolsClusterCreate  -Type null `
 -Name null `
 -Description null `
 -Labels null `
 -Group null `
 -Cloud null `
 -Layout null `
 -Server null
```

- Convert the resource to JSON
```powershell
$ClusterCreate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

