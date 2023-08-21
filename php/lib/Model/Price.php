<?php
/**
 * Price
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
 * Price Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class Price implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'price';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'name' => 'string',
        'code' => 'string',
        'active' => 'bool',
        'price_type' => 'string',
        'price_unit' => 'string',
        'additional_price_unit' => 'string',
        'price' => 'float',
        'custom_price' => 'float',
        'markup_type' => 'string',
        'markup' => 'float',
        'markup_percent' => 'float',
        'cost' => 'float',
        'currency' => 'string',
        'incur_charges' => 'string',
        'platform' => 'string',
        'software' => 'string',
        'volume_type' => '\OpenAPI\Client\Model\PriceSetVolumeType',
        'datastore' => '\OpenAPI\Client\Model\InlineResponse20082LoadBalancerInstanceSslCert',
        'cross_cloud_apply' => 'bool',
        'account' => 'string'
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
        'code' => null,
        'active' => null,
        'price_type' => null,
        'price_unit' => null,
        'additional_price_unit' => null,
        'price' => null,
        'custom_price' => null,
        'markup_type' => null,
        'markup' => null,
        'markup_percent' => null,
        'cost' => null,
        'currency' => null,
        'incur_charges' => null,
        'platform' => null,
        'software' => null,
        'volume_type' => null,
        'datastore' => null,
        'cross_cloud_apply' => null,
        'account' => null
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
        'code' => 'code',
        'active' => 'active',
        'price_type' => 'priceType',
        'price_unit' => 'priceUnit',
        'additional_price_unit' => 'additionalPriceUnit',
        'price' => 'price',
        'custom_price' => 'customPrice',
        'markup_type' => 'markupType',
        'markup' => 'markup',
        'markup_percent' => 'markupPercent',
        'cost' => 'cost',
        'currency' => 'currency',
        'incur_charges' => 'incurCharges',
        'platform' => 'platform',
        'software' => 'software',
        'volume_type' => 'volumeType',
        'datastore' => 'datastore',
        'cross_cloud_apply' => 'crossCloudApply',
        'account' => 'account'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'name' => 'setName',
        'code' => 'setCode',
        'active' => 'setActive',
        'price_type' => 'setPriceType',
        'price_unit' => 'setPriceUnit',
        'additional_price_unit' => 'setAdditionalPriceUnit',
        'price' => 'setPrice',
        'custom_price' => 'setCustomPrice',
        'markup_type' => 'setMarkupType',
        'markup' => 'setMarkup',
        'markup_percent' => 'setMarkupPercent',
        'cost' => 'setCost',
        'currency' => 'setCurrency',
        'incur_charges' => 'setIncurCharges',
        'platform' => 'setPlatform',
        'software' => 'setSoftware',
        'volume_type' => 'setVolumeType',
        'datastore' => 'setDatastore',
        'cross_cloud_apply' => 'setCrossCloudApply',
        'account' => 'setAccount'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'name' => 'getName',
        'code' => 'getCode',
        'active' => 'getActive',
        'price_type' => 'getPriceType',
        'price_unit' => 'getPriceUnit',
        'additional_price_unit' => 'getAdditionalPriceUnit',
        'price' => 'getPrice',
        'custom_price' => 'getCustomPrice',
        'markup_type' => 'getMarkupType',
        'markup' => 'getMarkup',
        'markup_percent' => 'getMarkupPercent',
        'cost' => 'getCost',
        'currency' => 'getCurrency',
        'incur_charges' => 'getIncurCharges',
        'platform' => 'getPlatform',
        'software' => 'getSoftware',
        'volume_type' => 'getVolumeType',
        'datastore' => 'getDatastore',
        'cross_cloud_apply' => 'getCrossCloudApply',
        'account' => 'getAccount'
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
        $this->container['code'] = $data['code'] ?? null;
        $this->container['active'] = $data['active'] ?? null;
        $this->container['price_type'] = $data['price_type'] ?? null;
        $this->container['price_unit'] = $data['price_unit'] ?? null;
        $this->container['additional_price_unit'] = $data['additional_price_unit'] ?? null;
        $this->container['price'] = $data['price'] ?? null;
        $this->container['custom_price'] = $data['custom_price'] ?? null;
        $this->container['markup_type'] = $data['markup_type'] ?? null;
        $this->container['markup'] = $data['markup'] ?? null;
        $this->container['markup_percent'] = $data['markup_percent'] ?? null;
        $this->container['cost'] = $data['cost'] ?? null;
        $this->container['currency'] = $data['currency'] ?? null;
        $this->container['incur_charges'] = $data['incur_charges'] ?? null;
        $this->container['platform'] = $data['platform'] ?? null;
        $this->container['software'] = $data['software'] ?? null;
        $this->container['volume_type'] = $data['volume_type'] ?? null;
        $this->container['datastore'] = $data['datastore'] ?? null;
        $this->container['cross_cloud_apply'] = $data['cross_cloud_apply'] ?? null;
        $this->container['account'] = $data['account'] ?? null;
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
     * Gets active
     *
     * @return bool|null
     */
    public function getActive()
    {
        return $this->container['active'];
    }

    /**
     * Sets active
     *
     * @param bool|null $active active
     *
     * @return self
     */
    public function setActive($active)
    {
        $this->container['active'] = $active;

        return $this;
    }

    /**
     * Gets price_type
     *
     * @return string|null
     */
    public function getPriceType()
    {
        return $this->container['price_type'];
    }

    /**
     * Sets price_type
     *
     * @param string|null $price_type price_type
     *
     * @return self
     */
    public function setPriceType($price_type)
    {
        $this->container['price_type'] = $price_type;

        return $this;
    }

    /**
     * Gets price_unit
     *
     * @return string|null
     */
    public function getPriceUnit()
    {
        return $this->container['price_unit'];
    }

    /**
     * Sets price_unit
     *
     * @param string|null $price_unit price_unit
     *
     * @return self
     */
    public function setPriceUnit($price_unit)
    {
        $this->container['price_unit'] = $price_unit;

        return $this;
    }

    /**
     * Gets additional_price_unit
     *
     * @return string|null
     */
    public function getAdditionalPriceUnit()
    {
        return $this->container['additional_price_unit'];
    }

    /**
     * Sets additional_price_unit
     *
     * @param string|null $additional_price_unit additional_price_unit
     *
     * @return self
     */
    public function setAdditionalPriceUnit($additional_price_unit)
    {
        $this->container['additional_price_unit'] = $additional_price_unit;

        return $this;
    }

    /**
     * Gets price
     *
     * @return float|null
     */
    public function getPrice()
    {
        return $this->container['price'];
    }

    /**
     * Sets price
     *
     * @param float|null $price price
     *
     * @return self
     */
    public function setPrice($price)
    {
        $this->container['price'] = $price;

        return $this;
    }

    /**
     * Gets custom_price
     *
     * @return float|null
     */
    public function getCustomPrice()
    {
        return $this->container['custom_price'];
    }

    /**
     * Sets custom_price
     *
     * @param float|null $custom_price custom_price
     *
     * @return self
     */
    public function setCustomPrice($custom_price)
    {
        $this->container['custom_price'] = $custom_price;

        return $this;
    }

    /**
     * Gets markup_type
     *
     * @return string|null
     */
    public function getMarkupType()
    {
        return $this->container['markup_type'];
    }

    /**
     * Sets markup_type
     *
     * @param string|null $markup_type markup_type
     *
     * @return self
     */
    public function setMarkupType($markup_type)
    {
        $this->container['markup_type'] = $markup_type;

        return $this;
    }

    /**
     * Gets markup
     *
     * @return float|null
     */
    public function getMarkup()
    {
        return $this->container['markup'];
    }

    /**
     * Sets markup
     *
     * @param float|null $markup markup
     *
     * @return self
     */
    public function setMarkup($markup)
    {
        $this->container['markup'] = $markup;

        return $this;
    }

    /**
     * Gets markup_percent
     *
     * @return float|null
     */
    public function getMarkupPercent()
    {
        return $this->container['markup_percent'];
    }

    /**
     * Sets markup_percent
     *
     * @param float|null $markup_percent markup_percent
     *
     * @return self
     */
    public function setMarkupPercent($markup_percent)
    {
        $this->container['markup_percent'] = $markup_percent;

        return $this;
    }

    /**
     * Gets cost
     *
     * @return float|null
     */
    public function getCost()
    {
        return $this->container['cost'];
    }

    /**
     * Sets cost
     *
     * @param float|null $cost cost
     *
     * @return self
     */
    public function setCost($cost)
    {
        $this->container['cost'] = $cost;

        return $this;
    }

    /**
     * Gets currency
     *
     * @return string|null
     */
    public function getCurrency()
    {
        return $this->container['currency'];
    }

    /**
     * Sets currency
     *
     * @param string|null $currency currency
     *
     * @return self
     */
    public function setCurrency($currency)
    {
        $this->container['currency'] = $currency;

        return $this;
    }

    /**
     * Gets incur_charges
     *
     * @return string|null
     */
    public function getIncurCharges()
    {
        return $this->container['incur_charges'];
    }

    /**
     * Sets incur_charges
     *
     * @param string|null $incur_charges incur_charges
     *
     * @return self
     */
    public function setIncurCharges($incur_charges)
    {
        $this->container['incur_charges'] = $incur_charges;

        return $this;
    }

    /**
     * Gets platform
     *
     * @return string|null
     */
    public function getPlatform()
    {
        return $this->container['platform'];
    }

    /**
     * Sets platform
     *
     * @param string|null $platform platform
     *
     * @return self
     */
    public function setPlatform($platform)
    {
        $this->container['platform'] = $platform;

        return $this;
    }

    /**
     * Gets software
     *
     * @return string|null
     */
    public function getSoftware()
    {
        return $this->container['software'];
    }

    /**
     * Sets software
     *
     * @param string|null $software software
     *
     * @return self
     */
    public function setSoftware($software)
    {
        $this->container['software'] = $software;

        return $this;
    }

    /**
     * Gets volume_type
     *
     * @return \OpenAPI\Client\Model\PriceSetVolumeType|null
     */
    public function getVolumeType()
    {
        return $this->container['volume_type'];
    }

    /**
     * Sets volume_type
     *
     * @param \OpenAPI\Client\Model\PriceSetVolumeType|null $volume_type volume_type
     *
     * @return self
     */
    public function setVolumeType($volume_type)
    {
        $this->container['volume_type'] = $volume_type;

        return $this;
    }

    /**
     * Gets datastore
     *
     * @return \OpenAPI\Client\Model\InlineResponse20082LoadBalancerInstanceSslCert|null
     */
    public function getDatastore()
    {
        return $this->container['datastore'];
    }

    /**
     * Sets datastore
     *
     * @param \OpenAPI\Client\Model\InlineResponse20082LoadBalancerInstanceSslCert|null $datastore datastore
     *
     * @return self
     */
    public function setDatastore($datastore)
    {
        $this->container['datastore'] = $datastore;

        return $this;
    }

    /**
     * Gets cross_cloud_apply
     *
     * @return bool|null
     */
    public function getCrossCloudApply()
    {
        return $this->container['cross_cloud_apply'];
    }

    /**
     * Sets cross_cloud_apply
     *
     * @param bool|null $cross_cloud_apply cross_cloud_apply
     *
     * @return self
     */
    public function setCrossCloudApply($cross_cloud_apply)
    {
        $this->container['cross_cloud_apply'] = $cross_cloud_apply;

        return $this;
    }

    /**
     * Gets account
     *
     * @return string|null
     */
    public function getAccount()
    {
        return $this->container['account'];
    }

    /**
     * Sets account
     *
     * @param string|null $account account
     *
     * @return self
     */
    public function setAccount($account)
    {
        $this->container['account'] = $account;

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

