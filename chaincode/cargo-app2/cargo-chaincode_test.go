package main

import(
  "fmt"
  "testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkInvoke(t *testing.T,stub *shim.MockStub,args [][]byte){
  res:=stub.MockInvoke("1",args)
  if res.Status !=shim.OK{
    fmt.Println("Invoke", args, "failed", string(res.Message))
    t.FailNow()
  }
}

func checkQuery(t *testing.T,stub *shim.MockStub,args [][]byte,expected string){
  res := stub.MockInvoke("1", args)
  if res.Status != shim.OK {
  fmt.Println("Query", args, "failed", string(res.Message))
  t.FailNow()
}
if res.Payload == nil {
  fmt.Println("Query", args, "failed to get result")
  t.FailNow()
}
if string(res.Payload) != expected {
  fmt.Println("Query result ", string(res.Payload), "was not", expected, "as expected")
  t.FailNow()
}
}
/**
Point 관련 유닛테스트
**/
func TestCargo_point01(t *testing.T){
  expected1:="{\"Username\":\"lim\",\"Total\":1000}"
  expected2:="{\"Username\":\"lim\",\"Total\":1200}"
  expected3:="{\"Username\":\"lim\",\"Total\":200}"
  expected4:="{\"Username\":\"anonymous\",\"Total\":0}"
  cargoc:=new(SmartContract)
  stub:=shim.NewMockStub("cargo",cargoc)
  checkInvoke(t,stub,[][]byte{[]byte("initLedger")})
  checkQuery(t,stub,[][]byte{[]byte("queryPoint"),[]byte("color0e")},expected1)
  checkInvoke(t,stub,[][]byte{[]byte("addPoint"),[]byte("color0e"),[]byte("200")})
  checkQuery(t,stub,[][]byte{[]byte("queryPoint"),[]byte("color0e")},expected2)
  checkInvoke(t,stub,[][]byte{[]byte("subtractPoint"),[]byte("color0e"),[]byte("1000")})
  checkQuery(t,stub,[][]byte{[]byte("queryPoint"),[]byte("color0e")},expected3)
  checkInvoke(t,stub,[][]byte{[]byte("createUser"),[]byte("newUser"),[]byte("anonymous")})
  checkQuery(t,stub,[][]byte{[]byte("queryPoint"),[]byte("newUser")},expected4)
}
/**
계약관련 유닛테스트 01
**/
func TestCargo_contract01(t *testing.T){
  expected1:=`[{"Key":"CARGO20180606_1", "Record":{"weight":3,"distance":5,"money":5,"date":"2018-06-06","Registrant":"you","Driver":"me","Recipient":"her","Status":"Yet"}},{"Key":"CARGO20180606_2", "Record":{"weight":3,"distance":5,"money":5,"date":"2018-06-06","Registrant":"you","Driver":"me","Recipient":"her","Status":"Fail"}},{"Key":"CARGO20180913_1", "Record":{"weight":3,"distance":5,"money":5,"date":"2018-09-13","Registrant":"you","Driver":"me","Recipient":"him","Status":"Success"}}]`

  cargoc:=new(SmartContract)
  stub:=shim.NewMockStub("cargo",cargoc)
  checkInvoke(t,stub,[][]byte{[]byte("initLedger")})
  checkQuery(t,stub,[][]byte{[]byte("queryCargo"),[]byte("20180606")},expected1)

}

/**
계약관련 유닛테스트 02
**/

func TestCargo_contract02(t *testing.T){
//cargo1 := CargoContract{Weight: 3, Distance: 5.0, Money: 5, Date: dateTestValue2, Registrant : "you", Driver : "me", Recipient : "her" , Status : YET}
expected1:=`[{"Key":"CARGO20180909_1", "Record":{"weight":3,"distance":5,"money":10000,"date":"2018-09-09","Registrant":"redtree0","Driver":"color0e","Recipient":"newUser","Status":"Yet"}}]`
cargoc:=new(SmartContract)
stub:=shim.NewMockStub("cargo",cargoc)
checkInvoke(t,stub,[][]byte{[]byte("createUser"),[]byte("newUser"),[]byte("anonymous")})
checkInvoke(t,stub,[][]byte{[]byte("createUser"),[]byte("color0e"),[]byte("lim")})
checkInvoke(t,stub,[][]byte{[]byte("createUser"),[]byte("redtree0"),[]byte("kim")})
checkInvoke(t,stub,[][]byte{[]byte("addPoint"),[]byte("newUser"),[]byte("1000")})
checkInvoke(t,stub,[][]byte{[]byte("addPoint"),[]byte("color0e"),[]byte("1000")})
checkInvoke(t,stub,[][]byte{[]byte("addPoint"),[]byte("redtree0"),[]byte("50000")})
checkInvoke(t,stub,[][]byte{[]byte("createContract"),[]byte("20180909"),[]byte("3"),[]byte("5.0"),
  []byte("10000"),[]byte("2018-09-09"),[]byte("redtree0"),[]byte("color0e"),[]byte("newUser"),[]byte("Yet")})
checkQuery(t,stub,[][]byte{[]byte("queryCargo"),[]byte("20180909")},expected1)
}

