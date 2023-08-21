# # ApiStorageBucketsStorageBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | A unique name scoped to your account for the storage bucket |
**provider_type** | **string** | The type of storage bucket |
**default_backup_target** | **bool** | Default Backup Target | [optional] [default to false]
**copy_to_store** | **bool** | Archive Snapshots | [optional]
**default_deployment_target** | **bool** | Default Deployment Target | [optional] [default to false]
**default_virtual_image_target** | **bool** | Default Virtual Image Store | [optional] [default to false]
**retention_policy_type** | **string** | Cleanup mode. &#x60;backup&#x60; - Move old files to a backup provider. &#x60;delete&#x60; - Delete old files. &#x60;none&#x60; - Keep all files. | [optional] [default to 'none']
**retention_policy_days** | **int** | The number of days old a file must be before it is deleted. | [optional]
**retention_provider** | **string** | The backup Storage Bucket where old files are moved to. | [optional]
**bucket_name** | **string** | The name of the bucket. Only applies to &#x60;Amazon&#x60;, &#x60;Azure&#x60;, &#x60;CIFS&#x60;, &#x60;NFSv3&#x60;, &#x60;Openstack Swift&#x60;, and &#x60;Rackspace CDN&#x60;. | [optional]
**create_bucket** | **bool** | Create the bucket if it does not exist. Only applies to &#x60;Amazon&#x60;, &#x60;Azure&#x60;, &#x60;Openstack Swift&#x60;, and &#x60;Rackspace CDN&#x60;. | [optional] [default to false]
**config** | [**OneOfObjectObjectObjectObjectObjectObjectObject**](OneOfObjectObjectObjectObjectObjectObjectObject.md) |  |

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
