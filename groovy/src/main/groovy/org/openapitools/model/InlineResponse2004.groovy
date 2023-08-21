package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ArchiveBucket;
import org.openapitools.model.ArchiveBucketFile;

@Canonical
class InlineResponse2004 {
    
    ArchiveBucket archiveBucket
    
    Boolean isOwner
    
    String parentDirectory
    
    List<ArchiveBucketFile> archiveFiles = new ArrayList<ArchiveBucketFile>()
    
    Long archiveFileCount
}
