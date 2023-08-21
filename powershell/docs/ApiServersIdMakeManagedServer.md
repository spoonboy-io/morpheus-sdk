# ApiServersIdMakeManagedServer
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshUsername** | **String** | SSH username to use when provisioning | [optional] 
**SshPassword** | **String** | SSH password to use, if not specified the account public key can be used | [optional] 
**ServerOs** | [**ApiServersIdInstallAgentServerServerOs**](ApiServersIdInstallAgentServerServerOs.md) |  | [optional] 
**Plan** | [**ApiServersIdMakeManagedServerPlan**](ApiServersIdMakeManagedServerPlan.md) |  | [optional] 
**Account** | [**ApiServersIdMakeManagedServerAccount**](ApiServersIdMakeManagedServerAccount.md) |  | [optional] 
**ProvisionSiteId** | **Int64** | Specific group to assign the server | [optional] 
**Tags** | [**ApiServersIdMakeManagedServerTags[]**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value, this adds or updates the specified tags and removes any tags not specified. | [optional] 
**Config** | [**ApiServersIdMakeManagedServerConfig**](ApiServersIdMakeManagedServerConfig.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiServersIdMakeManagedServer = Initialize-PSOpenAPIToolsApiServersIdMakeManagedServer  -SshUsername null `
 -SshPassword null `
 -ServerOs null `
 -Plan null `
 -Account null `
 -ProvisionSiteId null `
 -Tags null `
 -Config null
```

- Convert the resource to JSON
```powershell
$ApiServersIdMakeManagedServer | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

