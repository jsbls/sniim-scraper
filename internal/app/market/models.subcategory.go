package market

/*
	SubMarket
*/

type SubCategory struct {
	Name string `json:"name"` // Frutas y hortalizas
	Url  string `json:"url"`  // http://www.economia-sniim.gob.mx/Nuevo/Home.aspx?opcion=Consultas/MercadosNacionales/PreciosDeMercado/Agricolas/ConsultaFrutasYHortalizas.aspx?SubOpcion=4
}

func (m *SubCategory) IsNotEmpty() bool {
	return m.Name != ""
}
