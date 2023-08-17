package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.ApplianceSettingsEnabledZoneTypesInner;
import org.openapitools.model.ArchiveBucketCreatedBy;

@Canonical
class ArchiveBucket {
    
    Long id
    
    String name
    
    String description
    
    ApplianceSettingsEnabledZoneTypesInner storageProvider
    
    ApplianceSettingsEnabledZoneTypesInner owner
    
    ArchiveBucketCreatedBy createdBy
    
    Boolean isPublic
    
    String visibility
    
    String code
    
    String filePath
    
    Long rawSize
    
    Long fileCount
    
    List<Object> accounts
    
    Date dateCreated
    
    Date lastUpdated
}
