# # VirtualImageUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | A name for the virtual image | [optional]
**labels** | **string[]** | Array of label strings, can be used for filtering. | [optional]
**image_type** | **string** | Code of image type. eg. vmware, ami, etc. | [optional]
**storage_provider** | [**\OpenAPI\Client\Model\VirtualImageCreateStorageProvider**](VirtualImageCreateStorageProvider.md) |  | [optional]
**is_cloud_init** | **bool** | Cloud Init Enabled? | [optional] [default to false]
**user_data** | **string** | Cloud-Init User Data, a bash script | [optional]
**install_agent** | **bool** | Install Agent? | [optional] [default to false]
**ssh_username** | **string** | SSH Username | [optional]
**ssh_password** | **string** | SSH Password | [optional]
**ssh_key** | **string** | SSH Key | [optional]
**os_type** | [**OneOfObjectString**](OneOfObjectString.md) | A Map containing the id of the OS Type. This can also be passed as a string (code or name) instead. | [optional]
**visibility** | **string** | private or public | [optional] [default to 'private']
**accounts** | **int[]** |  | [optional]
**is_auto_join_domain** | **bool** | Auto Join Domain? | [optional] [default to false]
**virtio_supported** | **bool** | VirtIO Drivers Loaded? | [optional] [default to true]
**vm_tools_installed** | **bool** | VM Tools Installed? | [optional] [default to true]
**is_force_customization** | **bool** | Force Guest Customization? | [optional] [default to false]
**trial_version** | **bool** | Trial Version | [optional] [default to false]
**is_sysprep** | **bool** | Sysprep Enabled? | [optional] [default to false]
**config** | [**OneOfObjectObject**](OneOfObjectObject.md) | Map of configuration properties, varies by image type. | [optional]
**tags** | [**\OpenAPI\Client\Model\VirtualImageCreateTags[]**](VirtualImageCreateTags.md) | Metadata tags, Array of objects having a name and value, this adds or updates the specified tags and removes any tags not specified. | [optional]
**add_tags** | [**\OpenAPI\Client\Model\VirtualImageCreateTags[]**](VirtualImageCreateTags.md) | Add or update value of Metadata tags, Array of objects having a name and value. | [optional]
**remove_tags** | [**\OpenAPI\Client\Model\VirtualImageUpdateRemoveTags[]**](VirtualImageUpdateRemoveTags.md) | Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
