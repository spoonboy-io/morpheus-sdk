<?php
/**
 * CredentialOauth2ConfigConfig
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
 * CredentialOauth2ConfigConfig Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class CredentialOauth2ConfigConfig implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'credentialOauth2Config_config';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'grant_type' => 'string',
        'access_token_url' => 'string',
        'client_id' => 'string',
        'client_secret' => 'string',
        'scope' => 'string',
        'client_auth' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'grant_type' => null,
        'access_token_url' => null,
        'client_id' => null,
        'client_secret' => null,
        'scope' => null,
        'client_auth' => null
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
        'grant_type' => 'grantType',
        'access_token_url' => 'accessTokenUrl',
        'client_id' => 'clientId',
        'client_secret' => 'clientSecret',
        'scope' => 'scope',
        'client_auth' => 'clientAuth'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'grant_type' => 'setGrantType',
        'access_token_url' => 'setAccessTokenUrl',
        'client_id' => 'setClientId',
        'client_secret' => 'setClientSecret',
        'scope' => 'setScope',
        'client_auth' => 'setClientAuth'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'grant_type' => 'getGrantType',
        'access_token_url' => 'getAccessTokenUrl',
        'client_id' => 'getClientId',
        'client_secret' => 'getClientSecret',
        'scope' => 'getScope',
        'client_auth' => 'getClientAuth'
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

    const GRANT_TYPE_PASSWORD = 'password';
    const GRANT_TYPE_CLIENT_CREDENTIALS = 'client_credentials';
    const CLIENT_AUTH_BODY = 'body';
    const CLIENT_AUTH_BASIC_AUTH = 'basic-auth';
    

    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getGrantTypeAllowableValues()
    {
        return [
            self::GRANT_TYPE_PASSWORD,
            self::GRANT_TYPE_CLIENT_CREDENTIALS,
        ];
    }
    
    /**
     * Gets allowable values of the enum
     *
     * @return string[]
     */
    public function getClientAuthAllowableValues()
    {
        return [
            self::CLIENT_AUTH_BODY,
            self::CLIENT_AUTH_BASIC_AUTH,
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
        $this->container['grant_type'] = $data['grant_type'] ?? null;
        $this->container['access_token_url'] = $data['access_token_url'] ?? null;
        $this->container['client_id'] = $data['client_id'] ?? null;
        $this->container['client_secret'] = $data['client_secret'] ?? null;
        $this->container['scope'] = $data['scope'] ?? null;
        $this->container['client_auth'] = $data['client_auth'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['grant_type'] === null) {
            $invalidProperties[] = "'grant_type' can't be null";
        }
        $allowedValues = $this->getGrantTypeAllowableValues();
        if (!is_null($this->container['grant_type']) && !in_array($this->container['grant_type'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'grant_type', must be one of '%s'",
                $this->container['grant_type'],
                implode("', '", $allowedValues)
            );
        }

        if ($this->container['access_token_url'] === null) {
            $invalidProperties[] = "'access_token_url' can't be null";
        }
        if ($this->container['client_auth'] === null) {
            $invalidProperties[] = "'client_auth' can't be null";
        }
        $allowedValues = $this->getClientAuthAllowableValues();
        if (!is_null($this->container['client_auth']) && !in_array($this->container['client_auth'], $allowedValues, true)) {
            $invalidProperties[] = sprintf(
                "invalid value '%s' for 'client_auth', must be one of '%s'",
                $this->container['client_auth'],
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
     * Gets grant_type
     *
     * @return string
     */
    public function getGrantType()
    {
        return $this->container['grant_type'];
    }

    /**
     * Sets grant_type
     *
     * @param string $grant_type OAuth 2.0 grant type
     *
     * @return self
     */
    public function setGrantType($grant_type)
    {
        $allowedValues = $this->getGrantTypeAllowableValues();
        if (!in_array($grant_type, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'grant_type', must be one of '%s'",
                    $grant_type,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['grant_type'] = $grant_type;

        return $this;
    }

    /**
     * Gets access_token_url
     *
     * @return string
     */
    public function getAccessTokenUrl()
    {
        return $this->container['access_token_url'];
    }

    /**
     * Sets access_token_url
     *
     * @param string $access_token_url Token endpoint
     *
     * @return self
     */
    public function setAccessTokenUrl($access_token_url)
    {
        $this->container['access_token_url'] = $access_token_url;

        return $this;
    }

    /**
     * Gets client_id
     *
     * @return string|null
     */
    public function getClientId()
    {
        return $this->container['client_id'];
    }

    /**
     * Sets client_id
     *
     * @param string|null $client_id Client ID
     *
     * @return self
     */
    public function setClientId($client_id)
    {
        $this->container['client_id'] = $client_id;

        return $this;
    }

    /**
     * Gets client_secret
     *
     * @return string|null
     */
    public function getClientSecret()
    {
        return $this->container['client_secret'];
    }

    /**
     * Sets client_secret
     *
     * @param string|null $client_secret Client Secret
     *
     * @return self
     */
    public function setClientSecret($client_secret)
    {
        $this->container['client_secret'] = $client_secret;

        return $this;
    }

    /**
     * Gets scope
     *
     * @return string|null
     */
    public function getScope()
    {
        return $this->container['scope'];
    }

    /**
     * Sets scope
     *
     * @param string|null $scope Scope
     *
     * @return self
     */
    public function setScope($scope)
    {
        $this->container['scope'] = $scope;

        return $this;
    }

    /**
     * Gets client_auth
     *
     * @return string
     */
    public function getClientAuth()
    {
        return $this->container['client_auth'];
    }

    /**
     * Sets client_auth
     *
     * @param string $client_auth Client Authentication Method
     *
     * @return self
     */
    public function setClientAuth($client_auth)
    {
        $allowedValues = $this->getClientAuthAllowableValues();
        if (!in_array($client_auth, $allowedValues, true)) {
            throw new \InvalidArgumentException(
                sprintf(
                    "Invalid value '%s' for 'client_auth', must be one of '%s'",
                    $client_auth,
                    implode("', '", $allowedValues)
                )
            );
        }
        $this->container['client_auth'] = $client_auth;

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


