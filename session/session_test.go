package session_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/internal"
	"github.com/ory/kratos/session"
)

func TestSession(t *testing.T) {
	conf, _ := internal.NewFastRegistryWithMocks(t)
	authAt := time.Now()

	i := &identity.Identity{}
	s := session.NewActiveSession(i, conf, authAt)
	assert.True(t, s.IsActive())

	assert.False(t, (&session.Session{ExpiresAt: time.Now().Add(time.Hour), Identity: i}).IsActive())
	assert.False(t, (&session.Session{Active: true, Identity: i}).IsActive())

	i.Disable()
	assert.False(t, (&session.Session{Active: true, ExpiresAt: time.Now().Add(time.Hour), Identity: i}).IsActive())
}
