# ClusterPods
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Category** | **String** |  | [optional] 
**ResourceLevel** | **String** |  | [optional] 
**ResourceType** | **String** |  | [optional] 
**Managed** | **Boolean** |  | [optional] 
**Status** | **String** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**Owner** | [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) |  | [optional] 
**TotalCpuUsage** | **Int64** |  | [optional] 
**Stats** | [**SystemCollectionsHashtable**](.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterPods = Initialize-PSOpenAPIToolsClusterPods  -Id null `
 -Name null `
 -Code null `
 -Description null `
 -Category null `
 -ResourceLevel null `
 -ResourceType null `
 -Managed null `
 -Status null `
 -LastUpdated null `
 -Owner null `
 -TotalCpuUsage null `
 -Stats null
```

- Convert the resource to JSON
```powershell
$ClusterPods | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

