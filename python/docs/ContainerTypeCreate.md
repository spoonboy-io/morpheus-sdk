# ContainerTypeCreate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Node type name | 
**short_name** | **str** | The short name is a name with no spaces used for display in your container list. | 
**container_version** | **str** | Version of the node type | 
**provision_type_code** | **str** | Provision type code, eg. &#x60;amazon&#x60;, etc. | 
**labels** | **[str]** |  | [optional] 
**description** | **str** | Node type description | [optional] 
**scripts** | **[int]** | Array of script IDs. | [optional] 
**templates** | **[int]** | Array of file template IDs. | [optional] 
**virtual_image_id** | **int** | Virtual image ID | [optional] 
**stat_type_code** | **str** | Stat type code.  Varies with node type, see Provision Types (customOptionTypes) for allowed values within selected type. | [optional] 
**log_type_code** | **str** | Log type code.  Varies with node type, see Provision Types (customOptionTypes) for allowed values within selected type. | [optional] 
**server_type** | **str** | Server type.  Always pass \&quot;vm\&quot;. | [optional] 
**container_ports** | [**[ContainerTypeCreateContainerPortsInner]**](ContainerTypeCreateContainerPortsInner.md) | List of exposed port definitions in the format NAME&#x3D;PORT|PROTOCOL | [optional] 
**environment_variables** | [**[ClusterLayoutCreateEnvironmentVariablesInner]**](ClusterLayoutCreateEnvironmentVariablesInner.md) | The environmentVariables parameter is array of env objects. | [optional] 
**config** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Config object varies with node type.  If using docker, scvmm, ARM, hyperv, or cloudformation, look up provision type details (customOptionTypes) for information. | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


