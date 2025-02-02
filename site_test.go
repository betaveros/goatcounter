// Copyright © 2019 Martin Tournoij <martin@arp242.net>
// This file is part of GoatCounter and published under the terms of the AGPLv3,
// which can be found in the LICENSE file or at gnu.org/licenses/agpl.html

package goatcounter

import "testing"

func TestSiteInsert(t *testing.T) {
	ctx, clean := StartTest(t)
	defer clean()

	s := Site{Code: "the_code", Name: "the-code.com", Plan: PlanPersonal}
	err := s.Insert(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if s.ID == 0 {
		t.Fatal("ID is 0")
	}
}
