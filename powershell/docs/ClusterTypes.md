# ClusterTypes
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**DeployTargetService** | **String** |  | [optional] 
**ShortName** | **String** |  | [optional] 
**ProviderType** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**HostService** | **String** |  | [optional] 
**Managed** | **Boolean** |  | [optional] 
**HasMasters** | **Boolean** |  | [optional] 
**HasWorkers** | **Boolean** |  | [optional] 
**ViewSet** | **String** |  | [optional] 
**ImageCode** | **String** |  | [optional] 
**KubeCtlLocal** | **Boolean** |  | [optional] 
**HasDatastore** | **Boolean** |  | [optional] 
**SupportsCloudScaling** | **Boolean** |  | [optional] 
**Name** | **String** |  | [optional] 
**HasDefaultDataDisk** | **Boolean** |  | [optional] 
**CanManage** | **Boolean** |  | [optional] 
**HasCluster** | **Boolean** |  | [optional] 
**Description** | **String** |  | [optional] 
**OptionTypes** | [**OptionType[]**](OptionType.md) |  | [optional] 
**ControllerTypes** | [**ClusterTypesControllerTypesInner[]**](ClusterTypesControllerTypesInner.md) |  | [optional] 
**WorkerTypes** | [**ClusterTypesControllerTypesInner[]**](ClusterTypesControllerTypesInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterTypes = Initialize-PSOpenAPIToolsClusterTypes  -Id null `
 -DeployTargetService null `
 -ShortName null `
 -ProviderType null `
 -Code null `
 -HostService null `
 -Managed null `
 -HasMasters null `
 -HasWorkers null `
 -ViewSet null `
 -ImageCode null `
 -KubeCtlLocal null `
 -HasDatastore null `
 -SupportsCloudScaling null `
 -Name null `
 -HasDefaultDataDisk null `
 -CanManage null `
 -HasCluster null `
 -Description null `
 -OptionTypes null `
 -ControllerTypes null `
 -WorkerTypes null
```

- Convert the resource to JSON
```powershell
$ClusterTypes | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

