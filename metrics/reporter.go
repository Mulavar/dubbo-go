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

package metrics

import (
	"context"
	"time"
)

import (
	"dubbo.apache.org/dubbogo/protocol"
)

const (
	NameSpace = "dubbo"
)

// Reporter will be used to report the invocation's duration
type Reporter interface {
	// report the duration of an invocation
	Report(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation,
		cost time.Duration, res protocol.Result)
}
