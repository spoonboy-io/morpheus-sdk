# ApiAppsIdWikiPage
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** |  | [optional] 
**Category** | **String** |  | [optional] 
**Content** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiAppsIdWikiPage = Initialize-PSOpenAPIToolsApiAppsIdWikiPage  -Name Sample Doc `
 -Category info `
 -Content A sample document in **markdown**.&#x60;
```

- Convert the resource to JSON
```powershell
$ApiAppsIdWikiPage | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

