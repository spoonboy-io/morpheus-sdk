package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.HealthRabbitQueues;

@Canonical
class HealthRabbit {
    
    Boolean success
    
    List<Object> busyQueues = new ArrayList<Object>()
    
    List<Object> errorQueues = new ArrayList<Object>()
    
    String status
    
    List<HealthRabbitQueues> queues = new ArrayList<HealthRabbitQueues>()
}
