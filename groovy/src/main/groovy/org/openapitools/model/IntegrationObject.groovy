package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.IntegrationObjectLayout;

@Canonical
class IntegrationObject {
    
    Long id
    
    String name
    
    String type
    
    String refType
    
    Long refId
    
    IntegrationObjectLayout layout
}
