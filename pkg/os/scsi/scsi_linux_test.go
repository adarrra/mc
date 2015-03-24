/*
 * Minimalist Object Storage, (C) 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// !build linux,amd64

package scsi

import (
	. "gopkg.in/check.v1"
	"testing"
)

type MySuite struct{}

var _ = Suite(&MySuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *MySuite) TestSCSI(c *C) {
	var d Devices
	err := d.Get()
	c.Assert(err, IsNil)
	c.Assert(len(d.List), Equals, 1)

	c.Assert(len(d.List[0].Scsiattrmap), Not(Equals), 0)
	c.Assert(len(d.List[0].Diskattrmap), Not(Equals), 0)
}