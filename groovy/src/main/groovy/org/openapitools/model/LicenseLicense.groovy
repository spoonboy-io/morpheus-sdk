package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.LicenseLicenseFeatures;

@Canonical
class LicenseLicense {
    
    String productTier
    
    Date startDate
    
    Date endDate
    
    Long maxInstances
    
    Long maxMemory
    
    Long maxStorage
    
    Boolean hardLimit
    
    Boolean freeTrial
    
    Boolean multiTenant
    
    Boolean whitelabel
    
    Boolean reportStatus
    
    String supportLevel
    
    String accountName
    
    Object config
    
    String amazonProductCodes
    
    LicenseLicenseFeatures features
    
    String zoneTypes
    
    Date lastUpdated
    
    Date dateCreated
}
