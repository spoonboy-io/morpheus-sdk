# GroupConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsIntegrationId** | **String** |  | [optional] 
**ConfigCmdbId** | **String** |  | [optional] 
**ConfigCmId** | **String** |  | [optional] 
**ServiceRegistryId** | **String** |  | [optional] 
**ConfigManagementId** | **String** |  | [optional] 
**ConfigCmdbDiscovery** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GroupConfig = Initialize-PSOpenAPIToolsGroupConfig  -DnsIntegrationId null `
 -ConfigCmdbId null `
 -ConfigCmId null `
 -ServiceRegistryId null `
 -ConfigManagementId null `
 -ConfigCmdbDiscovery null
```

- Convert the resource to JSON
```powershell
$GroupConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

