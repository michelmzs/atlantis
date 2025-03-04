// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events/vcs (interfaces: GithubPullRequestGetter)

package mocks

import (
	github "github.com/google/go-github/v52/github"
	pegomock "github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockGithubPullRequestGetter struct {
	fail func(message string, callerSkip ...int)
}

func NewMockGithubPullRequestGetter(options ...pegomock.Option) *MockGithubPullRequestGetter {
	mock := &MockGithubPullRequestGetter{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockGithubPullRequestGetter) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockGithubPullRequestGetter) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockGithubPullRequestGetter) GetPullRequest(_param0 models.Repo, _param1 int) (*github.PullRequest, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockGithubPullRequestGetter().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetPullRequest", params, []reflect.Type{reflect.TypeOf((**github.PullRequest)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *github.PullRequest
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*github.PullRequest)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockGithubPullRequestGetter) VerifyWasCalledOnce() *VerifierMockGithubPullRequestGetter {
	return &VerifierMockGithubPullRequestGetter{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockGithubPullRequestGetter) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockGithubPullRequestGetter {
	return &VerifierMockGithubPullRequestGetter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockGithubPullRequestGetter) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockGithubPullRequestGetter {
	return &VerifierMockGithubPullRequestGetter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockGithubPullRequestGetter) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockGithubPullRequestGetter {
	return &VerifierMockGithubPullRequestGetter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockGithubPullRequestGetter struct {
	mock                   *MockGithubPullRequestGetter
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockGithubPullRequestGetter) GetPullRequest(_param0 models.Repo, _param1 int) *MockGithubPullRequestGetter_GetPullRequest_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetPullRequest", params, verifier.timeout)
	return &MockGithubPullRequestGetter_GetPullRequest_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockGithubPullRequestGetter_GetPullRequest_OngoingVerification struct {
	mock              *MockGithubPullRequestGetter
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockGithubPullRequestGetter_GetPullRequest_OngoingVerification) GetCapturedArguments() (models.Repo, int) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockGithubPullRequestGetter_GetPullRequest_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
	}
	return
}
