// This file was generated by counterfeiter
package etcdclientfakes

import (
	"sync"

	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

type FakeKeysAPI struct {
	GetStub        func(ctx context.Context, key string, opts *client.GetOptions) (*client.Response, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		ctx  context.Context
		key  string
		opts *client.GetOptions
	}
	getReturns struct {
		result1 *client.Response
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *client.Response
		result2 error
	}
	SetStub        func(ctx context.Context, key, value string, opts *client.SetOptions) (*client.Response, error)
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		ctx   context.Context
		key   string
		value string
		opts  *client.SetOptions
	}
	setReturns struct {
		result1 *client.Response
		result2 error
	}
	setReturnsOnCall map[int]struct {
		result1 *client.Response
		result2 error
	}
	DeleteStub        func(ctx context.Context, key string, opts *client.DeleteOptions) (*client.Response, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		ctx  context.Context
		key  string
		opts *client.DeleteOptions
	}
	deleteReturns struct {
		result1 *client.Response
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 *client.Response
		result2 error
	}
	CreateStub        func(ctx context.Context, key, value string) (*client.Response, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		ctx   context.Context
		key   string
		value string
	}
	createReturns struct {
		result1 *client.Response
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *client.Response
		result2 error
	}
	CreateInOrderStub        func(ctx context.Context, dir, value string, opts *client.CreateInOrderOptions) (*client.Response, error)
	createInOrderMutex       sync.RWMutex
	createInOrderArgsForCall []struct {
		ctx   context.Context
		dir   string
		value string
		opts  *client.CreateInOrderOptions
	}
	createInOrderReturns struct {
		result1 *client.Response
		result2 error
	}
	createInOrderReturnsOnCall map[int]struct {
		result1 *client.Response
		result2 error
	}
	UpdateStub        func(ctx context.Context, key, value string) (*client.Response, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		ctx   context.Context
		key   string
		value string
	}
	updateReturns struct {
		result1 *client.Response
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 *client.Response
		result2 error
	}
	WatcherStub        func(key string, opts *client.WatcherOptions) client.Watcher
	watcherMutex       sync.RWMutex
	watcherArgsForCall []struct {
		key  string
		opts *client.WatcherOptions
	}
	watcherReturns struct {
		result1 client.Watcher
	}
	watcherReturnsOnCall map[int]struct {
		result1 client.Watcher
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKeysAPI) Get(ctx context.Context, key string, opts *client.GetOptions) (*client.Response, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		ctx  context.Context
		key  string
		opts *client.GetOptions
	}{ctx, key, opts})
	fake.recordInvocation("Get", []interface{}{ctx, key, opts})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(ctx, key, opts)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getReturns.result1, fake.getReturns.result2
}

func (fake *FakeKeysAPI) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeKeysAPI) GetArgsForCall(i int) (context.Context, string, *client.GetOptions) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].ctx, fake.getArgsForCall[i].key, fake.getArgsForCall[i].opts
}

func (fake *FakeKeysAPI) GetReturns(result1 *client.Response, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *client.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeKeysAPI) GetReturnsOnCall(i int, result1 *client.Response, result2 error) {
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *client.Response
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *client.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeKeysAPI) Set(ctx context.Context, key string, value string, opts *client.SetOptions) (*client.Response, error) {
	fake.setMutex.Lock()
	ret, specificReturn := fake.setReturnsOnCall[len(fake.setArgsForCall)]
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		ctx   context.Context
		key   string
		value string
		opts  *client.SetOptions
	}{ctx, key, value, opts})
	fake.recordInvocation("Set", []interface{}{ctx, key, value, opts})
	fake.setMutex.Unlock()
	if fake.SetStub != nil {
		return fake.SetStub(ctx, key, value, opts)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.setReturns.result1, fake.setReturns.result2
}

func (fake *FakeKeysAPI) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *FakeKeysAPI) SetArgsForCall(i int) (context.Context, string, string, *client.SetOptions) {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return fake.setArgsForCall[i].ctx, fake.setArgsForCall[i].key, fake.setArgsForCall[i].value, fake.setArgsForCall[i].opts
}

func (fake *FakeKeysAPI) SetReturns(result1 *client.Response, result2 error) {
	fake.SetStub = nil
	fake.setReturns = struct {
		result1 *client.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeKeysAPI) SetReturnsOnCall(i int, result1 *client.Response, result2 error) {
	fake.SetStub = nil
	if fake.setReturnsOnCall == nil {
		fake.setReturnsOnCall = make(map[int]struct {
			result1 *client.Response
			result2 error
		})
	}
	fake.setReturnsOnCall[i] = struct {
		result1 *client.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeKeysAPI) Delete(ctx context.Context, key string, opts *client.DeleteOptions) (*client.Response, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		ctx  context.Context
		key  string
		opts *client.DeleteOptions
	}{ctx, key, opts})
	fake.recordInvocation("Delete", []interface{}{ctx, key, opts})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(ctx, key, opts)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deleteReturns.result1, fake.deleteReturns.result2
}

func (fake *FakeKeysAPI) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeKeysAPI) DeleteArgsForCall(i int) (context.Context, string, *client.DeleteOptions) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].ctx, fake.deleteArgsForCall[i].key, fake.deleteArgsForCall[i].opts
}

func (fake *FakeKeysAPI) DeleteReturns(result1 *client.Response, result2 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 *client.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeKeysAPI) DeleteReturnsOnCall(i int, result1 *client.Response, result2 error) {
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 *client.Response
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 *client.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeKeysAPI) Create(ctx context.Context, key string, value string) (*client.Response, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		ctx   context.Context
		key   string
		value string
	}{ctx, key, value})
	fake.recordInvocation("Create", []interface{}{ctx, key, value})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(ctx, key, value)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createReturns.result1, fake.createReturns.result2
}

