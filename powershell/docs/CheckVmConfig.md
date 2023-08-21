# CheckVmConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerName** | **String** |  | 
**ExternalId** | **String** |  | [optional] 
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
$CheckVmConfig = Initialize-PSOpenAPIToolsCheckVmConfig  -ContainerName null `
 -ExternalId null `
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
$CheckVmConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

