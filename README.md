# SNIIM Scraper

Scraper para obtener los datos de precios sobre algunos productos como frutos y carnes en México.


## 📄 Variables de entorno

| Nombre | Descripción | Default |
|--|--|--|
| SNIIM_ADDR | Dirección del sitio fuente de la información | http://www.economia-sniim.gob.mx |
| CATALOGUE_SRC | Nombre de la base de datos o directorio del  filesystem para guardar los catálogos | SNIIM_DATA |
| DEBUG | Bandera para habilitar el modo debug | false |
| MONGO_URI* | Dirección de la base de datos mongo, solo se intentará conectar si está presente. | '' |


\* [TODO](./docs/README.md)
## 📦 Instalación

Para realizar una instalación de la herramienta puede visitar la [página de releases](https://github.com/everitosan/sniim-scrapper/releases) o bien descargar el código fuente y seguir los pasos de [compilación](./docs/Compilation.md).

Si decide usar una versión [precompilada]((https://github.com/everitosan/sniim-scrapper/releases)), una vez que descargue el archivo debe descomprimir el archivo .tar.gz

<br>

*🐧 Ejemplo en Linux*
```bash
# Descomprimir con tar
$ tar -xf sniim-scraper_0.0.1_Linux_arm64.tar.gz

# Ejecución de prueba
$ ./sniim-cli -h                        
Usage:
  sniim-cli [flags]
  sniim-cli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  consult     Manage consults
  help        Help about any command
  init        Create catalogues
  request     Request information

Flags:
  -h, --help      help for sniim-cli
  -v, --version   Show version of the cli

Use "sniim-cli [command] --help" for more information about a command.

```




## 🍇 Uso

Inicialización
---

El CLI necesita tener los catálogos disponibles para poder operar.

Para esto existe el comando init. 

🕹️ *Carga de catálogos.*

```bash
sniim-cli init
```
![](./docs/img/init.png)

Registro de consultas
---

Para poder saber el precio de algún artículo, podemos crear una consulta. 

El comando **consult** con la bandera **-c** ó **--create** invocará una serie de preguntas basadas en la categoría y artículo seleccionados.

Si deseamos guardar esa consulta para volver a ejecutarla posteriormente, podemos agregar la bandera **-s** ó **--save**.

🕹️ *Crea un registro de consulta.*

```bash
sniim-cli consult -c
```
![](./docs/img/consult--create.png)

> 📝 Nota:  
*Al usar la palabra reservada **now**, estaremos indicando que la fecha de interés será la fecha en la que se ejecute la petición con el comando **request**.*

También podemos mostar una lista de las consultas que tenemos guardadas en registros.

🕹️ *Listar las consultas registradas.*

```bash
sniim-cli consult -l
```

![](./docs/img/consult--list.png)

Ejecución de Consultas
---

Si deseamos repetir alguna de las consultas guardadas, podemos usar el comando request.

🕹️ *Realiza una consulta.*

```bash
sniim-cli request -i [índice]
```

![](./docs/img/request--index.png)

> 📝 Nota:  
***-s** Es una bandera opcional que se puede agregar al comando de request para guardar en resultado obtenido en lugar de mostrarlo en la consola*