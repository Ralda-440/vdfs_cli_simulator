# USER MANUAL

* ### [Ir a GUI de la aplicación](#gui-de-la-aplicación)
* ### [Ir a Comandos disponibles para la cli](#comandos-disponibles-para-la-cli)


## GUI de la aplicación

* ### CLI
    ![cli](/manuals/images/cli.png)

    * **Consola**: aquí es donde se ingresan los comandos. Soporta múltiples comandos a la vez separados por salto de lineá. Presione **Enter** para ejecutar el/los comando(s). 

    * **Reset**: En la parte superior se encuentra el botón para revertir todos los comandos ingresados y volver a un estado como si nunca se hubiera ejecutado ningún comando.

* ### Explorer

    ![explorer](/manuals/images/explorer.png)

    Puede navegar en todos los discos, particiones y sistemas de archivos que se han creado. 

    Para acceder a las particiones de un disco es necesario formatear la partición de un disco con un sistema de archivos. En el sistema de archivos se creara automáticamente un archivo con lo usuarios que pueden acceder a la partición. Por default se crea un superusuario con usuario: **root** y password: **123**. 
* ### Reportes

    ![reportes](/manuals/images/reportes.png)

    Puede ver todos los tipos de reportes que se han creado, y con un doble clic se mostrara el reporte en una nueva ventana del navegador.

## Comandos disponibles para la cli

### Administración de discos

Estos comandos permitirán crear archivos que simularán discos duros en los que se
podrá crear particiones y formatear con el sistema de archivos ext2 o ext3.

* ### mkdisk

    Este comando creará un archivo binario que simulará un disco, se crearan con una extensión **.dsk** y su contenido serán 0x00 (00000000b en binario), los necesarios para obtener al tamaño del solicitado. Estos archivos se crearan en el **Backend**

    Utiliza los siguientes parámetros:

    |Parámetro|Categoría|Descripción|
    |-|-|-|
    |-size|**Obligatorio**|Este parámetro recibirá un número que indicará el tamaño del disco a crear. Debe ser positivo y mayor que cero.|
    |-fit|Optcional|Indicará el ajuste que utilizará el disco para crear las particiones dentro del disco Podrá tener los siguientes valores: <br><br>BF: Indicará el mejor ajuste (Best Fit).<br>FF: Utilizará el primer ajuste (First Fit). <br>WF: Utilizará el peor ajuste (Worst Fit).<br><br>Ya que es opcional, se tomará el primer ajuste (FF) si no está especificado en el comando.|
    |-unit|Opcional|Este parámetro recibirá una letra que indicará las unidades que utilizará el parámetro size. Podrá tener los siguientes valores:<br>K: Indicará que se utilizarán Kilobytes (1024 bytes)<br>M: Indicará que se utilizarán Megabytes (1024 * 1024 bytes)<br>Este parámetro es opcional, si no se encuentra se creará un disco con tamaño en Megabytes.|

* ### rmdisk

    Este parámetro elimina un archivo binario que representa a un disco. 

    Utiliza los siguientes Parámetros

    |Parámetro|Categoría|Descripción|
    |-|-|-|
    |-driveletter|**Obligatorio**|Este parámetro será la letra del disco a eliminar|

