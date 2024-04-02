grammar FileSysCLI;
//import FileSysCLILex;

commands: command* EOF;

command 
    : mkdisk
    | rmdisk
    | fdisk
    | execute
    | mount
    | unmount
    | mkfs
    | login
    | logout
    | mkgrp
    | rmgrp
    | mkusr
    | rmusr
    | mkfile
    | cat
    | remove
    | edit
    | rename
    | mkdir
    | chgrp
    | rep
    | move
    | loss
    | recovery
    | chown
    | chmod
    | pause
;

execute
    : 'execute' parametro* #Command_execute
;

mkdisk
    : 'mkdisk' parametro* #Command_mkdisk 
;

rmdisk
    : 'rmdisk' parametro* #Command_rmdisk
;

fdisk
    : 'fdisk' parametro* #Command_fdisk
;

mount: 'mount' parametro* #Command_mount
;

unmount: 'unmount' parametro* #Command_unmount
;

mkfs : 'mkfs' parametro* #Command_mkfs
;

login: 'login' parametro* #Command_login
;

logout: 'logout' parametro* #Command_logout
;

mkgrp: 'mkgrp' parametro* #Command_mkgrp
;

rmgrp: 'rmgrp' parametro* #Command_rmgrp
;

mkusr: 'mkusr' parametro* #Command_mkusr
;

rmusr: 'rmusr' parametro* #Command_rmusr
;

mkfile: 'mkfile' parametro* #Command_mkfile
;

rep: 'rep' parametro* #Command_rep
;

cat: 'cat' parametro* #Command_cat
;

remove : 'remove' parametro* #Command_rm
; 

edit : 'edit' parametro* #Command_edit
;

rename : 'rename' parametro* #Command_rename
;

mkdir : 'mkdir' parametro* #Command_mkdir
;

pause : 'pause' parametro* #Command_pause
;

chgrp : 'chgrp' parametro* #Command_chgrp
;

move : 'move' parametro* #Command_move
;

loss: 'loss' parametro* #Command_loss
;

recovery: 'recovery' parametro* #Command_recovery
; 

chown: 'chown' parametro* #Command_chown
;

chmod: 'chmod' parametro* #Command_chmod
;

parametro: QUOTED_TEXT_ ('=' value_parametro)? #Parameter
;

//name_parametro: parameter=( '-size='  | '-fit=' | '-unit=' | '-path=')  #Name_parameter
//;

value_parametro: (QUOTED_TEXT_ | STRING) #Value_parameter
;

//IGNORAR
ENBLANCO:
	[ \t\n\r]+ -> skip; //Se ignoran todos los espacios en blanco
LINE_COMMENT:
	'#' ~[\r\n]* -> skip; //Comentario Simple

//fragmentos
//fragment DIG: [0-9];
//fragment LETRA: [a-zA-Z];

//Expresiones
//DECIMAL: DIG+ '.' DIG+;
//ENTERO: DIG+;
STRING: '"' QUOTED_TEXT? '"';
fragment QUOTED_TEXT: QUOTED_TEXT_ITEM+;
fragment QUOTED_TEXT_ITEM: ~["\n\r\t];
QUOTED_TEXT_: QUOTED_TEXT_ITEM_+;
fragment QUOTED_TEXT_ITEM_: ~[ ="\n\r\t];
//