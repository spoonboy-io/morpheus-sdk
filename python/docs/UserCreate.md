# UserCreate

Payload for creating a new user

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**username** | **str** | Username (unique per tenant). | 
**email** | **str** | Email address | 
**password** | **str** | Password to apply to the user | 
**roles** | [**[UpdateBlueprintPermissionsRequestResourcePermissionSitesInner]**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) | Array of objects with id of the role(s) to assign to the user. | 
**first_name** | **str** | The user&#39;s first name (optional) | [optional] 
**last_name** | **str** | The user&#39;s last name (optional) | [optional] 
**receive_notifications** | **bool** | Receive Notifications? | [optional]  if omitted the server will use the default value of True
**linux_username** | **str** | Linux Username, user settings for provisioning | [optional] 
**linux_password** | **str** | Linux Password, user settings for provisioning | [optional] 
**linux_key_pair_id** | **int** | Linux SSH Key, user settings for provisioning | [optional] 
**windows_username** | **str** | Windows Username, user settings for provisioning | [optional] 
**windows_password** | **str** | Windows Password, user settings for provisioning | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


