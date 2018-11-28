// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	sync "sync"

	exporter "github.com/alphagov/paas-prometheus-exporter/exporter"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

type FakeCFClient struct {
	ListAppsWithSpaceAndOrgStub        func() ([]cfclient.App, error)
	listAppsWithSpaceAndOrgMutex       sync.RWMutex
	listAppsWithSpaceAndOrgArgsForCall []struct {
	}
	listAppsWithSpaceAndOrgReturns struct {
		result1 []cfclient.App
		result2 error
	}
	listAppsWithSpaceAndOrgReturnsOnCall map[int]struct {
		result1 []cfclient.App
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCFClient) ListAppsWithSpaceAndOrg() ([]cfclient.App, error) {
	fake.listAppsWithSpaceAndOrgMutex.Lock()
	ret, specificReturn := fake.listAppsWithSpaceAndOrgReturnsOnCall[len(fake.listAppsWithSpaceAndOrgArgsForCall)]
	fake.listAppsWithSpaceAndOrgArgsForCall = append(fake.listAppsWithSpaceAndOrgArgsForCall, struct {
	}{})
	fake.recordInvocation("ListAppsWithSpaceAndOrg", []interface{}{})
	fake.listAppsWithSpaceAndOrgMutex.Unlock()
	if fake.ListAppsWithSpaceAndOrgStub != nil {
		return fake.ListAppsWithSpaceAndOrgStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listAppsWithSpaceAndOrgReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCFClient) ListAppsWithSpaceAndOrgCallCount() int {
	fake.listAppsWithSpaceAndOrgMutex.RLock()
	defer fake.listAppsWithSpaceAndOrgMutex.RUnlock()
	return len(fake.listAppsWithSpaceAndOrgArgsForCall)
}

func (fake *FakeCFClient) ListAppsWithSpaceAndOrgCalls(stub func() ([]cfclient.App, error)) {
	fake.listAppsWithSpaceAndOrgMutex.Lock()
	defer fake.listAppsWithSpaceAndOrgMutex.Unlock()
	fake.ListAppsWithSpaceAndOrgStub = stub
}

func (fake *FakeCFClient) ListAppsWithSpaceAndOrgReturns(result1 []cfclient.App, result2 error) {
	fake.listAppsWithSpaceAndOrgMutex.Lock()
	defer fake.listAppsWithSpaceAndOrgMutex.Unlock()
	fake.ListAppsWithSpaceAndOrgStub = nil
	fake.listAppsWithSpaceAndOrgReturns = struct {
		result1 []cfclient.App
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) ListAppsWithSpaceAndOrgReturnsOnCall(i int, result1 []cfclient.App, result2 error) {
	fake.listAppsWithSpaceAndOrgMutex.Lock()
	defer fake.listAppsWithSpaceAndOrgMutex.Unlock()
	fake.ListAppsWithSpaceAndOrgStub = nil
	if fake.listAppsWithSpaceAndOrgReturnsOnCall == nil {
		fake.listAppsWithSpaceAndOrgReturnsOnCall = make(map[int]struct {
			result1 []cfclient.App
			result2 error
		})
	}
	fake.listAppsWithSpaceAndOrgReturnsOnCall[i] = struct {
		result1 []cfclient.App
		result2 error
	}{result1, result2}
}

func (fake *FakeCFClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listAppsWithSpaceAndOrgMutex.RLock()
	defer fake.listAppsWithSpaceAndOrgMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCFClient) recordInvocation(key string, args []interface{}) {
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

var _ exporter.CFClient = new(FakeCFClient)