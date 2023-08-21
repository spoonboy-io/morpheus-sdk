# # ContainerTypeCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Node type name |
**labels** | **string[]** |  | [optional]
**short_name** | **string** | The short name is a name with no spaces used for display in your container list. |
**description** | **string** | Node type description | [optional]
**container_version** | **string** | Version of the node type |
**provision_type_code** | **string** | Provision type code, eg. &#x60;amazon&#x60;, etc. |
**scripts** | **int[]** | Array of script IDs. | [optional]
**templates** | **int[]** | Array of file template IDs. | [optional]
**virtual_image_id** | **int** | Virtual image ID | [optional]
**stat_type_code** | **string** | Stat type code.  Varies with node type, see Provision Types (customOptionTypes) for allowed values within selected type. | [optional]
**log_type_code** | **string** | Log type code.  Varies with node type, see Provision Types (customOptionTypes) for allowed values within selected type. | [optional]
**server_type** | **string** | Server type.  Always pass \&quot;vm\&quot;. | [optional]
**container_ports** | [**\OpenAPI\Client\Model\ContainerTypeCreateContainerPorts[]**](ContainerTypeCreateContainerPorts.md) | List of exposed port definitions in the format NAME&#x3D;PORT|PROTOCOL | [optional]
**environment_variables** | [**\OpenAPI\Client\Model\ClusterLayoutCreateEnvironmentVariables[]**](ClusterLayoutCreateEnvironmentVariables.md) | The environmentVariables parameter is array of env objects. | [optional]
**config** | **object** | Config object varies with node type.  If using docker, scvmm, ARM, hyperv, or cloudformation, look up provision type details (customOptionTypes) for information. | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
