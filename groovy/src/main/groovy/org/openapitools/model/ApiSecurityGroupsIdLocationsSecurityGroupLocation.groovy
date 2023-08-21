package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.AnyOfsecurityGroupLocationAzureCustomOptionssecurityGroupLocationAwsCustomOptionssecurityGroupLocationOpenstackCustomOptions;

@Canonical
class ApiSecurityGroupsIdLocationsSecurityGroupLocation {
    /* The ID of the Zone (Cloud) */
    Long zoneId
    
    AnyOfsecurityGroupLocationAzureCustomOptionssecurityGroupLocationAwsCustomOptionssecurityGroupLocationOpenstackCustomOptions customOptions = null
}
