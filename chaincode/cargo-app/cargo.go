
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
	 특정 화물 계약을 조회하는 메소트
	**********************************************************/
//peer chaincode query -n cargo-app -c '{"Args":["queryCargo", "CARGOS20180606"]}' -C mychannel
 func (s *SmartContract) queryCargo(stub shim.ChaincodeStubInterface, args []string) peer.Response {
 
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}


	cargoAsBytes, _ := stub.GetState("CARGOS" + args[0])
	// cargoAsBytes, _ := stub.GetState(args[0])
	if cargoAsBytes == nil {
		return shim.Error("Could not locate tuna")
	}
	return shim.Success(cargoAsBytes)
}
 

 	/********************************************************
	 args[0] - "all"
	 화물 계약을 전체 조회하는 메소트
	**********************************************************/
//peer chaincode query -n cargo-app -c '{"Args":["queryAllCargo"]}' -C mychannel
// peer chaincode query -n cargo-app -c '{"Args":["queryAllCargo", "all"]}' -C mychannel
 func  (t *SmartContract) queryAllCargo(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// var logger = shim.NewLogger("example_cc0")
	// logger("###########queryAllCargo#############")

  
	if args[0] == "all" {
		now := time.Now()
		now = now.AddDate(0,0,1)
		date := now.Format("20060102")
		
		startKey := "CARGOS20180606"
		endKey := "CARGOS" + date 
	

		resultsIterator, err := stub.GetStateByRange(startKey, endKey)
		if err != nil {
			return shim.Error(err.Error())
		}
		defer resultsIterator.Close()


		var buffer bytes.Buffer
		buffer.WriteString("[")

		bArrayMemberAlreadyWritten := false
		for resultsIterator.HasNext() {
			queryResponse, err := resultsIterator.Next()
			if err != nil {
				return shim.Error(err.Error())
			}
			// Add a comma before array members, suppress it for the first array member
			if bArrayMemberAlreadyWritten == true {
				buffer.WriteString(",")
			}
			buffer.WriteString("{\"Key\":")
			buffer.WriteString("\"")
			buffer.WriteString(queryResponse.Key)
			buffer.WriteString("\"")

			buffer.WriteString(", \"Record\":")
			// Record is a JSON object, so we write as-is
			buffer.WriteString(string(queryResponse.Value))
			buffer.WriteString("}")
			bArrayMemberAlreadyWritten = true
		}
		buffer.WriteString("]")

		fmt.Printf("- queryAllCARGO:\n%s\n", buffer.String())
	 
		 return shim.Success(buffer.Bytes())

	}


	 return shim.Success(nil)
 }
 

 	/********************************************************
	 args[0] - key 날짜 YYYYMMDD
	 args[1] - 이전 txId
	 args[2] - 상태값 Success, Complete, Yet, Fail
	 화물 계약을 등록하는 메소드
	**********************************************************/
 func (t *SmartContract) createContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {
 
	 if len(args) != 8 {
		 return shim.Error("Incorrect number of arguments. Expecting 8")
	 }

	key := "CARGOS"+args[0]
	cargoDataSets, _ := stub.GetState(key)

	if cargoDataSets == nil {
		return shim.Error("Could not locate tuna")
	}

	w, _ := strconv.Atoi(args[1])
	d, _ := strconv.ParseFloat(args[2], 64)
	m, _ := strconv.ParseFloat(args[3], 64)
	// s, _ := strconv.Atoi(args[7])

	// state :=	Status(s);
	var cargo = CargoContract{TxId : stub.GetTxID(), Weight: w, Distance: d, Money: m, 
		Date: args[4], Registrant : args[5], Driver : args[6], Status : args[7] }

	var cargos []CargoContract 	
	_ = json.Unmarshal( cargoDataSets, &cargos )
	cargos = append( cargos, cargo )
  
	// Encode as JSON
	// Put back on the block
	bytes, _ := json.Marshal( cargos )
	_ = stub.PutState( key, bytes )
	return shim.Success(nil)
 }
 
 	/********************************************************
	 args[0] - key 날짜 YYYYMMDD
	 args[1] - 이전 txId
	 args[2] - 상태값 Success, Complete, Yet, Fail
	 화물 계약 상태를 변경하는 메소드
	**********************************************************/
 func (t *SmartContract) changeStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	 if len(args) != 3 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }
 

	 key := "CARGOS"+ args[0]
	 cargoAsbytes, err := stub.GetState(key)
  
	 if err != nil {
		return shim.Error( "Unable to get accounts." )
	 }
   
	 var cargos []CargoContract
   
	 // From JSON to data structure
	 _ = json.Unmarshal( cargoAsbytes, &cargos )
   
	 // Look for match
	 for a := 0; a < len( cargos ); a++ {
	   // Match
	   if cargos[a].TxId == args[1] {
			cargos[a].Status = args[2]
			cargos[a].TxId =  stub.GetTxID()
			// cargos[a].Password = args[2]
		 break
	   }
	 }
   
	 // Encode as JSON
	 // Put back on the block
	 bytes, err := json.Marshal( cargos )
	 _ = stub.PutState(key, bytes )
	 //fmt.Printf("Query Response:%s\n", bytes)
	 return shim.Success(nil)
 }
 