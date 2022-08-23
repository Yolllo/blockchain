package systemSmartContracts

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go/vm"
	vmcommon "github.com/ElrondNetwork/elrond-vm-common"
)

type customStaking struct {
	eei     vm.SystemEI
	gasCost vm.GasCost

	mutExecution sync.RWMutex
}

type ArgsNewCustomStaking struct {
	Eei               vm.SystemEI
	GasCost           vm.GasCost
	StakingAccessAddr []byte
}

func NewCustomStakingSystemSC(args ArgsNewCustomStaking) (*customStaking, error) {
	/*
		fmt.Println("==================================+")
		var addressEncoder, _ = pubkeyConverter.NewBech32PubkeyConverter(32, &mock.LoggerMock{})
		customStakingSCAddress := addressEncoder.Encode(vm.CustomStakingSCAddress)
		//fmt.Println(args.StakingSCAddress)
		fmt.Println("CustomStakingSCAddress: " + customStakingSCAddress)
		fmt.Println("==================================+")
	*/

	cs := &customStaking{
		eei: args.Eei,
	}

	return cs, nil
}

func (cs *customStaking) Execute(args *vmcommon.ContractCallInput) vmcommon.ReturnCode {
	cs.mutExecution.RLock()
	defer cs.mutExecution.RUnlock()
	fmt.Println(args.Function)

	switch args.Function {
	case core.SCDeployInitFunctionName:
		return cs.init(args)
	case "TestA":
		return cs.TestA(args)
	case "TestB":
		return cs.TestB(args)
	case "TestC":
		return cs.TestC(args)
	}

	cs.eei.AddReturnMessage(args.Function + " is an unknown function")
	return vmcommon.UserError
}

func (cs *customStaking) init(args *vmcommon.ContractCallInput) vmcommon.ReturnCode {

	return vmcommon.Ok
}

func (cs *customStaking) TestA(args *vmcommon.ContractCallInput) vmcommon.ReturnCode {
	err := cs.eei.UseGas(10000)
	if err != nil {
		cs.eei.AddReturnMessage(err.Error())
		return vmcommon.OutOfGas
	}
	data := args.Arguments[0]
	cs.eei.SetStorage([]byte("test_key"), data)
	cs.eei.AddReturnMessage(string(data))

	return vmcommon.Ok
}

func (cs *customStaking) TestB(args *vmcommon.ContractCallInput) vmcommon.ReturnCode {
	cs.eei.SetStorage([]byte("test_key2"), []byte("lolo"))

	data1 := cs.eei.GetStorage([]byte("test_key"))
	data2 := cs.eei.GetStorage([]byte("test_key2"))
	data3 := cs.eei.GetStorage([]byte("test_key3"))

	dataResp := string(data1) + string(data2) + string(data3)

	cs.eei.AddReturnMessage(dataResp)

	return vmcommon.Ok
}

func (cs *customStaking) TestC(args *vmcommon.ContractCallInput) vmcommon.ReturnCode {
	err := cs.eei.UseGas(cs.gasCost.MetaChainSystemSCsCost.DelegationOps)
	if err != nil {
		cs.eei.AddReturnMessage(err.Error())
		return vmcommon.OutOfGas
	}

	transferValue := big.NewInt(12345678)
	err = cs.eei.Transfer([]byte("yol1vuem526uwdnzure9y40hvxmepz86lf8f3dppx8fmcjpcl9ayy59sljgfmt"), []byte("yol14dc7vlg4we0s6rw9z7cwht8zza4mnmfavv8gttt0qvnvxytm5s7sxrg5pp"), transferValue, nil, 0)
	if err != nil {
		cs.eei.AddReturnMessage(err.Error())
		return vmcommon.UserError
	}

	return vmcommon.Ok
}

// CanUseContract returns true if contract can be used
func (cs *customStaking) CanUseContract() bool {
	return true
}

// SetNewGasCost is called whenever a gas cost was changed
func (cs *customStaking) SetNewGasCost(gasCost vm.GasCost) {
	cs.mutExecution.Lock()
	cs.gasCost = gasCost
	cs.mutExecution.Unlock()
}

// IsInterfaceNil returns true if underlying object is nil
func (cs *customStaking) IsInterfaceNil() bool {
	return cs == nil
}
