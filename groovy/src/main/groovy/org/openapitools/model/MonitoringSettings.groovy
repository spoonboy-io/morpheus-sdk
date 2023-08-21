package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.MonitoringSettingsServiceNow;

@Canonical
class MonitoringSettings {
    /* Auto Create Checks */
    Boolean autoManageChecks
    /* Availability Time Frame. The number of days availability should be calculated for. Changes will not take effect until your checks have passed their check interval. */
    Integer availabilityTimeFrame
    /* Availability Precision. The number of decimal places availability should be displayed in. Can be anywhere between 0 and 5. */
    Integer availabilityPrecision
    /* Default Check Interval. The number of minutes to use as the default interval to use when creating new checks. */
    Integer defaultCheckInterval
    
    MonitoringSettingsServiceNow serviceNow
}
