/*******************************************************************************
Copyright (c) 2016 IBM Corporation.


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.


Contributors:

Harrison Kurtz - Initial Contribution
Keerthi Challabotla - Initial Contribution

******************************************************************************/
//SN: March 2016



var log4js = require("log4js");

/**
 * Creates a logger which keeps log entries in a circular,
 * in-memory buffer.  This is useful for performance-critical
 * or load tests, where you don't want to generate lots of I/O.
 * @param {Number} number of entries to keep in the history
 * @return {Object} a new memory logger object
 */
function createLogger(logToMemoryOnly, numEntries, loggerName, config) {
    var that = {};
    var size = numEntries || 20;
    var history = new Array(size);
    var index = 0;
    
    if(config != undefined){
    	 log4js.configure(config);
    }
   
    var log4jsLogger = log4js.getLogger(loggerName);
    
    
    that.logger = log4jsLogger;
    var memoryOnly = false;
	var formatEntry = function(entry) {
		var result = "";
		if (entry) {
			result += entry.time.toISOString() + " " + entry.msg + "\n";
		}
		return result;
	};
	
    var log = function(msg) {
        history[index] = {time: new Date(), msg: msg};
        index++;
        if (index === history.length) {
            index = 0;
        }
    }; 
    
    if (logToMemoryOnly !== undefined) {
    	memoryOnly = logToMemoryOnly;
    }
    that.debug = function(msg) {
    	log(msg);
        if (!memoryOnly) {
        	log4jsLogger.debug(msg);
        }
    };
	
    that.info = function(msg) {
    	log(msg);
        if (!memoryOnly) {
        	log4jsLogger.info(msg);
        }
    };

    that.warn = function(msg) {
    	log(msg);
        if (!memoryOnly) {
        	log4jsLogger.warn(msg);
        }
    };
    
    that.error = function(msg) {
    	log(msg);
        if (!memoryOnly) {
        	log4jsLogger.error(msg);
        }
    };
	
    that.toString = function() {
        var result = "";
        for (var i=index; i<history.length; i++) {
			result += formatEntry(history[i]);
        }
        for (var j=0; j<index; j++) {
			result += formatEntry(history[j]);
        }
		return result;
    };
    
    that.dump = function() {
    	if (memoryOnly) {
    		log4jsLogger.debug(that.toString());
    	}
    };
	
    return that;
};

exports.createLogger = createLogger;