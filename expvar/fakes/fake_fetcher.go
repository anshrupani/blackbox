// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/blackbox/expvar"
)

type FakeFetcher struct {
	FetchStub        func() (expvar.Expvars, error)
	fetchMutex       sync.RWMutex
	fetchArgsForCall []struct{}
	fetchReturns     struct {
		result1 expvar.Expvars
		result2 error
	}
}

func (fake *FakeFetcher) Fetch() (expvar.Expvars, error) {
	fake.fetchMutex.Lock()
	fake.fetchArgsForCall = append(fake.fetchArgsForCall, struct{}{})
	fake.fetchMutex.Unlock()
	if fake.FetchStub != nil {
		return fake.FetchStub()
	} else {
		return fake.fetchReturns.result1, fake.fetchReturns.result2
	}
}

func (fake *FakeFetcher) FetchCallCount() int {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return len(fake.fetchArgsForCall)
}

func (fake *FakeFetcher) FetchReturns(result1 expvar.Expvars, result2 error) {
	fake.FetchStub = nil
	fake.fetchReturns = struct {
		result1 expvar.Expvars
		result2 error
	}{result1, result2}
}

var _ expvar.Fetcher = new(FakeFetcher)
