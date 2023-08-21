package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InstanceCreateSuccessInstance;

@Canonical
class InstanceUpdateSuccess {
    
    InstanceCreateSuccessInstance instance
    /* The Cloud ID to provision the instance onto. */
    Long zoneId
}
