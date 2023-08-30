# AddVDIAppsRequestVdiApp
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | VDI App name | 
**Description** | **String** | Description | [optional] 
**IconPath** | **System.IO.FileInfo** | Icon Path. A relative location of an icon image | [optional] 
**LaunchPrefix** | **String** | The RDS App Name Prefix.  Must start with &#39;||&#39; | [optional] 

## Examples

- Prepare the resource
```powershell
$AddVDIAppsRequestVdiApp = Initialize-PSOpenAPIToolsAddVDIAppsRequestVdiApp  -Name null `
 -Description null `
 -IconPath null `
 -LaunchPrefix ||chrome
```

- Convert the resource to JSON
```powershell
$AddVDIAppsRequestVdiApp | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

