# ClusterLayoutUpdate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Cluster layout name | [optional] 
**Description** | **String** | Cluster layout description | [optional] 
**Labels** | **String[]** | Array of label strings, can be used for filtering. | [optional] 
**ComputeVersion** | **String** | Version of the cluster layout | [optional] 
**Creatable** | **Boolean** | Can be used to enable / disable the creatability of the cluster layout. | [optional] [default to $true]
**HasAutoScale** | **Boolean** | Can be used to enable / disable the horizontal scaling. | [optional] [default to $false]
**InstallContainerRuntime** | **Boolean** | Install Docker (container runtime). | [optional] [default to $false]
**MemoryRequirement** | **Int64** | Memory requirement in bytes | [optional] 
**GroupType** | [**ClusterLayoutCreateGroupType**](ClusterLayoutCreateGroupType.md) |  | [optional] 
**ProvisionType** | [**AddServicePlansRequestServicePlanProvisionTypeInner**](AddServicePlansRequestServicePlanProvisionTypeInner.md) |  | [optional] 
**OptionTypes** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner[]**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) | Array of cluster layout option types | [optional] 
**TaskSets** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner[]**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) | Array of cluster layout task sets | [optional] 
**EnvironmentVariables** | [**ClusterLayoutCreateEnvironmentVariablesInner[]**](ClusterLayoutCreateEnvironmentVariablesInner.md) | Array of cluster layout env variables | [optional] 
**Masters** | [**ClusterLayoutCreateMastersInner[]**](ClusterLayoutCreateMastersInner.md) | Array of cluster layout master nodes | [optional] 
**Workers** | [**ClusterLayoutCreateMastersInner[]**](ClusterLayoutCreateMastersInner.md) | Array of cluster layout worker nodes | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterLayoutUpdate = Initialize-PSOpenAPIToolsClusterLayoutUpdate  -Name null `
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
$ClusterLayoutUpdate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

