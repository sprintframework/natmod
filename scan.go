/*
 * Copyright (c) 2023 Zander Schwid & Co. LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package natmod

import (
	"github.com/codeallergy/glue"
	"github.com/sprintframework/nat"
)

type natScanner struct {
	Scan     []interface{}
}

func Scanner(scan... interface{}) glue.Scanner {
	return &natScanner{
		Scan: scan,
	}
}

func (t *natScanner) Beans() []interface{} {

	beans := []interface{}{
		NatServiceFactory(),
		&struct {
			NatService []nat.NatService `inject`
		}{},
	}

	return append(beans, t.Scan...)
}

