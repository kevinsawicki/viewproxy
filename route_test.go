package viewproxy

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouteMatch(t *testing.T) {
	tests := map[string]struct {
		routePath   string
		providedUrl string
		want        bool
	}{
		"simple":            {routePath: "/", providedUrl: "/", want: true},
		"simple false":      {routePath: "/", providedUrl: "/hello-world", want: false},
		"multi":             {routePath: "/hello/world", providedUrl: "/hello/world", want: true},
		"multi false":       {routePath: "/hello/world", providedUrl: "/hello/false", want: false},
		"named param":       {routePath: "/hello/:name", providedUrl: "/hello/world", want: true},
		"named param false": {routePath: "/hello/:name", providedUrl: "/hello/world/wow", want: false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			route := newRoute(test.routePath, NewFragment(""), []*Fragment{})
			providedUrlParts := strings.Split(test.providedUrl, "/")
			got := route.matchParts(providedUrlParts)

			if got != test.want {
				t.Fatalf("expected route %s to match URL %s", test.routePath, test.providedUrl)
			}
		})
	}
}

func TestRouteParameters(t *testing.T) {
	tests := map[string]struct {
		routePath   string
		providedUrl string
		want        map[string]string
	}{
		"simple":      {routePath: "/", providedUrl: "/", want: map[string]string{}},
		"multi false": {routePath: "/hello/:name", providedUrl: "/hello/world", want: map[string]string{"name": "world"}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			route := newRoute(test.routePath, NewFragment(""), []*Fragment{})
			providedUrlParts := strings.Split(test.providedUrl, "/")
			got := route.parametersFor(providedUrlParts)

			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("expected route %v with URL %s to have params: %v\n got: %v", test.routePath, test.providedUrl, test.want, got)
			}
		})
	}
}

func TestLayout(t *testing.T) {
	route := newRoute("/", NewFragment("my_layout"), []*Fragment{})

	assert.Equal(t, *route.Layout, Fragment{
		Path:     "my_layout",
		Url:      "",
		Metadata: map[string]string{},
	})
}
