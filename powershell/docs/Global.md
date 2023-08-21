# Global
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Global** | **Boolean** | Global (All Tenants), load users from all tenants. The default is to only see your own tenant. This is only available to master tenant users with permission to manage tenants and users. | [optional] [default to $false]

## Examples

- Prepare the resource
```powershell
$Global = Initialize-PSOpenAPIToolsGlobal  -Global null
```

- Convert the resource to JSON
```powershell
$Global | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

