package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.OneOfstringobject;

@Canonical
class AppStateInputData {
    
    String key
    
    OneOfstringobject name = null
    
    String type
}
