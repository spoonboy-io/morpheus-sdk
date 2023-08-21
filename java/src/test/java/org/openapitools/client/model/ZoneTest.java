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
import org.openapitools.client.model.AnyOfobjectobject;
import org.openapitools.client.model.AnyOfzoneVcenterConfigzoneAwsConfigzoneAzureConfigzoneGcpConfig;
import org.openapitools.client.model.InlineResponse20040AppDeployInstance;
import org.openapitools.client.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;
import org.openapitools.client.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.openapitools.client.model.ZoneGroups;
import org.openapitools.client.model.ZoneStats;
import org.threeten.bp.OffsetDateTime;
import org.junit.Assert;
import org.junit.Ignore;
import org.junit.Test;


/**
 * Model tests for Zone
 */
public class ZoneTest {
    private final Zone model = new Zone();

    /**
     * Model tests for Zone
     */
    @Test
    public void testZone() {
        // TODO: test Zone
    }

    /**
     * Test the property 'id'
     */
    @Test
    public void idTest() {
        // TODO: test id
    }

    /**
     * Test the property 'uuid'
     */
    @Test
    public void uuidTest() {
        // TODO: test uuid
    }

    /**
     * Test the property 'externalId'
     */
    @Test
    public void externalIdTest() {
        // TODO: test externalId
    }

    /**
     * Test the property 'name'
     */
    @Test
    public void nameTest() {
        // TODO: test name
    }

    /**
     * Test the property 'code'
     */
    @Test
    public void codeTest() {
        // TODO: test code
    }

    /**
     * Test the property 'location'
     */
    @Test
    public void locationTest() {
        // TODO: test location
    }

    /**
     * Test the property 'owner'
     */
    @Test
    public void ownerTest() {
        // TODO: test owner
    }

    /**
     * Test the property 'accountId'
     */
    @Test
    public void accountIdTest() {
        // TODO: test accountId
    }

    /**
     * Test the property 'account'
     */
    @Test
    public void accountTest() {
        // TODO: test account
    }

    /**
     * Test the property 'visibility'
     */
    @Test
    public void visibilityTest() {
        // TODO: test visibility
    }

    /**
     * Test the property 'enabled'
     */
    @Test
    public void enabledTest() {
        // TODO: test enabled
    }

    /**
     * Test the property 'status'
     */
    @Test
    public void statusTest() {
        // TODO: test status
    }

    /**
     * Test the property 'statusMessage'
     */
    @Test
    public void statusMessageTest() {
        // TODO: test statusMessage
    }

    /**
     * Test the property 'statusDate'
     */
    @Test
    public void statusDateTest() {
        // TODO: test statusDate
    }

    /**
     * Test the property 'costStatus'
     */
    @Test
    public void costStatusTest() {
        // TODO: test costStatus
    }

    /**
     * Test the property 'costStatusMessage'
     */
    @Test
    public void costStatusMessageTest() {
        // TODO: test costStatusMessage
    }

    /**
     * Test the property 'costStatusDate'
     */
    @Test
    public void costStatusDateTest() {
        // TODO: test costStatusDate
    }

    /**
     * Test the property 'costLastSyncDuration'
     */
    @Test
    public void costLastSyncDurationTest() {
        // TODO: test costLastSyncDuration
    }

    /**
     * Test the property 'costLastSync'
     */
    @Test
    public void costLastSyncTest() {
        // TODO: test costLastSync
    }

    /**
     * Test the property 'zoneType'
     */
    @Test
    public void zoneTypeTest() {
        // TODO: test zoneType
    }

    /**
     * Test the property 'zoneTypeId'
     */
    @Test
    public void zoneTypeIdTest() {
        // TODO: test zoneTypeId
    }

    /**
     * Test the property 'guidanceMode'
     */
    @Test
    public void guidanceModeTest() {
        // TODO: test guidanceMode
    }

    /**
     * Test the property 'storageMode'
     */
    @Test
    public void storageModeTest() {
        // TODO: test storageMode
    }

    /**
     * Test the property 'agentMode'
     */
    @Test
    public void agentModeTest() {
        // TODO: test agentMode
    }

