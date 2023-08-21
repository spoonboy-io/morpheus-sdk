package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiMonitoringContactsIdContact {
    /* Unique name scoped to your account for the contact */
    String name
    /* Email notification address */
    String emailAddress
    /* SMS notification address */
    String smsAddress
    /* Slack Hook */
    String slackHook
}
