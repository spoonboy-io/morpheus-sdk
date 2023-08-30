# BlueprintCFTCreateSuccessCloudFormation
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **String** | Configuration Type | 
**Json** | **String** | CloudFormation Template in JSON | [optional] 
**Yaml** | **String** | CloudFormation Template in YAML | [optional] 
**Git** | [**BlueprintCFTCreateCloudFormationGit**](BlueprintCFTCreateCloudFormationGit.md) |  | [optional] 
**IAM** | [**BlueprintCFTCreateSuccessCloudFormationIAM**](BlueprintCFTCreateSuccessCloudFormationIAM.md) |  | [optional] 
**CAPABILITYNAMEDIAM** | [**BlueprintCFTCreateSuccessCloudFormationCAPABILITYNAMEDIAM**](BlueprintCFTCreateSuccessCloudFormationCAPABILITYNAMEDIAM.md) |  | [optional] 
**CAPABILITYAUTOEXPAND** | [**BlueprintCFTCreateSuccessCloudFormationCAPABILITYAUTOEXPAND**](BlueprintCFTCreateSuccessCloudFormationCAPABILITYAUTOEXPAND.md) |  | [optional] 
**InstallAgent** | [**BlueprintARMCreateArmInstallAgent**](BlueprintARMCreateArmInstallAgent.md) |  | [optional] 
**CloudInitEnabled** | [**BlueprintARMCreateArmCloudInitEnabled**](BlueprintARMCreateArmCloudInitEnabled.md) |  | [optional] 

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

