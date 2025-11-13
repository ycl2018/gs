// Code generated from D:/go/src/gs/Gs.g4 by ANTLR 4.13.2. DO NOT EDIT.

package gen // Gs
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

type GsParser struct {
	*antlr.BaseParser
}

var GsParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func gsParserInit() {
	staticData := &GsParserStaticData
	staticData.LiteralNames = []string{
		"", "'printf'", "'='", "'+='", "'-='", "'*='", "'/='", "'%='", "'$'",
		"'true'", "'false'", "'nil'", "'&&'", "'||'", "'!'", "'if'", "'else'",
		"'print'", "'for'", "'range'", "'return'", "'func'", "'type'", "'struct'",
		"'new'", "'break'", "'continue'", "'?.'", "'?['", "'+'", "'-'", "'*'",
		"'/'", "'%'", "'=='", "'<'", "'>'", "'>='", "'<='", "'!='", "'&'", "'|'",
		"'^'", "'++'", "'--'", "'.'", "'['", "'('", "')'", "'{'", "'}'", "']'",
		"':'", "';'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "ENV", "TRUE", "FALSE", "NIL", "AND",
		"OR", "NOT", "IF", "ELSE", "PRINT", "FOR", "RANGE", "RETURN", "FUNC",
		"TYPE", "STRUCT", "NEW", "BREAK", "CONTINUE", "SAFE_DOT", "SAFE_LBRACK",
		"ADD", "SUB", "MUL", "DIV", "MOD", "EQ", "LT", "GT", "GEQ", "LEQ", "NEQ",
		"BITAND", "BITOR", "XOR", "INCR", "DECR", "DOT", "LBRACK", "LPAREN",
		"RPAREN", "LBRACE", "RBRACE", "RBRACK", "COLON", "SEMICOLON", "COMMA",
		"INT", "FLOAT", "STRING", "WS", "NEWLINE", "SL_COMMENT", "ML_COMMENT",
		"ID",
	}
	staticData.RuleNames = []string{
		"program", "structDefinition", "functionDefinition", "block", "statement",
		"assign", "incrDecr", "iterVar", "forInit", "forUpdate", "selfAssign",
		"updateItem", "selfAssignOp", "call", "expr", "logicalOrExpr", "logicalAndExpr",
		"comparisonExpr", "addExpr", "binExpr", "mulExpr", "atom", "lvalue",
		"arrayLiteral", "indexAccess", "sliceExpr", "dictLiteral", "dictEntry",
		"instance", "qid", "primary", "compOp", "addOp", "bitOp", "mulOp",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 62, 445, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 1, 0, 1, 0, 4, 0, 74,
		8, 0, 11, 0, 12, 0, 75, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 5, 1, 87, 8, 1, 10, 1, 12, 1, 90, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 5, 2, 100, 8, 2, 10, 2, 12, 2, 103, 9, 2, 3, 2, 105,
		8, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 4, 3, 113, 8, 3, 11, 3, 12, 3,
		114, 1, 3, 3, 3, 118, 8, 3, 5, 3, 120, 8, 3, 10, 3, 12, 3, 123, 9, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 136,
		8, 4, 10, 4, 12, 4, 139, 9, 4, 3, 4, 141, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		5, 4, 147, 8, 4, 10, 4, 12, 4, 150, 9, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4,
		156, 8, 4, 10, 4, 12, 4, 159, 9, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 165,
		8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 171, 8, 4, 1, 4, 1, 4, 3, 4, 175, 8,
		4, 1, 4, 1, 4, 3, 4, 179, 8, 4, 1, 4, 1, 4, 3, 4, 183, 8, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 3, 4, 200, 8, 4, 1, 5, 1, 5, 1, 5, 5, 5, 205, 8, 5, 10, 5, 12, 5,
		208, 9, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 214, 8, 5, 10, 5, 12, 5, 217,
		9, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 226, 8, 7, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 3, 11,
		239, 8, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 248,
		8, 13, 10, 13, 12, 13, 251, 9, 13, 1, 13, 3, 13, 254, 8, 13, 3, 13, 256,
		8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 5, 15, 265, 8,
		15, 10, 15, 12, 15, 268, 9, 15, 1, 16, 1, 16, 1, 16, 5, 16, 273, 8, 16,
		10, 16, 12, 16, 276, 9, 16, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 282, 8,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 288, 8, 18, 10, 18, 12, 18, 291,
		9, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 297, 8, 19, 10, 19, 12, 19, 300,
		9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 306, 8, 20, 10, 20, 12, 20, 309,
		9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 335, 8, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 3, 22, 343, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23,
		349, 8, 23, 10, 23, 12, 23, 352, 9, 23, 1, 23, 3, 23, 355, 8, 23, 3, 23,
		357, 8, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 365, 8, 24,
		1, 24, 1, 24, 1, 25, 3, 25, 370, 8, 25, 1, 25, 1, 25, 3, 25, 374, 8, 25,
		1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 380, 8, 26, 10, 26, 12, 26, 383, 9,
		26, 1, 26, 3, 26, 386, 8, 26, 3, 26, 388, 8, 26, 1, 26, 1, 26, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 399, 8, 27, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 411, 8,
		28, 10, 28, 12, 28, 414, 9, 28, 1, 28, 3, 28, 417, 8, 28, 3, 28, 419, 8,
		28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29,
		430, 8, 29, 10, 29, 12, 29, 433, 9, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1,
		32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 0, 0, 35, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 0, 11, 2, 0, 53, 53, 59,
		59, 1, 0, 43, 44, 1, 0, 3, 7, 2, 0, 9, 10, 55, 57, 2, 0, 27, 27, 45, 45,
		2, 0, 28, 28, 46, 46, 2, 0, 8, 8, 62, 62, 1, 0, 34, 39, 1, 0, 29, 30, 1,
		0, 40, 42, 1, 0, 31, 33, 490, 0, 73, 1, 0, 0, 0, 2, 79, 1, 0, 0, 0, 4,
		93, 1, 0, 0, 0, 6, 109, 1, 0, 0, 0, 8, 199, 1, 0, 0, 0, 10, 201, 1, 0,
		0, 0, 12, 218, 1, 0, 0, 0, 14, 225, 1, 0, 0, 0, 16, 227, 1, 0, 0, 0, 18,
		229, 1, 0, 0, 0, 20, 231, 1, 0, 0, 0, 22, 238, 1, 0, 0, 0, 24, 240, 1,
		0, 0, 0, 26, 242, 1, 0, 0, 0, 28, 259, 1, 0, 0, 0, 30, 261, 1, 0, 0, 0,
		32, 269, 1, 0, 0, 0, 34, 277, 1, 0, 0, 0, 36, 283, 1, 0, 0, 0, 38, 292,
		1, 0, 0, 0, 40, 301, 1, 0, 0, 0, 42, 334, 1, 0, 0, 0, 44, 342, 1, 0, 0,
		0, 46, 344, 1, 0, 0, 0, 48, 360, 1, 0, 0, 0, 50, 369, 1, 0, 0, 0, 52, 375,
		1, 0, 0, 0, 54, 398, 1, 0, 0, 0, 56, 400, 1, 0, 0, 0, 58, 422, 1, 0, 0,
		0, 60, 434, 1, 0, 0, 0, 62, 436, 1, 0, 0, 0, 64, 438, 1, 0, 0, 0, 66, 440,
		1, 0, 0, 0, 68, 442, 1, 0, 0, 0, 70, 74, 3, 4, 2, 0, 71, 74, 3, 2, 1, 0,
		72, 74, 3, 8, 4, 0, 73, 70, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 72, 1,
		0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76,
		77, 1, 0, 0, 0, 77, 78, 5, 0, 0, 1, 78, 1, 1, 0, 0, 0, 79, 80, 5, 22, 0,
		0, 80, 81, 5, 62, 0, 0, 81, 82, 5, 23, 0, 0, 82, 83, 5, 49, 0, 0, 83, 88,
		5, 62, 0, 0, 84, 85, 5, 54, 0, 0, 85, 87, 5, 62, 0, 0, 86, 84, 1, 0, 0,
		0, 87, 90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 91,
		1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 91, 92, 5, 50, 0, 0, 92, 3, 1, 0, 0, 0,
		93, 94, 5, 21, 0, 0, 94, 95, 5, 62, 0, 0, 95, 104, 5, 47, 0, 0, 96, 101,
		5, 62, 0, 0, 97, 98, 5, 54, 0, 0, 98, 100, 5, 62, 0, 0, 99, 97, 1, 0, 0,
		0, 100, 103, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102,
		105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 104, 96, 1, 0, 0, 0, 104, 105, 1,
		0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 107, 5, 48, 0, 0, 107, 108, 3, 6, 3,
		0, 108, 5, 1, 0, 0, 0, 109, 121, 5, 49, 0, 0, 110, 117, 3, 8, 4, 0, 111,
		113, 7, 0, 0, 0, 112, 111, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 112,
		1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 118, 5, 0,
		0, 1, 117, 112, 1, 0, 0, 0, 117, 116, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0,
		118, 120, 1, 0, 0, 0, 119, 110, 1, 0, 0, 0, 120, 123, 1, 0, 0, 0, 121,
		119, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 124, 1, 0, 0, 0, 123, 121,
		1, 0, 0, 0, 124, 125, 5, 50, 0, 0, 125, 7, 1, 0, 0, 0, 126, 200, 5, 53,
		0, 0, 127, 200, 3, 2, 1, 0, 128, 200, 3, 10, 5, 0, 129, 200, 3, 20, 10,
		0, 130, 200, 3, 12, 6, 0, 131, 140, 5, 20, 0, 0, 132, 137, 3, 28, 14, 0,
		133, 134, 5, 54, 0, 0, 134, 136, 3, 28, 14, 0, 135, 133, 1, 0, 0, 0, 136,
		139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 141,
		1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 140, 132, 1, 0, 0, 0, 140, 141, 1, 0,
		0, 0, 141, 200, 1, 0, 0, 0, 142, 143, 5, 17, 0, 0, 143, 148, 3, 28, 14,
		0, 144, 145, 5, 54, 0, 0, 145, 147, 3, 28, 14, 0, 146, 144, 1, 0, 0, 0,
		147, 150, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149,
		200, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 151, 152, 5, 1, 0, 0, 152, 157,
		3, 28, 14, 0, 153, 154, 5, 54, 0, 0, 154, 156, 3, 28, 14, 0, 155, 153,
		1, 0, 0, 0, 156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0,
		0, 0, 158, 200, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 164, 5, 15, 0, 0,
		161, 162, 3, 10, 5, 0, 162, 163, 5, 53, 0, 0, 163, 165, 1, 0, 0, 0, 164,
		161, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 167,
		3, 28, 14, 0, 167, 170, 3, 6, 3, 0, 168, 169, 5, 16, 0, 0, 169, 171, 3,
		6, 3, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 200, 1, 0, 0,
		0, 172, 174, 5, 18, 0, 0, 173, 175, 3, 16, 8, 0, 174, 173, 1, 0, 0, 0,
		174, 175, 1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 178, 5, 53, 0, 0, 177,
		179, 3, 28, 14, 0, 178, 177, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180,
		1, 0, 0, 0, 180, 182, 5, 53, 0, 0, 181, 183, 3, 18, 9, 0, 182, 181, 1,
		0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0, 184, 200, 3, 6, 3,
		0, 185, 186, 5, 18, 0, 0, 186, 187, 3, 14, 7, 0, 187, 188, 5, 2, 0, 0,
		188, 189, 5, 19, 0, 0, 189, 190, 3, 28, 14, 0, 190, 191, 3, 6, 3, 0, 191,
		200, 1, 0, 0, 0, 192, 193, 5, 18, 0, 0, 193, 194, 3, 28, 14, 0, 194, 195,
		3, 6, 3, 0, 195, 200, 1, 0, 0, 0, 196, 200, 3, 26, 13, 0, 197, 200, 5,
		25, 0, 0, 198, 200, 5, 26, 0, 0, 199, 126, 1, 0, 0, 0, 199, 127, 1, 0,
		0, 0, 199, 128, 1, 0, 0, 0, 199, 129, 1, 0, 0, 0, 199, 130, 1, 0, 0, 0,
		199, 131, 1, 0, 0, 0, 199, 142, 1, 0, 0, 0, 199, 151, 1, 0, 0, 0, 199,
		160, 1, 0, 0, 0, 199, 172, 1, 0, 0, 0, 199, 185, 1, 0, 0, 0, 199, 192,
		1, 0, 0, 0, 199, 196, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199, 198, 1, 0,
		0, 0, 200, 9, 1, 0, 0, 0, 201, 206, 3, 58, 29, 0, 202, 203, 5, 54, 0, 0,
		203, 205, 3, 58, 29, 0, 204, 202, 1, 0, 0, 0, 205, 208, 1, 0, 0, 0, 206,
		204, 1, 0, 0, 0, 206, 207, 1, 0, 0, 0, 207, 209, 1, 0, 0, 0, 208, 206,
		1, 0, 0, 0, 209, 210, 5, 2, 0, 0, 210, 215, 3, 28, 14, 0, 211, 212, 5,
		54, 0, 0, 212, 214, 3, 28, 14, 0, 213, 211, 1, 0, 0, 0, 214, 217, 1, 0,
		0, 0, 215, 213, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 11, 1, 0, 0, 0,
		217, 215, 1, 0, 0, 0, 218, 219, 3, 58, 29, 0, 219, 220, 7, 1, 0, 0, 220,
		13, 1, 0, 0, 0, 221, 226, 5, 62, 0, 0, 222, 223, 5, 62, 0, 0, 223, 224,
		5, 54, 0, 0, 224, 226, 5, 62, 0, 0, 225, 221, 1, 0, 0, 0, 225, 222, 1,
		0, 0, 0, 226, 15, 1, 0, 0, 0, 227, 228, 3, 10, 5, 0, 228, 17, 1, 0, 0,
		0, 229, 230, 3, 22, 11, 0, 230, 19, 1, 0, 0, 0, 231, 232, 3, 58, 29, 0,
		232, 233, 3, 24, 12, 0, 233, 234, 3, 28, 14, 0, 234, 21, 1, 0, 0, 0, 235,
		239, 3, 20, 10, 0, 236, 239, 3, 12, 6, 0, 237, 239, 3, 10, 5, 0, 238, 235,
		1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 238, 237, 1, 0, 0, 0, 239, 23, 1, 0,
		0, 0, 240, 241, 7, 2, 0, 0, 241, 25, 1, 0, 0, 0, 242, 243, 5, 62, 0, 0,
		243, 255, 5, 47, 0, 0, 244, 249, 3, 28, 14, 0, 245, 246, 5, 54, 0, 0, 246,
		248, 3, 28, 14, 0, 247, 245, 1, 0, 0, 0, 248, 251, 1, 0, 0, 0, 249, 247,
		1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 253, 1, 0, 0, 0, 251, 249, 1, 0,
		0, 0, 252, 254, 5, 54, 0, 0, 253, 252, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0,
		254, 256, 1, 0, 0, 0, 255, 244, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256,
		257, 1, 0, 0, 0, 257, 258, 5, 48, 0, 0, 258, 27, 1, 0, 0, 0, 259, 260,
		3, 30, 15, 0, 260, 29, 1, 0, 0, 0, 261, 266, 3, 32, 16, 0, 262, 263, 5,
		13, 0, 0, 263, 265, 3, 32, 16, 0, 264, 262, 1, 0, 0, 0, 265, 268, 1, 0,
		0, 0, 266, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 31, 1, 0, 0, 0,
		268, 266, 1, 0, 0, 0, 269, 274, 3, 34, 17, 0, 270, 271, 5, 12, 0, 0, 271,
		273, 3, 34, 17, 0, 272, 270, 1, 0, 0, 0, 273, 276, 1, 0, 0, 0, 274, 272,
		1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275, 33, 1, 0, 0, 0, 276, 274, 1, 0,
		0, 0, 277, 281, 3, 36, 18, 0, 278, 279, 3, 62, 31, 0, 279, 280, 3, 36,
		18, 0, 280, 282, 1, 0, 0, 0, 281, 278, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0,
		282, 35, 1, 0, 0, 0, 283, 289, 3, 38, 19, 0, 284, 285, 3, 64, 32, 0, 285,
		286, 3, 38, 19, 0, 286, 288, 1, 0, 0, 0, 287, 284, 1, 0, 0, 0, 288, 291,
		1, 0, 0, 0, 289, 287, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 37, 1, 0,
		0, 0, 291, 289, 1, 0, 0, 0, 292, 298, 3, 40, 20, 0, 293, 294, 3, 66, 33,
		0, 294, 295, 3, 40, 20, 0, 295, 297, 1, 0, 0, 0, 296, 293, 1, 0, 0, 0,
		297, 300, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299,
		39, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0, 301, 307, 3, 42, 21, 0, 302, 303,
		3, 68, 34, 0, 303, 304, 3, 42, 21, 0, 304, 306, 1, 0, 0, 0, 305, 302, 1,
		0, 0, 0, 306, 309, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0, 307, 308, 1, 0, 0,
		0, 308, 41, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 310, 311, 5, 30, 0, 0, 311,
		335, 3, 42, 21, 0, 312, 335, 5, 55, 0, 0, 313, 335, 5, 56, 0, 0, 314, 335,
		5, 57, 0, 0, 315, 335, 5, 9, 0, 0, 316, 335, 5, 10, 0, 0, 317, 335, 5,
		11, 0, 0, 318, 319, 5, 14, 0, 0, 319, 335, 3, 28, 14, 0, 320, 335, 3, 58,
		29, 0, 321, 335, 3, 26, 13, 0, 322, 335, 3, 56, 28, 0, 323, 335, 3, 46,
		23, 0, 324, 335, 3, 48, 24, 0, 325, 335, 3, 52, 26, 0, 326, 327, 5, 47,
		0, 0, 327, 328, 3, 28, 14, 0, 328, 329, 5, 48, 0, 0, 329, 335, 1, 0, 0,
		0, 330, 331, 5, 31, 0, 0, 331, 335, 3, 44, 22, 0, 332, 333, 5, 40, 0, 0,
		333, 335, 3, 44, 22, 0, 334, 310, 1, 0, 0, 0, 334, 312, 1, 0, 0, 0, 334,
		313, 1, 0, 0, 0, 334, 314, 1, 0, 0, 0, 334, 315, 1, 0, 0, 0, 334, 316,
		1, 0, 0, 0, 334, 317, 1, 0, 0, 0, 334, 318, 1, 0, 0, 0, 334, 320, 1, 0,
		0, 0, 334, 321, 1, 0, 0, 0, 334, 322, 1, 0, 0, 0, 334, 323, 1, 0, 0, 0,
		334, 324, 1, 0, 0, 0, 334, 325, 1, 0, 0, 0, 334, 326, 1, 0, 0, 0, 334,
		330, 1, 0, 0, 0, 334, 332, 1, 0, 0, 0, 335, 43, 1, 0, 0, 0, 336, 343, 3,
		58, 29, 0, 337, 343, 3, 48, 24, 0, 338, 339, 5, 31, 0, 0, 339, 343, 3,
		44, 22, 0, 340, 341, 5, 40, 0, 0, 341, 343, 3, 44, 22, 0, 342, 336, 1,
		0, 0, 0, 342, 337, 1, 0, 0, 0, 342, 338, 1, 0, 0, 0, 342, 340, 1, 0, 0,
		0, 343, 45, 1, 0, 0, 0, 344, 356, 5, 46, 0, 0, 345, 350, 3, 28, 14, 0,
		346, 347, 5, 54, 0, 0, 347, 349, 3, 28, 14, 0, 348, 346, 1, 0, 0, 0, 349,
		352, 1, 0, 0, 0, 350, 348, 1, 0, 0, 0, 350, 351, 1, 0, 0, 0, 351, 354,
		1, 0, 0, 0, 352, 350, 1, 0, 0, 0, 353, 355, 5, 54, 0, 0, 354, 353, 1, 0,
		0, 0, 354, 355, 1, 0, 0, 0, 355, 357, 1, 0, 0, 0, 356, 345, 1, 0, 0, 0,
		356, 357, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358, 359, 5, 51, 0, 0, 359,
		47, 1, 0, 0, 0, 360, 361, 3, 58, 29, 0, 361, 364, 5, 46, 0, 0, 362, 365,
		3, 28, 14, 0, 363, 365, 3, 50, 25, 0, 364, 362, 1, 0, 0, 0, 364, 363, 1,
		0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 367, 5, 51, 0, 0, 367, 49, 1, 0, 0,
		0, 368, 370, 3, 28, 14, 0, 369, 368, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0,
		370, 371, 1, 0, 0, 0, 371, 373, 5, 52, 0, 0, 372, 374, 3, 28, 14, 0, 373,
		372, 1, 0, 0, 0, 373, 374, 1, 0, 0, 0, 374, 51, 1, 0, 0, 0, 375, 387, 5,
		49, 0, 0, 376, 381, 3, 54, 27, 0, 377, 378, 5, 54, 0, 0, 378, 380, 3, 54,
		27, 0, 379, 377, 1, 0, 0, 0, 380, 383, 1, 0, 0, 0, 381, 379, 1, 0, 0, 0,
		381, 382, 1, 0, 0, 0, 382, 385, 1, 0, 0, 0, 383, 381, 1, 0, 0, 0, 384,
		386, 5, 54, 0, 0, 385, 384, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 388,
		1, 0, 0, 0, 387, 376, 1, 0, 0, 0, 387, 388, 1, 0, 0, 0, 388, 389, 1, 0,
		0, 0, 389, 390, 5, 50, 0, 0, 390, 53, 1, 0, 0, 0, 391, 392, 7, 3, 0, 0,
		392, 393, 5, 52, 0, 0, 393, 399, 3, 28, 14, 0, 394, 395, 3, 58, 29, 0,
		395, 396, 5, 52, 0, 0, 396, 397, 3, 28, 14, 0, 397, 399, 1, 0, 0, 0, 398,
		391, 1, 0, 0, 0, 398, 394, 1, 0, 0, 0, 399, 55, 1, 0, 0, 0, 400, 401, 5,
		24, 0, 0, 401, 402, 5, 62, 0, 0, 402, 418, 5, 49, 0, 0, 403, 404, 5, 62,
		0, 0, 404, 405, 5, 52, 0, 0, 405, 412, 3, 28, 14, 0, 406, 407, 5, 54, 0,
		0, 407, 408, 5, 62, 0, 0, 408, 409, 5, 52, 0, 0, 409, 411, 3, 28, 14, 0,
		410, 406, 1, 0, 0, 0, 411, 414, 1, 0, 0, 0, 412, 410, 1, 0, 0, 0, 412,
		413, 1, 0, 0, 0, 413, 416, 1, 0, 0, 0, 414, 412, 1, 0, 0, 0, 415, 417,
		5, 54, 0, 0, 416, 415, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 419, 1, 0,
		0, 0, 418, 403, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 420, 1, 0, 0, 0,
		420, 421, 5, 50, 0, 0, 421, 57, 1, 0, 0, 0, 422, 431, 3, 60, 30, 0, 423,
		424, 7, 4, 0, 0, 424, 430, 5, 62, 0, 0, 425, 426, 7, 5, 0, 0, 426, 427,
		3, 28, 14, 0, 427, 428, 5, 51, 0, 0, 428, 430, 1, 0, 0, 0, 429, 423, 1,
		0, 0, 0, 429, 425, 1, 0, 0, 0, 430, 433, 1, 0, 0, 0, 431, 429, 1, 0, 0,
		0, 431, 432, 1, 0, 0, 0, 432, 59, 1, 0, 0, 0, 433, 431, 1, 0, 0, 0, 434,
		435, 7, 6, 0, 0, 435, 61, 1, 0, 0, 0, 436, 437, 7, 7, 0, 0, 437, 63, 1,
		0, 0, 0, 438, 439, 7, 8, 0, 0, 439, 65, 1, 0, 0, 0, 440, 441, 7, 9, 0,
		0, 441, 67, 1, 0, 0, 0, 442, 443, 7, 10, 0, 0, 443, 69, 1, 0, 0, 0, 48,
		73, 75, 88, 101, 104, 114, 117, 121, 137, 140, 148, 157, 164, 170, 174,
		178, 182, 199, 206, 215, 225, 238, 249, 253, 255, 266, 274, 281, 289, 298,
		307, 334, 342, 350, 354, 356, 364, 369, 373, 381, 385, 387, 398, 412, 416,
		418, 429, 431,
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

// GsParserInit initializes any static state used to implement GsParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGsParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GsParserInit() {
	staticData := &GsParserStaticData
	staticData.once.Do(gsParserInit)
}

// NewGsParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGsParser(input antlr.TokenStream) *GsParser {
	GsParserInit()
	this := new(GsParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GsParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Gs.g4"

	return this
}

// GsParser tokens.
const (
	GsParserEOF         = antlr.TokenEOF
	GsParserT__0        = 1
	GsParserT__1        = 2
	GsParserT__2        = 3
	GsParserT__3        = 4
	GsParserT__4        = 5
	GsParserT__5        = 6
	GsParserT__6        = 7
	GsParserENV         = 8
	GsParserTRUE        = 9
	GsParserFALSE       = 10
	GsParserNIL         = 11
	GsParserAND         = 12
	GsParserOR          = 13
	GsParserNOT         = 14
	GsParserIF          = 15
	GsParserELSE        = 16
	GsParserPRINT       = 17
	GsParserFOR         = 18
	GsParserRANGE       = 19
	GsParserRETURN      = 20
	GsParserFUNC        = 21
	GsParserTYPE        = 22
	GsParserSTRUCT      = 23
	GsParserNEW         = 24
	GsParserBREAK       = 25
	GsParserCONTINUE    = 26
	GsParserSAFE_DOT    = 27
	GsParserSAFE_LBRACK = 28
	GsParserADD         = 29
	GsParserSUB         = 30
	GsParserMUL         = 31
	GsParserDIV         = 32
	GsParserMOD         = 33
	GsParserEQ          = 34
	GsParserLT          = 35
	GsParserGT          = 36
	GsParserGEQ         = 37
	GsParserLEQ         = 38
	GsParserNEQ         = 39
	GsParserBITAND      = 40
	GsParserBITOR       = 41
	GsParserXOR         = 42
	GsParserINCR        = 43
	GsParserDECR        = 44
	GsParserDOT         = 45
	GsParserLBRACK      = 46
	GsParserLPAREN      = 47
	GsParserRPAREN      = 48
	GsParserLBRACE      = 49
	GsParserRBRACE      = 50
	GsParserRBRACK      = 51
	GsParserCOLON       = 52
	GsParserSEMICOLON   = 53
	GsParserCOMMA       = 54
	GsParserINT         = 55
	GsParserFLOAT       = 56
	GsParserSTRING      = 57
	GsParserWS          = 58
	GsParserNEWLINE     = 59
	GsParserSL_COMMENT  = 60
	GsParserML_COMMENT  = 61
	GsParserID          = 62
)

// GsParser rules.
const (
	GsParserRULE_program            = 0
	GsParserRULE_structDefinition   = 1
	GsParserRULE_functionDefinition = 2
	GsParserRULE_block              = 3
	GsParserRULE_statement          = 4
	GsParserRULE_assign             = 5
	GsParserRULE_incrDecr           = 6
	GsParserRULE_iterVar            = 7
	GsParserRULE_forInit            = 8
	GsParserRULE_forUpdate          = 9
	GsParserRULE_selfAssign         = 10
	GsParserRULE_updateItem         = 11
	GsParserRULE_selfAssignOp       = 12
	GsParserRULE_call               = 13
	GsParserRULE_expr               = 14
	GsParserRULE_logicalOrExpr      = 15
	GsParserRULE_logicalAndExpr     = 16
	GsParserRULE_comparisonExpr     = 17
	GsParserRULE_addExpr            = 18
	GsParserRULE_binExpr            = 19
	GsParserRULE_mulExpr            = 20
	GsParserRULE_atom               = 21
	GsParserRULE_lvalue             = 22
	GsParserRULE_arrayLiteral       = 23
	GsParserRULE_indexAccess        = 24
	GsParserRULE_sliceExpr          = 25
	GsParserRULE_dictLiteral        = 26
	GsParserRULE_dictEntry          = 27
	GsParserRULE_instance           = 28
	GsParserRULE_qid                = 29
	GsParserRULE_primary            = 30
	GsParserRULE_compOp             = 31
	GsParserRULE_addOp              = 32
	GsParserRULE_bitOp              = 33
	GsParserRULE_mulOp              = 34
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllFunctionDefinition() []IFunctionDefinitionContext
	FunctionDefinition(i int) IFunctionDefinitionContext
	AllStructDefinition() []IStructDefinitionContext
	StructDefinition(i int) IStructDefinitionContext
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(GsParserEOF, 0)
}

func (s *ProgramContext) AllFunctionDefinition() []IFunctionDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IFunctionDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionDefinitionContext); ok {
			tst[i] = t.(IFunctionDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) FunctionDefinition(i int) IFunctionDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefinitionContext); ok {
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

	return t.(IFunctionDefinitionContext)
}

