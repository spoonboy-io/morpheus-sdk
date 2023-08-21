package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.GroupStats;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class TenantGroup {
    
    Long id
    
    String name
    
    String code
    
    String location
    
    Long accountId
    
    String visibility
    
    Boolean active
    
    Date dateCreated
    
    Date lastUpdated
    
    List<InlineResponse20040AppDeployInstance> zones = new ArrayList<InlineResponse20040AppDeployInstance>()
    
    GroupStats stats
    
    Long serverCount
}
