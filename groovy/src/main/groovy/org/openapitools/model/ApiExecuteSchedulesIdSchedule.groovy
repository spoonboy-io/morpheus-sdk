package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiExecuteSchedulesIdSchedule {
    /* A name for the execute schedule */
    String name
    /* A description for the execute schedule */
    String description
    /* Type of schedule */
    String scheduleType
    /* Time Zone eg. America/New_York, Europe/Amsterdam, etc. */
    String scheduleTimezone = "UTC"
    /* Cron Expression. The default is daily at midnight */
    String cron = "0 0 * * *"
    /* Is enabled */
    Boolean enabled = true
}
