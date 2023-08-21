package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InstanceUpdateInstance;
import org.openapitools.model.InstancesConfigCustomOptions;

@Canonical
class InstanceUpdate {
    
    InstanceUpdateInstance instance
    
    InstancesConfigCustomOptions config
}
