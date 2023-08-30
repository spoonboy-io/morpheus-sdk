# ForgotPasswordRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **String** | Specified as &quot;&quot;username&quot;&quot; or &quot;&quot;tenantId\username&quot;&quot; with proper HTML encoding and used in conjunction with &#x60;password&#x60;.  | 

## Examples

- Prepare the resource
```powershell
$ForgotPasswordRequest = Initialize-PSOpenAPIToolsForgotPasswordRequest  -Username null
```

- Convert the resource to JSON
```powershell
$ForgotPasswordRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

