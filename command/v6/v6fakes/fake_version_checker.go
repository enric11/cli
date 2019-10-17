// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	"sync"

	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeVersionChecker struct {
	AuthorizationEndpointStub        func() string
	authorizationEndpointMutex       sync.RWMutex
	authorizationEndpointArgsForCall []struct {
	}
	authorizationEndpointReturns struct {
		result1 string
	}
	authorizationEndpointReturnsOnCall map[int]struct {
		result1 string
	}
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct {
	}
	cloudControllerAPIVersionReturns struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	MinCLIVersionStub        func() string
	minCLIVersionMutex       sync.RWMutex
	minCLIVersionArgsForCall []struct {
	}
	minCLIVersionReturns struct {
		result1 string
	}
	minCLIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVersionChecker) AuthorizationEndpoint() string {
	fake.authorizationEndpointMutex.Lock()
	ret, specificReturn := fake.authorizationEndpointReturnsOnCall[len(fake.authorizationEndpointArgsForCall)]
	fake.authorizationEndpointArgsForCall = append(fake.authorizationEndpointArgsForCall, struct {
	}{})
	fake.recordInvocation("AuthorizationEndpoint", []interface{}{})
	fake.authorizationEndpointMutex.Unlock()
	if fake.AuthorizationEndpointStub != nil {
		return fake.AuthorizationEndpointStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.authorizationEndpointReturns
	return fakeReturns.result1
}

func (fake *FakeVersionChecker) AuthorizationEndpointCallCount() int {
	fake.authorizationEndpointMutex.RLock()
	defer fake.authorizationEndpointMutex.RUnlock()
	return len(fake.authorizationEndpointArgsForCall)
}

func (fake *FakeVersionChecker) AuthorizationEndpointCalls(stub func() string) {
	fake.authorizationEndpointMutex.Lock()
	defer fake.authorizationEndpointMutex.Unlock()
	fake.AuthorizationEndpointStub = stub
}

func (fake *FakeVersionChecker) AuthorizationEndpointReturns(result1 string) {
	fake.authorizationEndpointMutex.Lock()
	defer fake.authorizationEndpointMutex.Unlock()
	fake.AuthorizationEndpointStub = nil
	fake.authorizationEndpointReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVersionChecker) AuthorizationEndpointReturnsOnCall(i int, result1 string) {
	fake.authorizationEndpointMutex.Lock()
	defer fake.authorizationEndpointMutex.Unlock()
	fake.AuthorizationEndpointStub = nil
	if fake.authorizationEndpointReturnsOnCall == nil {
		fake.authorizationEndpointReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.authorizationEndpointReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeVersionChecker) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cloudControllerAPIVersionReturns
	return fakeReturns.result1
}

func (fake *FakeVersionChecker) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeVersionChecker) CloudControllerAPIVersionCalls(stub func() string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = stub
}

func (fake *FakeVersionChecker) CloudControllerAPIVersionReturns(result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVersionChecker) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeVersionChecker) MinCLIVersion() string {
	fake.minCLIVersionMutex.Lock()
	ret, specificReturn := fake.minCLIVersionReturnsOnCall[len(fake.minCLIVersionArgsForCall)]
	fake.minCLIVersionArgsForCall = append(fake.minCLIVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("MinCLIVersion", []interface{}{})
	fake.minCLIVersionMutex.Unlock()
	if fake.MinCLIVersionStub != nil {
		return fake.MinCLIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.minCLIVersionReturns
	return fakeReturns.result1
}

func (fake *FakeVersionChecker) MinCLIVersionCallCount() int {
	fake.minCLIVersionMutex.RLock()
	defer fake.minCLIVersionMutex.RUnlock()
	return len(fake.minCLIVersionArgsForCall)
}

func (fake *FakeVersionChecker) MinCLIVersionCalls(stub func() string) {
	fake.minCLIVersionMutex.Lock()
	defer fake.minCLIVersionMutex.Unlock()
	fake.MinCLIVersionStub = stub
}

func (fake *FakeVersionChecker) MinCLIVersionReturns(result1 string) {
	fake.minCLIVersionMutex.Lock()
	defer fake.minCLIVersionMutex.Unlock()
	fake.MinCLIVersionStub = nil
	fake.minCLIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVersionChecker) MinCLIVersionReturnsOnCall(i int, result1 string) {
	fake.minCLIVersionMutex.Lock()
	defer fake.minCLIVersionMutex.Unlock()
	fake.MinCLIVersionStub = nil
	if fake.minCLIVersionReturnsOnCall == nil {
		fake.minCLIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.minCLIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeVersionChecker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.authorizationEndpointMutex.RLock()
	defer fake.authorizationEndpointMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.minCLIVersionMutex.RLock()
	defer fake.minCLIVersionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVersionChecker) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ v6.VersionChecker = new(FakeVersionChecker)
