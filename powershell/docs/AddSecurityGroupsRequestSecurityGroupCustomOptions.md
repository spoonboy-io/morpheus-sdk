# AddSecurityGroupsRequestSecurityGroupCustomOptions
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceGroup** | **String** | External ID of Azure Resource Group | [optional] 
**Vpc** | **String** | External ID of Amazon VPC | [optional] 
**ResourcePoolId** | **Int64** | Resource Pool ID (applicable to cloud types Openstack/OpenTelekom/Huawei) | [optional] 

## Examples

- Prepare the resource
```powershell
$AddSecurityGroupsRequestSecurityGroupCustomOptions = Initialize-PSOpenAPIToolsAddSecurityGroupsRequestSecurityGroupCustomOptions  -ResourceGroup demo-lab `
 -Vpc vpc-0ae14e3b1e3bd3197 `
 -ResourcePoolId 3
```

- Convert the resource to JSON
```powershell
$AddSecurityGroupsRequestSecurityGroupCustomOptions | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

