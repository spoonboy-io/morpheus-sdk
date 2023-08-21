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
import ApiSecurityGroupsIdRulesRuleDestinationGroup from './ApiSecurityGroupsIdRulesRuleDestinationGroup';
import ApiSecurityGroupsIdRulesRuleDestinationTier from './ApiSecurityGroupsIdRulesRuleDestinationTier';
import ApiSecurityGroupsIdRulesRuleSourceGroup from './ApiSecurityGroupsIdRulesRuleSourceGroup';
import ApiSecurityGroupsIdRulesRuleSourceTier from './ApiSecurityGroupsIdRulesRuleSourceTier';

/**
 * The ApiSecurityGroupsIdRulesSgIdRule model module.
 * @module model/ApiSecurityGroupsIdRulesSgIdRule
 * @version 6.2.1
 */
class ApiSecurityGroupsIdRulesSgIdRule {
    /**
     * Constructs a new <code>ApiSecurityGroupsIdRulesSgIdRule</code>.
     * @alias module:model/ApiSecurityGroupsIdRulesSgIdRule
     * @param protocol {module:model/ApiSecurityGroupsIdRulesSgIdRule.ProtocolEnum} Either tcp, udp, icmp. Required if customRule is true, otherwise, ignored.
     * @param ruleType {String} Either `customRule` or an `instance type` code.
     */
    constructor(protocol, ruleType) { 
        
        ApiSecurityGroupsIdRulesSgIdRule.initialize(this, protocol, ruleType);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, protocol, ruleType) { 
        obj['protocol'] = protocol;
        obj['ruleType'] = ruleType || 'customRule';
    }

