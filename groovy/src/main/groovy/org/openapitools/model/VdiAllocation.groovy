package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.VdiAllocationInstance;
import org.openapitools.model.VdiPoolOwner;

@Canonical
class VdiAllocation {
    
    Long id
    
    Long poolId
    
    InlineResponse20040AppDeployInstance pool
    
    VdiAllocationInstance instance
    
    VdiPoolOwner user
    
    Boolean localUserCreated
    
    Boolean persistent
    
    Boolean recyclable
    
    String status
    
    Date dateCreated
    
    Date lastUpdated
    
    Date lastReserved
    
    Date releaseDate
}
