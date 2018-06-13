# TEvent

## Example
```go
import (
	"fmt"
	te "./tevent"
	"time"
)

var e =te.NewEvent()

const (
	DoorOpen = iota
	DoorClose
)

func OnDoorOpen(i int,str string){
	fmt.Println(i,str,"DoorOpen")
}

func OnDoorOpen2(i int,s string){
	fmt.Println(s,i,"Hello")
}

func AutoDoor(){
	for{
		e.Happen(DoorOpen,1,"Red")
		time.Sleep(time.Second)
	}
}

func main(){
	go AutoDoor()
	e.AddEventHandler(DoorOpen, OnDoorOpen)
	time.Sleep(time.Second*3)
	e.AddEventHandler(DoorOpen, OnDoorOpen2)
	time.Sleep(time.Second*4)
	e.ReMoveEventHandler(DoorOpen, OnDoorOpen)
	time.Sleep(time.Second*3)
}
```
### OUTPUT
1 Red DoorOpen
1 Red DoorOpen
Red 1 Hello
1 Red DoorOpen
1 Red DoorOpen
Red 1 Hello
1 Red DoorOpen
Red 1 Hello
1 Red DoorOpen
Red 1 Hello
Red 1 Hello
Red 1 Hello
Red 1 Hello
