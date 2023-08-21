package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.IntegrationSNOWConfigServiceNowCmdbClassMapping;

@Canonical
class IntegrationSNOWConfig {
    
    Boolean incidentAccess
    
    Boolean requestAccess
    
    String serviceNowCMDBBusinessObject
    
    String serviceNowCustomCmdbMapping
    
    List<IntegrationSNOWConfigServiceNowCmdbClassMapping> serviceNowCmdbClassMapping = new ArrayList<IntegrationSNOWConfigServiceNowCmdbClassMapping>()
    
    String webServiceImportUrl
    
    String webServiceImportSysId
    
    String webServiceOperationUrl
    
    Boolean preparedForSync
}
