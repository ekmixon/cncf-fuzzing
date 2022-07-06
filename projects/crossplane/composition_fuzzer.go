// Copyright 2022 ADA Logics Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package composition

import (
	fuzz "github.com/AdaLogics/go-fuzz-headers"
	v1 "github.com/crossplane/crossplane/apis/apiextensions/v1"
)

func FuzzNewCompositionRevision(data []byte) int {
	f := fuzz.NewConsumer(data)
	c := &v1.Composition{}
	f.GenerateStruct(c)
	revision, err := f.GetInt()
	if err != nil {
		return 0
	}
	compSpecHash, err := f.GetString()
	if err != nil {
		return 0
	}

	_ = NewCompositionRevision(c, int64(revision), compSpecHash)
	return 1
}
