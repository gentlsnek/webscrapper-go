package functions

import (
	"fmt"
	"net/http"
	//"strings"
	//"time"
	//"github.com/cilium/ebpf/link"
)

func ReadWebsite(sites []string) {

	c := make(chan string)

	for _, list := range sites {
		go checkLink("http://"+list, c)
	}

}

func checkLink(link string, c chan string) {
	status, err := http.Get(link)

	if err != nil {
		fmt.Println("could not establish connection to", err)
		c <- link
		return
	}
	fmt.Println(status)
	c <- link
}
