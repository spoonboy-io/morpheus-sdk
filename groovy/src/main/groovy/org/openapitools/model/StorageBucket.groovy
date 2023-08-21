package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.StorageBucketConfig;

@Canonical
class StorageBucket {
    
    Long id
    
    String name
    
    Boolean active
    
    Long accountId
    
    String providerType
    
    StorageBucketConfig config
    
    String bucketName
    
    Boolean readOnly
    
    Boolean defaultBackupTarget
    
    Boolean defaultDeploymentTarget
    
    Boolean defaultVirtualImageTarget
    
    Boolean copyToStore
    
    String retentionPolicyType
    
    String retentionPolicyDays
    
    String retentionProvider
}
