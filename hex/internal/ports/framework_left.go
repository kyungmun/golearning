package ports

type GRPCPort interface {
	Run()
	GetAddition()
	GetSubstraction()
	GetMultipulication()
	GetDivision()
}
