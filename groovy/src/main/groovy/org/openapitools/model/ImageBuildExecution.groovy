package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ArchiveBucketFileCreatedBy;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class ImageBuildExecution {
    
    Long id
    
    InlineResponse20040AppDeployInstance imageBuild
    
    Long buildNumber
    
    Date startDate
    
    Date endDate
    
    String statusMessage
    
    Long statusPercent
    
    String statusEta
    
    String status
    
    String errorMessage
    
    ArchiveBucketFileCreatedBy createdBy
    
    String tempInstance
    
    List<Object> virtualImages = new ArrayList<Object>()
}
