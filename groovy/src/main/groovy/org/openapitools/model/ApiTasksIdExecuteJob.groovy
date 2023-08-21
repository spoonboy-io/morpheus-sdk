package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class ApiTasksIdExecuteJob {
    /* A name for the execution job. Can be used to find execution results with `/api/processes?name=` */
    String name
    /* The target context for task execution. This is required for tasks with `executeTarget` set to `resource`. */
    String targetType
    /* Array of Instance IDs. Only applicable if `targetType` is instance. */
    List<Long> instances = new ArrayList<Long>()
    /* Array of Server IDs. Only applicable if `targetType` is `server`. */
    List<Long> servers = new ArrayList<Long>()
    /* Instance Label. Only applicable if `targetType` is `instance-label`. */
    String instanceLabel
    /* Server Label. Only applicable if `targetType` is `server-label`. */
    String serverLabel
    /* Map of options to be used as values in the task. These correspond to option types. */
    Object customOptions
    /* String of custom configuration values as JSON. */
    String customConfig
}
