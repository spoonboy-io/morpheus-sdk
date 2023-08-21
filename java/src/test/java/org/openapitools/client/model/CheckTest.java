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
import java.math.BigDecimal;
import org.openapitools.client.model.AnyOfcheckWebConfigcheckSqlConfigcheckElasticsearchConfigcheckSocketConfigobjectcheckVmConfig;
import org.openapitools.client.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.client.model.CheckCheckType;
import org.openapitools.client.model.CheckContainer;
import org.openapitools.client.model.InlineResponse20083LoadBalancerNodeCreatedBy;
import org.threeten.bp.OffsetDateTime;
import org.junit.Assert;
import org.junit.Ignore;
import org.junit.Test;


/**
 * Model tests for Check
 */
public class CheckTest {
    private final Check model = new Check();

    /**
     * Model tests for Check
     */
    @Test
    public void testCheck() {
        // TODO: test Check
    }

    /**
     * Test the property 'id'
     */
    @Test
    public void idTest() {
        // TODO: test id
    }

    /**
     * Test the property 'account'
     */
    @Test
    public void accountTest() {
        // TODO: test account
    }

    /**
     * Test the property 'active'
     */
    @Test
    public void activeTest() {
        // TODO: test active
    }

    /**
     * Test the property 'apiKey'
     */
    @Test
    public void apiKeyTest() {
        // TODO: test apiKey
    }

    /**
     * Test the property 'availability'
     */
    @Test
    public void availabilityTest() {
        // TODO: test availability
    }

    /**
     * Test the property 'checkAgent'
     */
    @Test
    public void checkAgentTest() {
        // TODO: test checkAgent
    }

    /**
     * Test the property 'checkInterval'
     */
    @Test
    public void checkIntervalTest() {
        // TODO: test checkInterval
    }

    /**
     * Test the property 'checkSpec'
     */
    @Test
    public void checkSpecTest() {
        // TODO: test checkSpec
    }

    /**
     * Test the property 'checkType'
     */
    @Test
    public void checkTypeTest() {
        // TODO: test checkType
    }

    /**
     * Test the property 'config'
     */
    @Test
    public void configTest() {
        // TODO: test config
    }

    /**
     * Test the property 'container'
     */
    @Test
    public void containerTest() {
        // TODO: test container
    }

    /**
     * Test the property 'createIncident'
     */
    @Test
    public void createIncidentTest() {
        // TODO: test createIncident
    }

    /**
     * Test the property 'muted'
     */
    @Test
    public void mutedTest() {
        // TODO: test muted
    }

    /**
     * Test the property 'createdBy'
     */
    @Test
    public void createdByTest() {
        // TODO: test createdBy
    }

    /**
     * Test the property 'dateCreated'
     */
    @Test
    public void dateCreatedTest() {
        // TODO: test dateCreated
    }

    /**
     * Test the property 'description'
     */
    @Test
    public void descriptionTest() {
        // TODO: test description
    }

    /**
     * Test the property 'endDate'
     */
    @Test
    public void endDateTest() {
        // TODO: test endDate
    }

    /**
     * Test the property 'health'
     */
    @Test
    public void healthTest() {
        // TODO: test health
    }

    /**
     * Test the property 'inUptime'
     */
    @Test
    public void inUptimeTest() {
        // TODO: test inUptime
    }

    /**
     * Test the property 'lastBoxStats'
     */
    @Test
    public void lastBoxStatsTest() {
        // TODO: test lastBoxStats
    }

    /**
     * Test the property 'lastCheckStatus'
     */
    @Test
    public void lastCheckStatusTest() {
        // TODO: test lastCheckStatus
    }

    /**
     * Test the property 'lastError'
     */
    @Test
    public void lastErrorTest() {
        // TODO: test lastError
    }

    /**
     * Test the property 'lastErrorDate'
     */
    @Test
    public void lastErrorDateTest() {
        // TODO: test lastErrorDate
    }

    /**
     * Test the property 'lastMessage'
     */
    @Test
    public void lastMessageTest() {
        // TODO: test lastMessage
    }

    /**
     * Test the property 'lastMetric'
     */
    @Test
    public void lastMetricTest() {
        // TODO: test lastMetric
    }

    /**
     * Test the property 'lastRunDate'
     */
    @Test
    public void lastRunDateTest() {
        // TODO: test lastRunDate
    }

    /**
     * Test the property 'lastStats'
     */
    @Test
    public void lastStatsTest() {
        // TODO: test lastStats
    }

    /**
     * Test the property 'lastSuccessDate'
     */
    @Test
    public void lastSuccessDateTest() {
        // TODO: test lastSuccessDate
    }

    /**
     * Test the property 'lastTimer'
     */
    @Test
    public void lastTimerTest() {
        // TODO: test lastTimer
    }

    /**
     * Test the property 'lastUpdated'
     */
    @Test
    public void lastUpdatedTest() {
        // TODO: test lastUpdated
    }

    /**
     * Test the property 'lastWarningDate'
     */
    @Test
    public void lastWarningDateTest() {
        // TODO: test lastWarningDate
    }

    /**
     * Test the property 'name'
     */
    @Test
    public void nameTest() {
        // TODO: test name
    }

    /**
     * Test the property 'nextRunDate'
     */
    @Test
    public void nextRunDateTest() {
        // TODO: test nextRunDate
    }

    /**
     * Test the property 'outageTime'
     */
    @Test
    public void outageTimeTest() {
        // TODO: test outageTime
    }

    /**
     * Test the property 'severity'
     */
    @Test
    public void severityTest() {
        // TODO: test severity
    }

    /**
     * Test the property 'startDate'
     */
    @Test
    public void startDateTest() {
        // TODO: test startDate
    }

    /**
     * Test the property 'deleted'
     */
    @Test
    public void deletedTest() {
        // TODO: test deleted
    }

}
