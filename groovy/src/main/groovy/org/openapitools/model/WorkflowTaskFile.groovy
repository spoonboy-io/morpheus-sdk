package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;

@Canonical
class WorkflowTaskFile {
    
    Long id
    
    String sourceType
    
    String contentRef
    
    String contentPath
    
    InlineResponse20082LoadBalancerInstanceSslCert repository
    
    String content
}
