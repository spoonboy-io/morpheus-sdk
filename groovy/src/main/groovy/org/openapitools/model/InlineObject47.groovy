package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class InlineObject47 {
    /* Refresh Mode. Run the `daily` or `costing` job instead of the default `hourly` refresh job. */
    String mode
    /* Rebuild. Pass `true` to purge existing invoices for the period before refreshing. Only applies to mode `costing`. */
    String rebuild
    /* Period. Invoice billing period to refresh in the format `YYYYMM`. The default period is the current month. Only applies to mode `costing`. */
    String period
}
