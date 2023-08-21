package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.IntegrationChefConfigDatabags;

@Canonical
class IntegrationChefConfig {
    
    List<IntegrationChefConfigDatabags> databags = new ArrayList<IntegrationChefConfigDatabags>()
    
    String endpoint
    
    String org
    
    String chefUser
    
    String userKey
    
    String orgKey
    
    String version
    
    Boolean chefUseFqdn
    
    String windowsVersion
    
    String windowsInstallUrl
    
    String userKeyHash
    
    String orgKeyHash
}
