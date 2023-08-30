# BlueprintCFTCreateCloudFormation
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **String** | Configuration Type | 
**Json** | **String** | CloudFormation Template in JSON | [optional] 
**Yaml** | **String** | CloudFormation Template in YAML | [optional] 
**Git** | [**BlueprintCFTCreateCloudFormationGit**](BlueprintCFTCreateCloudFormationGit.md) |  | [optional] 
**IAM** | **Boolean** | CloudFormation Attribute CAPABILITY_IAM | [optional] [default to $false]
**CAPABILITYNAMEDIAM** | **Boolean** | CloudFormation Attribute CAPABILITY_NAMED_IAM | [optional] [default to $false]
**CAPABILITYAUTOEXPAND** | **Boolean** | CloudFormation Attribute CAPABILITY_AUTO_EXPAND | [optional] [default to $false]
**InstallAgent** | **Boolean** | Install Morpheus Agent | [optional] [default to $false]
**CloudInitEnabled** | **Boolean** | Cloud Init Enabled | [optional] [default to $false]

## Examples

- Prepare the resource
```powershell
$BlueprintCFTCreateCloudFormation = Initialize-PSOpenAPIToolsBlueprintCFTCreateCloudFormation  -ConfigType null `
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
$BlueprintCFTCreateCloudFormation | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

