# ClusterLayout
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**InternalId** | **String** |  | [optional] 
**ServerCount** | **Int64** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**Code** | **String** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**HasAutoScale** | **Boolean** |  | [optional] 
**MemoryRequirement** | **Int64** |  | [optional] 
**ClusterVersion** | **String** |  | [optional] 
**ComputeVersion** | **String** |  | [optional] 
**HasSettings** | **Boolean** |  | [optional] 
**SortOrder** | **Int64** |  | [optional] 
**HasConfig** | **Boolean** |  | [optional] 
**Name** | **String** |  | [optional] 
**Creatable** | **Boolean** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**Description** | **String** |  | [optional] 
**GroupType** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType.md) |  | [optional] 
**Labels** | **String[]** |  | [optional] 
**EnvironmentVariables** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**OptionTypes** | [**OptionType[]**](OptionType.md) |  | [optional] 
**Actions** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**ComputeServers** | [**ClusterLayoutComputeServersInner[]**](ClusterLayoutComputeServersInner.md) |  | [optional] 
**InstallContainerRuntime** | **Boolean** |  | [optional] 
**ProvisionType** | [**GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork**](GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork.md) |  | [optional] 
**SpecTemplates** | [**ClusterLayoutSpecTemplatesInner[]**](ClusterLayoutSpecTemplatesInner.md) |  | [optional] 
**TaskSets** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Type** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterLayout = Initialize-PSOpenAPIToolsClusterLayout  -Id null `
 -InternalId null `
 -ServerCount null `
 -DateCreated null `
 -Code null `
 -LastUpdated null `
 -HasAutoScale null `
 -MemoryRequirement null `
 -ClusterVersion null `
 -ComputeVersion null `
 -HasSettings null `
 -SortOrder null `
 -HasConfig null `
 -Name null `
 -Creatable null `
 -Enabled null `
 -Description null `
 -GroupType null `
 -Labels null `
 -EnvironmentVariables null `
 -OptionTypes null `
 -Actions null `
 -ComputeServers null `
 -InstallContainerRuntime null `
 -ProvisionType null `
 -SpecTemplates null `
 -TaskSets null `
 -Type null
```

- Convert the resource to JSON
```powershell
$ClusterLayout | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

