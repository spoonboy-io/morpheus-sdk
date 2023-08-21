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
import java.util.ArrayList;
import java.util.List;
import org.openapitools.client.model.InlineResponse20040AppDeployInstance;
import org.openapitools.client.model.InvoiceCloud;
import org.openapitools.client.model.InvoiceLineItems;
import org.threeten.bp.OffsetDateTime;
import org.junit.Assert;
import org.junit.Ignore;
import org.junit.Test;


/**
 * Model tests for Invoice
 */
public class InvoiceTest {
    private final Invoice model = new Invoice();

    /**
     * Model tests for Invoice
     */
    @Test
    public void testInvoice() {
        // TODO: test Invoice
    }

    /**
     * Test the property 'id'
     */
    @Test
    public void idTest() {
        // TODO: test id
    }

    /**
     * Test the property 'ownerId'
     */
    @Test
    public void ownerIdTest() {
        // TODO: test ownerId
    }

    /**
     * Test the property 'account'
     */
    @Test
    public void accountTest() {
        // TODO: test account
    }

    /**
     * Test the property 'group'
     */
    @Test
    public void groupTest() {
        // TODO: test group
    }

    /**
     * Test the property 'cloud'
     */
    @Test
    public void cloudTest() {
        // TODO: test cloud
    }

    /**
     * Test the property 'instance'
     */
    @Test
    public void instanceTest() {
        // TODO: test instance
    }

    /**
     * Test the property 'server'
     */
    @Test
    public void serverTest() {
        // TODO: test server
    }

    /**
     * Test the property 'cluster'
     */
    @Test
    public void clusterTest() {
        // TODO: test cluster
    }

    /**
     * Test the property 'user'
     */
    @Test
    public void userTest() {
        // TODO: test user
    }

    /**
     * Test the property 'plan'
     */
    @Test
    public void planTest() {
        // TODO: test plan
    }

    /**
     * Test the property 'tags'
     */
    @Test
    public void tagsTest() {
        // TODO: test tags
    }

    /**
     * Test the property 'project'
     */
    @Test
    public void projectTest() {
        // TODO: test project
    }

    /**
     * Test the property 'refType'
     */
    @Test
    public void refTypeTest() {
        // TODO: test refType
    }

    /**
     * Test the property 'refId'
     */
    @Test
    public void refIdTest() {
        // TODO: test refId
    }

    /**
     * Test the property 'refUuid'
     */
    @Test
    public void refUuidTest() {
        // TODO: test refUuid
    }

    /**
     * Test the property 'refName'
     */
    @Test
    public void refNameTest() {
        // TODO: test refName
    }

    /**
     * Test the property 'refCategory'
     */
    @Test
    public void refCategoryTest() {
        // TODO: test refCategory
    }

    /**
     * Test the property 'resourceId'
     */
    @Test
    public void resourceIdTest() {
        // TODO: test resourceId
    }

    /**
     * Test the property 'resourceUuid'
     */
    @Test
    public void resourceUuidTest() {
        // TODO: test resourceUuid
    }

    /**
     * Test the property 'resourceType'
     */
    @Test
    public void resourceTypeTest() {
        // TODO: test resourceType
    }

    /**
     * Test the property 'resourceName'
     */
    @Test
    public void resourceNameTest() {
        // TODO: test resourceName
    }

    /**
     * Test the property 'resourceExternalId'
     */
    @Test
    public void resourceExternalIdTest() {
        // TODO: test resourceExternalId
    }

    /**
     * Test the property 'resourceInternalId'
     */
    @Test
    public void resourceInternalIdTest() {
        // TODO: test resourceInternalId
    }

    /**
     * Test the property 'interval'
     */
    @Test
    public void intervalTest() {
        // TODO: test interval
    }

    /**
     * Test the property 'period'
     */
    @Test
    public void periodTest() {
        // TODO: test period
    }

    /**
     * Test the property 'estimate'
     */
    @Test
    public void estimateTest() {
        // TODO: test estimate
    }

    /**
     * Test the property 'summaryInvoice'
     */
    @Test
    public void summaryInvoiceTest() {
        // TODO: test summaryInvoice
    }

    /**
     * Test the property 'active'
     */
    @Test
    public void activeTest() {
        // TODO: test active
    }

    /**
     * Test the property 'startDate'
     */
    @Test
    public void startDateTest() {
        // TODO: test startDate
    }

    /**
     * Test the property 'endDate'
     */
    @Test
    public void endDateTest() {
        // TODO: test endDate
    }

    /**
     * Test the property 'refStart'
     */
    @Test
    public void refStartTest() {
        // TODO: test refStart
    }

    /**
     * Test the property 'refEnd'
     */
    @Test
    public void refEndTest() {
        // TODO: test refEnd
    }

    /**
     * Test the property 'estimatedComputePrice'
     */
    @Test
    public void estimatedComputePriceTest() {
        // TODO: test estimatedComputePrice
    }

