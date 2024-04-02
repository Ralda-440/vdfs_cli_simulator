package comandos

import (
	structs "app/Interprete/Structs"
	"encoding/binary"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// GenerateRandomSignature genera una firma aleatoria
func generateRandomSignature() int64 {
	timestamp := time.Now().UnixNano()
	randSource := rand.NewSource(time.Now().UnixNano())
	randomValue := rand.New(randSource).Int63n(1000000)
	return timestamp + randomValue
}

// Obtiene el MBR del disco
func getMBRDisk(name string) (*structs.MBR, error) {
	// Leer el MBR del disco
	file, err := os.Open("./MIA/P1/" + name + ".dsk")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Leer el MBR
	mbr := structs.NewMBR()
	//Posicionar el puntero al inicio del archivo
	file.Seek(0, io.SeekStart)
	err = binary.Read(file, binary.BigEndian, &mbr)
	if err != nil {
		return nil, err
	}
	return &mbr, nil
}

// Escribir MBR en el disco
func writeMBRInDisk(driveLetter string, mbr *structs.MBR) error {
	//Abrir el archivo del disco
	file, err := os.OpenFile("./MIA/P1/"+driveLetter+".dsk", os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	//Mover el puntero al inicio del archivo
	file.Seek(0, io.SeekStart)
	//Escribir el MBR
	err = binary.Write(file, binary.BigEndian, mbr)
	if err != nil {
		return err
	}
	return nil
}

// Limpiar un rango de bytes en el disco
func Clear(driveletter string, start int64, size int64) error {
	//Abrir el archivo del disco
	file, err := os.OpenFile("./MIA/P1/"+driveletter+".dsk", os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	//Mover el puntero al inicio de la particion
	file.Seek(start, io.SeekStart)
	//Escribir 0s en la particion
	zeros := make([]byte, size)
	err = binary.Write(file, binary.BigEndian, zeros)
	if err != nil {
		return err
	}
	return nil
}

// User Particion
type UsuarioParticion struct {
	UID     string
	Tipo    string
	Grupo   string
	Usuario string
	Pass    string
}

// Convertir UsuarioParticion a []byte
func (u *UsuarioParticion) ToBytes() []byte {
	//Convertir a []byte
	user := u.UID + "," + u.Tipo + "," + u.Grupo + "," + u.Usuario + "," + u.Pass + "\n"
	//Convertir a []byte
	bytes := []byte(user)
	return bytes
}

// Grupo Particion
type GrupoParticion struct {
	GID   string
	Tipo  string
	Grupo string
}

// Convertir GrupoParticion a []byte
func (g *GrupoParticion) ToBytes() []byte {
	//Convertir a []byte
	group := g.GID + "," + g.Tipo + "," + g.Grupo + "\n"
	//Convertir a []byte
	bytes := []byte(group)
	return bytes
}

// Obtener el super bloque de la particion
func getSuperBloque(ctx *Contexto, driveletter string, partName string) (*Super_Bloque, error) {
	//Obtener MBR del disco donde esta la particion
	mbr, err := getMBRDisk(driveletter)
	if err != nil {
		ctx.AgregarError("Error : El disco no existe o hubo error al leer su MBR :"+err.Error(), 0, 0)
		return nil, err
	}
	//Buscar particion con el id
	existe, particion, _, ebr, _, _, _, err := mbr.BuscarParticion(driveletter, partName)
	if err != nil {
		ctx.AgregarError("Error : No se pudo buscar la particion "+err.Error(), 0, 0)
		return nil, err
	}
	if !existe {
		ctx.AgregarError("Error : No existe la particion con el nombre: "+partName, 0, 0)
		return nil, err
	}
	var start_part int64
	//Obtener el byte de inicio y el tamaño de la particion
	if particion != nil {
		//Es particion primaria
		start_part = particion.Part_start
	} else {
		//Es particion logica
		start_part = ebr.Part_start + int64(binary.Size(structs.EBR{}))
	}
	//Leer el super bloque
	file, err := os.OpenFile("./MIA/P1/"+driveletter+".dsk", os.O_RDONLY, 0644)
	if err != nil {
		ctx.AgregarError("Error : No se pudo abrir el disco :"+err.Error(), 0, 0)
		return nil, err
	}
	defer file.Close()
	//Mover el puntero al inicio del super bloque
	file.Seek(start_part, io.SeekStart)
	//Leer el super bloque
	superBloque := &Super_Bloque{}
	err = binary.Read(file, binary.BigEndian, superBloque)
	if err != nil {
		ctx.AgregarError("Error : No se pudo leer el super bloque :"+err.Error(), 0, 0)
		return nil, err
	}
	return superBloque, nil
}

// Error
type ERROR struct {
	Mensaje string
}

func (e *ERROR) Error() string {
	return e.Mensaje
}

func NewError(mensaje string) *ERROR {
	return &ERROR{Mensaje: mensaje}
}

// Calcular cantidad de bloques necesarios para guardar un archivo
func CalcularCantidadBloquesNecesarios(sizeCont int) int {
	bloques := float64(sizeCont) / float64(64)
	return int(math.Ceil(bloques))
}

// Convertir []UsuarioParticion y []GrupoParticion a []byte
func ConvertirUsuariosGrupos(users []UsuarioParticion, groups []GrupoParticion) []byte { //users []UsuarioParticion, groups []GrupoParticion
	//Convertir usuarios
	usersBytes := make([]byte, 0)
	for _, user := range users {
		//Convertir usuario a []byte
		usuarioBytes := user.ToBytes()
		//Agregar usuario a la lista de usuarios
		usersBytes = append(usersBytes, usuarioBytes...)
	}
	//Convertir grupos
	groupsBytes := make([]byte, 0)
	for _, group := range groups {
		//Convertir grupo a []byte
		grupoBytes := group.ToBytes()
		//Agregar grupo a la lista de grupos
		groupsBytes = append(groupsBytes, grupoBytes...)
	}
	//Unir usuarios y grupos
	usersBytes = append(usersBytes, groupsBytes...)
	return usersBytes
}
