// This file was generated by counterfeiter
package basefakes

import (
	"sync"

	"github.com/9corp/9volt/base"
)

type FakeIComponent struct {
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct{}
	startReturns     struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	StopStub        func() error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct{}
	stopReturns     struct {
		result1 error
	}
	stopReturnsOnCall map[int]struct {
		result1 error
	}
	IdentifyStub        func() string
	identifyMutex       sync.RWMutex
	identifyArgsForCall []struct{}
	identifyReturns     struct {
		result1 string
	}
	identifyReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIComponent) Start() error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct{}{})
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.startReturns.result1
}

func (fake *FakeIComponent) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeIComponent) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIComponent) StartReturnsOnCall(i int, result1 error) {
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIComponent) Stop() error {
	fake.stopMutex.Lock()
	ret, specificReturn := fake.stopReturnsOnCall[len(fake.stopArgsForCall)]
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct{}{})
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stopReturns.result1
}

func (fake *FakeIComponent) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeIComponent) StopReturns(result1 error) {
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIComponent) StopReturnsOnCall(i int, result1 error) {
	fake.StopStub = nil
	if fake.stopReturnsOnCall == nil {
		fake.stopReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIComponent) Identify() string {
	fake.identifyMutex.Lock()
	ret, specificReturn := fake.identifyReturnsOnCall[len(fake.identifyArgsForCall)]
	fake.identifyArgsForCall = append(fake.identifyArgsForCall, struct{}{})
	fake.recordInvocation("Identify", []interface{}{})
	fake.identifyMutex.Unlock()
	if fake.IdentifyStub != nil {
		return fake.IdentifyStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.identifyReturns.result1
}

func (fake *FakeIComponent) IdentifyCallCount() int {
	fake.identifyMutex.RLock()
	defer fake.identifyMutex.RUnlock()
	return len(fake.identifyArgsForCall)
}

func (fake *FakeIComponent) IdentifyReturns(result1 string) {
	fake.IdentifyStub = nil
	fake.identifyReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeIComponent) IdentifyReturnsOnCall(i int, result1 string) {
	fake.IdentifyStub = nil
	if fake.identifyReturnsOnCall == nil {
		fake.identifyReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.identifyReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeIComponent) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.identifyMutex.RLock()
	defer fake.identifyMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeIComponent) recordInvocation(key string, args []interface{}) {
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

var _ base.IComponent = new(FakeIComponent)
