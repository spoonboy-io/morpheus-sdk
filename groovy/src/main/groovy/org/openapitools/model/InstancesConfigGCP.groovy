package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class InstancesConfigGCP {
    /* Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. */
    Boolean noAgent = false
    /* External ID of the Google zone to use for instance. */
    Long googleZoneId
    /* External IP Type.  `-1` for ephemeral IP. */
    Long externalIpType
    /* Network Tags */
    String networkTags
    /* Service Account */
    String serviceAccount
    /* Access Scope */
    String accessScope
}