    /**
     * Test the property 'userDataLinux'
     */
    @Test
    public void userDataLinuxTest() {
        // TODO: test userDataLinux
    }

    /**
     * Test the property 'userDataWindows'
     */
    @Test
    public void userDataWindowsTest() {
        // TODO: test userDataWindows
    }

    /**
     * Test the property 'consoleKeymap'
     */
    @Test
    public void consoleKeymapTest() {
        // TODO: test consoleKeymap
    }

    /**
     * Test the property 'containerMode'
     */
    @Test
    public void containerModeTest() {
        // TODO: test containerMode
    }

    /**
     * Test the property 'costingMode'
     */
    @Test
    public void costingModeTest() {
        // TODO: test costingMode
    }

    /**
     * Test the property 'serviceVersion'
     */
    @Test
    public void serviceVersionTest() {
        // TODO: test serviceVersion
    }

    /**
     * Test the property 'securityMode'
     */
    @Test
    public void securityModeTest() {
        // TODO: test securityMode
    }

    /**
     * Test the property 'inventoryLevel'
     */
    @Test
    public void inventoryLevelTest() {
        // TODO: test inventoryLevel
    }

    /**
     * Test the property 'timezone'
     */
    @Test
    public void timezoneTest() {
        // TODO: test timezone
    }

    /**
     * Test the property 'apiProxy'
     */
    @Test
    public void apiProxyTest() {
        // TODO: test apiProxy
    }

    /**
     * Test the property 'provisioningProxy'
     */
    @Test
    public void provisioningProxyTest() {
        // TODO: test provisioningProxy
    }

    /**
     * Test the property 'networkDomain'
     */
    @Test
    public void networkDomainTest() {
        // TODO: test networkDomain
    }

    /**
     * Test the property 'domainName'
     */
    @Test
    public void domainNameTest() {
        // TODO: test domainName
    }

    /**
     * Test the property 'regionCode'
     */
    @Test
    public void regionCodeTest() {
        // TODO: test regionCode
    }

    /**
     * Test the property 'autoRecoverPowerState'
     */
    @Test
    public void autoRecoverPowerStateTest() {
        // TODO: test autoRecoverPowerState
    }

    /**
     * Test the property 'scalePriority'
     */
    @Test
    public void scalePriorityTest() {
        // TODO: test scalePriority
    }

    /**
     * Test the property 'config'
     */
    @Test
    public void configTest() {
        // TODO: test config
    }

    /**
     * Test the property 'credential'
     */
    @Test
    public void credentialTest() {
        // TODO: test credential
    }

    /**
     * Test the property 'imagePath'
     */
    @Test
    public void imagePathTest() {
        // TODO: test imagePath
    }

    /**
     * Test the property 'darkImagePath'
     */
    @Test
    public void darkImagePathTest() {
        // TODO: test darkImagePath
    }

    /**
     * Test the property 'dateCreated'
     */
    @Test
    public void dateCreatedTest() {
        // TODO: test dateCreated
    }

    /**
     * Test the property 'lastUpdated'
     */
    @Test
    public void lastUpdatedTest() {
        // TODO: test lastUpdated
    }

    /**
     * Test the property 'lastSync'
     */
    @Test
    public void lastSyncTest() {
        // TODO: test lastSync
    }

    /**
     * Test the property 'lastSyncDuration'
     */
    @Test
    public void lastSyncDurationTest() {
        // TODO: test lastSyncDuration
    }

    /**
     * Test the property 'nextRunDate'
     */
    @Test
    public void nextRunDateTest() {
        // TODO: test nextRunDate
    }

    /**
     * Test the property 'groups'
     */
    @Test
    public void groupsTest() {
        // TODO: test groups
    }

    /**
     * Test the property 'securityServer'
     */
    @Test
    public void securityServerTest() {
        // TODO: test securityServer
    }

    /**
     * Test the property 'networkServer'
     */
    @Test
    public void networkServerTest() {
        // TODO: test networkServer
    }

    /**
     * Test the property 'stats'
     */
    @Test
    public void statsTest() {
        // TODO: test stats
    }

    /**
     * Test the property 'serverCount'
     */
    @Test
    public void serverCountTest() {
        // TODO: test serverCount
    }

}