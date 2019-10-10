// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeStartActor struct {
	AppNeedsToStageStub        func(v7action.Application) (bool, error)
	appNeedsToStageMutex       sync.RWMutex
	appNeedsToStageArgsForCall []struct {
		arg1 v7action.Application
	}
	appNeedsToStageReturns struct {
		result1 bool
		result2 error
	}
	appNeedsToStageReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	GetApplicationByNameAndSpaceStub        func(string, string) (v7action.Application, v7action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	GetDetailedAppSummaryStub        func(string, string, bool) (v7action.DetailedApplicationSummary, v7action.Warnings, error)
	getDetailedAppSummaryMutex       sync.RWMutex
	getDetailedAppSummaryArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
	}
	getDetailedAppSummaryReturns struct {
		result1 v7action.DetailedApplicationSummary
		result2 v7action.Warnings
		result3 error
	}
	getDetailedAppSummaryReturnsOnCall map[int]struct {
		result1 v7action.DetailedApplicationSummary
		result2 v7action.Warnings
		result3 error
	}
	GetStreamingLogsForApplicationByNameAndSpaceStub        func(string, string, v7action.NOAAClient) (<-chan *v7action.LogMessage, <-chan error, v7action.Warnings, error)
	getStreamingLogsForApplicationByNameAndSpaceMutex       sync.RWMutex
	getStreamingLogsForApplicationByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 v7action.NOAAClient
	}
	getStreamingLogsForApplicationByNameAndSpaceReturns struct {
		result1 <-chan *v7action.LogMessage
		result2 <-chan error
		result3 v7action.Warnings
		result4 error
	}
	getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 <-chan *v7action.LogMessage
		result2 <-chan error
		result3 v7action.Warnings
		result4 error
	}
	PollStartStub        func(string, bool) (v7action.Warnings, error)
	pollStartMutex       sync.RWMutex
	pollStartArgsForCall []struct {
		arg1 string
		arg2 bool
	}
	pollStartReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	pollStartReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	SetApplicationDropletStub        func(string, string) (v7action.Warnings, error)
	setApplicationDropletMutex       sync.RWMutex
	setApplicationDropletArgsForCall []struct {
		arg1 string
		arg2 string
	}
	setApplicationDropletReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	setApplicationDropletReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	StagePackageStub        func(string, string, string) (<-chan v7action.Droplet, <-chan v7action.Warnings, <-chan error)
	stagePackageMutex       sync.RWMutex
	stagePackageArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	stagePackageReturns struct {
		result1 <-chan v7action.Droplet
		result2 <-chan v7action.Warnings
		result3 <-chan error
	}
	stagePackageReturnsOnCall map[int]struct {
		result1 <-chan v7action.Droplet
		result2 <-chan v7action.Warnings
		result3 <-chan error
	}
	StartApplicationStub        func(string) (v7action.Warnings, error)
	startApplicationMutex       sync.RWMutex
	startApplicationArgsForCall []struct {
		arg1 string
	}
	startApplicationReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	startApplicationReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStartActor) AppNeedsToStage(arg1 v7action.Application) (bool, error) {
	fake.appNeedsToStageMutex.Lock()
	ret, specificReturn := fake.appNeedsToStageReturnsOnCall[len(fake.appNeedsToStageArgsForCall)]
	fake.appNeedsToStageArgsForCall = append(fake.appNeedsToStageArgsForCall, struct {
		arg1 v7action.Application
	}{arg1})
	fake.recordInvocation("AppNeedsToStage", []interface{}{arg1})
	fake.appNeedsToStageMutex.Unlock()
	if fake.AppNeedsToStageStub != nil {
		return fake.AppNeedsToStageStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.appNeedsToStageReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStartActor) AppNeedsToStageCallCount() int {
	fake.appNeedsToStageMutex.RLock()
	defer fake.appNeedsToStageMutex.RUnlock()
	return len(fake.appNeedsToStageArgsForCall)
}

func (fake *FakeStartActor) AppNeedsToStageCalls(stub func(v7action.Application) (bool, error)) {
	fake.appNeedsToStageMutex.Lock()
	defer fake.appNeedsToStageMutex.Unlock()
	fake.AppNeedsToStageStub = stub
}

func (fake *FakeStartActor) AppNeedsToStageArgsForCall(i int) v7action.Application {
	fake.appNeedsToStageMutex.RLock()
	defer fake.appNeedsToStageMutex.RUnlock()
	argsForCall := fake.appNeedsToStageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStartActor) AppNeedsToStageReturns(result1 bool, result2 error) {
	fake.appNeedsToStageMutex.Lock()
	defer fake.appNeedsToStageMutex.Unlock()
	fake.AppNeedsToStageStub = nil
	fake.appNeedsToStageReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeStartActor) AppNeedsToStageReturnsOnCall(i int, result1 bool, result2 error) {
	fake.appNeedsToStageMutex.Lock()
	defer fake.appNeedsToStageMutex.Unlock()
	fake.AppNeedsToStageStub = nil
	if fake.appNeedsToStageReturnsOnCall == nil {
		fake.appNeedsToStageReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.appNeedsToStageReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeStartActor) GetApplicationByNameAndSpace(arg1 string, arg2 string) (v7action.Application, v7action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{arg1, arg2})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeStartActor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeStartActor) GetApplicationByNameAndSpaceCalls(stub func(string, string) (v7action.Application, v7action.Warnings, error)) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = stub
}

