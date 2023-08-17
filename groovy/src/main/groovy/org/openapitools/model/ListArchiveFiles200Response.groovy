package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.ArchiveBucket;
import org.openapitools.model.ArchiveBucketFile;
import org.openapitools.model.MetaMeta;

@Canonical
class ListArchiveFiles200Response {
    
    MetaMeta meta
    
    ArchiveBucket archiveBucket
    
    Boolean isOwner
    
    String parentDirectory
    
    List<ArchiveBucketFile> archiveFiles
}
