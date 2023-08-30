# Clusters
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**Category** | **String** |  | [optional] 
**Visibility** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Location** | **String** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**ServiceUrl** | **String** |  | [optional] 
**ServiceHost** | **String** |  | [optional] 
**ServicePath** | **String** |  | [optional] 
**ServiceHostname** | **String** |  | [optional] 
**ServicePort** | **Int64** |  | [optional] 
**ServiceUsername** | **String** |  | [optional] 
**ServicePassword** | **String** |  | [optional] 
**ServicePasswordHash** | **String** |  | [optional] 
**ServiceToken** | **String** |  | [optional] 
**ServiceTokenHash** | **String** |  | [optional] 
**ServiceAccess** | **String** |  | [optional] 
**ServiceAccessHash** | **String** |  | [optional] 
**ServiceCert** | **String** |  | [optional] 
**ServiceCertHash** | **String** |  | [optional] 
**ServiceVersion** | **String** |  | [optional] 
**SearchDomains** | **String** |  | [optional] 
**EnableInternalDns** | **Boolean** |  | [optional] 
**InternalId** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**DatacenterId** | **String** |  | [optional] 
**Status** | **String** |  | [optional] 
**StatusDate** | **System.DateTime** |  | [optional] 
**StatusMessage** | **String** |  | [optional] 
**InventoryLevel** | **String** |  | [optional] 
**LastSync** | **System.DateTime** |  | [optional] 
**NextRunDate** | **System.DateTime** |  | [optional] 
**LastSyncDuration** | **Int64** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**Managed** | **Boolean** |  | [optional] 
**Labels** | **String[]** |  | [optional] 
**ServiceEntry** | **String** |  | [optional] 
**CreatedBy** | [**GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInnerCreatedBy**](GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInnerCreatedBy.md) |  | [optional] 
**UserGroup** | **String** |  | [optional] 
**Layout** | [**ClustersLayout**](ClustersLayout.md) |  | [optional] 
**Owner** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**Servers** | [**ClustersServersInner[]**](ClustersServersInner.md) |  | [optional] 
**Accounts** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Integrations** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Site** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**Type** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**Zone** | [**ClustersZone**](ClustersZone.md) |  | [optional] 
**WorkerStats** | [**ClustersWorkerStats**](ClustersWorkerStats.md) |  | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Clusters = Initialize-PSOpenAPIToolsClusters  -Id null `
 -Name null `
 -Code null `
 -Category null `
 -Visibility null `
 -Description null `
 -Location null `
 -Enabled null `
 -ServiceUrl null `
 -ServiceHost null `
 -ServicePath null `
 -ServiceHostname null `
 -ServicePort null `
 -ServiceUsername null `
 -ServicePassword null `
 -ServicePasswordHash null `
 -ServiceToken null `
 -ServiceTokenHash null `
 -ServiceAccess null `
 -ServiceAccessHash null `
 -ServiceCert null `
 -ServiceCertHash null `
 -ServiceVersion null `
 -SearchDomains null `
 -EnableInternalDns null `
 -InternalId null `
 -ExternalId null `
 -DatacenterId null `
 -Status null `
 -StatusDate null `
 -StatusMessage null `
 -InventoryLevel null `
 -LastSync null `
 -NextRunDate null `
 -LastSyncDuration null `
 -DateCreated null `
 -LastUpdated null `
 -Managed null `
 -Labels null `
 -ServiceEntry null `
 -CreatedBy null `
 -UserGroup null `
 -Layout null `
 -Owner null `
 -Servers null `
 -Accounts null `
 -Integrations null `
 -Site null `
 -Type null `
 -Zone null `
 -WorkerStats null `
 -Config null
```

- Convert the resource to JSON
```powershell
$Clusters | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

