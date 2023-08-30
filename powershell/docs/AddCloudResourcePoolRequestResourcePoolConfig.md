# AddCloudResourcePoolRequestResourcePoolConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | **String** | Provide the base CIDR Block to use for this VPC (must be between a /16 and /28 Block) | [optional] 
**Tenancy** | **String** | default or dedicated | [optional] [default to "default"]
**Managers** | **String[]** | Array of manager usernames | [optional] 
**Developers** | **String[]** | Array of developer usernames | [optional] 
**Auditors** | **String[]** | Array of auditor usernames | [optional] 

## Examples

- Prepare the resource
```powershell
$AddCloudResourcePoolRequestResourcePoolConfig = Initialize-PSOpenAPIToolsAddCloudResourcePoolRequestResourcePoolConfig  -CidrBlock null `
 -Tenancy null `
 -Managers null `
 -Developers null `
 -Auditors null
```

- Convert the resource to JSON
```powershell
$AddCloudResourcePoolRequestResourcePoolConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

