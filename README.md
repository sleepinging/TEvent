# TEvent

## Example
```go
import (
	"fmt"
	te "./tevent"
	"time"
)

var e =te.NewEvent()

//事件id
const (
	DoorOpen = iota
	DoorClose
)

//开门事件处理
func OnDoorOpen(i int,str string){
	fmt.Println(i,str,"DoorOpen")
}

//开门事件处理2
func OnDoorOpen2(i int,s string){
	fmt.Println(s,i,"Hello")
}

//模拟开门事件
func AutoDoor(){
	for{
		e.Happen(DoorOpen,1,"Red")//发生一个事件
		time.Sleep(time.Second)
	}
}

func main(){
	go AutoDoor()
	e.AddEventHandler(DoorOpen, OnDoorOpen)//添加事件数理函数
	time.Sleep(time.Second*3)
	e.AddEventHandler(DoorOpen, OnDoorOpen2)
	time.Sleep(time.Second*4)
	e.ReMoveEventHandler(DoorOpen, OnDoorOpen)//移除事件处理函数
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

