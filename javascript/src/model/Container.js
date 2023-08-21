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
import ContainerContainerType from './ContainerContainerType';
import ContainerContainerTypeSet from './ContainerContainerTypeSet';
import ContainerInstance from './ContainerInstance';
import ContainerPlan from './ContainerPlan';
import ContainerPort from './ContainerPort';
import ContainerStats from './ContainerStats';

/**
 * The Container model module.
 * @module model/Container
 * @version 6.2.1
 */
class Container {
    /**
     * Constructs a new <code>Container</code>.
     * @alias module:model/Container
     */
    constructor() { 
        
        Container.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Container</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Container} obj Optional instance to populate.
     * @return {module:model/Container} The populated <code>Container</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Container();

            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'Number');
            }
            if (data.hasOwnProperty('uuid')) {
                obj['uuid'] = ApiClient.convertToType(data['uuid'], 'String');
            }
            if (data.hasOwnProperty('accountId')) {
                obj['accountId'] = ApiClient.convertToType(data['accountId'], 'Number');
            }
            if (data.hasOwnProperty('instance')) {
                obj['instance'] = ContainerInstance.constructFromObject(data['instance']);
            }
            if (data.hasOwnProperty('containerType')) {
                obj['containerType'] = ContainerContainerType.constructFromObject(data['containerType']);
            }
            if (data.hasOwnProperty('containerTypeSet')) {
                obj['containerTypeSet'] = ContainerContainerTypeSet.constructFromObject(data['containerTypeSet']);
            }
            if (data.hasOwnProperty('server')) {
                obj['server'] = ContainerInstance.constructFromObject(data['server']);
            }
            if (data.hasOwnProperty('cloud')) {
                obj['cloud'] = ContainerInstance.constructFromObject(data['cloud']);
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('ip')) {
                obj['ip'] = ApiClient.convertToType(data['ip'], 'String');
            }
            if (data.hasOwnProperty('internalIp')) {
                obj['internalIp'] = ApiClient.convertToType(data['internalIp'], 'String');
            }
            if (data.hasOwnProperty('internalHostname')) {
                obj['internalHostname'] = ApiClient.convertToType(data['internalHostname'], 'String');
            }
            if (data.hasOwnProperty('externalHostname')) {
                obj['externalHostname'] = ApiClient.convertToType(data['externalHostname'], 'String');
            }
            if (data.hasOwnProperty('externalDomain')) {
                obj['externalDomain'] = ApiClient.convertToType(data['externalDomain'], 'String');
            }
            if (data.hasOwnProperty('externalFqdn')) {
                obj['externalFqdn'] = ApiClient.convertToType(data['externalFqdn'], 'String');
            }
            if (data.hasOwnProperty('ports')) {
                obj['ports'] = ApiClient.convertToType(data['ports'], [ContainerPort]);
            }
            if (data.hasOwnProperty('plan')) {
                obj['plan'] = ContainerPlan.constructFromObject(data['plan']);
            }
            if (data.hasOwnProperty('dateCreated')) {
                obj['dateCreated'] = ApiClient.convertToType(data['dateCreated'], 'Date');
            }
            if (data.hasOwnProperty('lastUpdated')) {
                obj['lastUpdated'] = ApiClient.convertToType(data['lastUpdated'], 'Date');
            }
            if (data.hasOwnProperty('statsEnabled')) {
                obj['statsEnabled'] = ApiClient.convertToType(data['statsEnabled'], 'Boolean');
            }
            if (data.hasOwnProperty('status')) {
                obj['status'] = ApiClient.convertToType(data['status'], 'String');
            }
            if (data.hasOwnProperty('userStatus')) {
                obj['userStatus'] = ApiClient.convertToType(data['userStatus'], 'String');
            }
            if (data.hasOwnProperty('environmentPrefix')) {
                obj['environmentPrefix'] = ApiClient.convertToType(data['environmentPrefix'], 'String');
            }
            if (data.hasOwnProperty('stats')) {
                obj['stats'] = ContainerStats.constructFromObject(data['stats']);
            }
            if (data.hasOwnProperty('runtimeInfo')) {
                obj['runtimeInfo'] = ApiClient.convertToType(data['runtimeInfo'], Object);
            }
            if (data.hasOwnProperty('containerVersion')) {
                obj['containerVersion'] = ApiClient.convertToType(data['containerVersion'], 'String');
            }
            if (data.hasOwnProperty('repositoryImage')) {
                obj['repositoryImage'] = ApiClient.convertToType(data['repositoryImage'], 'String');
            }
            if (data.hasOwnProperty('planCategory')) {
                obj['planCategory'] = ApiClient.convertToType(data['planCategory'], 'String');
            }
            if (data.hasOwnProperty('hostname')) {
                obj['hostname'] = ApiClient.convertToType(data['hostname'], 'String');
            }
            if (data.hasOwnProperty('domainName')) {
                obj['domainName'] = ApiClient.convertToType(data['domainName'], 'String');
            }
            if (data.hasOwnProperty('volumeCreated')) {
                obj['volumeCreated'] = ApiClient.convertToType(data['volumeCreated'], 'Boolean');
            }
            if (data.hasOwnProperty('containerCreated')) {
                obj['containerCreated'] = ApiClient.convertToType(data['containerCreated'], 'Boolean');
            }
            if (data.hasOwnProperty('maxStorage')) {
                obj['maxStorage'] = ApiClient.convertToType(data['maxStorage'], 'Number');
            }
            if (data.hasOwnProperty('maxMemory')) {
                obj['maxMemory'] = ApiClient.convertToType(data['maxMemory'], 'Number');
            }
            if (data.hasOwnProperty('maxCores')) {
                obj['maxCores'] = ApiClient.convertToType(data['maxCores'], 'Number');
            }
            if (data.hasOwnProperty('maxCpu')) {
                obj['maxCpu'] = ApiClient.convertToType(data['maxCpu'], 'Number');
            }
            if (data.hasOwnProperty('availableActions')) {
                obj['availableActions'] = ApiClient.convertToType(data['availableActions'], [Object]);
            }
            if (data.hasOwnProperty('configGroup')) {
                obj['configGroup'] = ApiClient.convertToType(data['configGroup'], 'String');
            }
            if (data.hasOwnProperty('configId')) {
                obj['configId'] = ApiClient.convertToType(data['configId'], 'String');
            }
            if (data.hasOwnProperty('configRole')) {
                obj['configRole'] = ApiClient.convertToType(data['configRole'], 'String');
            }
            if (data.hasOwnProperty('hourlyCost')) {
                obj['hourlyCost'] = ApiClient.convertToType(data['hourlyCost'], 'Number');
            }
            if (data.hasOwnProperty('hourlyPrice')) {
                obj['hourlyPrice'] = ApiClient.convertToType(data['hourlyPrice'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * @member {Number} id
 */
Container.prototype['id'] = undefined;

/**
 * @member {String} uuid
 */
Container.prototype['uuid'] = undefined;

/**
 * @member {Number} accountId
 */
Container.prototype['accountId'] = undefined;

/**
 * @member {module:model/ContainerInstance} instance
 */
Container.prototype['instance'] = undefined;

/**
 * @member {module:model/ContainerContainerType} containerType
 */
Container.prototype['containerType'] = undefined;

/**
 * @member {module:model/ContainerContainerTypeSet} containerTypeSet
 */
Container.prototype['containerTypeSet'] = undefined;

/**
 * @member {module:model/ContainerInstance} server
 */
Container.prototype['server'] = undefined;

/**
 * @member {module:model/ContainerInstance} cloud
 */
Container.prototype['cloud'] = undefined;

/**
 * @member {String} name
 */
Container.prototype['name'] = undefined;

/**
 * @member {String} ip
 */
Container.prototype['ip'] = undefined;

/**
 * @member {String} internalIp
 */
Container.prototype['internalIp'] = undefined;

/**
 * @member {String} internalHostname
 */
Container.prototype['internalHostname'] = undefined;

/**
 * @member {String} externalHostname
 */
Container.prototype['externalHostname'] = undefined;

/**
 * @member {String} externalDomain
 */
Container.prototype['externalDomain'] = undefined;

/**
 * @member {String} externalFqdn
 */
Container.prototype['externalFqdn'] = undefined;

/**
 * @member {Array.<module:model/ContainerPort>} ports
 */
Container.prototype['ports'] = undefined;

/**
 * @member {module:model/ContainerPlan} plan
 */
Container.prototype['plan'] = undefined;

/**
 * @member {Date} dateCreated
 */
Container.prototype['dateCreated'] = undefined;

/**
 * @member {Date} lastUpdated
 */
Container.prototype['lastUpdated'] = undefined;

/**
 * @member {Boolean} statsEnabled
 */
Container.prototype['statsEnabled'] = undefined;

/**
 * @member {String} status
 */
Container.prototype['status'] = undefined;

/**
 * @member {String} userStatus
 */
Container.prototype['userStatus'] = undefined;

/**
 * @member {String} environmentPrefix
 */
Container.prototype['environmentPrefix'] = undefined;

/**
 * @member {module:model/ContainerStats} stats
 */
Container.prototype['stats'] = undefined;

/**
 * @member {Object} runtimeInfo
 */
Container.prototype['runtimeInfo'] = undefined;

/**
 * @member {String} containerVersion
 */
Container.prototype['containerVersion'] = undefined;

/**
 * @member {String} repositoryImage
 */
Container.prototype['repositoryImage'] = undefined;

/**
 * @member {String} planCategory
 */
Container.prototype['planCategory'] = undefined;

/**
 * @member {String} hostname
 */
Container.prototype['hostname'] = undefined;

/**
 * @member {String} domainName
 */
Container.prototype['domainName'] = undefined;

/**
 * @member {Boolean} volumeCreated
 */
Container.prototype['volumeCreated'] = undefined;

/**
 * @member {Boolean} containerCreated
 */
Container.prototype['containerCreated'] = undefined;

/**
 * @member {Number} maxStorage
 */
Container.prototype['maxStorage'] = undefined;

/**
 * @member {Number} maxMemory
 */
Container.prototype['maxMemory'] = undefined;

/**
 * @member {Number} maxCores
 */
Container.prototype['maxCores'] = undefined;

/**
 * @member {Number} maxCpu
 */
Container.prototype['maxCpu'] = undefined;

/**
 * @member {Array.<Object>} availableActions
 */
Container.prototype['availableActions'] = undefined;

/**
 * @member {String} configGroup
 */
Container.prototype['configGroup'] = undefined;

/**
 * @member {String} configId
 */
Container.prototype['configId'] = undefined;

/**
 * @member {String} configRole
 */
Container.prototype['configRole'] = undefined;

/**
 * @member {Number} hourlyCost
 */
Container.prototype['hourlyCost'] = undefined;

/**
 * @member {Number} hourlyPrice
 */
Container.prototype['hourlyPrice'] = undefined;






export default Container;

