# AddCypherKeyRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttl** | [**AddCypherKeyRequestTtl**](AddCypherKeyRequestTtl.md) |  | [optional] 
**Value** | **String** | The secret value to be stored. This is only parsed if type is passed as &#x60;string&#x60;. | [optional] 

## Examples

- Prepare the resource
```powershell
$AddCypherKeyRequest = Initialize-PSOpenAPIToolsAddCypherKeyRequest  -Ttl null `
 -Value null
```

- Convert the resource to JSON
```powershell
$AddCypherKeyRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

