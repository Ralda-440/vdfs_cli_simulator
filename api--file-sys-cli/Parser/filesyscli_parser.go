// Code generated from FileSysCLI.g4 by ANTLR 4.13.1. DO NOT EDIT.

package Parser // FileSysCLI
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type FileSysCLIParser struct {
	*antlr.BaseParser
}

var FileSysCLIParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func filesyscliParserInit() {
	staticData := &FileSysCLIParserStaticData
	staticData.LiteralNames = []string{
		"", "'execute'", "'mkdisk'", "'rmdisk'", "'fdisk'", "'mount'", "'unmount'",
		"'mkfs'", "'login'", "'logout'", "'mkgrp'", "'rmgrp'", "'mkusr'", "'rmusr'",
		"'mkfile'", "'rep'", "'cat'", "'remove'", "'edit'", "'rename'", "'mkdir'",
		"'pause'", "'chgrp'", "'move'", "'loss'", "'recovery'", "'chown'", "'chmod'",
		"'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "ENBLANCO", "LINE_COMMENT",
		"STRING", "QUOTED_TEXT_",
	}
	staticData.RuleNames = []string{
		"commands", "command", "execute", "mkdisk", "rmdisk", "fdisk", "mount",
		"unmount", "mkfs", "login", "logout", "mkgrp", "rmgrp", "mkusr", "rmusr",
		"mkfile", "rep", "cat", "remove", "edit", "rename", "mkdir", "pause",
		"chgrp", "move", "loss", "recovery", "chown", "chmod", "parametro",
		"value_parametro",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 32, 296, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 5,
		0, 64, 8, 0, 10, 0, 12, 0, 67, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1,
		98, 8, 1, 1, 2, 1, 2, 5, 2, 102, 8, 2, 10, 2, 12, 2, 105, 9, 2, 1, 3, 1,
		3, 5, 3, 109, 8, 3, 10, 3, 12, 3, 112, 9, 3, 1, 4, 1, 4, 5, 4, 116, 8,
		4, 10, 4, 12, 4, 119, 9, 4, 1, 5, 1, 5, 5, 5, 123, 8, 5, 10, 5, 12, 5,
		126, 9, 5, 1, 6, 1, 6, 5, 6, 130, 8, 6, 10, 6, 12, 6, 133, 9, 6, 1, 7,
		1, 7, 5, 7, 137, 8, 7, 10, 7, 12, 7, 140, 9, 7, 1, 8, 1, 8, 5, 8, 144,
		8, 8, 10, 8, 12, 8, 147, 9, 8, 1, 9, 1, 9, 5, 9, 151, 8, 9, 10, 9, 12,
		9, 154, 9, 9, 1, 10, 1, 10, 5, 10, 158, 8, 10, 10, 10, 12, 10, 161, 9,
		10, 1, 11, 1, 11, 5, 11, 165, 8, 11, 10, 11, 12, 11, 168, 9, 11, 1, 12,
		1, 12, 5, 12, 172, 8, 12, 10, 12, 12, 12, 175, 9, 12, 1, 13, 1, 13, 5,
		13, 179, 8, 13, 10, 13, 12, 13, 182, 9, 13, 1, 14, 1, 14, 5, 14, 186, 8,
		14, 10, 14, 12, 14, 189, 9, 14, 1, 15, 1, 15, 5, 15, 193, 8, 15, 10, 15,
		12, 15, 196, 9, 15, 1, 16, 1, 16, 5, 16, 200, 8, 16, 10, 16, 12, 16, 203,
		9, 16, 1, 17, 1, 17, 5, 17, 207, 8, 17, 10, 17, 12, 17, 210, 9, 17, 1,
		18, 1, 18, 5, 18, 214, 8, 18, 10, 18, 12, 18, 217, 9, 18, 1, 19, 1, 19,
		5, 19, 221, 8, 19, 10, 19, 12, 19, 224, 9, 19, 1, 20, 1, 20, 5, 20, 228,
		8, 20, 10, 20, 12, 20, 231, 9, 20, 1, 21, 1, 21, 5, 21, 235, 8, 21, 10,
		21, 12, 21, 238, 9, 21, 1, 22, 1, 22, 5, 22, 242, 8, 22, 10, 22, 12, 22,
		245, 9, 22, 1, 23, 1, 23, 5, 23, 249, 8, 23, 10, 23, 12, 23, 252, 9, 23,
		1, 24, 1, 24, 5, 24, 256, 8, 24, 10, 24, 12, 24, 259, 9, 24, 1, 25, 1,
		25, 5, 25, 263, 8, 25, 10, 25, 12, 25, 266, 9, 25, 1, 26, 1, 26, 5, 26,
		270, 8, 26, 10, 26, 12, 26, 273, 9, 26, 1, 27, 1, 27, 5, 27, 277, 8, 27,
		10, 27, 12, 27, 280, 9, 27, 1, 28, 1, 28, 5, 28, 284, 8, 28, 10, 28, 12,
		28, 287, 9, 28, 1, 29, 1, 29, 1, 29, 3, 29, 292, 8, 29, 1, 30, 1, 30, 1,
		30, 0, 0, 31, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 0, 1, 1, 0,
		31, 32, 319, 0, 65, 1, 0, 0, 0, 2, 97, 1, 0, 0, 0, 4, 99, 1, 0, 0, 0, 6,
		106, 1, 0, 0, 0, 8, 113, 1, 0, 0, 0, 10, 120, 1, 0, 0, 0, 12, 127, 1, 0,
		0, 0, 14, 134, 1, 0, 0, 0, 16, 141, 1, 0, 0, 0, 18, 148, 1, 0, 0, 0, 20,
		155, 1, 0, 0, 0, 22, 162, 1, 0, 0, 0, 24, 169, 1, 0, 0, 0, 26, 176, 1,
		0, 0, 0, 28, 183, 1, 0, 0, 0, 30, 190, 1, 0, 0, 0, 32, 197, 1, 0, 0, 0,
		34, 204, 1, 0, 0, 0, 36, 211, 1, 0, 0, 0, 38, 218, 1, 0, 0, 0, 40, 225,
		1, 0, 0, 0, 42, 232, 1, 0, 0, 0, 44, 239, 1, 0, 0, 0, 46, 246, 1, 0, 0,
		0, 48, 253, 1, 0, 0, 0, 50, 260, 1, 0, 0, 0, 52, 267, 1, 0, 0, 0, 54, 274,
		1, 0, 0, 0, 56, 281, 1, 0, 0, 0, 58, 288, 1, 0, 0, 0, 60, 293, 1, 0, 0,
		0, 62, 64, 3, 2, 1, 0, 63, 62, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63,
		1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 68, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0,
		68, 69, 5, 0, 0, 1, 69, 1, 1, 0, 0, 0, 70, 98, 3, 6, 3, 0, 71, 98, 3, 8,
		4, 0, 72, 98, 3, 10, 5, 0, 73, 98, 3, 4, 2, 0, 74, 98, 3, 12, 6, 0, 75,
		98, 3, 14, 7, 0, 76, 98, 3, 16, 8, 0, 77, 98, 3, 18, 9, 0, 78, 98, 3, 20,
		10, 0, 79, 98, 3, 22, 11, 0, 80, 98, 3, 24, 12, 0, 81, 98, 3, 26, 13, 0,
		82, 98, 3, 28, 14, 0, 83, 98, 3, 30, 15, 0, 84, 98, 3, 34, 17, 0, 85, 98,
		3, 36, 18, 0, 86, 98, 3, 38, 19, 0, 87, 98, 3, 40, 20, 0, 88, 98, 3, 42,
		21, 0, 89, 98, 3, 46, 23, 0, 90, 98, 3, 32, 16, 0, 91, 98, 3, 48, 24, 0,
		92, 98, 3, 50, 25, 0, 93, 98, 3, 52, 26, 0, 94, 98, 3, 54, 27, 0, 95, 98,
		3, 56, 28, 0, 96, 98, 3, 44, 22, 0, 97, 70, 1, 0, 0, 0, 97, 71, 1, 0, 0,
		0, 97, 72, 1, 0, 0, 0, 97, 73, 1, 0, 0, 0, 97, 74, 1, 0, 0, 0, 97, 75,
		1, 0, 0, 0, 97, 76, 1, 0, 0, 0, 97, 77, 1, 0, 0, 0, 97, 78, 1, 0, 0, 0,
		97, 79, 1, 0, 0, 0, 97, 80, 1, 0, 0, 0, 97, 81, 1, 0, 0, 0, 97, 82, 1,
		0, 0, 0, 97, 83, 1, 0, 0, 0, 97, 84, 1, 0, 0, 0, 97, 85, 1, 0, 0, 0, 97,
		86, 1, 0, 0, 0, 97, 87, 1, 0, 0, 0, 97, 88, 1, 0, 0, 0, 97, 89, 1, 0, 0,
		0, 97, 90, 1, 0, 0, 0, 97, 91, 1, 0, 0, 0, 97, 92, 1, 0, 0, 0, 97, 93,
		1, 0, 0, 0, 97, 94, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 97, 96, 1, 0, 0, 0,
		98, 3, 1, 0, 0, 0, 99, 103, 5, 1, 0, 0, 100, 102, 3, 58, 29, 0, 101, 100,
		1, 0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0,
		0, 0, 104, 5, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 110, 5, 2, 0, 0, 107,
		109, 3, 58, 29, 0, 108, 107, 1, 0, 0, 0, 109, 112, 1, 0, 0, 0, 110, 108,
		1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 7, 1, 0, 0, 0, 112, 110, 1, 0, 0,
		0, 113, 117, 5, 3, 0, 0, 114, 116, 3, 58, 29, 0, 115, 114, 1, 0, 0, 0,
		116, 119, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118,
		9, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 124, 5, 4, 0, 0, 121, 123, 3,
		58, 29, 0, 122, 121, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124, 122, 1, 0,
		0, 0, 124, 125, 1, 0, 0, 0, 125, 11, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0,
		127, 131, 5, 5, 0, 0, 128, 130, 3, 58, 29, 0, 129, 128, 1, 0, 0, 0, 130,
		133, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 13, 1,
		0, 0, 0, 133, 131, 1, 0, 0, 0, 134, 138, 5, 6, 0, 0, 135, 137, 3, 58, 29,
		0, 136, 135, 1, 0, 0, 0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138,
		139, 1, 0, 0, 0, 139, 15, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 145, 5,
		7, 0, 0, 142, 144, 3, 58, 29, 0, 143, 142, 1, 0, 0, 0, 144, 147, 1, 0,
		0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 17, 1, 0, 0, 0,
		147, 145, 1, 0, 0, 0, 148, 152, 5, 8, 0, 0, 149, 151, 3, 58, 29, 0, 150,
		149, 1, 0, 0, 0, 151, 154, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152, 153,
		1, 0, 0, 0, 153, 19, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 155, 159, 5, 9,
		0, 0, 156, 158, 3, 58, 29, 0, 157, 156, 1, 0, 0, 0, 158, 161, 1, 0, 0,
		0, 159, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 21, 1, 0, 0, 0, 161,
		159, 1, 0, 0, 0, 162, 166, 5, 10, 0, 0, 163, 165, 3, 58, 29, 0, 164, 163,
		1, 0, 0, 0, 165, 168, 1, 0, 0, 0, 166, 164, 1, 0, 0, 0, 166, 167, 1, 0,
		0, 0, 167, 23, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 169, 173, 5, 11, 0, 0,
		170, 172, 3, 58, 29, 0, 171, 170, 1, 0, 0, 0, 172, 175, 1, 0, 0, 0, 173,
		171, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 25, 1, 0, 0, 0, 175, 173, 1,
		0, 0, 0, 176, 180, 5, 12, 0, 0, 177, 179, 3, 58, 29, 0, 178, 177, 1, 0,
		0, 0, 179, 182, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0,
		181, 27, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 183, 187, 5, 13, 0, 0, 184,
		186, 3, 58, 29, 0, 185, 184, 1, 0, 0, 0, 186, 189, 1, 0, 0, 0, 187, 185,
		1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 29, 1, 0, 0, 0, 189, 187, 1, 0,
		0, 0, 190, 194, 5, 14, 0, 0, 191, 193, 3, 58, 29, 0, 192, 191, 1, 0, 0,
		0, 193, 196, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195,
		31, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 197, 201, 5, 15, 0, 0, 198, 200,
		3, 58, 29, 0, 199, 198, 1, 0, 0, 0, 200, 203, 1, 0, 0, 0, 201, 199, 1,
		0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 33, 1, 0, 0, 0, 203, 201, 1, 0, 0,
		0, 204, 208, 5, 16, 0, 0, 205, 207, 3, 58, 29, 0, 206, 205, 1, 0, 0, 0,
		207, 210, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209,
		35, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 211, 215, 5, 17, 0, 0, 212, 214,
		3, 58, 29, 0, 213, 212, 1, 0, 0, 0, 214, 217, 1, 0, 0, 0, 215, 213, 1,
		0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 37, 1, 0, 0, 0, 217, 215, 1, 0, 0,
		0, 218, 222, 5, 18, 0, 0, 219, 221, 3, 58, 29, 0, 220, 219, 1, 0, 0, 0,
		221, 224, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223,
		39, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 225, 229, 5, 19, 0, 0, 226, 228,
		3, 58, 29, 0, 227, 226, 1, 0, 0, 0, 228, 231, 1, 0, 0, 0, 229, 227, 1,
		0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 41, 1, 0, 0, 0, 231, 229, 1, 0, 0,
		0, 232, 236, 5, 20, 0, 0, 233, 235, 3, 58, 29, 0, 234, 233, 1, 0, 0, 0,
		235, 238, 1, 0, 0, 0, 236, 234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237,
		43, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 239, 243, 5, 21, 0, 0, 240, 242,
		3, 58, 29, 0, 241, 240, 1, 0, 0, 0, 242, 245, 1, 0, 0, 0, 243, 241, 1,
		0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 45, 1, 0, 0, 0, 245, 243, 1, 0, 0,
		0, 246, 250, 5, 22, 0, 0, 247, 249, 3, 58, 29, 0, 248, 247, 1, 0, 0, 0,
		249, 252, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251,
		47, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 253, 257, 5, 23, 0, 0, 254, 256,
		3, 58, 29, 0, 255, 254, 1, 0, 0, 0, 256, 259, 1, 0, 0, 0, 257, 255, 1,
		0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 49, 1, 0, 0, 0, 259, 257, 1, 0, 0,
		0, 260, 264, 5, 24, 0, 0, 261, 263, 3, 58, 29, 0, 262, 261, 1, 0, 0, 0,
		263, 266, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265,
		51, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 267, 271, 5, 25, 0, 0, 268, 270,
		3, 58, 29, 0, 269, 268, 1, 0, 0, 0, 270, 273, 1, 0, 0, 0, 271, 269, 1,
		0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 53, 1, 0, 0, 0, 273, 271, 1, 0, 0,
		0, 274, 278, 5, 26, 0, 0, 275, 277, 3, 58, 29, 0, 276, 275, 1, 0, 0, 0,
		277, 280, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279,
		55, 1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 281, 285, 5, 27, 0, 0, 282, 284,
		3, 58, 29, 0, 283, 282, 1, 0, 0, 0, 284, 287, 1, 0, 0, 0, 285, 283, 1,
		0, 0, 0, 285, 286, 1, 0, 0, 0, 286, 57, 1, 0, 0, 0, 287, 285, 1, 0, 0,
		0, 288, 291, 5, 32, 0, 0, 289, 290, 5, 28, 0, 0, 290, 292, 3, 60, 30, 0,
		291, 289, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 59, 1, 0, 0, 0, 293, 294,
		7, 0, 0, 0, 294, 61, 1, 0, 0, 0, 30, 65, 97, 103, 110, 117, 124, 131, 138,
		145, 152, 159, 166, 173, 180, 187, 194, 201, 208, 215, 222, 229, 236, 243,
		250, 257, 264, 271, 278, 285, 291,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// FileSysCLIParserInit initializes any static state used to implement FileSysCLIParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFileSysCLIParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FileSysCLIParserInit() {
	staticData := &FileSysCLIParserStaticData
	staticData.once.Do(filesyscliParserInit)
}

// NewFileSysCLIParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFileSysCLIParser(input antlr.TokenStream) *FileSysCLIParser {
	FileSysCLIParserInit()
	this := new(FileSysCLIParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &FileSysCLIParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "FileSysCLI.g4"

	return this
}

// FileSysCLIParser tokens.
const (
	FileSysCLIParserEOF          = antlr.TokenEOF
	FileSysCLIParserT__0         = 1
	FileSysCLIParserT__1         = 2
	FileSysCLIParserT__2         = 3
	FileSysCLIParserT__3         = 4
	FileSysCLIParserT__4         = 5
	FileSysCLIParserT__5         = 6
	FileSysCLIParserT__6         = 7
	FileSysCLIParserT__7         = 8
	FileSysCLIParserT__8         = 9
	FileSysCLIParserT__9         = 10
	FileSysCLIParserT__10        = 11
	FileSysCLIParserT__11        = 12
	FileSysCLIParserT__12        = 13
	FileSysCLIParserT__13        = 14
	FileSysCLIParserT__14        = 15
	FileSysCLIParserT__15        = 16
	FileSysCLIParserT__16        = 17
	FileSysCLIParserT__17        = 18
	FileSysCLIParserT__18        = 19
	FileSysCLIParserT__19        = 20
	FileSysCLIParserT__20        = 21
	FileSysCLIParserT__21        = 22
	FileSysCLIParserT__22        = 23
	FileSysCLIParserT__23        = 24
	FileSysCLIParserT__24        = 25
	FileSysCLIParserT__25        = 26
	FileSysCLIParserT__26        = 27
	FileSysCLIParserT__27        = 28
	FileSysCLIParserENBLANCO     = 29
	FileSysCLIParserLINE_COMMENT = 30
	FileSysCLIParserSTRING       = 31
	FileSysCLIParserQUOTED_TEXT_ = 32
)

// FileSysCLIParser rules.
const (
	FileSysCLIParserRULE_commands        = 0
	FileSysCLIParserRULE_command         = 1
	FileSysCLIParserRULE_execute         = 2
	FileSysCLIParserRULE_mkdisk          = 3
	FileSysCLIParserRULE_rmdisk          = 4
	FileSysCLIParserRULE_fdisk           = 5
	FileSysCLIParserRULE_mount           = 6
	FileSysCLIParserRULE_unmount         = 7
	FileSysCLIParserRULE_mkfs            = 8
	FileSysCLIParserRULE_login           = 9
	FileSysCLIParserRULE_logout          = 10
	FileSysCLIParserRULE_mkgrp           = 11
	FileSysCLIParserRULE_rmgrp           = 12
	FileSysCLIParserRULE_mkusr           = 13
	FileSysCLIParserRULE_rmusr           = 14
	FileSysCLIParserRULE_mkfile          = 15
	FileSysCLIParserRULE_rep             = 16
	FileSysCLIParserRULE_cat             = 17
	FileSysCLIParserRULE_remove          = 18
	FileSysCLIParserRULE_edit            = 19
	FileSysCLIParserRULE_rename          = 20
	FileSysCLIParserRULE_mkdir           = 21
	FileSysCLIParserRULE_pause           = 22
	FileSysCLIParserRULE_chgrp           = 23
	FileSysCLIParserRULE_move            = 24
	FileSysCLIParserRULE_loss            = 25
	FileSysCLIParserRULE_recovery        = 26
	FileSysCLIParserRULE_chown           = 27
	FileSysCLIParserRULE_chmod           = 28
	FileSysCLIParserRULE_parametro       = 29
	FileSysCLIParserRULE_value_parametro = 30
)

// ICommandsContext is an interface to support dynamic dispatch.
type ICommandsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllCommand() []ICommandContext
	Command(i int) ICommandContext

	// IsCommandsContext differentiates from other interfaces.
	IsCommandsContext()
}

type CommandsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandsContext() *CommandsContext {
	var p = new(CommandsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_commands
	return p
}

func InitEmptyCommandsContext(p *CommandsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_commands
}

func (*CommandsContext) IsCommandsContext() {}

func NewCommandsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandsContext {
	var p = new(CommandsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_commands

	return p
}

func (s *CommandsContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandsContext) EOF() antlr.TerminalNode {
	return s.GetToken(FileSysCLIParserEOF, 0)
}

func (s *CommandsContext) AllCommand() []ICommandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICommandContext); ok {
			len++
		}
	}

	tst := make([]ICommandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICommandContext); ok {
			tst[i] = t.(ICommandContext)
			i++
		}
	}

	return tst
}

func (s *CommandsContext) Command(i int) ICommandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICommandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *CommandsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommands(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Commands() (localctx ICommandsContext) {
	localctx = NewCommandsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FileSysCLIParserRULE_commands)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&268435454) != 0 {
		{
			p.SetState(62)
			p.Command()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
		p.Match(FileSysCLIParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Mkdisk() IMkdiskContext
	Rmdisk() IRmdiskContext
	Fdisk() IFdiskContext
	Execute() IExecuteContext
	Mount() IMountContext
	Unmount() IUnmountContext
	Mkfs() IMkfsContext
	Login() ILoginContext
	Logout() ILogoutContext
	Mkgrp() IMkgrpContext
	Rmgrp() IRmgrpContext
	Mkusr() IMkusrContext
	Rmusr() IRmusrContext
	Mkfile() IMkfileContext
	Cat() ICatContext
	Remove() IRemoveContext
	Edit() IEditContext
	Rename() IRenameContext
	Mkdir() IMkdirContext
	Chgrp() IChgrpContext
	Rep() IRepContext
	Move() IMoveContext
	Loss() ILossContext
	Recovery() IRecoveryContext
	Chown() IChownContext
	Chmod() IChmodContext
	Pause() IPauseContext

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_command
	return p
}

func InitEmptyCommandContext(p *CommandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_command
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) Mkdisk() IMkdiskContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdiskContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdiskContext)
}

