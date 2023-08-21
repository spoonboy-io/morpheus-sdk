# IdentitySourcesLDAPConfigRoleMappings
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceRoleName** | **String** |  | [optional] 
**SourceRoleFqn** | **String** |  | [optional] 
**MappedRole** | [**IdentitySourcesLDAPConfigDefaultAccountRole**](IdentitySourcesLDAPConfigDefaultAccountRole.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$IdentitySourcesLDAPConfigRoleMappings = Initialize-PSOpenAPIToolsIdentitySourcesLDAPConfigRoleMappings  -SourceRoleName null `
 -SourceRoleFqn null `
 -MappedRole null
```

- Convert the resource to JSON
```powershell
$IdentitySourcesLDAPConfigRoleMappings | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

