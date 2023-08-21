# ApiWikiPagesPage
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** |  | 
**Category** | **String** |  | 
**Content** | **String** |  | 

## Examples

- Prepare the resource
```powershell
$ApiWikiPagesPage = Initialize-PSOpenAPIToolsApiWikiPagesPage  -Name Sample Doc `
 -Category info `
 -Content A sample document in **markdown**.
```

- Convert the resource to JSON
```powershell
$ApiWikiPagesPage | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

