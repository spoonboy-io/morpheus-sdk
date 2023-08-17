package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ApplianceSettingsEnabledZoneTypesInner;
import org.openapitools.model.BackupSettingsDefaultSchedule;

@Canonical
class BackupSettings {
    
    Boolean backupsEnabled
    
    Boolean createBackups
    
    Boolean backupAppliance
    
    ApplianceSettingsEnabledZoneTypesInner defaultStorageBucket
    
    BackupSettingsDefaultSchedule defaultSchedule
    
    Long retentionCount
}
