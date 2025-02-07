// Code generated by counterfeiter. DO NOT EDIT.
package apkfakes

import (
	"archive/tar"
	"sync"
	"time"

	"chainguard.dev/apko/pkg/apk/impl"
)

type FakeApkImplementation struct {
	FixateWorldStub        func(bool, bool, bool, *time.Time) error
	fixateWorldMutex       sync.RWMutex
	fixateWorldArgsForCall []struct {
		arg1 bool
		arg2 bool
		arg3 bool
		arg4 *time.Time
	}
	fixateWorldReturns struct {
		result1 error
	}
	fixateWorldReturnsOnCall map[int]struct {
		result1 error
	}
	GetInstalledStub        func() ([]*impl.InstalledPackage, error)
	getInstalledMutex       sync.RWMutex
	getInstalledArgsForCall []struct {
	}
	getInstalledReturns struct {
		result1 []*impl.InstalledPackage
		result2 error
	}
	getInstalledReturnsOnCall map[int]struct {
		result1 []*impl.InstalledPackage
		result2 error
	}
	GetRepositoriesStub        func() ([]string, error)
	getRepositoriesMutex       sync.RWMutex
	getRepositoriesArgsForCall []struct {
	}
	getRepositoriesReturns struct {
		result1 []string
		result2 error
	}
	getRepositoriesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	GetWorldStub        func() ([]string, error)
	getWorldMutex       sync.RWMutex
	getWorldArgsForCall []struct {
	}
	getWorldReturns struct {
		result1 []string
		result2 error
	}
	getWorldReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	InitDBStub        func(...string) error
	initDBMutex       sync.RWMutex
	initDBArgsForCall []struct {
		arg1 []string
	}
	initDBReturns struct {
		result1 error
	}
	initDBReturnsOnCall map[int]struct {
		result1 error
	}
	InitKeyringStub        func([]string, []string) error
	initKeyringMutex       sync.RWMutex
	initKeyringArgsForCall []struct {
		arg1 []string
		arg2 []string
	}
	initKeyringReturns struct {
		result1 error
	}
	initKeyringReturnsOnCall map[int]struct {
		result1 error
	}
	ListInitFilesStub        func() []tar.Header
	listInitFilesMutex       sync.RWMutex
	listInitFilesArgsForCall []struct {
	}
	listInitFilesReturns struct {
		result1 []tar.Header
	}
	listInitFilesReturnsOnCall map[int]struct {
		result1 []tar.Header
	}
	SetRepositoriesStub        func([]string) error
	setRepositoriesMutex       sync.RWMutex
	setRepositoriesArgsForCall []struct {
		arg1 []string
	}
	setRepositoriesReturns struct {
		result1 error
	}
	setRepositoriesReturnsOnCall map[int]struct {
		result1 error
	}
	SetWorldStub        func([]string) error
	setWorldMutex       sync.RWMutex
	setWorldArgsForCall []struct {
		arg1 []string
	}
	setWorldReturns struct {
		result1 error
	}
	setWorldReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeApkImplementation) FixateWorld(arg1 bool, arg2 bool, arg3 bool, arg4 *time.Time) error {
	fake.fixateWorldMutex.Lock()
	ret, specificReturn := fake.fixateWorldReturnsOnCall[len(fake.fixateWorldArgsForCall)]
	fake.fixateWorldArgsForCall = append(fake.fixateWorldArgsForCall, struct {
		arg1 bool
		arg2 bool
		arg3 bool
		arg4 *time.Time
	}{arg1, arg2, arg3, arg4})
	stub := fake.FixateWorldStub
	fakeReturns := fake.fixateWorldReturns
	fake.recordInvocation("FixateWorld", []interface{}{arg1, arg2, arg3, arg4})
	fake.fixateWorldMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeApkImplementation) FixateWorldCallCount() int {
	fake.fixateWorldMutex.RLock()
	defer fake.fixateWorldMutex.RUnlock()
	return len(fake.fixateWorldArgsForCall)
}

