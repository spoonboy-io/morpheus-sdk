# BlueprintKubernetesCreateSuccess


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A name for the blueprint | [optional] 
**image** | **str** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**type** | **str** | Blueprint Type | [optional]  if omitted the server will use the default value of "kubernetes"
**kubernetes** | [**BlueprintKubernetesCreateKubernetes**](BlueprintKubernetesCreateKubernetes.md) |  | [optional] 
**config** | [**BlueprintKubernetesCreateConfig**](BlueprintKubernetesCreateConfig.md) |  | [optional] 
**visibility** | **str** | Private or Public Access | [optional]  if omitted the server will use the default value of "private"
**resource_permission** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Resource Permission Block | [optional] 
**owner** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Owner | [optional] 
**tenant** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Tenant | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


