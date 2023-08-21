package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20094Network;

@Canonical
class IntegrationObjectLayout {
    
    Long id
    
    String name
    
    String code
    
    InlineResponse20094Network provisionType
    
    InlineResponse20094Network instanceType
    
    String instanceVersion
}