* ### fdisk

    Este comando administra las particiones en el archivo que representa el disco.

    |Parámetro|Categoría|Descripción|
    |-|-|-|
    |-size|**Obligatorio al crear una partición**|Este parámetro recibirá un número que indicará el tamaño de la partición a crear. Debe ser positivo y mayor a cero|
    |-driveletter|**Obligatorio**|Este parámetro será la letra del disco donde se aplicara el comando.|
    |-name|**Obligatorio**|Indicará el nombre de la partición.El nombre no debe repetirse dentro de las particiones de cada disco. Si se va a eliminar, la partición ya debe existir.|
    |-unit|Opcional|Este parámetro recibirá una letra que indicará las unidades que utilizará el parámetro s. Podrá tener los siguientes valores:<br>B: indicará que se utilizarán bytes.<br>K: indicará que se utilizarán Kilobytes(1024 bytes).<br>M: indicará que se utilizarán Megabytes(1024 * 1024 bytes).<br>Este parámetro es opcional, si no se encuentra se creará una partición en Kilobytes.|
    |-type|Opcional|Indicará que tipo de partición se creará. Ya que es opcional, se tomará como primaria en caso de que no se indique. Podrá tener los siguientes valores:<br>P: Se creará una partición primaria.<br>E: Se creará una partición extendida.<br>L: Se creará una partición lógica.<br><br>Las particiones lógicas sólo pueden estar dentro de la extendida sin sobrepasar su tamaño. Deberá tener en cuenta las restricciones de teoría de particiones:<br><br>●  La suma de primarias y extendidas debe ser como máximo 4.<br>● Solo puede haber una partición extendida por disco.<br>● No se puede crear una partición lógica si no hay una extendida|
    |-fit|Opcional|Indicará el ajuste que utilizará la partición para asignar espacio. Podrá tener los siguientes valores:<br><br>BF: Indicará el mejor ajuste (Best Fit)<br>FF: Utilizará el primer ajuste (First Fit)<br>WF: Utilizará el peor ajuste (Worst Fit)<br><br>Ya que es opcional, se tomará el peor ajuste (WF) si no está especificado en el comando|
    |-delete|Opcional|Este parámetro indica que se eliminará una partición. Este parámetro se utiliza junto con **-name**.<br>Recibirá el único siguiente valor:<br><br>**Full**: Esta opción además marcar como vació el espacio en la tabla de particiones, rellena el espacio con el carácter \0. Si se utiliza otro valor diferente, mostrará un mensaje de error|
    |-add|Opcional|Este parámetro se utilizará para agregar o quitar espacio de la partición. Puede ser positivo o negativo. Tomará el parámetro **-unit** para las unidades a agregar o eliminar|


* ### mount

    Este comando montará una partición del disco en el sistema.
    Si no se utiliza ningún parámetro, se mostraran todas las particiones montadas en el sistema y el **id** que se le asignó.

    Utiliza los siguientes parámetros:

    |Parámetro|Categoría|Descripción|
    |-|-|-|
    |-driveletter|**Obligatorio**|Sera la letra del disco donde se encuentra la partición que se desea montar.|
    |-name|**Obligatorio**|Indica el nombre de la partición a montar.|

* ### unmount 

    Desmonta una partición el sistema. Se utiliza el **id** que se le asignó a la partición al momento de montarla en el sistema.

    Utiliza el siguiente parámetro:

    |Parámetro|Categoría|Descripción|
    |-|-|-|
    |-id|**Obligatorio**|Especifica el **id** de la partición que se desmontará.|

* ### mkfs

    Este comando realiza un formateo completo de la partición, se formatea como **ext2** por defecto si en caso el parámetro **-fs** no está definido. También creará un archivo en la raíz llamado **users.txt** que tendrá los usuarios y contraseñas del sistema de archivos. Estos usuarios son con los que se puede acceder a la partición.

    |Parámetro|Categoría|Descripción|
    |-|-|-|
    |-id|**Obligatorio**|Indicará el id que se generó con el comando **mount**|
    |-type|Opcional|Indicará que tipo de formateo se realizará. Podrá tener los siguientes valores:<br><br>**Full**: en este caso se realizará un formateo completo. Ya que es opcional, se tomará como un formateo completo si no se especifica esta opción|
    |-fs|Opcional|Indica el sistema de archivos a formatear. Podrá tener los siguientes valores:<br><br>**2fs**: Para el sistema EXT2<br>**3fs**: Para el sistema EXT3<br><br>En caso no se indique, por defecto será **2fs**|

### Administración de Usuarios Y Grupos

* ### login

    Este comando se utiliza para iniciar sesión en el sistema. No se puede iniciar otra sesión sin haber hecho un **logout** antes.

    Utiliza los siguiente parámetros:

    |Parámetro|Categoría|Descripción|
    |-|-|-|
    |-user|**Obligatorio**|Especifica el nombre del usuario que iniciará sesión.|
    |-pass|**Obligatorio**|Indicará la contraseña del usuario que inicia sesión.|
    |-id|**Obligatorio**|Indicará el id de la partición montada de la cual se va a iniciar sesión. De lograr iniciar sesión todas las acciones se realizarán sobre este id.|

* ### logout

    Este comando se utiliza para cerrar sesión. Debe haber una sesión activa anteriormente para poder utilizarlo, si no, debe mostrar un mensaje de error. Este comando no recibe parámetros.


#### Nota: Todos los siguientes comandos a continuación (a excepción de la sección de Reportes), necesitan que exista una sesión iniciada en el sistema ya que se ejecutan sobre la partición en la que se inicio sesión.

