# BlueprintKubernetesCreate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A name for the blueprint | 
**kubernetes** | [**BlueprintKubernetesCreateKubernetes**](BlueprintKubernetesCreateKubernetes.md) |  | 
**type** | **str** | Blueprint Type | defaults to "kubernetes"
**image** | **str** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**labels** | **[str], none_type** | Array of label strings, can be used for filtering. | [optional] 
**config** | [**BlueprintKubernetesCreateConfig**](BlueprintKubernetesCreateConfig.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


