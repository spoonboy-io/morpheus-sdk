# CheckSshConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshPort** | **Int64** |  | [optional] 
**CheckUser** | **String** |  | [optional] 
**TunnelOn** | **String** |  | [optional] 
**TextCheckOn** | **String** |  | [optional] 
**CheckPassword** | **String** |  | [optional] 
**SshHost** | **String** |  | [optional] 
**SshUser** | **String** |  | [optional] 
**WebTextMatch** | **String** |  | [optional] 
**CheckPasswordHash** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CheckSshConfig = Initialize-PSOpenAPIToolsCheckSshConfig  -SshPort null `
 -CheckUser null `
 -TunnelOn null `
 -TextCheckOn null `
 -CheckPassword null `
 -SshHost null `
 -SshUser null `
 -WebTextMatch null `
 -CheckPasswordHash null
```

- Convert the resource to JSON
```powershell
$CheckSshConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

