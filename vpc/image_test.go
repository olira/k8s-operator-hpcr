// Copyright 2023 IBM Corp.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.package datasource

package vpc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindImageService(t *testing.T) {
	env, err := envFromDotEnv()
	if err != nil {
		t.Skipf("No .env file")
	}

	service, err := CreateVpcServiceFromEnv(env)
	require.NoError(t, err)

	img, err := FindStockImages(service)
	require.NoError(t, err)

	fmt.Println(img)
}
