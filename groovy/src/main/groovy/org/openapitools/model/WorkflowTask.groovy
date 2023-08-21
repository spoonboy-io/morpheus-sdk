package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.PriceSetVolumeType;
import org.openapitools.model.WorkflowTaskFile;
import org.openapitools.model.WorkflowTaskTaskOptions;

@Canonical
class WorkflowTask {
    
    Long id
    
    Long accountId
    
    String name
    
    String code
    
    PriceSetVolumeType taskType
    
    List<String> labels = new ArrayList<String>()
    
    WorkflowTaskTaskOptions taskOptions
    
    WorkflowTaskFile file
    
    String resultType
    
    String executeTarget
    
    Boolean retryable
    
    Long retryCount
    
    Long retryDelaySeconds
    
    Boolean allowCustomConfig
    
    Date dateCreated
    
    Date lastUpdated
}
