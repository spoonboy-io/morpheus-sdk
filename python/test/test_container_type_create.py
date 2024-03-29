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
from openapi_client.model.cluster_layout_create_environment_variables_inner import ClusterLayoutCreateEnvironmentVariablesInner
from openapi_client.model.container_type_create_container_ports_inner import ContainerTypeCreateContainerPortsInner
globals()['ClusterLayoutCreateEnvironmentVariablesInner'] = ClusterLayoutCreateEnvironmentVariablesInner
globals()['ContainerTypeCreateContainerPortsInner'] = ContainerTypeCreateContainerPortsInner
from openapi_client.model.container_type_create import ContainerTypeCreate


class TestContainerTypeCreate(unittest.TestCase):
    """ContainerTypeCreate unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testContainerTypeCreate(self):
        """Test ContainerTypeCreate"""
        # FIXME: construct object with mandatory attributes with example values
        # model = ContainerTypeCreate()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
