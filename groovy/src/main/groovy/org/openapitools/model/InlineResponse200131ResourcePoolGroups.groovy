package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse200131ResourcePermission;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class InlineResponse200131ResourcePoolGroups {
    
    Long id
    
    String name
    
    String description
    
    String visibility
    /* Pool selection mode. Valid values are `roundrobin` or `availablecapacity`. */
    String mode
    /* Array of Resource Pool IDs */
    List<Long> pools = new ArrayList<Long>()
    
    List<InlineResponse20040AppDeployInstance> tenants = new ArrayList<InlineResponse20040AppDeployInstance>()
    
    InlineResponse200131ResourcePermission resourcePermission
}
