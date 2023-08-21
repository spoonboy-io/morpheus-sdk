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
import org.openapitools.client.model.InlineResponse20040AppDeployInstance;
import org.threeten.bp.OffsetDateTime;

/**
 * Approvals
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class Approvals {
  public static final String SERIALIZED_NAME_ID = "id";
  @SerializedName(SERIALIZED_NAME_ID)
  private Long id;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_INTERNAL_ID = "internalId";
  @SerializedName(SERIALIZED_NAME_INTERNAL_ID)
  private String internalId;

  public static final String SERIALIZED_NAME_EXTERNAL_ID = "externalId";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_ID)
  private String externalId;

  public static final String SERIALIZED_NAME_EXTERNAL_NAME = "externalName";
  @SerializedName(SERIALIZED_NAME_EXTERNAL_NAME)
  private String externalName;

  public static final String SERIALIZED_NAME_REQUEST_TYPE = "requestType";
  @SerializedName(SERIALIZED_NAME_REQUEST_TYPE)
  private String requestType;

  public static final String SERIALIZED_NAME_ACCOUNT = "account";
  @SerializedName(SERIALIZED_NAME_ACCOUNT)
  private InlineResponse20040AppDeployInstance account;

  public static final String SERIALIZED_NAME_APPROVER = "approver";
  @SerializedName(SERIALIZED_NAME_APPROVER)
  private InlineResponse20040AppDeployInstance approver;

  public static final String SERIALIZED_NAME_ACCOUNT_INTEGRATION = "accountIntegration";
  @SerializedName(SERIALIZED_NAME_ACCOUNT_INTEGRATION)
  private String accountIntegration;

  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private String status;

  public static final String SERIALIZED_NAME_ERROR_MESSAGE = "errorMessage";
  @SerializedName(SERIALIZED_NAME_ERROR_MESSAGE)
  private String errorMessage;

  public static final String SERIALIZED_NAME_DATE_CREATED = "dateCreated";
  @SerializedName(SERIALIZED_NAME_DATE_CREATED)
  private OffsetDateTime dateCreated;

  public static final String SERIALIZED_NAME_LAST_UPDATED = "lastUpdated";
  @SerializedName(SERIALIZED_NAME_LAST_UPDATED)
  private OffsetDateTime lastUpdated;

  public static final String SERIALIZED_NAME_REQUEST_BY = "requestBy";
  @SerializedName(SERIALIZED_NAME_REQUEST_BY)
  private String requestBy;


  public Approvals id(Long id) {
    
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


  public Approvals name(String name) {
    
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


  public Approvals internalId(String internalId) {
    
    this.internalId = internalId;
    return this;
  }

   /**
   * Get internalId
   * @return internalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getInternalId() {
    return internalId;
  }


  public void setInternalId(String internalId) {
    this.internalId = internalId;
  }


  public Approvals externalId(String externalId) {
    
    this.externalId = externalId;
    return this;
  }

   /**
   * Get externalId
   * @return externalId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getExternalId() {
    return externalId;
  }


  public void setExternalId(String externalId) {
    this.externalId = externalId;
  }


  public Approvals externalName(String externalName) {
    
    this.externalName = externalName;
    return this;
  }

   /**
   * Get externalName
   * @return externalName
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getExternalName() {
    return externalName;
  }


  public void setExternalName(String externalName) {
    this.externalName = externalName;
  }


  public Approvals requestType(String requestType) {
    
    this.requestType = requestType;
    return this;
  }

   /**
   * Get requestType
   * @return requestType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRequestType() {
    return requestType;
  }


  public void setRequestType(String requestType) {
    this.requestType = requestType;
  }


  public Approvals account(InlineResponse20040AppDeployInstance account) {
    
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


  public Approvals approver(InlineResponse20040AppDeployInstance approver) {
    
    this.approver = approver;
    return this;
  }

   /**
   * Get approver
   * @return approver
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public InlineResponse20040AppDeployInstance getApprover() {
    return approver;
  }


  public void setApprover(InlineResponse20040AppDeployInstance approver) {
    this.approver = approver;
  }


  public Approvals accountIntegration(String accountIntegration) {
    
    this.accountIntegration = accountIntegration;
    return this;
  }

   /**
   * Get accountIntegration
   * @return accountIntegration
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getAccountIntegration() {
    return accountIntegration;
  }


  public void setAccountIntegration(String accountIntegration) {
    this.accountIntegration = accountIntegration;
  }


  public Approvals status(String status) {
    
    this.status = status;
    return this;
  }

   /**
   * Get status
   * @return status
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getStatus() {
    return status;
  }


  public void setStatus(String status) {
    this.status = status;
  }


  public Approvals errorMessage(String errorMessage) {
    
    this.errorMessage = errorMessage;
    return this;
  }

   /**
   * Get errorMessage
   * @return errorMessage
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getErrorMessage() {
    return errorMessage;
  }


  public void setErrorMessage(String errorMessage) {
    this.errorMessage = errorMessage;
  }


  public Approvals dateCreated(OffsetDateTime dateCreated) {
    
    this.dateCreated = dateCreated;
    return this;
  }

   /**
   * Get dateCreated
   * @return dateCreated
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getDateCreated() {
    return dateCreated;
  }


  public void setDateCreated(OffsetDateTime dateCreated) {
    this.dateCreated = dateCreated;
  }


  public Approvals lastUpdated(OffsetDateTime lastUpdated) {
    
    this.lastUpdated = lastUpdated;
    return this;
  }

   /**
   * Get lastUpdated
   * @return lastUpdated
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getLastUpdated() {
    return lastUpdated;
  }


  public void setLastUpdated(OffsetDateTime lastUpdated) {
    this.lastUpdated = lastUpdated;
  }


  public Approvals requestBy(String requestBy) {
    
    this.requestBy = requestBy;
    return this;
  }

   /**
   * Get requestBy
   * @return requestBy
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getRequestBy() {
    return requestBy;
  }


  public void setRequestBy(String requestBy) {
    this.requestBy = requestBy;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    Approvals approvals = (Approvals) o;
    return Objects.equals(this.id, approvals.id) &&
        Objects.equals(this.name, approvals.name) &&
        Objects.equals(this.internalId, approvals.internalId) &&
        Objects.equals(this.externalId, approvals.externalId) &&
        Objects.equals(this.externalName, approvals.externalName) &&
        Objects.equals(this.requestType, approvals.requestType) &&
        Objects.equals(this.account, approvals.account) &&
        Objects.equals(this.approver, approvals.approver) &&
        Objects.equals(this.accountIntegration, approvals.accountIntegration) &&
        Objects.equals(this.status, approvals.status) &&
        Objects.equals(this.errorMessage, approvals.errorMessage) &&
        Objects.equals(this.dateCreated, approvals.dateCreated) &&
        Objects.equals(this.lastUpdated, approvals.lastUpdated) &&
        Objects.equals(this.requestBy, approvals.requestBy);
  }

  @Override
  public int hashCode() {
    return Objects.hash(id, name, internalId, externalId, externalName, requestType, account, approver, accountIntegration, status, errorMessage, dateCreated, lastUpdated, requestBy);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class Approvals {\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    internalId: ").append(toIndentedString(internalId)).append("\n");
    sb.append("    externalId: ").append(toIndentedString(externalId)).append("\n");
    sb.append("    externalName: ").append(toIndentedString(externalName)).append("\n");
    sb.append("    requestType: ").append(toIndentedString(requestType)).append("\n");
    sb.append("    account: ").append(toIndentedString(account)).append("\n");
    sb.append("    approver: ").append(toIndentedString(approver)).append("\n");
    sb.append("    accountIntegration: ").append(toIndentedString(accountIntegration)).append("\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    errorMessage: ").append(toIndentedString(errorMessage)).append("\n");
    sb.append("    dateCreated: ").append(toIndentedString(dateCreated)).append("\n");
    sb.append("    lastUpdated: ").append(toIndentedString(lastUpdated)).append("\n");
    sb.append("    requestBy: ").append(toIndentedString(requestBy)).append("\n");
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

