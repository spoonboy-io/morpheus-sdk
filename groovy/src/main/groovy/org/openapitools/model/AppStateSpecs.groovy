package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.AppStateTemplate;

@Canonical
class AppStateSpecs {
    
    Long id
    
    String name
    
    AppStateTemplate template
    
    Boolean isolated
}
