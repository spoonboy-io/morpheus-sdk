package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class BackupTypeServer {
    
    String locationType
    /* A name for the backup */
    String name
    /* The ID of the server to backup */
    Long serverId
    /* The backup type code, options vary by the type of cloud and source */
    String backupType
    /* Create a new backup job, clone an existing job or add the new backup to an existing job */
    String jobAction
    /* The ID of the job to clone or add to. Only applies to jobAction `clone` and `addTo`. */
    Long jobId
    /* Name for new job. Only applies to jobAction `new` and `clone`. */
    String jobName
    /* The ID of the execute schedule for new job. See Execute Schedules. Only applies to jobAction `new` and `clone`. */
    Long jobSchedule
    /* Retention Count for new job. By default the backup settings value will be used. Only applies to jobAction `new` and `clone`. */
    Long retentionCount
}
