package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class Cypher {
    
    Integer id
    
    String itemKey
    
    Long leaseTimeout
    
    Date expireDate
    
    Date dateCreated
    
    Date lastUpdated
    
    Date lastAccessed
    
    String createdBy
}
