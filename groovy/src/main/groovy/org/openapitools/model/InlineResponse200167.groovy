package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse200167Appliance;
import org.openapitools.model.InlineResponse200167Permissions;
import org.openapitools.model.User;

@Canonical
class InlineResponse200167 {
    
    User user
    
    Boolean isMasterAccount
    
    List<InlineResponse200167Permissions> permissions = new ArrayList<InlineResponse200167Permissions>()
    
    InlineResponse200167Appliance appliance
}
