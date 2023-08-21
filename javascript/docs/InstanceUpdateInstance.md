# MorpheusApi.InstanceUpdateInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Unique name scoped to your account for the instance. | [optional] 
**description** | **String** | Optional description field. | [optional] 
**instanceContext** | **String** | Environment | [optional] 
**labels** | **[String]** | Array of strings (keywords). | [optional] 
**tags** | [**[ApiServersIdMakeManagedServerTags]**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**addTags** | [**[ApiServersIdMakeManagedServerTags]**](ApiServersIdMakeManagedServerTags.md) | Add or update value of Metadata tags, Array of objects having a name and value. | [optional] 
**removeTags** | [**[InstanceUpdateInstanceRemoveTags]**](InstanceUpdateInstanceRemoveTags.md) | Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed. | [optional] 
**powerScheduleType** | **Number** | Power schedule ID. | [optional] 
**site** | [**InstanceUpdateInstanceSite**](InstanceUpdateInstanceSite.md) |  | [optional] 
**ownerId** | **Number** | User ID, can be used to change instance owner. | [optional] 
**displayName** | **String** | Name used in the UI for display | [optional] 


