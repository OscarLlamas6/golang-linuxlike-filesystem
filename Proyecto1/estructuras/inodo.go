package estructuras

//Inodo struct
type Inodo struct {
	Proper             [20]byte
	NumeroInodo        int32
	FileSize           int32
	NumeroBloques      int32
	ApuntadoresBloques [4]int32
	ApuntadorIndirecto int32
}

//BloqueDatos struct
type BloqueDatos struct {
	Data [25]byte
}
