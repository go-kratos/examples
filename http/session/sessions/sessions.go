package sessions

import (
	"context"
	"encoding/gob"
	"fmt"
	"net/http"
	"time"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
)

func init() {
	gob.Register([]any{})
}

// Default flashes key.
const flashesKey = "_flash"

// Options stores configuration for a session or session store.
//
// Fields are a subset of http.Cookie fields.
type Options struct {
	Path   string
	Domain string
	// MaxAge=0 means no Max-Age attribute specified and the cookie will be
	// deleted after the browser session ends.
	// MaxAge<0 means delete cookie immediately.
	// MaxAge>0 means Max-Age attribute present and given in seconds.
	MaxAge   int
	Secure   bool
	HttpOnly bool
	SameSite http.SameSite
}

// Session stores the values and optional configuration for a session.
type Session struct {
	ID      string
	Values  map[any]any
	Options *Options
	IsNew   bool
	store   Store
	name    string
}

// NewSession is called by session stores to create a new session instance.
func NewSession(store Store, name string) *Session {
	return &Session{
		Values:  make(map[any]any),
		store:   store,
		name:    name,
		Options: new(Options),
	}
}

// Flashes returns a slice of flash messages from the session.
//
// A single variadic argument is accepted, and it is optional: it defines
// the flash key. If not defined "_flash" is used by default.
func (s *Session) Flashes(vars ...string) []any {
	var flashes []any
	key := flashesKey
	if len(vars) > 0 {
		key = vars[0]
	}
	if v, ok := s.Values[key]; ok {
		// Drop the flashes and return it.
		delete(s.Values, key)
		flashes = v.([]any)
	}
	return flashes
}

// AddFlash adds a flash message to the session.
//
// A single variadic argument is accepted, and it is optional: it defines
// the flash key. If not defined "_flash" is used by default.
func (s *Session) AddFlash(value any, vars ...string) {
	key := flashesKey
	if len(vars) > 0 {
		key = vars[0]
	}
	var flashes []any
	if v, ok := s.Values[key]; ok {
		flashes = v.([]any)
	}
	s.Values[key] = append(flashes, value)
}

// Save is a convenience method to save this session. It is the same as calling
// store.Save(request, response, session). You should call Save before writing to
// the response or returning from the handler.
func (s *Session) Save(ht *khttp.Transport) error {
	return s.store.Save(ht, s)
}

// Name returns the name used to register the session.
func (s *Session) Name() string {
	return s.name
}

// Store returns the session store used to register the session.
func (s *Session) Store() Store {
	return s.store
}

// sessionInfo stores a session tracked by the registry.
type sessionInfo struct {
	s *Session
	e error
}

// contextKey is the type used to store the registry in the context.
type contextKey int

// registryKey is the key used to store the registry in the context.
const registryKey contextKey = 0

// Registry stores sessions used during a request.
type Registry struct {
	ht       *khttp.Transport
	sessions map[string]sessionInfo
}

// GetRegistry returns a registry instance for the current request.
func GetRegistry(ht *khttp.Transport) *Registry {
	ctx := ht.Request().Context()
	registry := ctx.Value(registryKey)
	if registry != nil {
		return registry.(*Registry)
	}
	newRegistry := &Registry{
		ht:       ht,
		sessions: make(map[string]sessionInfo),
	}
	*ht.Request() = *ht.Request().WithContext(context.WithValue(ctx, registryKey, newRegistry))
	return newRegistry
}

// Get registers and returns a session for the given name and session store.
//
// It returns a new session if there are no sessions registered for the name.
func (s *Registry) Get(store Store, name string) (session *Session, err error) {
	if !isCookieNameValid(name) {
		return nil, fmt.Errorf("sessions: invalid character in cookie name: %s", name)
	}
	if info, ok := s.sessions[name]; ok {
		session, err = info.s, info.e
	} else {
		session, err = store.New(s.ht, name)
		session.name = name
		s.sessions[name] = sessionInfo{s: session, e: err}
	}
	session.store = store
	return
}

// Save saves all sessions registered for the current request.
func (s *Registry) Save(ht *khttp.Transport) error {
	var errMulti MultiError
	for name, info := range s.sessions {
		session := info.s
		if session.store == nil {
			errMulti = append(errMulti, fmt.Errorf(
				"sessions: missing store for session %q", name))
		} else if err := session.store.Save(ht, session); err != nil {
			errMulti = append(errMulti, fmt.Errorf(
				"sessions: error saving session %q -- %v", name, err))
		}
	}
	if errMulti != nil {
		return errMulti
	}
	return nil
}

// Save saves all sessions used during the current request.
func Save(ht *khttp.Transport) error {
	return GetRegistry(ht).Save(ht)
}

// NewCookie returns an http.Cookie with the options set. It also sets
// the Expires field calculated based on the MaxAge value, for Internet
// Explorer compatibility.
func NewCookie(name, value string, options *Options) *http.Cookie {
	cookie := newCookieFromOptions(name, value, options)
	if options.MaxAge > 0 {
		d := time.Duration(options.MaxAge) * time.Second
		cookie.Expires = time.Now().Add(d)
	} else if options.MaxAge < 0 {
		// Set it to the past to expire now.
		cookie.Expires = time.Unix(1, 0)
	}
	return cookie
}

// MultiError stores multiple errors.
//
// Borrowed from the App Engine SDK.
type MultiError []error

func (m MultiError) Error() string {
	s, n := "", 0
	for _, e := range m {
		if e != nil {
			if n == 0 {
				s = e.Error()
			}
			n++
		}
	}
	switch n {
	case 0:
		return "(0 errors)"
	case 1:
		return s
	case 2:
		return s + " (and 1 other error)"
	}
	return fmt.Sprintf("%s (and %d other errors)", s, n-1)
}
