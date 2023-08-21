/*
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

/**
 * InstancesConfigVMWare
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InstancesConfigVMWare {
  public static final String SERIALIZED_NAME_NO_AGENT = "noAgent";
  @SerializedName(SERIALIZED_NAME_NO_AGENT)
  private Boolean noAgent = false;

  public static final String SERIALIZED_NAME_RESOURCE_POOL_ID = "resourcePoolId";
  @SerializedName(SERIALIZED_NAME_RESOURCE_POOL_ID)
  private String resourcePoolId;

  public static final String SERIALIZED_NAME_HOST_ID = "hostId";
  @SerializedName(SERIALIZED_NAME_HOST_ID)
  private String hostId;

  public static final String SERIALIZED_NAME_SMBIOS_ASSET_TAG = "smbiosAssetTag";
  @SerializedName(SERIALIZED_NAME_SMBIOS_ASSET_TAG)
  private String smbiosAssetTag;

  /**
   * Enable Nested Virtualization
   */
  @JsonAdapter(NestedVirtualizationEnum.Adapter.class)
  public enum NestedVirtualizationEnum {
    ON("on"),
    
    OFF("off");

    private String value;

    NestedVirtualizationEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static NestedVirtualizationEnum fromValue(String value) {
      for (NestedVirtualizationEnum b : NestedVirtualizationEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<NestedVirtualizationEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final NestedVirtualizationEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public NestedVirtualizationEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return NestedVirtualizationEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_NESTED_VIRTUALIZATION = "nestedVirtualization";
  @SerializedName(SERIALIZED_NAME_NESTED_VIRTUALIZATION)
  private NestedVirtualizationEnum nestedVirtualization = NestedVirtualizationEnum.OFF;

  public static final String SERIALIZED_NAME_VMWARE_FOLDER_ID = "vmwareFolderId";
  @SerializedName(SERIALIZED_NAME_VMWARE_FOLDER_ID)
  private String vmwareFolderId;


  public InstancesConfigVMWare noAgent(Boolean noAgent) {
    
    this.noAgent = noAgent;
    return this;
  }

   /**
   * Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected.
   * @return noAgent
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected.")

  public Boolean getNoAgent() {
    return noAgent;
  }


  public void setNoAgent(Boolean noAgent) {
    this.noAgent = noAgent;
  }


  public InstancesConfigVMWare resourcePoolId(String resourcePoolId) {
    
    this.resourcePoolId = resourcePoolId;
    return this;
  }

   /**
   * id of the resource group to be used, can be prefixed with &#x60;pool-&#x60;. A resource pool group can be specified instead by prefixing its ID with &#x60;poolGroup-&#x60;.
   * @return resourcePoolId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "id of the resource group to be used, can be prefixed with `pool-`. A resource pool group can be specified instead by prefixing its ID with `poolGroup-`.")

  public String getResourcePoolId() {
    return resourcePoolId;
  }


  public void setResourcePoolId(String resourcePoolId) {
    this.resourcePoolId = resourcePoolId;
  }


  public InstancesConfigVMWare hostId(String hostId) {
    
    this.hostId = hostId;
    return this;
  }

   /**
   * Specific host to deploy to if so desired.
   * @return hostId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Specific host to deploy to if so desired.")

  public String getHostId() {
    return hostId;
  }


  public void setHostId(String hostId) {
    this.hostId = hostId;
  }


  public InstancesConfigVMWare smbiosAssetTag(String smbiosAssetTag) {
    
    this.smbiosAssetTag = smbiosAssetTag;
    return this;
  }

   /**
   * Sets the asset tag on the SMBIOS for use by the guest operating system. If left blank, the virtual machine name will be used.
   * @return smbiosAssetTag
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Sets the asset tag on the SMBIOS for use by the guest operating system. If left blank, the virtual machine name will be used.")

  public String getSmbiosAssetTag() {
    return smbiosAssetTag;
  }


  public void setSmbiosAssetTag(String smbiosAssetTag) {
    this.smbiosAssetTag = smbiosAssetTag;
  }


  public InstancesConfigVMWare nestedVirtualization(NestedVirtualizationEnum nestedVirtualization) {
    
    this.nestedVirtualization = nestedVirtualization;
    return this;
  }

   /**
   * Enable Nested Virtualization
   * @return nestedVirtualization
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Enable Nested Virtualization")

  public NestedVirtualizationEnum getNestedVirtualization() {
    return nestedVirtualization;
  }


  public void setNestedVirtualization(NestedVirtualizationEnum nestedVirtualization) {
    this.nestedVirtualization = nestedVirtualization;
  }


  public InstancesConfigVMWare vmwareFolderId(String vmwareFolderId) {
    
    this.vmwareFolderId = vmwareFolderId;
    return this;
  }

   /**
   * VMWare Folder External ID (as a String) or ID (as an Integer or String)
   * @return vmwareFolderId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "VMWare Folder External ID (as a String) or ID (as an Integer or String)")

  public String getVmwareFolderId() {
    return vmwareFolderId;
  }


  public void setVmwareFolderId(String vmwareFolderId) {
    this.vmwareFolderId = vmwareFolderId;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InstancesConfigVMWare instancesConfigVMWare = (InstancesConfigVMWare) o;
    return Objects.equals(this.noAgent, instancesConfigVMWare.noAgent) &&
        Objects.equals(this.resourcePoolId, instancesConfigVMWare.resourcePoolId) &&
        Objects.equals(this.hostId, instancesConfigVMWare.hostId) &&
        Objects.equals(this.smbiosAssetTag, instancesConfigVMWare.smbiosAssetTag) &&
        Objects.equals(this.nestedVirtualization, instancesConfigVMWare.nestedVirtualization) &&
        Objects.equals(this.vmwareFolderId, instancesConfigVMWare.vmwareFolderId);
  }

  @Override
  public int hashCode() {
    return Objects.hash(noAgent, resourcePoolId, hostId, smbiosAssetTag, nestedVirtualization, vmwareFolderId);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InstancesConfigVMWare {\n");
    sb.append("    noAgent: ").append(toIndentedString(noAgent)).append("\n");
    sb.append("    resourcePoolId: ").append(toIndentedString(resourcePoolId)).append("\n");
    sb.append("    hostId: ").append(toIndentedString(hostId)).append("\n");
    sb.append("    smbiosAssetTag: ").append(toIndentedString(smbiosAssetTag)).append("\n");
    sb.append("    nestedVirtualization: ").append(toIndentedString(nestedVirtualization)).append("\n");
    sb.append("    vmwareFolderId: ").append(toIndentedString(vmwareFolderId)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