func (s *ProgramContext) AllStructDefinition() []IStructDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStructDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IStructDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStructDefinitionContext); ok {
			tst[i] = t.(IStructDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) StructDefinition(i int) IStructDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDefinitionContext); ok {
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

	return t.(IStructDefinitionContext)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GsParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4620693217790558466) != 0) {
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(70)
				p.FunctionDefinition()
			}

		case 2:
			{
				p.SetState(71)
				p.StructDefinition()
			}

		case 3:
			{
				p.SetState(72)
				p.Statement()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(77)
		p.Match(GsParserEOF)
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

// IStructDefinitionContext is an interface to support dynamic dispatch.
type IStructDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TYPE() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	STRUCT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsStructDefinitionContext differentiates from other interfaces.
	IsStructDefinitionContext()
}

type StructDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDefinitionContext() *StructDefinitionContext {
	var p = new(StructDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_structDefinition
	return p
}

func InitEmptyStructDefinitionContext(p *StructDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_structDefinition
}

func (*StructDefinitionContext) IsStructDefinitionContext() {}

func NewStructDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDefinitionContext {
	var p = new(StructDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_structDefinition

	return p
}

func (s *StructDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDefinitionContext) TYPE() antlr.TerminalNode {
	return s.GetToken(GsParserTYPE, 0)
}

func (s *StructDefinitionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GsParserID)
}

func (s *StructDefinitionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GsParserID, i)
}

