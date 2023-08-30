# CreateNetworkRouterRequest

The parameters for creating a network router is type dependent. The following lists the common parameters. See get a specific type to list available options for that network router type. Note: when creating a router on NSX v3.0+ some BGP configuration settings require BGP to be disabled during initial creation. The BGP feature can be enabled in a subsequent router update API call. 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**network_router** | [**NetworkRoutersCreate**](NetworkRoutersCreate.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


