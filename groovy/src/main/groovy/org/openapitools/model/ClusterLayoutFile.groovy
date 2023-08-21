package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ClusterLayoutFile {
    
    Long id
    
    String sourceType
    
    String contentRef
    
    String contentPath
    
    String repository
    
    String content
}
