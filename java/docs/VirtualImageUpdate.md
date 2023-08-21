

# VirtualImageUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A name for the virtual image |  [optional]
**labels** | **List&lt;String&gt;** | Array of label strings, can be used for filtering. |  [optional]
**imageType** | **String** | Code of image type. eg. vmware, ami, etc. |  [optional]
**storageProvider** | [**VirtualImageCreateStorageProvider**](VirtualImageCreateStorageProvider.md) |  |  [optional]
**isCloudInit** | **Boolean** | Cloud Init Enabled? |  [optional]
**userData** | **String** | Cloud-Init User Data, a bash script |  [optional]
**installAgent** | **Boolean** | Install Agent? |  [optional]
**sshUsername** | **String** | SSH Username |  [optional]
**sshPassword** | **String** | SSH Password |  [optional]
**sshKey** | **String** | SSH Key |  [optional]
**osType** | [**OneOfobjectstring**](OneOfobjectstring.md) | A Map containing the id of the OS Type. This can also be passed as a string (code or name) instead. |  [optional]
**visibility** | **String** | private or public |  [optional]
**accounts** | **List&lt;Long&gt;** |  |  [optional]
**isAutoJoinDomain** | **Boolean** | Auto Join Domain? |  [optional]
**virtioSupported** | **Boolean** | VirtIO Drivers Loaded? |  [optional]
**vmToolsInstalled** | **Boolean** | VM Tools Installed? |  [optional]
**isForceCustomization** | **Boolean** | Force Guest Customization? |  [optional]
**trialVersion** | **Boolean** | Trial Version |  [optional]
**isSysprep** | **Boolean** | Sysprep Enabled? |  [optional]
**config** | [**OneOfobjectobject**](OneOfobjectobject.md) | Map of configuration properties, varies by image type. |  [optional]
**tags** | [**List&lt;VirtualImageCreateTags&gt;**](VirtualImageCreateTags.md) | Metadata tags, Array of objects having a name and value, this adds or updates the specified tags and removes any tags not specified. |  [optional]
**addTags** | [**List&lt;VirtualImageCreateTags&gt;**](VirtualImageCreateTags.md) | Add or update value of Metadata tags, Array of objects having a name and value. |  [optional]
**removeTags** | [**List&lt;VirtualImageUpdateRemoveTags&gt;**](VirtualImageUpdateRemoveTags.md) | Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed. |  [optional]



