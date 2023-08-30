# HostnamePolicyTypeConfiguration
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostNamingType** | **String** |  | [optional] 
**HostNamingPattern** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$HostnamePolicyTypeConfiguration = Initialize-PSOpenAPIToolsHostnamePolicyTypeConfiguration  -HostNamingType null `
 -HostNamingPattern null
```

- Convert the resource to JSON
```powershell
$HostnamePolicyTypeConfiguration | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

