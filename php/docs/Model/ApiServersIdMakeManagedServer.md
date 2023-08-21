# # ApiServersIdMakeManagedServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ssh_username** | **string** | SSH username to use when provisioning | [optional]
**ssh_password** | **string** | SSH password to use, if not specified the account public key can be used | [optional]
**server_os** | [**\OpenAPI\Client\Model\ApiServersIdInstallAgentServerServerOs**](ApiServersIdInstallAgentServerServerOs.md) |  | [optional]
**plan** | [**\OpenAPI\Client\Model\ApiServersIdMakeManagedServerPlan**](ApiServersIdMakeManagedServerPlan.md) |  | [optional]
**account** | [**\OpenAPI\Client\Model\ApiServersIdMakeManagedServerAccount**](ApiServersIdMakeManagedServerAccount.md) |  | [optional]
**provision_site_id** | **int** | Specific group to assign the server | [optional]
**tags** | [**\OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value, this adds or updates the specified tags and removes any tags not specified. | [optional]
**config** | [**\OpenAPI\Client\Model\ApiServersIdMakeManagedServerConfig**](ApiServersIdMakeManagedServerConfig.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
