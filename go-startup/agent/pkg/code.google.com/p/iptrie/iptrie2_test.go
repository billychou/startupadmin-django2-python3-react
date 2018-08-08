/* iptrie2_test.go - test case for new added functions	*/
/*
modification history
--------------------
2015/6/4, by Zhang Miao, create
*/
package iptrie

import (
	"testing"
)

func TestCidrToRange2(t *testing.T) {
	s, e := cidrToRange2("172.111.68.0/29")
	if s != nil && e != nil {
		if s.String() != "172.111.68.0" || e.String() != "172.111.68.7" {
			t.Errorf("should be (172.111.68.0, 172.111.68.7), now (%s, %s)\n", 
					s.String(), e.String())
		}
	} else {
		t.Errorf("1: return nil\n")
	}

	s, e = cidrToRange2("172.111.68.0/30")
	if s != nil && e != nil {
		if s.String() != "172.111.68.0" || e.String() != "172.111.68.3" {
			t.Errorf("should be (172.111.68.0, 172.111.68.3), now (%s, %s)\n", 
					s.String(), e.String())
		}
	} else {
		t.Errorf("2: return nil\n")
	}

	s, e = cidrToRange2("172.111.68.0/31")
	if s != nil && e != nil {
		if s.String() != "172.111.68.0" || e.String() != "172.111.68.1" {
			t.Errorf("should be (172.111.68.0, 172.111.68.1), now (%s, %s)\n", 
					s.String(), e.String())
		}
	} else {
		t.Errorf("3: return nil\n")
	}

	s, e = cidrToRange2("172.111.68.2/32")
	if s != nil && e != nil {
		if s.String() != "172.111.68.2" || e.String() != "172.111.68.2" {
			t.Errorf("should be (172.111.68.2, 172.111.68.2), now (%s, %s)\n", 
					s.String(), e.String())
		}
	} else {
		t.Errorf("4: return nil\n")
	}
}

func TestAddCIDRRange2(t *testing.T) {
	tt := NewIPTrie()
	da := &testData{10}	
	tt.AddCIDRRange2("192.168.42.0/24", da)
	
	tr := tt.Get("192.168.42.0")
	if tr == nil {
		t.Error("fail to get 192.168.42.0")
	}

	tr = tt.Get("192.168.42.255")
	if tr == nil {
		t.Error("fail to get 192.168.42.255")
	}

	tr = tt.Get("192.168.42.10")
	if tr == nil {
		t.Error("fail to get 192.168.42.10")
	}
}
