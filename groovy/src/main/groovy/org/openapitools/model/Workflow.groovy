package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.WorkflowOptionTypes;
import org.openapitools.model.WorkflowTaskSetTasks;

@Canonical
class Workflow {
    
    Long id
    
    String name
    
    String type
    
    String description
    
    List<String> labels = new ArrayList<String>()
    
    Date dateCreated
    
    Date lastUpdated
    
    Long accountId
    
    String platform
    
    String visibility
    
    Boolean allowCustomConfig
    
    List<Long> tasks = new ArrayList<Long>()
    
    List<WorkflowOptionTypes> optionTypes = new ArrayList<WorkflowOptionTypes>()
    
    List<WorkflowTaskSetTasks> taskSetTasks = new ArrayList<WorkflowTaskSetTasks>()
}
