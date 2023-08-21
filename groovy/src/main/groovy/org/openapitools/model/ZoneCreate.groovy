package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.AnyOfobjectobject;
import org.openapitools.model.ZoneCreateCredential;

@Canonical
class ZoneCreate {
    /* A unique name scoped to your account for the cloud */
    String name
    /* Optional description field if you want to put more info there */
    String description
    /* Optional code for use with policies */
    String code
    /* Optional location for your cloud */
    String location
    /* private or public */
    String visibility = VisibilityEnum.PRIVATE
    
    AnyOfobjectobject zoneType = null
    /* Specifies which Server group this cloud should be assigned to */
    Long groupId
    /* Specifies which Tenant this cloud should be assigned to */
    Long accountId
    /* Can be used to disable the cloud */
    Boolean enabled = true
    /* Automatically Power on VMs */
    Boolean autoRecoverPowerState = false
    /* Scale Priority */
    Long scalePriority = 1l
    /* Linked Account ID (enter commercial ID to get costing for AWS Govcloud) */
    Long linkedAccountId
    /* Map containing zone configuration settings. See the section on specific zone types for details. */
    Object config
    /* host firewall. `off` or `internal`. a.k.a. \"local firewall\" */
    String securityMode = "off"
    
    ZoneCreateCredential credential
}
