package session

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CoockieLifetime string
	CoockiePersist  string
	CoockieName     string
	CoockieDomain   string
	CoockieSecure   string
	SessionType     string
}

func (c *Session) InitSession() *scs.SessionManager {
	var persist, secure bool

	// how long should sessions last?

	minutes, err := strconv.Atoi(c.CoockieLifetime)

	if err != nil {
		minutes = 60
	}

	// should coockies persist?

	if strings.ToLower(c.CoockiePersist) == "true" {
		persist = true
	} else {
		persist = false
	}

	// must coockies be secure?
	if strings.ToLower(c.CoockieSecure) == "true" {
		secure = true
	}

	// create session
	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = c.CoockieName
	session.Cookie.Secure = secure
	session.Cookie.Domain = c.CoockieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	// witch session store?
	switch strings.ToLower(c.SessionType) {
	case "redis":

	case "mysql", "mariadb":

	case "postgres", "postgresql":

	default:
		//coockie
	}

	return session
}
