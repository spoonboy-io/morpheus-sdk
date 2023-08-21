package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiBackupsJobsJob {
    /* A name for the backup job */
    String name
    /* A code for the backup job */
    String code
    /* Execute Schedule ID to use for the backup job */
    Long scheduleId
    /* Retention Count. By default the backup settings value will be used. */
    Long retentionCount
}
