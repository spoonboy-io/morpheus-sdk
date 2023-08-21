package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ImageBuildCreateBootScript;
import org.openapitools.model.ImageBuildCreatePreseedScript;
import org.openapitools.model.ImageBuildCreateSite;
import org.openapitools.model.ImageBuildCreateZone;

@Canonical
class ImageBuildCreate {
    /* A name for the image build */
    String name
    /* A description for the image build */
    String description
    /* The image builder type. */
    String type
    
    ImageBuildCreateSite site
    
    ImageBuildCreateZone zone
    /* A map of config values. This is the instance config that is used for provisioning. See Provisioning Types. */
    Object config
    
    ImageBuildCreateBootScript bootScript
    
    ImageBuildCreatePreseedScript preseedScript
    /* SSH Username */
    String sshUsername
    /* SSH Password */
    String sshPassword
    /* Storage Provider */
    String storageProvider
    /* Cloud Init */
    String isCloudInit
    /* Build Output Name */
    String buildOutputName
    /* Conversion Formats */
    String conversionFormats
    /* Keep Results - Keep only the most recent builds. Older executions will be deleted along with their associated Virtual Images. The value 0 disables this functionality. */
    Long keepResults = 0l
}
