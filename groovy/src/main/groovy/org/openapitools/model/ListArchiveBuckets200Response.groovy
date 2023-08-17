package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.model.ArchiveBucket;
import org.openapitools.model.MetaMeta;

@Canonical
class ListArchiveBuckets200Response {
    
    MetaMeta meta
    
    List<ArchiveBucket> archiveBuckets
    
    Long archiveBucketCount
}
