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

package feed

import (
	"net/url"
	"strings"

	"github.com/vogo/logger"
)

func FormatURL(link string) (string, bool) {
	if isURLContainsInvalidChars(link) {
		return "", false
	}

	u, err := url.Parse(link)
	if err != nil {
		logger.Info("format url error!", err)

		return "", false
	}

	params := u.Query()
	for key := range params {
		if strings.HasPrefix(key, "utm_") {
			params.Del(key)
		}
	}

	u.RawQuery = params.Encode()

	return u.String(), true
}

func isURLContainsInvalidChars(link string) bool {
	return strings.Contains(link, "%22") ||
		strings.Contains(link, "%20") ||
		strings.Contains(link, "%3C")
}