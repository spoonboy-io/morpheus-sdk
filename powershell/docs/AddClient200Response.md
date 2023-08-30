# AddClient200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | [**Client**](Client.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddClient200Response = Initialize-PSOpenAPIToolsAddClient200Response  -Client null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddClient200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

