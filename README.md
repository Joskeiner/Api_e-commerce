
# API eCommerce

Esta es una simple API Restful escrita en GO. Para esta API, se utilizó Fiber como framework web, GORM como ORM y PostgreSQL como base de datos. También se empleó ZAP para los registros (logs) y Viper para el manejo de variables de entorno.

La intención de este proyecto es meramente educativa, por lo que se irán agregando mejoras con el tiempo
## Installation
En esta guía se da por hecho que usted sabe cómo clonar un repositorio en Git.

1.segúrese de tener GoLang v1.21.3 o superior instalado.
link [GO](https://go.dev/doc/install)
```bash
 go version
```
2.Asegurese de tener [Task](https://taskfile.dev/installation/)
instalado 

```bash
 task --verion
```

3.instalar todas las herramientas necesarias para el proyecto  

```bash
 task install
```

4.Copie el contenido o renombra el archivo .env.example a .env  

```bash
cp .env.example .env
```
## Documentation

Esquema de base de datos [Aqui](https://github.com/Joskeiner/Api_e-commerce/tree/main/UML) 

swagger  en progreso


## Referencias

 - [go-clean-template](https://awesomeopensource.com/project/elangosundar/awesome-README-templates) by Evrone
 - [README](https://readme.so/es/editor)
 - [go-clean-arch ](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) by Uncle Bob
 - [gonmmerce](https://github.com/bagashiz/gommerce) by  bagashiz
