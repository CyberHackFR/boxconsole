// Copyright Jason Ertel (github.com/jertel).
// Copyright Security Onion Solutions LLC and/or licensed to Security Onion Solutions LLC under one
// or more contributor license agreements. Licensed under the Elastic License 2.0 as shown at
// https://securityonion.net/license; you may not use this file except in compliance with the
// Elastic License 2.0.

package sostatus

import (
	"testing"

	"github.com/cyberhackfr/boxconsole/server"
	"github.com/stretchr/testify/assert"
)

func TestSoStatusInit(tester *testing.T) {
	status := NewSoStatus(server.NewFakeUnauthorizedServer())
	cfg := make(map[string]interface{})
	cfg["refreshIntervalMs"] = float64(1000)
	cfg["offlineThresholdMs"] = float64(2000)
	err := status.Init(cfg)
	if assert.Nil(tester, err) {
		assert.Equal(tester, 1000, status.refreshIntervalMs)
		assert.Equal(tester, 2000, status.offlineThresholdMs)
	}
}
