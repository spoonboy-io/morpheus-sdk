package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class StorageBucketConfig {
    
    String accessKey
    
    String secretKey
    
    String stsAssumeRole
    
    Boolean useHostCredentials
    
    String endpoint
    
    String secretKeyHash
}
