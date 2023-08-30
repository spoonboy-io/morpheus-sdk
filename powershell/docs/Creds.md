# Creds
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Type** | **String** |  | [optional] 
**Types** | **String[]** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Creds = Initialize-PSOpenAPIToolsCreds  -Id null `
 -Name null `
 -Type null `
 -Types null
```

- Convert the resource to JSON
```powershell
$Creds | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

