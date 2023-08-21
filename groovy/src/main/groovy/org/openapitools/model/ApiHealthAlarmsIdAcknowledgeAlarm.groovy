package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiHealthAlarmsIdAcknowledgeAlarm {
    /* Pass `true` to ackowledge an alarm, or pass `false` to unacknowledge it. */
    Boolean acknowledged
}
