# AddBlueprintRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**image** | **str** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**labels** | **[str], none_type** | Array of label strings, can be used for filtering. | [optional] 
**config** | [**BlueprintTerraformCreateConfig**](BlueprintTerraformCreateConfig.md) |  | [optional] 
**name** | **str** | A name for the blueprint | [optional] 
**type** | **str** | Blueprint Type | [optional]  if omitted the server will use the default value of "terraform"
**arm** | [**BlueprintARMCreateArm**](BlueprintARMCreateArm.md) |  | [optional] 
**cloud_formation** | [**BlueprintCFTCreateCloudFormation**](BlueprintCFTCreateCloudFormation.md) |  | [optional] 
**helm** | [**BlueprintHelmCreateHelm**](BlueprintHelmCreateHelm.md) |  | [optional] 
**kubernetes** | [**BlueprintKubernetesCreateKubernetes**](BlueprintKubernetesCreateKubernetes.md) |  | [optional] 
**tiers** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Tier definitions - Create in UI to view a baseline for object | [optional] 
**terraform** | [**BlueprintTerraformCreateTerraform**](BlueprintTerraformCreateTerraform.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


