package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.HashMap;
import java.util.List;
import org.openapitools.model.GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions;

@Canonical
class GuidanceAzureReservationsConfigServicesAzureVms {
    
    String code
    
    String name
    
    Map<String, GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions> paymentOptions = new HashMap<String, GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions>()
}
