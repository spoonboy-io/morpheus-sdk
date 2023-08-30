# Container
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int32** |  | [optional] 
**Uuid** | **String** |  | [optional] 
**AccountId** | **Int32** |  | [optional] 
**Instance** | [**ContainerInstance**](ContainerInstance.md) |  | [optional] 
**ContainerType** | [**ContainerContainerType**](ContainerContainerType.md) |  | [optional] 
**ContainerTypeSet** | [**ContainerContainerTypeSet**](ContainerContainerTypeSet.md) |  | [optional] 
**Server** | [**ContainerInstance**](ContainerInstance.md) |  | [optional] 
**Cloud** | [**ContainerInstance**](ContainerInstance.md) |  | [optional] 
**Name** | **String** |  | [optional] 
**Ip** | **String** |  | [optional] 
**InternalIp** | **String** |  | [optional] 
**InternalHostname** | **String** |  | [optional] 
**ExternalHostname** | **String** |  | [optional] 
**ExternalDomain** | **String** |  | [optional] 
**ExternalFqdn** | **String** |  | [optional] 
**Ports** | [**ContainerPort[]**](ContainerPort.md) |  | [optional] 
**Plan** | [**ContainerPlan**](ContainerPlan.md) |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**StatsEnabled** | **Boolean** |  | [optional] 
**Status** | **String** |  | [optional] 
**UserStatus** | **String** |  | [optional] 
**EnvironmentPrefix** | **String** |  | [optional] 
**Stats** | [**ContainerStats**](ContainerStats.md) |  | [optional] 
**RuntimeInfo** | [**SystemCollectionsHashtable**](.md) |  | [optional] 
**ContainerVersion** | **String** |  | [optional] 
**RepositoryImage** | **String** |  | [optional] 
**PlanCategory** | **String** |  | [optional] 
**Hostname** | **String** |  | [optional] 
**DomainName** | **String** |  | [optional] 
**VolumeCreated** | **Boolean** |  | [optional] 
**ContainerCreated** | **Boolean** |  | [optional] 
**MaxStorage** | **Int32** |  | [optional] 
**MaxMemory** | **Int32** |  | [optional] 
**MaxCores** | **Int32** |  | [optional] 
**MaxCpu** | **Int32** |  | [optional] 
**AvailableActions** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**ConfigGroup** | **String** |  | [optional] 
**ConfigId** | **String** |  | [optional] 
**ConfigRole** | **String** |  | [optional] 
**HourlyCost** | **Double** |  | [optional] 
**HourlyPrice** | **Double** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Container = Initialize-PSOpenAPIToolsContainer  -Id null `
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
 -AvailableActions null `
 -ConfigGroup null `
 -ConfigId null `
 -ConfigRole null `
 -HourlyCost null `
 -HourlyPrice null
```

- Convert the resource to JSON
```powershell
$Container | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

