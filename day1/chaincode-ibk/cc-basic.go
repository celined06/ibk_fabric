package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"	//gopath에 shim 패키지가 있어야함..
	"github.com/hyperledger/fabric/protos/peer"
)

type User struct {
	userId   string
	userName string
}

type SmartContract struct {
}

func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}


//자신이 호출하려는 펑션은 무조건 Invoke 메소드에 있어야함
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()	//펑션하고 밸류를 구분해줌..GetFunctionAndParameters()

	if function == "Create" {	//["Create" , "user01", "허찬"]
		return t.Create(stub, args)
	}

	if function == "Read" {		//["Read", "user01"]
		return t.Read(stub, args)
	}

	return shim.Success(nil)

}

//
func (t *SmartContract) Create(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	err := stub.PutState(args[0], []byte(args[1]))	//저장
	//value, err := stub.PutState(args[0], []byte(args[1])) 저장 후 바로읽을 경우?

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SmartContract) Read(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	value, err := stub.GetState(args[0])

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(value)
}

func main() {	//패키지 안에는 무조건 main() method 가 있어야함
	err := shim.Start(new(SmartContract))

	if err != nil {
		fmt.Printf("Error starting SmartContract: %s", err)
	}
}
