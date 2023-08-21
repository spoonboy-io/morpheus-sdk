# IdentitySourcesAzureADConfigConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **String** |  | [optional] 
**LogoutUrl** | **String** |  | [optional] 
**DoNotIncludeSAMLRequest** | **Boolean** |  | [optional] 
**SAMLSignatureMode** | **String** |  | [optional] 
**DoNotValidateSignature** | **Boolean** |  | [optional] 
**DoNotValidateStatusCode** | **Boolean** |  | [optional] 
**DoNotValidateDestination** | **Boolean** |  | [optional] 
**DoNotValidateIssueInstants** | **Boolean** |  | [optional] 
**DoNotValidateAssertions** | **Boolean** |  | [optional] 
**GivenNameAttribute** | **String** |  | [optional] 
**SurnameAttribute** | **String** |  | [optional] 
**EmailAttribute** | **String** |  | [optional] 
**RequiredAttributeValue** | **String** |  | [optional] 
**RoleAttributeName** | **String** |  | [optional] 
**AzureTenantId** | **String** |  | [optional] 
**AzureAppId** | **String** |  | [optional] 
**AzureAppSecret** | **String** |  | [optional] 
**RoleLinkAttributeName** | **String** |  | [optional] 
**PublicKey** | **String** |  | [optional] 
**AzureAppSecretHash** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$IdentitySourcesAzureADConfigConfig = Initialize-PSOpenAPIToolsIdentitySourcesAzureADConfigConfig  -Url null `
 -LogoutUrl null `
 -DoNotIncludeSAMLRequest null `
 -SAMLSignatureMode null `
 -DoNotValidateSignature null `
 -DoNotValidateStatusCode null `
 -DoNotValidateDestination null `
 -DoNotValidateIssueInstants null `
 -DoNotValidateAssertions null `
 -GivenNameAttribute null `
 -SurnameAttribute null `
 -EmailAttribute null `
 -RequiredAttributeValue null `
 -RoleAttributeName null `
 -AzureTenantId null `
 -AzureAppId null `
 -AzureAppSecret null `
 -RoleLinkAttributeName null `
 -PublicKey null `
 -AzureAppSecretHash null
```

- Convert the resource to JSON
```powershell
$IdentitySourcesAzureADConfigConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

