// Copyright 2015,2016,2017,2018,2019,2020,2021 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
	"time"

	"github.com/kasworld/log/genlog/basiclog"
	"github.com/kasworld/signalhandlewin"
)

// var Ver = ""

// func init() {
// 	version.Set(Ver)
// }

func main() {
	basiclog.Debug("Start Main")
	defer func() {
		basiclog.Debug("End Main")
	}()
	signalhandlewin.AddArgs()
	flag.Parse()
	srv := &SignalTestServer{}
	if err := signalhandlewin.StartByArgs(srv); err != nil {
		basiclog.Error("%v", err)
	}
	// signalhandlewin.RunWithSignalHandle(srv, basiclog.GlobalLogger)
}

type SignalTestServer struct {
}

// called from signal handler
func (ws *SignalTestServer) GetServiceLockFilename() string {
	filename, err := filepath.Abs("signalhandletest.pid")
	if err != nil {
		fmt.Printf("abs path fail %v", err)
	}
	return filename
}

func (sts *SignalTestServer) GetLogger() interface{} {
	return basiclog.GlobalLogger
}

// called from signal handler
func (sts *SignalTestServer) ServiceInit() error {
	basiclog.Debug("Start ServiceInit %v", sts)
	defer func() {
		basiclog.Debug("End ServiceInit %v", sts)
	}()
	return nil
}

// called from signal handler
func (sts *SignalTestServer) ServiceMain(ctx context.Context) {
	basiclog.Debug("Start ServiceMain %v", sts)
	defer func() {
		basiclog.Debug("End ServiceMain %v", sts)
	}()
	ch1sec := time.After(1000 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ch1sec:
			return
		}
	}
}

// called from signal handler
func (sts *SignalTestServer) ServiceCleanup() {
	basiclog.Debug("Start ServiceCleanup %v", sts)
	defer func() {
		basiclog.Debug("End ServiceCleanup %v", sts)
	}()
}
