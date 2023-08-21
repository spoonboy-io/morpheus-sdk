package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.GuidanceAzureReservationsConfigDetailList;
import org.openapitools.model.GuidanceAzureReservationsConfigServices;
import org.openapitools.model.GuidanceAzureReservationsConfigServicesAzureVmsSummary;

@Canonical
class GuidanceAzureReservationsConfig {
    
    Boolean success
    
    List<GuidanceAzureReservationsConfigDetailList> detailList = new ArrayList<GuidanceAzureReservationsConfigDetailList>()
    
    GuidanceAzureReservationsConfigServices services
    
    GuidanceAzureReservationsConfigServicesAzureVmsSummary summary
}
