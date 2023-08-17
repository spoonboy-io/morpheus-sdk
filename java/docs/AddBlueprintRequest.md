

# AddBlueprintRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**name** | **String** | A name for the blueprint |  |
|**image** | **String** | Path to display image. Defaults to an internal Morpheus image. |  [optional] |
|**type** | [**TypeEnum**](#TypeEnum) | Blueprint Type |  |
|**labels** | **List&lt;String&gt;** | Array of label strings, can be used for filtering. |  [optional] |
|**arm** | [**BlueprintARMCreateArm**](BlueprintARMCreateArm.md) |  |  |
|**cloudFormation** | [**BlueprintCFTCreateCloudFormation**](BlueprintCFTCreateCloudFormation.md) |  |  |
|**helm** | [**BlueprintHelmCreateHelm**](BlueprintHelmCreateHelm.md) |  |  |
|**kubernetes** | [**BlueprintKubernetesCreateKubernetes**](BlueprintKubernetesCreateKubernetes.md) |  |  |
|**config** | [**BlueprintTerraformCreateConfig**](BlueprintTerraformCreateConfig.md) |  |  [optional] |
|**tiers** | **Object** | Tier definitions - Create in UI to view a baseline for object |  |
|**terraform** | [**BlueprintTerraformCreateTerraform**](BlueprintTerraformCreateTerraform.md) |  |  |



## Enum: TypeEnum

| Name | Value |
|---- | -----|
| ARM | &quot;arm&quot; |
| CLOUDFORMATION | &quot;cloudFormation&quot; |
| HELM | &quot;helm&quot; |
| KUBERNETES | &quot;kubernetes&quot; |
| MORPHEUS | &quot;morpheus&quot; |
| TERRAFORM | &quot;terraform&quot; |



