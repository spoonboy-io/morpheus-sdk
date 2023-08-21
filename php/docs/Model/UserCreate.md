# # UserCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**first_name** | **string** | The user&#39;s first name (optional) | [optional]
**last_name** | **string** | The user&#39;s last name (optional) | [optional]
**username** | **string** | Username (unique per tenant). |
**email** | **string** | Email address |
**password** | **string** | Password to apply to the user |
**roles** | [**\OpenAPI\Client\Model\ApiBlueprintsIdUpdatePermissionsResourcePermissionSites[]**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of objects with id of the role(s) to assign to the user. |
**receive_notifications** | **bool** | Receive Notifications? | [optional] [default to true]
**linux_username** | **string** | Linux Username, user settings for provisioning | [optional]
**linux_password** | **string** | Linux Password, user settings for provisioning | [optional]
**linux_key_pair_id** | **int** | Linux SSH Key, user settings for provisioning | [optional]
**windows_username** | **string** | Windows Username, user settings for provisioning | [optional]
**windows_password** | **string** | Windows Password, user settings for provisioning | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
