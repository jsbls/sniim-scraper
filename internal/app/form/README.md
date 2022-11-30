# Formulario 0
Los inputs están situadoe dentro de tablas.  
Las tablas pueden tener los siguientes ids.
- table#[tblDatos]
- table#[tblFiltro]

Los parámetros son enviados por método *HTTP-GET* a este endpoint /Nuevo/Consultas/MercadosNacionales/PreciosDeMercado/Agricolas/

- El rango de fechas siempre es necesario.


## Datos generales

| Filter | Selector | UrlParam |
|--|--|--|
| Producto | select[id*=ddlProducto] | ProductoId |
| Origen | select[id*=ddlOrigen] | OrigenId |
| Destino | select[id*=ddlDestino] | DestinoId |
| Precios por | select[id*=ddlPrecios] | PreciosPorId |

## Reportes diarios

| Filter | Selector | UrlParam |
|--|--|--|
| Desde | -- | fechainicio | 
| Hasta | -- | fechafinal |

## Reportes semanales
| Filter | Selector | UrlParam |
|--|--|--|
| Semana | select[id=ddlSemanaSemanal] | Semana |
| Mes | select[id=ddlMesSemanal] | Mes |
| Año | select[id=ddlAnioSemana] | Anio |