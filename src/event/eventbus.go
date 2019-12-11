package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DataEvent struct {
	Data  interface{}
	Topic string
}

// 接收DataEvent的channel
type DataChannel chan DataEvent

// 包含DataChannels数据的切片
type DataChannelSlice []DataChannel

type EventBus struct {
	// topic -> 订阅者
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

func (eb *EventBus) Subscribe(topic string, ch DataChannel) {
	eb.rm.Lock()
	defer eb.rm.Unlock()
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}

}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.Lock()
	defer eb.rm.Unlock()
	if chans, found := eb.subscribers[topic]; found {
		channels := append(DataChannelSlice{}, chans...)
		go func(data DataEvent, dataChannelSlice DataChannelSlice) {
			for _, ch := range dataChannelSlice {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
}

var eb = &EventBus{
	subscribers: map[string]DataChannelSlice{},
}

func publishTo(topic, data string) {
	for {
		eb.Publish(topic, data)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
func printDataEvent(ch string, data DataEvent) {
	fmt.Printf("Channel: %s; Topic %s; DataEvent: %v\n", ch, data.Topic, data.Data)
}
func main() {
	ch1 := make(chan DataEvent)
	ch2 := make(chan DataEvent)
	ch3 := make(chan DataEvent)
	eb.Subscribe("topic1", ch1)
	eb.Subscribe("topic2", ch2)
	eb.Subscribe("topic2", ch3)
	go publishTo("topic1", "hi topic1")
	go publishTo("topic2", "hi topic2")

	for {
		select {
		case d := <-ch1:
			go printDataEvent("ch1", d)
		case d := <-ch2:
			go printDataEvent("ch2", d)
		case d := <-ch3:
			go printDataEvent("ch3", d)

		}
	}

}
