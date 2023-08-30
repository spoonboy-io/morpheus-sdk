# AddStorageBucketsRequestStorageBucketConfigOneOf
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | **String** | Access Key | [optional] 
**SecretKey** | **String** | Secret Key | [optional] 
**Region** | **String** | Optional Amazon region if creating a new bucket | [optional] 
**Endpoint** | **String** | Optional endpoint URL if pointing to an object store other than amazon that mimics the Amazon S3 APIs. | [optional] 

## Examples

- Prepare the resource
```powershell
$AddStorageBucketsRequestStorageBucketConfigOneOf = Initialize-PSOpenAPIToolsAddStorageBucketsRequestStorageBucketConfigOneOf  -AccessKey null `
 -SecretKey null `
 -Region null `
 -Endpoint null
```

- Convert the resource to JSON
```powershell
$AddStorageBucketsRequestStorageBucketConfigOneOf | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

