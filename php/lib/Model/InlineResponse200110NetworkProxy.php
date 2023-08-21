<?php
/**
 * InlineResponse200110NetworkProxy
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
 * InlineResponse200110NetworkProxy Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class InlineResponse200110NetworkProxy implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'inline_response_200_110_networkProxy';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'name' => 'string',
        'proxy_host' => 'string',
        'proxy_port' => 'int',
        'proxy_user' => 'string',
        'proxy_password' => 'string',
        'proxy_domain' => 'string',
        'proxy_workstation' => 'string',
        'visibility' => 'string',
        'account' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance',
        'owner' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance'
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
        'proxy_host' => null,
        'proxy_port' => 'int64',
        'proxy_user' => null,
        'proxy_password' => null,
        'proxy_domain' => null,
        'proxy_workstation' => null,
        'visibility' => null,
        'account' => null,
        'owner' => null
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
        'proxy_host' => 'proxyHost',
        'proxy_port' => 'proxyPort',
        'proxy_user' => 'proxyUser',
        'proxy_password' => 'proxyPassword',
        'proxy_domain' => 'proxyDomain',
        'proxy_workstation' => 'proxyWorkstation',
        'visibility' => 'visibility',
        'account' => 'account',
        'owner' => 'owner'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'name' => 'setName',
        'proxy_host' => 'setProxyHost',
        'proxy_port' => 'setProxyPort',
        'proxy_user' => 'setProxyUser',
        'proxy_password' => 'setProxyPassword',
        'proxy_domain' => 'setProxyDomain',
        'proxy_workstation' => 'setProxyWorkstation',
        'visibility' => 'setVisibility',
        'account' => 'setAccount',
        'owner' => 'setOwner'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'name' => 'getName',
        'proxy_host' => 'getProxyHost',
        'proxy_port' => 'getProxyPort',
        'proxy_user' => 'getProxyUser',
        'proxy_password' => 'getProxyPassword',
        'proxy_domain' => 'getProxyDomain',
        'proxy_workstation' => 'getProxyWorkstation',
        'visibility' => 'getVisibility',
        'account' => 'getAccount',
        'owner' => 'getOwner'
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
        $this->container['proxy_host'] = $data['proxy_host'] ?? null;
        $this->container['proxy_port'] = $data['proxy_port'] ?? null;
        $this->container['proxy_user'] = $data['proxy_user'] ?? null;
        $this->container['proxy_password'] = $data['proxy_password'] ?? null;
        $this->container['proxy_domain'] = $data['proxy_domain'] ?? null;
        $this->container['proxy_workstation'] = $data['proxy_workstation'] ?? null;
        $this->container['visibility'] = $data['visibility'] ?? null;
        $this->container['account'] = $data['account'] ?? null;
        $this->container['owner'] = $data['owner'] ?? null;
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
     * Gets proxy_host
     *
     * @return string|null
     */
    public function getProxyHost()
    {
        return $this->container['proxy_host'];
    }

    /**
     * Sets proxy_host
     *
     * @param string|null $proxy_host proxy_host
     *
     * @return self
     */
    public function setProxyHost($proxy_host)
    {
        $this->container['proxy_host'] = $proxy_host;

        return $this;
    }

    /**
     * Gets proxy_port
     *
     * @return int|null
     */
    public function getProxyPort()
    {
        return $this->container['proxy_port'];
    }

    /**
     * Sets proxy_port
     *
     * @param int|null $proxy_port proxy_port
     *
     * @return self
     */
    public function setProxyPort($proxy_port)
    {
        $this->container['proxy_port'] = $proxy_port;

        return $this;
    }

    /**
     * Gets proxy_user
     *
     * @return string|null
     */
    public function getProxyUser()
    {
        return $this->container['proxy_user'];
    }

    /**
     * Sets proxy_user
     *
     * @param string|null $proxy_user proxy_user
     *
     * @return self
     */
    public function setProxyUser($proxy_user)
    {
        $this->container['proxy_user'] = $proxy_user;

        return $this;
    }

    /**
     * Gets proxy_password
     *
     * @return string|null
     */
    public function getProxyPassword()
    {
        return $this->container['proxy_password'];
    }

    /**
     * Sets proxy_password
     *
     * @param string|null $proxy_password proxy_password
     *
     * @return self
     */
    public function setProxyPassword($proxy_password)
    {
        $this->container['proxy_password'] = $proxy_password;

        return $this;
    }

    /**
     * Gets proxy_domain
     *
     * @return string|null
     */
    public function getProxyDomain()
    {
        return $this->container['proxy_domain'];
    }

    /**
     * Sets proxy_domain
     *
     * @param string|null $proxy_domain proxy_domain
     *
     * @return self
     */
    public function setProxyDomain($proxy_domain)
    {
        $this->container['proxy_domain'] = $proxy_domain;

        return $this;
    }

    /**
     * Gets proxy_workstation
     *
     * @return string|null
     */
    public function getProxyWorkstation()
    {
        return $this->container['proxy_workstation'];
    }

    /**
     * Sets proxy_workstation
     *
     * @param string|null $proxy_workstation proxy_workstation
     *
     * @return self
     */
    public function setProxyWorkstation($proxy_workstation)
    {
        $this->container['proxy_workstation'] = $proxy_workstation;

        return $this;
    }

    /**
     * Gets visibility
     *
     * @return string|null
     */
    public function getVisibility()
    {
        return $this->container['visibility'];
    }

    /**
     * Sets visibility
     *
     * @param string|null $visibility visibility
     *
     * @return self
     */
    public function setVisibility($visibility)
    {
        $this->container['visibility'] = $visibility;

        return $this;
    }

    /**
     * Gets account
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null
     */
    public function getAccount()
    {
        return $this->container['account'];
    }

    /**
     * Sets account
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null $account account
     *
     * @return self
     */
    public function setAccount($account)
    {
        $this->container['account'] = $account;

        return $this;
    }

    /**
     * Gets owner
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null
     */
    public function getOwner()
    {
        return $this->container['owner'];
    }

    /**
     * Sets owner
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null $owner owner
     *
     * @return self
     */
    public function setOwner($owner)
    {
        $this->container['owner'] = $owner;

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


