# GetHost200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | [**Server**](Server.md) |  | [optional] 
**Stats** | [**SystemCollectionsHashtable**](.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetHost200Response = Initialize-PSOpenAPIToolsGetHost200Response  -Server null `
 -Stats null
```

- Convert the resource to JSON
```powershell
$GetHost200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

