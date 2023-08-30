# IdentitySourcesLDAPConfigConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **String** |  | [optional] 
**BindingUsername** | **String** |  | [optional] 
**BindingPassword** | **String** |  | [optional] 
**UserFqnExpression** | **String** |  | [optional] 
**RequiredRoleFqn** | **String** |  | [optional] 
**UsernameAttribute** | **String** |  | [optional] 
**CommonNameAttribute** | **String** |  | [optional] 
**FirstNameAttribute** | **String** |  | [optional] 
**LastNameAttribute** | **String** |  | [optional] 
**EmailAttribute** | **String** |  | [optional] 
**UniqueMemberAttribute** | **String** |  | [optional] 
**MemberOfAttribute** | **String** |  | [optional] 
**BindingPasswordHash** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$IdentitySourcesLDAPConfigConfig = Initialize-PSOpenAPIToolsIdentitySourcesLDAPConfigConfig  -Url null `
 -BindingUsername null `
 -BindingPassword null `
 -UserFqnExpression null `
 -RequiredRoleFqn null `
 -UsernameAttribute null `
 -CommonNameAttribute null `
 -FirstNameAttribute null `
 -LastNameAttribute null `
 -EmailAttribute null `
 -UniqueMemberAttribute null `
 -MemberOfAttribute null `
 -BindingPasswordHash null
```

- Convert the resource to JSON
```powershell
$IdentitySourcesLDAPConfigConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

