package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.DeploymentVersions;

@Canonical
class Deployment {
    
    Long id
    
    String name
    
    String description
    
    Long accountId
    
    String externalId
    
    Date dateCreated
    
    Date lastUpdated
    
    Long versionCount
    
    List<DeploymentVersions> versions = new ArrayList<DeploymentVersions>()
}
