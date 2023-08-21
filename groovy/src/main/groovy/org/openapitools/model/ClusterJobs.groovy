package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20083LoadBalancerNodeCreatedBy;
import org.openapitools.model.InlineResponse20094Network;

@Canonical
class ClusterJobs {
    
    Long id
    
    String name
    
    List<String> labels = new ArrayList<String>()
    
    InlineResponse20094Network type
    
    String jobSummary
    
    String scheduleMode
    
    Date dateTime
    
    String status
    
    String namespace
    
    String category
    
    String description
    
    Boolean enabled
    
    Date dateCreated
    
    Date lastUpdated
    
    Date lastRun
    
    String lastResult
    
    InlineResponse20083LoadBalancerNodeCreatedBy createdBy
    
    String targetType
    
    List<Object> targets = new ArrayList<Object>()
    
    Object customConfig
    
    Object customOptions
}
