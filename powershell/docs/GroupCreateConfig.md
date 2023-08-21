# GroupCreateConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsIntegrationId** | **String** | Optional DNS Integration ID | [optional] 
**ConfigCmdbId** | **String** | Optional CMDB Integration ID | [optional] 
**ConfigCmId** | **String** | Optional Change Management Integration ID | [optional] 
**ServiceRegistryId** | **String** | Optional Service Registry Integration ID | [optional] 
**ConfigManagementId** | **String** | Optional Configuration Management Integration ID | [optional] 
**ConfigCmdbDiscovery** | **Boolean** | Enable or disable CMDB Discovery | [optional] 

## Examples

- Prepare the resource
```powershell
$GroupCreateConfig = Initialize-PSOpenAPIToolsGroupCreateConfig  -DnsIntegrationId null `
 -ConfigCmdbId null `
 -ConfigCmId null `
 -ServiceRegistryId null `
 -ConfigManagementId null `
 -ConfigCmdbDiscovery null
```

- Convert the resource to JSON
```powershell
$GroupCreateConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