* ### mkgrp

    Este comando creará un grupo para los usuarios de la partición y se guardará en el archivo **users.txt** de la partición, este comando solo lo puede utilizar el usuario **root**. Si el grupo a crear ya existe, se mostrará un
    mensaje de error.

    Utiliza los siguientes parámetros:

    |Parámetro|Categoría|Descripción|
    |-|-|-|
    |-name|**Obligatorio**|Indicará el nombre que tendrá el grupo.|

* ### rmgrp

    Este comando eliminará un grupo para los usuarios de la partición. Solo lo puede utilizar el usuario **root**.

    Utiliza los siguientes parámetros:

    |Parámetro|Categoría|Descripción|
    |-|-|-|
    |-name|**Obligatorio**|Indicará el nombre del grupo a eliminar. Si el grupo no se encuentra dentro de la partición se mostrara un error.|

* ### mkusr

    Este comando crea un usuario en la partición. Solo lo puede ejecutar el usuario root.

    Utiliza los siguientes parámetros:

    |Parámetro|Categoría|Descripción|
    |-|-|-|
    |-user|**Obligatorio**|Indicará el nombre del usuario a crear, si ya existe, mostrará un error indicando que ya existe el usuario.<br>**Máximo: 10 caracteres**|
    |-pass|**Obligatorio**|Indicará la contraseña del usuario<br>**Máximo 10 Caracteres**|
    |-grp|**Obligatorio**|Indicará el grupo al que pertenece el usuario. Debe de existir en la partición en la que se está creando el usuario, si no se mostrará un mensaje de error.<br>**Máximo 10 Caracteres**|

* ### rmusr

    Este comando elimina un usuario en la partición. Solo lo puede ejecutar el usuario root.

    Utiliza los siguientes parámetros:

    |Parámetro|Categoría|Descripción|
    |-|-|-|
    |-user|**Obligatorio**|Indicará el nombre del usuario a eliminar. Si el usuario no se encuentra dentro de la partición, mostrará un error.|

### Administración de Carpetas, Archivos y Permisos

Estos comandos permitirán crear archivos y carpetas, así como editarlos, copiarlos, moverlos y eliminarlos. Los permisos serán para el usuario propietario del archivo, para el grupo al que pertenece y para otros usuarios, así como en Linux.

* ### mkfile

    Este comando permitirá crear un archivo, el propietario será el usuario que
    actualmente ha iniciado sesión. Tendrá los permisos **664**. Por temas de simplificar la aplicación, la creación de archivos se hará con contenido creado automáticamente. El contenido serán números del 0 a 9 cuantas veces sea necesario para cumplir el tamaño ingresado. 

    Utiliza los siguientes parámetros:

    |Parámetro|Categoría|Descripción|
    |-|-|-|
    |-path|**Obligatorio**|Este parámetro será la ruta del archivo que se creará|
    |-r|Opcional|Si se utiliza este parámetro y las carpetas especificadas por el parámetro **-path** no existen, entonces se crearan las carpetas padres. Si ya existen, no se crearan las carpetas.|
    |-size|Opcional|Este parámetro indicará el tamaño en bytes del archivo, El contenido serán números del 0 al 9 cuantas veces sea necesario hasta cumplir el tamaño ingresado. Si no se utiliza este parámetro, el tamaño será 0 bytes.|


* ### cat

    Este comando permitirá mostrar el contenido del archivo, si el usuario que actualmente está logueado tiene acceso al permiso de lectura.

    Utiliza los siguientes parámetros:

    |Parámetro|Categoría|Descripción|
    |-|-|-|
    |-fileN|**Obligatorio**|Permitirá admitir como argumentos una lista de N ficheros que hay que enlazar. Estos se deben encadenar en el mismo orden en el cual fueron especificados<br><br>Ejemplo:<br><br>**cat** -file1="/home/a.txt" -file2="/home/b.txt" -file3="/home/c.txt"|

* ### remove

    Este comando permitirá eliminar un archivo o carpeta y todo su  contenido, si el usuario que actualmente está logueado tiene acceso al permiso de escritura sobre el archivo y en el caso de carpetas, eliminará todos los archivos o subcarpetas en los que el usuario tenga permiso de escritura.

    Utiliza los siguientes parámetros:

    |Parámetro|Categoría|Descripción|
    |-|-|-|
    |-path|**Obligatorio**|Este parámetro será la ruta del archivo o carpeta que se eliminará. Si lleva espacios en blanco deberá encerrarse entre comillas. Si no existe el archivo o no tiene permisos de escritura en la carpeta o en el archivo, se mostrará un mensaje de error. Si no se puedo eliminar algún archivo o carpeta no se eliminará los padres.|

