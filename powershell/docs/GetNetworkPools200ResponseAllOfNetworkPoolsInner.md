# GetNetworkPools200ResponseAllOfNetworkPoolsInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Type** | [**GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork**](GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork.md) |  | [optional] 
**Account** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**Category** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**DisplayName** | **String** |  | [optional] 
**InternalId** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**DnsDomain** | **String** |  | [optional] 
**DnsSearchPath** | **String** |  | [optional] 
**HostPrefix** | **String** |  | [optional] 
**HttpProxy** | **String** |  | [optional] 
**DnsServers** | **String[]** |  | [optional] 
**DnsSuffixList** | **String[]** |  | [optional] 
**DhcpServer** | **Boolean** |  | [optional] 
**DhcpIp** | **String** |  | [optional] 
**Gateway** | **String** |  | [optional] 
**Netmask** | **String** |  | [optional] 
**SubnetAddress** | **String** |  | [optional] 
**IpCount** | **Int64** |  | [optional] 
**FreeCount** | **Int64** |  | [optional] 
**PoolEnabled** | **Boolean** |  | [optional] 
**TftpServer** | **String** |  | [optional] 
**BootFile** | **String** |  | [optional] 
**RefType** | **String** |  | [optional] 
**RefId** | **String** |  | [optional] 
**ParentType** | **String** |  | [optional] 
**ParentId** | **String** |  | [optional] 
**PoolGroup** | **String** |  | [optional] 
**IpRanges** | [**GetNetworkPools200ResponseAllOfNetworkPoolsInnerIpRangesInner[]**](GetNetworkPools200ResponseAllOfNetworkPoolsInnerIpRangesInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkPools200ResponseAllOfNetworkPoolsInner = Initialize-PSOpenAPIToolsGetNetworkPools200ResponseAllOfNetworkPoolsInner  -Id null `
 -Type null `
 -Account null `
 -Category null `
 -Code null `
 -Name null `
 -DisplayName null `
 -InternalId null `
 -ExternalId null `
 -DnsDomain null `
 -DnsSearchPath null `
 -HostPrefix null `
 -HttpProxy null `
 -DnsServers null `
 -DnsSuffixList null `
 -DhcpServer null `
 -DhcpIp null `
 -Gateway null `
 -Netmask null `
 -SubnetAddress null `
 -IpCount null `
 -FreeCount null `
 -PoolEnabled null `
 -TftpServer null `
 -BootFile null `
 -RefType null `
 -RefId null `
 -ParentType null `
 -ParentId null `
 -PoolGroup null `
 -IpRanges null
```

- Convert the resource to JSON
```powershell
$GetNetworkPools200ResponseAllOfNetworkPoolsInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

