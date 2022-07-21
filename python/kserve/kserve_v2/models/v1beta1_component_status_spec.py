# Copyright 2022 The KServe Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8

"""
    KServe

    Python SDK for KServe  # noqa: E501

    The version of the OpenAPI document: v0.1
    Generated by: https://openapi-generator.tech
"""


import pprint
import re  # noqa: F401

import six

from kserve.configuration import Configuration


class V1beta1ComponentStatusSpec(object):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    """
    Attributes:
      openapi_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    openapi_types = {
        'address': 'KnativeAddressable',
        'grpc_url': 'KnativeURL',
        'latest_created_revision': 'str',
        'latest_ready_revision': 'str',
        'latest_rolledout_revision': 'str',
        'previous_rolledout_revision': 'str',
        'rest_url': 'KnativeURL',
        'traffic': 'list[KnativeDevServingPkgApisServingV1TrafficTarget]',
        'url': 'KnativeURL'
    }

    attribute_map = {
        'address': 'address',
        'grpc_url': 'grpcUrl',
        'latest_created_revision': 'latestCreatedRevision',
        'latest_ready_revision': 'latestReadyRevision',
        'latest_rolledout_revision': 'latestRolledoutRevision',
        'previous_rolledout_revision': 'previousRolledoutRevision',
        'rest_url': 'restUrl',
        'traffic': 'traffic',
        'url': 'url'
    }

    def __init__(self, address=None, grpc_url=None, latest_created_revision=None, latest_ready_revision=None, latest_rolledout_revision=None, previous_rolledout_revision=None, rest_url=None, traffic=None, url=None, local_vars_configuration=None):  # noqa: E501
        """V1beta1ComponentStatusSpec - a model defined in OpenAPI"""  # noqa: E501
        if local_vars_configuration is None:
            local_vars_configuration = Configuration()
        self.local_vars_configuration = local_vars_configuration

        self._address = None
        self._grpc_url = None
        self._latest_created_revision = None
        self._latest_ready_revision = None
        self._latest_rolledout_revision = None
        self._previous_rolledout_revision = None
        self._rest_url = None
        self._traffic = None
        self._url = None
        self.discriminator = None

        if address is not None:
            self.address = address
        if grpc_url is not None:
            self.grpc_url = grpc_url
        if latest_created_revision is not None:
            self.latest_created_revision = latest_created_revision
        if latest_ready_revision is not None:
            self.latest_ready_revision = latest_ready_revision
        if latest_rolledout_revision is not None:
            self.latest_rolledout_revision = latest_rolledout_revision
        if previous_rolledout_revision is not None:
            self.previous_rolledout_revision = previous_rolledout_revision
        if rest_url is not None:
            self.rest_url = rest_url
        if traffic is not None:
            self.traffic = traffic
        if url is not None:
            self.url = url

    @property
    def address(self):
        """Gets the address of this V1beta1ComponentStatusSpec.  # noqa: E501


        :return: The address of this V1beta1ComponentStatusSpec.  # noqa: E501
        :rtype: KnativeAddressable
        """
        return self._address

    @address.setter
    def address(self, address):
        """Sets the address of this V1beta1ComponentStatusSpec.


        :param address: The address of this V1beta1ComponentStatusSpec.  # noqa: E501
        :type: KnativeAddressable
        """

        self._address = address

    @property
    def grpc_url(self):
        """Gets the grpc_url of this V1beta1ComponentStatusSpec.  # noqa: E501


        :return: The grpc_url of this V1beta1ComponentStatusSpec.  # noqa: E501
        :rtype: KnativeURL
        """
        return self._grpc_url

    @grpc_url.setter
    def grpc_url(self, grpc_url):
        """Sets the grpc_url of this V1beta1ComponentStatusSpec.


        :param grpc_url: The grpc_url of this V1beta1ComponentStatusSpec.  # noqa: E501
        :type: KnativeURL
        """

        self._grpc_url = grpc_url

    @property
    def latest_created_revision(self):
        """Gets the latest_created_revision of this V1beta1ComponentStatusSpec.  # noqa: E501

        Latest revision name that is created  # noqa: E501

        :return: The latest_created_revision of this V1beta1ComponentStatusSpec.  # noqa: E501
        :rtype: str
        """
        return self._latest_created_revision

    @latest_created_revision.setter
    def latest_created_revision(self, latest_created_revision):
        """Sets the latest_created_revision of this V1beta1ComponentStatusSpec.

        Latest revision name that is created  # noqa: E501

        :param latest_created_revision: The latest_created_revision of this V1beta1ComponentStatusSpec.  # noqa: E501
        :type: str
        """

        self._latest_created_revision = latest_created_revision

    @property
    def latest_ready_revision(self):
        """Gets the latest_ready_revision of this V1beta1ComponentStatusSpec.  # noqa: E501

        Latest revision name that is in ready state  # noqa: E501

        :return: The latest_ready_revision of this V1beta1ComponentStatusSpec.  # noqa: E501
        :rtype: str
        """
        return self._latest_ready_revision

    @latest_ready_revision.setter
    def latest_ready_revision(self, latest_ready_revision):
        """Sets the latest_ready_revision of this V1beta1ComponentStatusSpec.

        Latest revision name that is in ready state  # noqa: E501

        :param latest_ready_revision: The latest_ready_revision of this V1beta1ComponentStatusSpec.  # noqa: E501
        :type: str
        """

        self._latest_ready_revision = latest_ready_revision

    @property
    def latest_rolledout_revision(self):
        """Gets the latest_rolledout_revision of this V1beta1ComponentStatusSpec.  # noqa: E501

        Latest revision name that is rolled out with 100 percent traffic  # noqa: E501

        :return: The latest_rolledout_revision of this V1beta1ComponentStatusSpec.  # noqa: E501
        :rtype: str
        """
        return self._latest_rolledout_revision

    @latest_rolledout_revision.setter
    def latest_rolledout_revision(self, latest_rolledout_revision):
        """Sets the latest_rolledout_revision of this V1beta1ComponentStatusSpec.

        Latest revision name that is rolled out with 100 percent traffic  # noqa: E501

        :param latest_rolledout_revision: The latest_rolledout_revision of this V1beta1ComponentStatusSpec.  # noqa: E501
        :type: str
        """

        self._latest_rolledout_revision = latest_rolledout_revision

    @property
    def previous_rolledout_revision(self):
        """Gets the previous_rolledout_revision of this V1beta1ComponentStatusSpec.  # noqa: E501

        Previous revision name that is rolled out with 100 percent traffic  # noqa: E501

        :return: The previous_rolledout_revision of this V1beta1ComponentStatusSpec.  # noqa: E501
        :rtype: str
        """
        return self._previous_rolledout_revision

    @previous_rolledout_revision.setter
    def previous_rolledout_revision(self, previous_rolledout_revision):
        """Sets the previous_rolledout_revision of this V1beta1ComponentStatusSpec.

        Previous revision name that is rolled out with 100 percent traffic  # noqa: E501

        :param previous_rolledout_revision: The previous_rolledout_revision of this V1beta1ComponentStatusSpec.  # noqa: E501
        :type: str
        """

        self._previous_rolledout_revision = previous_rolledout_revision

    @property
    def rest_url(self):
        """Gets the rest_url of this V1beta1ComponentStatusSpec.  # noqa: E501


        :return: The rest_url of this V1beta1ComponentStatusSpec.  # noqa: E501
        :rtype: KnativeURL
        """
        return self._rest_url

    @rest_url.setter
    def rest_url(self, rest_url):
        """Sets the rest_url of this V1beta1ComponentStatusSpec.


        :param rest_url: The rest_url of this V1beta1ComponentStatusSpec.  # noqa: E501
        :type: KnativeURL
        """

        self._rest_url = rest_url

    @property
    def traffic(self):
        """Gets the traffic of this V1beta1ComponentStatusSpec.  # noqa: E501

        Traffic holds the configured traffic distribution for latest ready revision and previous rolled out revision.  # noqa: E501

        :return: The traffic of this V1beta1ComponentStatusSpec.  # noqa: E501
        :rtype: list[KnativeDevServingPkgApisServingV1TrafficTarget]
        """
        return self._traffic

    @traffic.setter
    def traffic(self, traffic):
        """Sets the traffic of this V1beta1ComponentStatusSpec.

        Traffic holds the configured traffic distribution for latest ready revision and previous rolled out revision.  # noqa: E501

        :param traffic: The traffic of this V1beta1ComponentStatusSpec.  # noqa: E501
        :type: list[KnativeDevServingPkgApisServingV1TrafficTarget]
        """

        self._traffic = traffic

    @property
    def url(self):
        """Gets the url of this V1beta1ComponentStatusSpec.  # noqa: E501


        :return: The url of this V1beta1ComponentStatusSpec.  # noqa: E501
        :rtype: KnativeURL
        """
        return self._url

    @url.setter
    def url(self, url):
        """Sets the url of this V1beta1ComponentStatusSpec.


        :param url: The url of this V1beta1ComponentStatusSpec.  # noqa: E501
        :type: KnativeURL
        """

        self._url = url

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.openapi_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1beta1ComponentStatusSpec):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, V1beta1ComponentStatusSpec):
            return True

        return self.to_dict() != other.to_dict()
