# MorpheusApi.AppCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**templateId** | **Number** |  | [optional] 
**blueprintId** | [**AppCreateBlueprintId**](AppCreateBlueprintId.md) |  | 
**name** | **String** | A unique name for the app | 
**description** | **String** | Description | [optional] 
**labels** | **[String]** | Array of label strings, can be used for filtering. | [optional] 
**group** | [**AppCreateGroup**](AppCreateGroup.md) |  | [optional] 
**defaultCloud** | [**AppCreateDefaultCloud**](AppCreateDefaultCloud.md) |  | [optional] 
**environment** | **String** | Environment code (appContext) | [optional] 
**tiers** | **Object** | Configuration of app elements | [optional] 

