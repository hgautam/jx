// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"

	github "github.com/google/go-github/v32/github"
	"github.com/petergtz/pegomock"
)

func AnyPtrToGithubResponse() *github.Response {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*github.Response))(nil)).Elem()))
	var nullValue *github.Response
	return nullValue
}

func EqPtrToGithubResponse(value *github.Response) *github.Response {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *github.Response
	return nullValue
}
