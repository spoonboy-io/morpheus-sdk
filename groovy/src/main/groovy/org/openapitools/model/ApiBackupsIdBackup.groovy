package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiBackupsIdBackup {
    /* A name for the backup */
    String name
    /* The Backup Job ID to assign the backup to. This determines when the backup is run. */
    Long jobId
    /* Can be used to enable or disable the backup */
    Boolean enabled
}
