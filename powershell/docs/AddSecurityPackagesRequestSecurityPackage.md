# AddSecurityPackagesRequestSecurityPackage
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the security package | 
**Labels** | **String[]** | Array of label strings, can be used for filtering. | [optional] 
**Type** | **String** | Security Package Type Code | [optional] [default to "scap-package"]
**Description** | **String** | A description for the security package | [optional] 
**Url** | **String** | URL to download the security package zip file from | 
**Enabled** | **Boolean** | Can be used to disable the security package | [optional] [default to $true]

## Examples

- Prepare the resource
```powershell
$AddSecurityPackagesRequestSecurityPackage = Initialize-PSOpenAPIToolsAddSecurityPackagesRequestSecurityPackage  -Name Sample Security Package `
 -Labels null `
 -Type null `
 -Description null `
 -Url http://10.0.2.2:8080/public-archives/download/SCAP/scap-security-guide-0.1.51.zip `
 -Enabled null
```

- Convert the resource to JSON
```powershell
$AddSecurityPackagesRequestSecurityPackage | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

