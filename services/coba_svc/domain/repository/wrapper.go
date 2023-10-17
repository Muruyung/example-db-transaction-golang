package repository

type Wrapper struct {
	CobaDuaRepo  CobaDuaRepository
	CobaSatuRepo CobaSatuRepository
	SqlTx
}
