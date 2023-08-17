# AddBlueprintRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the blueprint | 
**Image** | **String** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **String** | Blueprint Type | 
**Labels** | **String[]** | Array of label strings, can be used for filtering. | [optional] 
**Arm** | [**BlueprintARMCreateArm**](BlueprintARMCreateArm.md) |  | 
**CloudFormation** | [**BlueprintCFTCreateCloudFormation**](BlueprintCFTCreateCloudFormation.md) |  | 
**Helm** | [**BlueprintHelmCreateHelm**](BlueprintHelmCreateHelm.md) |  | 
**Kubernetes** | [**BlueprintKubernetesCreateKubernetes**](BlueprintKubernetesCreateKubernetes.md) |  | 
**Config** | [**BlueprintTerraformCreateConfig**](BlueprintTerraformCreateConfig.md) |  | [optional] 
**Tiers** | [**SystemCollectionsHashtable**](.md) | Tier definitions - Create in UI to view a baseline for object | 
**Terraform** | [**BlueprintTerraformCreateTerraform**](BlueprintTerraformCreateTerraform.md) |  | 

## Examples

- Prepare the resource
```powershell
$AddBlueprintRequest = Initialize-PSOpenAPIToolsAddBlueprintRequest  -Name null `
 -Image null `
 -Type null `
 -Labels null `
 -Arm null `
 -CloudFormation null `
 -Helm null `
 -Kubernetes null `
 -Config null `
 -Tiers null `
 -Terraform null
```

- Convert the resource to JSON
```powershell
$AddBlueprintRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

