

# AppCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**templateId** | **Long** |  |  [optional]
**blueprintId** | [**OneOflongstring**](OneOflongstring.md) | The ID of the Blueprint. Use \&quot;existing\&quot; to create a blank app. | 
**name** | **String** | A unique name for the app | 
**description** | **String** | Description |  [optional]
**labels** | **List&lt;String&gt;** | Array of label strings, can be used for filtering. |  [optional]
**group** | [**AppCreateGroup**](AppCreateGroup.md) |  |  [optional]
**defaultCloud** | [**AppCreateDefaultCloud**](AppCreateDefaultCloud.md) |  |  [optional]
**environment** | **String** | Environment code (appContext) |  [optional]
**tiers** | **Object** | Configuration of app elements |  [optional]