    /**
     * Test the property 'estimatedComputeCost'
     */
    @Test
    public void estimatedComputeCostTest() {
        // TODO: test estimatedComputeCost
    }

    /**
     * Test the property 'estimatedMemoryPrice'
     */
    @Test
    public void estimatedMemoryPriceTest() {
        // TODO: test estimatedMemoryPrice
    }

    /**
     * Test the property 'estimatedMemoryCost'
     */
    @Test
    public void estimatedMemoryCostTest() {
        // TODO: test estimatedMemoryCost
    }

    /**
     * Test the property 'estimatedStoragePrice'
     */
    @Test
    public void estimatedStoragePriceTest() {
        // TODO: test estimatedStoragePrice
    }

    /**
     * Test the property 'estimatedStorageCost'
     */
    @Test
    public void estimatedStorageCostTest() {
        // TODO: test estimatedStorageCost
    }

    /**
     * Test the property 'estimatedNetworkPrice'
     */
    @Test
    public void estimatedNetworkPriceTest() {
        // TODO: test estimatedNetworkPrice
    }

    /**
     * Test the property 'estimatedNetworkCost'
     */
    @Test
    public void estimatedNetworkCostTest() {
        // TODO: test estimatedNetworkCost
    }

    /**
     * Test the property 'estimatedLicensePrice'
     */
    @Test
    public void estimatedLicensePriceTest() {
        // TODO: test estimatedLicensePrice
    }

    /**
     * Test the property 'estimatedLicenseCost'
     */
    @Test
    public void estimatedLicenseCostTest() {
        // TODO: test estimatedLicenseCost
    }

    /**
     * Test the property 'estimatedExtraPrice'
     */
    @Test
    public void estimatedExtraPriceTest() {
        // TODO: test estimatedExtraPrice
    }

    /**
     * Test the property 'estimatedExtraCost'
     */
    @Test
    public void estimatedExtraCostTest() {
        // TODO: test estimatedExtraCost
    }

    /**
     * Test the property 'estimatedTotalPrice'
     */
    @Test
    public void estimatedTotalPriceTest() {
        // TODO: test estimatedTotalPrice
    }

    /**
     * Test the property 'estimatedTotalCost'
     */
    @Test
    public void estimatedTotalCostTest() {
        // TODO: test estimatedTotalCost
    }

    /**
     * Test the property 'estimatedRunningPrice'
     */
    @Test
    public void estimatedRunningPriceTest() {
        // TODO: test estimatedRunningPrice
    }

    /**
     * Test the property 'estimatedRunningCost'
     */
    @Test
    public void estimatedRunningCostTest() {
        // TODO: test estimatedRunningCost
    }

    /**
     * Test the property 'estimatedCurrency'
     */
    @Test
    public void estimatedCurrencyTest() {
        // TODO: test estimatedCurrency
    }

    /**
     * Test the property 'estimatedConversionRate'
     */
    @Test
    public void estimatedConversionRateTest() {
        // TODO: test estimatedConversionRate
    }

    /**
     * Test the property 'actualComputePrice'
     */
    @Test
    public void actualComputePriceTest() {
        // TODO: test actualComputePrice
    }

    /**
     * Test the property 'actualComputeCost'
     */
    @Test
    public void actualComputeCostTest() {
        // TODO: test actualComputeCost
    }

    /**
     * Test the property 'actualMemoryPrice'
     */
    @Test
    public void actualMemoryPriceTest() {
        // TODO: test actualMemoryPrice
    }

    /**
     * Test the property 'actualMemoryCost'
     */
    @Test
    public void actualMemoryCostTest() {
        // TODO: test actualMemoryCost
    }

    /**
     * Test the property 'actualStoragePrice'
     */
    @Test
    public void actualStoragePriceTest() {
        // TODO: test actualStoragePrice
    }

    /**
     * Test the property 'actualStorageCost'
     */
    @Test
    public void actualStorageCostTest() {
        // TODO: test actualStorageCost
    }

    /**
     * Test the property 'actualNetworkPrice'
     */
    @Test
    public void actualNetworkPriceTest() {
        // TODO: test actualNetworkPrice
    }

    /**
     * Test the property 'actualNetworkCost'
     */
    @Test
    public void actualNetworkCostTest() {
        // TODO: test actualNetworkCost
    }

    /**
     * Test the property 'actualLicensePrice'
     */
    @Test
    public void actualLicensePriceTest() {
        // TODO: test actualLicensePrice
    }

    /**
     * Test the property 'actualLicenseCost'
     */
    @Test
    public void actualLicenseCostTest() {
        // TODO: test actualLicenseCost
    }

    /**
     * Test the property 'actualExtraPrice'
     */
    @Test
    public void actualExtraPriceTest() {
        // TODO: test actualExtraPrice
    }

    /**
     * Test the property 'actualExtraCost'
     */
    @Test
    public void actualExtraCostTest() {
        // TODO: test actualExtraCost
    }

