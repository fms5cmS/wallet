package constant

type ChainArchitecture uint8

const (
	ChainArchitectureUnknown ChainArchitecture = iota
	ChainArchitectureEVM                       // evm 架构
	ChainArchitectureSVM                       // Soltana
	ChainArchitectureTVM                       // TON
)

const (
	EVMZeroAddress = "0x0000000000000000000000000000000000000000"
	SVMZeroAddress = ""
	TVMZeroAddress = ""
)
