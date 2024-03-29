<?php
/**
 * SecurityGroupRules
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
 * SecurityGroupRules Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class SecurityGroupRules implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'securityGroup_rules';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'name' => 'string',
        'rule_type' => 'string',
        'custom_rule' => 'bool',
        'instance_type_id' => 'string',
        'direction' => 'string',
        'policy' => 'string',
        'source_type' => 'string',
        'source' => 'string',
        'source_group' => 'string',
        'source_tier' => 'string',
        'port_range' => 'string',
        'protocol' => 'string',
        'destination_type' => 'string',
        'destination' => 'string',
        'destination_group' => 'string',
        'destination_tier' => 'string',
        'external_id' => 'string',
        'enabled' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'id' => 'int64',
        'name' => null,
        'rule_type' => null,
        'custom_rule' => null,
        'instance_type_id' => null,
        'direction' => null,
        'policy' => null,
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
        'external_id' => null,
        'enabled' => null
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
        'id' => 'id',
        'name' => 'name',
        'rule_type' => 'ruleType',
        'custom_rule' => 'customRule',
        'instance_type_id' => 'instanceTypeId',
        'direction' => 'direction',
        'policy' => 'policy',
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
        'external_id' => 'externalId',
        'enabled' => 'enabled'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'name' => 'setName',
        'rule_type' => 'setRuleType',
        'custom_rule' => 'setCustomRule',
        'instance_type_id' => 'setInstanceTypeId',
        'direction' => 'setDirection',
        'policy' => 'setPolicy',
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
        'external_id' => 'setExternalId',
        'enabled' => 'setEnabled'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'name' => 'getName',
        'rule_type' => 'getRuleType',
        'custom_rule' => 'getCustomRule',
        'instance_type_id' => 'getInstanceTypeId',
        'direction' => 'getDirection',
        'policy' => 'getPolicy',
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
        'external_id' => 'getExternalId',
        'enabled' => 'getEnabled'
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
        $this->container['id'] = $data['id'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['rule_type'] = $data['rule_type'] ?? null;
        $this->container['custom_rule'] = $data['custom_rule'] ?? null;
        $this->container['instance_type_id'] = $data['instance_type_id'] ?? null;
        $this->container['direction'] = $data['direction'] ?? null;
        $this->container['policy'] = $data['policy'] ?? null;
        $this->container['source_type'] = $data['source_type'] ?? null;
        $this->container['source'] = $data['source'] ?? null;
        $this->container['source_group'] = $data['source_group'] ?? null;
        $this->container['source_tier'] = $data['source_tier'] ?? null;
        $this->container['port_range'] = $data['port_range'] ?? null;
        $this->container['protocol'] = $data['protocol'] ?? null;
        $this->container['destination_type'] = $data['destination_type'] ?? null;
        $this->container['destination'] = $data['destination'] ?? null;
        $this->container['destination_group'] = $data['destination_group'] ?? null;
        $this->container['destination_tier'] = $data['destination_tier'] ?? null;
        $this->container['external_id'] = $data['external_id'] ?? null;
        $this->container['enabled'] = $data['enabled'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

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
     * Gets id
     *
     * @return int|null
     */
    public function getId()
    {
        return $this->container['id'];
    }

    /**
     * Sets id
     *
     * @param int|null $id id
     *
     * @return self
     */
    public function setId($id)
    {
        $this->container['id'] = $id;

        return $this;
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
     * @param string|null $name name
     *
     * @return self
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets rule_type
     *
     * @return string|null
     */
    public function getRuleType()
    {
        return $this->container['rule_type'];
    }

    /**
     * Sets rule_type
     *
     * @param string|null $rule_type rule_type
     *
     * @return self
     */
    public function setRuleType($rule_type)
    {
        $this->container['rule_type'] = $rule_type;

        return $this;
    }

    /**
     * Gets custom_rule
     *
     * @return bool|null
     */
    public function getCustomRule()
    {
        return $this->container['custom_rule'];
    }

    /**
     * Sets custom_rule
     *
     * @param bool|null $custom_rule custom_rule
     *
     * @return self
     */
    public function setCustomRule($custom_rule)
    {
        $this->container['custom_rule'] = $custom_rule;

        return $this;
    }

    /**
     * Gets instance_type_id
     *
     * @return string|null
     */
    public function getInstanceTypeId()
    {
        return $this->container['instance_type_id'];
    }

    /**
     * Sets instance_type_id
     *
     * @param string|null $instance_type_id instance_type_id
     *
     * @return self
     */
    public function setInstanceTypeId($instance_type_id)
    {
        $this->container['instance_type_id'] = $instance_type_id;

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
     * @param string|null $direction direction
     *
     * @return self
     */
    public function setDirection($direction)
    {
        $this->container['direction'] = $direction;

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
     * @param string|null $policy policy
     *
     * @return self
     */
    public function setPolicy($policy)
    {
        $this->container['policy'] = $policy;

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
     * @param string|null $source_type source_type
     *
     * @return self
     */
    public function setSourceType($source_type)
    {
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
     * @param string|null $source source
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
     * @return string|null
     */
    public function getSourceGroup()
    {
        return $this->container['source_group'];
    }

    /**
     * Sets source_group
     *
     * @param string|null $source_group source_group
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
     * @return string|null
     */
    public function getSourceTier()
    {
        return $this->container['source_tier'];
    }

    /**
     * Sets source_tier
     *
     * @param string|null $source_tier source_tier
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
     * @param string|null $port_range port_range
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
     * @return string|null
     */
    public function getProtocol()
    {
        return $this->container['protocol'];
    }

    /**
     * Sets protocol
     *
     * @param string|null $protocol protocol
     *
     * @return self
     */
    public function setProtocol($protocol)
    {
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
     * @param string|null $destination_type destination_type
     *
     * @return self
     */
    public function setDestinationType($destination_type)
    {
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
     * @param string|null $destination destination
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
     * @return string|null
     */
    public function getDestinationGroup()
    {
        return $this->container['destination_group'];
    }

    /**
     * Sets destination_group
     *
     * @param string|null $destination_group destination_group
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
     * @return string|null
     */
    public function getDestinationTier()
    {
        return $this->container['destination_tier'];
    }

    /**
     * Sets destination_tier
     *
     * @param string|null $destination_tier destination_tier
     *
     * @return self
     */
    public function setDestinationTier($destination_tier)
    {
        $this->container['destination_tier'] = $destination_tier;

        return $this;
    }

    /**
     * Gets external_id
     *
     * @return string|null
     */
    public function getExternalId()
    {
        return $this->container['external_id'];
    }

    /**
     * Sets external_id
     *
     * @param string|null $external_id external_id
     *
     * @return self
     */
    public function setExternalId($external_id)
    {
        $this->container['external_id'] = $external_id;

        return $this;
    }

    /**
     * Gets enabled
     *
     * @return string|null
     */
    public function getEnabled()
    {
        return $this->container['enabled'];
    }

    /**
     * Sets enabled
     *
     * @param string|null $enabled enabled
     *
     * @return self
     */
    public function setEnabled($enabled)
    {
        $this->container['enabled'] = $enabled;

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


