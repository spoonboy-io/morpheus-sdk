package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class VdiPoolConfigConfig {
    
    Boolean createUser
    
    Boolean isEC2
    
    Boolean isVpcSelectable
    
    Boolean noAgent
    
    String smbiosAssetTag
    
    String nestedVirtualization
    
    String vmwareFolderId
    
    Long resourcePoolId
    
    String poolProviderType
}
