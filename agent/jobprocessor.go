// Copyright Jason Ertel (github.com/jertel).
// Copyright Security Onion Solutions LLC and/or licensed to Security Onion Solutions LLC under one
// or more contributor license agreements. Licensed under the Elastic License 2.0 as shown at
// https://securityonion.net/license; you may not use this file except in compliance with the
// Elastic License 2.0.

package agent

import (
	"io"
	"time"

	"github.com/CyberHackFR/boxconsole/model"
)

type JobProcessor interface {
	ProcessJob(*model.Job, io.ReadCloser) (io.ReadCloser, error)
	CleanupJob(*model.Job)
	GetDataEpoch() time.Time
}
