

# ApiStorageBucketsStorageBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A unique name scoped to your account for the storage bucket | 
**providerType** | [**ProviderTypeEnum**](#ProviderTypeEnum) | The type of storage bucket | 
**defaultBackupTarget** | **Boolean** | Default Backup Target |  [optional]
**copyToStore** | **Boolean** | Archive Snapshots |  [optional]
**defaultDeploymentTarget** | **Boolean** | Default Deployment Target |  [optional]
**defaultVirtualImageTarget** | **Boolean** | Default Virtual Image Store |  [optional]
**retentionPolicyType** | [**RetentionPolicyTypeEnum**](#RetentionPolicyTypeEnum) | Cleanup mode. &#x60;backup&#x60; - Move old files to a backup provider. &#x60;delete&#x60; - Delete old files. &#x60;none&#x60; - Keep all files. |  [optional]
**retentionPolicyDays** | **Long** | The number of days old a file must be before it is deleted. |  [optional]
**retentionProvider** | **String** | The backup Storage Bucket where old files are moved to. |  [optional]
**bucketName** | **String** | The name of the bucket. Only applies to &#x60;Amazon&#x60;, &#x60;Azure&#x60;, &#x60;CIFS&#x60;, &#x60;NFSv3&#x60;, &#x60;Openstack Swift&#x60;, and &#x60;Rackspace CDN&#x60;. |  [optional]
**createBucket** | **Boolean** | Create the bucket if it does not exist. Only applies to &#x60;Amazon&#x60;, &#x60;Azure&#x60;, &#x60;Openstack Swift&#x60;, and &#x60;Rackspace CDN&#x60;. |  [optional]
**config** | [**OneOfobjectobjectobjectobjectobjectobjectobject**](OneOfobjectobjectobjectobjectobjectobjectobject.md) |  | 



## Enum: ProviderTypeEnum

Name | Value
---- | -----
S3 | &quot;s3&quot;
AZURE | &quot;azure&quot;
CIFS | &quot;cifs&quot;
LOCAL | &quot;local&quot;
NFS | &quot;nfs&quot;
OPENSTACK | &quot;openstack&quot;
RACKSPACE | &quot;rackspace&quot;



## Enum: RetentionPolicyTypeEnum

Name | Value
---- | -----
BACKUP | &quot;backup&quot;
DELETE | &quot;delete&quot;
NONE | &quot;none&quot;



