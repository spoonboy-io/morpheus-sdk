# GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**NetworkPoolId** | **Int64** |  | [optional] 
**IpType** | **String** |  | [optional] 
**IpAddress** | **String** |  | [optional] 
**GatewayAddress** | **String** |  | [optional] 
**SubnetMask** | **String** |  | [optional] 
**DnsServer** | **String** |  | [optional] 
**InterfaceName** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Active** | **Boolean** |  | [optional] 
**StaticIp** | **Boolean** |  | [optional] 
**Fqdn** | **String** |  | [optional] 
**DomainName** | **String** |  | [optional] 
**Hostname** | **String** |  | [optional] 
**InternalId** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**PtrId** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**RefType** | **String** |  | [optional] 
**RefId** | **String** |  | [optional] 
**SubRefId** | **String** |  | [optional] 
**NetworkDomain** | **String** |  | [optional] 
**CreatedBy** | [**GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInnerCreatedBy**](GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInnerCreatedBy.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInner = Initialize-PSOpenAPIToolsGetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInner  -Id null `
 -NetworkPoolId null `
 -IpType null `
 -IpAddress null `
 -GatewayAddress null `
 -SubnetMask null `
 -DnsServer null `
 -InterfaceName null `
 -Description null `
 -Active null `
 -StaticIp null `
 -Fqdn null `
 -DomainName null `
 -Hostname null `
 -InternalId null `
 -ExternalId null `
 -PtrId null `
 -DateCreated null `
 -LastUpdated null `
 -StartDate null `
 -EndDate null `
 -RefType null `
 -RefId null `
 -SubRefId null `
 -NetworkDomain null `
 -CreatedBy null
```

- Convert the resource to JSON
```powershell
$GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

