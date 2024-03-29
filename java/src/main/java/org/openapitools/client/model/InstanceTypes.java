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
import org.openapitools.client.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.openapitools.client.model.InstanceTypesInstanceTypeLayouts;

/**
 * InstanceTypes
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InstanceTypes {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_ACCOUNT = "account";
  @SerializedName(SERIALIZED_NAME_ACCOUNT)
  private InlineResponse20082LoadBalancerInstanceSslCert account;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels = null;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_PROVISION_TYPE_CODE = "provisionTypeCode";
  @SerializedName(SERIALIZED_NAME_PROVISION_TYPE_CODE)
  private String provisionTypeCode;

  public static final String SERIALIZED_NAME_CATEGORY = "category";
  @SerializedName(SERIALIZED_NAME_CATEGORY)
  private String category;

  public static final String SERIALIZED_NAME_ACTIVE = "active";
  @SerializedName(SERIALIZED_NAME_ACTIVE)
  private Boolean active;

  public static final String SERIALIZED_NAME_ENVIRONMENT_PREFIX = "environmentPrefix";
  @SerializedName(SERIALIZED_NAME_ENVIRONMENT_PREFIX)
  private String environmentPrefix;

  public static final String SERIALIZED_NAME_VISIBILITY = "visibility";
  @SerializedName(SERIALIZED_NAME_VISIBILITY)
  private String visibility;

  public static final String SERIALIZED_NAME_FEATURED = "featured";
  @SerializedName(SERIALIZED_NAME_FEATURED)
  private Boolean featured;

  public static final String SERIALIZED_NAME_VERSIONS = "versions";
  @SerializedName(SERIALIZED_NAME_VERSIONS)
  private List<String> versions = null;

  public static final String SERIALIZED_NAME_INSTANCE_TYPE_LAYOUTS = "instanceTypeLayouts";
  @SerializedName(SERIALIZED_NAME_INSTANCE_TYPE_LAYOUTS)
  private List<InstanceTypesInstanceTypeLayouts> instanceTypeLayouts = null;

  public static final String SERIALIZED_NAME_IMAGE_PATH = "imagePath";
  @SerializedName(SERIALIZED_NAME_IMAGE_PATH)
  private String imagePath;

  public static final String SERIALIZED_NAME_DARK_IMAGE_PATH = "darkImagePath";
  @SerializedName(SERIALIZED_NAME_DARK_IMAGE_PATH)
  private String darkImagePath;


  public InstanceTypes id(Long id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getId() {
    return id;
  }


  public void setId(Long id) {
    this.id = id;
  }


  public InstanceTypes account(InlineResponse20082LoadBalancerInstanceSslCert account) {
    
    this.account = account;
    return this;
  }

   /**
   * Get account
   * @return account
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20082LoadBalancerInstanceSslCert getAccount() {
    return account;
  }


  public void setAccount(InlineResponse20082LoadBalancerInstanceSslCert account) {
    this.account = account;
  }


  public InstanceTypes name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public InstanceTypes labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public InstanceTypes addLabelsItem(String labelsItem) {
    if (this.labels == null) {
      this.labels = new ArrayList<String>();
    }
    this.labels.add(labelsItem);
    return this;
  }

   /**
   * Get labels
   * @return labels
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getLabels() {
    return labels;
  }


  public void setLabels(List<String> labels) {
    this.labels = labels;
  }


  public InstanceTypes code(String code) {
    
    this.code = code;
    return this;
  }

   /**
   * Get code
   * @return code
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCode() {
    return code;
  }


  public void setCode(String code) {
    this.code = code;
  }


  public InstanceTypes description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public InstanceTypes provisionTypeCode(String provisionTypeCode) {
    
    this.provisionTypeCode = provisionTypeCode;
    return this;
  }

   /**
   * Get provisionTypeCode
   * @return provisionTypeCode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getProvisionTypeCode() {
    return provisionTypeCode;
  }


  public void setProvisionTypeCode(String provisionTypeCode) {
    this.provisionTypeCode = provisionTypeCode;
  }


  public InstanceTypes category(String category) {
    
    this.category = category;
    return this;
  }

   /**
   * Get category
   * @return category
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCategory() {
    return category;
  }


  public void setCategory(String category) {
    this.category = category;
  }


  public InstanceTypes active(Boolean active) {
    
    this.active = active;
    return this;
  }

   /**
   * Get active
   * @return active
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getActive() {
    return active;
  }


  public void setActive(Boolean active) {
    this.active = active;
  }


  public InstanceTypes environmentPrefix(String environmentPrefix) {
    
    this.environmentPrefix = environmentPrefix;
    return this;
  }

   /**
   * Get environmentPrefix
   * @return environmentPrefix
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getEnvironmentPrefix() {
    return environmentPrefix;
  }


  public void setEnvironmentPrefix(String environmentPrefix) {
    this.environmentPrefix = environmentPrefix;
  }


  public InstanceTypes visibility(String visibility) {
    
    this.visibility = visibility;
    return this;
  }

   /**
   * Get visibility
   * @return visibility
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVisibility() {
    return visibility;
  }


  public void setVisibility(String visibility) {
    this.visibility = visibility;
  }


  public InstanceTypes featured(Boolean featured) {
    
    this.featured = featured;
    return this;
  }

   /**
   * Get featured
   * @return featured
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getFeatured() {
    return featured;
  }


  public void setFeatured(Boolean featured) {
    this.featured = featured;
  }


  public InstanceTypes versions(List<String> versions) {
    
    this.versions = versions;
    return this;
  }

  public InstanceTypes addVersionsItem(String versionsItem) {
    if (this.versions == null) {
      this.versions = new ArrayList<String>();
    }
    this.versions.add(versionsItem);
    return this;
  }

   /**
   * Get versions
   * @return versions
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getVersions() {
    return versions;
  }


  public void setVersions(List<String> versions) {
    this.versions = versions;
  }


  public InstanceTypes instanceTypeLayouts(List<InstanceTypesInstanceTypeLayouts> instanceTypeLayouts) {
    
    this.instanceTypeLayouts = instanceTypeLayouts;
    return this;
  }

  public InstanceTypes addInstanceTypeLayoutsItem(InstanceTypesInstanceTypeLayouts instanceTypeLayoutsItem) {
    if (this.instanceTypeLayouts == null) {
      this.instanceTypeLayouts = new ArrayList<InstanceTypesInstanceTypeLayouts>();
    }
    this.instanceTypeLayouts.add(instanceTypeLayoutsItem);
    return this;
  }

   /**
   * Get instanceTypeLayouts
   * @return instanceTypeLayouts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<InstanceTypesInstanceTypeLayouts> getInstanceTypeLayouts() {
    return instanceTypeLayouts;
  }


  public void setInstanceTypeLayouts(List<InstanceTypesInstanceTypeLayouts> instanceTypeLayouts) {
    this.instanceTypeLayouts = instanceTypeLayouts;
  }


  public InstanceTypes imagePath(String imagePath) {
    
    this.imagePath = imagePath;
    return this;
  }

   /**
   * Logo image URL
   * @return imagePath
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Logo image URL")

  public String getImagePath() {
    return imagePath;
  }


  public void setImagePath(String imagePath) {
    this.imagePath = imagePath;
  }


  public InstanceTypes darkImagePath(String darkImagePath) {
    
    this.darkImagePath = darkImagePath;
    return this;
  }

   /**
   * Dark logo image URL
   * @return darkImagePath
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Dark logo image URL")

  public String getDarkImagePath() {
    return darkImagePath;
  }


  public void setDarkImagePath(String darkImagePath) {
    this.darkImagePath = darkImagePath;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InstanceTypes instanceTypes = (InstanceTypes) o;
    return Objects.equals(this.id, instanceTypes.id) &&
        Objects.equals(this.account, instanceTypes.account) &&
        Objects.equals(this.name, instanceTypes.name) &&
        Objects.equals(this.labels, instanceTypes.labels) &&
        Objects.equals(this.code, instanceTypes.code) &&
        Objects.equals(this.description, instanceTypes.description) &&
        Objects.equals(this.provisionTypeCode, instanceTypes.provisionTypeCode) &&
        Objects.equals(this.category, instanceTypes.category) &&
        Objects.equals(this.active, instanceTypes.active) &&
        Objects.equals(this.environmentPrefix, instanceTypes.environmentPrefix) &&
        Objects.equals(this.visibility, instanceTypes.visibility) &&
        Objects.equals(this.featured, instanceTypes.featured) &&
        Objects.equals(this.versions, instanceTypes.versions) &&
        Objects.equals(this.instanceTypeLayouts, instanceTypes.instanceTypeLayouts) &&
        Objects.equals(this.imagePath, instanceTypes.imagePath) &&
        Objects.equals(this.darkImagePath, instanceTypes.darkImagePath);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, account, name, labels, code, description, provisionTypeCode, category, active, environmentPrefix, visibility, featured, versions, instanceTypeLayouts, imagePath, darkImagePath);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InstanceTypes {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    account: ").append(toIndentedString(account)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    provisionTypeCode: ").append(toIndentedString(provisionTypeCode)).append("\n");
    sb.append("    category: ").append(toIndentedString(category)).append("\n");
    sb.append("    active: ").append(toIndentedString(active)).append("\n");
    sb.append("    environmentPrefix: ").append(toIndentedString(environmentPrefix)).append("\n");
    sb.append("    visibility: ").append(toIndentedString(visibility)).append("\n");
    sb.append("    featured: ").append(toIndentedString(featured)).append("\n");
    sb.append("    versions: ").append(toIndentedString(versions)).append("\n");
    sb.append("    instanceTypeLayouts: ").append(toIndentedString(instanceTypeLayouts)).append("\n");
    sb.append("    imagePath: ").append(toIndentedString(imagePath)).append("\n");
    sb.append("    darkImagePath: ").append(toIndentedString(darkImagePath)).append("\n");
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

