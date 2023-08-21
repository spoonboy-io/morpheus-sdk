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
import org.openapitools.client.model.AnyOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject;

/**
 * PolicyUpdate
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class PolicyUpdate {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private AnyOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject config = null;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled = true;

  /**
   * Scope object type
   */
  @JsonAdapter(RefTypeEnum.Adapter.class)
  public enum RefTypeEnum {
    COMPUTESITE("ComputeSite"),
    
    COMPUTEZONE("ComputeZone"),
    
    USER("User"),
    
    ROLE("Role"),
    
    NETWORK("Network"),
    
    PLAN("Plan");

    private String value;

    RefTypeEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static RefTypeEnum fromValue(String value) {
      for (RefTypeEnum b : RefTypeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<RefTypeEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final RefTypeEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public RefTypeEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return RefTypeEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_REF_TYPE = "refType";
  @SerializedName(SERIALIZED_NAME_REF_TYPE)
  private RefTypeEnum refType;

  public static final String SERIALIZED_NAME_REF_ID = "refId";
  @SerializedName(SERIALIZED_NAME_REF_ID)
  private Long refId;

  public static final String SERIALIZED_NAME_ACCOUNTS = "accounts";
  @SerializedName(SERIALIZED_NAME_ACCOUNTS)
  private List<Long> accounts = null;

  public static final String SERIALIZED_NAME_EACH_USER = "eachUser";
  @SerializedName(SERIALIZED_NAME_EACH_USER)
  private Boolean eachUser;


  public PolicyUpdate name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * A name for the policy
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(example = "Sample Policy", value = "A name for the policy")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public PolicyUpdate description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * A description for the policy
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "A description for the policy")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public PolicyUpdate config(AnyOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject config) {
    
    this.config = config;
    return this;
  }

   /**
   * A map of config values. The expected values vary by policy type. See &#x60;Retrieves all Policy Types&#x60; endpoint for &#x60;fieldName&#x60;(s) of required options.
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "A map of config values. The expected values vary by policy type. See `Retrieves all Policy Types` endpoint for `fieldName`(s) of required options.")

  public AnyOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject getConfig() {
    return config;
  }


  public void setConfig(AnyOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject config) {
    this.config = config;
  }


  public PolicyUpdate enabled(Boolean enabled) {
    
    this.enabled = enabled;
    return this;
  }

   /**
   * Set to false to disable
   * @return enabled
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Set to false to disable")

  public Boolean getEnabled() {
    return enabled;
  }


  public void setEnabled(Boolean enabled) {
    this.enabled = enabled;
  }


  public PolicyUpdate refType(RefTypeEnum refType) {
    
    this.refType = refType;
    return this;
  }

   /**
   * Scope object type
   * @return refType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Scope object type")

  public RefTypeEnum getRefType() {
    return refType;
  }


  public void setRefType(RefTypeEnum refType) {
    this.refType = refType;
  }


  public PolicyUpdate refId(Long refId) {
    
    this.refId = refId;
    return this;
  }

   /**
   * Scope object ID (&#x60;group&#x60;,&#x60;cloud&#x60;,&#x60;user&#x60;, etc)
   * @return refId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Scope object ID (`group`,`cloud`,`user`, etc)")

  public Long getRefId() {
    return refId;
  }


  public void setRefId(Long refId) {
    this.refId = refId;
  }


  public PolicyUpdate accounts(List<Long> accounts) {
    
    this.accounts = accounts;
    return this;
  }

  public PolicyUpdate addAccountsItem(Long accountsItem) {
    if (this.accounts == null) {
      this.accounts = new ArrayList<Long>();
    }
    this.accounts.add(accountsItem);
    return this;
  }

   /**
   * Array of tenants to scope the policy to
   * @return accounts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Array of tenants to scope the policy to")

  public List<Long> getAccounts() {
    return accounts;
  }


  public void setAccounts(List<Long> accounts) {
    this.accounts = accounts;
  }


  public PolicyUpdate eachUser(Boolean eachUser) {
    
    this.eachUser = eachUser;
    return this;
  }

   /**
   * Apply individually to each user in role.  Only when &#x60;refType&#x60; equals &#x60;Role&#x60;
   * @return eachUser
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Apply individually to each user in role.  Only when `refType` equals `Role`")

  public Boolean getEachUser() {
    return eachUser;
  }


  public void setEachUser(Boolean eachUser) {
    this.eachUser = eachUser;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    PolicyUpdate policyUpdate = (PolicyUpdate) o;
    return Objects.equals(this.name, policyUpdate.name) &&
        Objects.equals(this.description, policyUpdate.description) &&
        Objects.equals(this.config, policyUpdate.config) &&
        Objects.equals(this.enabled, policyUpdate.enabled) &&
        Objects.equals(this.refType, policyUpdate.refType) &&
        Objects.equals(this.refId, policyUpdate.refId) &&
        Objects.equals(this.accounts, policyUpdate.accounts) &&
        Objects.equals(this.eachUser, policyUpdate.eachUser);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, description, config, enabled, refType, refId, accounts, eachUser);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class PolicyUpdate {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    config: ").append(toIndentedString(config)).append("\n");
    sb.append("    enabled: ").append(toIndentedString(enabled)).append("\n");
    sb.append("    refType: ").append(toIndentedString(refType)).append("\n");
    sb.append("    refId: ").append(toIndentedString(refId)).append("\n");
    sb.append("    accounts: ").append(toIndentedString(accounts)).append("\n");
    sb.append("    eachUser: ").append(toIndentedString(eachUser)).append("\n");
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

