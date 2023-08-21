package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class BackupSchedule {
    /* Schedule ID */
    Long id
    /* Schedule Name */
    String name
    /* Schedule Cron Expression */
    String cron
}
