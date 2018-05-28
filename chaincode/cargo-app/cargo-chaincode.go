
 package main

 import (
	 "fmt"
 //	"time"
	 // "reflect"
	 "bytes"
	 "encoding/json"
	 "strconv"
	 "github.com/hyperledger/fabric/core/chaincode/shim"
	 "github.com/hyperledger/fabric/protos/peer"
 )
 
 type Status int
 
 const (
	 SUCCESS Status = 1+ iota
	 FAIL
	 YET
	 COMPLETE
 )
 
 var status = [...]string{
	 "Success",
	 "Fail",
	 "Yet",
	 "Complete" ,
 }
 
 func (s Status) String() string{ return status[(s-1)%4]}
 
 // SmartContract implements a simple chaincode to manage an asset
 type SmartContract struct {
 }
 
 // Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
 type CargoContract struct {
	 // Make   string `json:"make"`
	 // Model  string `json:"model"`
	 // Colour string `json:"colour"`
	 // Owner  string `json:"owner"`
	 Weight int `json:"weight"`
	 Distance float64 `json:"distance"`
	 Money float64 `json:"money"`
	 Date string `json:"date"`
	 Registrant string `json:registrant`
	 Driver string `json:driver`
	 Status Status `json:status`
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
			 return t.queryAllCargo(stub)
		 case "initLedger":
			 return t.initLedger(stub)
		 case "createContract":
			 return t.createContract(stub, args)
		 case "changeStatus":
			 return t.changeStatus(stub, args)
		 case "queryCargo" :
			  return t.queryCargo(stub, args)
		 default :
		      return shim.Success([]byte(nil))
	 }

 }


 func (s *SmartContract) queryCargo(stub shim.ChaincodeStubInterface, args []string) peer.Response {
 
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	tunaAsBytes, _ := stub.GetState(args[0])
	if tunaAsBytes == nil {
		return shim.Error("Could not locate tuna")
	}
	return shim.Success(tunaAsBytes)
}
 
 func  (t *SmartContract) queryAllCargo(stub shim.ChaincodeStubInterface) peer.Response {
	 startKey := "CARGO0"
	 endKey := "CARGO999"
 
	 resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	 if err != nil {
		 return shim.Error(err.Error())
	 }
	 defer resultsIterator.Close()
 
	 // buffer is a JSON array containing QueryResults
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
 
 func (t *SmartContract) initLedger(stub shim.ChaincodeStubInterface) peer.Response {
 
	 contracts := []CargoContract{
		 CargoContract{Weight: 3, Distance: 5.0, Money: 5.0, Date: "2018-05-26", Registrant : "you", Driver : "me", Status : Status(1)},
		 CargoContract{Weight: 6, Distance: 5.0, Money: 5.0, Date: "2018-05-26", Registrant : "you", Driver : "me", Status : Status(2)},
		 CargoContract{Weight: 4, Distance: 5.0, Money: 5.0, Date: "2018-05-26", Registrant : "you", Driver : "me", Status : Status(0)},
		 CargoContract{Weight: 3, Distance: 5.0, Money: 5.0, Date: "2018-05-26", Registrant : "you", Driver : "me", Status : Status(2)},
		 CargoContract{Weight: 3, Distance: 5.0, Money: 5.0, Date: "2018-05-26", Registrant : "you", Driver : "me", Status : Status(3)},
	 }
 
	 i := 0
	 for i < len(contracts) {
		 fmt.Println("i is ", i)
		 cargoAsBytes, _ := json.Marshal(contracts[i])
		 stub.PutState("CARGO"+strconv.Itoa(i), cargoAsBytes)
		 fmt.Println("Added", contracts[i])
		 i = i + 1
	 }
 
	 return shim.Success(nil)
 }
 
 
 func (t *SmartContract) createContract(stub shim.ChaincodeStubInterface, args []string) peer.Response {
 
	 if len(args) != 8 {
		 return shim.Error("Incorrect number of arguments. Expecting 5")
	 }
	 w, _ := strconv.Atoi(args[1])
	 d, _ := strconv.ParseFloat(args[2], 64)
	 m, _ := strconv.ParseFloat(args[3], 64)
	 s, _ := strconv.Atoi(args[7])
	 var cargo = CargoContract{Weight: w, Distance: d, Money: m, 
		 Date: args[4], Registrant : args[5], Driver : args[6], Status : Status(s)}
 
	 cargoAsBytes, _ := json.Marshal(cargo)
	 key := args[0]
	 stub.PutState(key, cargoAsBytes)
 
	 return shim.Success(nil)
 }
 
 func (t *SmartContract) changeStatus(stub shim.ChaincodeStubInterface, args []string) peer.Response {
 
	 if len(args) != 2 {
		 return shim.Error("Incorrect number of arguments. Expecting 2")
	 }
 
	 cargoAsBytes, _ := stub.GetState(args[0])
	 cargo := CargoContract{}
 
	 json.Unmarshal(cargoAsBytes, &cargo)
	 s, _ := strconv.Atoi(args[1])
	 cargo.Status = Status(s)
 
	 cargoAsBytes, _ = json.Marshal(cargo)
	 stub.PutState(args[0], cargoAsBytes)
 
	 return shim.Success(nil)
 }
 
 
 // main function starts up the chaincode in the container during instantiate
 func main() {
	 if err := shim.Start(new(SmartContract)); err != nil {
		 fmt.Printf("Error starting SmartContract chaincode: %s", err)
	 }
 }
 