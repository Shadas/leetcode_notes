package _93_Restore_IP_Addresses

import "testing"

func TestRestoreIpAddresses(t *testing.T) {
	var (
		ret []string
	)
	ret = restoreIpAddresses("25525511135")
	t.Logf("ret is %v", ret) // ["255.255.11.135","255.255.111.35"]
	ret = restoreIpAddresses("0000")
	t.Logf("ret is %v", ret) // ["0.0.0.0"]
}
