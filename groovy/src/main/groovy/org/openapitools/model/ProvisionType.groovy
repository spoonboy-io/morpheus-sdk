package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ControllerType;
import org.openapitools.model.NetworkType;
import org.openapitools.model.OptionType;
import org.openapitools.model.StorageType;

@Canonical
class ProvisionType {
    
    Long id
    
    String name
    
    String description
    
    String code
    
    Boolean aclEnabled
    
    Boolean multiTenant
    
    Boolean managed
    
    Boolean hostNetwork
    
    Boolean customSupported
    
    Boolean mapPorts
    
    Boolean exportServer
    
    String viewSet
    
    String serverType
    
    String hostType
    
    Boolean addVolumes
    
    Boolean hasVolumes
    
    Boolean hasDatastore
    
    Boolean hasNetworks
    
    Long maxNetworks
    
    Boolean customizeVolume
    
    Boolean rootDiskCustomizable
    
    Boolean rootDiskSizeKnown
    
    Boolean rootDiskResizable
    
    Boolean lvmSupported
    
    String hostDiskMode
    
    Long minDisk
    
    String maxDisk
    
    Boolean resizeCopiesVolumes
    
    Boolean supportsAutoDatastore
    
    Boolean hasZonePools
    
    Boolean hasSecurityGroups
    
    Boolean hasParameters
    
    Boolean canEnforceTags
    
    Boolean disableRootDatastore
    
    Boolean hasSnapshots
    
    Boolean hasSpecTemplates
    
    Boolean hasPreview
    
    Boolean zonePoolRequired
    
    Boolean planRequiresPool
    
    Boolean hasFolders
    
    List<OptionType> optionTypes = new ArrayList<OptionType>()
    
    List<OptionType> customOptionTypes = new ArrayList<OptionType>()
    
    List<NetworkType> networkTypes = new ArrayList<NetworkType>()
    
    List<StorageType> storageTypes = new ArrayList<StorageType>()
    
    List<StorageType> rootStorageTypes = new ArrayList<StorageType>()
    
    List<ControllerType> controllerTypes = new ArrayList<ControllerType>()
}
