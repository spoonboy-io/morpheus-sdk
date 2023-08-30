# GuidanceVmwareSizingPlanBeforeActionProvisionType
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**RootDiskCustomizable** | **Boolean** |  | [optional] 
**AddVolumes** | **Boolean** |  | [optional] 
**CustomizeVolume** | **Boolean** |  | [optional] 
**HasConfigurableCpuSockets** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceVmwareSizingPlanBeforeActionProvisionType = Initialize-PSOpenAPIToolsGuidanceVmwareSizingPlanBeforeActionProvisionType  -Id null `
 -Name null `
 -Code null `
 -RootDiskCustomizable null `
 -AddVolumes null `
 -CustomizeVolume null `
 -HasConfigurableCpuSockets null
```

- Convert the resource to JSON
```powershell
$GuidanceVmwareSizingPlanBeforeActionProvisionType | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

