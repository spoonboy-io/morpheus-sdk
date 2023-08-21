package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.OneOfobjectobjectobjectobjectobjectobjectobject;

@Canonical
class ApiStorageBucketsStorageBucket {
    /* A unique name scoped to your account for the storage bucket */
    String name
    /* The type of storage bucket */
    String providerType
    /* Default Backup Target */
    Boolean defaultBackupTarget = false
    /* Archive Snapshots */
    Boolean copyToStore
    /* Default Deployment Target */
    Boolean defaultDeploymentTarget = false
    /* Default Virtual Image Store */
    Boolean defaultVirtualImageTarget = false
    /* Cleanup mode. `backup` - Move old files to a backup provider. `delete` - Delete old files. `none` - Keep all files. */
    String retentionPolicyType = RetentionPolicyTypeEnum.NONE
    /* The number of days old a file must be before it is deleted. */
    Long retentionPolicyDays
    /* The backup Storage Bucket where old files are moved to. */
    String retentionProvider
    /* The name of the bucket. Only applies to `Amazon`, `Azure`, `CIFS`, `NFSv3`, `Openstack Swift`, and `Rackspace CDN`. */
    String bucketName
    /* Create the bucket if it does not exist. Only applies to `Amazon`, `Azure`, `Openstack Swift`, and `Rackspace CDN`. */
    Boolean createBucket = false
    
    OneOfobjectobjectobjectobjectobjectobjectobject config = null
}
