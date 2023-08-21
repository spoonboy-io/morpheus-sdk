package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.TenantRole;
import org.openapitools.model.TenantStats;

@Canonical
class Tenant {
    
    Long id
    
    String name
    
    String description
    
    String subdomain
    
    String currency
    
    String externalId
    
    String customerNumber
    
    String accountNumber
    
    String accountName
    
    Boolean active
    
    Boolean master
    
    TenantRole role
    
    TenantStats stats
    
    Date dateCreated
    
    Date lastUpdated
}
