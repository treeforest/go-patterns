package main

import (
	"fmt"
	"time"
)

type (
	// Event 定义了时间点发生的指示
	Event struct {
		// Data 在这个示例中是 int，而实际实现取决于应用程序
		Data int64
	}

	// Observer 观察者，定义了通知事件发生的接口
	Observer interface {
		// OnNotify 允许一个事件被发布到接口实现。在实际应用中，错误的处理可能会被实现。
		OnNotify(Event)
	}

	// Notifier 通知者，是将会被观察(observed)的示例。
	Notifier interface {
		// Register 注册监听的事件。
		Register(Observer)

		// Deregister 移除监听的事件。
		Deregister(Observer)

		// Notify 发布一个新的事件给监听者
		// 该方法不一定成功，因为每个实现都可以在不丢失功能的情况下定义自身。
		Notify(Event)
	}
)

type eventObserver struct {
	id int
}

func (o *eventObserver) OnNotify(e Event) {
	fmt.Printf("*** Observer %d received: %d\n", o.id, e.Data)
}

type eventNotifier struct {
	// 在一个 map 中使用一个空的 struct，可以使 observers 唯一，减少内存的占用。
	observers map[Observer]struct{}
}

func (o *eventNotifier) Register(l Observer) {
	o.observers[l] = struct{}{}
}

func (o *eventNotifier) Deregister(l Observer) {
	delete(o.observers, l)
}

func (o *eventNotifier) Notify(e Event) {
	for observer, _ := range o.observers {
		observer.OnNotify(e)
	}
}

func main() {
	// 初始化一个新的 通知者
	var n Notifier = &eventNotifier{
		observers: map[Observer]struct{}{},
	}

	// 注册一组观察者
	n.Register(&eventObserver{id: 1})
	n.Register(&eventObserver{id: 2})

	e := &eventObserver{id: 3}
	n.Register(e)
	n.Deregister(e)

	// 一个简单的循环将当前的Unix时间戳发布给观察者。
	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C
	for {
		select {
		case <-stop:
			return
		case t := <-tick:
			n.Notify(Event{Data: t.UnixNano()})
		}
	}
}
