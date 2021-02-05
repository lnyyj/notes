package tetcd

import (
	"log"
	"testing"

	srvregister "github.com/lnyyj/utils/service/register"
)

func Test_SrvReg(t *testing.T) {
	{
		ser, err := srvregister.New(`{"endpoints":["localhost:2379"]}`)
		if err != nil {
			log.Fatalln(err)
		}
		defer ser.Close()
	}
}
