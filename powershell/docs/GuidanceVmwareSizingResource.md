# GuidanceVmwareSizingResource
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Uuid** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**InternalId** | **String** |  | [optional] 
**ExternalUniqueId** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**ExternalName** | **String** |  | [optional] 
**Hostname** | **String** |  | [optional] 
**ParentServer** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**AccountId** | **Int64** |  | [optional] 
**Account** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**Owner** | [**GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInnerCreatedBy**](GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInnerCreatedBy.md) |  | [optional] 
**Zone** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**Plan** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType.md) |  | [optional] 
**ComputeServerType** | [**ClusterLayoutComputeServersInnerComputeServerType**](ClusterLayoutComputeServersInnerComputeServerType.md) |  | [optional] 
**Visibility** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**ZoneId** | **Int64** |  | [optional] 
**SiteId** | **Int64** |  | [optional] 
**ResourcePoolId** | **Int64** |  | [optional] 
**FolderId** | **Int64** |  | [optional] 
**SshHost** | **String** |  | [optional] 
**SshPort** | **Int64** |  | [optional] 
**ExternalIp** | **String** |  | [optional] 
**InternalIp** | **String** |  | [optional] 
**VolumeId** | **String** |  | [optional] 
**Platform** | **String** |  | [optional] 
**PlatformVersion** | **String** |  | [optional] 
**SshUsername** | **String** |  | [optional] 
**SshPassword** | **String** |  | [optional] 
**SshPasswordHash** | **String** |  | [optional] 
**OsDevice** | **String** |  | [optional] 
**OsType** | **String** |  | [optional] 
**DataDevice** | **String** |  | [optional] 
**LvmEnabled** | **Boolean** |  | [optional] 
**ApiKey** | **String** |  | [optional] 
**SoftwareRaid** | **Boolean** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**Stats** | [**GuidanceVmwareSizingResourceStats**](GuidanceVmwareSizingResourceStats.md) |  | [optional] 
**Status** | **String** |  | [optional] 
**StatusMessage** | **String** |  | [optional] 
**ErrorMessage** | **String** |  | [optional] 
**StatusDate** | **System.DateTime** |  | [optional] 
**StatusPercent** | **String** |  | [optional] 
**StatusEta** | **String** |  | [optional] 
**PowerState** | **String** |  | [optional] 
**AgentInstalled** | **Boolean** |  | [optional] 
**LastAgentUpdate** | **System.DateTime** |  | [optional] 
**AgentVersion** | **String** |  | [optional] 
**MaxCores** | **Int64** |  | [optional] 
**CoresPerSocket** | **Int64** |  | [optional] 
**MaxMemory** | **Int64** |  | [optional] 
**MaxStorage** | **Int64** |  | [optional] 
**MaxCpu** | **String** |  | [optional] 
**HourlyPrice** | **Decimal** |  | [optional] 
**SourceImage** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType.md) |  | [optional] 
**ServerOs** | [**GuidanceVmwareSizingResourceServerOs**](GuidanceVmwareSizingResourceServerOs.md) |  | [optional] 
**Volumes** | [**GuidanceVmwareSizingResourceVolumesInner[]**](GuidanceVmwareSizingResourceVolumesInner.md) |  | [optional] 
**Controllers** | [**GuidanceVmwareSizingResourceControllersInner[]**](GuidanceVmwareSizingResourceControllersInner.md) |  | [optional] 
**Interfaces** | [**GuidanceVmwareSizingResourceInterfacesInner[]**](GuidanceVmwareSizingResourceInterfacesInner.md) |  | [optional] 
**Labels** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Tags** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**TagCompliant** | **String** |  | [optional] 
**Containers** | **Int64[]** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceVmwareSizingResource = Initialize-PSOpenAPIToolsGuidanceVmwareSizingResource  -Id null `
 -Uuid null `
 -ExternalId null `
 -InternalId null `
 -ExternalUniqueId null `
 -Name null `
 -ExternalName null `
 -Hostname null `
 -ParentServer null `
 -AccountId null `
 -Account null `
 -Owner null `
 -Zone null `
 -Plan null `
 -ComputeServerType null `
 -Visibility null `
 -Description null `
 -ZoneId null `
 -SiteId null `
 -ResourcePoolId null `
 -FolderId null `
 -SshHost null `
 -SshPort null `
 -ExternalIp null `
 -InternalIp null `
 -VolumeId null `
 -Platform null `
 -PlatformVersion null `
 -SshUsername null `
 -SshPassword null `
 -SshPasswordHash null `
 -OsDevice null `
 -OsType null `
 -DataDevice null `
 -LvmEnabled null `
 -ApiKey null `
 -SoftwareRaid null `
 -DateCreated null `
 -LastUpdated null `
 -Stats null `
 -Status null `
 -StatusMessage null `
 -ErrorMessage null `
 -StatusDate null `
 -StatusPercent null `
 -StatusEta null `
 -PowerState null `
 -AgentInstalled null `
 -LastAgentUpdate null `
 -AgentVersion null `
 -MaxCores null `
 -CoresPerSocket null `
 -MaxMemory null `
 -MaxStorage null `
 -MaxCpu null `
 -HourlyPrice null `
 -SourceImage null `
 -ServerOs null `
 -Volumes null `
 -Controllers null `
 -Interfaces null `
 -Labels null `
 -Tags null `
 -Enabled null `
 -TagCompliant null `
 -Containers null
```

- Convert the resource to JSON
```powershell
$GuidanceVmwareSizingResource | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

