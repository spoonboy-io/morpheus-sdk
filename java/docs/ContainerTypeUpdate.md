

# ContainerTypeUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Node type name |  [optional]
**labels** | **List&lt;String&gt;** |  |  [optional]
**shortName** | **String** | The short name is a name with no spaces used for display in your container list. |  [optional]
**description** | **String** | Node type description |  [optional]
**containerVersion** | **String** | Version of the node type |  [optional]
**provisionTypeCode** | **String** | Provision type code, eg. &#x60;amazon&#x60;, etc. |  [optional]
**scripts** | **List&lt;Long&gt;** | Array of script IDs. |  [optional]
**templates** | **List&lt;Long&gt;** | Array of file template IDs. |  [optional]
**virtualImageId** | **Long** | Virtual image ID |  [optional]
**statTypeCode** | **String** | Stat type code.  Varies with node type, see Provision Types (customOptionTypes) for allowed values within selected type. |  [optional]
**logTypeCode** | **String** | Log type code.  Varies with node type, see Provision Types (customOptionTypes) for allowed values within selected type. |  [optional]
**serverType** | **String** | Server type.  Always pass \&quot;vm\&quot;. |  [optional]
**containerPorts** | [**List&lt;ContainerTypeCreateContainerPorts&gt;**](ContainerTypeCreateContainerPorts.md) | List of exposed port definitions in the format NAME&#x3D;PORT|PROTOCOL |  [optional]
**environmentVariables** | [**List&lt;ClusterLayoutCreateEnvironmentVariables&gt;**](ClusterLayoutCreateEnvironmentVariables.md) | The environmentVariables parameter is array of env objects. |  [optional]
**config** | **Object** | Config object varies with node type.  If using docker, scvmm, ARM, hyperv, or cloudformation, look up provision type details (customOptionTypes) for information. |  [optional]



