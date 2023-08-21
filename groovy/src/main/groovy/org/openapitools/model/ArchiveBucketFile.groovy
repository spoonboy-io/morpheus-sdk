package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ArchiveBucketFileArchiveBucket;
import org.openapitools.model.ArchiveBucketFileCreatedBy;

@Canonical
class ArchiveBucketFile {
    
    Long id
    
    String name
    
    String filePath
    
    ArchiveBucketFileArchiveBucket archiveBucket
    
    ArchiveBucketFileCreatedBy createdBy
    
    Boolean isDirectory
    
    String status
    
    Long rawSize
    
    String contentType
    
    Long downloadCount
    
    Date dateCreated
    
    Date lastUpdated
}
