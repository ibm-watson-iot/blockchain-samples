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
	CRITICAL LogLevel = iota
	ERROR
	WARNING
	NOTICE
	INFO
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

const DEFAULTLOGGINGLEVEL = INFO

type ContractLogger struct {
    module      string
    level       LogLevel
}

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

// exported New function
func NewContractLogger(module string, level LogLevel) (*ContractLogger) {
    l := &ContractLogger{module, level}
    l.SetLoggingLevel(level)
    l.SetModule(module)
    return l
}

func (cl *ContractLogger) SetLoggingLevel(level LogLevel) {
    if level < CRITICAL || level > DEBUG {
        cl.level = DEFAULTLOGGINGLEVEL
    } else {
        cl.level = level
    }
}

func (cl *ContractLogger) SetModule(module string) {
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
    var a string = fmt.Sprint(msg)
    var t string = time.Now().Format("2006/01/02 15:04:05") 
    return fmt.Sprintf(pf, t, module, logLevelNames[level], a) 
}

// Critical logs a message using CRITICAL as log level.
func (l *ContractLogger) Critical(msg interface{}) {
    if CRITICAL > l.level { return }
	logMessage(CRITICAL, buildLogString(l.module, CRITICAL, msg))
}

// because older version of go-logger do not support the "f" functions
func (l *ContractLogger) Criticalf(format string, args ...interface{}) {
    if CRITICAL > l.level { return }
	logMessage(CRITICAL, buildLogString(l.module, CRITICAL, fmt.Sprintf(format, args)))
}

// Error logs a message using ERROR as log level.
func (l *ContractLogger) Error(msg interface{}) {
    if ERROR > l.level { return }
	logMessage(ERROR, buildLogString(l.module, ERROR, msg))
}

// Errorf logs a message using ERROR as log level.
func (l *ContractLogger) Errorf(format string, args ...interface{}) {
    if ERROR > l.level { return }
	logMessage(ERROR, buildLogString(l.module, ERROR, fmt.Sprintf(format, args)))
}

// Warning logs a message using WARNING as log level.
func (l *ContractLogger) Warning(msg interface{}) {
    if WARNING > l.level { return }
	logMessage(WARNING, buildLogString(l.module, WARNING, msg))
}

// Warningf logs a message using WARNING as log level.
func (l *ContractLogger) Warningf(format string, args ...interface{}) {
    if WARNING > l.level { return }
	logMessage(WARNING, buildLogString(l.module, WARNING, fmt.Sprintf(format, args)))
}

// Notice logs a message using NOTICE as log level.
func (l *ContractLogger) Notice(msg interface{}) {
    if NOTICE > l.level { return }
	logMessage(NOTICE, buildLogString(l.module, NOTICE, msg))
}

// Noticef logs a message using NOTICE as log level.
func (l *ContractLogger) Noticef(format string, args ...interface{}) {
    if NOTICE > l.level { return }
	logMessage(NOTICE, buildLogString(l.module, NOTICE, fmt.Sprintf(format, args)))
}

// Info logs a message using INFO as log level.
func (l *ContractLogger) Info(msg interface{}) {
    if INFO > l.level { return }
    logMessage(INFO, buildLogString(l.module, INFO, msg))
}

// Infof logs a message using INFO as log level.
func (l *ContractLogger) Infof(format string, args ...interface{}) {
    if INFO > l.level { return }
	logMessage(INFO, buildLogString(l.module, INFO, fmt.Sprintf(format, args)))
}

// Debug logs a message using DEBUG as log level.
func (l *ContractLogger) Debug(msg interface{}) {
    if DEBUG > l.level { return }
	logMessage(DEBUG, buildLogString(l.module, DEBUG, msg))
}

// Debugf logs a message using DEBUG as log level.
func (l *ContractLogger) Debugf(format string, args ...interface{}) {
    if DEBUG > l.level { return }
	logMessage(DEBUG, buildLogString(l.module, DEBUG, fmt.Sprintf(format, args)))
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
