package Http

import (
	"fmt"
	_ "net"
)

type HttpWorker struct {
}

func (this *HttpWorker) Start() {
	fmt.Println("IdCreator开启http服务")
}

func (this *HttpWorker) Stop() {

}
