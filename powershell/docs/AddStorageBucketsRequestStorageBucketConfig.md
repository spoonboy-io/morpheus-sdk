# AddStorageBucketsRequestStorageBucketConfig
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **String** | Access Key | [optional] 
**SecretKey** | **String** | Secret Key | [optional] 
**Region** | **String** | Region | [optional] 
**Endpoint** | **String** | Optional endpoint URL if pointing to an object store other than amazon that mimics the Amazon S3 APIs. | [optional] 
**StorageAccount** | **String** | Storage Account | [optional] 
**StorageKey** | **String** | Storage Key | [optional] 
**VarHost** | **String** | Host | [optional] 
**Username** | **String** | username | [optional] 
**Password** | **String** | password | [optional] 
**BasePath** | **String** | Storage Path | [optional] 
**ExportFolder** | **String** | Export Folder | [optional] 
**ApiKey** | **String** | API Key | [optional] 
**IdentityUrl** | **String** | Identity URL | [optional] 

## Examples

- Prepare the resource
```powershell
$AddStorageBucketsRequestStorageBucketConfig = Initialize-PSOpenAPIToolsAddStorageBucketsRequestStorageBucketConfig  -AccessKey null `
 -SecretKey null `
 -Region null `
 -Endpoint null `
 -StorageAccount null `
 -StorageKey null `
 -VarHost null `
 -Username null `
 -Password null `
 -BasePath null `
 -ExportFolder null `
 -ApiKey null `
 -IdentityUrl null
```

- Convert the resource to JSON
```powershell
$AddStorageBucketsRequestStorageBucketConfig | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

