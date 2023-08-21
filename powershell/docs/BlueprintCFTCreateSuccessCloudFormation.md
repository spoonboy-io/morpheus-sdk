# BlueprintCFTCreateSuccessCloudFormation
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **String** | Configuration Type | 
**Json** | **String** | CloudFormation Template in JSON | [optional] 
**Yaml** | **String** | CloudFormation Template in YAML | [optional] 
**Git** | [**BlueprintCFTCreateCloudFormationGit**](BlueprintCFTCreateCloudFormationGit.md) |  | [optional] 
**IAM** | [**OneOfbooleanstring**](OneOfbooleanstring.md) | CloudFormation Attribute CAPABILITY_IAM | [optional] 
**CAPABILITYNAMEDIAM** | [**OneOfbooleanstring**](OneOfbooleanstring.md) | CloudFormation Attribute CAPABILITY_NAMED_IAM | [optional] 
**CAPABILITYAUTOEXPAND** | [**OneOfbooleanstring**](OneOfbooleanstring.md) | CloudFormation Attribute CAPABILITY_AUTO_EXPAND | [optional] 
**InstallAgent** | [**OneOfbooleanstring**](OneOfbooleanstring.md) | Install Morpheus Agent | [optional] 
**CloudInitEnabled** | [**OneOfbooleanstring**](OneOfbooleanstring.md) | Cloud Init Enabled | [optional] 

## Examples

- Prepare the resource
```powershell
$BlueprintCFTCreateSuccessCloudFormation = Initialize-PSOpenAPIToolsBlueprintCFTCreateSuccessCloudFormation  -ConfigType null `
 -Json null `
 -Yaml null `
 -Git null `
 -IAM null `
 -CAPABILITYNAMEDIAM null `
 -CAPABILITYAUTOEXPAND null `
 -InstallAgent null `
 -CloudInitEnabled null
```

- Convert the resource to JSON
```powershell
$BlueprintCFTCreateSuccessCloudFormation | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

