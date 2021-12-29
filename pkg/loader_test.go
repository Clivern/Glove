// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package pkg

import (
	"fmt"
	"testing"

	"github.com/franela/goblin"
	"github.com/spf13/viper"
)

// TestLoadConfigs test cases
func TestLoadConfigs(t *testing.T) {
	g := goblin.Goblin(t)

	LoadConfigs(fmt.Sprintf("%s/config.dist.yml", GetBaseDir("cache")))

	g.Describe("LoadConfigs", func() {
		g.It("It should run and return the expected string", func() {
			g.Assert(viper.GetString("app.name")).Equal("glove")
		})
	})

	g.Describe("GetBaseDir", func() {
		g.It("It should run and return the expected string", func() {
			g.Assert(GetBaseDir("cache") != "").Equal(true)
		})
	})
}
