# CheckSocketConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VarHost** | **String** | Hostname or IP address of the socket server | 
**Port** | **String** | TCP port | 
**Send** | **String** | Connection string you might want to send to the service | 
**ResponseMatch** | **String** | Response from the service to match against | 
**CheckUser** | **String** |  | [optional] 
**TextCheckOn** | **String** |  | [optional] 
**CheckPassword** | **String** |  | [optional] 
**WebTextMatch** | **String** |  | [optional] 
**CheckPasswordHash** | **String** |  | [optional] 
**TunnelOn** | **String** | Set to on to turn on tunneling | [optional] [default to "off"]
**SshHost** | **String** | Hostname or IP address of the proxy host | [optional] 
**SshPort** | **Int64** | Port for SSH on the proxy host, defaults to 22 | [optional] 
**SshUser** | **String** | SSH user on the proxy host to login as | [optional] 
**SshPassword** | **String** | Password for user, if not using key based authentication | [optional] 

## Examples

- Prepare the resource
```powershell
$CheckSocketConfig = Initialize-PSOpenAPIToolsCheckSocketConfig  -VarHost test.example.org `
 -Port 3306 `
 -Send blah `
 -ResponseMatch OK `
 -CheckUser null `
 -TextCheckOn null `
 -CheckPassword null `
 -WebTextMatch null `
 -CheckPasswordHash null `
 -TunnelOn null `
 -SshHost null `
 -SshPort null `
 -SshUser null `
 -SshPassword null
```

- Convert the resource to JSON
```powershell
$CheckSocketConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

