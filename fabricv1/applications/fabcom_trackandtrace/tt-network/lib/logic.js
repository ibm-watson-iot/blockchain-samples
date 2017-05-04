/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * CRUD create asset
 * @param {com.ibm.watson-iot.samples.tracktrace.CreateSkit} skit - the surgicalkit
 * @transaction
 */
function createSkit(skit) {
    // test for existence
    var skitRegistry = getAssetRegistry('com.ibm.watson-iot.samples.tracktrace.SurgicalKit');
    var skitPrev = skitRegistry.
    return getAssetRegistry('com.ibm.watson-iot.samples.tracktrace.SurgicalKit')
        .then(function(skitRegistry) {
            // save the surgical kit
            return skitRegistry.update(skit);
        });
}