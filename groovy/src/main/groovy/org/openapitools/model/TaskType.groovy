package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.TaskTypeOptionTypes;

@Canonical
class TaskType {
    
    Long id
    
    String code
    
    String name
    
    String category
    
    String description
    
    Boolean scriptable
    
    Boolean enabled
    
    Boolean hasResults
    
    Boolean allowExecuteLocal
    
    Boolean allowExecuteRemote
    
    Boolean allowExecuteResource
    
    Boolean allowLocalRepo
    
    Boolean allowRemoteKeyAuth
    
    List<TaskTypeOptionTypes> optionTypes = new ArrayList<TaskTypeOptionTypes>()
}