    /**
     * Constructs a <code>ApiSecurityGroupsIdRulesSgIdRule</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ApiSecurityGroupsIdRulesSgIdRule} obj Optional instance to populate.
     * @return {module:model/ApiSecurityGroupsIdRulesSgIdRule} The populated <code>ApiSecurityGroupsIdRulesSgIdRule</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ApiSecurityGroupsIdRulesSgIdRule();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('direction')) {
                obj['direction'] = ApiClient.convertToType(data['direction'], 'String');
            }
            if (data.hasOwnProperty('sourceType')) {
                obj['sourceType'] = ApiClient.convertToType(data['sourceType'], 'String');
            }
            if (data.hasOwnProperty('source')) {
                obj['source'] = ApiClient.convertToType(data['source'], 'String');
            }
            if (data.hasOwnProperty('sourceGroup')) {
                obj['sourceGroup'] = ApiSecurityGroupsIdRulesRuleSourceGroup.constructFromObject(data['sourceGroup']);
            }
            if (data.hasOwnProperty('sourceTier')) {
                obj['sourceTier'] = ApiSecurityGroupsIdRulesRuleSourceTier.constructFromObject(data['sourceTier']);
            }
            if (data.hasOwnProperty('portRange')) {
                obj['portRange'] = ApiClient.convertToType(data['portRange'], 'String');
            }
            if (data.hasOwnProperty('protocol')) {
                obj['protocol'] = ApiClient.convertToType(data['protocol'], 'String');
            }
            if (data.hasOwnProperty('destinationType')) {
                obj['destinationType'] = ApiClient.convertToType(data['destinationType'], 'String');
            }
            if (data.hasOwnProperty('destination')) {
                obj['destination'] = ApiClient.convertToType(data['destination'], 'String');
            }
            if (data.hasOwnProperty('destinationGroup')) {
                obj['destinationGroup'] = ApiSecurityGroupsIdRulesRuleDestinationGroup.constructFromObject(data['destinationGroup']);
            }
            if (data.hasOwnProperty('destinationTier')) {
                obj['destinationTier'] = ApiSecurityGroupsIdRulesRuleDestinationTier.constructFromObject(data['destinationTier']);
            }
            if (data.hasOwnProperty('ruleType')) {
                obj['ruleType'] = ApiClient.convertToType(data['ruleType'], 'String');
            }
            if (data.hasOwnProperty('policy')) {
                obj['policy'] = ApiClient.convertToType(data['policy'], 'String');
            }
            if (data.hasOwnProperty('instanceTypeId')) {
                obj['instanceTypeId'] = ApiClient.convertToType(data['instanceTypeId'], 'Number');
            }
        }
        return obj;
    }


}

/**
 * A name for the rule
 * @member {String} name
 */
ApiSecurityGroupsIdRulesSgIdRule.prototype['name'] = undefined;

/**
 * Either `ingress` or `egress`
 * @member {module:model/ApiSecurityGroupsIdRulesSgIdRule.DirectionEnum} direction
 * @default 'ingress'
 */
ApiSecurityGroupsIdRulesSgIdRule.prototype['direction'] = 'ingress';

/**
 * Either `cidr`, `group`, `tier`, `all`.
 * @member {module:model/ApiSecurityGroupsIdRulesSgIdRule.SourceTypeEnum} sourceType
 * @default 'cidr'
 */
ApiSecurityGroupsIdRulesSgIdRule.prototype['sourceType'] = 'cidr';

/**
 * CIDR representing the source IP(s) which should receive access. Required for `sourceType`=cidr
 * @member {String} source
 */
ApiSecurityGroupsIdRulesSgIdRule.prototype['source'] = undefined;

/**
 * @member {module:model/ApiSecurityGroupsIdRulesRuleSourceGroup} sourceGroup
 */
ApiSecurityGroupsIdRulesSgIdRule.prototype['sourceGroup'] = undefined;

/**
 * @member {module:model/ApiSecurityGroupsIdRulesRuleSourceTier} sourceTier
 */
ApiSecurityGroupsIdRulesSgIdRule.prototype['sourceTier'] = undefined;

/**
 * Either a single value (i.e. 55) or a port range (i.e. 1-65535) for which to open access to the source. Required if customRule is true, otherwise, ignored. 
 * @member {String} portRange
 */
ApiSecurityGroupsIdRulesSgIdRule.prototype['portRange'] = undefined;

/**
 * Either tcp, udp, icmp. Required if customRule is true, otherwise, ignored.
 * @member {module:model/ApiSecurityGroupsIdRulesSgIdRule.ProtocolEnum} protocol
 */
ApiSecurityGroupsIdRulesSgIdRule.prototype['protocol'] = undefined;

/**
 * Either cidr, group, tier, instance.
 * @member {module:model/ApiSecurityGroupsIdRulesSgIdRule.DestinationTypeEnum} destinationType
 * @default 'cidr'
 */
ApiSecurityGroupsIdRulesSgIdRule.prototype['destinationType'] = 'cidr';

/**
 * CIDR representing the destination IP(s) which should receive access. Required for `destinationType`=cidr.
 * @member {String} destination
 */
ApiSecurityGroupsIdRulesSgIdRule.prototype['destination'] = undefined;

/**
 * @member {module:model/ApiSecurityGroupsIdRulesRuleDestinationGroup} destinationGroup
 */
ApiSecurityGroupsIdRulesSgIdRule.prototype['destinationGroup'] = undefined;

/**
 * @member {module:model/ApiSecurityGroupsIdRulesRuleDestinationTier} destinationTier
 */
ApiSecurityGroupsIdRulesSgIdRule.prototype['destinationTier'] = undefined;

/**
 * Either `customRule` or an `instance type` code.
 * @member {String} ruleType
 * @default 'customRule'
 */
ApiSecurityGroupsIdRulesSgIdRule.prototype['ruleType'] = 'customRule';

/**
 * Either `accept` or `deny`.
 * @member {module:model/ApiSecurityGroupsIdRulesSgIdRule.PolicyEnum} policy
 */
ApiSecurityGroupsIdRulesSgIdRule.prototype['policy'] = undefined;

/**
 * The id of an Instance Type. If specified, the source CIDR will have access to all ports exposed by the particular instance in the cloud, app, or instance. Required if customRule is false, otherwise ignored. 
 * @member {Number} instanceTypeId
 */
ApiSecurityGroupsIdRulesSgIdRule.prototype['instanceTypeId'] = undefined;





/**
 * Allowed values for the <code>direction</code> property.
 * @enum {String}
 * @readonly
 */
ApiSecurityGroupsIdRulesSgIdRule['DirectionEnum'] = {

    /**
     * value: "ingress"
     * @const
     */
    "ingress": "ingress",

    /**
     * value: "egress"
     * @const
     */
    "egress": "egress"
};


/**
 * Allowed values for the <code>sourceType</code> property.
 * @enum {String}
 * @readonly
 */
ApiSecurityGroupsIdRulesSgIdRule['SourceTypeEnum'] = {

    /**
     * value: "cidr"
     * @const
     */
    "cidr": "cidr",

    /**
     * value: "group"
     * @const
     */
    "group": "group",

    /**
     * value: "tier"
     * @const
     */
    "tier": "tier",

    /**
     * value: "all"
     * @const
     */
    "all": "all"
};


/**
 * Allowed values for the <code>protocol</code> property.
 * @enum {String}
 * @readonly
 */
ApiSecurityGroupsIdRulesSgIdRule['ProtocolEnum'] = {

    /**
     * value: "tcp"
     * @const
     */
    "tcp": "tcp",

    /**
     * value: "udp"
     * @const
     */
    "udp": "udp",

    /**
     * value: "icmp"
     * @const
     */
    "icmp": "icmp"
};


/**
 * Allowed values for the <code>destinationType</code> property.
 * @enum {String}
 * @readonly
 */
ApiSecurityGroupsIdRulesSgIdRule['DestinationTypeEnum'] = {

    /**
     * value: "cidr"
     * @const
     */
    "cidr": "cidr",

    /**
     * value: "group"
     * @const
     */
    "group": "group",

    /**
     * value: "tier"
     * @const
     */
    "tier": "tier",

    /**
     * value: "instance"
     * @const
     */
    "instance": "instance"
};


/**
 * Allowed values for the <code>policy</code> property.
 * @enum {String}
 * @readonly
 */
ApiSecurityGroupsIdRulesSgIdRule['PolicyEnum'] = {

    /**
     * value: "accept"
     * @const
     */
    "accept": "accept",

    /**
     * value: "deny"
     * @const
     */
    "deny": "deny"
};



export default ApiSecurityGroupsIdRulesSgIdRule;

