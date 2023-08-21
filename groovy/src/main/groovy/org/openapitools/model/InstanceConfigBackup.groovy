package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class InstanceConfigBackup {
    
    Boolean createBackup
    
    String jobAction
    
    String jobRetentionCount
    
    Long providerBackupType
}
