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

/**
 * CheckSshConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class CheckSshConfig {
  public static final String SERIALIZED_NAME_SSH_PORT = "sshPort";
  @SerializedName(SERIALIZED_NAME_SSH_PORT)
  private Long sshPort;

  public static final String SERIALIZED_NAME_CHECK_USER = "checkUser";
  @SerializedName(SERIALIZED_NAME_CHECK_USER)
  private String checkUser;

  public static final String SERIALIZED_NAME_TUNNEL_ON = "tunnelOn";
  @SerializedName(SERIALIZED_NAME_TUNNEL_ON)
  private String tunnelOn;

  public static final String SERIALIZED_NAME_TEXT_CHECK_ON = "textCheckOn";
  @SerializedName(SERIALIZED_NAME_TEXT_CHECK_ON)
  private String textCheckOn;

  public static final String SERIALIZED_NAME_CHECK_PASSWORD = "checkPassword";
  @SerializedName(SERIALIZED_NAME_CHECK_PASSWORD)
  private String checkPassword;

  public static final String SERIALIZED_NAME_SSH_HOST = "sshHost";
  @SerializedName(SERIALIZED_NAME_SSH_HOST)
  private String sshHost;

  public static final String SERIALIZED_NAME_SSH_USER = "sshUser";
  @SerializedName(SERIALIZED_NAME_SSH_USER)
  private String sshUser;

  public static final String SERIALIZED_NAME_WEB_TEXT_MATCH = "webTextMatch";
  @SerializedName(SERIALIZED_NAME_WEB_TEXT_MATCH)
  private String webTextMatch;

  public static final String SERIALIZED_NAME_CHECK_PASSWORD_HASH = "checkPasswordHash";
  @SerializedName(SERIALIZED_NAME_CHECK_PASSWORD_HASH)
  private String checkPasswordHash;


  public CheckSshConfig sshPort(Long sshPort) {
    
    this.sshPort = sshPort;
    return this;
  }

   /**
   * Get sshPort
   * @return sshPort
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Long getSshPort() {
    return sshPort;
  }


  public void setSshPort(Long sshPort) {
    this.sshPort = sshPort;
  }


  public CheckSshConfig checkUser(String checkUser) {
    
    this.checkUser = checkUser;
    return this;
  }

   /**
   * Get checkUser
   * @return checkUser
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCheckUser() {
    return checkUser;
  }


  public void setCheckUser(String checkUser) {
    this.checkUser = checkUser;
  }


  public CheckSshConfig tunnelOn(String tunnelOn) {
    
    this.tunnelOn = tunnelOn;
    return this;
  }

   /**
   * Get tunnelOn
   * @return tunnelOn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTunnelOn() {
    return tunnelOn;
  }


  public void setTunnelOn(String tunnelOn) {
    this.tunnelOn = tunnelOn;
  }


  public CheckSshConfig textCheckOn(String textCheckOn) {
    
    this.textCheckOn = textCheckOn;
    return this;
  }

   /**
   * Get textCheckOn
   * @return textCheckOn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getTextCheckOn() {
    return textCheckOn;
  }


  public void setTextCheckOn(String textCheckOn) {
    this.textCheckOn = textCheckOn;
  }


  public CheckSshConfig checkPassword(String checkPassword) {
    
    this.checkPassword = checkPassword;
    return this;
  }

   /**
   * Get checkPassword
   * @return checkPassword
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCheckPassword() {
    return checkPassword;
  }


  public void setCheckPassword(String checkPassword) {
    this.checkPassword = checkPassword;
  }


  public CheckSshConfig sshHost(String sshHost) {
    
    this.sshHost = sshHost;
    return this;
  }

   /**
   * Get sshHost
   * @return sshHost
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSshHost() {
    return sshHost;
  }


  public void setSshHost(String sshHost) {
    this.sshHost = sshHost;
  }


  public CheckSshConfig sshUser(String sshUser) {
    
    this.sshUser = sshUser;
    return this;
  }

   /**
   * Get sshUser
   * @return sshUser
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getSshUser() {
    return sshUser;
  }


  public void setSshUser(String sshUser) {
    this.sshUser = sshUser;
  }


  public CheckSshConfig webTextMatch(String webTextMatch) {
    
    this.webTextMatch = webTextMatch;
    return this;
  }

   /**
   * Get webTextMatch
   * @return webTextMatch
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getWebTextMatch() {
    return webTextMatch;
  }


  public void setWebTextMatch(String webTextMatch) {
    this.webTextMatch = webTextMatch;
  }


  public CheckSshConfig checkPasswordHash(String checkPasswordHash) {
    
    this.checkPasswordHash = checkPasswordHash;
    return this;
  }

   /**
   * Get checkPasswordHash
   * @return checkPasswordHash
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getCheckPasswordHash() {
    return checkPasswordHash;
  }


  public void setCheckPasswordHash(String checkPasswordHash) {
    this.checkPasswordHash = checkPasswordHash;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CheckSshConfig checkSshConfig = (CheckSshConfig) o;
    return Objects.equals(this.sshPort, checkSshConfig.sshPort) &&
        Objects.equals(this.checkUser, checkSshConfig.checkUser) &&
        Objects.equals(this.tunnelOn, checkSshConfig.tunnelOn) &&
        Objects.equals(this.textCheckOn, checkSshConfig.textCheckOn) &&
        Objects.equals(this.checkPassword, checkSshConfig.checkPassword) &&
        Objects.equals(this.sshHost, checkSshConfig.sshHost) &&
        Objects.equals(this.sshUser, checkSshConfig.sshUser) &&
        Objects.equals(this.webTextMatch, checkSshConfig.webTextMatch) &&
        Objects.equals(this.checkPasswordHash, checkSshConfig.checkPasswordHash);
  }

  @Override
  public int hashCode() {
    return Objects.hash(sshPort, checkUser, tunnelOn, textCheckOn, checkPassword, sshHost, sshUser, webTextMatch, checkPasswordHash);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CheckSshConfig {\n");
    sb.append("    sshPort: ").append(toIndentedString(sshPort)).append("\n");
    sb.append("    checkUser: ").append(toIndentedString(checkUser)).append("\n");
    sb.append("    tunnelOn: ").append(toIndentedString(tunnelOn)).append("\n");
    sb.append("    textCheckOn: ").append(toIndentedString(textCheckOn)).append("\n");
    sb.append("    checkPassword: ").append(toIndentedString(checkPassword)).append("\n");
    sb.append("    sshHost: ").append(toIndentedString(sshHost)).append("\n");
    sb.append("    sshUser: ").append(toIndentedString(sshUser)).append("\n");
    sb.append("    webTextMatch: ").append(toIndentedString(webTextMatch)).append("\n");
    sb.append("    checkPasswordHash: ").append(toIndentedString(checkPasswordHash)).append("\n");
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

