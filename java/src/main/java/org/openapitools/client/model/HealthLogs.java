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
import org.threeten.bp.OffsetDateTime;

/**
 * HealthLogs
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class HealthLogs {
  public static final String SERIALIZED_NAME_TYPE_CODE = "typeCode";
  @SerializedName(SERIALIZED_NAME_TYPE_CODE)
  private String typeCode;

  public static final String SERIALIZED_NAME_TS = "ts";
  @SerializedName(SERIALIZED_NAME_TS)
  private OffsetDateTime ts;

  public static final String SERIALIZED_NAME_LEVEL = "level";
  @SerializedName(SERIALIZED_NAME_LEVEL)
  private String level;

  public static final String SERIALIZED_NAME_SOURCE_TYPE = "sourceType";
  @SerializedName(SERIALIZED_NAME_SOURCE_TYPE)
  private String sourceType;

  public static final String SERIALIZED_NAME_MESSAGE = "message";
  @SerializedName(SERIALIZED_NAME_MESSAGE)
  private String message;

  public static final String SERIALIZED_NAME_HOSTNAME = "hostname";
  @SerializedName(SERIALIZED_NAME_HOSTNAME)
  private String hostname;

  public static final String SERIALIZED_NAME_TITLE = "title";
  @SerializedName(SERIALIZED_NAME_TITLE)
  private String title;

  public static final String SERIALIZED_NAME_LOG_SIGNATURE = "logSignature";
  @SerializedName(SERIALIZED_NAME_LOG_SIGNATURE)
  private String logSignature;

  public static final String SERIALIZED_NAME_OBJECT_ID = "objectId";
  @SerializedName(SERIALIZED_NAME_OBJECT_ID)
  private String objectId;

  public static final String SERIALIZED_NAME_SEQ = "seq";
  @SerializedName(SERIALIZED_NAME_SEQ)
  private Long seq;

  public static final String SERIALIZED_NAME_ID = "_id";
  @SerializedName(SERIALIZED_NAME_ID)
  private String id;

  public static final String SERIALIZED_NAME_SIGNATURE_VERIFIED = "signatureVerified";
  @SerializedName(SERIALIZED_NAME_SIGNATURE_VERIFIED)
  private Boolean signatureVerified;


  public HealthLogs typeCode(String typeCode) {
    
    this.typeCode = typeCode;
    return this;
  }

   /**
   * Get typeCode
   * @return typeCode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTypeCode() {
    return typeCode;
  }


  public void setTypeCode(String typeCode) {
    this.typeCode = typeCode;
  }


  public HealthLogs ts(OffsetDateTime ts) {
    
    this.ts = ts;
    return this;
  }

   /**
   * Get ts
   * @return ts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getTs() {
    return ts;
  }


  public void setTs(OffsetDateTime ts) {
    this.ts = ts;
  }


  public HealthLogs level(String level) {
    
    this.level = level;
    return this;
  }

   /**
   * Get level
   * @return level
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLevel() {
    return level;
  }


  public void setLevel(String level) {
    this.level = level;
  }


  public HealthLogs sourceType(String sourceType) {
    
    this.sourceType = sourceType;
    return this;
  }

   /**
   * Get sourceType
   * @return sourceType
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSourceType() {
    return sourceType;
  }


  public void setSourceType(String sourceType) {
    this.sourceType = sourceType;
  }


  public HealthLogs message(String message) {
    
    this.message = message;
    return this;
  }

   /**
   * Get message
   * @return message
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMessage() {
    return message;
  }


  public void setMessage(String message) {
    this.message = message;
  }


  public HealthLogs hostname(String hostname) {
    
    this.hostname = hostname;
    return this;
  }

   /**
   * Get hostname
   * @return hostname
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getHostname() {
    return hostname;
  }


  public void setHostname(String hostname) {
    this.hostname = hostname;
  }


  public HealthLogs title(String title) {
    
    this.title = title;
    return this;
  }

   /**
   * Get title
   * @return title
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTitle() {
    return title;
  }


  public void setTitle(String title) {
    this.title = title;
  }


  public HealthLogs logSignature(String logSignature) {
    
    this.logSignature = logSignature;
    return this;
  }

   /**
   * Get logSignature
   * @return logSignature
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLogSignature() {
    return logSignature;
  }


  public void setLogSignature(String logSignature) {
    this.logSignature = logSignature;
  }


  public HealthLogs objectId(String objectId) {
    
    this.objectId = objectId;
    return this;
  }

   /**
   * Get objectId
   * @return objectId
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getObjectId() {
    return objectId;
  }


  public void setObjectId(String objectId) {
    this.objectId = objectId;
  }


  public HealthLogs seq(Long seq) {
    
    this.seq = seq;
    return this;
  }

   /**
   * Get seq
   * @return seq
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getSeq() {
    return seq;
  }


  public void setSeq(Long seq) {
    this.seq = seq;
  }


  public HealthLogs id(String id) {
    
    this.id = id;
    return this;
  }

   /**
   * Get id
   * @return id
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getId() {
    return id;
  }


  public void setId(String id) {
    this.id = id;
  }


  public HealthLogs signatureVerified(Boolean signatureVerified) {
    
    this.signatureVerified = signatureVerified;
    return this;
  }

   /**
   * Get signatureVerified
   * @return signatureVerified
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getSignatureVerified() {
    return signatureVerified;
  }


  public void setSignatureVerified(Boolean signatureVerified) {
    this.signatureVerified = signatureVerified;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    HealthLogs healthLogs = (HealthLogs) o;
    return Objects.equals(this.typeCode, healthLogs.typeCode) &&
        Objects.equals(this.ts, healthLogs.ts) &&
        Objects.equals(this.level, healthLogs.level) &&
        Objects.equals(this.sourceType, healthLogs.sourceType) &&
        Objects.equals(this.message, healthLogs.message) &&
        Objects.equals(this.hostname, healthLogs.hostname) &&
        Objects.equals(this.title, healthLogs.title) &&
        Objects.equals(this.logSignature, healthLogs.logSignature) &&
        Objects.equals(this.objectId, healthLogs.objectId) &&
        Objects.equals(this.seq, healthLogs.seq) &&
        Objects.equals(this.id, healthLogs.id) &&
        Objects.equals(this.signatureVerified, healthLogs.signatureVerified);
  }

  @Override
  public int hashCode() {
    return Objects.hash(typeCode, ts, level, sourceType, message, hostname, title, logSignature, objectId, seq, id, signatureVerified);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class HealthLogs {\n");
    sb.append("    typeCode: ").append(toIndentedString(typeCode)).append("\n");
    sb.append("    ts: ").append(toIndentedString(ts)).append("\n");
    sb.append("    level: ").append(toIndentedString(level)).append("\n");
    sb.append("    sourceType: ").append(toIndentedString(sourceType)).append("\n");
    sb.append("    message: ").append(toIndentedString(message)).append("\n");
    sb.append("    hostname: ").append(toIndentedString(hostname)).append("\n");
    sb.append("    title: ").append(toIndentedString(title)).append("\n");
    sb.append("    logSignature: ").append(toIndentedString(logSignature)).append("\n");
    sb.append("    objectId: ").append(toIndentedString(objectId)).append("\n");
    sb.append("    seq: ").append(toIndentedString(seq)).append("\n");
    sb.append("    id: ").append(toIndentedString(id)).append("\n");
    sb.append("    signatureVerified: ").append(toIndentedString(signatureVerified)).append("\n");
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

