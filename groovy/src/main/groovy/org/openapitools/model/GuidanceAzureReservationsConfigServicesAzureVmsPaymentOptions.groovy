package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.HashMap;
import java.util.List;
import org.openapitools.model.GuidanceAzureReservationsConfigServicesAzureVmsTermOptions;

@Canonical
class GuidanceAzureReservationsConfigServicesAzureVmsPaymentOptions {
    
    String code
    
    String name
    
    Map<String, GuidanceAzureReservationsConfigServicesAzureVmsTermOptions> termOptions = new HashMap<String, GuidanceAzureReservationsConfigServicesAzureVmsTermOptions>()
}