func (s *CommandContext) Rmdisk() IRmdiskContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRmdiskContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRmdiskContext)
}

func (s *CommandContext) Fdisk() IFdiskContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdiskContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFdiskContext)
}

func (s *CommandContext) Execute() IExecuteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExecuteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExecuteContext)
}

func (s *CommandContext) Mount() IMountContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMountContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMountContext)
}

func (s *CommandContext) Unmount() IUnmountContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnmountContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnmountContext)
}

func (s *CommandContext) Mkfs() IMkfsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfsContext)
}

func (s *CommandContext) Login() ILoginContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoginContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoginContext)
}

func (s *CommandContext) Logout() ILogoutContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogoutContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogoutContext)
}

func (s *CommandContext) Mkgrp() IMkgrpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkgrpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkgrpContext)
}

func (s *CommandContext) Rmgrp() IRmgrpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRmgrpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRmgrpContext)
}

func (s *CommandContext) Mkusr() IMkusrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkusrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkusrContext)
}

func (s *CommandContext) Rmusr() IRmusrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRmusrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRmusrContext)
}

func (s *CommandContext) Mkfile() IMkfileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfileContext)
}

func (s *CommandContext) Cat() ICatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICatContext)
}

func (s *CommandContext) Remove() IRemoveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemoveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemoveContext)
}

func (s *CommandContext) Edit() IEditContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEditContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEditContext)
}

func (s *CommandContext) Rename() IRenameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRenameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRenameContext)
}

func (s *CommandContext) Mkdir() IMkdirContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdirContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdirContext)
}

func (s *CommandContext) Chgrp() IChgrpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChgrpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChgrpContext)
}

func (s *CommandContext) Rep() IRepContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepContext)
}

func (s *CommandContext) Move() IMoveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMoveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMoveContext)
}

func (s *CommandContext) Loss() ILossContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILossContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILossContext)
}

