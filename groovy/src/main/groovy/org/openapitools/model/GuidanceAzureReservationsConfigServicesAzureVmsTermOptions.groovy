package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.GuidanceAzureReservationsConfigDetailList;
import org.openapitools.model.GuidanceAzureReservationsConfigServicesAzureVmsSummary;

@Canonical
class GuidanceAzureReservationsConfigServicesAzureVmsTermOptions {
    
    String code
    
    String name
    
    List<GuidanceAzureReservationsConfigDetailList> detailList = new ArrayList<GuidanceAzureReservationsConfigDetailList>()
    
    GuidanceAzureReservationsConfigServicesAzureVmsSummary summary
}