func (fake *FakeStartActor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getApplicationByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStartActor) GetApplicationByNameAndSpaceReturns(result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStartActor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v7action.Application
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStartActor) GetDetailedAppSummary(arg1 string, arg2 string, arg3 bool) (v7action.DetailedApplicationSummary, v7action.Warnings, error) {
	fake.getDetailedAppSummaryMutex.Lock()
	ret, specificReturn := fake.getDetailedAppSummaryReturnsOnCall[len(fake.getDetailedAppSummaryArgsForCall)]
	fake.getDetailedAppSummaryArgsForCall = append(fake.getDetailedAppSummaryArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetDetailedAppSummary", []interface{}{arg1, arg2, arg3})
	fake.getDetailedAppSummaryMutex.Unlock()
	if fake.GetDetailedAppSummaryStub != nil {
		return fake.GetDetailedAppSummaryStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getDetailedAppSummaryReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeStartActor) GetDetailedAppSummaryCallCount() int {
	fake.getDetailedAppSummaryMutex.RLock()
	defer fake.getDetailedAppSummaryMutex.RUnlock()
	return len(fake.getDetailedAppSummaryArgsForCall)
}

func (fake *FakeStartActor) GetDetailedAppSummaryCalls(stub func(string, string, bool) (v7action.DetailedApplicationSummary, v7action.Warnings, error)) {
	fake.getDetailedAppSummaryMutex.Lock()
	defer fake.getDetailedAppSummaryMutex.Unlock()
	fake.GetDetailedAppSummaryStub = stub
}

func (fake *FakeStartActor) GetDetailedAppSummaryArgsForCall(i int) (string, string, bool) {
	fake.getDetailedAppSummaryMutex.RLock()
	defer fake.getDetailedAppSummaryMutex.RUnlock()
	argsForCall := fake.getDetailedAppSummaryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeStartActor) GetDetailedAppSummaryReturns(result1 v7action.DetailedApplicationSummary, result2 v7action.Warnings, result3 error) {
	fake.getDetailedAppSummaryMutex.Lock()
	defer fake.getDetailedAppSummaryMutex.Unlock()
	fake.GetDetailedAppSummaryStub = nil
	fake.getDetailedAppSummaryReturns = struct {
		result1 v7action.DetailedApplicationSummary
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStartActor) GetDetailedAppSummaryReturnsOnCall(i int, result1 v7action.DetailedApplicationSummary, result2 v7action.Warnings, result3 error) {
	fake.getDetailedAppSummaryMutex.Lock()
	defer fake.getDetailedAppSummaryMutex.Unlock()
	fake.GetDetailedAppSummaryStub = nil
	if fake.getDetailedAppSummaryReturnsOnCall == nil {
		fake.getDetailedAppSummaryReturnsOnCall = make(map[int]struct {
			result1 v7action.DetailedApplicationSummary
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getDetailedAppSummaryReturnsOnCall[i] = struct {
		result1 v7action.DetailedApplicationSummary
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeStartActor) GetStreamingLogsForApplicationByNameAndSpace(arg1 string, arg2 string, arg3 v7action.NOAAClient) (<-chan *v7action.LogMessage, <-chan error, v7action.Warnings, error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)]
	fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall = append(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 v7action.NOAAClient
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetStreamingLogsForApplicationByNameAndSpace", []interface{}{arg1, arg2, arg3})
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetStreamingLogsForApplicationByNameAndSpaceStub != nil {
		return fake.GetStreamingLogsForApplicationByNameAndSpaceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	fakeReturns := fake.getStreamingLogsForApplicationByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4
}

func (fake *FakeStartActor) GetStreamingLogsForApplicationByNameAndSpaceCallCount() int {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeStartActor) GetStreamingLogsForApplicationByNameAndSpaceCalls(stub func(string, string, v7action.NOAAClient) (<-chan *v7action.LogMessage, <-chan error, v7action.Warnings, error)) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = stub
}

func (fake *FakeStartActor) GetStreamingLogsForApplicationByNameAndSpaceArgsForCall(i int) (string, string, v7action.NOAAClient) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeStartActor) GetStreamingLogsForApplicationByNameAndSpaceReturns(result1 <-chan *v7action.LogMessage, result2 <-chan error, result3 v7action.Warnings, result4 error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	fake.getStreamingLogsForApplicationByNameAndSpaceReturns = struct {
		result1 <-chan *v7action.LogMessage
		result2 <-chan error
		result3 v7action.Warnings
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeStartActor) GetStreamingLogsForApplicationByNameAndSpaceReturnsOnCall(i int, result1 <-chan *v7action.LogMessage, result2 <-chan error, result3 v7action.Warnings, result4 error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	if fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 <-chan *v7action.LogMessage
			result2 <-chan error
			result3 v7action.Warnings
			result4 error
		})
	}
	fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 <-chan *v7action.LogMessage
		result2 <-chan error
		result3 v7action.Warnings
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeStartActor) PollStart(arg1 string, arg2 bool) (v7action.Warnings, error) {
	fake.pollStartMutex.Lock()
	ret, specificReturn := fake.pollStartReturnsOnCall[len(fake.pollStartArgsForCall)]
	fake.pollStartArgsForCall = append(fake.pollStartArgsForCall, struct {
		arg1 string
		arg2 bool
	}{arg1, arg2})
	fake.recordInvocation("PollStart", []interface{}{arg1, arg2})
	fake.pollStartMutex.Unlock()
	if fake.PollStartStub != nil {
		return fake.PollStartStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.pollStartReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStartActor) PollStartCallCount() int {
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	return len(fake.pollStartArgsForCall)
}

func (fake *FakeStartActor) PollStartCalls(stub func(string, bool) (v7action.Warnings, error)) {
	fake.pollStartMutex.Lock()
	defer fake.pollStartMutex.Unlock()
	fake.PollStartStub = stub
}

func (fake *FakeStartActor) PollStartArgsForCall(i int) (string, bool) {
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	argsForCall := fake.pollStartArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStartActor) PollStartReturns(result1 v7action.Warnings, result2 error) {
	fake.pollStartMutex.Lock()
	defer fake.pollStartMutex.Unlock()
	fake.PollStartStub = nil
	fake.pollStartReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeStartActor) PollStartReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.pollStartMutex.Lock()
	defer fake.pollStartMutex.Unlock()
	fake.PollStartStub = nil
	if fake.pollStartReturnsOnCall == nil {
		fake.pollStartReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.pollStartReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeStartActor) SetApplicationDroplet(arg1 string, arg2 string) (v7action.Warnings, error) {
	fake.setApplicationDropletMutex.Lock()
	ret, specificReturn := fake.setApplicationDropletReturnsOnCall[len(fake.setApplicationDropletArgsForCall)]
	fake.setApplicationDropletArgsForCall = append(fake.setApplicationDropletArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("SetApplicationDroplet", []interface{}{arg1, arg2})
	fake.setApplicationDropletMutex.Unlock()
	if fake.SetApplicationDropletStub != nil {
		return fake.SetApplicationDropletStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.setApplicationDropletReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStartActor) SetApplicationDropletCallCount() int {
	fake.setApplicationDropletMutex.RLock()
	defer fake.setApplicationDropletMutex.RUnlock()
	return len(fake.setApplicationDropletArgsForCall)
}

func (fake *FakeStartActor) SetApplicationDropletCalls(stub func(string, string) (v7action.Warnings, error)) {
	fake.setApplicationDropletMutex.Lock()
	defer fake.setApplicationDropletMutex.Unlock()
	fake.SetApplicationDropletStub = stub
}

func (fake *FakeStartActor) SetApplicationDropletArgsForCall(i int) (string, string) {
	fake.setApplicationDropletMutex.RLock()
	defer fake.setApplicationDropletMutex.RUnlock()
	argsForCall := fake.setApplicationDropletArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStartActor) SetApplicationDropletReturns(result1 v7action.Warnings, result2 error) {
	fake.setApplicationDropletMutex.Lock()
	defer fake.setApplicationDropletMutex.Unlock()
	fake.SetApplicationDropletStub = nil
	fake.setApplicationDropletReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeStartActor) SetApplicationDropletReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.setApplicationDropletMutex.Lock()
	defer fake.setApplicationDropletMutex.Unlock()
	fake.SetApplicationDropletStub = nil
	if fake.setApplicationDropletReturnsOnCall == nil {
		fake.setApplicationDropletReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.setApplicationDropletReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeStartActor) StagePackage(arg1 string, arg2 string, arg3 string) (<-chan v7action.Droplet, <-chan v7action.Warnings, <-chan error) {
	fake.stagePackageMutex.Lock()
	ret, specificReturn := fake.stagePackageReturnsOnCall[len(fake.stagePackageArgsForCall)]
	fake.stagePackageArgsForCall = append(fake.stagePackageArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("StagePackage", []interface{}{arg1, arg2, arg3})
	fake.stagePackageMutex.Unlock()
	if fake.StagePackageStub != nil {
		return fake.StagePackageStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.stagePackageReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeStartActor) StagePackageCallCount() int {
	fake.stagePackageMutex.RLock()
	defer fake.stagePackageMutex.RUnlock()
	return len(fake.stagePackageArgsForCall)
}

func (fake *FakeStartActor) StagePackageCalls(stub func(string, string, string) (<-chan v7action.Droplet, <-chan v7action.Warnings, <-chan error)) {
	fake.stagePackageMutex.Lock()
	defer fake.stagePackageMutex.Unlock()
	fake.StagePackageStub = stub
}

func (fake *FakeStartActor) StagePackageArgsForCall(i int) (string, string, string) {
	fake.stagePackageMutex.RLock()
	defer fake.stagePackageMutex.RUnlock()
	argsForCall := fake.stagePackageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeStartActor) StagePackageReturns(result1 <-chan v7action.Droplet, result2 <-chan v7action.Warnings, result3 <-chan error) {
	fake.stagePackageMutex.Lock()
	defer fake.stagePackageMutex.Unlock()
	fake.StagePackageStub = nil
	fake.stagePackageReturns = struct {
		result1 <-chan v7action.Droplet
		result2 <-chan v7action.Warnings
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeStartActor) StagePackageReturnsOnCall(i int, result1 <-chan v7action.Droplet, result2 <-chan v7action.Warnings, result3 <-chan error) {
	fake.stagePackageMutex.Lock()
	defer fake.stagePackageMutex.Unlock()
	fake.StagePackageStub = nil
	if fake.stagePackageReturnsOnCall == nil {
		fake.stagePackageReturnsOnCall = make(map[int]struct {
			result1 <-chan v7action.Droplet
			result2 <-chan v7action.Warnings
			result3 <-chan error
		})
	}
	fake.stagePackageReturnsOnCall[i] = struct {
		result1 <-chan v7action.Droplet
		result2 <-chan v7action.Warnings
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeStartActor) StartApplication(arg1 string) (v7action.Warnings, error) {
	fake.startApplicationMutex.Lock()
	ret, specificReturn := fake.startApplicationReturnsOnCall[len(fake.startApplicationArgsForCall)]
	fake.startApplicationArgsForCall = append(fake.startApplicationArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("StartApplication", []interface{}{arg1})
	fake.startApplicationMutex.Unlock()
	if fake.StartApplicationStub != nil {
		return fake.StartApplicationStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.startApplicationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStartActor) StartApplicationCallCount() int {
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	return len(fake.startApplicationArgsForCall)
}

func (fake *FakeStartActor) StartApplicationCalls(stub func(string) (v7action.Warnings, error)) {
	fake.startApplicationMutex.Lock()
	defer fake.startApplicationMutex.Unlock()
	fake.StartApplicationStub = stub
}

func (fake *FakeStartActor) StartApplicationArgsForCall(i int) string {
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	argsForCall := fake.startApplicationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStartActor) StartApplicationReturns(result1 v7action.Warnings, result2 error) {
	fake.startApplicationMutex.Lock()
	defer fake.startApplicationMutex.Unlock()
	fake.StartApplicationStub = nil
	fake.startApplicationReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeStartActor) StartApplicationReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.startApplicationMutex.Lock()
	defer fake.startApplicationMutex.Unlock()
	fake.StartApplicationStub = nil
	if fake.startApplicationReturnsOnCall == nil {
		fake.startApplicationReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.startApplicationReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeStartActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appNeedsToStageMutex.RLock()
	defer fake.appNeedsToStageMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getDetailedAppSummaryMutex.RLock()
	defer fake.getDetailedAppSummaryMutex.RUnlock()
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	fake.setApplicationDropletMutex.RLock()
	defer fake.setApplicationDropletMutex.RUnlock()
	fake.stagePackageMutex.RLock()
	defer fake.stagePackageMutex.RUnlock()
	fake.startApplicationMutex.RLock()
	defer fake.startApplicationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStartActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.StartActor = new(FakeStartActor)
