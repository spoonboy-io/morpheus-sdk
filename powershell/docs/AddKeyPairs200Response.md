# AddKeyPairs200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**KeyPair**](KeyPair.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddKeyPairs200Response = Initialize-PSOpenAPIToolsAddKeyPairs200Response  -Account null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddKeyPairs200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

