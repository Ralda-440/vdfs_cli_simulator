# Virtual Disk & Filesystem Command Line Interface Simulator (vdfs_cli_simulator)

Interfaz de lineá de comandos que simula la creación de discos, particiones, sistemas de archivos y manejo de archivos. 

Los discos son simulados utilizando archivos binarios, al crearlos se llenan con 0x00 (00000000b en binario) para inicializar cada byte y simular su capacidad total. En el, se almacenan estructuras serializadas para simular el funcionamiento de las particiones, sistema de archivos y archivos que pudiera contener.

Los comandos que se utilizan en la **cli** son reconocidos por un analizador en **go**, generado con la herramienta **ANTLR**.

Los archivos binarios se crean y manipulan en el Backend.

La aplicación web se encuentra desplegada en la plataforma **Cloud Run** de **Google Cloud**. [Acceda aquí](https://vdfs-cli-simulator-gui-5320668054.us-east4.run.app)

## Arquitectura

![Arquitectura](/manuals/images/Arquitectura.svg)

## Ejemplo de Uso

```python
#Crear un Disco
mkdisk -unit=M -size=1

#Crear Una Particion en el Disco
fdisk -driveletter=A -unit=K -size=250 -name=Part1

#Montar la Particion Creada Para Trabajar con ella
mount -driveletter=A -name=Part1

#Ver las Particiones Montadas
mount

#Formatear la Particion con un sistema de Archivos EXT2
#Utilizar el Id que aparece en el comando anterior
mkfs -id=A116 -fs=2fs

#Crear un Reporte para ver Graficamente el estado actual del disco Creado
rep -name=disk -path=ReporteDisk -id=A 

rep -name=tree -path=ReporteTree -id=A116
```
## Manuales

* Para aprender sobre los **comandos** disponibles y la **cli**, vea [User Manual](/manuals/USER_MANUAL.md)

* Para aprender acerca de las **estructuras** que se utilizan para simular: la tabla de particiones MBR, las particiones, los sistemas de archivos y el manejo de archivos, vea [Technical Manual](/manuals/TECHNICAL_MANUAL.md)