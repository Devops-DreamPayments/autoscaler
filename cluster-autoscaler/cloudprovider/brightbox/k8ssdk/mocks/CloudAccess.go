// Copyright 2020 Brightbox Systems Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	gobrightbox "k8s.io/autoscaler/cluster-autoscaler/cloudprovider/brightbox/gobrightbox"

	mock "github.com/stretchr/testify/mock"
)

// CloudAccess is an autogenerated mock type for the CloudAccess type
type CloudAccess struct {
	mock.Mock
}

// AddServersToServerGroup provides a mock function with given fields: identifier, serverIds
func (_m *CloudAccess) AddServersToServerGroup(identifier string, serverIds []string) (*gobrightbox.ServerGroup, error) {
	ret := _m.Called(identifier, serverIds)

	var r0 *gobrightbox.ServerGroup
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []string) (*gobrightbox.ServerGroup, error)); ok {
		return rf(identifier, serverIds)
	}
	if rf, ok := ret.Get(0).(func(string, []string) *gobrightbox.ServerGroup); ok {
		r0 = rf(identifier, serverIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gobrightbox.ServerGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(identifier, serverIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloudIP provides a mock function with given fields: identifier
func (_m *CloudAccess) CloudIP(identifier string) (*gobrightbox.CloudIP, error) {
	ret := _m.Called(identifier)

	var r0 *gobrightbox.CloudIP
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*gobrightbox.CloudIP, error)); ok {
		return rf(identifier)
	}
	if rf, ok := ret.Get(0).(func(string) *gobrightbox.CloudIP); ok {
		r0 = rf(identifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gobrightbox.CloudIP)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(identifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloudIPs provides a mock function with given fields:
func (_m *CloudAccess) CloudIPs() ([]gobrightbox.CloudIP, error) {
	ret := _m.Called()

	var r0 []gobrightbox.CloudIP
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]gobrightbox.CloudIP, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []gobrightbox.CloudIP); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gobrightbox.CloudIP)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfigMap provides a mock function with given fields: identifier
func (_m *CloudAccess) ConfigMap(identifier string) (*gobrightbox.ConfigMap, error) {
	ret := _m.Called(identifier)

	var r0 *gobrightbox.ConfigMap
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*gobrightbox.ConfigMap, error)); ok {
		return rf(identifier)
	}
	if rf, ok := ret.Get(0).(func(string) *gobrightbox.ConfigMap); ok {
		r0 = rf(identifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gobrightbox.ConfigMap)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(identifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConfigMaps provides a mock function with given fields:
func (_m *CloudAccess) ConfigMaps() ([]gobrightbox.ConfigMap, error) {
	ret := _m.Called()

	var r0 []gobrightbox.ConfigMap
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]gobrightbox.ConfigMap, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []gobrightbox.ConfigMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gobrightbox.ConfigMap)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCloudIP provides a mock function with given fields: newCloudIP
func (_m *CloudAccess) CreateCloudIP(newCloudIP *gobrightbox.CloudIPOptions) (*gobrightbox.CloudIP, error) {
	ret := _m.Called(newCloudIP)

	var r0 *gobrightbox.CloudIP
	var r1 error
	if rf, ok := ret.Get(0).(func(*gobrightbox.CloudIPOptions) (*gobrightbox.CloudIP, error)); ok {
		return rf(newCloudIP)
	}
	if rf, ok := ret.Get(0).(func(*gobrightbox.CloudIPOptions) *gobrightbox.CloudIP); ok {
		r0 = rf(newCloudIP)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gobrightbox.CloudIP)
		}
	}

	if rf, ok := ret.Get(1).(func(*gobrightbox.CloudIPOptions) error); ok {
		r1 = rf(newCloudIP)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFirewallPolicy provides a mock function with given fields: policyOptions
func (_m *CloudAccess) CreateFirewallPolicy(policyOptions *gobrightbox.FirewallPolicyOptions) (*gobrightbox.FirewallPolicy, error) {
	ret := _m.Called(policyOptions)

	var r0 *gobrightbox.FirewallPolicy
	var r1 error
	if rf, ok := ret.Get(0).(func(*gobrightbox.FirewallPolicyOptions) (*gobrightbox.FirewallPolicy, error)); ok {
		return rf(policyOptions)
	}
	if rf, ok := ret.Get(0).(func(*gobrightbox.FirewallPolicyOptions) *gobrightbox.FirewallPolicy); ok {
		r0 = rf(policyOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gobrightbox.FirewallPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func(*gobrightbox.FirewallPolicyOptions) error); ok {
		r1 = rf(policyOptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFirewallRule provides a mock function with given fields: ruleOptions
func (_m *CloudAccess) CreateFirewallRule(ruleOptions *gobrightbox.FirewallRuleOptions) (*gobrightbox.FirewallRule, error) {
	ret := _m.Called(ruleOptions)

	var r0 *gobrightbox.FirewallRule
	var r1 error
	if rf, ok := ret.Get(0).(func(*gobrightbox.FirewallRuleOptions) (*gobrightbox.FirewallRule, error)); ok {
		return rf(ruleOptions)
	}
	if rf, ok := ret.Get(0).(func(*gobrightbox.FirewallRuleOptions) *gobrightbox.FirewallRule); ok {
		r0 = rf(ruleOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gobrightbox.FirewallRule)
		}
	}

	if rf, ok := ret.Get(1).(func(*gobrightbox.FirewallRuleOptions) error); ok {
		r1 = rf(ruleOptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateLoadBalancer provides a mock function with given fields: newDetails
func (_m *CloudAccess) CreateLoadBalancer(newDetails *gobrightbox.LoadBalancerOptions) (*gobrightbox.LoadBalancer, error) {
	ret := _m.Called(newDetails)

	var r0 *gobrightbox.LoadBalancer
	var r1 error
	if rf, ok := ret.Get(0).(func(*gobrightbox.LoadBalancerOptions) (*gobrightbox.LoadBalancer, error)); ok {
		return rf(newDetails)
	}
	if rf, ok := ret.Get(0).(func(*gobrightbox.LoadBalancerOptions) *gobrightbox.LoadBalancer); ok {
		r0 = rf(newDetails)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gobrightbox.LoadBalancer)
		}
	}

	if rf, ok := ret.Get(1).(func(*gobrightbox.LoadBalancerOptions) error); ok {
		r1 = rf(newDetails)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateServer provides a mock function with given fields: newServer
func (_m *CloudAccess) CreateServer(newServer *gobrightbox.ServerOptions) (*gobrightbox.Server, error) {
	ret := _m.Called(newServer)

	var r0 *gobrightbox.Server
	var r1 error
	if rf, ok := ret.Get(0).(func(*gobrightbox.ServerOptions) (*gobrightbox.Server, error)); ok {
		return rf(newServer)
	}
	if rf, ok := ret.Get(0).(func(*gobrightbox.ServerOptions) *gobrightbox.Server); ok {
		r0 = rf(newServer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gobrightbox.Server)
		}
	}

	if rf, ok := ret.Get(1).(func(*gobrightbox.ServerOptions) error); ok {
		r1 = rf(newServer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateServerGroup provides a mock function with given fields: newServerGroup
func (_m *CloudAccess) CreateServerGroup(newServerGroup *gobrightbox.ServerGroupOptions) (*gobrightbox.ServerGroup, error) {
	ret := _m.Called(newServerGroup)

	var r0 *gobrightbox.ServerGroup
	var r1 error
	if rf, ok := ret.Get(0).(func(*gobrightbox.ServerGroupOptions) (*gobrightbox.ServerGroup, error)); ok {
		return rf(newServerGroup)
	}
	if rf, ok := ret.Get(0).(func(*gobrightbox.ServerGroupOptions) *gobrightbox.ServerGroup); ok {
		r0 = rf(newServerGroup)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gobrightbox.ServerGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(*gobrightbox.ServerGroupOptions) error); ok {
		r1 = rf(newServerGroup)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DestroyCloudIP provides a mock function with given fields: identifier
func (_m *CloudAccess) DestroyCloudIP(identifier string) error {
	ret := _m.Called(identifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(identifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DestroyFirewallPolicy provides a mock function with given fields: identifier
func (_m *CloudAccess) DestroyFirewallPolicy(identifier string) error {
	ret := _m.Called(identifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(identifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DestroyLoadBalancer provides a mock function with given fields: identifier
func (_m *CloudAccess) DestroyLoadBalancer(identifier string) error {
	ret := _m.Called(identifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(identifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DestroyServer provides a mock function with given fields: identifier
func (_m *CloudAccess) DestroyServer(identifier string) error {
	ret := _m.Called(identifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(identifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DestroyServerGroup provides a mock function with given fields: identifier
func (_m *CloudAccess) DestroyServerGroup(identifier string) error {
	ret := _m.Called(identifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(identifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirewallPolicies provides a mock function with given fields:
func (_m *CloudAccess) FirewallPolicies() ([]gobrightbox.FirewallPolicy, error) {
	ret := _m.Called()

	var r0 []gobrightbox.FirewallPolicy
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]gobrightbox.FirewallPolicy, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []gobrightbox.FirewallPolicy); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gobrightbox.FirewallPolicy)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Images provides a mock function with given fields:
func (_m *CloudAccess) Images() ([]gobrightbox.Image, error) {
	ret := _m.Called()

	var r0 []gobrightbox.Image
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]gobrightbox.Image, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []gobrightbox.Image); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gobrightbox.Image)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadBalancer provides a mock function with given fields: identifier
func (_m *CloudAccess) LoadBalancer(identifier string) (*gobrightbox.LoadBalancer, error) {
	ret := _m.Called(identifier)

	var r0 *gobrightbox.LoadBalancer
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*gobrightbox.LoadBalancer, error)); ok {
		return rf(identifier)
	}
	if rf, ok := ret.Get(0).(func(string) *gobrightbox.LoadBalancer); ok {
		r0 = rf(identifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gobrightbox.LoadBalancer)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(identifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadBalancers provides a mock function with given fields:
func (_m *CloudAccess) LoadBalancers() ([]gobrightbox.LoadBalancer, error) {
	ret := _m.Called()

	var r0 []gobrightbox.LoadBalancer
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]gobrightbox.LoadBalancer, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []gobrightbox.LoadBalancer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gobrightbox.LoadBalancer)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MapCloudIP provides a mock function with given fields: identifier, destination
func (_m *CloudAccess) MapCloudIP(identifier string, destination string) error {
	ret := _m.Called(identifier, destination)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(identifier, destination)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveServersFromServerGroup provides a mock function with given fields: identifier, serverIds
func (_m *CloudAccess) RemoveServersFromServerGroup(identifier string, serverIds []string) (*gobrightbox.ServerGroup, error) {
	ret := _m.Called(identifier, serverIds)

	var r0 *gobrightbox.ServerGroup
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []string) (*gobrightbox.ServerGroup, error)); ok {
		return rf(identifier, serverIds)
	}
	if rf, ok := ret.Get(0).(func(string, []string) *gobrightbox.ServerGroup); ok {
		r0 = rf(identifier, serverIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gobrightbox.ServerGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []string) error); ok {
		r1 = rf(identifier, serverIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Server provides a mock function with given fields: identifier
func (_m *CloudAccess) Server(identifier string) (*gobrightbox.Server, error) {
	ret := _m.Called(identifier)

	var r0 *gobrightbox.Server
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*gobrightbox.Server, error)); ok {
		return rf(identifier)
	}
	if rf, ok := ret.Get(0).(func(string) *gobrightbox.Server); ok {
		r0 = rf(identifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gobrightbox.Server)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(identifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServerGroup provides a mock function with given fields: identifier
func (_m *CloudAccess) ServerGroup(identifier string) (*gobrightbox.ServerGroup, error) {
	ret := _m.Called(identifier)

	var r0 *gobrightbox.ServerGroup
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*gobrightbox.ServerGroup, error)); ok {
		return rf(identifier)
	}
	if rf, ok := ret.Get(0).(func(string) *gobrightbox.ServerGroup); ok {
		r0 = rf(identifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gobrightbox.ServerGroup)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(identifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServerGroups provides a mock function with given fields:
func (_m *CloudAccess) ServerGroups() ([]gobrightbox.ServerGroup, error) {
	ret := _m.Called()

	var r0 []gobrightbox.ServerGroup
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]gobrightbox.ServerGroup, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []gobrightbox.ServerGroup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gobrightbox.ServerGroup)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServerType provides a mock function with given fields: identifier
func (_m *CloudAccess) ServerType(identifier string) (*gobrightbox.ServerType, error) {
	ret := _m.Called(identifier)

	var r0 *gobrightbox.ServerType
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*gobrightbox.ServerType, error)); ok {
		return rf(identifier)
	}
	if rf, ok := ret.Get(0).(func(string) *gobrightbox.ServerType); ok {
		r0 = rf(identifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gobrightbox.ServerType)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(identifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServerTypes provides a mock function with given fields:
func (_m *CloudAccess) ServerTypes() ([]gobrightbox.ServerType, error) {
	ret := _m.Called()

	var r0 []gobrightbox.ServerType
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]gobrightbox.ServerType, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []gobrightbox.ServerType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gobrightbox.ServerType)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnMapCloudIP provides a mock function with given fields: identifier
func (_m *CloudAccess) UnMapCloudIP(identifier string) error {
	ret := _m.Called(identifier)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(identifier)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateFirewallRule provides a mock function with given fields: ruleOptions
func (_m *CloudAccess) UpdateFirewallRule(ruleOptions *gobrightbox.FirewallRuleOptions) (*gobrightbox.FirewallRule, error) {
	ret := _m.Called(ruleOptions)

	var r0 *gobrightbox.FirewallRule
	var r1 error
	if rf, ok := ret.Get(0).(func(*gobrightbox.FirewallRuleOptions) (*gobrightbox.FirewallRule, error)); ok {
		return rf(ruleOptions)
	}
	if rf, ok := ret.Get(0).(func(*gobrightbox.FirewallRuleOptions) *gobrightbox.FirewallRule); ok {
		r0 = rf(ruleOptions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gobrightbox.FirewallRule)
		}
	}

	if rf, ok := ret.Get(1).(func(*gobrightbox.FirewallRuleOptions) error); ok {
		r1 = rf(ruleOptions)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateLoadBalancer provides a mock function with given fields: newDetails
func (_m *CloudAccess) UpdateLoadBalancer(newDetails *gobrightbox.LoadBalancerOptions) (*gobrightbox.LoadBalancer, error) {
	ret := _m.Called(newDetails)

	var r0 *gobrightbox.LoadBalancer
	var r1 error
	if rf, ok := ret.Get(0).(func(*gobrightbox.LoadBalancerOptions) (*gobrightbox.LoadBalancer, error)); ok {
		return rf(newDetails)
	}
	if rf, ok := ret.Get(0).(func(*gobrightbox.LoadBalancerOptions) *gobrightbox.LoadBalancer); ok {
		r0 = rf(newDetails)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gobrightbox.LoadBalancer)
		}
	}

	if rf, ok := ret.Get(1).(func(*gobrightbox.LoadBalancerOptions) error); ok {
		r1 = rf(newDetails)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCloudAccess interface {
	mock.TestingT
	Cleanup(func())
}

// NewCloudAccess creates a new instance of CloudAccess. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCloudAccess(t mockConstructorTestingTNewCloudAccess) *CloudAccess {
	mock := &CloudAccess{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
