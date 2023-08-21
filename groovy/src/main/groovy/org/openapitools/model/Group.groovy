package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.GroupConfig;
import org.openapitools.model.GroupStats;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class Group {
    
    Long id
    
    String uuid
    
    String name
    
    String code
    
    String location
    
    Long accountId
    
    Boolean active
    
    GroupConfig config
    
    Date dateCreated
    
    Date lastUpdated
    
    List<InlineResponse20040AppDeployInstance> zones = new ArrayList<InlineResponse20040AppDeployInstance>()
    
    GroupStats stats
    
    Long serverCount
}