func (s *StructDefinitionContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(GsParserSTRUCT, 0)
}

func (s *StructDefinitionContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GsParserLBRACE, 0)
}

func (s *StructDefinitionContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GsParserRBRACE, 0)
}

func (s *StructDefinitionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *StructDefinitionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}

func (s *StructDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitStructDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) StructDefinition() (localctx IStructDefinitionContext) {
	localctx = NewStructDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GsParserRULE_structDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(GsParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Match(GsParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(GsParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Match(GsParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.Match(GsParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GsParserCOMMA {
		{
			p.SetState(84)
			p.Match(GsParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(85)
			p.Match(GsParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(91)
		p.Match(GsParserRBRACE)
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

// IFunctionDefinitionContext is an interface to support dynamic dispatch.
type IFunctionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Block() IBlockContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunctionDefinitionContext differentiates from other interfaces.
	IsFunctionDefinitionContext()
}

type FunctionDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefinitionContext() *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_functionDefinition
	return p
}

func InitEmptyFunctionDefinitionContext(p *FunctionDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_functionDefinition
}

func (*FunctionDefinitionContext) IsFunctionDefinitionContext() {}

func NewFunctionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_functionDefinition

	return p
}

func (s *FunctionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefinitionContext) FUNC() antlr.TerminalNode {
	return s.GetToken(GsParserFUNC, 0)
}

func (s *FunctionDefinitionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GsParserID)
}

func (s *FunctionDefinitionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GsParserID, i)
}

func (s *FunctionDefinitionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *FunctionDefinitionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *FunctionDefinitionContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDefinitionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *FunctionDefinitionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}

func (s *FunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitFunctionDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) FunctionDefinition() (localctx IFunctionDefinitionContext) {
	localctx = NewFunctionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GsParserRULE_functionDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(GsParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Match(GsParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Match(GsParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GsParserID {
		{
			p.SetState(96)
			p.Match(GsParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GsParserCOMMA {
			{
				p.SetState(97)
				p.Match(GsParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(98)
				p.Match(GsParserID)
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
		}

	}
	{
		p.SetState(106)
		p.Match(GsParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Block()
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllEOF() []antlr.TerminalNode
	EOF(i int) antlr.TerminalNode
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode
	AllNEWLINE() []antlr.TerminalNode
	NEWLINE(i int) antlr.TerminalNode

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GsParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GsParserRBRACE, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
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

	return t.(IStatementContext)
}

func (s *BlockContext) AllEOF() []antlr.TerminalNode {
	return s.GetTokens(GsParserEOF)
}

func (s *BlockContext) EOF(i int) antlr.TerminalNode {
	return s.GetToken(GsParserEOF, i)
}

func (s *BlockContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(GsParserSEMICOLON)
}

func (s *BlockContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(GsParserSEMICOLON, i)
}

func (s *BlockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(GsParserNEWLINE)
}

func (s *BlockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(GsParserNEWLINE, i)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GsParserRULE_block)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(GsParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4620693217788461314) != 0 {
		{
			p.SetState(110)
			p.Statement()
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(111)
						_la = p.GetTokenStream().LA(1)

						if !(_la == GsParserSEMICOLON || _la == GsParserNEWLINE) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}

				p.SetState(114)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(116)
				p.Match(GsParserEOF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(124)
		p.Match(GsParserRBRACE)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForCondStmtContext struct {
	StatementContext
}

func NewForCondStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForCondStmtContext {
	var p = new(ForCondStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ForCondStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForCondStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(GsParserFOR, 0)
}

func (s *ForCondStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForCondStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForCondStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitForCondStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructStmtContext struct {
	StatementContext
}

func NewStructStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructStmtContext {
	var p = new(StructStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *StructStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructStmtContext) StructDefinition() IStructDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDefinitionContext)
}

func (s *StructStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitStructStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignStmtContext struct {
	StatementContext
}

func NewAssignStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignStmtContext {
	var p = new(AssignStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *AssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStmtContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *AssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitAssignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type IncrDecrStmtContext struct {
	StatementContext
}

func NewIncrDecrStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncrDecrStmtContext {
	var p = new(IncrDecrStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *IncrDecrStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrDecrStmtContext) IncrDecr() IIncrDecrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncrDecrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncrDecrContext)
}

func (s *IncrDecrStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitIncrDecrStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForCStyleStmtContext struct {
	StatementContext
}

func NewForCStyleStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForCStyleStmtContext {
	var p = new(ForCStyleStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ForCStyleStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForCStyleStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(GsParserFOR, 0)
}

func (s *ForCStyleStmtContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(GsParserSEMICOLON)
}

func (s *ForCStyleStmtContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(GsParserSEMICOLON, i)
}

func (s *ForCStyleStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForCStyleStmtContext) ForInit() IForInitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForInitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForInitContext)
}

func (s *ForCStyleStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForCStyleStmtContext) ForUpdate() IForUpdateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForUpdateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForUpdateContext)
}

func (s *ForCStyleStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitForCStyleStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type EmptyStmtContext struct {
	StatementContext
}

func NewEmptyStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EmptyStmtContext {
	var p = new(EmptyStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *EmptyStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GsParserSEMICOLON, 0)
}

func (s *EmptyStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitEmptyStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelfOpAssignStmtContext struct {
	StatementContext
}

func NewSelfOpAssignStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelfOpAssignStmtContext {
	var p = new(SelfOpAssignStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *SelfOpAssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelfOpAssignStmtContext) SelfAssign() ISelfAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelfAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelfAssignContext)
}

func (s *SelfOpAssignStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitSelfOpAssignStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStmtContext struct {
	StatementContext
}

func NewReturnStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(GsParserRETURN, 0)
}

func (s *ReturnStmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ReturnStmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ReturnStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *ReturnStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintStmtContext struct {
	StatementContext
}

func NewPrintStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintStmtContext {
	var p = new(PrintStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PrintStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStmtContext) PRINT() antlr.TerminalNode {
	return s.GetToken(GsParserPRINT, 0)
}

func (s *PrintStmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *PrintStmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *PrintStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *PrintStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}

func (s *PrintStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitPrintStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallStmtContext struct {
	StatementContext
}

func NewCallStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallStmtContext {
	var p = new(CallStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *CallStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallStmtContext) Call() ICallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *CallStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitCallStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfStmtContext struct {
	StatementContext
}

func NewIfStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStmtContext {
	var p = new(IfStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(GsParserIF, 0)
}

func (s *IfStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfStmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *IfStmtContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *IfStmtContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GsParserSEMICOLON, 0)
}

func (s *IfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GsParserELSE, 0)
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForRangeStmtContext struct {
	StatementContext
}

func NewForRangeStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForRangeStmtContext {
	var p = new(ForRangeStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ForRangeStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForRangeStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(GsParserFOR, 0)
}

func (s *ForRangeStmtContext) IterVar() IIterVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIterVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIterVarContext)
}

func (s *ForRangeStmtContext) RANGE() antlr.TerminalNode {
	return s.GetToken(GsParserRANGE, 0)
}

func (s *ForRangeStmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForRangeStmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForRangeStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitForRangeStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakStmtContext struct {
	StatementContext
}

func NewBreakStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStmtContext {
	var p = new(BreakStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(GsParserBREAK, 0)
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintfStmtContext struct {
	StatementContext
}

func NewPrintfStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintfStmtContext {
	var p = new(PrintfStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PrintfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintfStmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *PrintfStmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *PrintfStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *PrintfStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}

func (s *PrintfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitPrintfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContinueStmtContext struct {
	StatementContext
}

func NewContinueStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(GsParserCONTINUE, 0)
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GsParserRULE_statement)
	var _la int

	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEmptyStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)
			p.Match(GsParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewStructStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.StructDefinition()
		}

	case 3:
		localctx = NewAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(128)
			p.Assign()
		}

	case 4:
		localctx = NewSelfOpAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(129)
			p.SelfAssign()
		}

	case 5:
		localctx = NewIncrDecrStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(130)
			p.IncrDecr()
		}

	case 6:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(131)
			p.Match(GsParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(132)
				p.Expr()
			}
			p.SetState(137)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == GsParserCOMMA {
				{
					p.SetState(133)
					p.Match(GsParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(134)
					p.Expr()
				}

				p.SetState(139)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 7:
		localctx = NewPrintStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(142)
			p.Match(GsParserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Expr()
		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GsParserCOMMA {
			{
				p.SetState(144)
				p.Match(GsParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(145)
				p.Expr()
			}

			p.SetState(150)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 8:
		localctx = NewPrintfStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(151)
			p.Match(GsParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(152)
			p.Expr()
		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GsParserCOMMA {
			{
				p.SetState(153)
				p.Match(GsParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(154)
				p.Expr()
			}

			p.SetState(159)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case 9:
		localctx = NewIfStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(160)
			p.Match(GsParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(161)
				p.Assign()
			}
			{
				p.SetState(162)
				p.Match(GsParserSEMICOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(166)
			p.Expr()
		}
		{
			p.SetState(167)
			p.Block()
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GsParserELSE {
			{
				p.SetState(168)
				p.Match(GsParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(169)
				p.Block()
			}

		}

	case 10:
		localctx = NewForCStyleStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(172)
			p.Match(GsParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GsParserENV || _la == GsParserID {
			{
				p.SetState(173)
				p.ForInit()
			}

		}
		{
			p.SetState(176)
			p.Match(GsParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4864662756495740672) != 0 {
			{
				p.SetState(177)
				p.Expr()
			}

		}
		{
			p.SetState(180)
			p.Match(GsParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GsParserENV || _la == GsParserID {
			{
				p.SetState(181)
				p.ForUpdate()
			}

		}
		{
			p.SetState(184)
			p.Block()
		}

	case 11:
		localctx = NewForRangeStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(185)
			p.Match(GsParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.IterVar()
		}
		{
			p.SetState(187)
			p.Match(GsParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Match(GsParserRANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Expr()
		}
		{
			p.SetState(190)
			p.Block()
		}

	case 12:
		localctx = NewForCondStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(192)
			p.Match(GsParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Expr()
		}
		{
			p.SetState(194)
			p.Block()
		}

	case 13:
		localctx = NewCallStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(196)
			p.Call()
		}

	case 14:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(197)
			p.Match(GsParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		localctx = NewContinueStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(198)
			p.Match(GsParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IAssignContext is an interface to support dynamic dispatch.
type IAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllQid() []IQidContext
	Qid(i int) IQidContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsAssignContext differentiates from other interfaces.
	IsAssignContext()
}

type AssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignContext() *AssignContext {
	var p = new(AssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_assign
	return p
}

func InitEmptyAssignContext(p *AssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_assign
}

func (*AssignContext) IsAssignContext() {}

func NewAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignContext {
	var p = new(AssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_assign

	return p
}

func (s *AssignContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignContext) AllQid() []IQidContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQidContext); ok {
			len++
		}
	}

	tst := make([]IQidContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQidContext); ok {
			tst[i] = t.(IQidContext)
			i++
		}
	}

	return tst
}

func (s *AssignContext) Qid(i int) IQidContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQidContext); ok {
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

	return t.(IQidContext)
}

func (s *AssignContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AssignContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *AssignContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *AssignContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) Assign() (localctx IAssignContext) {
	localctx = NewAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GsParserRULE_assign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Qid()
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GsParserCOMMA {
		{
			p.SetState(202)
			p.Match(GsParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.Qid()
		}

		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(209)
		p.Match(GsParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.Expr()
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GsParserCOMMA {
		{
			p.SetState(211)
			p.Match(GsParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)
			p.Expr()
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

// IIncrDecrContext is an interface to support dynamic dispatch.
type IIncrDecrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Qid() IQidContext
	INCR() antlr.TerminalNode
	DECR() antlr.TerminalNode

	// IsIncrDecrContext differentiates from other interfaces.
	IsIncrDecrContext()
}

type IncrDecrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncrDecrContext() *IncrDecrContext {
	var p = new(IncrDecrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_incrDecr
	return p
}

func InitEmptyIncrDecrContext(p *IncrDecrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_incrDecr
}

func (*IncrDecrContext) IsIncrDecrContext() {}

func NewIncrDecrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncrDecrContext {
	var p = new(IncrDecrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_incrDecr

	return p
}

func (s *IncrDecrContext) GetParser() antlr.Parser { return s.parser }

func (s *IncrDecrContext) Qid() IQidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQidContext)
}

func (s *IncrDecrContext) INCR() antlr.TerminalNode {
	return s.GetToken(GsParserINCR, 0)
}

func (s *IncrDecrContext) DECR() antlr.TerminalNode {
	return s.GetToken(GsParserDECR, 0)
}

func (s *IncrDecrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrDecrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncrDecrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitIncrDecr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) IncrDecr() (localctx IIncrDecrContext) {
	localctx = NewIncrDecrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GsParserRULE_incrDecr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Qid()
	}
	{
		p.SetState(219)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GsParserINCR || _la == GsParserDECR) {
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

// IIterVarContext is an interface to support dynamic dispatch.
type IIterVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIterVarContext differentiates from other interfaces.
	IsIterVarContext()
}

type IterVarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterVarContext() *IterVarContext {
	var p = new(IterVarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_iterVar
	return p
}

func InitEmptyIterVarContext(p *IterVarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_iterVar
}

func (*IterVarContext) IsIterVarContext() {}

func NewIterVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterVarContext {
	var p = new(IterVarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_iterVar

	return p
}

func (s *IterVarContext) GetParser() antlr.Parser { return s.parser }

func (s *IterVarContext) CopyAll(ctx *IterVarContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IterVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleIterContext struct {
	IterVarContext
}

func NewSingleIterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleIterContext {
	var p = new(SingleIterContext)

	InitEmptyIterVarContext(&p.IterVarContext)
	p.parser = parser
	p.CopyAll(ctx.(*IterVarContext))

	return p
}

func (s *SingleIterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleIterContext) ID() antlr.TerminalNode {
	return s.GetToken(GsParserID, 0)
}

func (s *SingleIterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitSingleIter(s)

	default:
		return t.VisitChildren(s)
	}
}

type DoubleIterContext struct {
	IterVarContext
}

func NewDoubleIterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoubleIterContext {
	var p = new(DoubleIterContext)

	InitEmptyIterVarContext(&p.IterVarContext)
	p.parser = parser
	p.CopyAll(ctx.(*IterVarContext))

	return p
}

func (s *DoubleIterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleIterContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GsParserID)
}

func (s *DoubleIterContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GsParserID, i)
}

func (s *DoubleIterContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, 0)
}

func (s *DoubleIterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitDoubleIter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) IterVar() (localctx IIterVarContext) {
	localctx = NewIterVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GsParserRULE_iterVar)
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingleIterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(221)
			p.Match(GsParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDoubleIterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(222)
			p.Match(GsParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Match(GsParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.Match(GsParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IForInitContext is an interface to support dynamic dispatch.
type IForInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assign() IAssignContext

	// IsForInitContext differentiates from other interfaces.
	IsForInitContext()
}

type ForInitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForInitContext() *ForInitContext {
	var p = new(ForInitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_forInit
	return p
}

func InitEmptyForInitContext(p *ForInitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_forInit
}

func (*ForInitContext) IsForInitContext() {}

func NewForInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForInitContext {
	var p = new(ForInitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_forInit

	return p
}

func (s *ForInitContext) GetParser() antlr.Parser { return s.parser }

func (s *ForInitContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *ForInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitForInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) ForInit() (localctx IForInitContext) {
	localctx = NewForInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GsParserRULE_forInit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Assign()
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

// IForUpdateContext is an interface to support dynamic dispatch.
type IForUpdateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UpdateItem() IUpdateItemContext

	// IsForUpdateContext differentiates from other interfaces.
	IsForUpdateContext()
}

type ForUpdateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForUpdateContext() *ForUpdateContext {
	var p = new(ForUpdateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_forUpdate
	return p
}

func InitEmptyForUpdateContext(p *ForUpdateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_forUpdate
}

func (*ForUpdateContext) IsForUpdateContext() {}

func NewForUpdateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForUpdateContext {
	var p = new(ForUpdateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_forUpdate

	return p
}

func (s *ForUpdateContext) GetParser() antlr.Parser { return s.parser }

func (s *ForUpdateContext) UpdateItem() IUpdateItemContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdateItemContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdateItemContext)
}

func (s *ForUpdateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForUpdateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForUpdateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitForUpdate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) ForUpdate() (localctx IForUpdateContext) {
	localctx = NewForUpdateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GsParserRULE_forUpdate)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.UpdateItem()
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

// ISelfAssignContext is an interface to support dynamic dispatch.
type ISelfAssignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Qid() IQidContext
	SelfAssignOp() ISelfAssignOpContext
	Expr() IExprContext

	// IsSelfAssignContext differentiates from other interfaces.
	IsSelfAssignContext()
}

type SelfAssignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelfAssignContext() *SelfAssignContext {
	var p = new(SelfAssignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_selfAssign
	return p
}

func InitEmptySelfAssignContext(p *SelfAssignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_selfAssign
}

func (*SelfAssignContext) IsSelfAssignContext() {}

func NewSelfAssignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelfAssignContext {
	var p = new(SelfAssignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_selfAssign

	return p
}

func (s *SelfAssignContext) GetParser() antlr.Parser { return s.parser }

func (s *SelfAssignContext) Qid() IQidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQidContext)
}

func (s *SelfAssignContext) SelfAssignOp() ISelfAssignOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelfAssignOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelfAssignOpContext)
}

func (s *SelfAssignContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SelfAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelfAssignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelfAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitSelfAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) SelfAssign() (localctx ISelfAssignContext) {
	localctx = NewSelfAssignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GsParserRULE_selfAssign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Qid()
	}
	{
		p.SetState(232)
		p.SelfAssignOp()
	}
	{
		p.SetState(233)
		p.Expr()
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

// IUpdateItemContext is an interface to support dynamic dispatch.
type IUpdateItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsUpdateItemContext differentiates from other interfaces.
	IsUpdateItemContext()
}

type UpdateItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateItemContext() *UpdateItemContext {
	var p = new(UpdateItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_updateItem
	return p
}

func InitEmptyUpdateItemContext(p *UpdateItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_updateItem
}

func (*UpdateItemContext) IsUpdateItemContext() {}

func NewUpdateItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateItemContext {
	var p = new(UpdateItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_updateItem

	return p
}

func (s *UpdateItemContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateItemContext) CopyAll(ctx *UpdateItemContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *UpdateItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SelfUpdateContext struct {
	UpdateItemContext
}

func NewSelfUpdateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelfUpdateContext {
	var p = new(SelfUpdateContext)

	InitEmptyUpdateItemContext(&p.UpdateItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*UpdateItemContext))

	return p
}

func (s *SelfUpdateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelfUpdateContext) SelfAssign() ISelfAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelfAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelfAssignContext)
}

func (s *SelfUpdateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitSelfUpdate(s)

	default:
		return t.VisitChildren(s)
	}
}

type IncrDecrUpdateContext struct {
	UpdateItemContext
}

func NewIncrDecrUpdateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncrDecrUpdateContext {
	var p = new(IncrDecrUpdateContext)

	InitEmptyUpdateItemContext(&p.UpdateItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*UpdateItemContext))

	return p
}

func (s *IncrDecrUpdateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrDecrUpdateContext) IncrDecr() IIncrDecrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncrDecrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncrDecrContext)
}

func (s *IncrDecrUpdateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitIncrDecrUpdate(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignUpdateContext struct {
	UpdateItemContext
}

func NewAssignUpdateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignUpdateContext {
	var p = new(AssignUpdateContext)

	InitEmptyUpdateItemContext(&p.UpdateItemContext)
	p.parser = parser
	p.CopyAll(ctx.(*UpdateItemContext))

	return p
}

func (s *AssignUpdateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignUpdateContext) Assign() IAssignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignContext)
}

func (s *AssignUpdateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitAssignUpdate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) UpdateItem() (localctx IUpdateItemContext) {
	localctx = NewUpdateItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GsParserRULE_updateItem)
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSelfUpdateContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(235)
			p.SelfAssign()
		}

	case 2:
		localctx = NewIncrDecrUpdateContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(236)
			p.IncrDecr()
		}

	case 3:
		localctx = NewAssignUpdateContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(237)
			p.Assign()
		}

	case antlr.ATNInvalidAltNumber:
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

// ISelfAssignOpContext is an interface to support dynamic dispatch.
type ISelfAssignOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSelfAssignOpContext differentiates from other interfaces.
	IsSelfAssignOpContext()
}

type SelfAssignOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelfAssignOpContext() *SelfAssignOpContext {
	var p = new(SelfAssignOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_selfAssignOp
	return p
}

func InitEmptySelfAssignOpContext(p *SelfAssignOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_selfAssignOp
}

func (*SelfAssignOpContext) IsSelfAssignOpContext() {}

func NewSelfAssignOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelfAssignOpContext {
	var p = new(SelfAssignOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_selfAssignOp

	return p
}

func (s *SelfAssignOpContext) GetParser() antlr.Parser { return s.parser }
func (s *SelfAssignOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelfAssignOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelfAssignOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitSelfAssignOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) SelfAssignOp() (localctx ISelfAssignOpContext) {
	localctx = NewSelfAssignOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GsParserRULE_selfAssignOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&248) != 0) {
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

// ICallContext is an interface to support dynamic dispatch.
type ICallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsCallContext differentiates from other interfaces.
	IsCallContext()
}

type CallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallContext() *CallContext {
	var p = new(CallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_call
	return p
}

func InitEmptyCallContext(p *CallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_call
}

func (*CallContext) IsCallContext() {}

func NewCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallContext {
	var p = new(CallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_call

	return p
}

func (s *CallContext) GetParser() antlr.Parser { return s.parser }

func (s *CallContext) ID() antlr.TerminalNode {
	return s.GetToken(GsParserID, 0)
}

func (s *CallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *CallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *CallContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CallContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *CallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *CallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GsParserRULE_call)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(GsParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Match(GsParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4864662756495740672) != 0 {
		{
			p.SetState(244)
			p.Expr()
		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(245)
					p.Match(GsParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(246)
					p.Expr()
				}

			}
			p.SetState(251)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GsParserCOMMA {
			{
				p.SetState(252)
				p.Match(GsParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(257)
		p.Match(GsParserRPAREN)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LogicalOrExpr() ILogicalOrExprContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) LogicalOrExpr() ILogicalOrExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalOrExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogicalOrExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GsParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.LogicalOrExpr()
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

// ILogicalOrExprContext is an interface to support dynamic dispatch.
type ILogicalOrExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLogicalAndExpr() []ILogicalAndExprContext
	LogicalAndExpr(i int) ILogicalAndExprContext
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsLogicalOrExprContext differentiates from other interfaces.
	IsLogicalOrExprContext()
}

type LogicalOrExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOrExprContext() *LogicalOrExprContext {
	var p = new(LogicalOrExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_logicalOrExpr
	return p
}

func InitEmptyLogicalOrExprContext(p *LogicalOrExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_logicalOrExpr
}

func (*LogicalOrExprContext) IsLogicalOrExprContext() {}

func NewLogicalOrExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOrExprContext {
	var p = new(LogicalOrExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_logicalOrExpr

	return p
}

func (s *LogicalOrExprContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOrExprContext) AllLogicalAndExpr() []ILogicalAndExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogicalAndExprContext); ok {
			len++
		}
	}

	tst := make([]ILogicalAndExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogicalAndExprContext); ok {
			tst[i] = t.(ILogicalAndExprContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOrExprContext) LogicalAndExpr(i int) ILogicalAndExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogicalAndExprContext); ok {
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

	return t.(ILogicalAndExprContext)
}

func (s *LogicalOrExprContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(GsParserOR)
}

func (s *LogicalOrExprContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(GsParserOR, i)
}

func (s *LogicalOrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitLogicalOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) LogicalOrExpr() (localctx ILogicalOrExprContext) {
	localctx = NewLogicalOrExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GsParserRULE_logicalOrExpr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.LogicalAndExpr()
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(262)
				p.Match(GsParserOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(263)
				p.LogicalAndExpr()
			}

		}
		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
		if p.HasError() {
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

// ILogicalAndExprContext is an interface to support dynamic dispatch.
type ILogicalAndExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllComparisonExpr() []IComparisonExprContext
	ComparisonExpr(i int) IComparisonExprContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode

	// IsLogicalAndExprContext differentiates from other interfaces.
	IsLogicalAndExprContext()
}

type LogicalAndExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalAndExprContext() *LogicalAndExprContext {
	var p = new(LogicalAndExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_logicalAndExpr
	return p
}

func InitEmptyLogicalAndExprContext(p *LogicalAndExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_logicalAndExpr
}

func (*LogicalAndExprContext) IsLogicalAndExprContext() {}

func NewLogicalAndExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalAndExprContext {
	var p = new(LogicalAndExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_logicalAndExpr

	return p
}

func (s *LogicalAndExprContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalAndExprContext) AllComparisonExpr() []IComparisonExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComparisonExprContext); ok {
			len++
		}
	}

	tst := make([]IComparisonExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComparisonExprContext); ok {
			tst[i] = t.(IComparisonExprContext)
			i++
		}
	}

	return tst
}

func (s *LogicalAndExprContext) ComparisonExpr(i int) IComparisonExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonExprContext); ok {
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

	return t.(IComparisonExprContext)
}

func (s *LogicalAndExprContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(GsParserAND)
}

func (s *LogicalAndExprContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(GsParserAND, i)
}

func (s *LogicalAndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalAndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitLogicalAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) LogicalAndExpr() (localctx ILogicalAndExprContext) {
	localctx = NewLogicalAndExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GsParserRULE_logicalAndExpr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.ComparisonExpr()
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(270)
				p.Match(GsParserAND)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(271)
				p.ComparisonExpr()
			}

		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
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

// IComparisonExprContext is an interface to support dynamic dispatch.
type IComparisonExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAddExpr() []IAddExprContext
	AddExpr(i int) IAddExprContext
	CompOp() ICompOpContext

	// IsComparisonExprContext differentiates from other interfaces.
	IsComparisonExprContext()
}

type ComparisonExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonExprContext() *ComparisonExprContext {
	var p = new(ComparisonExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_comparisonExpr
	return p
}

func InitEmptyComparisonExprContext(p *ComparisonExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_comparisonExpr
}

func (*ComparisonExprContext) IsComparisonExprContext() {}

func NewComparisonExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonExprContext {
	var p = new(ComparisonExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_comparisonExpr

	return p
}

func (s *ComparisonExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonExprContext) AllAddExpr() []IAddExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAddExprContext); ok {
			len++
		}
	}

	tst := make([]IAddExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAddExprContext); ok {
			tst[i] = t.(IAddExprContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonExprContext) AddExpr(i int) IAddExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddExprContext); ok {
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

	return t.(IAddExprContext)
}

func (s *ComparisonExprContext) CompOp() ICompOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompOpContext)
}

func (s *ComparisonExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitComparisonExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) ComparisonExpr() (localctx IComparisonExprContext) {
	localctx = NewComparisonExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GsParserRULE_comparisonExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.AddExpr()
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(278)
			p.CompOp()
		}
		{
			p.SetState(279)
			p.AddExpr()
		}

	} else if p.HasError() { // JIM
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

// IAddExprContext is an interface to support dynamic dispatch.
type IAddExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllBinExpr() []IBinExprContext
	BinExpr(i int) IBinExprContext
	AllAddOp() []IAddOpContext
	AddOp(i int) IAddOpContext

	// IsAddExprContext differentiates from other interfaces.
	IsAddExprContext()
}

type AddExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddExprContext() *AddExprContext {
	var p = new(AddExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_addExpr
	return p
}

func InitEmptyAddExprContext(p *AddExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_addExpr
}

func (*AddExprContext) IsAddExprContext() {}

func NewAddExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddExprContext {
	var p = new(AddExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_addExpr

	return p
}

func (s *AddExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AddExprContext) AllBinExpr() []IBinExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBinExprContext); ok {
			len++
		}
	}

	tst := make([]IBinExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBinExprContext); ok {
			tst[i] = t.(IBinExprContext)
			i++
		}
	}

	return tst
}

func (s *AddExprContext) BinExpr(i int) IBinExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinExprContext); ok {
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

	return t.(IBinExprContext)
}

func (s *AddExprContext) AllAddOp() []IAddOpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAddOpContext); ok {
			len++
		}
	}

	tst := make([]IAddOpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAddOpContext); ok {
			tst[i] = t.(IAddOpContext)
			i++
		}
	}

	return tst
}

func (s *AddExprContext) AddOp(i int) IAddOpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddOpContext); ok {
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

	return t.(IAddOpContext)
}

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitAddExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) AddExpr() (localctx IAddExprContext) {
	localctx = NewAddExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GsParserRULE_addExpr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.BinExpr()
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(284)
				p.AddOp()
			}
			{
				p.SetState(285)
				p.BinExpr()
			}

		}
		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
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

// IBinExprContext is an interface to support dynamic dispatch.
type IBinExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMulExpr() []IMulExprContext
	MulExpr(i int) IMulExprContext
	AllBitOp() []IBitOpContext
	BitOp(i int) IBitOpContext

	// IsBinExprContext differentiates from other interfaces.
	IsBinExprContext()
}

type BinExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinExprContext() *BinExprContext {
	var p = new(BinExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_binExpr
	return p
}

func InitEmptyBinExprContext(p *BinExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_binExpr
}

func (*BinExprContext) IsBinExprContext() {}

func NewBinExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinExprContext {
	var p = new(BinExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_binExpr

	return p
}

func (s *BinExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BinExprContext) AllMulExpr() []IMulExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMulExprContext); ok {
			len++
		}
	}

	tst := make([]IMulExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMulExprContext); ok {
			tst[i] = t.(IMulExprContext)
			i++
		}
	}

	return tst
}

func (s *BinExprContext) MulExpr(i int) IMulExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulExprContext); ok {
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

	return t.(IMulExprContext)
}

func (s *BinExprContext) AllBitOp() []IBitOpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBitOpContext); ok {
			len++
		}
	}

	tst := make([]IBitOpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBitOpContext); ok {
			tst[i] = t.(IBitOpContext)
			i++
		}
	}

	return tst
}

func (s *BinExprContext) BitOp(i int) IBitOpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBitOpContext); ok {
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

	return t.(IBitOpContext)
}

func (s *BinExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitBinExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) BinExpr() (localctx IBinExprContext) {
	localctx = NewBinExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GsParserRULE_binExpr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(292)
		p.MulExpr()
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(293)
				p.BitOp()
			}
			{
				p.SetState(294)
				p.MulExpr()
			}

		}
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
		if p.HasError() {
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

// IMulExprContext is an interface to support dynamic dispatch.
type IMulExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAtom() []IAtomContext
	Atom(i int) IAtomContext
	AllMulOp() []IMulOpContext
	MulOp(i int) IMulOpContext

	// IsMulExprContext differentiates from other interfaces.
	IsMulExprContext()
}

type MulExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulExprContext() *MulExprContext {
	var p = new(MulExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_mulExpr
	return p
}

func InitEmptyMulExprContext(p *MulExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_mulExpr
}

func (*MulExprContext) IsMulExprContext() {}

func NewMulExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulExprContext {
	var p = new(MulExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_mulExpr

	return p
}

func (s *MulExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MulExprContext) AllAtom() []IAtomContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAtomContext); ok {
			len++
		}
	}

	tst := make([]IAtomContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAtomContext); ok {
			tst[i] = t.(IAtomContext)
			i++
		}
	}

	return tst
}

func (s *MulExprContext) Atom(i int) IAtomContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
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

	return t.(IAtomContext)
}

func (s *MulExprContext) AllMulOp() []IMulOpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMulOpContext); ok {
			len++
		}
	}

	tst := make([]IMulOpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMulOpContext); ok {
			tst[i] = t.(IMulOpContext)
			i++
		}
	}

	return tst
}

func (s *MulExprContext) MulOp(i int) IMulOpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulOpContext); ok {
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

	return t.(IMulOpContext)
}

func (s *MulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitMulExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) MulExpr() (localctx IMulExprContext) {
	localctx = NewMulExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GsParserRULE_mulExpr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Atom()
	}
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(302)
				p.MulOp()
			}
			{
				p.SetState(303)
				p.Atom()
			}

		}
		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
		if p.HasError() {
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

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_atom
	return p
}

func InitEmptyAtomContext(p *AtomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_atom
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) CopyAll(ctx *AtomContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CallAtomContext struct {
	AtomContext
}

func NewCallAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallAtomContext {
	var p = new(CallAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *CallAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallAtomContext) Call() ICallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallContext)
}

func (s *CallAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitCallAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayAtomContext struct {
	AtomContext
}

func NewArrayAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayAtomContext {
	var p = new(ArrayAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *ArrayAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayAtomContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *ArrayAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitArrayAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntAtomContext struct {
	AtomContext
}

func NewIntAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntAtomContext {
	var p = new(IntAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *IntAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntAtomContext) INT() antlr.TerminalNode {
	return s.GetToken(GsParserINT, 0)
}

func (s *IntAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitIntAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type NilAtomContext struct {
	AtomContext
}

func NewNilAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NilAtomContext {
	var p = new(NilAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *NilAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilAtomContext) NIL() antlr.TerminalNode {
	return s.GetToken(GsParserNIL, 0)
}

func (s *NilAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitNilAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddrAtomContext struct {
	AtomContext
}

func NewAddrAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddrAtomContext {
	var p = new(AddrAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *AddrAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddrAtomContext) BITAND() antlr.TerminalNode {
	return s.GetToken(GsParserBITAND, 0)
}

func (s *AddrAtomContext) Lvalue() ILvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *AddrAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitAddrAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type DerefAtomContext struct {
	AtomContext
}

func NewDerefAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DerefAtomContext {
	var p = new(DerefAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *DerefAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DerefAtomContext) MUL() antlr.TerminalNode {
	return s.GetToken(GsParserMUL, 0)
}

func (s *DerefAtomContext) Lvalue() ILvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *DerefAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitDerefAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegAtomContext struct {
	AtomContext
}

func NewNegAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegAtomContext {
	var p = new(NegAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *NegAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegAtomContext) SUB() antlr.TerminalNode {
	return s.GetToken(GsParserSUB, 0)
}

func (s *NegAtomContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *NegAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitNegAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type QidAtomContext struct {
	AtomContext
}

func NewQidAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QidAtomContext {
	var p = new(QidAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *QidAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QidAtomContext) Qid() IQidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQidContext)
}

func (s *QidAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitQidAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type DictAtomContext struct {
	AtomContext
}

func NewDictAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DictAtomContext {
	var p = new(DictAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *DictAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictAtomContext) DictLiteral() IDictLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDictLiteralContext)
}

func (s *DictAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitDictAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloatAtomContext struct {
	AtomContext
}

func NewFloatAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatAtomContext {
	var p = new(FloatAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *FloatAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatAtomContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(GsParserFLOAT, 0)
}

func (s *FloatAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitFloatAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenAtomContext struct {
	AtomContext
}

func NewParenAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenAtomContext {
	var p = new(ParenAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *ParenAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenAtomContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *ParenAtomContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenAtomContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *ParenAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitParenAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringAtomContext struct {
	AtomContext
}

func NewStringAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringAtomContext {
	var p = new(StringAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *StringAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringAtomContext) STRING() antlr.TerminalNode {
	return s.GetToken(GsParserSTRING, 0)
}

func (s *StringAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitStringAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type FalseAtomContext struct {
	AtomContext
}

func NewFalseAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseAtomContext {
	var p = new(FalseAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *FalseAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseAtomContext) FALSE() antlr.TerminalNode {
	return s.GetToken(GsParserFALSE, 0)
}

func (s *FalseAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitFalseAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type InstanceAtomContext struct {
	AtomContext
}

func NewInstanceAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InstanceAtomContext {
	var p = new(InstanceAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *InstanceAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanceAtomContext) Instance() IInstanceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstanceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstanceContext)
}

func (s *InstanceAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitInstanceAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type TrueAtomContext struct {
	AtomContext
}

func NewTrueAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueAtomContext {
	var p = new(TrueAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *TrueAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueAtomContext) TRUE() antlr.TerminalNode {
	return s.GetToken(GsParserTRUE, 0)
}

func (s *TrueAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitTrueAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexAccessAtomContext struct {
	AtomContext
}

func NewIndexAccessAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexAccessAtomContext {
	var p = new(IndexAccessAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *IndexAccessAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexAccessAtomContext) IndexAccess() IIndexAccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexAccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexAccessContext)
}

func (s *IndexAccessAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitIndexAccessAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotAtomContext struct {
	AtomContext
}

func NewNotAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotAtomContext {
	var p = new(NotAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *NotAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotAtomContext) NOT() antlr.TerminalNode {
	return s.GetToken(GsParserNOT, 0)
}

func (s *NotAtomContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitNotAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GsParserRULE_atom)
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNegAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(310)
			p.Match(GsParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(311)
			p.Atom()
		}

	case 2:
		localctx = NewIntAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(312)
			p.Match(GsParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewFloatAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(313)
			p.Match(GsParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewStringAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(314)
			p.Match(GsParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewTrueAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(315)
			p.Match(GsParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewFalseAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(316)
			p.Match(GsParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewNilAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(317)
			p.Match(GsParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewNotAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(318)
			p.Match(GsParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(319)
			p.Expr()
		}

	case 9:
		localctx = NewQidAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(320)
			p.Qid()
		}

	case 10:
		localctx = NewCallAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(321)
			p.Call()
		}

	case 11:
		localctx = NewInstanceAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(322)
			p.Instance()
		}

	case 12:
		localctx = NewArrayAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(323)
			p.ArrayLiteral()
		}

	case 13:
		localctx = NewIndexAccessAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(324)
			p.IndexAccess()
		}

	case 14:
		localctx = NewDictAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(325)
			p.DictLiteral()
		}

	case 15:
		localctx = NewParenAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(326)
			p.Match(GsParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
			p.Expr()
		}
		{
			p.SetState(328)
			p.Match(GsParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		localctx = NewDerefAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(330)
			p.Match(GsParserMUL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.Lvalue()
		}

	case 17:
		localctx = NewAddrAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(332)
			p.Match(GsParserBITAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(333)
			p.Lvalue()
		}

	case antlr.ATNInvalidAltNumber:
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

// ILvalueContext is an interface to support dynamic dispatch.
type ILvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Qid() IQidContext
	IndexAccess() IIndexAccessContext
	MUL() antlr.TerminalNode
	Lvalue() ILvalueContext
	BITAND() antlr.TerminalNode

	// IsLvalueContext differentiates from other interfaces.
	IsLvalueContext()
}

type LvalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLvalueContext() *LvalueContext {
	var p = new(LvalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_lvalue
	return p
}

func InitEmptyLvalueContext(p *LvalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_lvalue
}

func (*LvalueContext) IsLvalueContext() {}

func NewLvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LvalueContext {
	var p = new(LvalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_lvalue

	return p
}

func (s *LvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *LvalueContext) Qid() IQidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQidContext)
}

func (s *LvalueContext) IndexAccess() IIndexAccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexAccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexAccessContext)
}

func (s *LvalueContext) MUL() antlr.TerminalNode {
	return s.GetToken(GsParserMUL, 0)
}

func (s *LvalueContext) Lvalue() ILvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *LvalueContext) BITAND() antlr.TerminalNode {
	return s.GetToken(GsParserBITAND, 0)
}

func (s *LvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LvalueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitLvalue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) Lvalue() (localctx ILvalueContext) {
	localctx = NewLvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GsParserRULE_lvalue)
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(336)
			p.Qid()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(337)
			p.IndexAccess()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(338)
			p.Match(GsParserMUL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.Lvalue()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(340)
			p.Match(GsParserBITAND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.Lvalue()
		}

	case antlr.ATNInvalidAltNumber:
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

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_arrayLiteral
	return p
}

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_arrayLiteral
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(GsParserLBRACK, 0)
}

func (s *ArrayLiteralContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(GsParserRBRACK, 0)
}

func (s *ArrayLiteralContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ArrayLiteralContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ArrayLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *ArrayLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GsParserRULE_arrayLiteral)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.Match(GsParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4864662756495740672) != 0 {
		{
			p.SetState(345)
			p.Expr()
		}
		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(346)
					p.Match(GsParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(347)
					p.Expr()
				}

			}
			p.SetState(352)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GsParserCOMMA {
			{
				p.SetState(353)
				p.Match(GsParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(358)
		p.Match(GsParserRBRACK)
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

// IIndexAccessContext is an interface to support dynamic dispatch.
type IIndexAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Qid() IQidContext
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	Expr() IExprContext
	SliceExpr() ISliceExprContext

	// IsIndexAccessContext differentiates from other interfaces.
	IsIndexAccessContext()
}

type IndexAccessContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexAccessContext() *IndexAccessContext {
	var p = new(IndexAccessContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_indexAccess
	return p
}

func InitEmptyIndexAccessContext(p *IndexAccessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_indexAccess
}

func (*IndexAccessContext) IsIndexAccessContext() {}

func NewIndexAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexAccessContext {
	var p = new(IndexAccessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_indexAccess

	return p
}

func (s *IndexAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexAccessContext) Qid() IQidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQidContext)
}

func (s *IndexAccessContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(GsParserLBRACK, 0)
}

func (s *IndexAccessContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(GsParserRBRACK, 0)
}

func (s *IndexAccessContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexAccessContext) SliceExpr() ISliceExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceExprContext)
}

func (s *IndexAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitIndexAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) IndexAccess() (localctx IIndexAccessContext) {
	localctx = NewIndexAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GsParserRULE_indexAccess)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
		p.Qid()
	}
	{
		p.SetState(361)
		p.Match(GsParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(362)
			p.Expr()
		}

	case 2:
		{
			p.SetState(363)
			p.SliceExpr()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(366)
		p.Match(GsParserRBRACK)
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

// ISliceExprContext is an interface to support dynamic dispatch.
type ISliceExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COLON() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsSliceExprContext differentiates from other interfaces.
	IsSliceExprContext()
}

type SliceExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceExprContext() *SliceExprContext {
	var p = new(SliceExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_sliceExpr
	return p
}

func InitEmptySliceExprContext(p *SliceExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_sliceExpr
}

func (*SliceExprContext) IsSliceExprContext() {}

func NewSliceExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceExprContext {
	var p = new(SliceExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_sliceExpr

	return p
}

func (s *SliceExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceExprContext) COLON() antlr.TerminalNode {
	return s.GetToken(GsParserCOLON, 0)
}

func (s *SliceExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *SliceExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *SliceExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitSliceExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) SliceExpr() (localctx ISliceExprContext) {
	localctx = NewSliceExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GsParserRULE_sliceExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4864662756495740672) != 0 {
		{
			p.SetState(368)
			p.Expr()
		}

	}
	{
		p.SetState(371)
		p.Match(GsParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4864662756495740672) != 0 {
		{
			p.SetState(372)
			p.Expr()
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

// IDictLiteralContext is an interface to support dynamic dispatch.
type IDictLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllDictEntry() []IDictEntryContext
	DictEntry(i int) IDictEntryContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsDictLiteralContext differentiates from other interfaces.
	IsDictLiteralContext()
}

type DictLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictLiteralContext() *DictLiteralContext {
	var p = new(DictLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_dictLiteral
	return p
}

func InitEmptyDictLiteralContext(p *DictLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_dictLiteral
}

func (*DictLiteralContext) IsDictLiteralContext() {}

func NewDictLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictLiteralContext {
	var p = new(DictLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_dictLiteral

	return p
}

func (s *DictLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DictLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GsParserLBRACE, 0)
}

func (s *DictLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GsParserRBRACE, 0)
}

func (s *DictLiteralContext) AllDictEntry() []IDictEntryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDictEntryContext); ok {
			len++
		}
	}

	tst := make([]IDictEntryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDictEntryContext); ok {
			tst[i] = t.(IDictEntryContext)
			i++
		}
	}

	return tst
}

func (s *DictLiteralContext) DictEntry(i int) IDictEntryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictEntryContext); ok {
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

	return t.(IDictEntryContext)
}

func (s *DictLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *DictLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}

func (s *DictLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DictLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitDictLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) DictLiteral() (localctx IDictLiteralContext) {
	localctx = NewDictLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GsParserRULE_dictLiteral)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(375)
		p.Match(GsParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4863887597560137472) != 0 {
		{
			p.SetState(376)
			p.DictEntry()
		}
		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(377)
					p.Match(GsParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(378)
					p.DictEntry()
				}

			}
			p.SetState(383)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GsParserCOMMA {
			{
				p.SetState(384)
				p.Match(GsParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(389)
		p.Match(GsParserRBRACE)
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

// IDictEntryContext is an interface to support dynamic dispatch.
type IDictEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDictEntryContext differentiates from other interfaces.
	IsDictEntryContext()
}

type DictEntryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDictEntryContext() *DictEntryContext {
	var p = new(DictEntryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_dictEntry
	return p
}

func InitEmptyDictEntryContext(p *DictEntryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_dictEntry
}

func (*DictEntryContext) IsDictEntryContext() {}

func NewDictEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DictEntryContext {
	var p = new(DictEntryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_dictEntry

	return p
}

func (s *DictEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *DictEntryContext) CopyAll(ctx *DictEntryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DictEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DictEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConstKeyEntryContext struct {
	DictEntryContext
}

func NewConstKeyEntryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstKeyEntryContext {
	var p = new(ConstKeyEntryContext)

	InitEmptyDictEntryContext(&p.DictEntryContext)
	p.parser = parser
	p.CopyAll(ctx.(*DictEntryContext))

	return p
}

func (s *ConstKeyEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstKeyEntryContext) COLON() antlr.TerminalNode {
	return s.GetToken(GsParserCOLON, 0)
}

func (s *ConstKeyEntryContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConstKeyEntryContext) STRING() antlr.TerminalNode {
	return s.GetToken(GsParserSTRING, 0)
}

func (s *ConstKeyEntryContext) INT() antlr.TerminalNode {
	return s.GetToken(GsParserINT, 0)
}

func (s *ConstKeyEntryContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(GsParserFLOAT, 0)
}

func (s *ConstKeyEntryContext) TRUE() antlr.TerminalNode {
	return s.GetToken(GsParserTRUE, 0)
}

func (s *ConstKeyEntryContext) FALSE() antlr.TerminalNode {
	return s.GetToken(GsParserFALSE, 0)
}

func (s *ConstKeyEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitConstKeyEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdKeyEntryContext struct {
	DictEntryContext
}

func NewIdKeyEntryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdKeyEntryContext {
	var p = new(IdKeyEntryContext)

	InitEmptyDictEntryContext(&p.DictEntryContext)
	p.parser = parser
	p.CopyAll(ctx.(*DictEntryContext))

	return p
}

func (s *IdKeyEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdKeyEntryContext) Qid() IQidContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQidContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQidContext)
}

func (s *IdKeyEntryContext) COLON() antlr.TerminalNode {
	return s.GetToken(GsParserCOLON, 0)
}

func (s *IdKeyEntryContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IdKeyEntryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitIdKeyEntry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) DictEntry() (localctx IDictEntryContext) {
	localctx = NewDictEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GsParserRULE_dictEntry)
	var _la int

	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GsParserTRUE, GsParserFALSE, GsParserINT, GsParserFLOAT, GsParserSTRING:
		localctx = NewConstKeyEntryContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(391)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&252201579132749312) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(392)
			p.Match(GsParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.Expr()
		}

	case GsParserENV, GsParserID:
		localctx = NewIdKeyEntryContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(394)
			p.Qid()
		}
		{
			p.SetState(395)
			p.Match(GsParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.Expr()
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

// IInstanceContext is an interface to support dynamic dispatch.
type IInstanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NEW() antlr.TerminalNode
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsInstanceContext differentiates from other interfaces.
	IsInstanceContext()
}

type InstanceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstanceContext() *InstanceContext {
	var p = new(InstanceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_instance
	return p
}

func InitEmptyInstanceContext(p *InstanceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_instance
}

func (*InstanceContext) IsInstanceContext() {}

func NewInstanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstanceContext {
	var p = new(InstanceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_instance

	return p
}

func (s *InstanceContext) GetParser() antlr.Parser { return s.parser }

func (s *InstanceContext) NEW() antlr.TerminalNode {
	return s.GetToken(GsParserNEW, 0)
}

func (s *InstanceContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GsParserID)
}

func (s *InstanceContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GsParserID, i)
}

func (s *InstanceContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(GsParserLBRACE, 0)
}

func (s *InstanceContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(GsParserRBRACE, 0)
}

func (s *InstanceContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOLON)
}

func (s *InstanceContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOLON, i)
}

func (s *InstanceContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *InstanceContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *InstanceContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *InstanceContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}

func (s *InstanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstanceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitInstance(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) Instance() (localctx IInstanceContext) {
	localctx = NewInstanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GsParserRULE_instance)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(400)
		p.Match(GsParserNEW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(401)
		p.Match(GsParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(402)
		p.Match(GsParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GsParserID {
		{
			p.SetState(403)
			p.Match(GsParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)
			p.Match(GsParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.Expr()
		}
		p.SetState(412)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(406)
					p.Match(GsParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(407)
					p.Match(GsParserID)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(408)
					p.Match(GsParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(409)
					p.Expr()
				}

			}
			p.SetState(414)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GsParserCOMMA {
			{
				p.SetState(415)
				p.Match(GsParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(420)
		p.Match(GsParserRBRACE)
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

// IQidContext is an interface to support dynamic dispatch.
type IQidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Primary() IPrimaryContext
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllRBRACK() []antlr.TerminalNode
	RBRACK(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllSAFE_DOT() []antlr.TerminalNode
	SAFE_DOT(i int) antlr.TerminalNode
	AllLBRACK() []antlr.TerminalNode
	LBRACK(i int) antlr.TerminalNode
	AllSAFE_LBRACK() []antlr.TerminalNode
	SAFE_LBRACK(i int) antlr.TerminalNode

	// IsQidContext differentiates from other interfaces.
	IsQidContext()
}

type QidContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQidContext() *QidContext {
	var p = new(QidContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_qid
	return p
}

func InitEmptyQidContext(p *QidContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_qid
}

func (*QidContext) IsQidContext() {}

func NewQidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QidContext {
	var p = new(QidContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_qid

	return p
}

func (s *QidContext) GetParser() antlr.Parser { return s.parser }

func (s *QidContext) Primary() IPrimaryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *QidContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GsParserID)
}

func (s *QidContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GsParserID, i)
}

func (s *QidContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *QidContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *QidContext) AllRBRACK() []antlr.TerminalNode {
	return s.GetTokens(GsParserRBRACK)
}

func (s *QidContext) RBRACK(i int) antlr.TerminalNode {
	return s.GetToken(GsParserRBRACK, i)
}

func (s *QidContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(GsParserDOT)
}

func (s *QidContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(GsParserDOT, i)
}

func (s *QidContext) AllSAFE_DOT() []antlr.TerminalNode {
	return s.GetTokens(GsParserSAFE_DOT)
}

func (s *QidContext) SAFE_DOT(i int) antlr.TerminalNode {
	return s.GetToken(GsParserSAFE_DOT, i)
}

func (s *QidContext) AllLBRACK() []antlr.TerminalNode {
	return s.GetTokens(GsParserLBRACK)
}

func (s *QidContext) LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(GsParserLBRACK, i)
}

func (s *QidContext) AllSAFE_LBRACK() []antlr.TerminalNode {
	return s.GetTokens(GsParserSAFE_LBRACK)
}

func (s *QidContext) SAFE_LBRACK(i int) antlr.TerminalNode {
	return s.GetToken(GsParserSAFE_LBRACK, i)
}

func (s *QidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitQid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) Qid() (localctx IQidContext) {
	localctx = NewQidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GsParserRULE_qid)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(422)
		p.Primary()
	}
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(429)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case GsParserSAFE_DOT, GsParserDOT:
				{
					p.SetState(423)
					_la = p.GetTokenStream().LA(1)

					if !(_la == GsParserSAFE_DOT || _la == GsParserDOT) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(424)
					p.Match(GsParserID)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case GsParserSAFE_LBRACK, GsParserLBRACK:
				{
					p.SetState(425)
					_la = p.GetTokenStream().LA(1)

					if !(_la == GsParserSAFE_LBRACK || _la == GsParserLBRACK) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(426)
					p.Expr()
				}
				{
					p.SetState(427)
					p.Match(GsParserRBRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}
		p.SetState(433)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
		if p.HasError() {
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

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	ENV() antlr.TerminalNode

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_primary
	return p
}

func InitEmptyPrimaryContext(p *PrimaryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_primary
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) ID() antlr.TerminalNode {
	return s.GetToken(GsParserID, 0)
}

func (s *PrimaryContext) ENV() antlr.TerminalNode {
	return s.GetToken(GsParserENV, 0)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GsParserRULE_primary)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(434)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GsParserENV || _la == GsParserID) {
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

// ICompOpContext is an interface to support dynamic dispatch.
type ICompOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	NEQ() antlr.TerminalNode
	GEQ() antlr.TerminalNode
	LEQ() antlr.TerminalNode

	// IsCompOpContext differentiates from other interfaces.
	IsCompOpContext()
}

type CompOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompOpContext() *CompOpContext {
	var p = new(CompOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_compOp
	return p
}

func InitEmptyCompOpContext(p *CompOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_compOp
}

func (*CompOpContext) IsCompOpContext() {}

func NewCompOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompOpContext {
	var p = new(CompOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_compOp

	return p
}

func (s *CompOpContext) GetParser() antlr.Parser { return s.parser }

func (s *CompOpContext) EQ() antlr.TerminalNode {
	return s.GetToken(GsParserEQ, 0)
}

func (s *CompOpContext) LT() antlr.TerminalNode {
	return s.GetToken(GsParserLT, 0)
}

func (s *CompOpContext) GT() antlr.TerminalNode {
	return s.GetToken(GsParserGT, 0)
}

func (s *CompOpContext) NEQ() antlr.TerminalNode {
	return s.GetToken(GsParserNEQ, 0)
}

func (s *CompOpContext) GEQ() antlr.TerminalNode {
	return s.GetToken(GsParserGEQ, 0)
}

func (s *CompOpContext) LEQ() antlr.TerminalNode {
	return s.GetToken(GsParserLEQ, 0)
}

func (s *CompOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitCompOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) CompOp() (localctx ICompOpContext) {
	localctx = NewCompOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GsParserRULE_compOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(436)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1082331758592) != 0) {
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

// IAddOpContext is an interface to support dynamic dispatch.
type IAddOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode

	// IsAddOpContext differentiates from other interfaces.
	IsAddOpContext()
}

type AddOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddOpContext() *AddOpContext {
	var p = new(AddOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_addOp
	return p
}

func InitEmptyAddOpContext(p *AddOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_addOp
}

func (*AddOpContext) IsAddOpContext() {}

func NewAddOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddOpContext {
	var p = new(AddOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_addOp

	return p
}

func (s *AddOpContext) GetParser() antlr.Parser { return s.parser }

func (s *AddOpContext) ADD() antlr.TerminalNode {
	return s.GetToken(GsParserADD, 0)
}

func (s *AddOpContext) SUB() antlr.TerminalNode {
	return s.GetToken(GsParserSUB, 0)
}

func (s *AddOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitAddOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) AddOp() (localctx IAddOpContext) {
	localctx = NewAddOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GsParserRULE_addOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(438)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GsParserADD || _la == GsParserSUB) {
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

// IBitOpContext is an interface to support dynamic dispatch.
type IBitOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BITAND() antlr.TerminalNode
	BITOR() antlr.TerminalNode
	XOR() antlr.TerminalNode

	// IsBitOpContext differentiates from other interfaces.
	IsBitOpContext()
}

type BitOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBitOpContext() *BitOpContext {
	var p = new(BitOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_bitOp
	return p
}

func InitEmptyBitOpContext(p *BitOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_bitOp
}

func (*BitOpContext) IsBitOpContext() {}

func NewBitOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitOpContext {
	var p = new(BitOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_bitOp

	return p
}

func (s *BitOpContext) GetParser() antlr.Parser { return s.parser }

func (s *BitOpContext) BITAND() antlr.TerminalNode {
	return s.GetToken(GsParserBITAND, 0)
}

func (s *BitOpContext) BITOR() antlr.TerminalNode {
	return s.GetToken(GsParserBITOR, 0)
}

func (s *BitOpContext) XOR() antlr.TerminalNode {
	return s.GetToken(GsParserXOR, 0)
}

func (s *BitOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BitOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitBitOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) BitOp() (localctx IBitOpContext) {
	localctx = NewBitOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GsParserRULE_bitOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(440)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7696581394432) != 0) {
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

// IMulOpContext is an interface to support dynamic dispatch.
type IMulOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MOD() antlr.TerminalNode

	// IsMulOpContext differentiates from other interfaces.
	IsMulOpContext()
}

type MulOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulOpContext() *MulOpContext {
	var p = new(MulOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_mulOp
	return p
}

func InitEmptyMulOpContext(p *MulOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_mulOp
}

func (*MulOpContext) IsMulOpContext() {}

func NewMulOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulOpContext {
	var p = new(MulOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_mulOp

	return p
}

func (s *MulOpContext) GetParser() antlr.Parser { return s.parser }

func (s *MulOpContext) MUL() antlr.TerminalNode {
	return s.GetToken(GsParserMUL, 0)
}

func (s *MulOpContext) DIV() antlr.TerminalNode {
	return s.GetToken(GsParserDIV, 0)
}

func (s *MulOpContext) MOD() antlr.TerminalNode {
	return s.GetToken(GsParserMOD, 0)
}

func (s *MulOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitMulOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) MulOp() (localctx IMulOpContext) {
	localctx = NewMulOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GsParserRULE_mulOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(442)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15032385536) != 0) {
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
