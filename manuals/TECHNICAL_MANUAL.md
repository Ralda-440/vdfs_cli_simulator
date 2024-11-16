# TECHNICAL MANUAL



## Estructura MBR

Cuando se crea un nuevo disco este contiene un MBR que provee de información de las particiones, este se encuentra en el
primer sector del disco.

|Atributo|Tipo|Descripción|
|-|-|-|
|mbr_tamano|int|Tamaño total del disco en bytes|
|mbr_fecha_creacion|time|Fecha y hora de creación del disco|
|mbr_dsk_signature|int|Número random, que identifica de forma única a cada disco|
|dsk_fit|char|Tipo de ajuste de la partición. Tendrá los valores B (Best), F (First) o W (worst)|
|mbr_partitions|partition[4]|Estructura con información de las 4 particiones|

#### Implementación en Go

```go
type MBR struct {
	Mbr_tamanio        int64        //Tamaño del disco en bytes
	Mbr_fecha_creacion [19]byte     //00/00/0000 00:00:00
	Mbr_disk_signature int64        //Numero random
	Mbr_disk_fit       byte         // B, F, W
	Particiones        [4]Particion //Particiones
}
```

## Estructura Partition
Una partición es una división lógica de un disco, en la cual se alojan y organizan los archivos mediante un sistema de archivos.

|Atributo|Tipo|Descripción|
|-|-|-|
|part_status|char|Indica si la partición está montada o no.|
|part_type|char|Indica el tipo de partición, primaria o extendida. Tendrá los valores P o E|
|part_fit|char|Tipo de ajuste de la partición. Tendrá los valores B (Best), F (First) o W (worst)|
|part_start|int|Indica en qué byte del disco inicia la partición|
|part_s|int|Contiene el tamaño total de la partición en bytes|
|part_name|char[16]|Nombre de la partición|
|part_correlative|int|Indica el correlativo de la partición|
|part_id|char[4]|Indica el ID de la partición generada al montar esta partición, esto se explicará más adelante|

#### Implementación en go

```go
type Particion struct {
	Part_status      byte     // 0 = Inactiva, 1 = Activa
	Part_type        byte     // P, E
	Part_fit         byte     // B, F, W
	Part_start       int64    // Byte de inicio
	Part_s           int64    // Tamaño de la partición
	Part_name        [16]byte // Nombre de la partición
	Part_correlative int64    // Correlativo de la partición
	Part_id          [4]byte  // Identificador de la partición
}
```

## Estructura EBR

es un descriptor para una partición lógica bajo el sistema común de particionamiento de unidades de disco. Si otra partición lógica le sigue, entonces el primer EBR contendrá una entrada apuntando al siguiente EBR; por lo tanto, múltiples EBR forman una Lista enlazada.

|Atributo|Tipo|Descripción|
|-|-|-|
|part_mount|char|Indica si la partición está montada o no|
|part_fit|char|Tipo de ajuste de la partición. Tendrá los valores B (Best), F (First) o W (worst)|
|part_start|int|Indica en qué byte del disco inicia la partición|
|part_s|int|Contiene el tamaño total de la partición en bytes|
|part_next|int|Byte en el que está el próximo EBR. -1 si no hay siguiente|
|part_name|char[16]|Nombre de la partición|

#### Implementación en go

```go
type EBR struct {
	Part_mount byte     // 0 = Inactiva, 1 = Activada
	Part_fit   byte     // B, F, W
	Part_start int64    // Byte de inicio
	Part_s     int64    // Tamaño de la partición
	Part_next  int64    // Byte de inicio de la siguiente partición
	Part_name  [16]byte // Nombre de la partición
}
```

## Estructura Super Bloque

Esta estructura contiene la información sobre el sistema de archivos de la partición. 

|Atributo|Tipo|Descripción|
|-|-|-|
| s_filesystem_type  | int   | Guarda el número que identifica el sistema de archivos utilizado      |
| s_inodes_count     | int   | Guarda el número total de inodos                                      |
| s_blocks_count     | int   | Guarda el número total de bloques                                     |
| s_free_blocks_count| int   | Contiene el número de bloques libres                                  |
| s_free_inodes_count| int   | Contiene el número de inodos libres                                   |
| s_mtime            | time  | Última fecha en el que el sistema fue montado                         |
| s_umtime           | time  | Última fecha en que el sistema fue desmontado                         |
| s_mnt_count        | int   | Indica cuántas veces se ha montado el sistema                         |
| s_magic            | int   | Valor que identifica al sistema de archivos, tendrá el valor 0xEF53   |
| s_inode_s          | int   | Tamaño del inodo                                                      |
| s_block_s          | int   | Tamaño del bloque                                                     |
| s_firts_ino        | int   | Primer inodo libre                                                    |
| s_first_blo        | int   | Primer bloque libre                                                   |
| s_bm_inode_start   | int   | Guardará el inicio del bitmap de inodos                               |
| s_bm_block_start   | int   | Guardará el inicio del bitmap de bloques                              |
| s_inode_start      | int   | Guardará el inicio de la tabla de inodos                              |
| s_block_start      | int   | Guardará el inicio de la tabla de bloques                             |


