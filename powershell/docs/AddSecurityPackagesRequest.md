# AddSecurityPackagesRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityPackage** | [**AddSecurityPackagesRequestSecurityPackage**](AddSecurityPackagesRequestSecurityPackage.md) |  | 

## Examples

- Prepare the resource
```powershell
$AddSecurityPackagesRequest = Initialize-PSOpenAPIToolsAddSecurityPackagesRequest  -SecurityPackage null
```

- Convert the resource to JSON
```powershell
$AddSecurityPackagesRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

