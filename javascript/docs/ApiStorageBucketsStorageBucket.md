# MorpheusApi.ApiStorageBucketsStorageBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A unique name scoped to your account for the storage bucket | 
**providerType** | **String** | The type of storage bucket | 
**defaultBackupTarget** | **Boolean** | Default Backup Target | [optional] [default to false]
**copyToStore** | **Boolean** | Archive Snapshots | [optional] 
**defaultDeploymentTarget** | **Boolean** | Default Deployment Target | [optional] [default to false]
**defaultVirtualImageTarget** | **Boolean** | Default Virtual Image Store | [optional] [default to false]
**retentionPolicyType** | **String** | Cleanup mode. &#x60;backup&#x60; - Move old files to a backup provider. &#x60;delete&#x60; - Delete old files. &#x60;none&#x60; - Keep all files. | [optional] [default to &#39;none&#39;]
**retentionPolicyDays** | **Number** | The number of days old a file must be before it is deleted. | [optional] 
**retentionProvider** | **String** | The backup Storage Bucket where old files are moved to. | [optional] 
**bucketName** | **String** | The name of the bucket. Only applies to &#x60;Amazon&#x60;, &#x60;Azure&#x60;, &#x60;CIFS&#x60;, &#x60;NFSv3&#x60;, &#x60;Openstack Swift&#x60;, and &#x60;Rackspace CDN&#x60;. | [optional] 
**createBucket** | **Boolean** | Create the bucket if it does not exist. Only applies to &#x60;Amazon&#x60;, &#x60;Azure&#x60;, &#x60;Openstack Swift&#x60;, and &#x60;Rackspace CDN&#x60;. | [optional] [default to false]
**config** | [**OneOfobjectobjectobjectobjectobjectobjectobject**](OneOfobjectobjectobjectobjectobjectobjectobject.md) |  | 



## Enum: ProviderTypeEnum


* `s3` (value: `"s3"`)

* `azure` (value: `"azure"`)

* `cifs` (value: `"cifs"`)

* `local` (value: `"local"`)

* `nfs` (value: `"nfs"`)

* `openstack` (value: `"openstack"`)

* `rackspace` (value: `"rackspace"`)





## Enum: RetentionPolicyTypeEnum


* `backup` (value: `"backup"`)

* `delete` (value: `"delete"`)

* `none` (value: `"none"`)




