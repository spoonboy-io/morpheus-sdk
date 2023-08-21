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
import org.openapitools.client.model.Snapshot;

/**
 * InstanceSnapshots
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-08-21T11:26:47.028Z[GMT]")
public class InstanceSnapshots {
  public static final String SERIALIZED_NAME_SNAPSHOTS = "snapshots";
  @SerializedName(SERIALIZED_NAME_SNAPSHOTS)
  private List<Snapshot> snapshots = null;


  public InstanceSnapshots snapshots(List<Snapshot> snapshots) {
    
    this.snapshots = snapshots;
    return this;
  }

  public InstanceSnapshots addSnapshotsItem(Snapshot snapshotsItem) {
    if (this.snapshots == null) {
      this.snapshots = new ArrayList<Snapshot>();
    }
    this.snapshots.add(snapshotsItem);
    return this;
  }

   /**
   * List of snapshot objects
   * @return snapshots
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "List of snapshot objects")

  public List<Snapshot> getSnapshots() {
    return snapshots;
  }


  public void setSnapshots(List<Snapshot> snapshots) {
    this.snapshots = snapshots;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    InstanceSnapshots instanceSnapshots = (InstanceSnapshots) o;
    return Objects.equals(this.snapshots, instanceSnapshots.snapshots);
  }

  @Override
  public int hashCode() {
    return Objects.hash(snapshots);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class InstanceSnapshots {\n");
    sb.append("    snapshots: ").append(toIndentedString(snapshots)).append("\n");
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

