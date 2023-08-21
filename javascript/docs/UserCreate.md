# MorpheusApi.UserCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**firstName** | **String** | The user&#39;s first name (optional) | [optional] 
**lastName** | **String** | The user&#39;s last name (optional) | [optional] 
**username** | **String** | Username (unique per tenant). | 
**email** | **String** | Email address | 
**password** | **String** | Password to apply to the user | 
**roles** | [**[ApiBlueprintsIdUpdatePermissionsResourcePermissionSites]**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of objects with id of the role(s) to assign to the user. | 
**receiveNotifications** | **Boolean** | Receive Notifications? | [optional] [default to true]
**linuxUsername** | **String** | Linux Username, user settings for provisioning | [optional] 
**linuxPassword** | **String** | Linux Password, user settings for provisioning | [optional] 
**linuxKeyPairId** | **Number** | Linux SSH Key, user settings for provisioning | [optional] 
**windowsUsername** | **String** | Windows Username, user settings for provisioning | [optional] 
**windowsPassword** | **String** | Windows Password, user settings for provisioning | [optional] 


