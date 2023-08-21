package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.LicenseCurrentUsage;
import org.openapitools.model.LicenseLicense;

@Canonical
class License {
    
    LicenseLicense license
    
    LicenseCurrentUsage currentUsage
}
