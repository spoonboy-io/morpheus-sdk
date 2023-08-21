<?php
/**
 * NetworkRouterType
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
 * NetworkRouterType Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class NetworkRouterType implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'networkRouterType';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'code' => 'string',
        'name' => 'string',
        'description' => 'string',
        'enabled' => 'bool',
        'creatable' => 'bool',
        'selectable' => 'bool',
        'has_firewall' => 'bool',
        'has_dhcp' => 'bool',
        'has_routing' => 'bool',
        'has_network_server' => 'bool',
        'option_types' => '\OpenAPI\Client\Model\OptionType[]',
        'rule_option_types' => '\OpenAPI\Client\Model\OptionType[]',
        'rule_group_option_types' => '\OpenAPI\Client\Model\OptionType[]',
        'nat_option_types' => '\OpenAPI\Client\Model\OptionType[]',
        'bgp_option_types' => '\OpenAPI\Client\Model\OptionType[]'
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
        'code' => null,
        'name' => null,
        'description' => null,
        'enabled' => null,
        'creatable' => null,
        'selectable' => null,
        'has_firewall' => null,
        'has_dhcp' => null,
        'has_routing' => null,
        'has_network_server' => null,
        'option_types' => null,
        'rule_option_types' => null,
        'rule_group_option_types' => null,
        'nat_option_types' => null,
        'bgp_option_types' => null
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
        'code' => 'code',
        'name' => 'name',
        'description' => 'description',
        'enabled' => 'enabled',
        'creatable' => 'creatable',
        'selectable' => 'selectable',
        'has_firewall' => 'hasFirewall',
        'has_dhcp' => 'hasDhcp',
        'has_routing' => 'hasRouting',
        'has_network_server' => 'hasNetworkServer',
        'option_types' => 'optionTypes',
        'rule_option_types' => 'ruleOptionTypes',
        'rule_group_option_types' => 'ruleGroupOptionTypes',
        'nat_option_types' => 'natOptionTypes',
        'bgp_option_types' => 'bgpOptionTypes'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'code' => 'setCode',
        'name' => 'setName',
        'description' => 'setDescription',
        'enabled' => 'setEnabled',
        'creatable' => 'setCreatable',
        'selectable' => 'setSelectable',
        'has_firewall' => 'setHasFirewall',
        'has_dhcp' => 'setHasDhcp',
        'has_routing' => 'setHasRouting',
        'has_network_server' => 'setHasNetworkServer',
        'option_types' => 'setOptionTypes',
        'rule_option_types' => 'setRuleOptionTypes',
        'rule_group_option_types' => 'setRuleGroupOptionTypes',
        'nat_option_types' => 'setNatOptionTypes',
        'bgp_option_types' => 'setBgpOptionTypes'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'code' => 'getCode',
        'name' => 'getName',
        'description' => 'getDescription',
        'enabled' => 'getEnabled',
        'creatable' => 'getCreatable',
        'selectable' => 'getSelectable',
        'has_firewall' => 'getHasFirewall',
        'has_dhcp' => 'getHasDhcp',
        'has_routing' => 'getHasRouting',
        'has_network_server' => 'getHasNetworkServer',
        'option_types' => 'getOptionTypes',
        'rule_option_types' => 'getRuleOptionTypes',
        'rule_group_option_types' => 'getRuleGroupOptionTypes',
        'nat_option_types' => 'getNatOptionTypes',
        'bgp_option_types' => 'getBgpOptionTypes'
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
        $this->container['code'] = $data['code'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['description'] = $data['description'] ?? null;
        $this->container['enabled'] = $data['enabled'] ?? null;
        $this->container['creatable'] = $data['creatable'] ?? null;
        $this->container['selectable'] = $data['selectable'] ?? null;
        $this->container['has_firewall'] = $data['has_firewall'] ?? null;
        $this->container['has_dhcp'] = $data['has_dhcp'] ?? null;
        $this->container['has_routing'] = $data['has_routing'] ?? null;
        $this->container['has_network_server'] = $data['has_network_server'] ?? null;
        $this->container['option_types'] = $data['option_types'] ?? null;
        $this->container['rule_option_types'] = $data['rule_option_types'] ?? null;
        $this->container['rule_group_option_types'] = $data['rule_group_option_types'] ?? null;
        $this->container['nat_option_types'] = $data['nat_option_types'] ?? null;
        $this->container['bgp_option_types'] = $data['bgp_option_types'] ?? null;
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
     * Gets code
     *
     * @return string|null
     */
    public function getCode()
    {
        return $this->container['code'];
    }

    /**
     * Sets code
     *
     * @param string|null $code code
     *
     * @return self
     */
    public function setCode($code)
    {
        $this->container['code'] = $code;

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
     * Gets description
     *
     * @return string|null
     */
    public function getDescription()
    {
        return $this->container['description'];
    }

    /**
     * Sets description
     *
     * @param string|null $description description
     *
     * @return self
     */
    public function setDescription($description)
    {
        $this->container['description'] = $description;

        return $this;
    }

    /**
     * Gets enabled
     *
     * @return bool|null
     */
    public function getEnabled()
    {
        return $this->container['enabled'];
    }

    /**
     * Sets enabled
     *
     * @param bool|null $enabled enabled
     *
     * @return self
     */
    public function setEnabled($enabled)
    {
        $this->container['enabled'] = $enabled;

        return $this;
    }

    /**
     * Gets creatable
     *
     * @return bool|null
     */
    public function getCreatable()
    {
        return $this->container['creatable'];
    }

    /**
     * Sets creatable
     *
     * @param bool|null $creatable creatable
     *
     * @return self
     */
    public function setCreatable($creatable)
    {
        $this->container['creatable'] = $creatable;

        return $this;
    }

    /**
     * Gets selectable
     *
     * @return bool|null
     */
    public function getSelectable()
    {
        return $this->container['selectable'];
    }

    /**
     * Sets selectable
     *
     * @param bool|null $selectable selectable
     *
     * @return self
     */
    public function setSelectable($selectable)
    {
        $this->container['selectable'] = $selectable;

        return $this;
    }

    /**
     * Gets has_firewall
     *
     * @return bool|null
     */
    public function getHasFirewall()
    {
        return $this->container['has_firewall'];
    }

    /**
     * Sets has_firewall
     *
     * @param bool|null $has_firewall has_firewall
     *
     * @return self
     */
    public function setHasFirewall($has_firewall)
    {
        $this->container['has_firewall'] = $has_firewall;

        return $this;
    }

    /**
     * Gets has_dhcp
     *
     * @return bool|null
     */
    public function getHasDhcp()
    {
        return $this->container['has_dhcp'];
    }

    /**
     * Sets has_dhcp
     *
     * @param bool|null $has_dhcp has_dhcp
     *
     * @return self
     */
    public function setHasDhcp($has_dhcp)
    {
        $this->container['has_dhcp'] = $has_dhcp;

        return $this;
    }

    /**
     * Gets has_routing
     *
     * @return bool|null
     */
    public function getHasRouting()
    {
        return $this->container['has_routing'];
    }

    /**
     * Sets has_routing
     *
     * @param bool|null $has_routing has_routing
     *
     * @return self
     */
    public function setHasRouting($has_routing)
    {
        $this->container['has_routing'] = $has_routing;

        return $this;
    }

    /**
     * Gets has_network_server
     *
     * @return bool|null
     */
    public function getHasNetworkServer()
    {
        return $this->container['has_network_server'];
    }

    /**
     * Sets has_network_server
     *
     * @param bool|null $has_network_server has_network_server
     *
     * @return self
     */
    public function setHasNetworkServer($has_network_server)
    {
        $this->container['has_network_server'] = $has_network_server;

        return $this;
    }

    /**
     * Gets option_types
     *
     * @return \OpenAPI\Client\Model\OptionType[]|null
     */
    public function getOptionTypes()
    {
        return $this->container['option_types'];
    }

    /**
     * Sets option_types
     *
     * @param \OpenAPI\Client\Model\OptionType[]|null $option_types option_types
     *
     * @return self
     */
    public function setOptionTypes($option_types)
    {
        $this->container['option_types'] = $option_types;

        return $this;
    }

    /**
     * Gets rule_option_types
     *
     * @return \OpenAPI\Client\Model\OptionType[]|null
     */
    public function getRuleOptionTypes()
    {
        return $this->container['rule_option_types'];
    }

    /**
     * Sets rule_option_types
     *
     * @param \OpenAPI\Client\Model\OptionType[]|null $rule_option_types rule_option_types
     *
     * @return self
     */
    public function setRuleOptionTypes($rule_option_types)
    {
        $this->container['rule_option_types'] = $rule_option_types;

        return $this;
    }

    /**
     * Gets rule_group_option_types
     *
     * @return \OpenAPI\Client\Model\OptionType[]|null
     */
    public function getRuleGroupOptionTypes()
    {
        return $this->container['rule_group_option_types'];
    }

    /**
     * Sets rule_group_option_types
     *
     * @param \OpenAPI\Client\Model\OptionType[]|null $rule_group_option_types rule_group_option_types
     *
     * @return self
     */
    public function setRuleGroupOptionTypes($rule_group_option_types)
    {
        $this->container['rule_group_option_types'] = $rule_group_option_types;

        return $this;
    }

    /**
     * Gets nat_option_types
     *
     * @return \OpenAPI\Client\Model\OptionType[]|null
     */
    public function getNatOptionTypes()
    {
        return $this->container['nat_option_types'];
    }

    /**
     * Sets nat_option_types
     *
     * @param \OpenAPI\Client\Model\OptionType[]|null $nat_option_types nat_option_types
     *
     * @return self
     */
    public function setNatOptionTypes($nat_option_types)
    {
        $this->container['nat_option_types'] = $nat_option_types;

        return $this;
    }

    /**
     * Gets bgp_option_types
     *
     * @return \OpenAPI\Client\Model\OptionType[]|null
     */
    public function getBgpOptionTypes()
    {
        return $this->container['bgp_option_types'];
    }

    /**
     * Sets bgp_option_types
     *
     * @param \OpenAPI\Client\Model\OptionType[]|null $bgp_option_types bgp_option_types
     *
     * @return self
     */
    public function setBgpOptionTypes($bgp_option_types)
    {
        $this->container['bgp_option_types'] = $bgp_option_types;

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