func (s *CommandContext) Recovery() IRecoveryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecoveryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecoveryContext)
}

func (s *CommandContext) Chown() IChownContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChownContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChownContext)
}

func (s *CommandContext) Chmod() IChmodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChmodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChmodContext)
}

func (s *CommandContext) Pause() IPauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPauseContext)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Command() (localctx ICommandContext) {
	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FileSysCLIParserRULE_command)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case FileSysCLIParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Mkdisk()
		}

	case FileSysCLIParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Rmdisk()
		}

	case FileSysCLIParserT__3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(72)
			p.Fdisk()
		}

	case FileSysCLIParserT__0:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(73)
			p.Execute()
		}

	case FileSysCLIParserT__4:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(74)
			p.Mount()
		}

	case FileSysCLIParserT__5:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(75)
			p.Unmount()
		}

	case FileSysCLIParserT__6:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(76)
			p.Mkfs()
		}

	case FileSysCLIParserT__7:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(77)
			p.Login()
		}

	case FileSysCLIParserT__8:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(78)
			p.Logout()
		}

	case FileSysCLIParserT__9:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(79)
			p.Mkgrp()
		}

	case FileSysCLIParserT__10:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(80)
			p.Rmgrp()
		}

	case FileSysCLIParserT__11:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(81)
			p.Mkusr()
		}

	case FileSysCLIParserT__12:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(82)
			p.Rmusr()
		}

	case FileSysCLIParserT__13:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(83)
			p.Mkfile()
		}

	case FileSysCLIParserT__15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(84)
			p.Cat()
		}

	case FileSysCLIParserT__16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(85)
			p.Remove()
		}

	case FileSysCLIParserT__17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(86)
			p.Edit()
		}

	case FileSysCLIParserT__18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(87)
			p.Rename()
		}

	case FileSysCLIParserT__19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(88)
			p.Mkdir()
		}

	case FileSysCLIParserT__21:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(89)
			p.Chgrp()
		}

	case FileSysCLIParserT__14:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(90)
			p.Rep()
		}

	case FileSysCLIParserT__22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(91)
			p.Move()
		}

	case FileSysCLIParserT__23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(92)
			p.Loss()
		}

	case FileSysCLIParserT__24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(93)
			p.Recovery()
		}

	case FileSysCLIParserT__25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(94)
			p.Chown()
		}

	case FileSysCLIParserT__26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(95)
			p.Chmod()
		}

	case FileSysCLIParserT__20:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(96)
			p.Pause()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExecuteContext is an interface to support dynamic dispatch.
type IExecuteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExecuteContext differentiates from other interfaces.
	IsExecuteContext()
}

type ExecuteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecuteContext() *ExecuteContext {
	var p = new(ExecuteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_execute
	return p
}

func InitEmptyExecuteContext(p *ExecuteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_execute
}

func (*ExecuteContext) IsExecuteContext() {}

func NewExecuteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecuteContext {
	var p = new(ExecuteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_execute

	return p
}

func (s *ExecuteContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecuteContext) CopyAll(ctx *ExecuteContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExecuteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecuteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_executeContext struct {
	ExecuteContext
}

func NewCommand_executeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_executeContext {
	var p = new(Command_executeContext)

	InitEmptyExecuteContext(&p.ExecuteContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExecuteContext))

	return p
}

func (s *Command_executeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_executeContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_executeContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_executeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_execute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Execute() (localctx IExecuteContext) {
	localctx = NewExecuteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FileSysCLIParserRULE_execute)
	var _la int

	localctx = NewCommand_executeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(FileSysCLIParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(100)
			p.Parametro()
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkdiskContext is an interface to support dynamic dispatch.
type IMkdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMkdiskContext differentiates from other interfaces.
	IsMkdiskContext()
}

type MkdiskContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkdiskContext() *MkdiskContext {
	var p = new(MkdiskContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_mkdisk
	return p
}

func InitEmptyMkdiskContext(p *MkdiskContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_mkdisk
}

func (*MkdiskContext) IsMkdiskContext() {}

func NewMkdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkdiskContext {
	var p = new(MkdiskContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_mkdisk

	return p
}

func (s *MkdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *MkdiskContext) CopyAll(ctx *MkdiskContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MkdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_mkdiskContext struct {
	MkdiskContext
}

func NewCommand_mkdiskContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_mkdiskContext {
	var p = new(Command_mkdiskContext)

	InitEmptyMkdiskContext(&p.MkdiskContext)
	p.parser = parser
	p.CopyAll(ctx.(*MkdiskContext))

	return p
}

func (s *Command_mkdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_mkdiskContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_mkdiskContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_mkdiskContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_mkdisk(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Mkdisk() (localctx IMkdiskContext) {
	localctx = NewMkdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FileSysCLIParserRULE_mkdisk)
	var _la int

	localctx = NewCommand_mkdiskContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(FileSysCLIParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(107)
			p.Parametro()
		}

		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRmdiskContext is an interface to support dynamic dispatch.
type IRmdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRmdiskContext differentiates from other interfaces.
	IsRmdiskContext()
}

type RmdiskContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRmdiskContext() *RmdiskContext {
	var p = new(RmdiskContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_rmdisk
	return p
}

func InitEmptyRmdiskContext(p *RmdiskContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_rmdisk
}

func (*RmdiskContext) IsRmdiskContext() {}

func NewRmdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RmdiskContext {
	var p = new(RmdiskContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_rmdisk

	return p
}

func (s *RmdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *RmdiskContext) CopyAll(ctx *RmdiskContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RmdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RmdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_rmdiskContext struct {
	RmdiskContext
}

func NewCommand_rmdiskContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_rmdiskContext {
	var p = new(Command_rmdiskContext)

	InitEmptyRmdiskContext(&p.RmdiskContext)
	p.parser = parser
	p.CopyAll(ctx.(*RmdiskContext))

	return p
}

func (s *Command_rmdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_rmdiskContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_rmdiskContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_rmdiskContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_rmdisk(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Rmdisk() (localctx IRmdiskContext) {
	localctx = NewRmdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FileSysCLIParserRULE_rmdisk)
	var _la int

	localctx = NewCommand_rmdiskContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Match(FileSysCLIParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(114)
			p.Parametro()
		}

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFdiskContext is an interface to support dynamic dispatch.
type IFdiskContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFdiskContext differentiates from other interfaces.
	IsFdiskContext()
}

type FdiskContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFdiskContext() *FdiskContext {
	var p = new(FdiskContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_fdisk
	return p
}

func InitEmptyFdiskContext(p *FdiskContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_fdisk
}

func (*FdiskContext) IsFdiskContext() {}

func NewFdiskContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FdiskContext {
	var p = new(FdiskContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_fdisk

	return p
}

func (s *FdiskContext) GetParser() antlr.Parser { return s.parser }

func (s *FdiskContext) CopyAll(ctx *FdiskContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FdiskContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_fdiskContext struct {
	FdiskContext
}

func NewCommand_fdiskContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_fdiskContext {
	var p = new(Command_fdiskContext)

	InitEmptyFdiskContext(&p.FdiskContext)
	p.parser = parser
	p.CopyAll(ctx.(*FdiskContext))

	return p
}

func (s *Command_fdiskContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_fdiskContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_fdiskContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_fdiskContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_fdisk(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Fdisk() (localctx IFdiskContext) {
	localctx = NewFdiskContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FileSysCLIParserRULE_fdisk)
	var _la int

	localctx = NewCommand_fdiskContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(FileSysCLIParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(121)
			p.Parametro()
		}

		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMountContext is an interface to support dynamic dispatch.
type IMountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMountContext differentiates from other interfaces.
	IsMountContext()
}

type MountContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMountContext() *MountContext {
	var p = new(MountContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_mount
	return p
}

func InitEmptyMountContext(p *MountContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_mount
}

func (*MountContext) IsMountContext() {}

func NewMountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MountContext {
	var p = new(MountContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_mount

	return p
}

func (s *MountContext) GetParser() antlr.Parser { return s.parser }

func (s *MountContext) CopyAll(ctx *MountContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_mountContext struct {
	MountContext
}

func NewCommand_mountContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_mountContext {
	var p = new(Command_mountContext)

	InitEmptyMountContext(&p.MountContext)
	p.parser = parser
	p.CopyAll(ctx.(*MountContext))

	return p
}

func (s *Command_mountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_mountContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_mountContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_mountContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_mount(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Mount() (localctx IMountContext) {
	localctx = NewMountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FileSysCLIParserRULE_mount)
	var _la int

	localctx = NewCommand_mountContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(FileSysCLIParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(128)
			p.Parametro()
		}

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnmountContext is an interface to support dynamic dispatch.
type IUnmountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsUnmountContext differentiates from other interfaces.
	IsUnmountContext()
}

type UnmountContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnmountContext() *UnmountContext {
	var p = new(UnmountContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_unmount
	return p
}

func InitEmptyUnmountContext(p *UnmountContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_unmount
}

func (*UnmountContext) IsUnmountContext() {}

func NewUnmountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnmountContext {
	var p = new(UnmountContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_unmount

	return p
}

func (s *UnmountContext) GetParser() antlr.Parser { return s.parser }

func (s *UnmountContext) CopyAll(ctx *UnmountContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *UnmountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnmountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_unmountContext struct {
	UnmountContext
}

func NewCommand_unmountContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_unmountContext {
	var p = new(Command_unmountContext)

	InitEmptyUnmountContext(&p.UnmountContext)
	p.parser = parser
	p.CopyAll(ctx.(*UnmountContext))

	return p
}

func (s *Command_unmountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_unmountContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_unmountContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_unmountContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_unmount(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Unmount() (localctx IUnmountContext) {
	localctx = NewUnmountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FileSysCLIParserRULE_unmount)
	var _la int

	localctx = NewCommand_unmountContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(FileSysCLIParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(135)
			p.Parametro()
		}

		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkfsContext is an interface to support dynamic dispatch.
type IMkfsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMkfsContext differentiates from other interfaces.
	IsMkfsContext()
}

type MkfsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkfsContext() *MkfsContext {
	var p = new(MkfsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_mkfs
	return p
}

func InitEmptyMkfsContext(p *MkfsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_mkfs
}

func (*MkfsContext) IsMkfsContext() {}

func NewMkfsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkfsContext {
	var p = new(MkfsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_mkfs

	return p
}

func (s *MkfsContext) GetParser() antlr.Parser { return s.parser }

func (s *MkfsContext) CopyAll(ctx *MkfsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MkfsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_mkfsContext struct {
	MkfsContext
}

func NewCommand_mkfsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_mkfsContext {
	var p = new(Command_mkfsContext)

	InitEmptyMkfsContext(&p.MkfsContext)
	p.parser = parser
	p.CopyAll(ctx.(*MkfsContext))

	return p
}

func (s *Command_mkfsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_mkfsContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_mkfsContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_mkfsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_mkfs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Mkfs() (localctx IMkfsContext) {
	localctx = NewMkfsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FileSysCLIParserRULE_mkfs)
	var _la int

	localctx = NewCommand_mkfsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(FileSysCLIParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(142)
			p.Parametro()
		}

		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoginContext is an interface to support dynamic dispatch.
type ILoginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLoginContext differentiates from other interfaces.
	IsLoginContext()
}

type LoginContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoginContext() *LoginContext {
	var p = new(LoginContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_login
	return p
}

func InitEmptyLoginContext(p *LoginContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_login
}

func (*LoginContext) IsLoginContext() {}

func NewLoginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoginContext {
	var p = new(LoginContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_login

	return p
}

func (s *LoginContext) GetParser() antlr.Parser { return s.parser }

func (s *LoginContext) CopyAll(ctx *LoginContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LoginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_loginContext struct {
	LoginContext
}

func NewCommand_loginContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_loginContext {
	var p = new(Command_loginContext)

	InitEmptyLoginContext(&p.LoginContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoginContext))

	return p
}

func (s *Command_loginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_loginContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_loginContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_loginContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_login(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Login() (localctx ILoginContext) {
	localctx = NewLoginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FileSysCLIParserRULE_login)
	var _la int

	localctx = NewCommand_loginContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(FileSysCLIParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(149)
			p.Parametro()
		}

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogoutContext is an interface to support dynamic dispatch.
type ILogoutContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLogoutContext differentiates from other interfaces.
	IsLogoutContext()
}

type LogoutContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogoutContext() *LogoutContext {
	var p = new(LogoutContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_logout
	return p
}

func InitEmptyLogoutContext(p *LogoutContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_logout
}

func (*LogoutContext) IsLogoutContext() {}

func NewLogoutContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogoutContext {
	var p = new(LogoutContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_logout

	return p
}

func (s *LogoutContext) GetParser() antlr.Parser { return s.parser }

func (s *LogoutContext) CopyAll(ctx *LogoutContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LogoutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogoutContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_logoutContext struct {
	LogoutContext
}

func NewCommand_logoutContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_logoutContext {
	var p = new(Command_logoutContext)

	InitEmptyLogoutContext(&p.LogoutContext)
	p.parser = parser
	p.CopyAll(ctx.(*LogoutContext))

	return p
}

func (s *Command_logoutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_logoutContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_logoutContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_logoutContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_logout(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Logout() (localctx ILogoutContext) {
	localctx = NewLogoutContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FileSysCLIParserRULE_logout)
	var _la int

	localctx = NewCommand_logoutContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(FileSysCLIParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(156)
			p.Parametro()
		}

		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkgrpContext is an interface to support dynamic dispatch.
type IMkgrpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMkgrpContext differentiates from other interfaces.
	IsMkgrpContext()
}

type MkgrpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkgrpContext() *MkgrpContext {
	var p = new(MkgrpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_mkgrp
	return p
}

func InitEmptyMkgrpContext(p *MkgrpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_mkgrp
}

func (*MkgrpContext) IsMkgrpContext() {}

func NewMkgrpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkgrpContext {
	var p = new(MkgrpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_mkgrp

	return p
}

func (s *MkgrpContext) GetParser() antlr.Parser { return s.parser }

func (s *MkgrpContext) CopyAll(ctx *MkgrpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MkgrpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkgrpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_mkgrpContext struct {
	MkgrpContext
}

func NewCommand_mkgrpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_mkgrpContext {
	var p = new(Command_mkgrpContext)

	InitEmptyMkgrpContext(&p.MkgrpContext)
	p.parser = parser
	p.CopyAll(ctx.(*MkgrpContext))

	return p
}

func (s *Command_mkgrpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_mkgrpContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_mkgrpContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_mkgrpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_mkgrp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Mkgrp() (localctx IMkgrpContext) {
	localctx = NewMkgrpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FileSysCLIParserRULE_mkgrp)
	var _la int

	localctx = NewCommand_mkgrpContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(FileSysCLIParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(163)
			p.Parametro()
		}

		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRmgrpContext is an interface to support dynamic dispatch.
type IRmgrpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRmgrpContext differentiates from other interfaces.
	IsRmgrpContext()
}

type RmgrpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRmgrpContext() *RmgrpContext {
	var p = new(RmgrpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_rmgrp
	return p
}

func InitEmptyRmgrpContext(p *RmgrpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_rmgrp
}

func (*RmgrpContext) IsRmgrpContext() {}

func NewRmgrpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RmgrpContext {
	var p = new(RmgrpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_rmgrp

	return p
}

func (s *RmgrpContext) GetParser() antlr.Parser { return s.parser }

func (s *RmgrpContext) CopyAll(ctx *RmgrpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RmgrpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RmgrpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_rmgrpContext struct {
	RmgrpContext
}

func NewCommand_rmgrpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_rmgrpContext {
	var p = new(Command_rmgrpContext)

	InitEmptyRmgrpContext(&p.RmgrpContext)
	p.parser = parser
	p.CopyAll(ctx.(*RmgrpContext))

	return p
}

func (s *Command_rmgrpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_rmgrpContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_rmgrpContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_rmgrpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_rmgrp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Rmgrp() (localctx IRmgrpContext) {
	localctx = NewRmgrpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FileSysCLIParserRULE_rmgrp)
	var _la int

	localctx = NewCommand_rmgrpContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(FileSysCLIParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(170)
			p.Parametro()
		}

		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkusrContext is an interface to support dynamic dispatch.
type IMkusrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMkusrContext differentiates from other interfaces.
	IsMkusrContext()
}

type MkusrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkusrContext() *MkusrContext {
	var p = new(MkusrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_mkusr
	return p
}

func InitEmptyMkusrContext(p *MkusrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_mkusr
}

func (*MkusrContext) IsMkusrContext() {}

func NewMkusrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkusrContext {
	var p = new(MkusrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_mkusr

	return p
}

func (s *MkusrContext) GetParser() antlr.Parser { return s.parser }

func (s *MkusrContext) CopyAll(ctx *MkusrContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MkusrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkusrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_mkusrContext struct {
	MkusrContext
}

func NewCommand_mkusrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_mkusrContext {
	var p = new(Command_mkusrContext)

	InitEmptyMkusrContext(&p.MkusrContext)
	p.parser = parser
	p.CopyAll(ctx.(*MkusrContext))

	return p
}

func (s *Command_mkusrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_mkusrContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_mkusrContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_mkusrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_mkusr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Mkusr() (localctx IMkusrContext) {
	localctx = NewMkusrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FileSysCLIParserRULE_mkusr)
	var _la int

	localctx = NewCommand_mkusrContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Match(FileSysCLIParserT__11)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(177)
			p.Parametro()
		}

		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRmusrContext is an interface to support dynamic dispatch.
type IRmusrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRmusrContext differentiates from other interfaces.
	IsRmusrContext()
}

type RmusrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRmusrContext() *RmusrContext {
	var p = new(RmusrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_rmusr
	return p
}

func InitEmptyRmusrContext(p *RmusrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_rmusr
}

func (*RmusrContext) IsRmusrContext() {}

func NewRmusrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RmusrContext {
	var p = new(RmusrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_rmusr

	return p
}

func (s *RmusrContext) GetParser() antlr.Parser { return s.parser }

func (s *RmusrContext) CopyAll(ctx *RmusrContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RmusrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RmusrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_rmusrContext struct {
	RmusrContext
}

func NewCommand_rmusrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_rmusrContext {
	var p = new(Command_rmusrContext)

	InitEmptyRmusrContext(&p.RmusrContext)
	p.parser = parser
	p.CopyAll(ctx.(*RmusrContext))

	return p
}

func (s *Command_rmusrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_rmusrContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_rmusrContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_rmusrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_rmusr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Rmusr() (localctx IRmusrContext) {
	localctx = NewRmusrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FileSysCLIParserRULE_rmusr)
	var _la int

	localctx = NewCommand_rmusrContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(FileSysCLIParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(184)
			p.Parametro()
		}

		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkfileContext is an interface to support dynamic dispatch.
type IMkfileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMkfileContext differentiates from other interfaces.
	IsMkfileContext()
}

type MkfileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkfileContext() *MkfileContext {
	var p = new(MkfileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_mkfile
	return p
}

func InitEmptyMkfileContext(p *MkfileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_mkfile
}

func (*MkfileContext) IsMkfileContext() {}

func NewMkfileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkfileContext {
	var p = new(MkfileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_mkfile

	return p
}

func (s *MkfileContext) GetParser() antlr.Parser { return s.parser }

func (s *MkfileContext) CopyAll(ctx *MkfileContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MkfileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_mkfileContext struct {
	MkfileContext
}

func NewCommand_mkfileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_mkfileContext {
	var p = new(Command_mkfileContext)

	InitEmptyMkfileContext(&p.MkfileContext)
	p.parser = parser
	p.CopyAll(ctx.(*MkfileContext))

	return p
}

func (s *Command_mkfileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_mkfileContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_mkfileContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_mkfileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_mkfile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Mkfile() (localctx IMkfileContext) {
	localctx = NewMkfileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FileSysCLIParserRULE_mkfile)
	var _la int

	localctx = NewCommand_mkfileContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(FileSysCLIParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(191)
			p.Parametro()
		}

		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRepContext is an interface to support dynamic dispatch.
type IRepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRepContext differentiates from other interfaces.
	IsRepContext()
}

type RepContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRepContext() *RepContext {
	var p = new(RepContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_rep
	return p
}

func InitEmptyRepContext(p *RepContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_rep
}

func (*RepContext) IsRepContext() {}

func NewRepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepContext {
	var p = new(RepContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_rep

	return p
}

func (s *RepContext) GetParser() antlr.Parser { return s.parser }

func (s *RepContext) CopyAll(ctx *RepContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_repContext struct {
	RepContext
}

func NewCommand_repContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_repContext {
	var p = new(Command_repContext)

	InitEmptyRepContext(&p.RepContext)
	p.parser = parser
	p.CopyAll(ctx.(*RepContext))

	return p
}

func (s *Command_repContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_repContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_repContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_repContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_rep(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Rep() (localctx IRepContext) {
	localctx = NewRepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FileSysCLIParserRULE_rep)
	var _la int

	localctx = NewCommand_repContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(FileSysCLIParserT__14)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(198)
			p.Parametro()
		}

		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICatContext is an interface to support dynamic dispatch.
type ICatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCatContext differentiates from other interfaces.
	IsCatContext()
}

type CatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCatContext() *CatContext {
	var p = new(CatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_cat
	return p
}

func InitEmptyCatContext(p *CatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_cat
}

func (*CatContext) IsCatContext() {}

func NewCatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatContext {
	var p = new(CatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_cat

	return p
}

func (s *CatContext) GetParser() antlr.Parser { return s.parser }

func (s *CatContext) CopyAll(ctx *CatContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_catContext struct {
	CatContext
}

func NewCommand_catContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_catContext {
	var p = new(Command_catContext)

	InitEmptyCatContext(&p.CatContext)
	p.parser = parser
	p.CopyAll(ctx.(*CatContext))

	return p
}

func (s *Command_catContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_catContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_catContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_catContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_cat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Cat() (localctx ICatContext) {
	localctx = NewCatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FileSysCLIParserRULE_cat)
	var _la int

	localctx = NewCommand_catContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(FileSysCLIParserT__15)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(205)
			p.Parametro()
		}

		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRemoveContext is an interface to support dynamic dispatch.
type IRemoveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRemoveContext differentiates from other interfaces.
	IsRemoveContext()
}

type RemoveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRemoveContext() *RemoveContext {
	var p = new(RemoveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_remove
	return p
}

func InitEmptyRemoveContext(p *RemoveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_remove
}

func (*RemoveContext) IsRemoveContext() {}

func NewRemoveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RemoveContext {
	var p = new(RemoveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_remove

	return p
}

func (s *RemoveContext) GetParser() antlr.Parser { return s.parser }

func (s *RemoveContext) CopyAll(ctx *RemoveContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RemoveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemoveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_rmContext struct {
	RemoveContext
}

func NewCommand_rmContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_rmContext {
	var p = new(Command_rmContext)

	InitEmptyRemoveContext(&p.RemoveContext)
	p.parser = parser
	p.CopyAll(ctx.(*RemoveContext))

	return p
}

func (s *Command_rmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_rmContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_rmContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_rmContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_rm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Remove() (localctx IRemoveContext) {
	localctx = NewRemoveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FileSysCLIParserRULE_remove)
	var _la int

	localctx = NewCommand_rmContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.Match(FileSysCLIParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(212)
			p.Parametro()
		}

		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEditContext is an interface to support dynamic dispatch.
type IEditContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEditContext differentiates from other interfaces.
	IsEditContext()
}

type EditContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEditContext() *EditContext {
	var p = new(EditContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_edit
	return p
}

func InitEmptyEditContext(p *EditContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_edit
}

func (*EditContext) IsEditContext() {}

func NewEditContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EditContext {
	var p = new(EditContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_edit

	return p
}

func (s *EditContext) GetParser() antlr.Parser { return s.parser }

func (s *EditContext) CopyAll(ctx *EditContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *EditContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EditContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_editContext struct {
	EditContext
}

func NewCommand_editContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_editContext {
	var p = new(Command_editContext)

	InitEmptyEditContext(&p.EditContext)
	p.parser = parser
	p.CopyAll(ctx.(*EditContext))

	return p
}

func (s *Command_editContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_editContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_editContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_editContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_edit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Edit() (localctx IEditContext) {
	localctx = NewEditContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FileSysCLIParserRULE_edit)
	var _la int

	localctx = NewCommand_editContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(FileSysCLIParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(219)
			p.Parametro()
		}

		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRenameContext is an interface to support dynamic dispatch.
type IRenameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRenameContext differentiates from other interfaces.
	IsRenameContext()
}

type RenameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRenameContext() *RenameContext {
	var p = new(RenameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_rename
	return p
}

func InitEmptyRenameContext(p *RenameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_rename
}

func (*RenameContext) IsRenameContext() {}

func NewRenameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RenameContext {
	var p = new(RenameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_rename

	return p
}

func (s *RenameContext) GetParser() antlr.Parser { return s.parser }

func (s *RenameContext) CopyAll(ctx *RenameContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RenameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RenameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_renameContext struct {
	RenameContext
}

func NewCommand_renameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_renameContext {
	var p = new(Command_renameContext)

	InitEmptyRenameContext(&p.RenameContext)
	p.parser = parser
	p.CopyAll(ctx.(*RenameContext))

	return p
}

func (s *Command_renameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_renameContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_renameContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_renameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_rename(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Rename() (localctx IRenameContext) {
	localctx = NewRenameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FileSysCLIParserRULE_rename)
	var _la int

	localctx = NewCommand_renameContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Match(FileSysCLIParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(226)
			p.Parametro()
		}

		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMkdirContext is an interface to support dynamic dispatch.
type IMkdirContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMkdirContext differentiates from other interfaces.
	IsMkdirContext()
}

type MkdirContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkdirContext() *MkdirContext {
	var p = new(MkdirContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_mkdir
	return p
}

func InitEmptyMkdirContext(p *MkdirContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_mkdir
}

func (*MkdirContext) IsMkdirContext() {}

func NewMkdirContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkdirContext {
	var p = new(MkdirContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_mkdir

	return p
}

func (s *MkdirContext) GetParser() antlr.Parser { return s.parser }

func (s *MkdirContext) CopyAll(ctx *MkdirContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MkdirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdirContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_mkdirContext struct {
	MkdirContext
}

func NewCommand_mkdirContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_mkdirContext {
	var p = new(Command_mkdirContext)

	InitEmptyMkdirContext(&p.MkdirContext)
	p.parser = parser
	p.CopyAll(ctx.(*MkdirContext))

	return p
}

func (s *Command_mkdirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_mkdirContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_mkdirContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_mkdirContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_mkdir(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Mkdir() (localctx IMkdirContext) {
	localctx = NewMkdirContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FileSysCLIParserRULE_mkdir)
	var _la int

	localctx = NewCommand_mkdirContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(FileSysCLIParserT__19)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(233)
			p.Parametro()
		}

		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPauseContext is an interface to support dynamic dispatch.
type IPauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPauseContext differentiates from other interfaces.
	IsPauseContext()
}

type PauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPauseContext() *PauseContext {
	var p = new(PauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_pause
	return p
}

func InitEmptyPauseContext(p *PauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_pause
}

func (*PauseContext) IsPauseContext() {}

func NewPauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PauseContext {
	var p = new(PauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_pause

	return p
}

func (s *PauseContext) GetParser() antlr.Parser { return s.parser }

func (s *PauseContext) CopyAll(ctx *PauseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_pauseContext struct {
	PauseContext
}

func NewCommand_pauseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_pauseContext {
	var p = new(Command_pauseContext)

	InitEmptyPauseContext(&p.PauseContext)
	p.parser = parser
	p.CopyAll(ctx.(*PauseContext))

	return p
}

func (s *Command_pauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_pauseContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_pauseContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_pauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_pause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Pause() (localctx IPauseContext) {
	localctx = NewPauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, FileSysCLIParserRULE_pause)
	var _la int

	localctx = NewCommand_pauseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(FileSysCLIParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(240)
			p.Parametro()
		}

		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IChgrpContext is an interface to support dynamic dispatch.
type IChgrpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsChgrpContext differentiates from other interfaces.
	IsChgrpContext()
}

type ChgrpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChgrpContext() *ChgrpContext {
	var p = new(ChgrpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_chgrp
	return p
}

func InitEmptyChgrpContext(p *ChgrpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_chgrp
}

func (*ChgrpContext) IsChgrpContext() {}

func NewChgrpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChgrpContext {
	var p = new(ChgrpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_chgrp

	return p
}

func (s *ChgrpContext) GetParser() antlr.Parser { return s.parser }

func (s *ChgrpContext) CopyAll(ctx *ChgrpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ChgrpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChgrpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_chgrpContext struct {
	ChgrpContext
}

func NewCommand_chgrpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_chgrpContext {
	var p = new(Command_chgrpContext)

	InitEmptyChgrpContext(&p.ChgrpContext)
	p.parser = parser
	p.CopyAll(ctx.(*ChgrpContext))

	return p
}

func (s *Command_chgrpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_chgrpContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_chgrpContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_chgrpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_chgrp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Chgrp() (localctx IChgrpContext) {
	localctx = NewChgrpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, FileSysCLIParserRULE_chgrp)
	var _la int

	localctx = NewCommand_chgrpContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(FileSysCLIParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(247)
			p.Parametro()
		}

		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMoveContext is an interface to support dynamic dispatch.
type IMoveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMoveContext differentiates from other interfaces.
	IsMoveContext()
}

type MoveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMoveContext() *MoveContext {
	var p = new(MoveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_move
	return p
}

func InitEmptyMoveContext(p *MoveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_move
}

func (*MoveContext) IsMoveContext() {}

func NewMoveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MoveContext {
	var p = new(MoveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_move

	return p
}

func (s *MoveContext) GetParser() antlr.Parser { return s.parser }

func (s *MoveContext) CopyAll(ctx *MoveContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *MoveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MoveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_moveContext struct {
	MoveContext
}

func NewCommand_moveContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_moveContext {
	var p = new(Command_moveContext)

	InitEmptyMoveContext(&p.MoveContext)
	p.parser = parser
	p.CopyAll(ctx.(*MoveContext))

	return p
}

func (s *Command_moveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_moveContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_moveContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_moveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_move(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Move() (localctx IMoveContext) {
	localctx = NewMoveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, FileSysCLIParserRULE_move)
	var _la int

	localctx = NewCommand_moveContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Match(FileSysCLIParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(254)
			p.Parametro()
		}

		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILossContext is an interface to support dynamic dispatch.
type ILossContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLossContext differentiates from other interfaces.
	IsLossContext()
}

type LossContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLossContext() *LossContext {
	var p = new(LossContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_loss
	return p
}

func InitEmptyLossContext(p *LossContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_loss
}

func (*LossContext) IsLossContext() {}

func NewLossContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LossContext {
	var p = new(LossContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_loss

	return p
}

func (s *LossContext) GetParser() antlr.Parser { return s.parser }

func (s *LossContext) CopyAll(ctx *LossContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LossContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LossContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_lossContext struct {
	LossContext
}

func NewCommand_lossContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_lossContext {
	var p = new(Command_lossContext)

	InitEmptyLossContext(&p.LossContext)
	p.parser = parser
	p.CopyAll(ctx.(*LossContext))

	return p
}

func (s *Command_lossContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_lossContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_lossContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_lossContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_loss(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Loss() (localctx ILossContext) {
	localctx = NewLossContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, FileSysCLIParserRULE_loss)
	var _la int

	localctx = NewCommand_lossContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		p.Match(FileSysCLIParserT__23)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(261)
			p.Parametro()
		}

		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRecoveryContext is an interface to support dynamic dispatch.
type IRecoveryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRecoveryContext differentiates from other interfaces.
	IsRecoveryContext()
}

type RecoveryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecoveryContext() *RecoveryContext {
	var p = new(RecoveryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_recovery
	return p
}

func InitEmptyRecoveryContext(p *RecoveryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_recovery
}

func (*RecoveryContext) IsRecoveryContext() {}

func NewRecoveryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecoveryContext {
	var p = new(RecoveryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_recovery

	return p
}

func (s *RecoveryContext) GetParser() antlr.Parser { return s.parser }

func (s *RecoveryContext) CopyAll(ctx *RecoveryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RecoveryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecoveryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_recoveryContext struct {
	RecoveryContext
}

func NewCommand_recoveryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_recoveryContext {
	var p = new(Command_recoveryContext)

	InitEmptyRecoveryContext(&p.RecoveryContext)
	p.parser = parser
	p.CopyAll(ctx.(*RecoveryContext))

	return p
}

func (s *Command_recoveryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_recoveryContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_recoveryContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_recoveryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_recovery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Recovery() (localctx IRecoveryContext) {
	localctx = NewRecoveryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, FileSysCLIParserRULE_recovery)
	var _la int

	localctx = NewCommand_recoveryContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Match(FileSysCLIParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(268)
			p.Parametro()
		}

		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IChownContext is an interface to support dynamic dispatch.
type IChownContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsChownContext differentiates from other interfaces.
	IsChownContext()
}

type ChownContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChownContext() *ChownContext {
	var p = new(ChownContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_chown
	return p
}

func InitEmptyChownContext(p *ChownContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_chown
}

func (*ChownContext) IsChownContext() {}

func NewChownContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChownContext {
	var p = new(ChownContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_chown

	return p
}

func (s *ChownContext) GetParser() antlr.Parser { return s.parser }

func (s *ChownContext) CopyAll(ctx *ChownContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ChownContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChownContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_chownContext struct {
	ChownContext
}

func NewCommand_chownContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_chownContext {
	var p = new(Command_chownContext)

	InitEmptyChownContext(&p.ChownContext)
	p.parser = parser
	p.CopyAll(ctx.(*ChownContext))

	return p
}

func (s *Command_chownContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_chownContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_chownContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_chownContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_chown(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Chown() (localctx IChownContext) {
	localctx = NewChownContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, FileSysCLIParserRULE_chown)
	var _la int

	localctx = NewCommand_chownContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.Match(FileSysCLIParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(275)
			p.Parametro()
		}

		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IChmodContext is an interface to support dynamic dispatch.
type IChmodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsChmodContext differentiates from other interfaces.
	IsChmodContext()
}

type ChmodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChmodContext() *ChmodContext {
	var p = new(ChmodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_chmod
	return p
}

func InitEmptyChmodContext(p *ChmodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_chmod
}

func (*ChmodContext) IsChmodContext() {}

func NewChmodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChmodContext {
	var p = new(ChmodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_chmod

	return p
}

func (s *ChmodContext) GetParser() antlr.Parser { return s.parser }

func (s *ChmodContext) CopyAll(ctx *ChmodContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ChmodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChmodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Command_chmodContext struct {
	ChmodContext
}

func NewCommand_chmodContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Command_chmodContext {
	var p = new(Command_chmodContext)

	InitEmptyChmodContext(&p.ChmodContext)
	p.parser = parser
	p.CopyAll(ctx.(*ChmodContext))

	return p
}

func (s *Command_chmodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_chmodContext) AllParametro() []IParametroContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParametroContext); ok {
			len++
		}
	}

	tst := make([]IParametroContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParametroContext); ok {
			tst[i] = t.(IParametroContext)
			i++
		}
	}

	return tst
}

func (s *Command_chmodContext) Parametro(i int) IParametroContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametroContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametroContext)
}

func (s *Command_chmodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitCommand_chmod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Chmod() (localctx IChmodContext) {
	localctx = NewChmodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, FileSysCLIParserRULE_chmod)
	var _la int

	localctx = NewCommand_chmodContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(FileSysCLIParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == FileSysCLIParserQUOTED_TEXT_ {
		{
			p.SetState(282)
			p.Parametro()
		}

		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametroContext is an interface to support dynamic dispatch.
type IParametroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParametroContext differentiates from other interfaces.
	IsParametroContext()
}

type ParametroContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametroContext() *ParametroContext {
	var p = new(ParametroContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_parametro
	return p
}

func InitEmptyParametroContext(p *ParametroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_parametro
}

func (*ParametroContext) IsParametroContext() {}

func NewParametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametroContext {
	var p = new(ParametroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_parametro

	return p
}

func (s *ParametroContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametroContext) CopyAll(ctx *ParametroContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ParametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParameterContext struct {
	ParametroContext
}

func NewParameterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParameterContext {
	var p = new(ParameterContext)

	InitEmptyParametroContext(&p.ParametroContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParametroContext))

	return p
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) QUOTED_TEXT_() antlr.TerminalNode {
	return s.GetToken(FileSysCLIParserQUOTED_TEXT_, 0)
}

func (s *ParameterContext) Value_parametro() IValue_parametroContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValue_parametroContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValue_parametroContext)
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Parametro() (localctx IParametroContext) {
	localctx = NewParametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, FileSysCLIParserRULE_parametro)
	var _la int

	localctx = NewParameterContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		p.Match(FileSysCLIParserQUOTED_TEXT_)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == FileSysCLIParserT__27 {
		{
			p.SetState(289)
			p.Match(FileSysCLIParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)
			p.Value_parametro()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValue_parametroContext is an interface to support dynamic dispatch.
type IValue_parametroContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValue_parametroContext differentiates from other interfaces.
	IsValue_parametroContext()
}

type Value_parametroContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValue_parametroContext() *Value_parametroContext {
	var p = new(Value_parametroContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_value_parametro
	return p
}

func InitEmptyValue_parametroContext(p *Value_parametroContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = FileSysCLIParserRULE_value_parametro
}

func (*Value_parametroContext) IsValue_parametroContext() {}

func NewValue_parametroContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Value_parametroContext {
	var p = new(Value_parametroContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = FileSysCLIParserRULE_value_parametro

	return p
}

func (s *Value_parametroContext) GetParser() antlr.Parser { return s.parser }

func (s *Value_parametroContext) CopyAll(ctx *Value_parametroContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Value_parametroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_parametroContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Value_parameterContext struct {
	Value_parametroContext
}

func NewValue_parameterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Value_parameterContext {
	var p = new(Value_parameterContext)

	InitEmptyValue_parametroContext(&p.Value_parametroContext)
	p.parser = parser
	p.CopyAll(ctx.(*Value_parametroContext))

	return p
}

func (s *Value_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_parameterContext) QUOTED_TEXT_() antlr.TerminalNode {
	return s.GetToken(FileSysCLIParserQUOTED_TEXT_, 0)
}

func (s *Value_parameterContext) STRING() antlr.TerminalNode {
	return s.GetToken(FileSysCLIParserSTRING, 0)
}

func (s *Value_parameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FileSysCLIVisitor:
		return t.VisitValue_parameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FileSysCLIParser) Value_parametro() (localctx IValue_parametroContext) {
	localctx = NewValue_parametroContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, FileSysCLIParserRULE_value_parametro)
	var _la int

	localctx = NewValue_parameterContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FileSysCLIParserSTRING || _la == FileSysCLIParserQUOTED_TEXT_) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
