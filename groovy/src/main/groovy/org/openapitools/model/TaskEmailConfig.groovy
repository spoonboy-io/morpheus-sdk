package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.OptionTypeListCredential;
import org.openapitools.model.TaskAnsiblePlaybookConfigFile;
import org.openapitools.model.TaskEmailConfigTaskOptions;
import org.openapitools.model.TaskEmailConfigTaskType;

@Canonical
class TaskEmailConfig {
    
    Long id
    
    Long accountId
    
    String name
    
    String code
    
    TaskEmailConfigTaskType taskType
    
    List<String> labels = new ArrayList<String>()
    
    String visibility
    
    TaskEmailConfigTaskOptions taskOptions
    
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
