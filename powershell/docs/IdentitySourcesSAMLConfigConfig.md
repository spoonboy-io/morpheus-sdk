# IdentitySourcesSAMLConfigConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleAttributeName** | **String** |  | [optional] 
**RequiredAttributeValue** | **String** |  | [optional] 
**GivenNameAttribute** | **String** |  | [optional] 
**SurnameAttribute** | **String** |  | [optional] 
**LogoutUrl** | **String** |  | [optional] 
**DoNotIncludeSAMLRequest** | **Boolean** |  | [optional] 
**PublicKey** | **String** |  | [optional] 
**EmailAttribute** | **String** |  | [optional] 
**Url** | **String** |  | [optional] 
**DoNotValidateSignature** | **Boolean** |  | [optional] 
**DoNotValidateStatusCode** | **Boolean** |  | [optional] 
**DoNotValidateDestination** | **Boolean** |  | [optional] 
**DoNotValidateIssueInstants** | **Boolean** |  | [optional] 
**DoNotValidateAssertions** | **Boolean** |  | [optional] 
**DoNotValidateAuthStatements** | **Boolean** |  | [optional] 
**DoNotValidateSubject** | **Boolean** |  | [optional] 
**DoNotValidateConditions** | **Boolean** |  | [optional] 
**DoNotValidateAudiences** | **Boolean** |  | [optional] 
**DoNotValidateSubjectRecipients** | **Boolean** |  | [optional] 
**SAMLSignatureMode** | **String** |  | [optional] 
**ForceAuthn** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$IdentitySourcesSAMLConfigConfig = Initialize-PSOpenAPIToolsIdentitySourcesSAMLConfigConfig  -RoleAttributeName null `
 -RequiredAttributeValue null `
 -GivenNameAttribute null `
 -SurnameAttribute null `
 -LogoutUrl null `
 -DoNotIncludeSAMLRequest null `
 -PublicKey null `
 -EmailAttribute null `
 -Url null `
 -DoNotValidateSignature null `
 -DoNotValidateStatusCode null `
 -DoNotValidateDestination null `
 -DoNotValidateIssueInstants null `
 -DoNotValidateAssertions null `
 -DoNotValidateAuthStatements null `
 -DoNotValidateSubject null `
 -DoNotValidateConditions null `
 -DoNotValidateAudiences null `
 -DoNotValidateSubjectRecipients null `
 -SAMLSignatureMode null `
 -ForceAuthn null
```

- Convert the resource to JSON
```powershell
$IdentitySourcesSAMLConfigConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

