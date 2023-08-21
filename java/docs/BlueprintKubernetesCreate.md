

# BlueprintKubernetesCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A name for the blueprint | 
**image** | **String** | Path to display image. Defaults to an internal Morpheus image. |  [optional]
**type** | [**TypeEnum**](#TypeEnum) | Blueprint Type | 
**labels** | **List&lt;String&gt;** | Array of label strings, can be used for filtering. |  [optional]
**kubernetes** | [**BlueprintKubernetesCreateKubernetes**](BlueprintKubernetesCreateKubernetes.md) |  | 
**config** | [**BlueprintKubernetesCreateConfig**](BlueprintKubernetesCreateConfig.md) |  |  [optional]



## Enum: TypeEnum

Name | Value
---- | -----
KUBERNETES | &quot;kubernetes&quot;



