// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cmd

import (
	"testing"

	"github.com/franela/goblin"
)

// TestLicenseCommand test cases
func TestLicenseCommand(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("LicenseCommand", func() {
		g.It("It should run and return the expected string", func() {
			g.Assert(licenseHandler()).Equal("MIT License, Copyright (c) 2022 Clivern")
		})
	})
}

// BenchmarkLicenseCommand benchmark
func BenchmarkLicenseCommand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		licenseHandler()
	}
}
