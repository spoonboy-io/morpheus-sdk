package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class ApiContainersIdAttachFloatingIpConfig {
    /* The Floating IP identifier in the format: \"ip-ID\" or \"pool-ID\".  The Options API /api/options/openStack/openstackFloatingIpOptions?containerId={{containerId}} can be used to see which options are available.  */
    String osExternalNetworkId
    /* Bandwidth (Mbit/s) Only the following cloud types support this parameter: Huawei, OpenTelekom  */
    BigDecimal floatingIpBandwidth
}
