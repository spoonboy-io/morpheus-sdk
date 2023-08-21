

# ApiServersIdMakeManagedServer

Object containing server configuration parameters
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sshUsername** | **String** | SSH username to use when provisioning |  [optional]
**sshPassword** | **String** | SSH password to use, if not specified the account public key can be used |  [optional]
**serverOs** | [**ApiServersIdInstallAgentServerServerOs**](ApiServersIdInstallAgentServerServerOs.md) |  |  [optional]
**plan** | [**ApiServersIdMakeManagedServerPlan**](ApiServersIdMakeManagedServerPlan.md) |  |  [optional]
**account** | [**ApiServersIdMakeManagedServerAccount**](ApiServersIdMakeManagedServerAccount.md) |  |  [optional]
**provisionSiteId** | **Long** | Specific group to assign the server |  [optional]
**tags** | [**List&lt;ApiServersIdMakeManagedServerTags&gt;**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value, this adds or updates the specified tags and removes any tags not specified. |  [optional]
**config** | [**ApiServersIdMakeManagedServerConfig**](ApiServersIdMakeManagedServerConfig.md) |  |  [optional]