    /**
     * Test the property 'actualTotalPrice'
     */
    @Test
    public void actualTotalPriceTest() {
        // TODO: test actualTotalPrice
    }

    /**
     * Test the property 'actualTotalCost'
     */
    @Test
    public void actualTotalCostTest() {
        // TODO: test actualTotalCost
    }

    /**
     * Test the property 'actualRunningPrice'
     */
    @Test
    public void actualRunningPriceTest() {
        // TODO: test actualRunningPrice
    }

    /**
     * Test the property 'actualRunningCost'
     */
    @Test
    public void actualRunningCostTest() {
        // TODO: test actualRunningCost
    }

    /**
     * Test the property 'actualCurrency'
     */
    @Test
    public void actualCurrencyTest() {
        // TODO: test actualCurrency
    }

    /**
     * Test the property 'actualConversionRate'
     */
    @Test
    public void actualConversionRateTest() {
        // TODO: test actualConversionRate
    }

    /**
     * Test the property 'computePrice'
     */
    @Test
    public void computePriceTest() {
        // TODO: test computePrice
    }

    /**
     * Test the property 'computeCost'
     */
    @Test
    public void computeCostTest() {
        // TODO: test computeCost
    }

    /**
     * Test the property 'memoryPrice'
     */
    @Test
    public void memoryPriceTest() {
        // TODO: test memoryPrice
    }

    /**
     * Test the property 'memoryCost'
     */
    @Test
    public void memoryCostTest() {
        // TODO: test memoryCost
    }

    /**
     * Test the property 'storagePrice'
     */
    @Test
    public void storagePriceTest() {
        // TODO: test storagePrice
    }

    /**
     * Test the property 'storageCost'
     */
    @Test
    public void storageCostTest() {
        // TODO: test storageCost
    }

    /**
     * Test the property 'networkPrice'
     */
    @Test
    public void networkPriceTest() {
        // TODO: test networkPrice
    }

    /**
     * Test the property 'networkCost'
     */
    @Test
    public void networkCostTest() {
        // TODO: test networkCost
    }

    /**
     * Test the property 'licensePrice'
     */
    @Test
    public void licensePriceTest() {
        // TODO: test licensePrice
    }

    /**
     * Test the property 'licenseCost'
     */
    @Test
    public void licenseCostTest() {
        // TODO: test licenseCost
    }

    /**
     * Test the property 'extraPrice'
     */
    @Test
    public void extraPriceTest() {
        // TODO: test extraPrice
    }

    /**
     * Test the property 'extraCost'
     */
    @Test
    public void extraCostTest() {
        // TODO: test extraCost
    }

    /**
     * Test the property 'totalPrice'
     */
    @Test
    public void totalPriceTest() {
        // TODO: test totalPrice
    }

    /**
     * Test the property 'totalCost'
     */
    @Test
    public void totalCostTest() {
        // TODO: test totalCost
    }

    /**
     * Test the property 'runningPrice'
     */
    @Test
    public void runningPriceTest() {
        // TODO: test runningPrice
    }

    /**
     * Test the property 'runningCost'
     */
    @Test
    public void runningCostTest() {
        // TODO: test runningCost
    }

    /**
     * Test the property 'currency'
     */
    @Test
    public void currencyTest() {
        // TODO: test currency
    }

    /**
     * Test the property 'conversionRate'
     */
    @Test
    public void conversionRateTest() {
        // TODO: test conversionRate
    }

    /**
     * Test the property 'costType'
     */
    @Test
    public void costTypeTest() {
        // TODO: test costType
    }

    /**
     * Test the property 'offTime'
     */
    @Test
    public void offTimeTest() {
        // TODO: test offTime
    }

    /**
     * Test the property 'powerState'
     */
    @Test
    public void powerStateTest() {
        // TODO: test powerState
    }

    /**
     * Test the property 'powerDate'
     */
    @Test
    public void powerDateTest() {
        // TODO: test powerDate
    }

    /**
     * Test the property 'runningMultiplier'
     */
    @Test
    public void runningMultiplierTest() {
        // TODO: test runningMultiplier
    }

    /**
     * Test the property 'usageType'
     */
    @Test
    public void usageTypeTest() {
        // TODO: test usageType
    }

    /**
     * Test the property 'usageCategory'
     */
    @Test
    public void usageCategoryTest() {
        // TODO: test usageCategory
    }

    /**
     * Test the property 'lastCostDate'
     */
    @Test
    public void lastCostDateTest() {
        // TODO: test lastCostDate
    }

    /**
     * Test the property 'lastActualDate'
     */
    @Test
    public void lastActualDateTest() {
        // TODO: test lastActualDate
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
     * Test the property 'lineItemCount'
     */
    @Test
    public void lineItemCountTest() {
        // TODO: test lineItemCount
    }

    /**
     * Test the property 'lineItems'
     */
    @Test
    public void lineItemsTest() {
        // TODO: test lineItems
    }

}