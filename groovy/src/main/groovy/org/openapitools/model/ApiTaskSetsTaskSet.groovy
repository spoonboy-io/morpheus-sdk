package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiTaskSetsTaskSetTasks;

@Canonical
class ApiTaskSetsTaskSet {
    /* A unique name for the workflow */
    String name
    /* A description of the workflow */
    String description
    /* An array of Category labels for filtering */
    List<String> labels = new ArrayList<String>()
    /* Workflow type */
    String type = TypeEnum.PROVISION
    /* private or public */
    String visibility = VisibilityEnum.PRIVATE
    /* List of option type IDs for use with operational workflow configuration. */
    List<Long> optionTypes = new ArrayList<Long>()
    
    ApiTaskSetsTaskSetTasks tasks
}
