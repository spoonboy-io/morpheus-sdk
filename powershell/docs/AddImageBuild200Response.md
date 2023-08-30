# AddImageBuild200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageBuild** | [**ImageBuild**](ImageBuild.md) |  | [optional] 
**ErrorCode** | **String** |  | [optional] 
**Success** | **Boolean** | Success indicator, true when the request succeeded and false when an error occurred | [optional] [default to $true]
**Msg** | **String** | Message containing a description of the result, usually a message about the error that occurred | [optional] 
**Errors** | [**SystemCollectionsHashtable**](.md) | Validation errors, with a key for Object containing error messages for each invalid parameter (key) | [optional] 

## Examples

- Prepare the resource
```powershell
$AddImageBuild200Response = Initialize-PSOpenAPIToolsAddImageBuild200Response  -ImageBuild null `
 -ErrorCode null `
 -Success null `
 -Msg null `
 -Errors null
```

- Convert the resource to JSON
```powershell
$AddImageBuild200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

