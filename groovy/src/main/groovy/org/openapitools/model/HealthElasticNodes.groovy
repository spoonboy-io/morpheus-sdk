package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class HealthElasticNodes {
    
    String ip
    
    String heapPercent
    
    String ramPercent
    
    String cpuCount
    
    String loadOne
    
    String loadFive
    
    String loadFifteen
    
    String role
    
    String master
    
    String name
}
