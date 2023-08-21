# ClusterLayoutCreate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Cluster layout name | 
**Description** | **String** | Cluster layout description | [optional] 
**Labels** | **String[]** |  | [optional] 
**ComputeVersion** | **String** | Version of the cluster layout | 
**Creatable** | **Boolean** | Can be used to enable / disable the creatability of the cluster layout. | [optional] [default to $true]
**HasAutoScale** | **Boolean** | Can be used to enable / disable the horizontal scaling. | [optional] [default to $false]
**InstallContainerRuntime** | **Boolean** | Install Docker (container runtime). | [optional] [default to $false]
**MemoryRequirement** | **Int64** | Memory requirement in bytes | [optional] 
**GroupType** | [**ClusterLayoutCreateGroupType**](ClusterLayoutCreateGroupType.md) |  | 
**ProvisionType** | [**ApiServicePlansServicePlanProvisionType**](ApiServicePlansServicePlanProvisionType.md) |  | 
**OptionTypes** | [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites[]**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of cluster layout option types | [optional] 
**TaskSets** | [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites[]**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of cluster layout task sets | [optional] 
**EnvironmentVariables** | [**ClusterLayoutCreateEnvironmentVariables[]**](ClusterLayoutCreateEnvironmentVariables.md) | Array of cluster layout env variables | [optional] 
**Masters** | [**ClusterLayoutCreateMasters[]**](ClusterLayoutCreateMasters.md) | Array of cluster layout master nodes | [optional] 
**Workers** | [**ClusterLayoutCreateMasters[]**](ClusterLayoutCreateMasters.md) | Array of cluster layout worker nodes | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterLayoutCreate = Initialize-PSOpenAPIToolsClusterLayoutCreate  -Name null `
 -Description null `
 -Labels null `
 -ComputeVersion null `
 -Creatable null `
 -HasAutoScale null `
 -InstallContainerRuntime null `
 -MemoryRequirement null `
 -GroupType null `
 -ProvisionType null `
 -OptionTypes null `
 -TaskSets null `
 -EnvironmentVariables null `
 -Masters null `
 -Workers null
```

- Convert the resource to JSON
```powershell
$ClusterLayoutCreate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