func (fake *FakeKeysAPI) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeKeysAPI) CreateArgsForCall(i int) (context.Context, string, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].ctx, fake.createArgsForCall[i].key, fake.createArgsForCall[i].value
}

func (fake *FakeKeysAPI) CreateReturns(result1 *client.Response, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *client.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeKeysAPI) CreateReturnsOnCall(i int, result1 *client.Response, result2 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *client.Response
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *client.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeKeysAPI) CreateInOrder(ctx context.Context, dir string, value string, opts *client.CreateInOrderOptions) (*client.Response, error) {
	fake.createInOrderMutex.Lock()
	ret, specificReturn := fake.createInOrderReturnsOnCall[len(fake.createInOrderArgsForCall)]
	fake.createInOrderArgsForCall = append(fake.createInOrderArgsForCall, struct {
		ctx   context.Context
		dir   string
		value string
		opts  *client.CreateInOrderOptions
	}{ctx, dir, value, opts})
	fake.recordInvocation("CreateInOrder", []interface{}{ctx, dir, value, opts})
	fake.createInOrderMutex.Unlock()
	if fake.CreateInOrderStub != nil {
		return fake.CreateInOrderStub(ctx, dir, value, opts)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.createInOrderReturns.result1, fake.createInOrderReturns.result2
}

func (fake *FakeKeysAPI) CreateInOrderCallCount() int {
	fake.createInOrderMutex.RLock()
	defer fake.createInOrderMutex.RUnlock()
	return len(fake.createInOrderArgsForCall)
}

func (fake *FakeKeysAPI) CreateInOrderArgsForCall(i int) (context.Context, string, string, *client.CreateInOrderOptions) {
	fake.createInOrderMutex.RLock()
	defer fake.createInOrderMutex.RUnlock()
	return fake.createInOrderArgsForCall[i].ctx, fake.createInOrderArgsForCall[i].dir, fake.createInOrderArgsForCall[i].value, fake.createInOrderArgsForCall[i].opts
}

func (fake *FakeKeysAPI) CreateInOrderReturns(result1 *client.Response, result2 error) {
	fake.CreateInOrderStub = nil
	fake.createInOrderReturns = struct {
		result1 *client.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeKeysAPI) CreateInOrderReturnsOnCall(i int, result1 *client.Response, result2 error) {
	fake.CreateInOrderStub = nil
	if fake.createInOrderReturnsOnCall == nil {
		fake.createInOrderReturnsOnCall = make(map[int]struct {
			result1 *client.Response
			result2 error
		})
	}
	fake.createInOrderReturnsOnCall[i] = struct {
		result1 *client.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeKeysAPI) Update(ctx context.Context, key string, value string) (*client.Response, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		ctx   context.Context
		key   string
		value string
	}{ctx, key, value})
	fake.recordInvocation("Update", []interface{}{ctx, key, value})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(ctx, key, value)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateReturns.result1, fake.updateReturns.result2
}

func (fake *FakeKeysAPI) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeKeysAPI) UpdateArgsForCall(i int) (context.Context, string, string) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].ctx, fake.updateArgsForCall[i].key, fake.updateArgsForCall[i].value
}

func (fake *FakeKeysAPI) UpdateReturns(result1 *client.Response, result2 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 *client.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeKeysAPI) UpdateReturnsOnCall(i int, result1 *client.Response, result2 error) {
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 *client.Response
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 *client.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeKeysAPI) Watcher(key string, opts *client.WatcherOptions) client.Watcher {
	fake.watcherMutex.Lock()
	ret, specificReturn := fake.watcherReturnsOnCall[len(fake.watcherArgsForCall)]
	fake.watcherArgsForCall = append(fake.watcherArgsForCall, struct {
		key  string
		opts *client.WatcherOptions
	}{key, opts})
	fake.recordInvocation("Watcher", []interface{}{key, opts})
	fake.watcherMutex.Unlock()
	if fake.WatcherStub != nil {
		return fake.WatcherStub(key, opts)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.watcherReturns.result1
}

func (fake *FakeKeysAPI) WatcherCallCount() int {
	fake.watcherMutex.RLock()
	defer fake.watcherMutex.RUnlock()
	return len(fake.watcherArgsForCall)
}

func (fake *FakeKeysAPI) WatcherArgsForCall(i int) (string, *client.WatcherOptions) {
	fake.watcherMutex.RLock()
	defer fake.watcherMutex.RUnlock()
	return fake.watcherArgsForCall[i].key, fake.watcherArgsForCall[i].opts
}

func (fake *FakeKeysAPI) WatcherReturns(result1 client.Watcher) {
	fake.WatcherStub = nil
	fake.watcherReturns = struct {
		result1 client.Watcher
	}{result1}
}

func (fake *FakeKeysAPI) WatcherReturnsOnCall(i int, result1 client.Watcher) {
	fake.WatcherStub = nil
	if fake.watcherReturnsOnCall == nil {
		fake.watcherReturnsOnCall = make(map[int]struct {
			result1 client.Watcher
		})
	}
	fake.watcherReturnsOnCall[i] = struct {
		result1 client.Watcher
	}{result1}
}

func (fake *FakeKeysAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.createInOrderMutex.RLock()
	defer fake.createInOrderMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.watcherMutex.RLock()
	defer fake.watcherMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeKeysAPI) recordInvocation(key string, args []interface{}) {
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

var _ client.KeysAPI = new(FakeKeysAPI)