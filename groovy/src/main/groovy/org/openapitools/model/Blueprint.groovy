package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse200107NetworkPoolCreatedBy;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class Blueprint {
    
    Long id
    
    String name
    
    List<String> labels = new ArrayList<String>()
    
    String type
    
    String description
    
    String category
    
    Object config
    
    String visibility
    
    Object resourcePermission
    
    InlineResponse200107NetworkPoolCreatedBy owner
    
    InlineResponse20040AppDeployInstance tenant
}
