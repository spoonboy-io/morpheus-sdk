# # HostUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Unique name scoped to your account for the server. | [optional]
**description** | **string** | Optional description field. | [optional]
**ssh_username** | **string** | SSH Username | [optional]
**ssh_password** | **string** | SSH Password | [optional]
**power_schedule_type** | **int** | Power schedule ID. | [optional]
**labels** | **string[]** |  | [optional]
**tags** | [**\OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value. | [optional]
**add_tags** | [**\OpenAPI\Client\Model\ApiServersIdMakeManagedServerTags[]**](ApiServersIdMakeManagedServerTags.md) | Add or update value of Metadata tags, Array of objects having a name and value. | [optional]
**remove_tags** | [**\OpenAPI\Client\Model\InstanceUpdateInstanceRemoveTags[]**](InstanceUpdateInstanceRemoveTags.md) | Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed. | [optional]
**guest_console_type** | **string** | The Type of guest console this server provides such as disabled, vnc, rdp, ssh | [optional]
**guest_console_username** | **string** | The optional guest console username if you don&#39;t want to use the user defaults | [optional]
**guest_console_password** | **string** | The optional guest console password if not using the accessing users creds | [optional]
**guest_console_port** | **string** | The port the guest console is being accessed from | [optional]
**guest_console_preferred** | **bool** | Can turn off guest console preferences on server in favor of hypervisor console | [optional] [default to true]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
