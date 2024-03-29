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
 * CheckElasticsearchConfig
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class CheckElasticsearchConfig {
  public static final String SERIALIZED_NAME_ES_HOST = "esHost";
  @SerializedName(SERIALIZED_NAME_ES_HOST)
  private String esHost;

  public static final String SERIALIZED_NAME_ES_PORT = "esPort";
  @SerializedName(SERIALIZED_NAME_ES_PORT)
  private String esPort;

  public static final String SERIALIZED_NAME_CHECK_USER = "checkUser";
  @SerializedName(SERIALIZED_NAME_CHECK_USER)
  private String checkUser;

  public static final String SERIALIZED_NAME_TEXT_CHECK_ON = "textCheckOn";
  @SerializedName(SERIALIZED_NAME_TEXT_CHECK_ON)
  private String textCheckOn;

  public static final String SERIALIZED_NAME_CHECK_PASSWORD = "checkPassword";
  @SerializedName(SERIALIZED_NAME_CHECK_PASSWORD)
  private String checkPassword;

  public static final String SERIALIZED_NAME_WEB_TEXT_MATCH = "webTextMatch";
  @SerializedName(SERIALIZED_NAME_WEB_TEXT_MATCH)
  private String webTextMatch;

  public static final String SERIALIZED_NAME_CHECK_PASSWORD_HASH = "checkPasswordHash";
  @SerializedName(SERIALIZED_NAME_CHECK_PASSWORD_HASH)
  private String checkPasswordHash;

  /**
   * Set to on to turn on tunneling
   */
  @JsonAdapter(TunnelOnEnum.Adapter.class)
  public enum TunnelOnEnum {
    ON("on"),
    
    OFF("off");

    private String value;

    TunnelOnEnum(String value) {
      this.value = value;
    }

    public String getValue() {
      return value;
    }

    @Override
    public String toString() {
      return String.valueOf(value);
    }

    public static TunnelOnEnum fromValue(String value) {
      for (TunnelOnEnum b : TunnelOnEnum.values()) {
        if (b.value.equals(value)) {
          return b;
        }
      }
      throw new IllegalArgumentException("Unexpected value '" + value + "'");
    }

    public static class Adapter extends TypeAdapter<TunnelOnEnum> {
      @Override
      public void write(final JsonWriter jsonWriter, final TunnelOnEnum enumeration) throws IOException {
        jsonWriter.value(enumeration.getValue());
      }

      @Override
      public TunnelOnEnum read(final JsonReader jsonReader) throws IOException {
        String value =  jsonReader.nextString();
        return TunnelOnEnum.fromValue(value);
      }
    }
  }

  public static final String SERIALIZED_NAME_TUNNEL_ON = "tunnelOn";
  @SerializedName(SERIALIZED_NAME_TUNNEL_ON)
  private TunnelOnEnum tunnelOn = TunnelOnEnum.OFF;

  public static final String SERIALIZED_NAME_SSH_HOST = "sshHost";
  @SerializedName(SERIALIZED_NAME_SSH_HOST)
  private String sshHost;

  public static final String SERIALIZED_NAME_SSH_PORT = "sshPort";
  @SerializedName(SERIALIZED_NAME_SSH_PORT)
  private Long sshPort;

  public static final String SERIALIZED_NAME_SSH_USER = "sshUser";
  @SerializedName(SERIALIZED_NAME_SSH_USER)
  private String sshUser;

  public static final String SERIALIZED_NAME_SSH_PASSWORD = "sshPassword";
  @SerializedName(SERIALIZED_NAME_SSH_PASSWORD)
  private String sshPassword;


  public CheckElasticsearchConfig esHost(String esHost) {
    
    this.esHost = esHost;
    return this;
  }

   /**
   * Hostname or IP address of the Elasticsearch server
   * @return esHost
  **/
  @ApiModelProperty(example = "test.example.org", required = true, value = "Hostname or IP address of the Elasticsearch server")

  public String getEsHost() {
    return esHost;
  }


  public void setEsHost(String esHost) {
    this.esHost = esHost;
  }


  public CheckElasticsearchConfig esPort(String esPort) {
    
    this.esPort = esPort;
    return this;
  }

   /**
   * Port to connect to the HTTP service
   * @return esPort
  **/
  @ApiModelProperty(example = "9200", required = true, value = "Port to connect to the HTTP service")

  public String getEsPort() {
    return esPort;
  }


  public void setEsPort(String esPort) {
    this.esPort = esPort;
  }


  public CheckElasticsearchConfig checkUser(String checkUser) {
    
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


  public CheckElasticsearchConfig textCheckOn(String textCheckOn) {
    
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


  public CheckElasticsearchConfig checkPassword(String checkPassword) {
    
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


  public CheckElasticsearchConfig webTextMatch(String webTextMatch) {
    
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


  public CheckElasticsearchConfig checkPasswordHash(String checkPasswordHash) {
    
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


  public CheckElasticsearchConfig tunnelOn(TunnelOnEnum tunnelOn) {
    
    this.tunnelOn = tunnelOn;
    return this;
  }

   /**
   * Set to on to turn on tunneling
   * @return tunnelOn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Set to on to turn on tunneling")

  public TunnelOnEnum getTunnelOn() {
    return tunnelOn;
  }


  public void setTunnelOn(TunnelOnEnum tunnelOn) {
    this.tunnelOn = tunnelOn;
  }


  public CheckElasticsearchConfig sshHost(String sshHost) {
    
    this.sshHost = sshHost;
    return this;
  }

   /**
   * Hostname or IP address of the proxy host
   * @return sshHost
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Hostname or IP address of the proxy host")

  public String getSshHost() {
    return sshHost;
  }


  public void setSshHost(String sshHost) {
    this.sshHost = sshHost;
  }


  public CheckElasticsearchConfig sshPort(Long sshPort) {
    
    this.sshPort = sshPort;
    return this;
  }

   /**
   * Port for SSH on the proxy host, defaults to 22
   * @return sshPort
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Port for SSH on the proxy host, defaults to 22")

  public Long getSshPort() {
    return sshPort;
  }


  public void setSshPort(Long sshPort) {
    this.sshPort = sshPort;
  }


  public CheckElasticsearchConfig sshUser(String sshUser) {
    
    this.sshUser = sshUser;
    return this;
  }

   /**
   * SSH user on the proxy host to login as
   * @return sshUser
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "SSH user on the proxy host to login as")

  public String getSshUser() {
    return sshUser;
  }


  public void setSshUser(String sshUser) {
    this.sshUser = sshUser;
  }


  public CheckElasticsearchConfig sshPassword(String sshPassword) {
    
    this.sshPassword = sshPassword;
    return this;
  }

   /**
   * Password for user, if not using key based authentication
   * @return sshPassword
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Password for user, if not using key based authentication")

  public String getSshPassword() {
    return sshPassword;
  }


  public void setSshPassword(String sshPassword) {
    this.sshPassword = sshPassword;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CheckElasticsearchConfig checkElasticsearchConfig = (CheckElasticsearchConfig) o;
    return Objects.equals(this.esHost, checkElasticsearchConfig.esHost) &&
        Objects.equals(this.esPort, checkElasticsearchConfig.esPort) &&
        Objects.equals(this.checkUser, checkElasticsearchConfig.checkUser) &&
        Objects.equals(this.textCheckOn, checkElasticsearchConfig.textCheckOn) &&
        Objects.equals(this.checkPassword, checkElasticsearchConfig.checkPassword) &&
        Objects.equals(this.webTextMatch, checkElasticsearchConfig.webTextMatch) &&
        Objects.equals(this.checkPasswordHash, checkElasticsearchConfig.checkPasswordHash) &&
        Objects.equals(this.tunnelOn, checkElasticsearchConfig.tunnelOn) &&
        Objects.equals(this.sshHost, checkElasticsearchConfig.sshHost) &&
        Objects.equals(this.sshPort, checkElasticsearchConfig.sshPort) &&
        Objects.equals(this.sshUser, checkElasticsearchConfig.sshUser) &&
        Objects.equals(this.sshPassword, checkElasticsearchConfig.sshPassword);
  }

  @Override
  public int hashCode() {
    return Objects.hash(esHost, esPort, checkUser, textCheckOn, checkPassword, webTextMatch, checkPasswordHash, tunnelOn, sshHost, sshPort, sshUser, sshPassword);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CheckElasticsearchConfig {\n");
    sb.append("    esHost: ").append(toIndentedString(esHost)).append("\n");
    sb.append("    esPort: ").append(toIndentedString(esPort)).append("\n");
    sb.append("    checkUser: ").append(toIndentedString(checkUser)).append("\n");
    sb.append("    textCheckOn: ").append(toIndentedString(textCheckOn)).append("\n");
    sb.append("    checkPassword: ").append(toIndentedString(checkPassword)).append("\n");
    sb.append("    webTextMatch: ").append(toIndentedString(webTextMatch)).append("\n");
    sb.append("    checkPasswordHash: ").append(toIndentedString(checkPasswordHash)).append("\n");
    sb.append("    tunnelOn: ").append(toIndentedString(tunnelOn)).append("\n");
    sb.append("    sshHost: ").append(toIndentedString(sshHost)).append("\n");
    sb.append("    sshPort: ").append(toIndentedString(sshPort)).append("\n");
    sb.append("    sshUser: ").append(toIndentedString(sshUser)).append("\n");
    sb.append("    sshPassword: ").append(toIndentedString(sshPassword)).append("\n");
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

