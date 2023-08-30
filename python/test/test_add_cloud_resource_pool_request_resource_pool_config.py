"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import openapi_client
from openapi_client.model.aws_resource_pool_config import AwsResourcePoolConfig
from openapi_client.model.cloud_foundry_resource_pool_config import CloudFoundryResourcePoolConfig
globals()['AwsResourcePoolConfig'] = AwsResourcePoolConfig
globals()['CloudFoundryResourcePoolConfig'] = CloudFoundryResourcePoolConfig
from openapi_client.model.add_cloud_resource_pool_request_resource_pool_config import AddCloudResourcePoolRequestResourcePoolConfig


class TestAddCloudResourcePoolRequestResourcePoolConfig(unittest.TestCase):
    """AddCloudResourcePoolRequestResourcePoolConfig unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testAddCloudResourcePoolRequestResourcePoolConfig(self):
        """Test AddCloudResourcePoolRequestResourcePoolConfig"""
        # FIXME: construct object with mandatory attributes with example values
        # model = AddCloudResourcePoolRequestResourcePoolConfig()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
