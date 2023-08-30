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
from openapi_client.model.container_type_account import ContainerTypeAccount
from openapi_client.model.container_type_container_ports_inner import ContainerTypeContainerPortsInner
from openapi_client.model.container_type_provision_type import ContainerTypeProvisionType
globals()['ContainerTypeAccount'] = ContainerTypeAccount
globals()['ContainerTypeContainerPortsInner'] = ContainerTypeContainerPortsInner
globals()['ContainerTypeProvisionType'] = ContainerTypeProvisionType
from openapi_client.model.container_type import ContainerType


class TestContainerType(unittest.TestCase):
    """ContainerType unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testContainerType(self):
        """Test ContainerType"""
        # FIXME: construct object with mandatory attributes with example values
        # model = ContainerType()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
