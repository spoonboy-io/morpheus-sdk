# IdentitySourcesCustomSSOConfigConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Noninteractive** | **Boolean** |  | [optional] 
**LoginUrl** | **String** |  | [optional] 
**LogoutUrl** | **String** |  | [optional] 
**EncryptionAlgo** | **String** |  | [optional] 
**EncryptionKey** | **String** |  | [optional] 
**RequiredAttributeValue** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$IdentitySourcesCustomSSOConfigConfig = Initialize-PSOpenAPIToolsIdentitySourcesCustomSSOConfigConfig  -Noninteractive null `
 -LoginUrl null `
 -LogoutUrl null `
 -EncryptionAlgo null `
 -EncryptionKey null `
 -RequiredAttributeValue null
```

- Convert the resource to JSON
```powershell
$IdentitySourcesCustomSSOConfigConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