func (fake *FakeApkImplementation) FixateWorldCalls(stub func(bool, bool, bool, *time.Time) error) {
	fake.fixateWorldMutex.Lock()
	defer fake.fixateWorldMutex.Unlock()
	fake.FixateWorldStub = stub
}

func (fake *FakeApkImplementation) FixateWorldArgsForCall(i int) (bool, bool, bool, *time.Time) {
	fake.fixateWorldMutex.RLock()
	defer fake.fixateWorldMutex.RUnlock()
	argsForCall := fake.fixateWorldArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeApkImplementation) FixateWorldReturns(result1 error) {
	fake.fixateWorldMutex.Lock()
	defer fake.fixateWorldMutex.Unlock()
	fake.FixateWorldStub = nil
	fake.fixateWorldReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) FixateWorldReturnsOnCall(i int, result1 error) {
	fake.fixateWorldMutex.Lock()
	defer fake.fixateWorldMutex.Unlock()
	fake.FixateWorldStub = nil
	if fake.fixateWorldReturnsOnCall == nil {
		fake.fixateWorldReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.fixateWorldReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) GetInstalled() ([]*impl.InstalledPackage, error) {
	fake.getInstalledMutex.Lock()
	ret, specificReturn := fake.getInstalledReturnsOnCall[len(fake.getInstalledArgsForCall)]
	fake.getInstalledArgsForCall = append(fake.getInstalledArgsForCall, struct {
	}{})
	stub := fake.GetInstalledStub
	fakeReturns := fake.getInstalledReturns
	fake.recordInvocation("GetInstalled", []interface{}{})
	fake.getInstalledMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeApkImplementation) GetInstalledCallCount() int {
	fake.getInstalledMutex.RLock()
	defer fake.getInstalledMutex.RUnlock()
	return len(fake.getInstalledArgsForCall)
}

func (fake *FakeApkImplementation) GetInstalledCalls(stub func() ([]*impl.InstalledPackage, error)) {
	fake.getInstalledMutex.Lock()
	defer fake.getInstalledMutex.Unlock()
	fake.GetInstalledStub = stub
}

func (fake *FakeApkImplementation) GetInstalledReturns(result1 []*impl.InstalledPackage, result2 error) {
	fake.getInstalledMutex.Lock()
	defer fake.getInstalledMutex.Unlock()
	fake.GetInstalledStub = nil
	fake.getInstalledReturns = struct {
		result1 []*impl.InstalledPackage
		result2 error
	}{result1, result2}
}

func (fake *FakeApkImplementation) GetInstalledReturnsOnCall(i int, result1 []*impl.InstalledPackage, result2 error) {
	fake.getInstalledMutex.Lock()
	defer fake.getInstalledMutex.Unlock()
	fake.GetInstalledStub = nil
	if fake.getInstalledReturnsOnCall == nil {
		fake.getInstalledReturnsOnCall = make(map[int]struct {
			result1 []*impl.InstalledPackage
			result2 error
		})
	}
	fake.getInstalledReturnsOnCall[i] = struct {
		result1 []*impl.InstalledPackage
		result2 error
	}{result1, result2}
}

func (fake *FakeApkImplementation) GetRepositories() ([]string, error) {
	fake.getRepositoriesMutex.Lock()
	ret, specificReturn := fake.getRepositoriesReturnsOnCall[len(fake.getRepositoriesArgsForCall)]
	fake.getRepositoriesArgsForCall = append(fake.getRepositoriesArgsForCall, struct {
	}{})
	stub := fake.GetRepositoriesStub
	fakeReturns := fake.getRepositoriesReturns
	fake.recordInvocation("GetRepositories", []interface{}{})
	fake.getRepositoriesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeApkImplementation) GetRepositoriesCallCount() int {
	fake.getRepositoriesMutex.RLock()
	defer fake.getRepositoriesMutex.RUnlock()
	return len(fake.getRepositoriesArgsForCall)
}

func (fake *FakeApkImplementation) GetRepositoriesCalls(stub func() ([]string, error)) {
	fake.getRepositoriesMutex.Lock()
	defer fake.getRepositoriesMutex.Unlock()
	fake.GetRepositoriesStub = stub
}

func (fake *FakeApkImplementation) GetRepositoriesReturns(result1 []string, result2 error) {
	fake.getRepositoriesMutex.Lock()
	defer fake.getRepositoriesMutex.Unlock()
	fake.GetRepositoriesStub = nil
	fake.getRepositoriesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeApkImplementation) GetRepositoriesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.getRepositoriesMutex.Lock()
	defer fake.getRepositoriesMutex.Unlock()
	fake.GetRepositoriesStub = nil
	if fake.getRepositoriesReturnsOnCall == nil {
		fake.getRepositoriesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.getRepositoriesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeApkImplementation) GetWorld() ([]string, error) {
	fake.getWorldMutex.Lock()
	ret, specificReturn := fake.getWorldReturnsOnCall[len(fake.getWorldArgsForCall)]
	fake.getWorldArgsForCall = append(fake.getWorldArgsForCall, struct {
	}{})
	stub := fake.GetWorldStub
	fakeReturns := fake.getWorldReturns
	fake.recordInvocation("GetWorld", []interface{}{})
	fake.getWorldMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeApkImplementation) GetWorldCallCount() int {
	fake.getWorldMutex.RLock()
	defer fake.getWorldMutex.RUnlock()
	return len(fake.getWorldArgsForCall)
}

func (fake *FakeApkImplementation) GetWorldCalls(stub func() ([]string, error)) {
	fake.getWorldMutex.Lock()
	defer fake.getWorldMutex.Unlock()
	fake.GetWorldStub = stub
}

func (fake *FakeApkImplementation) GetWorldReturns(result1 []string, result2 error) {
	fake.getWorldMutex.Lock()
	defer fake.getWorldMutex.Unlock()
	fake.GetWorldStub = nil
	fake.getWorldReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeApkImplementation) GetWorldReturnsOnCall(i int, result1 []string, result2 error) {
	fake.getWorldMutex.Lock()
	defer fake.getWorldMutex.Unlock()
	fake.GetWorldStub = nil
	if fake.getWorldReturnsOnCall == nil {
		fake.getWorldReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.getWorldReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeApkImplementation) InitDB(arg1 ...string) error {
	fake.initDBMutex.Lock()
	ret, specificReturn := fake.initDBReturnsOnCall[len(fake.initDBArgsForCall)]
	fake.initDBArgsForCall = append(fake.initDBArgsForCall, struct {
		arg1 []string
	}{arg1})
	stub := fake.InitDBStub
	fakeReturns := fake.initDBReturns
	fake.recordInvocation("InitDB", []interface{}{arg1})
	fake.initDBMutex.Unlock()
	if stub != nil {
		return stub(arg1...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeApkImplementation) InitDBCallCount() int {
	fake.initDBMutex.RLock()
	defer fake.initDBMutex.RUnlock()
	return len(fake.initDBArgsForCall)
}

func (fake *FakeApkImplementation) InitDBCalls(stub func(...string) error) {
	fake.initDBMutex.Lock()
	defer fake.initDBMutex.Unlock()
	fake.InitDBStub = stub
}

func (fake *FakeApkImplementation) InitDBArgsForCall(i int) []string {
	fake.initDBMutex.RLock()
	defer fake.initDBMutex.RUnlock()
	argsForCall := fake.initDBArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeApkImplementation) InitDBReturns(result1 error) {
	fake.initDBMutex.Lock()
	defer fake.initDBMutex.Unlock()
	fake.InitDBStub = nil
	fake.initDBReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) InitDBReturnsOnCall(i int, result1 error) {
	fake.initDBMutex.Lock()
	defer fake.initDBMutex.Unlock()
	fake.InitDBStub = nil
	if fake.initDBReturnsOnCall == nil {
		fake.initDBReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initDBReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) InitKeyring(arg1 []string, arg2 []string) error {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.initKeyringMutex.Lock()
	ret, specificReturn := fake.initKeyringReturnsOnCall[len(fake.initKeyringArgsForCall)]
	fake.initKeyringArgsForCall = append(fake.initKeyringArgsForCall, struct {
		arg1 []string
		arg2 []string
	}{arg1Copy, arg2Copy})
	stub := fake.InitKeyringStub
	fakeReturns := fake.initKeyringReturns
	fake.recordInvocation("InitKeyring", []interface{}{arg1Copy, arg2Copy})
	fake.initKeyringMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeApkImplementation) InitKeyringCallCount() int {
	fake.initKeyringMutex.RLock()
	defer fake.initKeyringMutex.RUnlock()
	return len(fake.initKeyringArgsForCall)
}

func (fake *FakeApkImplementation) InitKeyringCalls(stub func([]string, []string) error) {
	fake.initKeyringMutex.Lock()
	defer fake.initKeyringMutex.Unlock()
	fake.InitKeyringStub = stub
}

func (fake *FakeApkImplementation) InitKeyringArgsForCall(i int) ([]string, []string) {
	fake.initKeyringMutex.RLock()
	defer fake.initKeyringMutex.RUnlock()
	argsForCall := fake.initKeyringArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeApkImplementation) InitKeyringReturns(result1 error) {
	fake.initKeyringMutex.Lock()
	defer fake.initKeyringMutex.Unlock()
	fake.InitKeyringStub = nil
	fake.initKeyringReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) InitKeyringReturnsOnCall(i int, result1 error) {
	fake.initKeyringMutex.Lock()
	defer fake.initKeyringMutex.Unlock()
	fake.InitKeyringStub = nil
	if fake.initKeyringReturnsOnCall == nil {
		fake.initKeyringReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initKeyringReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) ListInitFiles() []tar.Header {
	fake.listInitFilesMutex.Lock()
	ret, specificReturn := fake.listInitFilesReturnsOnCall[len(fake.listInitFilesArgsForCall)]
	fake.listInitFilesArgsForCall = append(fake.listInitFilesArgsForCall, struct {
	}{})
	stub := fake.ListInitFilesStub
	fakeReturns := fake.listInitFilesReturns
	fake.recordInvocation("ListInitFiles", []interface{}{})
	fake.listInitFilesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeApkImplementation) ListInitFilesCallCount() int {
	fake.listInitFilesMutex.RLock()
	defer fake.listInitFilesMutex.RUnlock()
	return len(fake.listInitFilesArgsForCall)
}

func (fake *FakeApkImplementation) ListInitFilesCalls(stub func() []tar.Header) {
	fake.listInitFilesMutex.Lock()
	defer fake.listInitFilesMutex.Unlock()
	fake.ListInitFilesStub = stub
}

func (fake *FakeApkImplementation) ListInitFilesReturns(result1 []tar.Header) {
	fake.listInitFilesMutex.Lock()
	defer fake.listInitFilesMutex.Unlock()
	fake.ListInitFilesStub = nil
	fake.listInitFilesReturns = struct {
		result1 []tar.Header
	}{result1}
}

func (fake *FakeApkImplementation) ListInitFilesReturnsOnCall(i int, result1 []tar.Header) {
	fake.listInitFilesMutex.Lock()
	defer fake.listInitFilesMutex.Unlock()
	fake.ListInitFilesStub = nil
	if fake.listInitFilesReturnsOnCall == nil {
		fake.listInitFilesReturnsOnCall = make(map[int]struct {
			result1 []tar.Header
		})
	}
	fake.listInitFilesReturnsOnCall[i] = struct {
		result1 []tar.Header
	}{result1}
}

func (fake *FakeApkImplementation) SetRepositories(arg1 []string) error {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setRepositoriesMutex.Lock()
	ret, specificReturn := fake.setRepositoriesReturnsOnCall[len(fake.setRepositoriesArgsForCall)]
	fake.setRepositoriesArgsForCall = append(fake.setRepositoriesArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.SetRepositoriesStub
	fakeReturns := fake.setRepositoriesReturns
	fake.recordInvocation("SetRepositories", []interface{}{arg1Copy})
	fake.setRepositoriesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeApkImplementation) SetRepositoriesCallCount() int {
	fake.setRepositoriesMutex.RLock()
	defer fake.setRepositoriesMutex.RUnlock()
	return len(fake.setRepositoriesArgsForCall)
}

func (fake *FakeApkImplementation) SetRepositoriesCalls(stub func([]string) error) {
	fake.setRepositoriesMutex.Lock()
	defer fake.setRepositoriesMutex.Unlock()
	fake.SetRepositoriesStub = stub
}

func (fake *FakeApkImplementation) SetRepositoriesArgsForCall(i int) []string {
	fake.setRepositoriesMutex.RLock()
	defer fake.setRepositoriesMutex.RUnlock()
	argsForCall := fake.setRepositoriesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeApkImplementation) SetRepositoriesReturns(result1 error) {
	fake.setRepositoriesMutex.Lock()
	defer fake.setRepositoriesMutex.Unlock()
	fake.SetRepositoriesStub = nil
	fake.setRepositoriesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) SetRepositoriesReturnsOnCall(i int, result1 error) {
	fake.setRepositoriesMutex.Lock()
	defer fake.setRepositoriesMutex.Unlock()
	fake.SetRepositoriesStub = nil
	if fake.setRepositoriesReturnsOnCall == nil {
		fake.setRepositoriesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setRepositoriesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) SetWorld(arg1 []string) error {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setWorldMutex.Lock()
	ret, specificReturn := fake.setWorldReturnsOnCall[len(fake.setWorldArgsForCall)]
	fake.setWorldArgsForCall = append(fake.setWorldArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.SetWorldStub
	fakeReturns := fake.setWorldReturns
	fake.recordInvocation("SetWorld", []interface{}{arg1Copy})
	fake.setWorldMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeApkImplementation) SetWorldCallCount() int {
	fake.setWorldMutex.RLock()
	defer fake.setWorldMutex.RUnlock()
	return len(fake.setWorldArgsForCall)
}

func (fake *FakeApkImplementation) SetWorldCalls(stub func([]string) error) {
	fake.setWorldMutex.Lock()
	defer fake.setWorldMutex.Unlock()
	fake.SetWorldStub = stub
}

func (fake *FakeApkImplementation) SetWorldArgsForCall(i int) []string {
	fake.setWorldMutex.RLock()
	defer fake.setWorldMutex.RUnlock()
	argsForCall := fake.setWorldArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeApkImplementation) SetWorldReturns(result1 error) {
	fake.setWorldMutex.Lock()
	defer fake.setWorldMutex.Unlock()
	fake.SetWorldStub = nil
	fake.setWorldReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) SetWorldReturnsOnCall(i int, result1 error) {
	fake.setWorldMutex.Lock()
	defer fake.setWorldMutex.Unlock()
	fake.SetWorldStub = nil
	if fake.setWorldReturnsOnCall == nil {
		fake.setWorldReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setWorldReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeApkImplementation) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fixateWorldMutex.RLock()
	defer fake.fixateWorldMutex.RUnlock()
	fake.getInstalledMutex.RLock()
	defer fake.getInstalledMutex.RUnlock()
	fake.getRepositoriesMutex.RLock()
	defer fake.getRepositoriesMutex.RUnlock()
	fake.getWorldMutex.RLock()
	defer fake.getWorldMutex.RUnlock()
	fake.initDBMutex.RLock()
	defer fake.initDBMutex.RUnlock()
	fake.initKeyringMutex.RLock()
	defer fake.initKeyringMutex.RUnlock()
	fake.listInitFilesMutex.RLock()
	defer fake.listInitFilesMutex.RUnlock()
	fake.setRepositoriesMutex.RLock()
	defer fake.setRepositoriesMutex.RUnlock()
	fake.setWorldMutex.RLock()
	defer fake.setWorldMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeApkImplementation) recordInvocation(key string, args []interface{}) {
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
