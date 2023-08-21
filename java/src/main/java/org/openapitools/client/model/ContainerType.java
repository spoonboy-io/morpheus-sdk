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
import org.openapitools.client.model.ContainerTypeAccount;
import org.openapitools.client.model.ContainerTypeContainerPorts;
import org.openapitools.client.model.ContainerTypeProvisionType;

/**
 * ContainerType
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class ContainerType {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Integer id;

  public static final String SERIALIZED_NAME_ACCOUNT = "account";
  @SerializedName(SERIALIZED_NAME_ACCOUNT)
  private ContainerTypeAccount account;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels = null;

  public static final String SERIALIZED_NAME_SHORT_NAME = "shortName";
  @SerializedName(SERIALIZED_NAME_SHORT_NAME)
  private String shortName;

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private String code;

  public static final String SERIALIZED_NAME_CONTAINER_VERSION = "containerVersion";
  @SerializedName(SERIALIZED_NAME_CONTAINER_VERSION)
  private String containerVersion;

  public static final String SERIALIZED_NAME_PROVISION_TYPE = "provisionType";
  @SerializedName(SERIALIZED_NAME_PROVISION_TYPE)
  private ContainerTypeProvisionType provisionType;

  public static final String SERIALIZED_NAME_VIRTUAL_IMAGE = "virtualImage";
  @SerializedName(SERIALIZED_NAME_VIRTUAL_IMAGE)
  private ContainerTypeAccount virtualImage;

  public static final String SERIALIZED_NAME_CATEGORY = "category";
  @SerializedName(SERIALIZED_NAME_CATEGORY)
  private String category;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private Object config;

  public static final String SERIALIZED_NAME_CONTAINER_PORTS = "containerPorts";
  @SerializedName(SERIALIZED_NAME_CONTAINER_PORTS)
  private List<ContainerTypeContainerPorts> containerPorts = null;

  public static final String SERIALIZED_NAME_CONTAINER_SCRIPTS = "containerScripts";
  @SerializedName(SERIALIZED_NAME_CONTAINER_SCRIPTS)
  private List<Object> containerScripts = null;

  public static final String SERIALIZED_NAME_CONTAINER_TEMPLATES = "containerTemplates";
  @SerializedName(SERIALIZED_NAME_CONTAINER_TEMPLATES)
  private List<Object> containerTemplates = null;

  public static final String SERIALIZED_NAME_ENVIRONMENT_VARIABLES = "environmentVariables";
  @SerializedName(SERIALIZED_NAME_ENVIRONMENT_VARIABLES)
  private List<Object> environmentVariables = null;


  public ContainerType id(Integer id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getId() {
    return id;
  }


  public void setId(Integer id) {
    this.id = id;
  }


  public ContainerType account(ContainerTypeAccount account) {
    
    this.account = account;
    return this;
  }

   /**
   * Get account
   * @return account
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ContainerTypeAccount getAccount() {
    return account;
  }


  public void setAccount(ContainerTypeAccount account) {
    this.account = account;
  }


  public ContainerType name(String name) {
    
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


  public ContainerType labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public ContainerType addLabelsItem(String labelsItem) {
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


  public ContainerType shortName(String shortName) {
    
    this.shortName = shortName;
    return this;
  }

   /**
   * Get shortName
   * @return shortName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getShortName() {
    return shortName;
  }


  public void setShortName(String shortName) {
    this.shortName = shortName;
  }


  public ContainerType code(String code) {
    
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


  public ContainerType containerVersion(String containerVersion) {
    
    this.containerVersion = containerVersion;
    return this;
  }

   /**
   * Get containerVersion
   * @return containerVersion
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getContainerVersion() {
    return containerVersion;
  }


  public void setContainerVersion(String containerVersion) {
    this.containerVersion = containerVersion;
  }


  public ContainerType provisionType(ContainerTypeProvisionType provisionType) {
    
    this.provisionType = provisionType;
    return this;
  }

   /**
   * Get provisionType
   * @return provisionType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ContainerTypeProvisionType getProvisionType() {
    return provisionType;
  }


  public void setProvisionType(ContainerTypeProvisionType provisionType) {
    this.provisionType = provisionType;
  }


  public ContainerType virtualImage(ContainerTypeAccount virtualImage) {
    
    this.virtualImage = virtualImage;
    return this;
  }

   /**
   * Get virtualImage
   * @return virtualImage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public ContainerTypeAccount getVirtualImage() {
    return virtualImage;
  }


  public void setVirtualImage(ContainerTypeAccount virtualImage) {
    this.virtualImage = virtualImage;
  }


  public ContainerType category(String category) {
    
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


  public ContainerType config(Object config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Object getConfig() {
    return config;
  }


  public void setConfig(Object config) {
    this.config = config;
  }


  public ContainerType containerPorts(List<ContainerTypeContainerPorts> containerPorts) {
    
    this.containerPorts = containerPorts;
    return this;
  }

  public ContainerType addContainerPortsItem(ContainerTypeContainerPorts containerPortsItem) {
    if (this.containerPorts == null) {
      this.containerPorts = new ArrayList<ContainerTypeContainerPorts>();
    }
    this.containerPorts.add(containerPortsItem);
    return this;
  }

   /**
   * Get containerPorts
   * @return containerPorts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<ContainerTypeContainerPorts> getContainerPorts() {
    return containerPorts;
  }


  public void setContainerPorts(List<ContainerTypeContainerPorts> containerPorts) {
    this.containerPorts = containerPorts;
  }


  public ContainerType containerScripts(List<Object> containerScripts) {
    
    this.containerScripts = containerScripts;
    return this;
  }

  public ContainerType addContainerScriptsItem(Object containerScriptsItem) {
    if (this.containerScripts == null) {
      this.containerScripts = new ArrayList<Object>();
    }
    this.containerScripts.add(containerScriptsItem);
    return this;
  }

   /**
   * Get containerScripts
   * @return containerScripts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getContainerScripts() {
    return containerScripts;
  }


  public void setContainerScripts(List<Object> containerScripts) {
    this.containerScripts = containerScripts;
  }


  public ContainerType containerTemplates(List<Object> containerTemplates) {
    
    this.containerTemplates = containerTemplates;
    return this;
  }

  public ContainerType addContainerTemplatesItem(Object containerTemplatesItem) {
    if (this.containerTemplates == null) {
      this.containerTemplates = new ArrayList<Object>();
    }
    this.containerTemplates.add(containerTemplatesItem);
    return this;
  }

   /**
   * Get containerTemplates
   * @return containerTemplates
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getContainerTemplates() {
    return containerTemplates;
  }


  public void setContainerTemplates(List<Object> containerTemplates) {
    this.containerTemplates = containerTemplates;
  }


  public ContainerType environmentVariables(List<Object> environmentVariables) {
    
    this.environmentVariables = environmentVariables;
    return this;
  }

  public ContainerType addEnvironmentVariablesItem(Object environmentVariablesItem) {
    if (this.environmentVariables == null) {
      this.environmentVariables = new ArrayList<Object>();
    }
    this.environmentVariables.add(environmentVariablesItem);
    return this;
  }

   /**
   * Get environmentVariables
   * @return environmentVariables
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getEnvironmentVariables() {
    return environmentVariables;
  }


  public void setEnvironmentVariables(List<Object> environmentVariables) {
    this.environmentVariables = environmentVariables;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    ContainerType containerType = (ContainerType) o;
    return Objects.equals(this.id, containerType.id) &&
        Objects.equals(this.account, containerType.account) &&
        Objects.equals(this.name, containerType.name) &&
        Objects.equals(this.labels, containerType.labels) &&
        Objects.equals(this.shortName, containerType.shortName) &&
        Objects.equals(this.code, containerType.code) &&
        Objects.equals(this.containerVersion, containerType.containerVersion) &&
        Objects.equals(this.provisionType, containerType.provisionType) &&
        Objects.equals(this.virtualImage, containerType.virtualImage) &&
        Objects.equals(this.category, containerType.category) &&
        Objects.equals(this.config, containerType.config) &&
        Objects.equals(this.containerPorts, containerType.containerPorts) &&
        Objects.equals(this.containerScripts, containerType.containerScripts) &&
        Objects.equals(this.containerTemplates, containerType.containerTemplates) &&
        Objects.equals(this.environmentVariables, containerType.environmentVariables);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, account, name, labels, shortName, code, containerVersion, provisionType, virtualImage, category, config, containerPorts, containerScripts, containerTemplates, environmentVariables);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class ContainerType {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    account: ").append(toIndentedString(account)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    shortName: ").append(toIndentedString(shortName)).append("\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
    sb.append("    containerVersion: ").append(toIndentedString(containerVersion)).append("\n");
    sb.append("    provisionType: ").append(toIndentedString(provisionType)).append("\n");
    sb.append("    virtualImage: ").append(toIndentedString(virtualImage)).append("\n");
    sb.append("    category: ").append(toIndentedString(category)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    containerPorts: ").append(toIndentedString(containerPorts)).append("\n");
    sb.append("    containerScripts: ").append(toIndentedString(containerScripts)).append("\n");
    sb.append("    containerTemplates: ").append(toIndentedString(containerTemplates)).append("\n");
    sb.append("    environmentVariables: ").append(toIndentedString(environmentVariables)).append("\n");
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

