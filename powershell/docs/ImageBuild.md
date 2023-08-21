# ImageBuild
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Account** | [**InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**Type** | [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](InlineResponse20079LoadBalancerMonitorLoadBalancerType.md) |  | [optional] 
**Site** | [**InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**Zone** | [**InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**BootScript** | [**ImageBuildsBootScript**](ImageBuildsBootScript.md) |  | [optional] 
**BootCommand** | **String** |  | [optional] 
**PreseedScript** | [**ImageBuildsBootScript**](ImageBuildsBootScript.md) |  | [optional] 
**Scripts** | [**ImageBuildsScripts[]**](ImageBuildsScripts.md) |  | [optional] 
**SshUsername** | **String** |  | [optional] 
**SshPassword** | **String** |  | [optional] 
**StorageProvider** | **String** |  | [optional] 
**BuildOutputName** | **String** |  | [optional] 
**ConversionFormats** | **String** |  | [optional] 
**IsCloudInit** | **Boolean** |  | [optional] 
**VmToolsInstalled** | **Boolean** |  | [optional] 
**KeepResults** | **Int64** |  | [optional] 
**Config** | [**ImageBuildConfig**](ImageBuildConfig.md) |  | [optional] 
**LastResult** | [**ImageBuildLastResult**](ImageBuildLastResult.md) |  | [optional] 
**ExecutionCount** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ImageBuild = Initialize-PSOpenAPIToolsImageBuild  -Id null `
 -Account null `
 -Type null `
 -Site null `
 -Zone null `
 -Name null `
 -Description null `
 -BootScript null `
 -BootCommand null `
 -PreseedScript null `
 -Scripts null `
 -SshUsername null `
 -SshPassword null `
 -StorageProvider null `
 -BuildOutputName null `
 -ConversionFormats null `
 -IsCloudInit null `
 -VmToolsInstalled null `
 -KeepResults null `
 -Config null `
 -LastResult null `
 -ExecutionCount null
```

- Convert the resource to JSON
```powershell
$ImageBuild | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

