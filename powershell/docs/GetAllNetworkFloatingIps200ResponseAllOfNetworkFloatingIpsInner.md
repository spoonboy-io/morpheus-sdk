# GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**Cloud** | [**GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInnerCloud**](GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInnerCloud.md) |  | [optional] 
**Server** | [**GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInnerServer**](GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInnerServer.md) |  | [optional] 
**IpStatus** | **String** |  | [optional] 
**IpAddress** | **String** | IP Address | [optional] 
**IpRange** | **String** |  | [optional] 
**PtrId** | **String** |  | [optional] 
**NetworkDomain** | [**GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInnerNetworkDomain**](GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInnerNetworkDomain.md) |  | [optional] 
**CreatedBy** | [**GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInnerCreatedBy**](GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInnerCreatedBy.md) |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInner = Initialize-PSOpenAPIToolsGetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInner  -Id null `
 -ExternalId null `
 -Cloud null `
 -Server null `
 -IpStatus null `
 -IpAddress null `
 -IpRange null `
 -PtrId null `
 -NetworkDomain null `
 -CreatedBy null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$GetAllNetworkFloatingIps200ResponseAllOfNetworkFloatingIpsInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