* ### rename

    Este comando permitirá cambiar el nombre de un archivo o carpeta, si el usuario actualmente logueado tiene permiso de escritura sobre el archivo o carpeta.

    Utiliza los siguientes parámetros:

    |Parámetro|Categoría|Descripción|
    |-|-|-|
    |-path|**Obligatorio**|Este parámetro será la ruta del archivo o carpeta al que se le cambiará el nombre. Si lleva espacios en blanco deberá encerrarse entre comillas. Si no existe el archivo o carpeta o no tiene permisos de escritura se mostrará un mensaje de error|
    |-name|**Obligatorio**|Especificará el nuevo nombre del archivo, se verificará que no exista un archivo con el mismo nombre, de ser así se mostrará un mensaje de error|

* ### mkdir

    Este comando es similar a mkfile, pero no crea archivos, sino carpetas. El propietario será el usuario que actualmente ha iniciado sesión. Tendrá los permisos 664. El usuario deberá tener el permiso de escritura en la carpeta padre.

    Utiliza los siguientes parámetros:

    |Parámetro|Categoría|Descripción|
    |-|-|-|
    |-path|**Obligatorio**|Este parámetro será la ruta de la carpeta que se creará. Si lleva espacios en blanco deberá encerrarse entre comillas|
    |-r|Opcional|Si se utiliza este parámetro y las carpetas padres en el parámetro **-path** no existen, entonces se crearan. No recibirá ningún valor|

* ### move

    Este comando moverá un archivo o carpeta y todo su contenido hacia otro destino.

    Utiliza los siguientes parámetros:

    |Parámetro|Categoría|Descripción|
    |-|-|-|
    |-path|**Obligatorio**| Este parámetro será la ruta del archivo o carpeta que se desea mover. Si lleva espacios en blanco deberá encerrarse entre comillas.|
    |-destino|**Obligatorio**|Este parámetro será la ruta de la carpeta a la que se moverá el archivo o carpeta. Debe tener permiso de escritura sobre la carpeta. Si lleva espacios en blanco deberá encerrarse entre comillas|

### Reportes

Recibirá el nombre del reporte que se desea y agregara a la Sección de Reportes de la GUI.

Utiliza los siguientes parámetros:

|Parámetro|Categoría|Descripción|
|-|-|-|
|-name|**Obligatorio**|Nombre del reporte a generar. Tendrá los siguientes valores:<br>● mbr<br>● disk<br>● inode<br>● block<br>● bm_inode<br>● bm_block<br>● tree<br>● sb<br>● file<br>● ls|
|-path|**Obligatorio**|Indica el nombre que tendrá el reporte|
|-id|**Obligatorio**| Indica el id de la partición que se utilizará. Si el reporte es sobre la información del disco, solo es necesario indicar la **Letra** del disco|
|-ruta|Opcional|Funcionará para el reporte **file** y **ls**. Será el nombre del archivo o carpeta del que se mostrará el reporte.|

#### Tipos de Reportes

* #### mbr
    Mostrará tablas con toda la información del MBR, así como de los EBR que se pudieron haber creado.

* #### disk
    Este reporte mostrará la estructura de las particiones, el mbr del disco y el porcentaje que cada partición o espacio libre tiene dentro del disco.

* #### inode
    Mostrará bloques con toda la información de los inodos utilizados. Si no están utilizados no se mostraran.

* #### block
    Mostrará la información de todos los bloques utilizados. Si no están utilizados no se mostraran.

* #### bm_inode
    Este reporte mostrará la información del bitmap de inodos, mostrará todos los bits, libres o utilizados. Este reporte se generará en un archivo de texto mostrando 20 registros por línea.

* #### bm_block
    Este reporte mostrará la información del bitmap de inodos, desplegará todos los bits, libres o utilizados. Este reporte se generará en un archivo de texto que mostrará 20 registros por línea.

* #### tree
    Este reporte genera el árbol de todo el sistema ext2/ext3. Se mostrará toda la información de los inodos o bloques.

* #### sb
    Muestra toda la información del superbloque en una tabla.

* #### file
    Este reporte muestra el contenido del archivo especificado en el parámetro **-ruta**.

* #### ls 
    Este reporte mostrará la información de los archivos y carpetas con permisos, propietario, grupo propietario, fecha de modificación, hora de modificación, tipo, fecha de creación.