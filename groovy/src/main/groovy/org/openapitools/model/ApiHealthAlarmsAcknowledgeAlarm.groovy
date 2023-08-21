package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class ApiHealthAlarmsAcknowledgeAlarm {
    /* Pass `true` to ackowledge an alarm, or pass `false` to unacknowledge it. */
    Boolean acknowledged
    /* Array of Alarm ID(s)to be updated. */
    List<Long> ids = new ArrayList<Long>()
    /* Pass `true` to update all alarms instead of passing ids. This will update any active alarm that is not already acknowledged.  */
    Boolean all = false
}
