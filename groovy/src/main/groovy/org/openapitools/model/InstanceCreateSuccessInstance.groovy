package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InstanceCreateInstanceInstanceType;
import org.openapitools.model.InstanceCreateInstanceLayout;
import org.openapitools.model.InstanceCreateInstancePlan;
import org.openapitools.model.InstanceCreateInstanceSite;

@Canonical
class InstanceCreateSuccessInstance {
    /* Name of the instance to be created. */
    String name
    
    InstanceCreateInstanceSite site
    
    InstanceCreateInstanceInstanceType instanceType
    
    InstanceCreateInstanceLayout layout
    
    InstanceCreateInstancePlan plan
}
