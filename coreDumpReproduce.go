package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"os"

	"github.com/containernetworking/plugins/pkg/ns"
)

func main() {
	master := "ib0"
	linkName := "ipoibtest0"
	netnesPath := "/var/run/netns/testing"

	netns, err := ns.GetNS(netnesPath)
	if err != nil {
		fmt.Printf("Error: failed to open netns %q: %v", netns, err)
	}
	defer netns.Close()

	m, err := netlink.LinkByName(master)
	if err != nil {
		fmt.Printf("Error: failed to lookup master %q: %v", master, err)
		os.Exit(1)
	}
	ipoibLink := &netlink.IPoIB{
		LinkAttrs: netlink.LinkAttrs{
			Name:        linkName,
			ParentIndex: m.Attrs().Index,
			Namespace:   netlink.NsFd(int(netns.Fd())),
		},
		Pkey:   0x7fff,
		Mode:   netlink.IPOIB_MODE_DATAGRAM,
		Umcast: 1,
	}

	if err := netlink.LinkAdd(ipoibLink); err != nil {
		fmt.Printf("Error: failed to create interface: %v", err)
		os.Exit(1)
	}
	fmt.Println("Passed, link", linkName, " created successfully")
}
