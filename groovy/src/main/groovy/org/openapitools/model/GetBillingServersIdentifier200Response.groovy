package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.BillingServer;

@Canonical
class GetBillingServersIdentifier200Response {
    
    Boolean success
    
    BillingServer billingInfo
}
