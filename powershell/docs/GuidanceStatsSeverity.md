# GuidanceStatsSeverity
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Low** | **Int64** |  | [optional] 
**Info** | **Int64** |  | [optional] 
**Warning** | **Int64** |  | [optional] 
**Critical** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceStatsSeverity = Initialize-PSOpenAPIToolsGuidanceStatsSeverity  -Low null `
 -Info null `
 -Warning null `
 -Critical null
```

- Convert the resource to JSON
```powershell
$GuidanceStatsSeverity | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

