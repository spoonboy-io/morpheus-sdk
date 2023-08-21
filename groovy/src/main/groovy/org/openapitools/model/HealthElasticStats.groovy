package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class HealthElasticStats {
    
    String status
    
    String clusterName
    
    String nodeTotal
    
    String nodeData
    
    String shards
    
    String primary
    
    String relocating
    
    String initializing
    
    String unassigned
    
    String pendingTasks
    
    String activePercent
}
