# CheckWebConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebMethod** | **String** | HTTP method to use for testing | 
**WebUrl** | **String** | Web URL you wish to use to run a check on | 
**IgnoreSSL** | **Boolean** | Ignore SSL Errors | [optional] [default to $false]
**CheckUser** | **String** |  | [optional] 
**CheckPassword** | **String** |  | [optional] 
**TextCheckOn** | **String** |  | [optional] 
**WebTextMatch** | **String** |  | [optional] 
**TunnelOn** | **String** |  | [optional] 
**SshHost** | **String** |  | [optional] 
**SshPort** | **Int64** |  | [optional] 
**SshUser** | **String** |  | [optional] 
**SshPassword** | **String** | Password for user, if not using key based authentication | [optional] 
**CheckPasswordHash** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CheckWebConfig = Initialize-PSOpenAPIToolsCheckWebConfig  -WebMethod null `
 -WebUrl https://google.com `
 -IgnoreSSL true `
 -CheckUser null `
 -CheckPassword null `
 -TextCheckOn null `
 -WebTextMatch null `
 -TunnelOn null `
 -SshHost null `
 -SshPort null `
 -SshUser null `
 -SshPassword null `
 -CheckPasswordHash null
```

- Convert the resource to JSON
```powershell
$CheckWebConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

