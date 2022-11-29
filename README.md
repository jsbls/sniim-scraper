# SNIIM Scrapper

Este es un scrapper que pretende obtener los datos de precios sobre algunos productos como frutos y carnes en México.

La estructura de este proyecto está basada en el [post de Medium](https://medium.com/geekculture/how-to-structure-your-project-in-golang-the-backend-developers-guide-31be05c6fdd9)

Para poder trabajar de forma más eficiente, la herramienta debe bajar los *catálogos* antes de poder ser usada.


## Variables de entorno

| Nombre | Descripción | Default |
|--|--|--|
| SNIIM_ADDR | Dirección del sitio fuente de la infrmación | http://www.economia-sniim.gob.mx |
| CATALOGUE_SRC | Nombre de la base de datos o directorio del  filesystem para guardar los catálogos | SNIIM_DATA |
| DEBUG | Bandera para habilitar el modo debug | false |
| MONGO_URI | Dirección de la base de datos mongo, solo se intentará conectar si está presente dicha variable. | '' |

## Cli

### Compilación
```bash
go build -o sniim-cli ./cmd/cli/main.go
```

### Uso

| > *Carga de catálogos*

```bash
sniim-cli init
```