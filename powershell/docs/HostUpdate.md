# HostUpdate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Unique name scoped to your account for the server. | [optional] 
**Description** | **String** | Optional description field. | [optional] 
**SshUsername** | **String** | SSH Username | [optional] 
**SshPassword** | **String** | SSH Password | [optional] 
**PowerScheduleType** | **Int64** | Power schedule ID. | [optional] 
**Labels** | **String[]** |  | [optional] 
**Tags** | [**ApiServersIdMakeManagedServerTags[]**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**AddTags** | [**ApiServersIdMakeManagedServerTags[]**](ApiServersIdMakeManagedServerTags.md) | Add or update value of Metadata tags, Array of objects having a name and value. | [optional] 
**RemoveTags** | [**InstanceUpdateInstanceRemoveTags[]**](InstanceUpdateInstanceRemoveTags.md) | Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed. | [optional] 
**GuestConsoleType** | **String** | The Type of guest console this server provides such as disabled, vnc, rdp, ssh | [optional] 
**GuestConsoleUsername** | **String** | The optional guest console username if you don&#39;t want to use the user defaults | [optional] 
**GuestConsolePassword** | **String** | The optional guest console password if not using the accessing users creds | [optional] 
**GuestConsolePort** | **String** | The port the guest console is being accessed from | [optional] 
**GuestConsolePreferred** | **Boolean** | Can turn off guest console preferences on server in favor of hypervisor console | [optional] [default to $true]

## Examples

- Prepare the resource
```powershell
$HostUpdate = Initialize-PSOpenAPIToolsHostUpdate  -Name myserver `
 -Description my new server `
 -SshUsername myuser `
 -SshPassword mypassword `
 -PowerScheduleType null `
 -Labels null `
 -Tags [{&quot;name&quot;:&quot;hello&quot;,&quot;value&quot;:&quot;world&quot;},{&quot;name&quot;:&quot;flash&quot;,&quot;value&quot;:&quot;bang&quot;}] `
 -AddTags [{&quot;name&quot;:&quot;hello&quot;,&quot;value&quot;:&quot;world&quot;},{&quot;name&quot;:&quot;flash&quot;,&quot;value&quot;:&quot;bang&quot;}] `
 -RemoveTags [{&quot;name&quot;:&quot;hello&quot;}] `
 -GuestConsoleType null `
 -GuestConsoleUsername null `
 -GuestConsolePassword null `
 -GuestConsolePort null `
 -GuestConsolePreferred null
```

- Convert the resource to JSON
```powershell
$HostUpdate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

