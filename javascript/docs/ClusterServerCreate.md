# MorpheusApi.ClusterServerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**config** | **Object** | Key for specific host type configuration  The config parameter is for configuration options that are specific to each Provision Type. The Provision Types api can be used to see which options are available.  | 
**serverType** | [**ClusterServerCreateServerType**](ClusterServerCreateServerType.md) |  | [optional] 
**name** | **String** | Name to be used for host(s) created in the cluster | 
**plan** | [**ClusterServerCreatePlan**](ClusterServerCreatePlan.md) |  | 
**volumes** | [**[ClusterServerCreateVolumes]**](ClusterServerCreateVolumes.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of Objects | [optional] 
**networkInterfaces** | [**[ClusterServerCreateNetworkInterfaces]**](ClusterServerCreateNetworkInterfaces.md) | The networkInterfaces parameter is for network configuration.  The Options API /api/options/zoneNetworkOptions can be used to see which options are available.  It should be passed as an array of Objects with the following attributes  | [optional] 
**securityGroups** | **[String]** | Key for security group configuration. | [optional] 
**visibility** | **String** | Visibility for server host | [optional] [default to &#39;private&#39;]
**userGroup** | [**ClusterServerCreateUserGroup**](ClusterServerCreateUserGroup.md) |  | [optional] 
**networkDomain** | **String** | Network domain | [optional] 
**hostname** | **String** | Hostname for server host | [optional] 
**nodeCount** | **Number** | Number of workers or hosts | [optional] 
**tags** | [**[ApiServersIdMakeManagedServerTags]**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**labels** | **[String]** | Array of strings (keywords). This will set labels on the server and also on the cluster as well by default. | [optional] 


