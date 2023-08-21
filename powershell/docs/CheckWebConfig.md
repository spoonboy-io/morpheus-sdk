# CheckWebConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebMethod** | **String** | HTTP method to use for testing | 
**WebUrl** | **String** | Web URL you wish to use to run a check on | 
**IgnoreSSL** | **Boolean** | Ignore SSL Errors | [optional] [default to $false]
**CheckUser** | **String** | If you want to use HTTP Basic Authentication, populate this field with the username | [optional] 
**CheckPassword** | **String** | If you want to use HTTP basic Authentication, populate this field with the password | [optional] 
**TextCheckOn** | **String** | Set value to &#x60;on&#x60; if you want to turn on text matching | [optional] 
**WebTextMatch** | **String** | Set the string you want to look for in the page source | [optional] 
**TunnelOn** | **String** | Set to on to turn on tunneling | [optional] 
**SshHost** | **String** | Hostname or IP address of the proxy host | [optional] 
**SshPort** | **Int64** | Port for SSH on the proxy host, defaults to 22 | [optional] 
**SshUser** | **String** | SSH user on the proxy host to login as | [optional] 
**SshPassword** | **String** | Password for user, if not using key based authentication | [optional] 

## Examples

- Prepare the resource
```powershell
$CheckWebConfig = Initialize-PSOpenAPIToolsCheckWebConfig  -WebMethod null `
 -WebUrl https://google.com `
 -IgnoreSSL true `
 -CheckUser basicUser `
 -CheckPassword basicPassword `
 -TextCheckOn on `
 -WebTextMatch Login `
 -TunnelOn null `
 -SshHost null `
 -SshPort null `
 -SshUser null `
 -SshPassword null
```

- Convert the resource to JSON
```powershell
$CheckWebConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

