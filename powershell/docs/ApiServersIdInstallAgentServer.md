# ApiServersIdInstallAgentServer
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshUsername** | **String** | SSH username to use when provisioning | [optional] 
**SshPassword** | **String** | SSH password to use, if not specified the account public key can be used | [optional] 
**ServerOs** | [**ApiServersIdInstallAgentServerServerOs**](ApiServersIdInstallAgentServerServerOs.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiServersIdInstallAgentServer = Initialize-PSOpenAPIToolsApiServersIdInstallAgentServer  -SshUsername null `
 -SshPassword null `
 -ServerOs null
```

- Convert the resource to JSON
```powershell
$ApiServersIdInstallAgentServer | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

