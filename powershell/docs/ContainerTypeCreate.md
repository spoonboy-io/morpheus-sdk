# ContainerTypeCreate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Node type name | 
**Labels** | **String[]** |  | [optional] 
**ShortName** | **String** | The short name is a name with no spaces used for display in your container list. | 
**Description** | **String** | Node type description | [optional] 
**ContainerVersion** | **String** | Version of the node type | 
**ProvisionTypeCode** | **String** | Provision type code, eg. &#x60;amazon&#x60;, etc. | 
**Scripts** | **Int64[]** | Array of script IDs. | [optional] 
**Templates** | **Int64[]** | Array of file template IDs. | [optional] 
**VirtualImageId** | **Int64** | Virtual image ID | [optional] 
**StatTypeCode** | **String** | Stat type code.  Varies with node type, see Provision Types (customOptionTypes) for allowed values within selected type. | [optional] 
**LogTypeCode** | **String** | Log type code.  Varies with node type, see Provision Types (customOptionTypes) for allowed values within selected type. | [optional] 
**ServerType** | **String** | Server type.  Always pass &quot;&quot;vm&quot;&quot;. | [optional] 
**ContainerPorts** | [**ContainerTypeCreateContainerPorts[]**](ContainerTypeCreateContainerPorts.md) | List of exposed port definitions in the format NAME&#x3D;PORT|PROTOCOL | [optional] 
**EnvironmentVariables** | [**ClusterLayoutCreateEnvironmentVariables[]**](ClusterLayoutCreateEnvironmentVariables.md) | The environmentVariables parameter is array of env objects. | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) | Config object varies with node type.  If using docker, scvmm, ARM, hyperv, or cloudformation, look up provision type details (customOptionTypes) for information. | [optional] 

## Examples

- Prepare the resource
```powershell
$ContainerTypeCreate = Initialize-PSOpenAPIToolsContainerTypeCreate  -Name null `
 -Labels null `
 -ShortName null `
 -Description null `
 -ContainerVersion null `
 -ProvisionTypeCode null `
 -Scripts null `
 -Templates null `
 -VirtualImageId null `
 -StatTypeCode null `
 -LogTypeCode null `
 -ServerType null `
 -ContainerPorts [{&quot;name&quot;:&quot;WEB&quot;,&quot;port&quot;:80,&quot;loadBalanceProtocol&quot;:&quot;HTTP&quot;},{&quot;name&quot;:&quot;SECURE&quot;,&quot;port&quot;:443,&quot;loadBalanceProtocol&quot;:&quot;HTTPS&quot;}] `
 -EnvironmentVariables null `
 -Config null
```

- Convert the resource to JSON
```powershell
$ContainerTypeCreate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

