package tetcd

import (
	"log"
	"testing"
)

func Test_SrvReg(t *testing.T) {
	var endpoints = []string{"localhost:2379"}
	{
		ser, err := srvregister.New(`{}`)
		if err != nil {
			log.Fatalln(err)
		}
		defer ser.Close()
		go ser.ListenLeaseRespChan()
	}

	select {}
}
