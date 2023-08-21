package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.MonitoringSettingsServiceNowIntegration;

@Canonical
class MonitoringSettingsServiceNow {
    /* Enabled */
    Boolean enabled
    
    MonitoringSettingsServiceNowIntegration integration
    /* New Incident Action */
    String newIncidentAction
    /* Close Incident Action */
    String closeIncidentAction
    /* Info Mapping */
    String infoMapping
    /* Warning Mapping */
    String warningMapping
    /* Critical Mapping */
    String criticalMapping
}
