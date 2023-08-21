package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class NetworkPoolServerUpdatePhpIpamConfig {
    /* App ID (Your App name in phpIPAM) */
    String appId
    /* Inventory Existing */
    String inventoryExisting = InventoryExistingEnum.OFF
}
