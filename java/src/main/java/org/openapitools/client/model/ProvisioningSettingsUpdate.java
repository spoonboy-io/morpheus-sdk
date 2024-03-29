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
import org.openapitools.client.model.ProvisioningSettingsUpdateCloudInitKeyPair;
import org.openapitools.client.model.ProvisioningSettingsUpdateDefaultTemplateType;
import org.openapitools.client.model.ProvisioningSettingsUpdateDeployStorageProvider;

/**
 * ProvisioningSettingsUpdate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ProvisioningSettingsUpdate {
  public static final String SERIALIZED_NAME_ALLOW_ZONE_SELECTION = "allowZoneSelection";
  @SerializedName(SERIALIZED_NAME_ALLOW_ZONE_SELECTION)
  private Boolean allowZoneSelection;

  public static final String SERIALIZED_NAME_ALLOW_SERVER_SELECTION = "allowServerSelection";
  @SerializedName(SERIALIZED_NAME_ALLOW_SERVER_SELECTION)
  private Boolean allowServerSelection;

  public static final String SERIALIZED_NAME_REQUIRE_ENVIRONMENTS = "requireEnvironments";
  @SerializedName(SERIALIZED_NAME_REQUIRE_ENVIRONMENTS)
  private Boolean requireEnvironments;

  public static final String SERIALIZED_NAME_SHOW_PRICING = "showPricing";
  @SerializedName(SERIALIZED_NAME_SHOW_PRICING)
  private Boolean showPricing;

  public static final String SERIALIZED_NAME_HIDE_DATASTORE_STATS = "hideDatastoreStats";
  @SerializedName(SERIALIZED_NAME_HIDE_DATASTORE_STATS)
  private Boolean hideDatastoreStats;

  public static final String SERIALIZED_NAME_CROSS_TENANT_NAMING_POLICIES = "crossTenantNamingPolicies";
  @SerializedName(SERIALIZED_NAME_CROSS_TENANT_NAMING_POLICIES)
  private Boolean crossTenantNamingPolicies;

  public static final String SERIALIZED_NAME_REUSE_SEQUENCE = "reuseSequence";
  @SerializedName(SERIALIZED_NAME_REUSE_SEQUENCE)
  private Boolean reuseSequence;

  public static final String SERIALIZED_NAME_CLOUD_INIT_USERNAME = "cloudInitUsername";
  @SerializedName(SERIALIZED_NAME_CLOUD_INIT_USERNAME)
  private String cloudInitUsername;

  public static final String SERIALIZED_NAME_CLOUD_INIT_PASSWORD = "cloudInitPassword";
  @SerializedName(SERIALIZED_NAME_CLOUD_INIT_PASSWORD)
  private String cloudInitPassword;

  public static final String SERIALIZED_NAME_CLOUD_INIT_KEY_PAIR = "cloudInitKeyPair";
  @SerializedName(SERIALIZED_NAME_CLOUD_INIT_KEY_PAIR)
  private ProvisioningSettingsUpdateCloudInitKeyPair cloudInitKeyPair;

  public static final String SERIALIZED_NAME_DEPLOY_STORAGE_PROVIDER = "deployStorageProvider";
  @SerializedName(SERIALIZED_NAME_DEPLOY_STORAGE_PROVIDER)
  private ProvisioningSettingsUpdateDeployStorageProvider deployStorageProvider;

  public static final String SERIALIZED_NAME_WINDOWS_PASSWORD = "windowsPassword";
  @SerializedName(SERIALIZED_NAME_WINDOWS_PASSWORD)
  private String windowsPassword;

  public static final String SERIALIZED_NAME_PXE_ROOT_PASSWORD = "pxeRootPassword";
  @SerializedName(SERIALIZED_NAME_PXE_ROOT_PASSWORD)
  private String pxeRootPassword;

  public static final String SERIALIZED_NAME_DEFAULT_TEMPLATE_TYPE = "defaultTemplateType";
  @SerializedName(SERIALIZED_NAME_DEFAULT_TEMPLATE_TYPE)
  private ProvisioningSettingsUpdateDefaultTemplateType defaultTemplateType;


  public ProvisioningSettingsUpdate allowZoneSelection(Boolean allowZoneSelection) {
    
    this.allowZoneSelection = allowZoneSelection;
    return this;
  }

   /**
   * Use this to enable / disable allowing cloud selection
   * @return allowZoneSelection
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Use this to enable / disable allowing cloud selection")

  public Boolean getAllowZoneSelection() {
    return allowZoneSelection;
  }


  public void setAllowZoneSelection(Boolean allowZoneSelection) {
    this.allowZoneSelection = allowZoneSelection;
  }


  public ProvisioningSettingsUpdate allowServerSelection(Boolean allowServerSelection) {
    
    this.allowServerSelection = allowServerSelection;
    return this;
  }

   /**
   * Use this to enable / disable allowing host selection
   * @return allowServerSelection
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Use this to enable / disable allowing host selection")

  public Boolean getAllowServerSelection() {
    return allowServerSelection;
  }


  public void setAllowServerSelection(Boolean allowServerSelection) {
    this.allowServerSelection = allowServerSelection;
  }


  public ProvisioningSettingsUpdate requireEnvironments(Boolean requireEnvironments) {
    
    this.requireEnvironments = requireEnvironments;
    return this;
  }

   /**
   * Use this to enable / disable requiring environment selection
   * @return requireEnvironments
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Use this to enable / disable requiring environment selection")

  public Boolean getRequireEnvironments() {
    return requireEnvironments;
  }


  public void setRequireEnvironments(Boolean requireEnvironments) {
    this.requireEnvironments = requireEnvironments;
  }


  public ProvisioningSettingsUpdate showPricing(Boolean showPricing) {
    
    this.showPricing = showPricing;
    return this;
  }

   /**
   * Use this to enable / disable showing pricing
   * @return showPricing
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Use this to enable / disable showing pricing")

  public Boolean getShowPricing() {
    return showPricing;
  }


  public void setShowPricing(Boolean showPricing) {
    this.showPricing = showPricing;
  }


  public ProvisioningSettingsUpdate hideDatastoreStats(Boolean hideDatastoreStats) {
    
    this.hideDatastoreStats = hideDatastoreStats;
    return this;
  }

   /**
   * Use this to enable / disable hiding datastore stats
   * @return hideDatastoreStats
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Use this to enable / disable hiding datastore stats")

  public Boolean getHideDatastoreStats() {
    return hideDatastoreStats;
  }


  public void setHideDatastoreStats(Boolean hideDatastoreStats) {
    this.hideDatastoreStats = hideDatastoreStats;
  }


  public ProvisioningSettingsUpdate crossTenantNamingPolicies(Boolean crossTenantNamingPolicies) {
    
    this.crossTenantNamingPolicies = crossTenantNamingPolicies;
    return this;
  }

   /**
   * Use this to enable / disable cross-tenant naming policies
   * @return crossTenantNamingPolicies
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Use this to enable / disable cross-tenant naming policies")

  public Boolean getCrossTenantNamingPolicies() {
    return crossTenantNamingPolicies;
  }


  public void setCrossTenantNamingPolicies(Boolean crossTenantNamingPolicies) {
    this.crossTenantNamingPolicies = crossTenantNamingPolicies;
  }


  public ProvisioningSettingsUpdate reuseSequence(Boolean reuseSequence) {
    
    this.reuseSequence = reuseSequence;
    return this;
  }

   /**
   * Use this to enable / disable reusing naming sequence numbers
   * @return reuseSequence
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Use this to enable / disable reusing naming sequence numbers")

  public Boolean getReuseSequence() {
    return reuseSequence;
  }


  public void setReuseSequence(Boolean reuseSequence) {
    this.reuseSequence = reuseSequence;
  }


  public ProvisioningSettingsUpdate cloudInitUsername(String cloudInitUsername) {
    
    this.cloudInitUsername = cloudInitUsername;
    return this;
  }

   /**
   * Cloud-init username
   * @return cloudInitUsername
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Cloud-init username")

  public String getCloudInitUsername() {
    return cloudInitUsername;
  }


  public void setCloudInitUsername(String cloudInitUsername) {
    this.cloudInitUsername = cloudInitUsername;
  }


  public ProvisioningSettingsUpdate cloudInitPassword(String cloudInitPassword) {
    
    this.cloudInitPassword = cloudInitPassword;
    return this;
  }

   /**
   * Cloud-init password
   * @return cloudInitPassword
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Cloud-init password")

  public String getCloudInitPassword() {
    return cloudInitPassword;
  }


  public void setCloudInitPassword(String cloudInitPassword) {
    this.cloudInitPassword = cloudInitPassword;
  }


  public ProvisioningSettingsUpdate cloudInitKeyPair(ProvisioningSettingsUpdateCloudInitKeyPair cloudInitKeyPair) {
    
    this.cloudInitKeyPair = cloudInitKeyPair;
    return this;
  }

   /**
   * Get cloudInitKeyPair
   * @return cloudInitKeyPair
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ProvisioningSettingsUpdateCloudInitKeyPair getCloudInitKeyPair() {
    return cloudInitKeyPair;
  }


  public void setCloudInitKeyPair(ProvisioningSettingsUpdateCloudInitKeyPair cloudInitKeyPair) {
    this.cloudInitKeyPair = cloudInitKeyPair;
  }


  public ProvisioningSettingsUpdate deployStorageProvider(ProvisioningSettingsUpdateDeployStorageProvider deployStorageProvider) {
    
    this.deployStorageProvider = deployStorageProvider;
    return this;
  }

   /**
   * Get deployStorageProvider
   * @return deployStorageProvider
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ProvisioningSettingsUpdateDeployStorageProvider getDeployStorageProvider() {
    return deployStorageProvider;
  }


  public void setDeployStorageProvider(ProvisioningSettingsUpdateDeployStorageProvider deployStorageProvider) {
    this.deployStorageProvider = deployStorageProvider;
  }


  public ProvisioningSettingsUpdate windowsPassword(String windowsPassword) {
    
    this.windowsPassword = windowsPassword;
    return this;
  }

   /**
   * Windows administrator password
   * @return windowsPassword
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Windows administrator password")

  public String getWindowsPassword() {
    return windowsPassword;
  }


  public void setWindowsPassword(String windowsPassword) {
    this.windowsPassword = windowsPassword;
  }


  public ProvisioningSettingsUpdate pxeRootPassword(String pxeRootPassword) {
    
    this.pxeRootPassword = pxeRootPassword;
    return this;
  }

   /**
   * PXE Boot default root password
   * @return pxeRootPassword
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "PXE Boot default root password")

  public String getPxeRootPassword() {
    return pxeRootPassword;
  }


  public void setPxeRootPassword(String pxeRootPassword) {
    this.pxeRootPassword = pxeRootPassword;
  }


  public ProvisioningSettingsUpdate defaultTemplateType(ProvisioningSettingsUpdateDefaultTemplateType defaultTemplateType) {
    
    this.defaultTemplateType = defaultTemplateType;
    return this;
  }

   /**
   * Get defaultTemplateType
   * @return defaultTemplateType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ProvisioningSettingsUpdateDefaultTemplateType getDefaultTemplateType() {
    return defaultTemplateType;
  }


  public void setDefaultTemplateType(ProvisioningSettingsUpdateDefaultTemplateType defaultTemplateType) {
    this.defaultTemplateType = defaultTemplateType;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ProvisioningSettingsUpdate provisioningSettingsUpdate = (ProvisioningSettingsUpdate) o;
    return Objects.equals(this.allowZoneSelection, provisioningSettingsUpdate.allowZoneSelection) &&
        Objects.equals(this.allowServerSelection, provisioningSettingsUpdate.allowServerSelection) &&
        Objects.equals(this.requireEnvironments, provisioningSettingsUpdate.requireEnvironments) &&
        Objects.equals(this.showPricing, provisioningSettingsUpdate.showPricing) &&
        Objects.equals(this.hideDatastoreStats, provisioningSettingsUpdate.hideDatastoreStats) &&
        Objects.equals(this.crossTenantNamingPolicies, provisioningSettingsUpdate.crossTenantNamingPolicies) &&
        Objects.equals(this.reuseSequence, provisioningSettingsUpdate.reuseSequence) &&
        Objects.equals(this.cloudInitUsername, provisioningSettingsUpdate.cloudInitUsername) &&
        Objects.equals(this.cloudInitPassword, provisioningSettingsUpdate.cloudInitPassword) &&
        Objects.equals(this.cloudInitKeyPair, provisioningSettingsUpdate.cloudInitKeyPair) &&
        Objects.equals(this.deployStorageProvider, provisioningSettingsUpdate.deployStorageProvider) &&
        Objects.equals(this.windowsPassword, provisioningSettingsUpdate.windowsPassword) &&
        Objects.equals(this.pxeRootPassword, provisioningSettingsUpdate.pxeRootPassword) &&
        Objects.equals(this.defaultTemplateType, provisioningSettingsUpdate.defaultTemplateType);
  }

  @Override
  public int hashCode() {
    return Objects.hash(allowZoneSelection, allowServerSelection, requireEnvironments, showPricing, hideDatastoreStats, crossTenantNamingPolicies, reuseSequence, cloudInitUsername, cloudInitPassword, cloudInitKeyPair, deployStorageProvider, windowsPassword, pxeRootPassword, defaultTemplateType);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ProvisioningSettingsUpdate {\n");
    sb.append("    allowZoneSelection: ").append(toIndentedString(allowZoneSelection)).append("\n");
    sb.append("    allowServerSelection: ").append(toIndentedString(allowServerSelection)).append("\n");
    sb.append("    requireEnvironments: ").append(toIndentedString(requireEnvironments)).append("\n");
    sb.append("    showPricing: ").append(toIndentedString(showPricing)).append("\n");
    sb.append("    hideDatastoreStats: ").append(toIndentedString(hideDatastoreStats)).append("\n");
    sb.append("    crossTenantNamingPolicies: ").append(toIndentedString(crossTenantNamingPolicies)).append("\n");
    sb.append("    reuseSequence: ").append(toIndentedString(reuseSequence)).append("\n");
    sb.append("    cloudInitUsername: ").append(toIndentedString(cloudInitUsername)).append("\n");
    sb.append("    cloudInitPassword: ").append(toIndentedString(cloudInitPassword)).append("\n");
    sb.append("    cloudInitKeyPair: ").append(toIndentedString(cloudInitKeyPair)).append("\n");
    sb.append("    deployStorageProvider: ").append(toIndentedString(deployStorageProvider)).append("\n");
    sb.append("    windowsPassword: ").append(toIndentedString(windowsPassword)).append("\n");
    sb.append("    pxeRootPassword: ").append(toIndentedString(pxeRootPassword)).append("\n");
    sb.append("    defaultTemplateType: ").append(toIndentedString(defaultTemplateType)).append("\n");
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

