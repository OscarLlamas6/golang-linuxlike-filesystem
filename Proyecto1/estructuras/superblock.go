package estructuras

//Superblock struct
type Superblock struct {
	Name                [16]byte
	TotalAVDS           int32
	TotalDDS            int32
	TotalInodos         int32
	TotalBloques        int32
	TotalBitacoras      int32
	FreeAVDS            int32
	FreeDDS             int32
	FreeInodos          int32
	FreeBloques         int32
	FreeBitacoras       int32
	DateCreacion        [20]byte
	DateLastMount       [20]byte
	MontajesCount       int32
	InicioBitmapAVDS    int32
	InicioAVDS          int32
	InicioBitMapDDS     int32
	InicioDDS           int32
	InicioBitmapInodos  int32
	InicioInodos        int32
	InicioBitmapBloques int32
	InicioBloques       int32
	InicioBitacora      int32
	SizeAVD             int32
	SizeDD              int32
	SizeInodo           int32
	SizeBloque          int32
	SizeBitacora        int32
	FirstFreeAVD        int32
	FirstFreeDD         int32
	FirstFreeInodo      int32
	FirstFreeBloque     int32
	MagicNum            int32
}
