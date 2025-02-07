package viewproxy

import (
	"strings"
)

type Route struct {
	Parts     []string
	Layout    *Fragment
	fragments []*Fragment
}

func newRoute(path string, layout *Fragment, fragments []*Fragment) *Route {
	return &Route{
		Parts:     strings.Split(path, "/"),
		Layout:    layout,
		fragments: fragments,
	}
}

func (r *Route) matchParts(pathParts []string) bool {
	if len(r.Parts) != len(pathParts) {
		return false
	}

	for i := 0; i < len(r.Parts); i++ {
		if r.Parts[i] != pathParts[i] && !strings.HasPrefix(r.Parts[i], ":") {
			return false
		}
	}

	return true
}

func (r *Route) parametersFor(pathParts []string) map[string]string {
	parameters := make(map[string]string)

	for i := 0; i < len(r.Parts); i++ {
		if strings.HasPrefix(r.Parts[i], ":") {
			paramName := r.Parts[i][1:]
			parameters[paramName] = pathParts[i]
		}
	}

	return parameters
}

func (r *Route) FragmentsToRequest() []*Fragment {
	fragments := make([]*Fragment, len(r.fragments)+1)
	fragments[0] = r.Layout

	for i, fragment := range r.fragments {
		fragments[i+1] = fragment
	}
	return fragments
}
