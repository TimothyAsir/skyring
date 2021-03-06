/*Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package util

import (
	"flag"
	"github.com/golang/glog"
	"skyring/conf"
	"strconv"
	"time"
)

//var logFlushFreq = pflag.Duration("log_flush_frequency", 5*time.Second, "Maximum number of seconds between log flushes")

// InitLogs initializes logs the way we want for SkyRing.
func InitLogs(logConf conf.SkyringLogging) {
	flag.Parse()
	flag.Set("logtostderr", strconv.FormatBool(logConf.Logtostderr))
	flag.Set("log_dir", logConf.Log_dir)
	flag.Set("v", strconv.Itoa(logConf.V))
	// The default glog flush interval is 30 seconds, which is frighteningly long.
	go Forever(glog.Flush, 5*time.Second)
}