func TestCargo_contract03(t *testing.T){
  cargoc:=new(SmartContract)
  stub:=shim.NewMockStub("cargo",cargoc)
  expected1:="{\"Username\":\"kim\",\"Total\":40000}"
  expected2:="{\"Username\":\"lim\",\"Total\":11000}"
  checkInvoke(t,stub,[][]byte{[]byte("createUser"),[]byte("newUser"),[]byte("anonymous")})
  checkInvoke(t,stub,[][]byte{[]byte("createUser"),[]byte("color0e"),[]byte("lim")})
  checkInvoke(t,stub,[][]byte{[]byte("createUser"),[]byte("redtree0"),[]byte("kim")})
  checkInvoke(t,stub,[][]byte{[]byte("addPoint"),[]byte("newUser"),[]byte("1000")})
  checkInvoke(t,stub,[][]byte{[]byte("addPoint"),[]byte("color0e"),[]byte("1000")})
  checkInvoke(t,stub,[][]byte{[]byte("addPoint"),[]byte("redtree0"),[]byte("50000")})
  checkInvoke(t,stub,[][]byte{[]byte("createContract"),[]byte("20180909"),[]byte("3"),[]byte("5.0"),
    []byte("10000"),[]byte("2018-09-09"),[]byte("redtree0"),[]byte(""),[]byte("newUser"),[]byte("Yet")})
  checkQuery(t,stub,[][]byte{[]byte("queryPoint"),[]byte("redtree0")},expected1)
  checkInvoke(t,stub,[][]byte{[]byte("signContract"),[]byte("20180909_1"),[]byte("color0e")})
  checkInvoke(t,stub,[][]byte{[]byte("completeContract"),[]byte("20180909_1")})
  checkQuery(t,stub,[][]byte{[]byte("queryPoint"),[]byte("color0e")},expected2)
}

func TestCargo_contract04(t *testing.T){
  expected1:=`[{"Key":"CARGO20180909_1", "Record":{"weight":3,"distance":5,"money":10000,"date":"2018-09-09","Registrant":"redtree0","Driver":"","Recipient":"newUser","Status":"Yet"}},{"Key":"CARGO20180909_2", "Record":{"weight":3,"distance":5,"money":20000,"date":"2018-09-09","Registrant":"redtree0","Driver":"","Recipient":"newUser","Status":"Yet"}}]`
  cargoc:=new(SmartContract)
  stub:=shim.NewMockStub("cargo",cargoc)
  checkInvoke(t,stub,[][]byte{[]byte("createUser"),[]byte("newUser"),[]byte("anonymous")})
  checkInvoke(t,stub,[][]byte{[]byte("createUser"),[]byte("color0e"),[]byte("lim")})
  checkInvoke(t,stub,[][]byte{[]byte("createUser"),[]byte("redtree0"),[]byte("kim")})
  checkInvoke(t,stub,[][]byte{[]byte("addPoint"),[]byte("newUser"),[]byte("1000")})
  checkInvoke(t,stub,[][]byte{[]byte("addPoint"),[]byte("color0e"),[]byte("1000")})
  checkInvoke(t,stub,[][]byte{[]byte("addPoint"),[]byte("redtree0"),[]byte("50000")})
  checkInvoke(t,stub,[][]byte{[]byte("createContract"),[]byte("20180909"),[]byte("3"),[]byte("5.0"),
    []byte("10000"),[]byte("2018-09-09"),[]byte("redtree0"),[]byte(""),[]byte("newUser"),[]byte("Yet")})
  checkInvoke(t,stub,[][]byte{[]byte("createContract"),[]byte("20180909"),[]byte("3"),[]byte("5.0"),
    []byte("20000"),[]byte("2018-09-09"),[]byte("redtree0"),[]byte(""),[]byte("newUser"),[]byte("Yet")})
  checkQuery(t,stub,[][]byte{[]byte("queryCargo"),[]byte("20180909")},expected1)
}

func TestCargo_contract05(t *testing.T){
  expected1:=`[{"Key":"CARGO20180909_1", "Record":{"weight":3,"distance":5,"money":10000,"date":"2018-09-09","Registrant":"redtree0","Driver":"","Recipient":"newUser","Status":"Yet"}}]`
  expected2:="{\"Username\":\"kim\",\"Total\":40000}"
  expected3:=`[{"Key":"CARGO20180909_1", "Record":{"weight":3,"distance":5,"money":0,"date":"2018-09-09","Registrant":"redtree0","Driver":"","Recipient":"newUser","Status":"Fail"}}]`
  expected4:="{\"Username\":\"kim\",\"Total\":50000}"
  cargoc:=new(SmartContract)
  stub:=shim.NewMockStub("cargo",cargoc)
  checkInvoke(t,stub,[][]byte{[]byte("createUser"),[]byte("newUser"),[]byte("anonymous")})
  checkInvoke(t,stub,[][]byte{[]byte("createUser"),[]byte("color0e"),[]byte("lim")})
  checkInvoke(t,stub,[][]byte{[]byte("createUser"),[]byte("redtree0"),[]byte("kim")})
  checkInvoke(t,stub,[][]byte{[]byte("addPoint"),[]byte("newUser"),[]byte("1000")})
  checkInvoke(t,stub,[][]byte{[]byte("addPoint"),[]byte("color0e"),[]byte("1000")})
  checkInvoke(t,stub,[][]byte{[]byte("addPoint"),[]byte("redtree0"),[]byte("50000")})
  checkInvoke(t,stub,[][]byte{[]byte("createContract"),[]byte("20180909"),[]byte("3"),[]byte("5.0"),
    []byte("10000"),[]byte("2018-09-09"),[]byte("redtree0"),[]byte(""),[]byte("newUser"),[]byte("Yet")})
  checkQuery(t,stub,[][]byte{[]byte("queryCargo"),[]byte("20180909")},expected1)
  checkQuery(t,stub,[][]byte{[]byte("queryPoint"),[]byte("redtree0")},expected2)
  checkInvoke(t,stub,[][]byte{[]byte("cancelContract"),[]byte("20180909_1")})
  checkQuery(t,stub,[][]byte{[]byte("queryCargo"),[]byte("20180909")},expected3)
  checkQuery(t,stub,[][]byte{[]byte("queryPoint"),[]byte("redtree0")},expected4)


}
