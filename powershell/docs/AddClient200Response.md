# AddClient200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**Client** | [**Client**](Client.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddClient200Response = Initialize-PSOpenAPIToolsAddClient200Response  -Success null `
 -Client null
```

- Convert the resource to JSON
```powershell
$AddClient200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

