// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"bosh-vmrun-cpi/driver"
	"sync"
)

type FakeVdiskmanagerRunner struct {
	CreateDiskStub        func(string, int) error
	createDiskMutex       sync.RWMutex
	createDiskArgsForCall []struct {
		arg1 string
		arg2 int
	}
	createDiskReturns struct {
		result1 error
	}
	createDiskReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVdiskmanagerRunner) CreateDisk(arg1 string, arg2 int) error {
	fake.createDiskMutex.Lock()
	ret, specificReturn := fake.createDiskReturnsOnCall[len(fake.createDiskArgsForCall)]
	fake.createDiskArgsForCall = append(fake.createDiskArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("CreateDisk", []interface{}{arg1, arg2})
	fake.createDiskMutex.Unlock()
	if fake.CreateDiskStub != nil {
		return fake.CreateDiskStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createDiskReturns.result1
}

func (fake *FakeVdiskmanagerRunner) CreateDiskCallCount() int {
	fake.createDiskMutex.RLock()
	defer fake.createDiskMutex.RUnlock()
	return len(fake.createDiskArgsForCall)
}

func (fake *FakeVdiskmanagerRunner) CreateDiskArgsForCall(i int) (string, int) {
	fake.createDiskMutex.RLock()
	defer fake.createDiskMutex.RUnlock()
	return fake.createDiskArgsForCall[i].arg1, fake.createDiskArgsForCall[i].arg2
}

func (fake *FakeVdiskmanagerRunner) CreateDiskReturns(result1 error) {
	fake.CreateDiskStub = nil
	fake.createDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVdiskmanagerRunner) CreateDiskReturnsOnCall(i int, result1 error) {
	fake.CreateDiskStub = nil
	if fake.createDiskReturnsOnCall == nil {
		fake.createDiskReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createDiskReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVdiskmanagerRunner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createDiskMutex.RLock()
	defer fake.createDiskMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVdiskmanagerRunner) recordInvocation(key string, args []interface{}) {
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

var _ driver.VdiskmanagerRunner = new(FakeVdiskmanagerRunner)