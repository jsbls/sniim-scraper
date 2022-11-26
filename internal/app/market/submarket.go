package market

/*
	SubMarket
*/

type SubMarket struct {
	Name string // Frutas y hortalizas
	Url  string // http://www.economia-sniim.gob.mx/Nuevo/Home.aspx?opcion=Consultas/MercadosNacionales/PreciosDeMercado/Agricolas/ConsultaFrutasYHortalizas.aspx?SubOpcion=4
}

func (m *SubMarket) IsNotEmpty() bool {
	return m.Name != ""
}
