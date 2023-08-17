package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BillingZone;

@Canonical
class GetBillingZoneIdentifier200Response {
    
    Boolean success
    
    BillingZone billingInfo
}
