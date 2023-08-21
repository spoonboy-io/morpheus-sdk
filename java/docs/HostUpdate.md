

# HostUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Unique name scoped to your account for the server. |  [optional]
**description** | **String** | Optional description field. |  [optional]
**sshUsername** | **String** | SSH Username |  [optional]
**sshPassword** | **String** | SSH Password |  [optional]
**powerScheduleType** | **Long** | Power schedule ID. |  [optional]
**labels** | **List&lt;String&gt;** |  |  [optional]
**tags** | [**List&lt;ApiServersIdMakeManagedServerTags&gt;**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value. |  [optional]
**addTags** | [**List&lt;ApiServersIdMakeManagedServerTags&gt;**](ApiServersIdMakeManagedServerTags.md) | Add or update value of Metadata tags, Array of objects having a name and value. |  [optional]
**removeTags** | [**List&lt;InstanceUpdateInstanceRemoveTags&gt;**](InstanceUpdateInstanceRemoveTags.md) | Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed. |  [optional]
**guestConsoleType** | **String** | The Type of guest console this server provides such as disabled, vnc, rdp, ssh |  [optional]
**guestConsoleUsername** | **String** | The optional guest console username if you don&#39;t want to use the user defaults |  [optional]
**guestConsolePassword** | **String** | The optional guest console password if not using the accessing users creds |  [optional]
**guestConsolePort** | **String** | The port the guest console is being accessed from |  [optional]
**guestConsolePreferred** | **Boolean** | Can turn off guest console preferences on server in favor of hypervisor console |  [optional]



