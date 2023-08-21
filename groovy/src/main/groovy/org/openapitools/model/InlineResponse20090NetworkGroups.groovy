package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class InlineResponse20090NetworkGroups {
    
    Long id
    
    String name
    
    String description
    
    String visibility
    
    Boolean active
    
    List<Long> networks = new ArrayList<Long>()
    
    List<Object> subnets = new ArrayList<Object>()
    
    List<InlineResponse20040AppDeployInstance> tenants = new ArrayList<InlineResponse20040AppDeployInstance>()
}
