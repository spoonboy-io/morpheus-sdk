package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;

@Canonical
class FileTemplate {
    
    Long id
    
    String code
    
    InlineResponse20082LoadBalancerInstanceSslCert account
    
    String name
    
    List<String> labels = new ArrayList<String>()
    
    String fileName
    
    String filePath
    
    String templateType
    
    String templatePhase
    
    String template
    
    String category
    
    String settingCategory
    
    String settingName
    
    Boolean autoRun
    
    Boolean runOnScale
    
    Boolean runOnDeploy
    
    String fileOwner
    
    String fileGroup
    
    String permissions
    
    Date dateCreated
    
    Date lastUpdated
}
