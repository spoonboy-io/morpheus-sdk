# ClusterLayoutComputeServers
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**PriorityOrder** | **Int64** |  | [optional] 
**NodeCount** | **Int64** |  | [optional] 
**NodeType** | **String** |  | [optional] 
**MinNodeCount** | **Int64** |  | [optional] 
**MaxNodeCount** | **String** |  | [optional] 
**DynamicCount** | **Boolean** |  | [optional] 
**InstallContainerRuntime** | **Boolean** |  | [optional] 
**InstallStorageRuntime** | **Boolean** |  | [optional] 
**Name** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**Category** | **String** |  | [optional] 
**Config** | **String** |  | [optional] 
**ContainerType** | [**ClusterLayoutContainerType**](ClusterLayoutContainerType.md) |  | [optional] 
**ComputeServerType** | [**ClusterLayoutComputeServerType**](ClusterLayoutComputeServerType.md) |  | [optional] 
**ProvisionService** | **String** |  | [optional] 
**PlanCategory** | **String** |  | [optional] 
**NamePrefix** | **String** |  | [optional] 
**NameSuffix** | **String** |  | [optional] 
**ForceNameIndex** | **Boolean** |  | [optional] 
**LoadBalance** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterLayoutComputeServers = Initialize-PSOpenAPIToolsClusterLayoutComputeServers  -Id null `
 -PriorityOrder null `
 -NodeCount null `
 -NodeType null `
 -MinNodeCount null `
 -MaxNodeCount null `
 -DynamicCount null `
 -InstallContainerRuntime null `
 -InstallStorageRuntime null `
 -Name null `
 -Code null `
 -Category null `
 -Config null `
 -ContainerType null `
 -ComputeServerType null `
 -ProvisionService null `
 -PlanCategory null `
 -NamePrefix null `
 -NameSuffix null `
 -ForceNameIndex null `
 -LoadBalance null
```

- Convert the resource to JSON
```powershell
$ClusterLayoutComputeServers | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

