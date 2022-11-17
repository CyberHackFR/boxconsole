// Copyright Jason Ertel (github.com/jertel).
// Copyright Security Onion Solutions LLC and/or licensed to Security Onion Solutions LLC under one
// or more contributor license agreements. Licensed under the Elastic License 2.0 as shown at
// https://securityonion.net/license; you may not use this file except in compliance with the
// Elastic License 2.0.

package modules

import (
	"github.com/cyberhackfr/boxconsole/agent"
	"github.com/cyberhackfr/boxconsole/agent/modules/analyze"
	"github.com/cyberhackfr/boxconsole/agent/modules/importer"
	"github.com/cyberhackfr/boxconsole/agent/modules/statickeyauth"
	"github.com/cyberhackfr/boxconsole/agent/modules/stenoquery"
	"github.com/cyberhackfr/boxconsole/module"
)

func BuildModuleMap(agt *agent.Agent) map[string]module.Module {
	moduleMap := make(map[string]module.Module)
	moduleMap["analyze"] = analyze.NewAnalyze(agt)
	moduleMap["importer"] = importer.NewImporter(agt)
	moduleMap["statickeyauth"] = statickeyauth.NewStaticKeyAuth(agt)
	moduleMap["stenoquery"] = stenoquery.NewStenoQuery(agt)
	return moduleMap
}
