package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiTaskSetsTaskSetTasks {
    /* Task ID */
    Long taskId
    /* Task Phase. Pass operation for `operational` workflows. */
    String taskPhase = TaskPhaseEnum.PROVISION
}
