# ClusterServerCreate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**SystemCollectionsHashtable**](.md) | Key for specific host type configuration  The config parameter is for configuration options that are specific to each Provision Type. The Provision Types api can be used to see which options are available.  | 
**ServerType** | [**ClusterServerCreateServerType**](ClusterServerCreateServerType.md) |  | [optional] 
**Name** | **String** | Name to be used for host(s) created in the cluster | 
**Plan** | [**ClusterServerCreatePlan**](ClusterServerCreatePlan.md) |  | 
**Volumes** | [**ClusterServerCreateVolumes[]**](ClusterServerCreateVolumes.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of Objects | [optional] 
**NetworkInterfaces** | [**ClusterServerCreateNetworkInterfaces[]**](ClusterServerCreateNetworkInterfaces.md) | The networkInterfaces parameter is for network configuration.  The Options API /api/options/zoneNetworkOptions can be used to see which options are available.  It should be passed as an array of Objects with the following attributes  | [optional] 
**SecurityGroups** | **String[]** | Key for security group configuration. | [optional] 
**Visibility** | **String** | Visibility for server host | [optional] [default to "private"]
**UserGroup** | [**ClusterServerCreateUserGroup**](ClusterServerCreateUserGroup.md) |  | [optional] 
**NetworkDomain** | **String** | Network domain | [optional] 
**Hostname** | **String** | Hostname for server host | [optional] 
**NodeCount** | **Int64** | Number of workers or hosts | [optional] 
**Tags** | [**ApiServersIdMakeManagedServerTags[]**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**Labels** | **String[]** | Array of strings (keywords). This will set labels on the server and also on the cluster as well by default. | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterServerCreate = Initialize-PSOpenAPIToolsClusterServerCreate  -Config null `
 -ServerType null `
 -Name null `
 -Plan null `
 -Volumes null `
 -NetworkInterfaces null `
 -SecurityGroups null `
 -Visibility null `
 -UserGroup null `
 -NetworkDomain null `
 -Hostname null `
 -NodeCount null `
 -Tags null `
 -Labels null
```

- Convert the resource to JSON
```powershell
$ClusterServerCreate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

