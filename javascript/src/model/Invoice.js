/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import InlineResponse20040AppDeployInstance from './InlineResponse20040AppDeployInstance';
import InvoiceCloud from './InvoiceCloud';
import InvoiceLineItems from './InvoiceLineItems';

/**
 * The Invoice model module.
 * @module model/Invoice
 * @version 6.2.1
 */
class Invoice {
    /**
     * Constructs a new <code>Invoice</code>.
     * @alias module:model/Invoice
     */
    constructor() { 
        
        Invoice.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Invoice</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Invoice} obj Optional instance to populate.
     * @return {module:model/Invoice} The populated <code>Invoice</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Invoice();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('ownerId')) {
                obj['ownerId'] = ApiClient.convertToType(data['ownerId'], 'Number');
            }
            if (data.hasOwnProperty('account')) {
                obj['account'] = InlineResponse20040AppDeployInstance.constructFromObject(data['account']);
            }
            if (data.hasOwnProperty('group')) {
                obj['group'] = ApiClient.convertToType(data['group'], Object);
            }
            if (data.hasOwnProperty('cloud')) {
                obj['cloud'] = InvoiceCloud.constructFromObject(data['cloud']);
            }
            if (data.hasOwnProperty('instance')) {
                obj['instance'] = ApiClient.convertToType(data['instance'], Object);
            }
            if (data.hasOwnProperty('server')) {
                obj['server'] = ApiClient.convertToType(data['server'], 'String');
            }
            if (data.hasOwnProperty('cluster')) {
                obj['cluster'] = ApiClient.convertToType(data['cluster'], 'String');
            }
            if (data.hasOwnProperty('user')) {
                obj['user'] = ApiClient.convertToType(data['user'], Object);
            }
            if (data.hasOwnProperty('plan')) {
                obj['plan'] = ApiClient.convertToType(data['plan'], Object);
            }
            if (data.hasOwnProperty('tags')) {
                obj['tags'] = ApiClient.convertToType(data['tags'], [Object]);
            }
            if (data.hasOwnProperty('project')) {
                obj['project'] = ApiClient.convertToType(data['project'], 'String');
            }
            if (data.hasOwnProperty('refType')) {
                obj['refType'] = ApiClient.convertToType(data['refType'], 'String');
            }
            if (data.hasOwnProperty('refId')) {
                obj['refId'] = ApiClient.convertToType(data['refId'], 'Number');
            }
            if (data.hasOwnProperty('refUuid')) {
                obj['refUuid'] = ApiClient.convertToType(data['refUuid'], 'String');
            }
            if (data.hasOwnProperty('refName')) {
                obj['refName'] = ApiClient.convertToType(data['refName'], 'String');
            }
            if (data.hasOwnProperty('refCategory')) {
                obj['refCategory'] = ApiClient.convertToType(data['refCategory'], 'String');
            }
            if (data.hasOwnProperty('resourceId')) {
                obj['resourceId'] = ApiClient.convertToType(data['resourceId'], 'String');
            }
            if (data.hasOwnProperty('resourceUuid')) {
                obj['resourceUuid'] = ApiClient.convertToType(data['resourceUuid'], 'String');
            }
            if (data.hasOwnProperty('resourceType')) {
                obj['resourceType'] = ApiClient.convertToType(data['resourceType'], 'String');
            }
            if (data.hasOwnProperty('resourceName')) {
                obj['resourceName'] = ApiClient.convertToType(data['resourceName'], 'String');
            }
            if (data.hasOwnProperty('resourceExternalId')) {
                obj['resourceExternalId'] = ApiClient.convertToType(data['resourceExternalId'], 'String');
            }
            if (data.hasOwnProperty('resourceInternalId')) {
                obj['resourceInternalId'] = ApiClient.convertToType(data['resourceInternalId'], 'String');
            }
            if (data.hasOwnProperty('interval')) {
                obj['interval'] = ApiClient.convertToType(data['interval'], 'String');
            }
            if (data.hasOwnProperty('period')) {
                obj['period'] = ApiClient.convertToType(data['period'], 'String');
            }
            if (data.hasOwnProperty('estimate')) {
                obj['estimate'] = ApiClient.convertToType(data['estimate'], 'Boolean');
            }
            if (data.hasOwnProperty('summaryInvoice')) {
                obj['summaryInvoice'] = ApiClient.convertToType(data['summaryInvoice'], 'Boolean');
            }
            if (data.hasOwnProperty('active')) {
                obj['active'] = ApiClient.convertToType(data['active'], 'Boolean');
            }
            if (data.hasOwnProperty('startDate')) {
                obj['startDate'] = ApiClient.convertToType(data['startDate'], 'Date');
            }
            if (data.hasOwnProperty('endDate')) {
                obj['endDate'] = ApiClient.convertToType(data['endDate'], 'Date');
            }
            if (data.hasOwnProperty('refStart')) {
                obj['refStart'] = ApiClient.convertToType(data['refStart'], 'Date');
            }
            if (data.hasOwnProperty('refEnd')) {
                obj['refEnd'] = ApiClient.convertToType(data['refEnd'], 'Date');
            }
            if (data.hasOwnProperty('estimatedComputePrice')) {
                obj['estimatedComputePrice'] = ApiClient.convertToType(data['estimatedComputePrice'], 'Number');
            }
            if (data.hasOwnProperty('estimatedComputeCost')) {
                obj['estimatedComputeCost'] = ApiClient.convertToType(data['estimatedComputeCost'], 'Number');
            }
            if (data.hasOwnProperty('estimatedMemoryPrice')) {
                obj['estimatedMemoryPrice'] = ApiClient.convertToType(data['estimatedMemoryPrice'], 'Number');
            }
            if (data.hasOwnProperty('estimatedMemoryCost')) {
                obj['estimatedMemoryCost'] = ApiClient.convertToType(data['estimatedMemoryCost'], 'Number');
            }
            if (data.hasOwnProperty('estimatedStoragePrice')) {
                obj['estimatedStoragePrice'] = ApiClient.convertToType(data['estimatedStoragePrice'], 'Number');
            }
            if (data.hasOwnProperty('estimatedStorageCost')) {
                obj['estimatedStorageCost'] = ApiClient.convertToType(data['estimatedStorageCost'], 'Number');
            }
            if (data.hasOwnProperty('estimatedNetworkPrice')) {
                obj['estimatedNetworkPrice'] = ApiClient.convertToType(data['estimatedNetworkPrice'], 'Number');
            }
            if (data.hasOwnProperty('estimatedNetworkCost')) {
                obj['estimatedNetworkCost'] = ApiClient.convertToType(data['estimatedNetworkCost'], 'Number');
            }
            if (data.hasOwnProperty('estimatedLicensePrice')) {
                obj['estimatedLicensePrice'] = ApiClient.convertToType(data['estimatedLicensePrice'], 'Number');
            }
            if (data.hasOwnProperty('estimatedLicenseCost')) {
                obj['estimatedLicenseCost'] = ApiClient.convertToType(data['estimatedLicenseCost'], 'Number');
            }
            if (data.hasOwnProperty('estimatedExtraPrice')) {
                obj['estimatedExtraPrice'] = ApiClient.convertToType(data['estimatedExtraPrice'], 'Number');
            }
            if (data.hasOwnProperty('estimatedExtraCost')) {
                obj['estimatedExtraCost'] = ApiClient.convertToType(data['estimatedExtraCost'], 'Number');
            }
            if (data.hasOwnProperty('estimatedTotalPrice')) {
                obj['estimatedTotalPrice'] = ApiClient.convertToType(data['estimatedTotalPrice'], 'Number');
            }
            if (data.hasOwnProperty('estimatedTotalCost')) {
                obj['estimatedTotalCost'] = ApiClient.convertToType(data['estimatedTotalCost'], 'Number');
            }
            if (data.hasOwnProperty('estimatedRunningPrice')) {
                obj['estimatedRunningPrice'] = ApiClient.convertToType(data['estimatedRunningPrice'], 'Number');
            }
            if (data.hasOwnProperty('estimatedRunningCost')) {
                obj['estimatedRunningCost'] = ApiClient.convertToType(data['estimatedRunningCost'], 'Number');
            }
            if (data.hasOwnProperty('estimatedCurrency')) {
                obj['estimatedCurrency'] = ApiClient.convertToType(data['estimatedCurrency'], 'String');
            }
            if (data.hasOwnProperty('estimatedConversionRate')) {
                obj['estimatedConversionRate'] = ApiClient.convertToType(data['estimatedConversionRate'], 'Number');
            }
            if (data.hasOwnProperty('actualComputePrice')) {
                obj['actualComputePrice'] = ApiClient.convertToType(data['actualComputePrice'], 'Number');
            }
            if (data.hasOwnProperty('actualComputeCost')) {
                obj['actualComputeCost'] = ApiClient.convertToType(data['actualComputeCost'], 'Number');
            }
            if (data.hasOwnProperty('actualMemoryPrice')) {
                obj['actualMemoryPrice'] = ApiClient.convertToType(data['actualMemoryPrice'], 'Number');
            }
            if (data.hasOwnProperty('actualMemoryCost')) {
                obj['actualMemoryCost'] = ApiClient.convertToType(data['actualMemoryCost'], 'Number');
            }
            if (data.hasOwnProperty('actualStoragePrice')) {
                obj['actualStoragePrice'] = ApiClient.convertToType(data['actualStoragePrice'], 'Number');
            }
            if (data.hasOwnProperty('actualStorageCost')) {
                obj['actualStorageCost'] = ApiClient.convertToType(data['actualStorageCost'], 'Number');
            }
            if (data.hasOwnProperty('actualNetworkPrice')) {
                obj['actualNetworkPrice'] = ApiClient.convertToType(data['actualNetworkPrice'], 'Number');
            }
            if (data.hasOwnProperty('actualNetworkCost')) {
                obj['actualNetworkCost'] = ApiClient.convertToType(data['actualNetworkCost'], 'Number');
            }
            if (data.hasOwnProperty('actualLicensePrice')) {
                obj['actualLicensePrice'] = ApiClient.convertToType(data['actualLicensePrice'], 'Number');
            }
            if (data.hasOwnProperty('actualLicenseCost')) {
                obj['actualLicenseCost'] = ApiClient.convertToType(data['actualLicenseCost'], 'Number');
            }
            if (data.hasOwnProperty('actualExtraPrice')) {
                obj['actualExtraPrice'] = ApiClient.convertToType(data['actualExtraPrice'], 'Number');
            }
            if (data.hasOwnProperty('actualExtraCost')) {
                obj['actualExtraCost'] = ApiClient.convertToType(data['actualExtraCost'], 'Number');
            }
            if (data.hasOwnProperty('actualTotalPrice')) {
                obj['actualTotalPrice'] = ApiClient.convertToType(data['actualTotalPrice'], 'Number');
            }
            if (data.hasOwnProperty('actualTotalCost')) {
                obj['actualTotalCost'] = ApiClient.convertToType(data['actualTotalCost'], 'Number');
            }
            if (data.hasOwnProperty('actualRunningPrice')) {
                obj['actualRunningPrice'] = ApiClient.convertToType(data['actualRunningPrice'], 'Number');
            }
            if (data.hasOwnProperty('actualRunningCost')) {
                obj['actualRunningCost'] = ApiClient.convertToType(data['actualRunningCost'], 'Number');
            }
            if (data.hasOwnProperty('actualCurrency')) {
                obj['actualCurrency'] = ApiClient.convertToType(data['actualCurrency'], 'String');
            }
            if (data.hasOwnProperty('actualConversionRate')) {
                obj['actualConversionRate'] = ApiClient.convertToType(data['actualConversionRate'], 'Number');
            }
            if (data.hasOwnProperty('computePrice')) {
                obj['computePrice'] = ApiClient.convertToType(data['computePrice'], 'Number');
            }
            if (data.hasOwnProperty('computeCost')) {
                obj['computeCost'] = ApiClient.convertToType(data['computeCost'], 'Number');
            }
            if (data.hasOwnProperty('memoryPrice')) {
                obj['memoryPrice'] = ApiClient.convertToType(data['memoryPrice'], 'Number');
            }
            if (data.hasOwnProperty('memoryCost')) {
                obj['memoryCost'] = ApiClient.convertToType(data['memoryCost'], 'Number');
            }
            if (data.hasOwnProperty('storagePrice')) {
                obj['storagePrice'] = ApiClient.convertToType(data['storagePrice'], 'Number');
            }
            if (data.hasOwnProperty('storageCost')) {
                obj['storageCost'] = ApiClient.convertToType(data['storageCost'], 'Number');
            }
            if (data.hasOwnProperty('networkPrice')) {
                obj['networkPrice'] = ApiClient.convertToType(data['networkPrice'], 'Number');
            }
            if (data.hasOwnProperty('networkCost')) {
                obj['networkCost'] = ApiClient.convertToType(data['networkCost'], 'Number');
            }
            if (data.hasOwnProperty('licensePrice')) {
                obj['licensePrice'] = ApiClient.convertToType(data['licensePrice'], 'Number');
            }
            if (data.hasOwnProperty('licenseCost')) {
                obj['licenseCost'] = ApiClient.convertToType(data['licenseCost'], 'Number');
            }
            if (data.hasOwnProperty('extraPrice')) {
                obj['extraPrice'] = ApiClient.convertToType(data['extraPrice'], 'Number');
            }
            if (data.hasOwnProperty('extraCost')) {
                obj['extraCost'] = ApiClient.convertToType(data['extraCost'], 'Number');
            }
            if (data.hasOwnProperty('totalPrice')) {
                obj['totalPrice'] = ApiClient.convertToType(data['totalPrice'], 'Number');
            }
            if (data.hasOwnProperty('totalCost')) {
                obj['totalCost'] = ApiClient.convertToType(data['totalCost'], 'Number');
            }
            if (data.hasOwnProperty('runningPrice')) {
                obj['runningPrice'] = ApiClient.convertToType(data['runningPrice'], 'Number');
            }
            if (data.hasOwnProperty('runningCost')) {
                obj['runningCost'] = ApiClient.convertToType(data['runningCost'], 'Number');
            }
            if (data.hasOwnProperty('currency')) {
                obj['currency'] = ApiClient.convertToType(data['currency'], 'String');
            }
            if (data.hasOwnProperty('conversionRate')) {
                obj['conversionRate'] = ApiClient.convertToType(data['conversionRate'], 'Number');
            }
            if (data.hasOwnProperty('costType')) {
                obj['costType'] = ApiClient.convertToType(data['costType'], 'String');
            }
            if (data.hasOwnProperty('offTime')) {
                obj['offTime'] = ApiClient.convertToType(data['offTime'], 'Number');
            }
            if (data.hasOwnProperty('powerState')) {
                obj['powerState'] = ApiClient.convertToType(data['powerState'], 'String');
            }
            if (data.hasOwnProperty('powerDate')) {
                obj['powerDate'] = ApiClient.convertToType(data['powerDate'], 'Date');
            }
            if (data.hasOwnProperty('runningMultiplier')) {
                obj['runningMultiplier'] = ApiClient.convertToType(data['runningMultiplier'], 'Number');
            }
            if (data.hasOwnProperty('usageType')) {
                obj['usageType'] = ApiClient.convertToType(data['usageType'], 'String');
            }
            if (data.hasOwnProperty('usageCategory')) {
                obj['usageCategory'] = ApiClient.convertToType(data['usageCategory'], 'String');
            }
            if (data.hasOwnProperty('lastCostDate')) {
                obj['lastCostDate'] = ApiClient.convertToType(data['lastCostDate'], 'Date');
            }
            if (data.hasOwnProperty('lastActualDate')) {
                obj['lastActualDate'] = ApiClient.convertToType(data['lastActualDate'], 'Date');
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
            if (data.hasOwnProperty('lineItemCount')) {
                obj['lineItemCount'] = ApiClient.convertToType(data['lineItemCount'], 'Number');
            }
            if (data.hasOwnProperty('lineItems')) {
                obj['lineItems'] = ApiClient.convertToType(data['lineItems'], [InvoiceLineItems]);
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
Invoice.prototype['id'] = undefined;

/**
 * @member {Number} ownerId
 */
Invoice.prototype['ownerId'] = undefined;

/**
 * @member {module:model/InlineResponse20040AppDeployInstance} account
 */
Invoice.prototype['account'] = undefined;

/**
 * @member {Object} group
 */
Invoice.prototype['group'] = undefined;

/**
 * @member {module:model/InvoiceCloud} cloud
 */
Invoice.prototype['cloud'] = undefined;

/**
 * @member {Object} instance
 */
Invoice.prototype['instance'] = undefined;

/**
 * @member {String} server
 */
Invoice.prototype['server'] = undefined;

/**
 * @member {String} cluster
 */
Invoice.prototype['cluster'] = undefined;

/**
 * @member {Object} user
 */
Invoice.prototype['user'] = undefined;

/**
 * @member {Object} plan
 */
Invoice.prototype['plan'] = undefined;

/**
 * @member {Array.<Object>} tags
 */
Invoice.prototype['tags'] = undefined;

/**
 * @member {String} project
 */
Invoice.prototype['project'] = undefined;

/**
 * @member {String} refType
 */
Invoice.prototype['refType'] = undefined;

/**
 * @member {Number} refId
 */
Invoice.prototype['refId'] = undefined;

/**
 * @member {String} refUuid
 */
Invoice.prototype['refUuid'] = undefined;

/**
 * @member {String} refName
 */
Invoice.prototype['refName'] = undefined;

/**
 * @member {String} refCategory
 */
Invoice.prototype['refCategory'] = undefined;

/**
 * @member {String} resourceId
 */
Invoice.prototype['resourceId'] = undefined;

/**
 * @member {String} resourceUuid
 */
Invoice.prototype['resourceUuid'] = undefined;

/**
 * @member {String} resourceType
 */
Invoice.prototype['resourceType'] = undefined;

/**
 * @member {String} resourceName
 */
Invoice.prototype['resourceName'] = undefined;

/**
 * @member {String} resourceExternalId
 */
Invoice.prototype['resourceExternalId'] = undefined;

/**
 * @member {String} resourceInternalId
 */
Invoice.prototype['resourceInternalId'] = undefined;

/**
 * @member {String} interval
 */
Invoice.prototype['interval'] = undefined;

/**
 * @member {String} period
 */
Invoice.prototype['period'] = undefined;

/**
 * @member {Boolean} estimate
 */
Invoice.prototype['estimate'] = undefined;

/**
 * @member {Boolean} summaryInvoice
 */
Invoice.prototype['summaryInvoice'] = undefined;

/**
 * @member {Boolean} active
 */
Invoice.prototype['active'] = undefined;

/**
 * @member {Date} startDate
 */
Invoice.prototype['startDate'] = undefined;

/**
 * @member {Date} endDate
 */
Invoice.prototype['endDate'] = undefined;

/**
 * @member {Date} refStart
 */
Invoice.prototype['refStart'] = undefined;

/**
 * @member {Date} refEnd
 */
Invoice.prototype['refEnd'] = undefined;

/**
 * @member {Number} estimatedComputePrice
 */
Invoice.prototype['estimatedComputePrice'] = undefined;

/**
 * @member {Number} estimatedComputeCost
 */
Invoice.prototype['estimatedComputeCost'] = undefined;

/**
 * @member {Number} estimatedMemoryPrice
 */
Invoice.prototype['estimatedMemoryPrice'] = undefined;

/**
 * @member {Number} estimatedMemoryCost
 */
Invoice.prototype['estimatedMemoryCost'] = undefined;

/**
 * @member {Number} estimatedStoragePrice
 */
Invoice.prototype['estimatedStoragePrice'] = undefined;

/**
 * @member {Number} estimatedStorageCost
 */
Invoice.prototype['estimatedStorageCost'] = undefined;

/**
 * @member {Number} estimatedNetworkPrice
 */
Invoice.prototype['estimatedNetworkPrice'] = undefined;

/**
 * @member {Number} estimatedNetworkCost
 */
Invoice.prototype['estimatedNetworkCost'] = undefined;

/**
 * @member {Number} estimatedLicensePrice
 */
Invoice.prototype['estimatedLicensePrice'] = undefined;

/**
 * @member {Number} estimatedLicenseCost
 */
Invoice.prototype['estimatedLicenseCost'] = undefined;

/**
 * @member {Number} estimatedExtraPrice
 */
Invoice.prototype['estimatedExtraPrice'] = undefined;

/**
 * @member {Number} estimatedExtraCost
 */
Invoice.prototype['estimatedExtraCost'] = undefined;

/**
 * @member {Number} estimatedTotalPrice
 */
Invoice.prototype['estimatedTotalPrice'] = undefined;

/**
 * @member {Number} estimatedTotalCost
 */
Invoice.prototype['estimatedTotalCost'] = undefined;

/**
 * @member {Number} estimatedRunningPrice
 */
Invoice.prototype['estimatedRunningPrice'] = undefined;

/**
 * @member {Number} estimatedRunningCost
 */
Invoice.prototype['estimatedRunningCost'] = undefined;

/**
 * @member {String} estimatedCurrency
 */
Invoice.prototype['estimatedCurrency'] = undefined;

/**
 * @member {Number} estimatedConversionRate
 */
Invoice.prototype['estimatedConversionRate'] = undefined;

/**
 * @member {Number} actualComputePrice
 */
Invoice.prototype['actualComputePrice'] = undefined;

/**
 * @member {Number} actualComputeCost
 */
Invoice.prototype['actualComputeCost'] = undefined;

/**
 * @member {Number} actualMemoryPrice
 */
Invoice.prototype['actualMemoryPrice'] = undefined;

/**
 * @member {Number} actualMemoryCost
 */
Invoice.prototype['actualMemoryCost'] = undefined;

/**
 * @member {Number} actualStoragePrice
 */
Invoice.prototype['actualStoragePrice'] = undefined;

/**
 * @member {Number} actualStorageCost
 */
Invoice.prototype['actualStorageCost'] = undefined;

/**
 * @member {Number} actualNetworkPrice
 */
Invoice.prototype['actualNetworkPrice'] = undefined;

/**
 * @member {Number} actualNetworkCost
 */
Invoice.prototype['actualNetworkCost'] = undefined;

/**
 * @member {Number} actualLicensePrice
 */
Invoice.prototype['actualLicensePrice'] = undefined;

/**
 * @member {Number} actualLicenseCost
 */
Invoice.prototype['actualLicenseCost'] = undefined;

/**
 * @member {Number} actualExtraPrice
 */
Invoice.prototype['actualExtraPrice'] = undefined;

/**
 * @member {Number} actualExtraCost
 */
Invoice.prototype['actualExtraCost'] = undefined;

/**
 * @member {Number} actualTotalPrice
 */
Invoice.prototype['actualTotalPrice'] = undefined;

/**
 * @member {Number} actualTotalCost
 */
Invoice.prototype['actualTotalCost'] = undefined;

/**
 * @member {Number} actualRunningPrice
 */
Invoice.prototype['actualRunningPrice'] = undefined;

/**
 * @member {Number} actualRunningCost
 */
Invoice.prototype['actualRunningCost'] = undefined;

/**
 * @member {String} actualCurrency
 */
Invoice.prototype['actualCurrency'] = undefined;

/**
 * @member {Number} actualConversionRate
 */
Invoice.prototype['actualConversionRate'] = undefined;

/**
 * @member {Number} computePrice
 */
Invoice.prototype['computePrice'] = undefined;

/**
 * @member {Number} computeCost
 */
Invoice.prototype['computeCost'] = undefined;

/**
 * @member {Number} memoryPrice
 */
Invoice.prototype['memoryPrice'] = undefined;

/**
 * @member {Number} memoryCost
 */
Invoice.prototype['memoryCost'] = undefined;

/**
 * @member {Number} storagePrice
 */
Invoice.prototype['storagePrice'] = undefined;

/**
 * @member {Number} storageCost
 */
Invoice.prototype['storageCost'] = undefined;

/**
 * @member {Number} networkPrice
 */
Invoice.prototype['networkPrice'] = undefined;

/**
 * @member {Number} networkCost
 */
Invoice.prototype['networkCost'] = undefined;

/**
 * @member {Number} licensePrice
 */
Invoice.prototype['licensePrice'] = undefined;

/**
 * @member {Number} licenseCost
 */
Invoice.prototype['licenseCost'] = undefined;

/**
 * @member {Number} extraPrice
 */
Invoice.prototype['extraPrice'] = undefined;

/**
 * @member {Number} extraCost
 */
Invoice.prototype['extraCost'] = undefined;

/**
 * @member {Number} totalPrice
 */
Invoice.prototype['totalPrice'] = undefined;

/**
 * @member {Number} totalCost
 */
Invoice.prototype['totalCost'] = undefined;

/**
 * @member {Number} runningPrice
 */
Invoice.prototype['runningPrice'] = undefined;

/**
 * @member {Number} runningCost
 */
Invoice.prototype['runningCost'] = undefined;

/**
 * @member {String} currency
 */
Invoice.prototype['currency'] = undefined;

/**
 * @member {Number} conversionRate
 */
Invoice.prototype['conversionRate'] = undefined;

/**
 * @member {String} costType
 */
Invoice.prototype['costType'] = undefined;

/**
 * @member {Number} offTime
 */
Invoice.prototype['offTime'] = undefined;

/**
 * @member {String} powerState
 */
Invoice.prototype['powerState'] = undefined;

/**
 * @member {Date} powerDate
 */
Invoice.prototype['powerDate'] = undefined;

/**
 * @member {Number} runningMultiplier
 */
Invoice.prototype['runningMultiplier'] = undefined;

/**
 * @member {String} usageType
 */
Invoice.prototype['usageType'] = undefined;

/**
 * @member {String} usageCategory
 */
Invoice.prototype['usageCategory'] = undefined;

/**
 * @member {Date} lastCostDate
 */
Invoice.prototype['lastCostDate'] = undefined;

/**
 * @member {Date} lastActualDate
 */
Invoice.prototype['lastActualDate'] = undefined;

/**
 * @member {Date} dateCreated
 */
Invoice.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
Invoice.prototype['lastUpdated'] = undefined;

/**
 * @member {Number} lineItemCount
 */
Invoice.prototype['lineItemCount'] = undefined;

/**
 * @member {Array.<module:model/InvoiceLineItems>} lineItems
 */
Invoice.prototype['lineItems'] = undefined;






export default Invoice;

