package plasma

import (
	"context"
	"crypto/ecdsa"
	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind/backends"
	"github.com/Onther-Tech/plasma-evm/contracts/plasma/ton"
	"github.com/Onther-Tech/plasma-evm/core"
	"github.com/Onther-Tech/plasma-evm/params"
	"github.com/Onther-Tech/plasma-evm/pls"
	"math/big"
	"testing"

	"github.com/Onther-Tech/plasma-evm/accounts/abi/bind"
	"github.com/Onther-Tech/plasma-evm/common"
	"github.com/Onther-Tech/plasma-evm/crypto"
)

var (
	operatorKey, _   = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	operator         = crypto.PubkeyToAddress(operatorKey.PublicKey)
	challengerKey, _ = crypto.HexToECDSA("78ae75d1cd5960d87e76a69760cb451a58928eee7890780c352186d23094a114")
	challenger       = crypto.PubkeyToAddress(challengerKey.PublicKey)
	operatorOpt      = bind.NewKeyedTransactor(operatorKey)

	addr1 = common.HexToAddress("0x5df7107c960320b90a3d7ed9a83203d1f98a811d")
	addr2 = common.HexToAddress("0x3cd9f729c8d882b851f8c70fb36d22b391a288cd")
	addr3 = common.HexToAddress("0x57ab89f4eabdffce316809d790d5c93a49908510")
	addr4 = common.HexToAddress("0x6c278df36922fea54cf6f65f725267e271f60dd9")
	addrs = []common.Address{addr1, addr2, addr3, addr4}

	key1, _ = crypto.HexToECDSA("78ae75d1cd5960d87e76a69760cb451a58928eee7890780c352186d23094a115")
	key2, _ = crypto.HexToECDSA("bfaa65473b85b3c33b2f5ddb511f0f4ef8459213ada2920765aaac25b4fe38c5")
	key3, _ = crypto.HexToECDSA("067394195895a82e685b000e592f771f7899d77e87cc8c79110e53a2f0b0b8fc")
	key4, _ = crypto.HexToECDSA("ae03e057a5b117295db86079ba4c8505df6074cdc54eec62f2050e677e5d4e66")
	keys    = []*ecdsa.PrivateKey{key1, key2, key3, key4}

	plsConfig     = pls.DefaultConfig
	staminaConfig = params.DefaultStaminaConfig
)

func TestDeployPlasmaContracts(t *testing.T) {
	g := core.DeveloperGenesisBlock(0, common.Address{}, common.Address{1}, params.DefaultStaminaConfig)
	contractBackend := backends.NewSimulatedBackend(g.Alloc, 10000000000)

	tonAddr, _, _, err := ton.DeployTON(operatorOpt, contractBackend)
	if err != nil {
		t.Fatalf("Failed to deploy TON: %v", err)
	}

	// with previously deployed TON
	testDeployPlasmaContracts(t, contractBackend, common.Address{})

	// without previously deployed TON
	testDeployPlasmaContracts(t, contractBackend, tonAddr)
}

func testDeployPlasmaContracts(t *testing.T, backend *backends.SimulatedBackend, tonAddr common.Address) {
	rootchainAddr, genesis, err := DeployPlasmaContracts(operatorOpt, backend, plsConfig, staminaConfig, tonAddr)

	if err != nil {
		t.Fatalf("Failed to deploy RootChain: %v", err)
	}

	genesis.ToBlock()
}
