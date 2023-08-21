package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ArchiveBucketCreatedBy;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class ArchiveBucket {
    
    Long id
    
    String name
    
    String description
    
    InlineResponse20040AppDeployInstance storageProvider
    
    InlineResponse20040AppDeployInstance owner
    
    ArchiveBucketCreatedBy createdBy
    
    Boolean isPublic
    
    String visibility
    
    String code
    
    String filePath
    
    Long rawSize
    
    Long fileCount
    
    List<Object> accounts = new ArrayList<Object>()
    
    Date dateCreated
    
    Date lastUpdated
}
