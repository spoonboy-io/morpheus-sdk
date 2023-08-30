# ClustersZone
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**ZoneType** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClustersZone = Initialize-PSOpenAPIToolsClustersZone  -Id null `
 -Name null `
 -ZoneType null
```

- Convert the resource to JSON
```powershell
$ClustersZone | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

