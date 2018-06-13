package tevent

import "reflect"

type EventFunc interface {}

type TEvent struct {
	EventList map[int][]EventFunc
}

func getArgs(args ...interface{}) []reflect.Value {
	var ma = make([]reflect.Value, len(args))
	for k, v := range args {
		ma[k] = reflect.ValueOf(v)
	}
	return ma
}

func (this *TEvent) Happen(eid int, args ...interface{}) {
	fs := this.EventList[eid]
	if len(fs) == 0 {
		return
	}
	argvs := getArgs(args...)
	for _, f := range fs {
		go reflect.ValueOf(f).Call(argvs)
	}
}

func (this *TEvent) AddEventHandler(eid int, fh EventFunc) {
	this.EventList[eid] = append(this.EventList[eid], fh)
}

func (this *TEvent) ReMoveEventHandler(eid int, fh EventFunc) {
	fs := this.EventList[eid]
	if len(fs) == 0 {
		return
	}
	for i, f := range fs {
		if reflect.ValueOf(fh) == reflect.ValueOf(f) {
			this.EventList[eid] = append(fs[:i], fs[i+1:]...)
			break
		}
	}
}

func (this *TEvent) Init() {
	this.EventList = make(map[int][]EventFunc)
}

func NewEvent() (e *TEvent) {
	e = new(TEvent)
	e.Init()
	return
}
