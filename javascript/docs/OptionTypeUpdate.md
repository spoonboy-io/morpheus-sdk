# MorpheusApi.OptionTypeUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | The name of the option type for handy reference | [optional] 
**description** | **String** | Short description of the option type | [optional] 
**labels** | **[String]** | Array of label strings, can be used for filtering. | [optional] 
**fieldName** | **String** | Field Name, the name for user input. This along with fieldContext determines the configuration property name.  The property key for when posting this option type to a JSON POST request | [optional] 
**type** | **String** | Type, the type of input. eg. text, checkbox, select, etc. | [optional] [default to &#39;text&#39;]
**fieldLabel** | **String** | Field Label, the label for user input. | [optional] 
**placeholder** | **String** | Any placeholder text when nothing is yet entered | [optional] 
**verifyPattern** | **String** | Verify Pattern, A regexp string that validates the input, use (?i) to make the matcher case insensitive | [optional] 
**defaultValue** | **String** | The default value if no user entry is specified. This value should be passed to the desired JSON Map if nothing else is entered | [optional] 
**required** | **Boolean** | Is this field entry required for the request | [optional] [default to false]
**exportMeta** | **Boolean** | Export as Tag | [optional] [default to false]
**editable** | **Boolean** | Used primarily on tasks and workflows. Basically wether or not the field can be overridden optionally when the object is run | [optional] [default to false]
**optionList** | [**OptionTypeCreateOptionList**](OptionTypeCreateOptionList.md) |  | [optional] 


