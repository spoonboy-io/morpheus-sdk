package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiServersIdResizeServer;
import org.openapitools.model.ApiServersIdResizeServicePlanOptions;
import org.openapitools.model.InstanceCreateNetwork;
import org.openapitools.model.InstanceCreateVolume;

@Canonical
class InlineObject224 {
    
    ApiServersIdResizeServer server
    
    ApiServersIdResizeServicePlanOptions servicePlanOptions
    /* List of volumes with their new sizes. */
    List<InstanceCreateVolume> volumes = new ArrayList<InstanceCreateVolume>()
    /* Delete the original volumes after resizing. (Amazon only) */
    Boolean deleteOriginalVolumes = false
    /* Key for network configurations. Include id to update an existing interface. */
    List<InstanceCreateNetwork> networkInterfaces = new ArrayList<InstanceCreateNetwork>()
}
