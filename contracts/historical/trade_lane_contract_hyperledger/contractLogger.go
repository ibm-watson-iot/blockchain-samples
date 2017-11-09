/*
Copyright (c) 2016 IBM Corporation and other Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.

Contributors:
Kim Letkeman - Initial Contribution
*/


// v1 KL 15 Mar 2016 Created to allow us to insulate from logger versions
//                   and to provide an efficient method of changing levels 
//                   from outside and to allow the level switching to actually work

package main

import (
    "fmt"
    "time"
    //"github.com/op/go-logging"
    "strings"
)

type LogLevel int

const (
	// CRITICAL means cannot function
    CRITICAL LogLevel = iota
    // ERROR means something is wrong
	ERROR
    // WARNING means something might be wrong
	WARNING
    // NOTICE means take note, this should be investigated
	NOTICE
    // INFO means this happened and might be of interest
	INFO
    // DEBUG allows for a peek into the guts of the app for debugging
	DEBUG
)

var logLevelNames = []string {
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
}

// DEFAULTLOGGINGLEVEL is normally INFO in test and WARNING in production
const DEFAULTLOGGINGLEVEL = DEBUG

// ContractLogger is our version of goLogger
type ContractLogger struct {
    module      string
    level       LogLevel
}

// ILogger the goLogger interface to which we are 100% compatible
type ILogger interface {
    Critical(args ...interface{})
    Criticalf(format string, args ...interface{})
    Error(args ...interface{})
    Errorf(format string, args ...interface{})
    Warning(args ...interface{}) 
    Warningf(format string, args ...interface{})
    Notice(args ...interface{})
    Noticef(format string, args ...interface{})
    Info(args ...interface{})
    Infof(format string, args ...interface{})
    Debug(args ...interface{})
    Debugf(format string, args ...interface{})
}

//var goLogger *logging.Logger

// NewContractLogger creates a logger for the contract to use
func NewContractLogger(module string, level LogLevel) (*ContractLogger) {
    l := &ContractLogger{module, level}
    l.SetLoggingLevel(level)
    l.setModule(module)
    return l
}

//SetLoggingLevel is used to change the logging level while the smart contract is running
func (cl *ContractLogger) SetLoggingLevel(level LogLevel) {
    if level < CRITICAL || level > DEBUG {
        cl.level = DEFAULTLOGGINGLEVEL
    } else {
        cl.level = level
    }
}

func (cl *ContractLogger) setModule(module string) {
    if module == "" { module = DEFAULTNICKNAME }
    module += "-" + MYVERSION
    (*cl).module = module
    //goLogger = logging.MustGetLogger(module)
}

//*************
// print logger
//*************

const pf string = "%s [%s] %.4s %s" 

func buildLogString(module string, level LogLevel, msg interface{}) (string) {
    var a = fmt.Sprint(msg)
    var t = time.Now().Format("2006/01/02 15:04:05") 
    return fmt.Sprintf(pf, t, module, logLevelNames[level], a) 
}

// Critical logs a message using CRITICAL as log level.
func (cl *ContractLogger) Critical(msg interface{}) {
    if CRITICAL > cl.level { return }
	logMessage(CRITICAL, buildLogString(cl.module, CRITICAL, msg))
}

// Criticalf logs a message using CRITICAL as log level.
func (cl *ContractLogger) Criticalf(format string, args ...interface{}) {
    if CRITICAL > cl.level { return }
	logMessage(CRITICAL, buildLogString(cl.module, CRITICAL, fmt.Sprintf(format, args)))
}

// Error logs a message using ERROR as log level.
func (cl *ContractLogger) Error(msg interface{}) {
    if ERROR > cl.level { return }
	logMessage(ERROR, buildLogString(cl.module, ERROR, msg))
}

// Errorf logs a message using ERROR as log level.
func (cl *ContractLogger) Errorf(format string, args ...interface{}) {
    if ERROR > cl.level { return }
	logMessage(ERROR, buildLogString(cl.module, ERROR, fmt.Sprintf(format, args)))
}

// Warning logs a message using WARNING as log level.
func (cl *ContractLogger) Warning(msg interface{}) {
    if WARNING > cl.level { return }
	logMessage(WARNING, buildLogString(cl.module, WARNING, msg))
}

// Warningf logs a message using WARNING as log level.
func (cl *ContractLogger) Warningf(format string, args ...interface{}) {
    if WARNING > cl.level { return }
	logMessage(WARNING, buildLogString(cl.module, WARNING, fmt.Sprintf(format, args)))
}

// Notice logs a message using NOTICE as log level.
func (cl *ContractLogger) Notice(msg interface{}) {
    if NOTICE > cl.level { return }
	logMessage(NOTICE, buildLogString(cl.module, NOTICE, msg))
}

// Noticef logs a message using NOTICE as log level.
func (cl *ContractLogger) Noticef(format string, args ...interface{}) {
    if NOTICE > cl.level { return }
	logMessage(NOTICE, buildLogString(cl.module, NOTICE, fmt.Sprintf(format, args)))
}

// Info logs a message using INFO as log level.
func (cl *ContractLogger) Info(msg interface{}) {
    if INFO > cl.level { return }
    logMessage(INFO, buildLogString(cl.module, INFO, msg))
}

// Infof logs a message using INFO as log level.
func (cl *ContractLogger) Infof(format string, args ...interface{}) {
    if INFO > cl.level { return }
	logMessage(INFO, buildLogString(cl.module, INFO, fmt.Sprintf(format, args)))
}

// Debug logs a message using DEBUG as log level.
func (cl *ContractLogger) Debug(msg interface{}) {
    if DEBUG > cl.level { return }
	logMessage(DEBUG, buildLogString(cl.module, DEBUG, msg))
}

// Debugf logs a message using DEBUG as log level.
func (cl *ContractLogger) Debugf(format string, args ...interface{}) {
    if DEBUG > cl.level { return }
	logMessage(DEBUG, buildLogString(cl.module, DEBUG, fmt.Sprintf(format, args)))
}

func logMessage(ll LogLevel, msg string) {
    if !strings.HasSuffix(msg, "\n") {
        msg += "\n"
    }
    fmt.Print(msg)
/*

    removing logger dependency as it is quite literally the only include difference from 3.0.3 to 3.0.4

    // for logger, time is added on front, so delete date and time from
    // our messgage ... space separated
    msg = strings.SplitN(msg, " ", 3)[2] 
    switch ll {
        case CRITICAL :
            goLogger.Critical(msg)            
        case ERROR :
            goLogger.Error(msg)            
        case WARNING :
            goLogger.Warning(msg)            
        case NOTICE :
            goLogger.Notice(msg)            
        case INFO :
            goLogger.Info(msg)            
        case DEBUG :
            goLogger.Debug(msg)            
    }
 */
}
