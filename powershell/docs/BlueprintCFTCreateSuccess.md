# BlueprintCFTCreateSuccess
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the blueprint | [optional] 
**Image** | **String** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **String** | Blueprint Type | [optional] 
**CloudFormation** | [**BlueprintCFTCreateSuccessCloudFormation**](BlueprintCFTCreateSuccessCloudFormation.md) |  | [optional] 
**Visibility** | **String** | Private or Public Access | [optional] [default to "private"]
**ResourcePermission** | [**SystemCollectionsHashtable**](.md) | Resource Permission Block | [optional] 
**Owner** | [**SystemCollectionsHashtable**](.md) | Owner | [optional] 
**Tenant** | [**SystemCollectionsHashtable**](.md) | Tenant | [optional] 

## Examples

- Prepare the resource
```powershell
$BlueprintCFTCreateSuccess = Initialize-PSOpenAPIToolsBlueprintCFTCreateSuccess  -Name null `
 -Image null `
 -Type null `
 -CloudFormation null `
 -Visibility null `
 -ResourcePermission null `
 -Owner null `
 -Tenant null
```

- Convert the resource to JSON
```powershell
$BlueprintCFTCreateSuccess | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

