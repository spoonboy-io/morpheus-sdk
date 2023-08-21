# IdentitySourcesSAMLConfigProviderSettings
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **String** |  | [optional] 
**AcsUrl** | **String** |  | [optional] 
**SpMetadata** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$IdentitySourcesSAMLConfigProviderSettings = Initialize-PSOpenAPIToolsIdentitySourcesSAMLConfigProviderSettings  -EntityId null `
 -AcsUrl null `
 -SpMetadata null
```

- Convert the resource to JSON
```powershell
$IdentitySourcesSAMLConfigProviderSettings | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

