# AddSecurityPackages200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityPackage** | [**SecurityPackage**](SecurityPackage.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddSecurityPackages200Response = Initialize-PSOpenAPIToolsAddSecurityPackages200Response  -SecurityPackage null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddSecurityPackages200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

