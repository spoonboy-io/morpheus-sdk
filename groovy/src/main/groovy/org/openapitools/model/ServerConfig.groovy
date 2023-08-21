package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ServerConfig {
    
    String poolProviderType
    
    Boolean isVpcSelectable
    
    String smbiosAssetTag
    
    Boolean isEC2
    
    Long resourcePoolId
    
    Long hostId
    
    Boolean createUser
    
    String nestedVirtualization
    
    String vmwareFolderId
    
    Boolean noAgent
    
    Long powerScheduleType
}
