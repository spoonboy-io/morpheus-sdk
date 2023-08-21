package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class BackupStats {
    /* Total size of all backups in bytes */
    Long totalSize
    /* Average size of each backup in bytes */
    Long avgSize
    /* Total completed count */
    Long totalCompleted
    /* Successful backup count */
    Long success
    /* Failed backup count */
    Long failed
    /* Success rate 0-1 */
    Double successRate
    /* Failure rate 0-1 */
    Double failRate
    /* List of the last 5 backup result statuses */
    List<String> lastFiveResults = new ArrayList<LastFiveResultsEnum>()
}
