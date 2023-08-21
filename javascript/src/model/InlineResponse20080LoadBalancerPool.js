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
import InlineResponse20079LoadBalancerMonitorLoadBalancer from './InlineResponse20079LoadBalancerMonitorLoadBalancer';

/**
 * The InlineResponse20080LoadBalancerPool model module.
 * @module model/InlineResponse20080LoadBalancerPool
 * @version 6.2.1
 */
class InlineResponse20080LoadBalancerPool {
    /**
     * Constructs a new <code>InlineResponse20080LoadBalancerPool</code>.
     * @alias module:model/InlineResponse20080LoadBalancerPool
     */
    constructor() { 
        
        InlineResponse20080LoadBalancerPool.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InlineResponse20080LoadBalancerPool</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InlineResponse20080LoadBalancerPool} obj Optional instance to populate.
     * @return {module:model/InlineResponse20080LoadBalancerPool} The populated <code>InlineResponse20080LoadBalancerPool</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InlineResponse20080LoadBalancerPool();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('loadBalancer')) {
                obj['loadBalancer'] = InlineResponse20079LoadBalancerMonitorLoadBalancer.constructFromObject(data['loadBalancer']);
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('category')) {
                obj['category'] = ApiClient.convertToType(data['category'], 'String');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('internalId')) {
                obj['internalId'] = ApiClient.convertToType(data['internalId'], 'String');
            }
            if (data.hasOwnProperty('externalId')) {
                obj['externalId'] = ApiClient.convertToType(data['externalId'], 'String');
            }
            if (data.hasOwnProperty('enabled')) {
                obj['enabled'] = ApiClient.convertToType(data['enabled'], 'Boolean');
            }
            if (data.hasOwnProperty('vipSticky')) {
                obj['vipSticky'] = ApiClient.convertToType(data['vipSticky'], 'String');
            }
            if (data.hasOwnProperty('vipBalance')) {
                obj['vipBalance'] = ApiClient.convertToType(data['vipBalance'], 'String');
            }
            if (data.hasOwnProperty('allowNat')) {
                obj['allowNat'] = ApiClient.convertToType(data['allowNat'], 'String');
            }
            if (data.hasOwnProperty('allowSnat')) {
                obj['allowSnat'] = ApiClient.convertToType(data['allowSnat'], 'String');
            }
            if (data.hasOwnProperty('vipClientIpMode')) {
                obj['vipClientIpMode'] = ApiClient.convertToType(data['vipClientIpMode'], 'String');
            }
            if (data.hasOwnProperty('vipServerIpMode')) {
                obj['vipServerIpMode'] = ApiClient.convertToType(data['vipServerIpMode'], 'String');
            }
            if (data.hasOwnProperty('minActive')) {
                obj['minActive'] = ApiClient.convertToType(data['minActive'], 'Number');
            }
            if (data.hasOwnProperty('minInService')) {
                obj['minInService'] = ApiClient.convertToType(data['minInService'], 'String');
            }
            if (data.hasOwnProperty('minUpMonitor')) {
                obj['minUpMonitor'] = ApiClient.convertToType(data['minUpMonitor'], 'String');
            }
            if (data.hasOwnProperty('minUpAction')) {
                obj['minUpAction'] = ApiClient.convertToType(data['minUpAction'], 'String');
            }
            if (data.hasOwnProperty('maxQueueDepth')) {
                obj['maxQueueDepth'] = ApiClient.convertToType(data['maxQueueDepth'], 'String');
            }
            if (data.hasOwnProperty('maxQueueTime')) {
                obj['maxQueueTime'] = ApiClient.convertToType(data['maxQueueTime'], 'String');
            }
            if (data.hasOwnProperty('numberActive')) {
                obj['numberActive'] = ApiClient.convertToType(data['numberActive'], 'Number');
            }
            if (data.hasOwnProperty('numberInService')) {
                obj['numberInService'] = ApiClient.convertToType(data['numberInService'], 'Number');
            }
            if (data.hasOwnProperty('healthScore')) {
                obj['healthScore'] = ApiClient.convertToType(data['healthScore'], 'Number');
            }
            if (data.hasOwnProperty('performanceScore')) {
                obj['performanceScore'] = ApiClient.convertToType(data['performanceScore'], 'Number');
            }
            if (data.hasOwnProperty('healthPenalty')) {
                obj['healthPenalty'] = ApiClient.convertToType(data['healthPenalty'], 'Number');
            }
            if (data.hasOwnProperty('securityPenalty')) {
                obj['securityPenalty'] = ApiClient.convertToType(data['securityPenalty'], 'Number');
            }
            if (data.hasOwnProperty('errorPenalty')) {
                obj['errorPenalty'] = ApiClient.convertToType(data['errorPenalty'], 'Number');
            }
            if (data.hasOwnProperty('downAction')) {
                obj['downAction'] = ApiClient.convertToType(data['downAction'], 'String');
            }
            if (data.hasOwnProperty('rampTime')) {
                obj['rampTime'] = ApiClient.convertToType(data['rampTime'], 'String');
            }
            if (data.hasOwnProperty('port')) {
                obj['port'] = ApiClient.convertToType(data['port'], 'String');
            }
            if (data.hasOwnProperty('portType')) {
                obj['portType'] = ApiClient.convertToType(data['portType'], 'String');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('nodes')) {
                obj['nodes'] = ApiClient.convertToType(data['nodes'], [InlineResponse20040AppDeployInstance]);
            }
            if (data.hasOwnProperty('monitors')) {
                obj['monitors'] = ApiClient.convertToType(data['monitors'], [InlineResponse20040AppDeployInstance]);
            }
            if (data.hasOwnProperty('members')) {
                obj['members'] = ApiClient.convertToType(data['members'], [Object]);
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = ApiClient.convertToType(data['config'], Object);
            }
            if (data.hasOwnProperty('createdBy')) {
                obj['createdBy'] = ApiClient.convertToType(data['createdBy'], 'String');
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
InlineResponse20080LoadBalancerPool.prototype['id'] = undefined;

/**
 * @member {module:model/InlineResponse20079LoadBalancerMonitorLoadBalancer} loadBalancer
 */
InlineResponse20080LoadBalancerPool.prototype['loadBalancer'] = undefined;

/**
 * @member {String} name
 */
InlineResponse20080LoadBalancerPool.prototype['name'] = undefined;

/**
 * @member {String} category
 */
InlineResponse20080LoadBalancerPool.prototype['category'] = undefined;

/**
 * @member {String} visibility
 */
InlineResponse20080LoadBalancerPool.prototype['visibility'] = undefined;

/**
 * @member {String} description
 */
InlineResponse20080LoadBalancerPool.prototype['description'] = undefined;

/**
 * @member {String} internalId
 */
InlineResponse20080LoadBalancerPool.prototype['internalId'] = undefined;

/**
 * @member {String} externalId
 */
InlineResponse20080LoadBalancerPool.prototype['externalId'] = undefined;

/**
 * @member {Boolean} enabled
 */
InlineResponse20080LoadBalancerPool.prototype['enabled'] = undefined;

/**
 * @member {String} vipSticky
 */
InlineResponse20080LoadBalancerPool.prototype['vipSticky'] = undefined;

/**
 * @member {String} vipBalance
 */
InlineResponse20080LoadBalancerPool.prototype['vipBalance'] = undefined;

/**
 * @member {String} allowNat
 */
InlineResponse20080LoadBalancerPool.prototype['allowNat'] = undefined;

/**
 * @member {String} allowSnat
 */
InlineResponse20080LoadBalancerPool.prototype['allowSnat'] = undefined;

/**
 * @member {String} vipClientIpMode
 */
InlineResponse20080LoadBalancerPool.prototype['vipClientIpMode'] = undefined;

/**
 * @member {String} vipServerIpMode
 */
InlineResponse20080LoadBalancerPool.prototype['vipServerIpMode'] = undefined;

/**
 * @member {Number} minActive
 */
InlineResponse20080LoadBalancerPool.prototype['minActive'] = undefined;

/**
 * @member {String} minInService
 */
InlineResponse20080LoadBalancerPool.prototype['minInService'] = undefined;

/**
 * @member {String} minUpMonitor
 */
InlineResponse20080LoadBalancerPool.prototype['minUpMonitor'] = undefined;

/**
 * @member {String} minUpAction
 */
InlineResponse20080LoadBalancerPool.prototype['minUpAction'] = undefined;

/**
 * @member {String} maxQueueDepth
 */
InlineResponse20080LoadBalancerPool.prototype['maxQueueDepth'] = undefined;

/**
 * @member {String} maxQueueTime
 */
InlineResponse20080LoadBalancerPool.prototype['maxQueueTime'] = undefined;

/**
 * @member {Number} numberActive
 */
InlineResponse20080LoadBalancerPool.prototype['numberActive'] = undefined;

/**
 * @member {Number} numberInService
 */
InlineResponse20080LoadBalancerPool.prototype['numberInService'] = undefined;

/**
 * @member {Number} healthScore
 */
InlineResponse20080LoadBalancerPool.prototype['healthScore'] = undefined;

/**
 * @member {Number} performanceScore
 */
InlineResponse20080LoadBalancerPool.prototype['performanceScore'] = undefined;

/**
 * @member {Number} healthPenalty
 */
InlineResponse20080LoadBalancerPool.prototype['healthPenalty'] = undefined;

/**
 * @member {Number} securityPenalty
 */
InlineResponse20080LoadBalancerPool.prototype['securityPenalty'] = undefined;

/**
 * @member {Number} errorPenalty
 */
InlineResponse20080LoadBalancerPool.prototype['errorPenalty'] = undefined;

/**
 * @member {String} downAction
 */
InlineResponse20080LoadBalancerPool.prototype['downAction'] = undefined;

/**
 * @member {String} rampTime
 */
InlineResponse20080LoadBalancerPool.prototype['rampTime'] = undefined;

/**
 * @member {String} port
 */
InlineResponse20080LoadBalancerPool.prototype['port'] = undefined;

/**
 * @member {String} portType
 */
InlineResponse20080LoadBalancerPool.prototype['portType'] = undefined;

/**
 * @member {String} status
 */
InlineResponse20080LoadBalancerPool.prototype['status'] = undefined;

/**
 * @member {Array.<module:model/InlineResponse20040AppDeployInstance>} nodes
 */
InlineResponse20080LoadBalancerPool.prototype['nodes'] = undefined;

/**
 * @member {Array.<module:model/InlineResponse20040AppDeployInstance>} monitors
 */
InlineResponse20080LoadBalancerPool.prototype['monitors'] = undefined;

/**
 * @member {Array.<Object>} members
 */
InlineResponse20080LoadBalancerPool.prototype['members'] = undefined;

/**
 * @member {Object} config
 */
InlineResponse20080LoadBalancerPool.prototype['config'] = undefined;

/**
 * @member {String} createdBy
 */
InlineResponse20080LoadBalancerPool.prototype['createdBy'] = undefined;

/**
 * @member {Date} dateCreated
 */
InlineResponse20080LoadBalancerPool.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
InlineResponse20080LoadBalancerPool.prototype['lastUpdated'] = undefined;






export default InlineResponse20080LoadBalancerPool;
