package enums

type DBTypeRegisters string

const (
	Transaction DBTypeRegisters = "TX"
	FileCharge  DBTypeRegisters = "FILE"
	Contact     DBTypeRegisters = "CONTACT"
)
