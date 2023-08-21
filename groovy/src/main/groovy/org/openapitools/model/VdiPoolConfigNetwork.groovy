package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class VdiPoolConfigNetwork {
    
    String id
    
    Boolean hasPool
    
    String idName
    
    InlineResponse20040AppDeployInstance pool
}
