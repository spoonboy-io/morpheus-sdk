package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.AppStateSpecsInnerTemplate;

@Canonical
class AppStateSpecsInner {
    
    Long id
    
    String name
    
    AppStateSpecsInnerTemplate template
    
    Boolean isolated
}
