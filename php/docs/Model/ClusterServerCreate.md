# # ClusterServerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**config** | **object** | Key for specific host type configuration  The config parameter is for configuration options that are specific to each Provision Type. The Provision Types api can be used to see which options are available. |
**server_type** | [**\OpenAPI\Client\Model\ClusterServerCreateServerType**](ClusterServerCreateServerType.md) |  | [optional]
**name** | **string** | Name to be used for host(s) created in the cluster |
**plan** | [**\OpenAPI\Client\Model\ClusterServerCreatePlan**](ClusterServerCreatePlan.md) |  |
**volumes** | [**\OpenAPI\Client\Model\ClusterServerCreateVolumes[]**](ClusterServerCreateVolumes.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of Objects | [optional]
**network_interfaces** | [**\OpenAPI\Client\Model\ClusterServerCreateNetworkInterfaces[]**](ClusterServerCreateNetworkInterfaces.md) | The networkInterfaces parameter is for network configuration.  The Options API /api/options/zoneNetworkOptions can be used to see which options are available.  It should be passed as an array of Objects with the following attributes | [optional]
**security_groups** | **string[]** | Key for security group configuration. | [optional]
**visibility** | **string** | Visibility for server host | [optional] [default to 'private']
**user_group** | [**\OpenAPI\Client\Model\ClusterServerCreateUserGroup**](ClusterServerCreateUserGroup.md) |  | [optional]
**network_domain** | **string** | Network domain | [optional]
**hostname** | **string** | Hostname for server host | [optional]
**node_count** | **int** | Number of workers or hosts | [optional]
**tags** | [**\OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value. | [optional]
**labels** | **string[]** | Array of strings (keywords). This will set labels on the server and also on the cluster as well by default. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
