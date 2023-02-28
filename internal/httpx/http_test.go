/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package httpx_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vogo/informer/internal/httpx"
)

func TestGetWechatLinkData(t *testing.T) {
	t.Parallel()

	data, err := httpx.GetWechatLinkData("https://mp.weixin.qq.com/mp/profile_ext?action=home&__biz=MzI1MzYzMjE0MQ==&scene=124#wechat_redirect")
	assert.Nil(t, err)
	t.Logf("data: %s", data)
}
