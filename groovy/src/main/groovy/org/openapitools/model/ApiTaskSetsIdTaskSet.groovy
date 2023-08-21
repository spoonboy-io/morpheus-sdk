package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiTaskSetsTaskSetTasks;

@Canonical
class ApiTaskSetsIdTaskSet {
    /* A unique name for the workflow */
    String name
    /* A description of the workflow */
    String description
    /* An array of Category labels for filtering */
    List<String> labels = new ArrayList<String>()
    /* Workflow type */
    String type = TypeEnum.PROVISION
    /* List of option type IDs for use with operational workflow configuration. */
    List<Long> optionTypes = new ArrayList<Long>()
    
    ApiTaskSetsTaskSetTasks tasks
}
