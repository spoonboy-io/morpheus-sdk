<?php
/**
 * ApiSecurityGroupsIdRulesSgIdRule
 *
 * PHP version 7.2
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.0.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace OpenAPI\Client\Model;

use \ArrayAccess;
use \OpenAPI\Client\ObjectSerializer;

/**
 * ApiSecurityGroupsIdRulesSgIdRule Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ApiSecurityGroupsIdRulesSgIdRule implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = '_api_security_groups__id__rules__sgId__rule';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'name' => 'string',
        'direction' => 'string',
        'source_type' => 'string',
        'source' => 'string',
        'source_group' => '\OpenAPI\Client\Model\ApiSecurityGroupsIdRulesRuleSourceGroup',
        'source_tier' => '\OpenAPI\Client\Model\ApiSecurityGroupsIdRulesRuleSourceTier',
        'port_range' => 'string',
        'protocol' => 'string',
        'destination_type' => 'string',
        'destination' => 'string',
        'destination_group' => '\OpenAPI\Client\Model\ApiSecurityGroupsIdRulesRuleDestinationGroup',
        'destination_tier' => '\OpenAPI\Client\Model\ApiSecurityGroupsIdRulesRuleDestinationTier',
        'rule_type' => 'string',
        'policy' => 'string',
        'instance_type_id' => 'int'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'name' => null,
        'direction' => null,
        'source_type' => null,
        'source' => null,
        'source_group' => null,
        'source_tier' => null,
        'port_range' => null,
        'protocol' => null,
        'destination_type' => null,
        'destination' => null,
        'destination_group' => null,
        'destination_tier' => null,
        'rule_type' => null,
        'policy' => null,
        'instance_type_id' => 'int64'
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'name' => 'name',
        'direction' => 'direction',
        'source_type' => 'sourceType',
        'source' => 'source',
        'source_group' => 'sourceGroup',
        'source_tier' => 'sourceTier',
        'port_range' => 'portRange',
        'protocol' => 'protocol',
        'destination_type' => 'destinationType',
        'destination' => 'destination',
        'destination_group' => 'destinationGroup',
        'destination_tier' => 'destinationTier',
        'rule_type' => 'ruleType',
        'policy' => 'policy',
        'instance_type_id' => 'instanceTypeId'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'name' => 'setName',
        'direction' => 'setDirection',
        'source_type' => 'setSourceType',
        'source' => 'setSource',
        'source_group' => 'setSourceGroup',
        'source_tier' => 'setSourceTier',
        'port_range' => 'setPortRange',
        'protocol' => 'setProtocol',
        'destination_type' => 'setDestinationType',
        'destination' => 'setDestination',
        'destination_group' => 'setDestinationGroup',
        'destination_tier' => 'setDestinationTier',
        'rule_type' => 'setRuleType',
        'policy' => 'setPolicy',
        'instance_type_id' => 'setInstanceTypeId'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'name' => 'getName',
        'direction' => 'getDirection',
        'source_type' => 'getSourceType',
        'source' => 'getSource',
        'source_group' => 'getSourceGroup',
        'source_tier' => 'getSourceTier',
        'port_range' => 'getPortRange',
        'protocol' => 'getProtocol',
        'destination_type' => 'getDestinationType',
        'destination' => 'getDestination',
        'destination_group' => 'getDestinationGroup',
        'destination_tier' => 'getDestinationTier',
        'rule_type' => 'getRuleType',
        'policy' => 'getPolicy',
        'instance_type_id' => 'getInstanceTypeId'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }

    const DIRECTION_INGRESS = 'ingress';
    const DIRECTION_EGRESS = 'egress';
    const SOURCE_TYPE_CIDR = 'cidr';
    const SOURCE_TYPE_GROUP = 'group';
    const SOURCE_TYPE_TIER = 'tier';
    const SOURCE_TYPE_ALL = 'all';
    const PROTOCOL_TCP = 'tcp';
    const PROTOCOL_UDP = 'udp';
    const PROTOCOL_ICMP = 'icmp';
    const DESTINATION_TYPE_CIDR = 'cidr';
    const DESTINATION_TYPE_GROUP = 'group';
    const DESTINATION_TYPE_TIER = 'tier';
    const DESTINATION_TYPE_INSTANCE = 'instance';
    const POLICY_ACCEPT = 'accept';
    const POLICY_DENY = 'deny';
    

    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getDirectionAllowableValues()
    {
        return [
            self::DIRECTION_INGRESS,
            self::DIRECTION_EGRESS,
        ];
    }
    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getSourceTypeAllowableValues()
    {
        return [
            self::SOURCE_TYPE_CIDR,
            self::SOURCE_TYPE_GROUP,
            self::SOURCE_TYPE_TIER,
            self::SOURCE_TYPE_ALL,
        ];
    }
    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getProtocolAllowableValues()
    {
        return [
            self::PROTOCOL_TCP,
            self::PROTOCOL_UDP,
            self::PROTOCOL_ICMP,
        ];
    }
    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getDestinationTypeAllowableValues()
    {
        return [
            self::DESTINATION_TYPE_CIDR,
            self::DESTINATION_TYPE_GROUP,
            self::DESTINATION_TYPE_TIER,
            self::DESTINATION_TYPE_INSTANCE,
        ];
    }
    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getPolicyAllowableValues()
    {
        return [
            self::POLICY_ACCEPT,
            self::POLICY_DENY,
        ];
    }
    

    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['name'] = $data['name'] ?? null;
        $this->container['direction'] = $data['direction'] ?? 'ingress';
        $this->container['source_type'] = $data['source_type'] ?? 'cidr';
        $this->container['source'] = $data['source'] ?? null;
        $this->container['source_group'] = $data['source_group'] ?? null;
        $this->container['source_tier'] = $data['source_tier'] ?? null;
        $this->container['port_range'] = $data['port_range'] ?? null;
        $this->container['protocol'] = $data['protocol'] ?? null;
        $this->container['destination_type'] = $data['destination_type'] ?? 'cidr';
        $this->container['destination'] = $data['destination'] ?? null;
        $this->container['destination_group'] = $data['destination_group'] ?? null;
        $this->container['destination_tier'] = $data['destination_tier'] ?? null;
        $this->container['rule_type'] = $data['rule_type'] ?? 'customRule';
        $this->container['policy'] = $data['policy'] ?? null;
        $this->container['instance_type_id'] = $data['instance_type_id'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        $allowedValues = $this->getDirectionAllowableValues();
        if (!is_null($this->container['direction']) && !in_array($this->container['direction'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'direction', must be one of '%s'",
                $this->container['direction'],
                implode("', '", $allowedValues)
            );
        }

        $allowedValues = $this->getSourceTypeAllowableValues();
        if (!is_null($this->container['source_type']) && !in_array($this->container['source_type'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'source_type', must be one of '%s'",
                $this->container['source_type'],
                implode("', '", $allowedValues)
            );
        }

        if ($this->container['protocol'] === null) {
            $invalidProperties[] = "'protocol' can't be null";
        }
        $allowedValues = $this->getProtocolAllowableValues();
        if (!is_null($this->container['protocol']) && !in_array($this->container['protocol'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'protocol', must be one of '%s'",
                $this->container['protocol'],
                implode("', '", $allowedValues)
            );
        }

        $allowedValues = $this->getDestinationTypeAllowableValues();
        if (!is_null($this->container['destination_type']) && !in_array($this->container['destination_type'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'destination_type', must be one of '%s'",
                $this->container['destination_type'],
                implode("', '", $allowedValues)
            );
        }

        if ($this->container['rule_type'] === null) {
            $invalidProperties[] = "'rule_type' can't be null";
        }
        $allowedValues = $this->getPolicyAllowableValues();
        if (!is_null($this->container['policy']) && !in_array($this->container['policy'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'policy', must be one of '%s'",
                $this->container['policy'],
                implode("', '", $allowedValues)
            );
        }

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets name
     *
     * @return string|null
     */
    public function getName()
    {
        return $this->container['name'];
    }

    /**
     * Sets name
     *
     * @param string|null $name A name for the rule
     *
     * @return self
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets direction
     *
     * @return string|null
     */
    public function getDirection()
    {
        return $this->container['direction'];
    }

    /**
     * Sets direction
     *
     * @param string|null $direction Either `ingress` or `egress`
     *
     * @return self
     */
    public function setDirection($direction)
    {
        $allowedValues = $this->getDirectionAllowableValues();
        if (!is_null($direction) && !in_array($direction, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'direction', must be one of '%s'",
                    $direction,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['direction'] = $direction;

        return $this;
    }

    /**
     * Gets source_type
     *
     * @return string|null
     */
    public function getSourceType()
    {
        return $this->container['source_type'];
    }

    /**
     * Sets source_type
     *
     * @param string|null $source_type Either `cidr`, `group`, `tier`, `all`.
     *
     * @return self
     */
    public function setSourceType($source_type)
    {
        $allowedValues = $this->getSourceTypeAllowableValues();
        if (!is_null($source_type) && !in_array($source_type, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'source_type', must be one of '%s'",
                    $source_type,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['source_type'] = $source_type;

        return $this;
    }

    /**
     * Gets source
     *
     * @return string|null
     */
    public function getSource()
    {
        return $this->container['source'];
    }

    /**
     * Sets source
     *
     * @param string|null $source CIDR representing the source IP(s) which should receive access. Required for `sourceType`=cidr
     *
     * @return self
     */
    public function setSource($source)
    {
        $this->container['source'] = $source;

        return $this;
    }

    /**
     * Gets source_group
     *
     * @return \OpenAPI\Client\Model\ApiSecurityGroupsIdRulesRuleSourceGroup|null
     */
    public function getSourceGroup()
    {
        return $this->container['source_group'];
    }

    /**
     * Sets source_group
     *
     * @param \OpenAPI\Client\Model\ApiSecurityGroupsIdRulesRuleSourceGroup|null $source_group source_group
     *
     * @return self
     */
    public function setSourceGroup($source_group)
    {
        $this->container['source_group'] = $source_group;

        return $this;
    }

    /**
     * Gets source_tier
     *
     * @return \OpenAPI\Client\Model\ApiSecurityGroupsIdRulesRuleSourceTier|null
     */
    public function getSourceTier()
    {
        return $this->container['source_tier'];
    }

    /**
     * Sets source_tier
     *
     * @param \OpenAPI\Client\Model\ApiSecurityGroupsIdRulesRuleSourceTier|null $source_tier source_tier
     *
     * @return self
     */
    public function setSourceTier($source_tier)
    {
        $this->container['source_tier'] = $source_tier;

        return $this;
    }

    /**
     * Gets port_range
     *
     * @return string|null
     */
    public function getPortRange()
    {
        return $this->container['port_range'];
    }

    /**
     * Sets port_range
     *
     * @param string|null $port_range Either a single value (i.e. 55) or a port range (i.e. 1-65535) for which to open access to the source. Required if customRule is true, otherwise, ignored.
     *
     * @return self
     */
    public function setPortRange($port_range)
    {
        $this->container['port_range'] = $port_range;

        return $this;
    }

    /**
     * Gets protocol
     *
     * @return string
     */
    public function getProtocol()
    {
        return $this->container['protocol'];
    }

    /**
     * Sets protocol
     *
     * @param string $protocol Either tcp, udp, icmp. Required if customRule is true, otherwise, ignored.
     *
     * @return self
     */
    public function setProtocol($protocol)
    {
        $allowedValues = $this->getProtocolAllowableValues();
        if (!in_array($protocol, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'protocol', must be one of '%s'",
                    $protocol,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['protocol'] = $protocol;

        return $this;
    }

    /**
     * Gets destination_type
     *
     * @return string|null
     */
    public function getDestinationType()
    {
        return $this->container['destination_type'];
    }

    /**
     * Sets destination_type
     *
     * @param string|null $destination_type Either cidr, group, tier, instance.
     *
     * @return self
     */
    public function setDestinationType($destination_type)
    {
        $allowedValues = $this->getDestinationTypeAllowableValues();
        if (!is_null($destination_type) && !in_array($destination_type, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'destination_type', must be one of '%s'",
                    $destination_type,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['destination_type'] = $destination_type;

        return $this;
    }

    /**
     * Gets destination
     *
     * @return string|null
     */
    public function getDestination()
    {
        return $this->container['destination'];
    }

    /**
     * Sets destination
     *
     * @param string|null $destination CIDR representing the destination IP(s) which should receive access. Required for `destinationType`=cidr.
     *
     * @return self
     */
    public function setDestination($destination)
    {
        $this->container['destination'] = $destination;

        return $this;
    }

    /**
     * Gets destination_group
     *
     * @return \OpenAPI\Client\Model\ApiSecurityGroupsIdRulesRuleDestinationGroup|null
     */
    public function getDestinationGroup()
    {
        return $this->container['destination_group'];
    }

    /**
     * Sets destination_group
     *
     * @param \OpenAPI\Client\Model\ApiSecurityGroupsIdRulesRuleDestinationGroup|null $destination_group destination_group
     *
     * @return self
     */
    public function setDestinationGroup($destination_group)
    {
        $this->container['destination_group'] = $destination_group;

        return $this;
    }

    /**
     * Gets destination_tier
     *
     * @return \OpenAPI\Client\Model\ApiSecurityGroupsIdRulesRuleDestinationTier|null
     */
    public function getDestinationTier()
    {
        return $this->container['destination_tier'];
    }

    /**
     * Sets destination_tier
     *
     * @param \OpenAPI\Client\Model\ApiSecurityGroupsIdRulesRuleDestinationTier|null $destination_tier destination_tier
     *
     * @return self
     */
    public function setDestinationTier($destination_tier)
    {
        $this->container['destination_tier'] = $destination_tier;

        return $this;
    }

    /**
     * Gets rule_type
     *
     * @return string
     */
    public function getRuleType()
    {
        return $this->container['rule_type'];
    }

    /**
     * Sets rule_type
     *
     * @param string $rule_type Either `customRule` or an `instance type` code.
     *
     * @return self
     */
    public function setRuleType($rule_type)
    {
        $this->container['rule_type'] = $rule_type;

        return $this;
    }

    /**
     * Gets policy
     *
     * @return string|null
     */
    public function getPolicy()
    {
        return $this->container['policy'];
    }

    /**
     * Sets policy
     *
     * @param string|null $policy Either `accept` or `deny`.
     *
     * @return self
     */
    public function setPolicy($policy)
    {
        $allowedValues = $this->getPolicyAllowableValues();
        if (!is_null($policy) && !in_array($policy, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'policy', must be one of '%s'",
                    $policy,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['policy'] = $policy;

        return $this;
    }

    /**
     * Gets instance_type_id
     *
     * @return int|null
     */
    public function getInstanceTypeId()
    {
        return $this->container['instance_type_id'];
    }

    /**
     * Sets instance_type_id
     *
     * @param int|null $instance_type_id The id of an Instance Type. If specified, the source CIDR will have access to all ports exposed by the particular instance in the cloud, app, or instance. Required if customRule is false, otherwise ignored.
     *
     * @return self
     */
    public function setInstanceTypeId($instance_type_id)
    {
        $this->container['instance_type_id'] = $instance_type_id;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}


