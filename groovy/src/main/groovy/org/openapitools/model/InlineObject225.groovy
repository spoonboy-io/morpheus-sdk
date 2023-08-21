package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ApiServersIdWorkflowTaskSet;

@Canonical
class InlineObject225 {
    
    ApiServersIdWorkflowTaskSet taskSet
    /* Task Phase to run for Provisioning workflows. The default is `provision`. */
    String taskPhase = TaskPhaseEnum.PROVISION
}
