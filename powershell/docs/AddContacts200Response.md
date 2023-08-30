# AddContacts200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | [**Contact**](Contact.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddContacts200Response = Initialize-PSOpenAPIToolsAddContacts200Response  -Contact null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddContacts200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

