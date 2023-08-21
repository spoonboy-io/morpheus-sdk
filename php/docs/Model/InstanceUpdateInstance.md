# # InstanceUpdateInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Unique name scoped to your account for the instance. | [optional]
**description** | **string** | Optional description field. | [optional]
**instance_context** | **string** | Environment | [optional]
**labels** | **string[]** | Array of strings (keywords). | [optional]
**tags** | [**\OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value. | [optional]
**add_tags** | [**\OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]**](ApiServersIdMakeManagedServerTags.md) | Add or update value of Metadata tags, Array of objects having a name and value. | [optional]
**remove_tags** | [**\OpenAPI\Client\Model\InstanceUpdateInstanceRemoveTags[]**](InstanceUpdateInstanceRemoveTags.md) | Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed. | [optional]
**power_schedule_type** | **int** | Power schedule ID. | [optional]
**site** | [**\OpenAPI\Client\Model\InstanceUpdateInstanceSite**](InstanceUpdateInstanceSite.md) |  | [optional]
**owner_id** | **int** | User ID, can be used to change instance owner. | [optional]
**display_name** | **string** | Name used in the UI for display | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
