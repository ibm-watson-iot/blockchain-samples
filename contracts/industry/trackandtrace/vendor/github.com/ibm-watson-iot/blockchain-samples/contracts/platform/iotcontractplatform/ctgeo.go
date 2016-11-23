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

// v0.1 KL -- created to handle geo calculations

package iotcontractplatform

import "math"

// from a tweet by Rob Pike
const x = math.Pi / 180
const rEarth = 6372.8 // radius earth in km

// Rad converts from degrees to radians
func Rad(d float64) float64 { return d * x }

// Deg converts from radians to degrees
func Deg(r float64) float64 { return r / x }

// Distance returns distance between geo coordinates in degrees, result is km
// translated from groovy on rosettacode.com
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	dLat := Rad(lat2 - lat1)
	dLon := Rad(lon2 - lon1)
	lat1 = Rad(lat1)
	lat2 = Rad(lat2)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Asin(math.Sqrt(a))
	return rEarth * c
}
