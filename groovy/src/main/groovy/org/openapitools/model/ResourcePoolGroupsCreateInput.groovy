package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse200131ResourcePermission;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class ResourcePoolGroupsCreateInput {
    
    String name
    
    String description
    
    String visibility
    /* Pool selection mode. Valid values are `roundrobin` or `availablecapacity`. */
    String mode
    
    List<Long> pools = new ArrayList<Long>()
    
    List<InlineResponse20040AppDeployInstance> tenants = new ArrayList<InlineResponse20040AppDeployInstance>()
    
    InlineResponse200131ResourcePermission resourcePermission
}
