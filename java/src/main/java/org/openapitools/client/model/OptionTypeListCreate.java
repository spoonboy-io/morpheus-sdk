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
import org.openapitools.client.model.OptionTypeListCreateConfig;
import org.openapitools.client.model.OptionTypeListCreateCredential;

/**
 * OptionTypeListCreate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class OptionTypeListCreate {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_LABELS = "labels";
  @SerializedName(SERIALIZED_NAME_LABELS)
  private List<String> labels = null;

  /**
   * Option List Type eg. &#x60;rest&#x60;, &#x60;api&#x60;, &#x60;ldap&#x60; or &#x60;manual&#x60;.
   */
  @JsonAdapter(TypeEnum.Adapter.class)
  public enum TypeEnum {
    REST("rest"),
    
    API("api"),
    
    LDAP("ldap"),
    
    MANUAL("manual");

    private String value;

    TypeEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static TypeEnum fromValue(String value) {
      for (TypeEnum b : TypeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<TypeEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final TypeEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public TypeEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return TypeEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private TypeEnum type = TypeEnum.REST;

  public static final String SERIALIZED_NAME_SOURCE_URL = "sourceUrl";
  @SerializedName(SERIALIZED_NAME_SOURCE_URL)
  private String sourceUrl;

  /**
   * Visibility
   */
  @JsonAdapter(VisibilityEnum.Adapter.class)
  public enum VisibilityEnum {
    PRIVATE("private"),
    
    PUBLIC("public");

    private String value;

    VisibilityEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static VisibilityEnum fromValue(String value) {
      for (VisibilityEnum b : VisibilityEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<VisibilityEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final VisibilityEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public VisibilityEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return VisibilityEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_VISIBILITY = "visibility";
  @SerializedName(SERIALIZED_NAME_VISIBILITY)
  private VisibilityEnum visibility = VisibilityEnum.PRIVATE;

  /**
   * Source Method, the HTTP method.
   */
  @JsonAdapter(SourceMethodEnum.Adapter.class)
  public enum SourceMethodEnum {
    GET("GET"),
    
    POST("POST");

    private String value;

    SourceMethodEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static SourceMethodEnum fromValue(String value) {
      for (SourceMethodEnum b : SourceMethodEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<SourceMethodEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final SourceMethodEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public SourceMethodEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return SourceMethodEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_SOURCE_METHOD = "sourceMethod";
  @SerializedName(SERIALIZED_NAME_SOURCE_METHOD)
  private SourceMethodEnum sourceMethod = SourceMethodEnum.GET;

  public static final String SERIALIZED_NAME_API_TYPE = "apiType";
  @SerializedName(SERIALIZED_NAME_API_TYPE)
  private String apiType;

  public static final String SERIALIZED_NAME_IGNORE_S_S_L_ERRORS = "ignoreSSLErrors";
  @SerializedName(SERIALIZED_NAME_IGNORE_S_S_L_ERRORS)
  private Boolean ignoreSSLErrors = false;

  public static final String SERIALIZED_NAME_REAL_TIME = "realTime";
  @SerializedName(SERIALIZED_NAME_REAL_TIME)
  private Boolean realTime = false;

  public static final String SERIALIZED_NAME_CREDENTIAL = "credential";
  @SerializedName(SERIALIZED_NAME_CREDENTIAL)
  private OptionTypeListCreateCredential credential;

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

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private OptionTypeListCreateConfig config;


  public OptionTypeListCreate name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Name
   * @return name
  **/
  @ApiModelProperty(required = true, value = "Name")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public OptionTypeListCreate description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Description")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public OptionTypeListCreate labels(List<String> labels) {
    
    this.labels = labels;
    return this;
  }

  public OptionTypeListCreate addLabelsItem(String labelsItem) {
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


  public OptionTypeListCreate type(TypeEnum type) {
    
    this.type = type;
    return this;
  }

   /**
   * Option List Type eg. &#x60;rest&#x60;, &#x60;api&#x60;, &#x60;ldap&#x60; or &#x60;manual&#x60;.
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Option List Type eg. `rest`, `api`, `ldap` or `manual`.")

  public TypeEnum getType() {
    return type;
  }


  public void setType(TypeEnum type) {
    this.type = type;
  }


  public OptionTypeListCreate sourceUrl(String sourceUrl) {
    
    this.sourceUrl = sourceUrl;
    return this;
  }

   /**
   * Source URL, the http(s) URL to request data from. Required when type is rest.
   * @return sourceUrl
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Source URL, the http(s) URL to request data from. Required when type is rest.")

  public String getSourceUrl() {
    return sourceUrl;
  }


  public void setSourceUrl(String sourceUrl) {
    this.sourceUrl = sourceUrl;
  }


  public OptionTypeListCreate visibility(VisibilityEnum visibility) {
    
    this.visibility = visibility;
    return this;
  }

   /**
   * Visibility
   * @return visibility
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Visibility")

  public VisibilityEnum getVisibility() {
    return visibility;
  }


  public void setVisibility(VisibilityEnum visibility) {
    this.visibility = visibility;
  }


  public OptionTypeListCreate sourceMethod(SourceMethodEnum sourceMethod) {
    
    this.sourceMethod = sourceMethod;
    return this;
  }

   /**
   * Source Method, the HTTP method.
   * @return sourceMethod
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Source Method, the HTTP method.")

  public SourceMethodEnum getSourceMethod() {
    return sourceMethod;
  }


  public void setSourceMethod(SourceMethodEnum sourceMethod) {
    this.sourceMethod = sourceMethod;
  }


  public OptionTypeListCreate apiType(String apiType) {
    
    this.apiType = apiType;
    return this;
  }

   /**
   * Api Type, The code of the api option list to use, eg. clouds, environments, groups, instances, instance-wiki, networks, servicePlans, resourcePools, securityGroups, servers, server-wiki. Required when type is api.
   * @return apiType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Api Type, The code of the api option list to use, eg. clouds, environments, groups, instances, instance-wiki, networks, servicePlans, resourcePools, securityGroups, servers, server-wiki. Required when type is api.")

  public String getApiType() {
    return apiType;
  }


  public void setApiType(String apiType) {
    this.apiType = apiType;
  }


  public OptionTypeListCreate ignoreSSLErrors(Boolean ignoreSSLErrors) {
    
    this.ignoreSSLErrors = ignoreSSLErrors;
    return this;
  }

   /**
   * Ignore SSL Errors.
   * @return ignoreSSLErrors
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Ignore SSL Errors.")

  public Boolean getIgnoreSSLErrors() {
    return ignoreSSLErrors;
  }


  public void setIgnoreSSLErrors(Boolean ignoreSSLErrors) {
    this.ignoreSSLErrors = ignoreSSLErrors;
  }


  public OptionTypeListCreate realTime(Boolean realTime) {
    
    this.realTime = realTime;
    return this;
  }

   /**
   * Real Time.
   * @return realTime
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Real Time.")

  public Boolean getRealTime() {
    return realTime;
  }


  public void setRealTime(Boolean realTime) {
    this.realTime = realTime;
  }


  public OptionTypeListCreate credential(OptionTypeListCreateCredential credential) {
    
    this.credential = credential;
    return this;
  }

   /**
   * Get credential
   * @return credential
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OptionTypeListCreateCredential getCredential() {
    return credential;
  }


  public void setCredential(OptionTypeListCreateCredential credential) {
    this.credential = credential;
  }


  public OptionTypeListCreate serviceUsername(String serviceUsername) {
    
    this.serviceUsername = serviceUsername;
    return this;
  }

   /**
   * Username for authenticating with Basic Auth when type is rest or ldap username.
   * @return serviceUsername
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Username for authenticating with Basic Auth when type is rest or ldap username.")

  public String getServiceUsername() {
    return serviceUsername;
  }


  public void setServiceUsername(String serviceUsername) {
    this.serviceUsername = serviceUsername;
  }


  public OptionTypeListCreate servicePassword(String servicePassword) {
    
    this.servicePassword = servicePassword;
    return this;
  }

   /**
   * Password for authenticating with Basic Auth when type is rest or ldap password.
   * @return servicePassword
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Password for authenticating with Basic Auth when type is rest or ldap password.")

  public String getServicePassword() {
    return servicePassword;
  }


  public void setServicePassword(String servicePassword) {
    this.servicePassword = servicePassword;
  }


  public OptionTypeListCreate initialDataset(String initialDataset) {
    
    this.initialDataset = initialDataset;
    return this;
  }

   /**
   * Initial Dataset. Create an initial JSON or CSV dataset to be used as the collection for this option list. It should be a list containing objects with properties &#39;name&#39;, and &#39;value&#39;. Required when type is manual.
   * @return initialDataset
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Initial Dataset. Create an initial JSON or CSV dataset to be used as the collection for this option list. It should be a list containing objects with properties 'name', and 'value'. Required when type is manual.")

  public String getInitialDataset() {
    return initialDataset;
  }


  public void setInitialDataset(String initialDataset) {
    this.initialDataset = initialDataset;
  }


  public OptionTypeListCreate translationScript(String translationScript) {
    
    this.translationScript = translationScript;
    return this;
  }

   /**
   * Translation Script. Create a js script to translate the result data object into an Array containing objects with properties &#39;name&#39; and &#39;value&#39;. The input data is provided as data and the result should be put on the global variable results.
   * @return translationScript
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Translation Script. Create a js script to translate the result data object into an Array containing objects with properties 'name' and 'value'. The input data is provided as data and the result should be put on the global variable results.")

  public String getTranslationScript() {
    return translationScript;
  }


  public void setTranslationScript(String translationScript) {
    this.translationScript = translationScript;
  }


  public OptionTypeListCreate requestScript(String requestScript) {
    
    this.requestScript = requestScript;
    return this;
  }

   /**
   * Request Script. Create a js script to prepare the request. Return a data object as the body for a post, and return an array containing properties &#39;name&#39; and &#39;value&#39; for a get. The input data is provided as data and the result should be put on the global variable results.
   * @return requestScript
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Request Script. Create a js script to prepare the request. Return a data object as the body for a post, and return an array containing properties 'name' and 'value' for a get. The input data is provided as data and the result should be put on the global variable results.")

  public String getRequestScript() {
    return requestScript;
  }


  public void setRequestScript(String requestScript) {
    this.requestScript = requestScript;
  }


  public OptionTypeListCreate config(OptionTypeListCreateConfig config) {
    
    this.config = config;
    return this;
  }

   /**
   * Get config
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OptionTypeListCreateConfig getConfig() {
    return config;
  }


  public void setConfig(OptionTypeListCreateConfig config) {
    this.config = config;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    OptionTypeListCreate optionTypeListCreate = (OptionTypeListCreate) o;
    return Objects.equals(this.name, optionTypeListCreate.name) &&
        Objects.equals(this.description, optionTypeListCreate.description) &&
        Objects.equals(this.labels, optionTypeListCreate.labels) &&
        Objects.equals(this.type, optionTypeListCreate.type) &&
        Objects.equals(this.sourceUrl, optionTypeListCreate.sourceUrl) &&
        Objects.equals(this.visibility, optionTypeListCreate.visibility) &&
        Objects.equals(this.sourceMethod, optionTypeListCreate.sourceMethod) &&
        Objects.equals(this.apiType, optionTypeListCreate.apiType) &&
        Objects.equals(this.ignoreSSLErrors, optionTypeListCreate.ignoreSSLErrors) &&
        Objects.equals(this.realTime, optionTypeListCreate.realTime) &&
        Objects.equals(this.credential, optionTypeListCreate.credential) &&
        Objects.equals(this.serviceUsername, optionTypeListCreate.serviceUsername) &&
        Objects.equals(this.servicePassword, optionTypeListCreate.servicePassword) &&
        Objects.equals(this.initialDataset, optionTypeListCreate.initialDataset) &&
        Objects.equals(this.translationScript, optionTypeListCreate.translationScript) &&
        Objects.equals(this.requestScript, optionTypeListCreate.requestScript) &&
        Objects.equals(this.config, optionTypeListCreate.config);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, description, labels, type, sourceUrl, visibility, sourceMethod, apiType, ignoreSSLErrors, realTime, credential, serviceUsername, servicePassword, initialDataset, translationScript, requestScript, config);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class OptionTypeListCreate {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    labels: ").append(toIndentedString(labels)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    sourceUrl: ").append(toIndentedString(sourceUrl)).append("\n");
    sb.append("    visibility: ").append(toIndentedString(visibility)).append("\n");
    sb.append("    sourceMethod: ").append(toIndentedString(sourceMethod)).append("\n");
    sb.append("    apiType: ").append(toIndentedString(apiType)).append("\n");
    sb.append("    ignoreSSLErrors: ").append(toIndentedString(ignoreSSLErrors)).append("\n");
    sb.append("    realTime: ").append(toIndentedString(realTime)).append("\n");
    sb.append("    credential: ").append(toIndentedString(credential)).append("\n");
    sb.append("    serviceUsername: ").append(toIndentedString(serviceUsername)).append("\n");
    sb.append("    servicePassword: ").append(toIndentedString(servicePassword)).append("\n");
    sb.append("    initialDataset: ").append(toIndentedString(initialDataset)).append("\n");
    sb.append("    translationScript: ").append(toIndentedString(translationScript)).append("\n");
    sb.append("    requestScript: ").append(toIndentedString(requestScript)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
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

