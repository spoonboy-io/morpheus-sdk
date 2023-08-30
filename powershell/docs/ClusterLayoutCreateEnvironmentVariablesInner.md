# ClusterLayoutCreateEnvironmentVariablesInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Environment variable name | 
**Value** | **String** | Sets fixed value for variable | [optional] 
**Masked** | **Boolean** | Can be used to enable / disable masking of variable | [optional] [default to $false]
**Export** | **Boolean** | Can be used to enable / disable export of variable | [optional] [default to $false]

## Examples

- Prepare the resource
```powershell
$ClusterLayoutCreateEnvironmentVariablesInner = Initialize-PSOpenAPIToolsClusterLayoutCreateEnvironmentVariablesInner  -Name null `
 -Value null `
 -Masked null `
 -Export null
```

- Convert the resource to JSON
```powershell
$ClusterLayoutCreateEnvironmentVariablesInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