#### Implementacion en Go

```go
type Super_Bloque struct {
	S_filesystem_type   int64    // 0 = EXT2, 1 = EXT3
	S_inodes_count      int64    // Cantidad total inodos
	S_blocks_count      int64    // Cantidad total bloques
	S_free_blocks_count int64    // Cantidad bloques libres
	S_free_inodes_count int64    // Cantidad inodos libres
	S_mtime             [19]byte // Ultima fecha de montaje
	S_umtime            [19]byte // Ultima fecha de desmontaje
	S_mnt_count         int64    // Cantidad de montajes
	S_magic             int64    // Numero magico 0xEF53
	S_inode_s           int64    // Tamaño de un inodo
	S_block_s           int64    // Tamaño de un bloque
	S_first_ino         int64    // Primer inodo libre
	S_first_blo         int64    // Primer bloque libre
	S_bm_inode_start    int64    // Inicio del bitmap de inodos
	S_bm_block_start    int64    // Inicio del bitmap de bloques
	S_inode_start       int64    // Inicio de la tabla de inodos
	S_block_start       int64    // Inicio de la tabla de bloques
}
```

## Estructura Inodo

Esta estructura contiene las características e información sobre un fichero usado por una carpeta o archivos.

| Atributo   | Tipo     | Descripción                                                                                                                   |
|----------|----------|-------------------------------------------------------------------------------------------------------------------------------|
| i_uid    | int      | UID del usuario propietario del archivo o carpeta                                                                             |
| i_gid    | int      | GID del grupo al que pertenece el archivo o carpeta                                                                           |
| i_s      | int      | Tamaño del archivo en bytes                                                                                                    |
| i_atime  | time     | Última fecha en que se leyó el inodo sin modificarlo                                                                          |
| i_ctime  | time     | Fecha en la que se creó el inodo                                                                                               |
| i_mtime  | time     | Última fecha en la que se modifica el inodo                                                                                   |
| i_block  | int      | Array en los que los primeros 12 registros son bloques directos. El 13 será el número del bloque simple indirecto. El 14 será el número del bloque doble indirecto. El 15 será el número del bloque triple indirecto. Si no son utilizados tendrá el valor: -1. |
| i_type   | char     | Indica si es archivo o carpeta. Tendrá los siguientes valores: 1 = Archivo, 0 = Carpeta                                       |
| i_perm   | char[3]  | Guardará los permisos del archivo o carpeta. Se trabajarán usando los permisos UGO (User Group Other) en su forma octal.      |

#### Implementacion en Go

```go
type Inodes struct {
	I_uid   int64     // Identificador del usuario propietario del archivo o carpeta
	I_gid   int64     // Identificador del grupo propietario del archivo o carpeta
	I_s     int64     // Tamaño del archivo en bytes
	I_atime [19]byte  // Ultima fecha lectura sin modificación
	I_ctime [19]byte  // Fecha de creación formato: YYYY/MM/DD HH:MM:SS
	I_mtime [19]byte  // Ultima fecha modificación
	I_block [15]int32 // Bloques directos e indirectos
	I_type  byte      // 0 = Carpeta, 1 = Archivo
	I_perm  [3]byte   // Permisos
}
```

## Estructuras de Bloques

Estas estructuras son la unidad mínima de almacenamiento a nivel lógico. Estos bloques son un conjunto de sectores contiguos que componen la unidad de almacenamiento más pequeña de un disco, para este proyecto los bloques tienen un tamaño de 64 bytes.

* ### Estructura Bloque de Carpetas

	Esta estructura guardará la información sobre el nombre de de los archivos que contiene y a que Inodo apuntan.

	|Atributo|Tipo|Descripción|
	|-|-|-|
	|b_content|content[4]|Array con el contenido de la carpeta|

	La estructura **b_content** tendrá los siguientes atributos:

	|Atributo|Tipo|Descripción|
	|-|-|-|
	|b_name|char[12]|Nombre de la carpeta o archivo|
	|b_inodo|int|Apuntador hacia un Inodo asociado al archivo o carpeta|


	#### Implementación en Go

	```go
	type Block_Carpeta struct {
		B_content [4]B_content //4 Apuntadores a Carpetas o Archivos
	}

	type B_content struct {
		B_name  [12]byte // Nombre del archivo o carpeta
		B_inodo int32    // Apuntador al inodo del archivo o carpeta
	}
	```
* ### Estructura Bloque de Archivo

	Esta estructura guardará la información sobre contenido de un archivo.

	|Atributo|Tipo|Descripción|
	|-|-|-|
	|b_content|char[64]|Array con el contenido del archivo|

	#### Implementación en Go

	```go
	type Block_Archivo struct {
		B_content [64]byte //64 bytes de contenido
	}
	```

* ### Estructura Bloque de Apuntadores

	Esta estructura guardará la información de los apuntadores indirectos (simples,
	dobles y triples)

	|Atributo|Tipo|Descripción|
	|-|-|-|
	|b_pointers|int[16]|Array con los apuntadores a bloques (de archivo, carpeta o de apuntadores)|

	#### Implementación en Go

	```go
	type Block_Indirecto struct {
		B_pointers [16]int32
	}
	```