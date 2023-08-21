package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20094Network;

@Canonical
class SecurityPackage {
    
    Long id
    
    String name
    
    List<String> labels = new ArrayList<String>()
    
    String description
    
    InlineResponse20094Network type
    
    Boolean enabled
    
    String url
    
    String uuid
    
    Date dateCreated
    
    Date lastUpdated
    
    Object config
}
