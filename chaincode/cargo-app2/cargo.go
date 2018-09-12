
 package main

 import (
	 "fmt"
 	"time"

     "bytes"
	 "encoding/json"
	 "strconv"
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 "github.com/hyperledger/fabric/protos/peer"
 )


 	/********************************************************
	 args[0] - 날짜 YYYYMMDD
	 특정 화물 계약을 조회하는 메소드
   매개변수로 전달된 날짜부터 현재날짜까지 운송의뢰들을 조회하는 메소드
	**********************************************************/
//peer chaincode query -n cargo-app -c '{"Args":["queryCargo", "CARGOS20180606"]}' -C mychannel
 func (s *SmartContract) queryCargo(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

  now := time.Now()
  now = now.AddDate(0,0,1)
  date := now.Format("20060102")

  startKey := "CARGO"+args[0]
  endKey := "CARGO"+date

  		resultsIterator, err := stub.GetStateByRange(startKey, endKey)
  		if err != nil {
  			return shim.Error(err.Error())
  		}
  		defer resultsIterator.Close()

      var buffer bytes.Buffer
      buffer.WriteString("[")
      bArrayMemberAlreadyWritten := false

      for resultsIterator.HasNext(){
        var response CargoContext
        queryResponse,_ := resultsIterator.Next()
        _ = json.Unmarshal(queryResponse.Value,&response)

        for i:=1;i<=response.End;i++{
		  num:=strconv.Itoa(i)
          qkey:=queryResponse.Key+"_"+num
          queryResponse2,_:= stub.GetState(qkey)
          if bArrayMemberAlreadyWritten == true {
            buffer.WriteString(",")
          }
          buffer.WriteString("{\"Key\":")
          buffer.WriteString("\"")
          buffer.WriteString(qkey)
          buffer.WriteString("\"")

          buffer.WriteString(", \"Record\":")
          // Record is a JSON object, so we write as-is
          buffer.WriteString(string(queryResponse2))
          buffer.WriteString("}")
          bArrayMemberAlreadyWritten = true
        }
      }
	  buffer.WriteString("]")

      fmt.Printf("- queryAllCARGO:\n%s\n", buffer.String())

       return shim.Success(buffer.Bytes())
}

 	/********************************************************
	 화물 계약을 등록하는 메소드
	**********************************************************/
 func (t *SmartContract) createContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	 if len(args) != 9 {
		 return shim.Error("Incorrect number of arguments. Expecting 8")
	 }

	 cm, _ := strconv.Atoi(args[3])
	 pointAsBytes,_:=stub.GetState(args[5])
	 var point PointContract
	 
	 json.Unmarshal(pointAsBytes,&point)
	 if point.Total < cm {
	      return shim.Error("not enough money")
	 }
	 point.Total=point.Total-cm
	 pbytes,_:=json.Marshal(point)
	 fmt.Println("==test==")
	 fmt.Println(string(pbytes))	 
     stub.PutState(args[5],pbytes)
	
	ckey := "CARGO"+args[0]
  cargocontext, _ := stub.GetState(ckey)

  cw, _ := strconv.Atoi(args[1])
  cd, _ := strconv.ParseFloat(args[2], 64)

  cargo:=CargoContract{Weight: cw, Distance: cd, Money: cm,
    Date: args[4], Registrant : args[5], Driver : args[6],Recipient:args[7], Status : args[8] }
    cargoAsBytes,_:=json.Marshal(cargo)
  if cargocontext == nil{
        context:=CargoContext{Start:1,End:1}
        cargocontextAsBytes,_:=json.Marshal(context)
        stub.PutState(ckey,cargocontextAsBytes)
        key:="CARGO"+args[0]+"_1"
        stub.PutState(key,cargoAsBytes)
        return shim.Success(nil)
  }
  //cargo context structure cargoContext{Start:1,End:1}

  var cargoc CargoContext
  _ =json.Unmarshal(cargocontext,&cargoc)

  nextnum:= cargoc.End+1
  num:=strconv.Itoa(nextnum)
  newkey:="CARGO"+args[0]+"_"+num

  stub.PutState(newkey,cargoAsBytes)
  cargoc.End=nextnum
  cargocAsBytes,_:=json.Marshal(cargoc)
  stub.PutState(ckey,cargocAsBytes)

	return shim.Success(nil)
 }

 	/********************************************************
	 화물 계약 상태를 취소로 변경하는 메소드
   SmartContract - cancelContract
   args[0]==Key
	**********************************************************/
 func (t *SmartContract) cancelContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }


	 key := "CARGO"+ args[0]
	 cargoAsbytes, err := stub.GetState(key)
	 

	 if err != nil {
		return shim.Error( "Unable to get accounts." )
	 }

	 var cargo CargoContract
	 var point PointContract

	 // From JSON to data structure
	 _ = json.Unmarshal( cargoAsbytes, &cargo )

   cargo.Status = FAIL
   
   money:=cargo.Money
   cargo.Money = 0
   registrant:=cargo.Registrant
   
   pointAsBytes,_:=stub.GetState(registrant)
   _ = json.Unmarshal(pointAsBytes,&point)
   
   point.Total=point.Total+money
   
   pbytes,_:=json.Marshal(point)
   _=stub.PutState(registrant,pbytes)
   
	 // Encode as JSON
	 // Put back on the block
	 bytes, err := json.Marshal( cargo )
	 _ = stub.PutState(key, bytes )
	 //fmt.Printf("Query Response:%s\n", bytes)
	 return shim.Success(nil)
 }
 /**********************************
 signContract - 운전자가 화물운송의뢰를 계약하는 메소드
 args[0] == Key
 args[1] == 운전자아이디

 ***********************************/

 func (t *SmartContract) signContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }


	 key := "CARGO"+ args[0]
	 cargoAsbytes, err := stub.GetState(key)

	 if err != nil {
		return shim.Error( "Unable to get accounts." )
	 }

	 var cargo CargoContract

	 // From JSON to data structure
	 _ = json.Unmarshal( cargoAsbytes, &cargo )

   cargo.Status = SUCCESS
   cargo.Driver = args[1]
	 // Encode as JSON
	 // Put back on the block
	 bytes, err := json.Marshal( cargo )
	 _ = stub.PutState(key, bytes )
	 //fmt.Printf("Query Response:%s\n", bytes)
	 return shim.Success(nil)
 }

 /**********************************
 completeContract - 수령확인시 호출되어 포인트를 운전자에게 이동시키고,상태를 변경하는 메소드
 args[0] == Key
 ***********************************/
 func (t *SmartContract) completeContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }


	 key := "CARGO"+ args[0]
	 cargoAsbytes, err := stub.GetState(key)

	 if err != nil {
		return shim.Error( "Unable to get accounts." )
	 }

	 var cargo CargoContract

	 // From JSON to data structure
	 _ = json.Unmarshal(cargoAsbytes,&cargo)

   if cargo.Status != SUCCESS {
     return shim.Error("cargo status is not SUCCESS")
   }

   cargo.Status = COMPLETE
   driver:=cargo.Driver

   pointAsBytes,_:=stub.GetState(driver)
   var point PointContract

   json.Unmarshal(pointAsBytes,&point)
   point.Total=point.Total+cargo.Money
   cargo.Money=0


	 // Encode as JSON
	 // Put back on the block
	 bytes,_:=json.Marshal(cargo)
   bytes2,_:=json.Marshal(point)
	 _=stub.PutState(key, bytes)
   _=stub.PutState(driver,bytes2)
	 //fmt.Printf("Query Response:%s\n", bytes)
	 return shim.Success(nil)
 }
