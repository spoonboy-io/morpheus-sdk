

# ClusterServerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**config** | **Object** | Key for specific host type configuration  The config parameter is for configuration options that are specific to each Provision Type. The Provision Types api can be used to see which options are available.  | 
**serverType** | [**ClusterServerCreateServerType**](ClusterServerCreateServerType.md) |  |  [optional]
**name** | **String** | Name to be used for host(s) created in the cluster | 
**plan** | [**ClusterServerCreatePlan**](ClusterServerCreatePlan.md) |  | 
**volumes** | [**List&lt;ClusterServerCreateVolumes&gt;**](ClusterServerCreateVolumes.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of Objects |  [optional]
**networkInterfaces** | [**List&lt;ClusterServerCreateNetworkInterfaces&gt;**](ClusterServerCreateNetworkInterfaces.md) | The networkInterfaces parameter is for network configuration.  The Options API /api/options/zoneNetworkOptions can be used to see which options are available.  It should be passed as an array of Objects with the following attributes  |  [optional]
**securityGroups** | **List&lt;String&gt;** | Key for security group configuration. |  [optional]
**visibility** | **String** | Visibility for server host |  [optional]
**userGroup** | [**ClusterServerCreateUserGroup**](ClusterServerCreateUserGroup.md) |  |  [optional]
**networkDomain** | **String** | Network domain |  [optional]
**hostname** | **String** | Hostname for server host |  [optional]
**nodeCount** | **Long** | Number of workers or hosts |  [optional]
**tags** | [**List&lt;ApiServersIdMakeManagedServerTags&gt;**](ApiServersIdMakeManagedServerTags.md) | Metadata tags, Array of objects having a name and value. |  [optional]
**labels** | **List&lt;String&gt;** | Array of strings (keywords). This will set labels on the server and also on the cluster as well by default. |  [optional]



