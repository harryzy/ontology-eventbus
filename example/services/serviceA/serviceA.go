/****************************************************
Copyright 2018 The ont-eventbus Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*****************************************************/


/***************************************************
Copyright 2016 https://github.com/AsynkronIT/protoactor-go

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*****************************************************/
package serviceA

import (
	"fmt"

	"github.com/ontio/ontology-eventbus/actor"
	. "github.com/ontio/ontology-eventbus/example/services/messages"
)

type ServiceA struct {
}

func (this *ServiceA) Receive(context actor.Context) {
	switch msg := context.Message().(type) {

	case *ServiceARequest:
		fmt.Println("Receive ServiceARequest:", msg.Message)
		context.Sender().Tell(&ServiceAResponse{"I got your message"})

	case *ServiceBResponse:
		fmt.Println("Receive ServiceBResponse:", msg.Message)

	case int:
		context.Sender().Tell(msg + 1)

	default:
		fmt.Printf("unknown message:%v\n", msg)
	}
}
