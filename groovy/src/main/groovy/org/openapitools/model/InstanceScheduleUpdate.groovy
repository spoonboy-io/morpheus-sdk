package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InstanceScheduleUpdateThreshold;

@Canonical
class InstanceScheduleUpdate {
    
    String scheduleType = ScheduleTypeEnum.DAYOFWEEK
    /* Time Zone eg. America/New_York, Europe/Amsterdam, etc. Only used and required for scheduleType `dayOfWeek` */
    String scheduleTimezone = "UTC"
    /* Start day of the week 1-7 (Sun-Sat). Only used and required for scheduleType `dayOfWeek` */
    Integer startDayOfWeek
    /* Start time of the day in 24-hour format. Only used and required for scheduleType `dayOfWeek` */
    String startTime
    /* End day of the week 1-7 (Sun-Sat). Only used and required for scheduleType `dayOfWeek` */
    Integer endDayOfWeek
    /* End time of the day in 24-hour format. Only used and required for scheduleType `dayOfWeek` */
    String endTime
    /* Start Date. Only used and required for scheduleType `exact` */
    Date startDate
    /* End Date. Only used and required for scheduleType `exact` */
    Date endDate
    
    InstanceScheduleUpdateThreshold threshold
}
