package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20040AppDeployDeployment;
import org.openapitools.model.InlineResponse20040AppDeployDeploymentVersion;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class InlineResponse20040AppDeploy {
    
    Long id
    
    Long instanceId
    
    InlineResponse20040AppDeployInstance instance
    
    InlineResponse20040AppDeployDeployment deployment
    
    Long deploymentVersionId
    
    InlineResponse20040AppDeployDeploymentVersion deploymentVersion
    
    Object config
    
    String status
    
    Date deployDate
    
    Date dateCreated
    
    Date lastUpdated
}
