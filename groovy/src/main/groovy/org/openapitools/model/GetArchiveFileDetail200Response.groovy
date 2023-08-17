package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ArchiveBucketFile;

@Canonical
class GetArchiveFileDetail200Response {
    
    ArchiveBucketFile archiveFile
    
    Boolean isOwner
}
