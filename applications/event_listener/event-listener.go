/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hyperledger/fabric/events/consumer"
	pb "github.com/hyperledger/fabric/protos"
)

type adapter struct {
	notfy     chan *pb.Event_Block
	rejection chan *pb.Event_Rejection
	cc        chan *pb.Event_ChaincodeEvent
}

//GetInterestedEvents implements consumer.EventAdapter interface for registering interested events
func (a *adapter) GetInterestedEvents() ([]*pb.Interest, error) {
	return []*pb.Interest{
		&pb.Interest{
			EventType: pb.EventType_BLOCK,
		},
		&pb.Interest{
			EventType: pb.EventType_REJECTION,
		},
		&pb.Interest{
			EventType: pb.EventType_CHAINCODE,
			RegInfo: &pb.Interest_ChaincodeRegInfo{
				ChaincodeRegInfo: &pb.ChaincodeReg{
					ChaincodeID: "mycc",
					EventName:   "EVTPING",
				},
			},
		},
		&pb.Interest{
			EventType: pb.EventType_CHAINCODE,
			RegInfo: &pb.Interest_ChaincodeRegInfo{
				ChaincodeRegInfo: &pb.ChaincodeReg{
					ChaincodeID: "mycc",
					EventName:   "EVTPONG",
				},
			},
		},
		&pb.Interest{
			EventType: pb.EventType_CHAINCODE,
			RegInfo: &pb.Interest_ChaincodeRegInfo{
				ChaincodeRegInfo: &pb.ChaincodeReg{
					ChaincodeID: "mycc",
					EventName:   "EVTINVOKEERR",
				},
			},
		},
		&pb.Interest{
			EventType: pb.EventType_CHAINCODE,
			RegInfo: &pb.Interest_ChaincodeRegInfo{
				ChaincodeRegInfo: &pb.ChaincodeReg{
					ChaincodeID: "mycc",
					EventName:   "EVT.IOTCP.INVOKE.RESULT",
				},
			},
		},
	}, nil
}

//Recv implements consumer.EventAdapter interface for receiving events
func (a *adapter) Recv(msg *pb.Event) (bool, error) {
	switch msg.Event.(type) {
	case *pb.Event_Block:
		a.notfy <- msg.Event.(*pb.Event_Block)
		return true, nil
	case *pb.Event_Rejection:
		a.rejection <- msg.Event.(*pb.Event_Rejection)
		return true, nil
	case *pb.Event_ChaincodeEvent:
		a.cc <- msg.Event.(*pb.Event_ChaincodeEvent)
		return true, nil
	default:
		fmt.Printf("RECV went through DEFAULT for some reason\n")
		// a.notfy <- nil
		// a.cc <- nil
		return false, nil
	}
}

//Disconnected implements consumer.EventAdapter interface for disconnecting
func (a *adapter) Disconnected(err error) {
	fmt.Printf("Disconnected...exiting\n")
	os.Exit(1)
}

func createEventClient(eventAddress string) *adapter {
	var obcEHClient *consumer.EventsClient

	done := make(chan *pb.Event_Block)
	done2 := make(chan *pb.Event_ChaincodeEvent)
	done3 := make(chan *pb.Event_Rejection)
	adapter := &adapter{notfy: done, cc: done2, rejection: done3}
	obcEHClient, _ = consumer.NewEventsClient(eventAddress, 5, adapter)
	if err := obcEHClient.Start(); err != nil {
		fmt.Printf("could not start chat %s\n", err)
		obcEHClient.Stop()
		return nil
	}

	return adapter
}

func main() {
	var eventAddress string
	flag.StringVar(&eventAddress, "events-address", "0.0.0.0:7053", "address of events server")
	flag.Parse()

	fmt.Printf("Event Address: %s\n", eventAddress)

	a := createEventClient(eventAddress)
	if a == nil {
		fmt.Printf("Error creating event client\n")
		return
	}
	fmt.Printf("Event client appears to have been succesfully created\n")

	for {
		select {
		case b := <-a.notfy:
			fmt.Printf("\nReceived block\n")
			fmt.Printf("%+v\n\n", b)
		case r := <-a.rejection:
			fmt.Printf("\nReceived rejection\n")
			fmt.Printf("%+v\n\n", r)
		case p := <-a.cc:
			fmt.Printf("\nReceived chaincode event\n")
			fmt.Printf("%+v\n\n", p)
		}
	}
}
