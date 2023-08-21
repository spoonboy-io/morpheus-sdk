# AwsResourcePoolConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | **String** | Provide the base CIDR Block to use for this VPC (must be between a /16 and /28 Block) | [optional] 
**Tenancy** | **String** | default or dedicated | [optional] [default to "default"]

## Examples

- Prepare the resource
```powershell
$AwsResourcePoolConfig = Initialize-PSOpenAPIToolsAwsResourcePoolConfig  -CidrBlock null `
 -Tenancy null
```

- Convert the resource to JSON
```powershell
$AwsResourcePoolConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

