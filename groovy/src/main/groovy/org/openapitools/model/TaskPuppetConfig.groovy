package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.OptionTypeListCredential;
import org.openapitools.model.TaskAnsiblePlaybookConfigFile;
import org.openapitools.model.TaskPuppetConfigTaskOptions;
import org.openapitools.model.TaskPuppetConfigTaskType;

@Canonical
class TaskPuppetConfig {
    
    Long id
    
    Long accountId
    
    String name
    
    String code
    
    TaskPuppetConfigTaskType taskType
    
    List<String> labels = new ArrayList<String>()
    
    String visibility
    
    TaskPuppetConfigTaskOptions taskOptions
    
    TaskAnsiblePlaybookConfigFile file
    
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
