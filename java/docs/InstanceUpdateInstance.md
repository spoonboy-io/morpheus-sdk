

# InstanceUpdateInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Unique name scoped to your account for the instance. |  [optional]
**description** | **String** | Optional description field. |  [optional]
**instanceContext** | **String** | Environment |  [optional]
**labels** | **List&lt;String&gt;** | Array of strings (keywords). |  [optional]
**tags** | [**List&lt;ApiServersIdMakeManagedServerTags&gt;**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value. |  [optional]
**addTags** | [**List&lt;ApiServersIdMakeManagedServerTags&gt;**](ApiServersIdMakeManagedServerTags.md) | Add or update value of Metadata tags, Array of objects having a name and value. |  [optional]
**removeTags** | [**List&lt;InstanceUpdateInstanceRemoveTags&gt;**](InstanceUpdateInstanceRemoveTags.md) | Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed. |  [optional]
**powerScheduleType** | **Long** | Power schedule ID. |  [optional]
**site** | [**InstanceUpdateInstanceSite**](InstanceUpdateInstanceSite.md) |  |  [optional]
**ownerId** | **Long** | User ID, can be used to change instance owner. |  [optional]
**displayName** | **String** | Name used in the UI for display |  [optional]



