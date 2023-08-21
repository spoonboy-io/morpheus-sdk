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
import org.openapitools.client.model.InlineResponse20040AppDeployInstance;
import org.openapitools.client.model.OptionTypeListConfig;
import org.openapitools.client.model.OptionTypeListCredential;

/**
 * OptionTypeList
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class OptionTypeList {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels = null;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;

  public static final String SERIALIZED_NAME_SOURCE_URL = "sourceUrl";
  @SerializedName(SERIALIZED_NAME_SOURCE_URL)
  private String sourceUrl;

  public static final String SERIALIZED_NAME_SOURCE_METHOD = "sourceMethod";
  @SerializedName(SERIALIZED_NAME_SOURCE_METHOD)
  private String sourceMethod;

  public static final String SERIALIZED_NAME_API_TYPE = "apiType";
  @SerializedName(SERIALIZED_NAME_API_TYPE)
  private String apiType;

  public static final String SERIALIZED_NAME_IGNORE_S_S_L_ERRORS = "ignoreSSLErrors";
  @SerializedName(SERIALIZED_NAME_IGNORE_S_S_L_ERRORS)
  private Boolean ignoreSSLErrors;

  public static final String SERIALIZED_NAME_REAL_TIME = "realTime";
  @SerializedName(SERIALIZED_NAME_REAL_TIME)
  private Boolean realTime;

  public static final String SERIALIZED_NAME_VISIBILITY = "visibility";
  @SerializedName(SERIALIZED_NAME_VISIBILITY)
  private String visibility;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private OptionTypeListConfig config;

  public static final String SERIALIZED_NAME_CREDENTIAL = "credential";
  @SerializedName(SERIALIZED_NAME_CREDENTIAL)
  private OptionTypeListCredential credential;

  public static final String SERIALIZED_NAME_SERVICE_USERNAME = "serviceUsername";
  @SerializedName(SERIALIZED_NAME_SERVICE_USERNAME)
  private String serviceUsername;

  public static final String SERIALIZED_NAME_SERVICE_PASSWORD = "servicePassword";
  @SerializedName(SERIALIZED_NAME_SERVICE_PASSWORD)
  private String servicePassword;

  public static final String SERIALIZED_NAME_INITIAL_DATASET = "initialDataset";
  @SerializedName(SERIALIZED_NAME_INITIAL_DATASET)
  private String initialDataset;

  public static final String SERIALIZED_NAME_TRANSLATION_SCRIPT = "translationScript";
  @SerializedName(SERIALIZED_NAME_TRANSLATION_SCRIPT)
  private String translationScript;

  public static final String SERIALIZED_NAME_REQUEST_SCRIPT = "requestScript";
  @SerializedName(SERIALIZED_NAME_REQUEST_SCRIPT)
  private String requestScript;

  public static final String SERIALIZED_NAME_ACCOUNT = "account";
  @SerializedName(SERIALIZED_NAME_ACCOUNT)
  private InlineResponse20040AppDeployInstance account;


  public OptionTypeList id(Long id) {
    
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


  public OptionTypeList name(String name) {
    
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


  public OptionTypeList description(String description) {
    
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


  public OptionTypeList labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public OptionTypeList addLabelsItem(String labelsItem) {
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


  public OptionTypeList type(String type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getType() {
    return type;
  }


  public void setType(String type) {
    this.type = type;
  }


  public OptionTypeList sourceUrl(String sourceUrl) {
    
    this.sourceUrl = sourceUrl;
    return this;
  }

   /**
   * Get sourceUrl
   * @return sourceUrl
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSourceUrl() {
    return sourceUrl;
  }


  public void setSourceUrl(String sourceUrl) {
    this.sourceUrl = sourceUrl;
  }


  public OptionTypeList sourceMethod(String sourceMethod) {
    
    this.sourceMethod = sourceMethod;
    return this;
  }

   /**
   * Get sourceMethod
   * @return sourceMethod
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSourceMethod() {
    return sourceMethod;
  }


  public void setSourceMethod(String sourceMethod) {
    this.sourceMethod = sourceMethod;
  }


  public OptionTypeList apiType(String apiType) {
    
    this.apiType = apiType;
    return this;
  }

   /**
   * Get apiType
   * @return apiType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getApiType() {
    return apiType;
  }


  public void setApiType(String apiType) {
    this.apiType = apiType;
  }


  public OptionTypeList ignoreSSLErrors(Boolean ignoreSSLErrors) {
    
    this.ignoreSSLErrors = ignoreSSLErrors;
    return this;
  }

   /**
   * Get ignoreSSLErrors
   * @return ignoreSSLErrors
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getIgnoreSSLErrors() {
    return ignoreSSLErrors;
  }


  public void setIgnoreSSLErrors(Boolean ignoreSSLErrors) {
    this.ignoreSSLErrors = ignoreSSLErrors;
  }


  public OptionTypeList realTime(Boolean realTime) {
    
    this.realTime = realTime;
    return this;
  }

   /**
   * Get realTime
   * @return realTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getRealTime() {
    return realTime;
  }


  public void setRealTime(Boolean realTime) {
    this.realTime = realTime;
  }


  public OptionTypeList visibility(String visibility) {
    
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


  public OptionTypeList config(OptionTypeListConfig config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OptionTypeListConfig getConfig() {
    return config;
  }


  public void setConfig(OptionTypeListConfig config) {
    this.config = config;
  }


  public OptionTypeList credential(OptionTypeListCredential credential) {
    
    this.credential = credential;
    return this;
  }

   /**
   * Get credential
   * @return credential
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OptionTypeListCredential getCredential() {
    return credential;
  }


  public void setCredential(OptionTypeListCredential credential) {
    this.credential = credential;
  }


  public OptionTypeList serviceUsername(String serviceUsername) {
    
    this.serviceUsername = serviceUsername;
    return this;
  }

   /**
   * Get serviceUsername
   * @return serviceUsername
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getServiceUsername() {
    return serviceUsername;
  }


  public void setServiceUsername(String serviceUsername) {
    this.serviceUsername = serviceUsername;
  }


  public OptionTypeList servicePassword(String servicePassword) {
    
    this.servicePassword = servicePassword;
    return this;
  }

   /**
   * Get servicePassword
   * @return servicePassword
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getServicePassword() {
    return servicePassword;
  }


  public void setServicePassword(String servicePassword) {
    this.servicePassword = servicePassword;
  }


  public OptionTypeList initialDataset(String initialDataset) {
    
    this.initialDataset = initialDataset;
    return this;
  }

   /**
   * Get initialDataset
   * @return initialDataset
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getInitialDataset() {
    return initialDataset;
  }


  public void setInitialDataset(String initialDataset) {
    this.initialDataset = initialDataset;
  }


  public OptionTypeList translationScript(String translationScript) {
    
    this.translationScript = translationScript;
    return this;
  }

   /**
   * Get translationScript
   * @return translationScript
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTranslationScript() {
    return translationScript;
  }


  public void setTranslationScript(String translationScript) {
    this.translationScript = translationScript;
  }


  public OptionTypeList requestScript(String requestScript) {
    
    this.requestScript = requestScript;
    return this;
  }

   /**
   * Get requestScript
   * @return requestScript
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRequestScript() {
    return requestScript;
  }


  public void setRequestScript(String requestScript) {
    this.requestScript = requestScript;
  }


  public OptionTypeList account(InlineResponse20040AppDeployInstance account) {
    
    this.account = account;
    return this;
  }

   /**
   * Get account
   * @return account
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20040AppDeployInstance getAccount() {
    return account;
  }


  public void setAccount(InlineResponse20040AppDeployInstance account) {
    this.account = account;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    OptionTypeList optionTypeList = (OptionTypeList) o;
    return Objects.equals(this.id, optionTypeList.id) &&
        Objects.equals(this.name, optionTypeList.name) &&
        Objects.equals(this.description, optionTypeList.description) &&
        Objects.equals(this.labels, optionTypeList.labels) &&
        Objects.equals(this.type, optionTypeList.type) &&
        Objects.equals(this.sourceUrl, optionTypeList.sourceUrl) &&
        Objects.equals(this.sourceMethod, optionTypeList.sourceMethod) &&
        Objects.equals(this.apiType, optionTypeList.apiType) &&
        Objects.equals(this.ignoreSSLErrors, optionTypeList.ignoreSSLErrors) &&
        Objects.equals(this.realTime, optionTypeList.realTime) &&
        Objects.equals(this.visibility, optionTypeList.visibility) &&
        Objects.equals(this.config, optionTypeList.config) &&
        Objects.equals(this.credential, optionTypeList.credential) &&
        Objects.equals(this.serviceUsername, optionTypeList.serviceUsername) &&
        Objects.equals(this.servicePassword, optionTypeList.servicePassword) &&
        Objects.equals(this.initialDataset, optionTypeList.initialDataset) &&
        Objects.equals(this.translationScript, optionTypeList.translationScript) &&
        Objects.equals(this.requestScript, optionTypeList.requestScript) &&
        Objects.equals(this.account, optionTypeList.account);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, description, labels, type, sourceUrl, sourceMethod, apiType, ignoreSSLErrors, realTime, visibility, config, credential, serviceUsername, servicePassword, initialDataset, translationScript, requestScript, account);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class OptionTypeList {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    sourceUrl: ").append(toIndentedString(sourceUrl)).append("\n");
    sb.append("    sourceMethod: ").append(toIndentedString(sourceMethod)).append("\n");
    sb.append("    apiType: ").append(toIndentedString(apiType)).append("\n");
    sb.append("    ignoreSSLErrors: ").append(toIndentedString(ignoreSSLErrors)).append("\n");
    sb.append("    realTime: ").append(toIndentedString(realTime)).append("\n");
    sb.append("    visibility: ").append(toIndentedString(visibility)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    credential: ").append(toIndentedString(credential)).append("\n");
    sb.append("    serviceUsername: ").append(toIndentedString(serviceUsername)).append("\n");
    sb.append("    servicePassword: ").append(toIndentedString(servicePassword)).append("\n");
    sb.append("    initialDataset: ").append(toIndentedString(initialDataset)).append("\n");
    sb.append("    translationScript: ").append(toIndentedString(translationScript)).append("\n");
    sb.append("    requestScript: ").append(toIndentedString(requestScript)).append("\n");
    sb.append("    account: ").append(toIndentedString(account)).append("\n");
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
