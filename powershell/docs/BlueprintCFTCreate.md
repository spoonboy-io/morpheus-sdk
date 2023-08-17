# BlueprintCFTCreate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the blueprint | 
**Type** | **String** | Blueprint Type | 
**Labels** | **String[]** | Array of label strings, can be used for filtering. | [optional] 
**CloudFormation** | [**BlueprintCFTCreateCloudFormation**](BlueprintCFTCreateCloudFormation.md) |  | 

## Examples

- Prepare the resource
```powershell
$BlueprintCFTCreate = Initialize-PSOpenAPIToolsBlueprintCFTCreate  -Name null `
 -Type null `
 -Labels null `
 -CloudFormation null
```

- Convert the resource to JSON
```powershell
$BlueprintCFTCreate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

