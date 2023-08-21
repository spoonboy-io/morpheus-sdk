package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.OptionTypeListCredential;
import org.openapitools.model.TaskNestedWorkflowConfigTaskOptions;
import org.openapitools.model.TaskNestedWorkflowConfigTaskType;

@Canonical
class TaskNestedWorkflowConfig {
    
    Long id
    
    Long accountId
    
    String name
    
    String code
    
    TaskNestedWorkflowConfigTaskType taskType
    
    List<String> labels = new ArrayList<String>()
    
    String visibility
    
    TaskNestedWorkflowConfigTaskOptions taskOptions
    
    String resultType
    
    String executeTarget
    
    Boolean retryable
    
    Long retryCount
    
    Long retryDelaySeconds
    
    Boolean allowCustomConfig
    
    OptionTypeListCredential credential
    
    Date dateCreated
    
    Date lastUpdated
}
