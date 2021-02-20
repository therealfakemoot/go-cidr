package cidr

import (
	"encoding/binary"
	"net"
)

func Range(r string) ([]net.IP, error) {
	var ns []net.IP
	ip, n, err := net.ParseCIDR("192.145.34.123/16")
	if err != nil {
		return nil, err
	}
	ipInt := binary.BigEndian.Uint32(ip.To4())
	mask := binary.BigEndian.Uint32(n.Mask)

	// start the ++ here
	// when do we stop?
	x := ipInt & mask
	ipBytes := make([]byte, 4)
	for i := x; i > 0; i++ {
		binary.BigEndian.PutUint32(ipBytes, i)
		ns = append(ns, net.IP(ipBytes))
	}

	return ns, nil
}
