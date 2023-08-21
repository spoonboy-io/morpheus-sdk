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
import java.util.ArrayList;
import java.util.List;
import org.openapitools.client.model.OneOfobjectobject;
import org.openapitools.client.model.OneOfobjectstring;
import org.openapitools.client.model.VirtualImageCreateStorageProvider;
import org.openapitools.client.model.VirtualImageCreateTags;

/**
 * VirtualImageCreate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class VirtualImageCreate {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels = null;

  public static final String SERIALIZED_NAME_IMAGE_TYPE = "imageType";
  @SerializedName(SERIALIZED_NAME_IMAGE_TYPE)
  private String imageType;

  public static final String SERIALIZED_NAME_STORAGE_PROVIDER = "storageProvider";
  @SerializedName(SERIALIZED_NAME_STORAGE_PROVIDER)
  private VirtualImageCreateStorageProvider storageProvider;

  public static final String SERIALIZED_NAME_IS_CLOUD_INIT = "isCloudInit";
  @SerializedName(SERIALIZED_NAME_IS_CLOUD_INIT)
  private Boolean isCloudInit = false;

  public static final String SERIALIZED_NAME_USER_DATA = "userData";
  @SerializedName(SERIALIZED_NAME_USER_DATA)
  private String userData;

  public static final String SERIALIZED_NAME_INSTALL_AGENT = "installAgent";
  @SerializedName(SERIALIZED_NAME_INSTALL_AGENT)
  private Boolean installAgent = false;

  public static final String SERIALIZED_NAME_SSH_USERNAME = "sshUsername";
  @SerializedName(SERIALIZED_NAME_SSH_USERNAME)
  private String sshUsername;

  public static final String SERIALIZED_NAME_SSH_PASSWORD = "sshPassword";
  @SerializedName(SERIALIZED_NAME_SSH_PASSWORD)
  private String sshPassword;

  public static final String SERIALIZED_NAME_SSH_KEY = "sshKey";
  @SerializedName(SERIALIZED_NAME_SSH_KEY)
  private String sshKey;

  public static final String SERIALIZED_NAME_OS_TYPE = "osType";
  @SerializedName(SERIALIZED_NAME_OS_TYPE)
  private OneOfobjectstring osType = null;

  public static final String SERIALIZED_NAME_VISIBILITY = "visibility";
  @SerializedName(SERIALIZED_NAME_VISIBILITY)
  private String visibility = "private";

  public static final String SERIALIZED_NAME_ACCOUNTS = "accounts";
  @SerializedName(SERIALIZED_NAME_ACCOUNTS)
  private List<Long> accounts = null;

  public static final String SERIALIZED_NAME_IS_AUTO_JOIN_DOMAIN = "isAutoJoinDomain";
  @SerializedName(SERIALIZED_NAME_IS_AUTO_JOIN_DOMAIN)
  private Boolean isAutoJoinDomain = false;

  public static final String SERIALIZED_NAME_VIRTIO_SUPPORTED = "virtioSupported";
  @SerializedName(SERIALIZED_NAME_VIRTIO_SUPPORTED)
  private Boolean virtioSupported = true;

  public static final String SERIALIZED_NAME_VM_TOOLS_INSTALLED = "vmToolsInstalled";
  @SerializedName(SERIALIZED_NAME_VM_TOOLS_INSTALLED)
  private Boolean vmToolsInstalled = true;

  public static final String SERIALIZED_NAME_IS_FORCE_CUSTOMIZATION = "isForceCustomization";
  @SerializedName(SERIALIZED_NAME_IS_FORCE_CUSTOMIZATION)
  private Boolean isForceCustomization = false;

  public static final String SERIALIZED_NAME_TRIAL_VERSION = "trialVersion";
  @SerializedName(SERIALIZED_NAME_TRIAL_VERSION)
  private Boolean trialVersion = false;

  public static final String SERIALIZED_NAME_IS_SYSPREP = "isSysprep";
  @SerializedName(SERIALIZED_NAME_IS_SYSPREP)
  private Boolean isSysprep = false;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private OneOfobjectobject config = null;

  public static final String SERIALIZED_NAME_TAGS = "tags";
  @SerializedName(SERIALIZED_NAME_TAGS)
  private List<VirtualImageCreateTags> tags = null;


  public VirtualImageCreate name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * A name for the virtual image
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "A name for the virtual image")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public VirtualImageCreate labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public VirtualImageCreate addLabelsItem(String labelsItem) {
    if (this.labels == null) {
      this.labels = new ArrayList<String>();
    }
    this.labels.add(labelsItem);
    return this;
  }

   /**
   * Array of label strings, can be used for filtering.
   * @return labels
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of label strings, can be used for filtering.")

  public List<String> getLabels() {
    return labels;
  }


  public void setLabels(List<String> labels) {
    this.labels = labels;
  }


  public VirtualImageCreate imageType(String imageType) {
    
    this.imageType = imageType;
    return this;
  }

   /**
   * Code of image type. eg. vmware, ami, etc.
   * @return imageType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Code of image type. eg. vmware, ami, etc.")

  public String getImageType() {
    return imageType;
  }


  public void setImageType(String imageType) {
    this.imageType = imageType;
  }


  public VirtualImageCreate storageProvider(VirtualImageCreateStorageProvider storageProvider) {
    
    this.storageProvider = storageProvider;
    return this;
  }

   /**
   * Get storageProvider
   * @return storageProvider
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public VirtualImageCreateStorageProvider getStorageProvider() {
    return storageProvider;
  }


  public void setStorageProvider(VirtualImageCreateStorageProvider storageProvider) {
    this.storageProvider = storageProvider;
  }


  public VirtualImageCreate isCloudInit(Boolean isCloudInit) {
    
    this.isCloudInit = isCloudInit;
    return this;
  }

   /**
   * Cloud Init Enabled?
   * @return isCloudInit
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Cloud Init Enabled?")

  public Boolean getIsCloudInit() {
    return isCloudInit;
  }


  public void setIsCloudInit(Boolean isCloudInit) {
    this.isCloudInit = isCloudInit;
  }


  public VirtualImageCreate userData(String userData) {
    
    this.userData = userData;
    return this;
  }

   /**
   * Cloud-Init User Data, a bash script
   * @return userData
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Cloud-Init User Data, a bash script")

  public String getUserData() {
    return userData;
  }


  public void setUserData(String userData) {
    this.userData = userData;
  }


  public VirtualImageCreate installAgent(Boolean installAgent) {
    
    this.installAgent = installAgent;
    return this;
  }

   /**
   * Install Agent?
   * @return installAgent
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Install Agent?")

  public Boolean getInstallAgent() {
    return installAgent;
  }


  public void setInstallAgent(Boolean installAgent) {
    this.installAgent = installAgent;
  }


  public VirtualImageCreate sshUsername(String sshUsername) {
    
    this.sshUsername = sshUsername;
    return this;
  }

   /**
   * SSH Username
   * @return sshUsername
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "SSH Username")

  public String getSshUsername() {
    return sshUsername;
  }


  public void setSshUsername(String sshUsername) {
    this.sshUsername = sshUsername;
  }


  public VirtualImageCreate sshPassword(String sshPassword) {
    
    this.sshPassword = sshPassword;
    return this;
  }

   /**
   * SSH Password
   * @return sshPassword
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "SSH Password")

  public String getSshPassword() {
    return sshPassword;
  }


  public void setSshPassword(String sshPassword) {
    this.sshPassword = sshPassword;
  }


  public VirtualImageCreate sshKey(String sshKey) {
    
    this.sshKey = sshKey;
    return this;
  }

   /**
   * SSH Key
   * @return sshKey
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "SSH Key")

  public String getSshKey() {
    return sshKey;
  }


  public void setSshKey(String sshKey) {
    this.sshKey = sshKey;
  }


  public VirtualImageCreate osType(OneOfobjectstring osType) {
    
    this.osType = osType;
    return this;
  }

   /**
   * A Map containing the id of the OS Type. This can also be passed as a string (code or name) instead.
   * @return osType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "A Map containing the id of the OS Type. This can also be passed as a string (code or name) instead.")

  public OneOfobjectstring getOsType() {
    return osType;
  }


  public void setOsType(OneOfobjectstring osType) {
    this.osType = osType;
  }


  public VirtualImageCreate visibility(String visibility) {
    
    this.visibility = visibility;
    return this;
  }

   /**
   * private or public
   * @return visibility
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "private or public")

  public String getVisibility() {
    return visibility;
  }


  public void setVisibility(String visibility) {
    this.visibility = visibility;
  }


  public VirtualImageCreate accounts(List<Long> accounts) {
    
    this.accounts = accounts;
    return this;
  }

  public VirtualImageCreate addAccountsItem(Long accountsItem) {
    if (this.accounts == null) {
      this.accounts = new ArrayList<Long>();
    }
    this.accounts.add(accountsItem);
    return this;
  }

   /**
   * Get accounts
   * @return accounts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Long> getAccounts() {
    return accounts;
  }


  public void setAccounts(List<Long> accounts) {
    this.accounts = accounts;
  }


  public VirtualImageCreate isAutoJoinDomain(Boolean isAutoJoinDomain) {
    
    this.isAutoJoinDomain = isAutoJoinDomain;
    return this;
  }

   /**
   * Auto Join Domain?
   * @return isAutoJoinDomain
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Auto Join Domain?")

  public Boolean getIsAutoJoinDomain() {
    return isAutoJoinDomain;
  }


  public void setIsAutoJoinDomain(Boolean isAutoJoinDomain) {
    this.isAutoJoinDomain = isAutoJoinDomain;
  }


  public VirtualImageCreate virtioSupported(Boolean virtioSupported) {
    
    this.virtioSupported = virtioSupported;
    return this;
  }

   /**
   * VirtIO Drivers Loaded?
   * @return virtioSupported
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "VirtIO Drivers Loaded?")

  public Boolean getVirtioSupported() {
    return virtioSupported;
  }


  public void setVirtioSupported(Boolean virtioSupported) {
    this.virtioSupported = virtioSupported;
  }


  public VirtualImageCreate vmToolsInstalled(Boolean vmToolsInstalled) {
    
    this.vmToolsInstalled = vmToolsInstalled;
    return this;
  }

   /**
   * VM Tools Installed?
   * @return vmToolsInstalled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "VM Tools Installed?")

  public Boolean getVmToolsInstalled() {
    return vmToolsInstalled;
  }


  public void setVmToolsInstalled(Boolean vmToolsInstalled) {
    this.vmToolsInstalled = vmToolsInstalled;
  }


  public VirtualImageCreate isForceCustomization(Boolean isForceCustomization) {
    
    this.isForceCustomization = isForceCustomization;
    return this;
  }

   /**
   * Force Guest Customization?
   * @return isForceCustomization
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Force Guest Customization?")

  public Boolean getIsForceCustomization() {
    return isForceCustomization;
  }


  public void setIsForceCustomization(Boolean isForceCustomization) {
    this.isForceCustomization = isForceCustomization;
  }


  public VirtualImageCreate trialVersion(Boolean trialVersion) {
    
    this.trialVersion = trialVersion;
    return this;
  }

   /**
   * Trial Version
   * @return trialVersion
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Trial Version")

  public Boolean getTrialVersion() {
    return trialVersion;
  }


  public void setTrialVersion(Boolean trialVersion) {
    this.trialVersion = trialVersion;
  }


  public VirtualImageCreate isSysprep(Boolean isSysprep) {
    
    this.isSysprep = isSysprep;
    return this;
  }

   /**
   * Sysprep Enabled?
   * @return isSysprep
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Sysprep Enabled?")

  public Boolean getIsSysprep() {
    return isSysprep;
  }


  public void setIsSysprep(Boolean isSysprep) {
    this.isSysprep = isSysprep;
  }


  public VirtualImageCreate config(OneOfobjectobject config) {
    
    this.config = config;
    return this;
  }

   /**
   * Map of configuration properties, varies by image type.
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Map of configuration properties, varies by image type.")

  public OneOfobjectobject getConfig() {
    return config;
  }


  public void setConfig(OneOfobjectobject config) {
    this.config = config;
  }


  public VirtualImageCreate tags(List<VirtualImageCreateTags> tags) {
    
    this.tags = tags;
    return this;
  }

  public VirtualImageCreate addTagsItem(VirtualImageCreateTags tagsItem) {
    if (this.tags == null) {
      this.tags = new ArrayList<VirtualImageCreateTags>();
    }
    this.tags.add(tagsItem);
    return this;
  }

   /**
   * Metadata tags, Array of objects having a name and value
   * @return tags
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Metadata tags, Array of objects having a name and value")

  public List<VirtualImageCreateTags> getTags() {
    return tags;
  }


  public void setTags(List<VirtualImageCreateTags> tags) {
    this.tags = tags;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    VirtualImageCreate virtualImageCreate = (VirtualImageCreate) o;
    return Objects.equals(this.name, virtualImageCreate.name) &&
        Objects.equals(this.labels, virtualImageCreate.labels) &&
        Objects.equals(this.imageType, virtualImageCreate.imageType) &&
        Objects.equals(this.storageProvider, virtualImageCreate.storageProvider) &&
        Objects.equals(this.isCloudInit, virtualImageCreate.isCloudInit) &&
        Objects.equals(this.userData, virtualImageCreate.userData) &&
        Objects.equals(this.installAgent, virtualImageCreate.installAgent) &&
        Objects.equals(this.sshUsername, virtualImageCreate.sshUsername) &&
        Objects.equals(this.sshPassword, virtualImageCreate.sshPassword) &&
        Objects.equals(this.sshKey, virtualImageCreate.sshKey) &&
        Objects.equals(this.osType, virtualImageCreate.osType) &&
        Objects.equals(this.visibility, virtualImageCreate.visibility) &&
        Objects.equals(this.accounts, virtualImageCreate.accounts) &&
        Objects.equals(this.isAutoJoinDomain, virtualImageCreate.isAutoJoinDomain) &&
        Objects.equals(this.virtioSupported, virtualImageCreate.virtioSupported) &&
        Objects.equals(this.vmToolsInstalled, virtualImageCreate.vmToolsInstalled) &&
        Objects.equals(this.isForceCustomization, virtualImageCreate.isForceCustomization) &&
        Objects.equals(this.trialVersion, virtualImageCreate.trialVersion) &&
        Objects.equals(this.isSysprep, virtualImageCreate.isSysprep) &&
        Objects.equals(this.config, virtualImageCreate.config) &&
        Objects.equals(this.tags, virtualImageCreate.tags);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, labels, imageType, storageProvider, isCloudInit, userData, installAgent, sshUsername, sshPassword, sshKey, osType, visibility, accounts, isAutoJoinDomain, virtioSupported, vmToolsInstalled, isForceCustomization, trialVersion, isSysprep, config, tags);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class VirtualImageCreate {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    imageType: ").append(toIndentedString(imageType)).append("\n");
    sb.append("    storageProvider: ").append(toIndentedString(storageProvider)).append("\n");
    sb.append("    isCloudInit: ").append(toIndentedString(isCloudInit)).append("\n");
    sb.append("    userData: ").append(toIndentedString(userData)).append("\n");
    sb.append("    installAgent: ").append(toIndentedString(installAgent)).append("\n");
    sb.append("    sshUsername: ").append(toIndentedString(sshUsername)).append("\n");
    sb.append("    sshPassword: ").append(toIndentedString(sshPassword)).append("\n");
    sb.append("    sshKey: ").append(toIndentedString(sshKey)).append("\n");
    sb.append("    osType: ").append(toIndentedString(osType)).append("\n");
    sb.append("    visibility: ").append(toIndentedString(visibility)).append("\n");
    sb.append("    accounts: ").append(toIndentedString(accounts)).append("\n");
    sb.append("    isAutoJoinDomain: ").append(toIndentedString(isAutoJoinDomain)).append("\n");
    sb.append("    virtioSupported: ").append(toIndentedString(virtioSupported)).append("\n");
    sb.append("    vmToolsInstalled: ").append(toIndentedString(vmToolsInstalled)).append("\n");
    sb.append("    isForceCustomization: ").append(toIndentedString(isForceCustomization)).append("\n");
    sb.append("    trialVersion: ").append(toIndentedString(trialVersion)).append("\n");
    sb.append("    isSysprep: ").append(toIndentedString(isSysprep)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    tags: ").append(toIndentedString(tags)).append("\n");
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

