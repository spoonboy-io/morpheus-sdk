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
import org.openapitools.client.model.OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject;

/**
 * PolicyGroupCreatePolicyType
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class PolicyGroupCreatePolicyType {
  /**
   * The policy type
   */
  @JsonAdapter(CodeEnum.Adapter.class)
  public enum CodeEnum {
    APPROVE_DELETE("Approve Delete"),
    
    APPROVE_PROVISION("Approve Provision"),
    
    APPROVE_RECONFIGURE("Approve Reconfigure"),
    
    BACKUP_CREATION("Backup Creation"),
    
    BACKUP_TARGETS("Backup Targets"),
    
    BUDGET("Budget"),
    
    CLUSTER_RESOURCE_NAME("Cluster Resource Name"),
    
    CYPHER_ACCESS("Cypher Access"),
    
    DELAYED_DELETE("Delayed Delete"),
    
    EXPIRATION("Expiration"),
    
    FILE_SHARE_STORAGE_QUOTA("File Share Storage Quota"),
    
    HOSTNAME("Hostname"),
    
    INSTANCE_NAME("Instance Name"),
    
    MAX_CONTAINERS("Max Containers"),
    
    MAX_CORES("Max Cores"),
    
    MAX_HOSTS("Max Hosts"),
    
    MAX_LOAD_BALANCER_POOLS("Max Load Balancer Pools"),
    
    MAX_MEMORY("Max Memory"),
    
    MAX_POOL_MEMBERS("Max Pool Members"),
    
    MAX_STORAGE("Max Storage"),
    
    MAX_VIRTUAL_SERVERS("Max Virtual Servers"),
    
    MAX_VMS("Max VMs"),
    
    MESSAGE_OF_THE_DAY("Message of the Day"),
    
    NETWORK_QUOTA("Network Quota"),
    
    OBJECT_STORAGE_QUOTA("Object Storage Quota"),
    
    POWER_SCHEDULE("Power Schedule"),
    
    ROUTER_QUOTA("Router Quota"),
    
    SHUTDOWN("Shutdown"),
    
    STORAGE_SERVER_STORAGE_QUOTA("Storage Server Storage Quota"),
    
    TAGS("Tags"),
    
    USER_CREATION("User Creation"),
    
    USER_GROUP_CREATION("User Group Creation"),
    
    WORKFLOW("Workflow");

    private String value;

    CodeEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static CodeEnum fromValue(String value) {
      for (CodeEnum b : CodeEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<CodeEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final CodeEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public CodeEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return CodeEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_CODE = "code";
  @SerializedName(SERIALIZED_NAME_CODE)
  private CodeEnum code;

  public static final String SERIALIZED_NAME_CONFIG = "config";
  @SerializedName(SERIALIZED_NAME_CONFIG)
  private OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject config = null;

  public static final String SERIALIZED_NAME_ENABLED = "enabled";
  @SerializedName(SERIALIZED_NAME_ENABLED)
  private Boolean enabled = true;

  /**
   * Scope object type
   */
  @JsonAdapter(RefTypeEnum.Adapter.class)
  public enum RefTypeEnum {
    COMPUTESITE("ComputeSite");

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


  public PolicyGroupCreatePolicyType code(CodeEnum code) {
    
    this.code = code;
    return this;
  }

   /**
   * The policy type
   * @return code
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The policy type")

  public CodeEnum getCode() {
    return code;
  }


  public void setCode(CodeEnum code) {
    this.code = code;
  }


  public PolicyGroupCreatePolicyType config(OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject config) {
    
    this.config = config;
    return this;
  }

   /**
   * A map of config values. The expected values vary by policyType.
   * @return config
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "A map of config values. The expected values vary by policyType.")

  public OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject getConfig() {
    return config;
  }


  public void setConfig(OneOfobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobjectobject config) {
    this.config = config;
  }


  public PolicyGroupCreatePolicyType enabled(Boolean enabled) {
    
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


  public PolicyGroupCreatePolicyType refType(RefTypeEnum refType) {
    
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


  public PolicyGroupCreatePolicyType refId(Long refId) {
    
    this.refId = refId;
    return this;
  }

   /**
   * Scope object ID (&#x60;group&#x60;)
   * @return refId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Scope object ID (`group`)")

  public Long getRefId() {
    return refId;
  }


  public void setRefId(Long refId) {
    this.refId = refId;
  }


  public PolicyGroupCreatePolicyType accounts(List<Long> accounts) {
    
    this.accounts = accounts;
    return this;
  }

  public PolicyGroupCreatePolicyType addAccountsItem(Long accountsItem) {
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


  public PolicyGroupCreatePolicyType eachUser(Boolean eachUser) {
    
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
    PolicyGroupCreatePolicyType policyGroupCreatePolicyType = (PolicyGroupCreatePolicyType) o;
    return Objects.equals(this.code, policyGroupCreatePolicyType.code) &&
        Objects.equals(this.config, policyGroupCreatePolicyType.config) &&
        Objects.equals(this.enabled, policyGroupCreatePolicyType.enabled) &&
        Objects.equals(this.refType, policyGroupCreatePolicyType.refType) &&
        Objects.equals(this.refId, policyGroupCreatePolicyType.refId) &&
        Objects.equals(this.accounts, policyGroupCreatePolicyType.accounts) &&
        Objects.equals(this.eachUser, policyGroupCreatePolicyType.eachUser);
  }

  @Override
  public int hashCode() {
    return Objects.hash(code, config, enabled, refType, refId, accounts, eachUser);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class PolicyGroupCreatePolicyType {\n");
    sb.append("    code: ").append(toIndentedString(code)).append("\n");
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

