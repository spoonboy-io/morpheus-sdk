# ClusterContainers
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Uuid** | **String** |  | [optional] 
**AccountId** | **Int64** |  | [optional] 
**Instance** | **String** |  | [optional] 
**ContainerType** | [**ClusterContainersContainerType**](ClusterContainersContainerType.md) |  | [optional] 
**ContainerTypeSet** | [**ClusterContainersContainerTypeSet**](ClusterContainersContainerTypeSet.md) |  | [optional] 
**Server** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**Cloud** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**Name** | **String** |  | [optional] 
**Ip** | **String** |  | [optional] 
**InternalIp** | **String** |  | [optional] 
**InternalHostname** | **String** |  | [optional] 
**ExternalHostname** | **String** |  | [optional] 
**ExternalDomain** | **String** |  | [optional] 
**ExternalFqdn** | **String** |  | [optional] 
**Ports** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Plan** | [**ClusterContainersPlan**](ClusterContainersPlan.md) |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**StatsEnabled** | **Boolean** |  | [optional] 
**Status** | **String** |  | [optional] 
**UserStatus** | **String** |  | [optional] 
**EnvironmentPrefix** | **String** |  | [optional] 
**ConfigGroup** | **String** |  | [optional] 
**ConfigId** | **String** |  | [optional] 
**ConfigRole** | **String** |  | [optional] 
**Stats** | [**ClusterContainersStats**](ClusterContainersStats.md) |  | [optional] 
**RuntimeInfo** | [**SystemCollectionsHashtable**](.md) |  | [optional] 
**ContainerVersion** | **String** |  | [optional] 
**RepositoryImage** | **String** |  | [optional] 
**PlanCategory** | **String** |  | [optional] 
**Hostname** | **String** |  | [optional] 
**DomainName** | **String** |  | [optional] 
**VolumeCreated** | **Boolean** |  | [optional] 
**ContainerCreated** | **Boolean** |  | [optional] 
**MaxStorage** | **String** |  | [optional] 
**MaxMemory** | **String** |  | [optional] 
**MaxCores** | **String** |  | [optional] 
**MaxCpu** | **String** |  | [optional] 
**HourlyPrice** | **Decimal** |  | [optional] 
**AvailableActions** | [**ClusterContainersAvailableActionsInner[]**](ClusterContainersAvailableActionsInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterContainers = Initialize-PSOpenAPIToolsClusterContainers  -Id null `
 -Uuid null `
 -AccountId null `
 -Instance null `
 -ContainerType null `
 -ContainerTypeSet null `
 -Server null `
 -Cloud null `
 -Name null `
 -Ip null `
 -InternalIp null `
 -InternalHostname null `
 -ExternalHostname null `
 -ExternalDomain null `
 -ExternalFqdn null `
 -Ports null `
 -Plan null `
 -DateCreated null `
 -LastUpdated null `
 -StatsEnabled null `
 -Status null `
 -UserStatus null `
 -EnvironmentPrefix null `
 -ConfigGroup null `
 -ConfigId null `
 -ConfigRole null `
 -Stats null `
 -RuntimeInfo null `
 -ContainerVersion null `
 -RepositoryImage null `
 -PlanCategory null `
 -Hostname null `
 -DomainName null `
 -VolumeCreated null `
 -ContainerCreated null `
 -MaxStorage null `
 -MaxMemory null `
 -MaxCores null `
 -MaxCpu null `
 -HourlyPrice null `
 -AvailableActions null
```

- Convert the resource to JSON
```powershell
$ClusterContainers | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

