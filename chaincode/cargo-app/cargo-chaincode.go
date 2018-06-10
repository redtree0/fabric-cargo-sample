
 package main

 import (
	 "fmt"
  	"time"
	 // "reflect"
	// "bytes"
	 "encoding/json"
	//  "strconv"
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 "github.com/hyperledger/fabric/protos/peer"
 )
 

 
 	/********************************************************
	 상태값
	 Succcess - 운전자와 화물 의뢰자 간 계약이 채결이 된 상태
	 FAIL - 화물 계약이 취소 및 파토?
     Yet - 운전자와 화물 의뢰자 간 계약이 채결이 되기 전 상태
	 COMPLETE - 화물 이송이 끝나고 계약이 완료됨
	**********************************************************/
 const (
	 SUCCESS string = "Success"
	 FAIL string = "Fail"
	 YET string = "Yet"
	 COMPLETE string = "Complete"
 )

 
 // SmartContract implements a simple chaincode to manage an asset
 type SmartContract struct {
 }

 	/********************************************************
	  화물 계약
	  화물계약은 CARGO + YYYYMMDD 형식이 키이다
 	  
	  TxId - 트랜젝션 ID, 화물 계약이 등록될 시 생성
	   Weight - 화물 무게
	   Distance - 거리
	   Money - 의뢰 금액
		Date - 의뢰 일
		Registrant - 화물 의뢰자
		Driver - 화물 운송업자
		Status - 계약 상태
	**********************************************************/
type CargoContract struct {
	 TxId string `json:"txId"`
	 Weight int `json:"weight"`
	 Distance float64 `json:"distance"`
	 Money float64 `json:"money"`
	 Date string `json:"date"`
	 Registrant string `json:registrant`
	 Driver string `json:driver`
	 Status string `json:status`
 }

 /********************************************************
	 계정 별 포인트 현황
	 User - 계정 ID
	 Total - 계정 보유 포인트
	**********************************************************/
 type PointContract struct {
	 User string `json:user`
	 Total int `json:total`
 }
 
 var logger = shim.NewLogger("example_cc0")
 
 // Init is called during chaincode instantiation to initialize any
 // data. Note that chaincode upgrade also calls this function to reset
 // or to migrate data.
 func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	 logger.Info("########### example_cc0 Init ###########")

	 return shim.Success(nil)
 }
 
 // Invoke is called per transaction on the chaincode. Each transaction is
 // either a 'get' or a 'set' on the asset created by Init function. The Set
 // method may create a new asset by specifying a new key-value pair.
 func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	 logger.Info("########### example_cc0 Invoke ###########")
	 // Extract the function and args from the transaction proposal
	 fn, args := stub.GetFunctionAndParameters()
 
	 logger.Info(stub.GetTxID())
	 logger.Info(stub.GetChannelID())
 

	 switch method := fn; method {
		 case "queryAllCargo":
			 return t.queryAllCargo(stub, args)
		 case "initLedger":
			 return t.initLedger(stub)
		 case "createContract":
			 return t.createContract(stub, args)
		 case "changeStatus":
			 return t.changeStatus(stub, args)
		 case "queryCargo" :
			  return t.queryCargo(stub, args)
		case "queryPoint" :
			return t.queryPoint(stub, args)
		case "addPoint":
			return t.addPoint(stub, args)
		case "subtractPoint":
			return t.subtractPoint(stub, args)
		 default :
		      return shim.Success([]byte(nil))
	 }

 }


 	/********************************************************
	 체인코드 실행 시 실행되는 초기 데이터 셋
	 docker-compose에 정의됨
	 cargo , point에 관련된 데이터 셋 정의
	**********************************************************/
 func (t *SmartContract) initLedger(stub shim.ChaincodeStubInterface) peer.Response {

	now := time.Now()
	dateTestValue := now.Format("2006-01-02")

	cargos := []CargoContract{
		CargoContract{TxId : "txId1", Weight: 3, Distance: 5.0, Money: 5.0, Date: dateTestValue, Registrant : "you", Driver : "me", Status : SUCCESS},
		CargoContract{TxId : "txId2",Weight: 6, Distance: 5.0, Money: 5.0, Date: dateTestValue, Registrant : "you", Driver : "me", Status : YET},
		CargoContract{TxId : "txId3",Weight: 4, Distance: 5.0, Money: 5.0, Date: dateTestValue, Registrant : "you", Driver : "me", Status : SUCCESS},
		CargoContract{TxId : "txId4",Weight: 3, Distance: 5.0, Money: 5.0, Date: dateTestValue, Registrant : "you", Driver : "me", Status : SUCCESS},
		CargoContract{TxId : "txId5",Weight: 3, Distance: 5.0, Money: 5.0, Date: dateTestValue, Registrant : "you", Driver : "me", Status : FAIL},
	 }
	 cargosAsBytes, _ := json.Marshal(cargos)
	 date := now.Format("20060102")
     fmt.Println(date)
	 stub.PutState("CARGOS" + date , cargosAsBytes)


	 testdate := time.Date(
        2018, 6, 6, 0, 0, 0, 0, time.UTC)
    
	 dateTestValue2 := testdate.Format("2006-01-02")
	 cargos1 := []CargoContract{
		CargoContract{TxId : "txId1",Weight: 3, Distance: 5.0, Money: 5.0, Date: dateTestValue2, Registrant : "you", Driver : "me", Status : FAIL},
		CargoContract{TxId : "txId2",Weight: 6, Distance: 5.0, Money: 5.0, Date: dateTestValue2, Registrant : "you", Driver : "me", Status : COMPLETE},
		CargoContract{TxId : "txId3",Weight: 4, Distance: 5.0, Money: 5.0, Date: dateTestValue2, Registrant : "you", Driver : "me", Status : SUCCESS},
		CargoContract{TxId : "txId4",Weight: 3, Distance: 5.0, Money: 5.0, Date: dateTestValue2, Registrant : "you", Driver : "me", Status : SUCCESS},
		CargoContract{TxId : "txId5",Weight: 3, Distance: 5.0, Money: 5.0, Date: dateTestValue2, Registrant : "you", Driver : "me", Status : YET},
	 }
	 cargosAsBytes1, _ := json.Marshal(cargos1)
	 date1 := testdate.Format("20060102")
	 stub.PutState("CARGOS" + date1 , cargosAsBytes1)

	
	 point := []PointContract{
		PointContract{User : "you", Total : 500},
		PointContract{User : "me", Total : 1000},
		PointContract{User : "kim", Total : 1500},
	 }

	 pointAsBytes, _ := json.Marshal(point[0])
	 stub.PutState("USER.you" , pointAsBytes)
	 pointAsBytes1, _ := json.Marshal(point[1])
	 stub.PutState("USER.me" , pointAsBytes1)
	 pointAsBytes2, _ := json.Marshal(point[2])
	 stub.PutState("USER.kim" , pointAsBytes2)


	return shim.Success(nil)
 }
 


 // main function starts up the chaincode in the container during instantiate
 func main() {
	 if err := shim.Start(new(SmartContract)); err != nil {
		 fmt.Printf("Error starting SmartContract chaincode: %s", err)
	 }
 }
 