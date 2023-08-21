package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ArchiveBucketCreatedBy;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class BootScript {
    
    Long id
    
    InlineResponse20040AppDeployInstance account
    
    String fileName
    
    String description
    
    String content
    
    ArchiveBucketCreatedBy createdBy
    
    String visibility
}
