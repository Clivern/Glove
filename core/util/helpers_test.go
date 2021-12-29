// Copyright 2022 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package util

import (
	"testing"

	"github.com/franela/goblin"
)

// Item type
type Item struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

// TestUnitURL
func TestUnitURL(t *testing.T) {
	g := goblin.Goblin(t)

	item1 := &Item{
		Key1: "value1",
		Key2: "value2",
	}

	item2 := &Item{
		Key1: "value3",
		Key2: "value4",
	}

	item3 := &Item{}

	g.Describe("#TestLoadFromJSON", func() {
		g.It("It should satisfy test cases", func() {
			out, err := ConvertToJSON(item1)

			g.Assert(out).Equal(`{"key1":"value1","key2":"value2"}`)
			g.Assert(err).Equal(nil)
		})
	})

	g.Describe("#TestConvertToJSON", func() {
		g.It("It should satisfy test cases", func() {
			out, err := ConvertToJSON(item2)

			g.Assert(err).Equal(nil)
			LoadFromJSON(item3, []byte(out))

			g.Assert(item3.Key1).Equal("value3")
			g.Assert(item3.Key2).Equal("value4")
		})
	})
}
