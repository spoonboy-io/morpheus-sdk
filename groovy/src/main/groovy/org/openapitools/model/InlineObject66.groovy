package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.OneOfstringlong;

@Canonical
class InlineObject66 {
    /* Time to Live. The lease duration in seconds, or a human readable format eg. 15m, 8h, 7d. The default is 0 meaning Never expires. This only is applied if the cypher does not yet exist and is created.  */
    OneOfstringlong ttl = null
    /* The secret value to be stored. This is only parsed if type is passed as `string`. */
    String value
}
