# ApiStorageBucketsIdStorageBucket
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A unique name scoped to your account for the storage bucket | [optional] 
**ProviderType** | **String** | The type of storage bucket | [optional] 
**DefaultBackupTarget** | **Boolean** | Default Backup Target | [optional] [default to $false]
**CopyToStore** | **Boolean** | Archive Snapshots | [optional] 
**DefaultDeploymentTarget** | **Boolean** | Default Deployment Target | [optional] [default to $false]
**DefaultVirtualImageTarget** | **Boolean** | Default Virtual Image Store | [optional] [default to $false]
**RetentionPolicyType** | **String** | Cleanup mode. &#x60;backup&#x60; - Move old files to a backup provider. &#x60;delete&#x60; - Delete old files. &#x60;none&#x60; - Keep all files. | [optional] [default to "none"]
**RetentionPolicyDays** | **Int64** | The number of days old a file must be before it is deleted. | [optional] 
**RetentionProvider** | **String** | The backup Storage Bucket where old files are moved to. | [optional] 
**BucketName** | **String** | The name of the bucket. Only applies to &#x60;Amazon&#x60;, &#x60;Azure&#x60;, &#x60;CIFS&#x60;, &#x60;NFSv3&#x60;, &#x60;Openstack Swift&#x60;, and &#x60;Rackspace CDN&#x60;. | [optional] 
**CreateBucket** | **Boolean** | Create the bucket if it does not exist. Only applies to &#x60;Amazon&#x60;, &#x60;Azure&#x60;, &#x60;Openstack Swift&#x60;, and &#x60;Rackspace CDN&#x60;. | [optional] [default to $false]
**Config** | [**OneOfobjectobjectobjectobjectobjectobjectobject**](OneOfobjectobjectobjectobjectobjectobjectobject.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiStorageBucketsIdStorageBucket = Initialize-PSOpenAPIToolsApiStorageBucketsIdStorageBucket  -Name null `
 -ProviderType null `
 -DefaultBackupTarget null `
 -CopyToStore null `
 -DefaultDeploymentTarget null `
 -DefaultVirtualImageTarget null `
 -RetentionPolicyType null `
 -RetentionPolicyDays null `
 -RetentionProvider null `
 -BucketName myBucket `
 -CreateBucket null `
 -Config null
```

- Convert the resource to JSON
```powershell
$ApiStorageBucketsIdStorageBucket | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

