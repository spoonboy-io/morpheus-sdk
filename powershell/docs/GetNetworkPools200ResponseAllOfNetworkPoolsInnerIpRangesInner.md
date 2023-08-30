# GetNetworkPools200ResponseAllOfNetworkPoolsInnerIpRangesInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**StartAddress** | **String** |  | [optional] 
**EndAddress** | **String** |  | [optional] 
**InternalId** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**AddressCount** | **Int64** |  | [optional] 
**Active** | **Boolean** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**Cidr** | **String** |  | [optional] 
**CidrIPv6** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkPools200ResponseAllOfNetworkPoolsInnerIpRangesInner = Initialize-PSOpenAPIToolsGetNetworkPools200ResponseAllOfNetworkPoolsInnerIpRangesInner  -Id null `
 -StartAddress null `
 -EndAddress null `
 -InternalId null `
 -ExternalId null `
 -Description null `
 -AddressCount null `
 -Active null `
 -DateCreated null `
 -LastUpdated null `
 -Cidr null `
 -CidrIPv6 null
```

- Convert the resource to JSON
```powershell
$GetNetworkPools200ResponseAllOfNetworkPoolsInnerIpRangesInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

