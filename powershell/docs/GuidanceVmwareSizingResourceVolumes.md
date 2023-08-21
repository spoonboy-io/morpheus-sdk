# GuidanceVmwareSizingResourceVolumes
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**ControllerId** | **Int64** |  | [optional] 
**ControllerMountPoint** | **String** |  | [optional] 
**Resizeable** | **Boolean** |  | [optional] 
**PlanResizable** | **Boolean** |  | [optional] 
**RootVolume** | **Boolean** |  | [optional] 
**UnitNumber** | **String** |  | [optional] 
**TypeId** | **Int64** |  | [optional] 
**ConfigurableIOPS** | **Boolean** |  | [optional] 
**DatastoreId** | **Int64** |  | [optional] 
**MaxStorage** | **Int64** |  | [optional] 
**DisplayOrder** | **Int64** |  | [optional] 
**MaxIOPS** | **String** |  | [optional] 
**Uuid** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceVmwareSizingResourceVolumes = Initialize-PSOpenAPIToolsGuidanceVmwareSizingResourceVolumes  -Id null `
 -Name null `
 -ControllerId null `
 -ControllerMountPoint null `
 -Resizeable null `
 -PlanResizable null `
 -RootVolume null `
 -UnitNumber null `
 -TypeId null `
 -ConfigurableIOPS null `
 -DatastoreId null `
 -MaxStorage null `
 -DisplayOrder null `
 -MaxIOPS null `
 -Uuid null
```

- Convert the resource to JSON
```powershell
$GuidanceVmwareSizingResourceVolumes | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

