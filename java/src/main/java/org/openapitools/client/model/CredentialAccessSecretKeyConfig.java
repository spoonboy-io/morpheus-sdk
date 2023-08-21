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
import org.openapitools.client.model.CredentialAccessSecretKeyConfigIntegration;

/**
 * CredentialAccessSecretKeyConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class CredentialAccessSecretKeyConfig {
  /**
   * Credential Type Code
   */
  @JsonAdapter(TypeEnum.Adapter.class)
  public enum TypeEnum {
    ACCESS_KEY_SECRET("access-key-secret");

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
  private TypeEnum type;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled = true;

  public static final String SERIALIZED_NAME_INTEGRATION = "integration";
  @SerializedName(SERIALIZED_NAME_INTEGRATION)
  private CredentialAccessSecretKeyConfigIntegration integration;

  public static final String SERIALIZED_NAME_USERNAME = "username";
  @SerializedName(SERIALIZED_NAME_USERNAME)
  private String username;

  public static final String SERIALIZED_NAME_PASSWORD = "password";
  @SerializedName(SERIALIZED_NAME_PASSWORD)
  private String password;


  public CredentialAccessSecretKeyConfig type(TypeEnum type) {
    
    this.type = type;
    return this;
  }

   /**
   * Credential Type Code
   * @return type
  **/
  @ApiModelProperty(required = true, value = "Credential Type Code")

  public TypeEnum getType() {
    return type;
  }


  public void setType(TypeEnum type) {
    this.type = type;
  }


  public CredentialAccessSecretKeyConfig name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * A unique name scoped to your account for the credential
   * @return name
  **/
  @ApiModelProperty(required = true, value = "A unique name scoped to your account for the credential")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public CredentialAccessSecretKeyConfig description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Optional Description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Optional Description")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public CredentialAccessSecretKeyConfig enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Credential enabled
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Credential enabled")

  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  public CredentialAccessSecretKeyConfig integration(CredentialAccessSecretKeyConfigIntegration integration) {
    
    this.integration = integration;
    return this;
  }

   /**
   * Get integration
   * @return integration
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public CredentialAccessSecretKeyConfigIntegration getIntegration() {
    return integration;
  }


  public void setIntegration(CredentialAccessSecretKeyConfigIntegration integration) {
    this.integration = integration;
  }


  public CredentialAccessSecretKeyConfig username(String username) {
    
    this.username = username;
    return this;
  }

   /**
   * Access Key
   * @return username
  **/
  @ApiModelProperty(example = "72c54d9b-1c73-4c73-91b9-1fb895f5fe5a", required = true, value = "Access Key")

  public String getUsername() {
    return username;
  }


  public void setUsername(String username) {
    this.username = username;
  }


  public CredentialAccessSecretKeyConfig password(String password) {
    
    this.password = password;
    return this;
  }

   /**
   * Secret Key
   * @return password
  **/
  @ApiModelProperty(example = "2b3450f3-b563-4a5f-ba96-91af0212fd69", required = true, value = "Secret Key")

  public String getPassword() {
    return password;
  }


  public void setPassword(String password) {
    this.password = password;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CredentialAccessSecretKeyConfig credentialAccessSecretKeyConfig = (CredentialAccessSecretKeyConfig) o;
    return Objects.equals(this.type, credentialAccessSecretKeyConfig.type) &&
        Objects.equals(this.name, credentialAccessSecretKeyConfig.name) &&
        Objects.equals(this.description, credentialAccessSecretKeyConfig.description) &&
        Objects.equals(this.enabled, credentialAccessSecretKeyConfig.enabled) &&
        Objects.equals(this.integration, credentialAccessSecretKeyConfig.integration) &&
        Objects.equals(this.username, credentialAccessSecretKeyConfig.username) &&
        Objects.equals(this.password, credentialAccessSecretKeyConfig.password);
  }

  @Override
  public int hashCode() {
    return Objects.hash(type, name, description, enabled, integration, username, password);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CredentialAccessSecretKeyConfig {\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    integration: ").append(toIndentedString(integration)).append("\n");
    sb.append("    username: ").append(toIndentedString(username)).append("\n");
    sb.append("    password: ").append(toIndentedString(password)).append("\n");
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

