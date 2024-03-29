# AddStorageBucketsRequestStorageBucket


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A unique name scoped to your account for the storage bucket | 
**provider_type** | **str** | The type of storage bucket | 
**config** | [**AddStorageBucketsRequestStorageBucketConfig**](AddStorageBucketsRequestStorageBucketConfig.md) |  | 
**default_backup_target** | **bool** | Default Backup Target | [optional]  if omitted the server will use the default value of False
**copy_to_store** | **bool** | Archive Snapshots | [optional] 
**default_deployment_target** | **bool** | Default Deployment Target | [optional]  if omitted the server will use the default value of False
**default_virtual_image_target** | **bool** | Default Virtual Image Store | [optional]  if omitted the server will use the default value of False
**retention_policy_type** | **str** | Cleanup mode. &#x60;backup&#x60; - Move old files to a backup provider. &#x60;delete&#x60; - Delete old files. &#x60;none&#x60; - Keep all files. | [optional]  if omitted the server will use the default value of "none"
**retention_policy_days** | **int** | The number of days old a file must be before it is deleted. | [optional] 
**retention_provider** | **str** | The backup Storage Bucket where old files are moved to. | [optional] 
**bucket_name** | **str** | The name of the bucket. Only applies to &#x60;Amazon&#x60;, &#x60;Azure&#x60;, &#x60;CIFS&#x60;, &#x60;NFSv3&#x60;, &#x60;Openstack Swift&#x60;, and &#x60;Rackspace CDN&#x60;. | [optional] 
**create_bucket** | **bool** | Create the bucket if it does not exist. Only applies to &#x60;Amazon&#x60;, &#x60;Azure&#x60;, &#x60;Openstack Swift&#x60;, and &#x60;Rackspace CDN&#x60;. | [optional]  if omitted the server will use the default value of False
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


