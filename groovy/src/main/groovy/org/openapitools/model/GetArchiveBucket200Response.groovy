package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.ArchiveBucket;
import org.openapitools.model.ArchiveBucketFile;

@Canonical
class GetArchiveBucket200Response {
    
    ArchiveBucket archiveBucket
    
    Boolean isOwner
    
    String parentDirectory
    
    List<ArchiveBucketFile> archiveFiles
    
    Long archiveFileCount
}
