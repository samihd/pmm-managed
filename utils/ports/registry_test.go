// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package ports

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegistry(t *testing.T) {
	// 10002 is marked as reserved, 10003 is busy, 10004 is free
	r := NewRegistry(10002, 10004, []uint16{10002})
	l1, err := net.Listen("tcp", "127.0.0.1:10003")
	require.NoError(t, err)
	defer l1.Close()

	p, err := r.Reserve()
	assert.NoError(t, err)
	assert.EqualValues(t, 10004, p)
	_, err = r.Reserve()
	assert.Equal(t, errNoFreePort, err)

	l2, err := net.Listen("tcp", "127.0.0.1:10004")
	require.NoError(t, err)
	defer l2.Close()

	err = r.Release(10002)
	assert.NoError(t, err)
	err = r.Release(10003)
	assert.Equal(t, errNotReservedPort, err)
	err = r.Release(10004)
	assert.Equal(t, errPortBusy, err)

	l1.Close()
	l2.Close()

	p, err = r.Reserve()
	assert.NoError(t, err)
	assert.EqualValues(t, 10002, p)
	p, err = r.Reserve()
	assert.NoError(t, err)
	assert.EqualValues(t, 10003, p)
	_, err = r.Reserve()
	assert.Equal(t, errNoFreePort, err)

	err = r.Release(10004)
	assert.NoError(t, err)

	p, err = r.Reserve()
	assert.NoError(t, err)
	assert.EqualValues(t, 10004, p)
	_, err = r.Reserve()
	assert.Equal(t, errNoFreePort, err)
}
