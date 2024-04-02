package structs

import (
	"bytes"
	"strconv"
)

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

// Graphviz de Inodes
func (i *Inodes) GetGraph(numInodo int, nombre string, soloInodo bool, owner string, grupo string) (string, string) {
	filas := ""
	for pos, content := range i.I_block {
		filas += `
		<tr>
			<td>I_block_` + strconv.Itoa(pos) + `</td><td port="f` + strconv.Itoa(pos) + `">` + strconv.Itoa(int(content)) + `</td>
		</tr>
		`
	}
	graphviz := `
	Inodo` + strconv.Itoa(numInodo) + ` [label=< 
            <table border="1" cellspacing="0" >
				<tr>
                    <td colspan="2" bgcolor="#0000CC"><font color="white"><b>` + nombre + `</b></font></td>
                </tr>
                <tr>
                    <td colspan="2" bgcolor="#0000CC"><font color="white"><b>Inodo ` + strconv.Itoa(numInodo) + `</b></font></td>
                </tr>
                <tr>
                    <td>I_uid</td><td>` + strconv.Itoa(int(i.I_uid)) + `</td>
                </tr>
                <tr>
                    <td>I_gid</td><td>` + strconv.Itoa(int(i.I_gid)) + `</td>
                </tr>
                <tr>
                    <td>I_s</td><td>` + strconv.Itoa(int(i.I_s)) + `</td>
                </tr>
				<tr>
                    <td>I_atime</td><td>` + string(bytes.Replace(i.I_atime[:], []byte{0}, []byte{}, -1)) + `</td>
                </tr>
				<tr>
                    <td>I_ctime</td><td>` + string(bytes.Replace(i.I_ctime[:], []byte{0}, []byte{}, -1)) + `</td>
                </tr>
				<tr>
                    <td>I_mtime</td><td>` + string(bytes.Replace(i.I_mtime[:], []byte{0}, []byte{}, -1)) + `</td>
                </tr>
				` + filas + `
				<tr>
                    <td>I_type</td><td>` + strconv.Itoa(int(i.I_type)) + `</td>
                </tr>
				<tr>
                    <td>I_perm</td><td>` + strconv.Itoa(int(i.I_perm[0])) + strconv.Itoa(int(i.I_perm[1])) + strconv.Itoa(int(i.I_perm[2])) + `</td>
                </tr>
                </table> 
            > shape=box style=invisible ]
	`
	if !soloInodo {
		//Conexiones
		for pos, content := range i.I_block {
			if content != -1 {
				graphviz += `Inodo` + strconv.Itoa(numInodo) + `:f` + strconv.Itoa(pos) + " -> Bloque" + strconv.Itoa(int(content)) + "\n"
			}
		}
	}
	//Reporte Ls
	tipo := "Carpeta"
	if i.I_type == 1 {
		tipo = "Archivo"
	}
	graphvizLs := `
	<tr>
	<td>` + i.ConvertirPermiso() + `</td>
	<td>` + owner + `</td>
	<td>` + grupo + `</td>
	<td>` + strconv.FormatInt(i.I_s, 10) + `</td>
	<td>` + string(i.I_ctime[:10]) + `</td>
	<td>` + string(i.I_ctime[11:]) + `</td>
	<td>` + tipo + `</td>
	<td>` + nombre + `</td>
	</tr>
	`
	return graphviz, graphvizLs
}

// Convertir Permiso de [3]byte a -rwxrwxrwx
// tipo = 0 -> Carpeta | tipo = 1 -> Archivo
func (inodo *Inodes) ConvertirPermiso() string {
	//Convertir a string
	permisoStr := ""
	//Convertir a string
	for i := 0; i < 3; i++ {
		//Convertir a string
		permisoStr += inodo.ConvertirByteAString(inodo.I_perm[i])
	}
	//Agregar tipo
	if inodo.I_type == 0 {
		permisoStr = "d" + permisoStr
	} else {
		permisoStr = "-" + permisoStr
	}
	return permisoStr
}

// Convertir byte a string
func (inodo *Inodes) ConvertirByteAString(b byte) string {
	//Convertir a string
	permisoStr := ""
	//Convertir a string
	for i := 0; i < 3; i++ {
		//Convertir a string
		if b&1 == 1 && i == 0 {
			permisoStr = "x" + permisoStr
		} else if b&1 == 1 && i == 1 {
			permisoStr = "w" + permisoStr
		} else if b&1 == 1 && i == 2 {
			permisoStr = "r" + permisoStr
		} else {
			permisoStr = "-" + permisoStr
		}
		b = b >> 1
	}
	return permisoStr
}

// Nuevo Inodo
func NewInode() *Inodes {
	inode := Inodes{}
	for i := 0; i < len(inode.I_block); i++ {
		inode.I_block[i] = -1
	}
	return &inode
}
