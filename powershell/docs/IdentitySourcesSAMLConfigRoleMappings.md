# IdentitySourcesSAMLConfigRoleMappings
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceRoleName** | **String** |  | [optional] 
**SourceRoleFqn** | **String** |  | [optional] 
**MappedRole** | [**IdentitySourcesLDAPConfigDefaultAccountRole**](IdentitySourcesLDAPConfigDefaultAccountRole.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$IdentitySourcesSAMLConfigRoleMappings = Initialize-PSOpenAPIToolsIdentitySourcesSAMLConfigRoleMappings  -SourceRoleName null `
 -SourceRoleFqn null `
 -MappedRole null
```

- Convert the resource to JSON
```powershell
$IdentitySourcesSAMLConfigRoleMappings | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

