package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiMonitoringIncidentsIncident {
    /* Description of the resolution to this incident */
    String resolution
    /* Comment on this incident, updates summary field */
    String comment
    /* Set status */
    String status
    /* Set severity */
    String severity
    /* Set display name */
    String name
    /* Set start time */
    Date startDate
    /* Set start time */
    Date endDate
    /* Set 'In Availability' */
    Boolean inUptime
}
