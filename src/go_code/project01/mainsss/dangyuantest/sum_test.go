package dangyuantest

import "testing"

func TestSunNum(t *testing.T) {

	 res := SunNum(1,2)
	if res != 3 {
		t.Fatalf("SunNum=\n执行错误,期望值=%d,实际值=%d",3,res)
	}
	t.Logf("SunNum执行正确")
}