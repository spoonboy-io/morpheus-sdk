# ExpirationPolicyTypeConfiguration
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LifecycleType** | **String** |  | [optional] 
**LifecycleAge** | **String** |  | [optional] 
**LifecycleRenewal** | **String** |  | [optional] 
**LifecycleNotify** | **String** |  | [optional] 
**LifecycleMessage** | **String** |  | [optional] 
**LifecycleAutoRenew** | **Boolean** |  | [optional] 
**LifecycleExtensionsBeforeApproval** | **String** |  | [optional] 
**AccountIntegrationId** | **String** |  | [optional] 
**LifecycleHideFixed** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ExpirationPolicyTypeConfiguration = Initialize-PSOpenAPIToolsExpirationPolicyTypeConfiguration  -LifecycleType null `
 -LifecycleAge null `
 -LifecycleRenewal null `
 -LifecycleNotify null `
 -LifecycleMessage null `
 -LifecycleAutoRenew null `
 -LifecycleExtensionsBeforeApproval null `
 -AccountIntegrationId null `
 -LifecycleHideFixed null
```

- Convert the resource to JSON
```powershell
$ExpirationPolicyTypeConfiguration | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

