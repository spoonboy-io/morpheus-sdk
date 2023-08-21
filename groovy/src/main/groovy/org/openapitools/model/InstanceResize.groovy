package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InstanceCreateNetwork;
import org.openapitools.model.InstanceCreateVolume;
import org.openapitools.model.InstanceResizeInstance;
import org.openapitools.model.ServicePlanOptions;

@Canonical
class InstanceResize {
    
    InstanceResizeInstance instance
    
    ServicePlanOptions servicePlanOptions
    /* Can be used to grow just the logical volume of the instance instead of choosing a plan */
    List<InstanceCreateVolume> volumes = new ArrayList<InstanceCreateVolume>()
    /* Delete the original volumes after resizing. (Amazon only) */
    Boolean deleteOriginalVolumes = false
    /* Key for network configuration. Include id to update an existing interface. The existing interfaces and their id can be retrieved with the hosts API. */
    List<InstanceCreateNetwork> networkInterfaces = new ArrayList<InstanceCreateNetwork>()
}
