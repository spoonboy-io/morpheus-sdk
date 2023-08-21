package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.OneOfstringlong;

@Canonical
class ApiVdiPoolsIdVdiPoolConfig {
    /* Name of instance */
    String name
    
    OneOfstringlong group = null
    
    OneOfstringlong cloud = null
    
    OneOfstringlong type = null
    
    OneOfstringlong layout = null
    
    OneOfstringlong plan = null
}
