package plugin_allowpath

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
)

// Config holds the plugin configuration.
type Config struct {
	Regex []string `json:"regex,omitempty"`
}

func CreateConfig() *Config {
	return &Config{}
}

type allowPath struct {
	name    string
	next    http.Handler
	regexps []*regexp.Regexp
}

// New creates and returns a plugin instance.
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	regexps := make([]*regexp.Regexp, len(config.Regex))

	for i, regex := range config.Regex {
		re, err := regexp.Compile(regex)
		if err != nil {
			return nil, fmt.Errorf("error compiling regex %q: %w", regex, err)
		}

		regexps[i] = re
	}

	return &allowPath{
		name:    name,
		next:    next,
		regexps: regexps,
	}, nil
}

func (b *allowPath) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	currentPath := req.URL.EscapedPath()

	for _, re := range b.regexps {
		if re.MatchString(currentPath) {
			b.next.ServeHTTP(rw, req)
			return
		}
	}

	rw.WriteHeader(http.StatusForbidden)
}
