package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InstanceServicePlanAutoOptions;
import org.openapitools.model.ServerServicePlansDatastores;

@Canonical
class ServerServicePlans {
    
    Long id
    
    String name
    
    Long value
    
    String code
    
    Long maxStorage
    
    Long maxMemory
    
    Long maxCpu
    
    Long maxCores
    
    Long maxDataStorage
    
    Boolean customCpu
    
    Boolean customMaxMemory
    
    Boolean customMaxStorage
    
    Boolean customMaxDataStorage
    
    Boolean customCoresPerSocket
    
    Long coresPerSocket
    
    List<Object> storageTypes = new ArrayList<Object>()
    
    List<Object> rootStorageTypes = new ArrayList<Object>()
    
    Boolean addVolumes
    
    Boolean customizeVolume
    
    Boolean rootDiskCustomizable
    
    String hostDiskMode
    
    String hasDatastore
    
    String lvmSupported
    
    String minDisk
    
    String maxDisk
    
    ServerServicePlansDatastores datastores
    
    Boolean supportsAutoDatastore
    
    List<InstanceServicePlanAutoOptions> autoOptions = new ArrayList<InstanceServicePlanAutoOptions>()
    
    List<Object> cpuOptions = new ArrayList<Object>()
    
    List<Object> memoryOptions = new ArrayList<Object>()
    
    Object rootCustomSizeOptions
    
    Object customSizeOptions
    
    Boolean customCores
    
    String maxDisks
    
    String memorySizeType
}
