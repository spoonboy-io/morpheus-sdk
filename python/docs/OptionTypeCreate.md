# OptionTypeCreate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | The name of the option type for handy reference | 
**description** | **str, none_type** | Short description of the option type | [optional] 
**labels** | **[str], none_type** | Array of label strings, can be used for filtering. | [optional] 
**field_name** | **str** | Field Name, the name for user input. This along with fieldContext determines the configuration property name.  The property key for when posting this option type to a JSON POST request | [optional] 
**type** | **str** | Type, the type of input. eg. text, checkbox, select, etc. | [optional]  if omitted the server will use the default value of "text"
**field_label** | **str** | Field Label, the label for user input. | [optional] 
**placeholder** | **str** | Any placeholder text when nothing is yet entered | [optional] 
**verify_pattern** | **str** | Verify Pattern, A regexp string that validates the input, use (?i) to make the matcher case insensitive | [optional] 
**default_value** | **str** | The default value if no user entry is specified. This value should be passed to the desired JSON Map if nothing else is entered | [optional] 
**required** | **bool** | Is this field entry required for the request | [optional]  if omitted the server will use the default value of False
**export_meta** | **bool** | Export as Tag | [optional]  if omitted the server will use the default value of False
**editable** | **bool** | Used primarily on tasks and workflows. Basically wether or not the field can be overridden optionally when the object is run | [optional]  if omitted the server will use the default value of False
**option_list** | [**OptionTypeCreateOptionList**](OptionTypeCreateOptionList.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


