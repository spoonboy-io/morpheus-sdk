package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiTasksTaskFile;
import org.openapitools.model.ApiTasksTaskTaskType;
import org.openapitools.model.OneOfobjectobject;

@Canonical
class ApiTasksIdTask {
    /* A unique name for the task */
    String name
    /* A unique code for the task */
    String code
    /* Visibility */
    String visibility = "private"
    
    ApiTasksTaskTaskType taskType
    /* An array of Category labels for filtering */
    List<String> labels = new ArrayList<String>()
    /* Map of options specific to each `task type`. eg. script */
    Object taskOptions
    
    String resultType
    /* The execution target. eg. local,remote,resource. The default value varies by task type.  */
    String executeTarget
    /* If the task should be retried or not. */
    Boolean retryable = false
    /* The number of times to retry. */
    Integer retryCount
    /* The delay, between retries. */
    BigDecimal retryDelaySeconds
    
    ApiTasksTaskFile file
    /* Map containing Credential ID or the default {\"type\": \"local\"}  which means use the values set in the local task options username and password instead of associating a credential.  */
    OneOfobjectobject credential = null
}
