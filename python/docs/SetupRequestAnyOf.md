# SetupRequestAnyOf


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**appliance_name** | **str** | Appliance Name. Choose a name for your Morpheus Appliance.  This is stored in the Appliance Settings. | 
**account_name** | **str** | Name of the Master Tenant account being created. | 
**username** | **str** | Username for the System Admin user being created. | 
**email** | **str** | Email for the System Admin user being created. | 
**password** | **str** | Password for the System Admin user being created. | 
**appliance_url** | **str** | Appliance URL. Specify the full URL for your Morpheus Appliance. This is stored in the Appliance Settings. | [optional] 
**first_name** | **str** | First Name for the System Admin user being created. | [optional] 
**last_name** | **str** | Last Name for the System Admin user being created. | [optional] 
**hubmode** | **str** | Hub Mode. Determines if and how the appliance should connect with the Morpheus Hub. The default value (skip) means do not connect with the hub, and you will be installing your license manually. If you login or register with the hub then a Community Edition license will be installed automatically. | [optional]  if omitted the server will use the default value of "skip"
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


