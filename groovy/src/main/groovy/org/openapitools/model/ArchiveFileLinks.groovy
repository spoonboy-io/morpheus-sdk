package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.ArchiveBucketFileCreatedBy;
import org.openapitools.model.ArchiveFileLinksArchiveFile;

@Canonical
class ArchiveFileLinks {
    
    Long id
    
    String secretAccessKey
    
    ArchiveFileLinksArchiveFile archiveFile
    
    ArchiveBucketFileCreatedBy createdBy
    
    Date dateCreated
    
    Date lastUpdated
    
    Date lastAccessDate
    
    Date expirationDate
    
    Long downloadCount
}
