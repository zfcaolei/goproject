package dangyuantest

import "testing"

func TestAddnum(t *testing.T) {
	res := Addnum(10)
	if res != 55 {
		t.Fatalf("Addnum执行错误")
	}

	//执行正确t
	t.Logf("Addnum执行正确=")
}


//func TestSunNum(t *testing.T) {
//
//	res := SunNum(1,2)
//	if res != 3 {
//		t.Fatalf("SunNum=\n执行错误,期望值=%d,实际值=%d",3,res)
//	}
//	t.Logf("SunNum执行正确")
//}