package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ClusterLayoutFile;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20094Network;

@Canonical
class SpecTemplate {
    
    Long id
    
    InlineResponse20040AppDeployInstance account
    
    String name
    
    List<String> labels = new ArrayList<String>()
    
    String code
    
    InlineResponse20094Network type
    
    String externalId
    
    String externalType
    
    String deploymentId
    
    String status
    
    ClusterLayoutFile file
    
    Object config
    
    String createdBy
    
    String updatedBy
    
    Date dateCreated
    
    Date lastUpdated
}
