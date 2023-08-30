# CredentialType
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Code** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**Creatable** | **Boolean** |  | [optional] 
**Editable** | **Boolean** |  | [optional] 
**OptionTypes** | [**OptionType[]**](OptionType.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CredentialType = Initialize-PSOpenAPIToolsCredentialType  -Id null `
 -Code null `
 -Name null `
 -Description null `
 -Enabled null `
 -Creatable null `
 -Editable null `
 -OptionTypes null
```

- Convert the resource to JSON
```powershell
$CredentialType | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

