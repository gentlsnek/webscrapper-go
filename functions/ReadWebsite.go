package functions

import (
	"fmt"
	"net/http"
	//"strings"
	//"time"
	//"github.com/cilium/ebpf/link"
)

// no outputs for some reason, need fix
func ReadWebsite(sites []string) {

	c := make(chan string)

	for _, list := range sites {
		go checkLink("http://www."+list, c)
	}

}

func checkLink(link string, c chan string) {
	status, err := http.Get(link)

	if err != nil {
		fmt.Println("could not establish connection to", err)
		c <- link
		return
	} else {
		fmt.Println(status)
		fmt.Println(link, " works fune")
		c <- link
	}
}
