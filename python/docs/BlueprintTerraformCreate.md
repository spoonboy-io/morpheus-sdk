# BlueprintTerraformCreate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A name for the blueprint | 
**terraform** | [**BlueprintTerraformCreateTerraform**](BlueprintTerraformCreateTerraform.md) |  | 
**type** | **str** | Blueprint Type | defaults to "terraform"
**image** | **str** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**labels** | **[str], none_type** | Array of label strings, can be used for filtering. | [optional] 
**config** | [**BlueprintTerraformCreateConfig**](BlueprintTerraformCreateConfig.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


