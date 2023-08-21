# # OptionTypeUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | The name of the option type for handy reference | [optional]
**description** | **string** | Short description of the option type | [optional]
**labels** | **string[]** | Array of label strings, can be used for filtering. | [optional]
**field_name** | **string** | Field Name, the name for user input. This along with fieldContext determines the configuration property name.  The property key for when posting this option type to a JSON POST request | [optional]
**type** | **string** | Type, the type of input. eg. text, checkbox, select, etc. | [optional] [default to 'text']
**field_label** | **string** | Field Label, the label for user input. | [optional]
**placeholder** | **string** | Any placeholder text when nothing is yet entered | [optional]
**verify_pattern** | **string** | Verify Pattern, A regexp string that validates the input, use (?i) to make the matcher case insensitive | [optional]
**default_value** | **string** | The default value if no user entry is specified. This value should be passed to the desired JSON Map if nothing else is entered | [optional]
**required** | **bool** | Is this field entry required for the request | [optional] [default to false]
**export_meta** | **bool** | Export as Tag | [optional] [default to false]
**editable** | **bool** | Used primarily on tasks and workflows. Basically wether or not the field can be overridden optionally when the object is run | [optional] [default to false]
**option_list** | [**\OpenAPI\Client\Model\OptionTypeCreateOptionList**](OptionTypeCreateOptionList.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
