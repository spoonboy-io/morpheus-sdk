package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BillingZones;

@Canonical
class ListBillingZone200Response {
    
    Boolean success
    
    BillingZones billingInfo
}
