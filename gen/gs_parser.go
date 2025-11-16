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
		"", "'='", "'+='", "'-='", "'*='", "'/='", "'%='", "'$'", "'true'",
		"'false'", "'nil'", "'&&'", "'||'", "'!'", "'if'", "'else'", "'for'",
		"'range'", "'return'", "'func'", "'type'", "'struct'", "'new'", "'break'",
		"'continue'", "'len'", "'append'", "'delete'", "'copy'", "'toString'",
		"'print'", "'printf'", "'println'", "'uint'", "'uint8'", "'uint16'",
		"'uint32'", "'uint64'", "'int'", "'int8'", "'int16'", "'int32'", "'int64'",
		"'float32'", "'float64'", "'string'", "'bool'", "'?.'", "'?['", "'+'",
		"'-'", "'*'", "'/'", "'%'", "'=='", "'<'", "'>'", "'>='", "'<='", "'!='",
		"'&'", "'|'", "'^'", "'++'", "'--'", "'.'", "'['", "'('", "')'", "'{'",
		"'}'", "']'", "':'", "';'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "ENV", "TRUE", "FALSE", "NIL", "AND", "OR",
		"NOT", "IF", "ELSE", "FOR", "RANGE", "RETURN", "FUNC", "TYPE", "STRUCT",
		"NEW", "BREAK", "CONTINUE", "LEN", "APPEND", "DELETE", "COPY", "TOSTRING",
		"PRINT", "PRINTF", "PRINTLN", "UINT", "UINT8", "UINT16", "UINT32", "UINT64",
		"INTS", "INT8", "INT16", "INT32", "INT64", "FLOAT32", "FLOAT64", "STRINGS",
		"BOOL", "SAFE_DOT", "SAFE_LBRACK", "ADD", "SUB", "MUL", "DIV", "MOD",
		"EQ", "LT", "GT", "GEQ", "LEQ", "NEQ", "BITAND", "BITOR", "XOR", "INCR",
		"DECR", "DOT", "LBRACK", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "RBRACK",
		"COLON", "SEMICOLON", "COMMA", "INT", "FLOAT", "STRING", "WS", "NEWLINE",
		"SL_COMMENT", "ML_COMMENT", "ID",
	}
	staticData.RuleNames = []string{
		"program", "structDefinition", "functionDefinition", "block", "statement",
		"assign", "incrDecr", "builtinCall", "iterVar", "forInit", "forUpdate",
		"selfAssign", "updateItem", "selfAssignOp", "call", "expr", "logicalOrExpr",
		"logicalAndExpr", "comparisonExpr", "addExpr", "binExpr", "mulExpr",
		"atom", "lvalue", "arrayLiteral", "sliceExpr", "dictLiteral", "dictEntry",
		"instance", "qid", "accessor", "primary", "compOp", "addOp", "bitOp",
		"mulOp",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 82, 502, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0, 1, 0,
		1, 0, 4, 0, 76, 8, 0, 11, 0, 12, 0, 77, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 5, 1, 89, 8, 1, 10, 1, 12, 1, 92, 9, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 102, 8, 2, 10, 2, 12, 2, 105,
		9, 2, 3, 2, 107, 8, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 4, 3, 115, 8,
		3, 11, 3, 12, 3, 116, 1, 3, 3, 3, 120, 8, 3, 5, 3, 122, 8, 3, 10, 3, 12,
		3, 125, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 5, 4, 138, 8, 4, 10, 4, 12, 4, 141, 9, 4, 3, 4, 143, 8, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 3, 4, 149, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 155,
		8, 4, 1, 4, 1, 4, 3, 4, 159, 8, 4, 1, 4, 1, 4, 3, 4, 163, 8, 4, 1, 4, 1,
		4, 3, 4, 167, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 185, 8, 4, 1, 5, 1,
		5, 1, 5, 5, 5, 190, 8, 5, 10, 5, 12, 5, 193, 9, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 5, 5, 199, 8, 5, 10, 5, 12, 5, 202, 9, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 219,
		8, 7, 10, 7, 12, 7, 222, 9, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 248, 8, 7, 10, 7, 12, 7, 251, 9, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 260, 8, 7, 10, 7, 12, 7,
		263, 9, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 272, 8, 7, 10,
		7, 12, 7, 275, 9, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 284,
		8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 290, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 3, 12, 303, 8, 12, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 312, 8, 14, 10, 14,
		12, 14, 315, 9, 14, 1, 14, 3, 14, 318, 8, 14, 3, 14, 320, 8, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 5, 16, 329, 8, 16, 10, 16, 12,
		16, 332, 9, 16, 1, 17, 1, 17, 1, 17, 5, 17, 337, 8, 17, 10, 17, 12, 17,
		340, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 346, 8, 18, 1, 19, 1, 19,
		1, 19, 1, 19, 5, 19, 352, 8, 19, 10, 19, 12, 19, 355, 9, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 5, 20, 361, 8, 20, 10, 20, 12, 20, 364, 9, 20, 1, 21,
		1, 21, 1, 21, 1, 21, 5, 21, 370, 8, 21, 10, 21, 12, 21, 373, 9, 21, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 3, 22, 397, 8, 22, 1, 23, 1, 23, 1, 23, 3, 23, 402, 8, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 5, 24, 408, 8, 24, 10, 24, 12, 24, 411, 9, 24, 1, 24,
		3, 24, 414, 8, 24, 3, 24, 416, 8, 24, 1, 24, 1, 24, 1, 25, 3, 25, 421,
		8, 25, 1, 25, 1, 25, 3, 25, 425, 8, 25, 1, 26, 1, 26, 1, 26, 1, 26, 5,
		26, 431, 8, 26, 10, 26, 12, 26, 434, 9, 26, 1, 26, 3, 26, 437, 8, 26, 3,
		26, 439, 8, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 3, 27, 450, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 5, 28, 462, 8, 28, 10, 28, 12, 28, 465, 9, 28,
		1, 28, 3, 28, 468, 8, 28, 3, 28, 470, 8, 28, 1, 28, 1, 28, 1, 29, 1, 29,
		5, 29, 476, 8, 29, 10, 29, 12, 29, 479, 9, 29, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 3, 30, 486, 8, 30, 1, 30, 1, 30, 3, 30, 490, 8, 30, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 0, 0,
		36, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
		0, 12, 2, 0, 73, 73, 79, 79, 1, 0, 63, 64, 1, 0, 33, 46, 1, 0, 2, 6, 2,
		0, 8, 9, 75, 77, 2, 0, 47, 47, 65, 65, 2, 0, 48, 48, 66, 66, 2, 0, 7, 7,
		82, 82, 1, 0, 54, 59, 1, 0, 49, 50, 1, 0, 60, 62, 1, 0, 51, 53, 552, 0,
		75, 1, 0, 0, 0, 2, 81, 1, 0, 0, 0, 4, 95, 1, 0, 0, 0, 6, 111, 1, 0, 0,
		0, 8, 184, 1, 0, 0, 0, 10, 186, 1, 0, 0, 0, 12, 203, 1, 0, 0, 0, 14, 283,
		1, 0, 0, 0, 16, 289, 1, 0, 0, 0, 18, 291, 1, 0, 0, 0, 20, 293, 1, 0, 0,
		0, 22, 295, 1, 0, 0, 0, 24, 302, 1, 0, 0, 0, 26, 304, 1, 0, 0, 0, 28, 306,
		1, 0, 0, 0, 30, 323, 1, 0, 0, 0, 32, 325, 1, 0, 0, 0, 34, 333, 1, 0, 0,
		0, 36, 341, 1, 0, 0, 0, 38, 347, 1, 0, 0, 0, 40, 356, 1, 0, 0, 0, 42, 365,
		1, 0, 0, 0, 44, 396, 1, 0, 0, 0, 46, 401, 1, 0, 0, 0, 48, 403, 1, 0, 0,
		0, 50, 420, 1, 0, 0, 0, 52, 426, 1, 0, 0, 0, 54, 449, 1, 0, 0, 0, 56, 451,
		1, 0, 0, 0, 58, 473, 1, 0, 0, 0, 60, 489, 1, 0, 0, 0, 62, 491, 1, 0, 0,
		0, 64, 493, 1, 0, 0, 0, 66, 495, 1, 0, 0, 0, 68, 497, 1, 0, 0, 0, 70, 499,
		1, 0, 0, 0, 72, 76, 3, 4, 2, 0, 73, 76, 3, 2, 1, 0, 74, 76, 3, 8, 4, 0,
		75, 72, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 74, 1, 0, 0, 0, 76, 77, 1,
		0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79,
		80, 5, 0, 0, 1, 80, 1, 1, 0, 0, 0, 81, 82, 5, 20, 0, 0, 82, 83, 5, 82,
		0, 0, 83, 84, 5, 21, 0, 0, 84, 85, 5, 69, 0, 0, 85, 90, 5, 82, 0, 0, 86,
		87, 5, 74, 0, 0, 87, 89, 5, 82, 0, 0, 88, 86, 1, 0, 0, 0, 89, 92, 1, 0,
		0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 93, 1, 0, 0, 0, 92, 90,
		1, 0, 0, 0, 93, 94, 5, 70, 0, 0, 94, 3, 1, 0, 0, 0, 95, 96, 5, 19, 0, 0,
		96, 97, 5, 82, 0, 0, 97, 106, 5, 67, 0, 0, 98, 103, 5, 82, 0, 0, 99, 100,
		5, 74, 0, 0, 100, 102, 5, 82, 0, 0, 101, 99, 1, 0, 0, 0, 102, 105, 1, 0,
		0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0,
		105, 103, 1, 0, 0, 0, 106, 98, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 108,
		1, 0, 0, 0, 108, 109, 5, 68, 0, 0, 109, 110, 3, 6, 3, 0, 110, 5, 1, 0,
		0, 0, 111, 123, 5, 69, 0, 0, 112, 119, 3, 8, 4, 0, 113, 115, 7, 0, 0, 0,
		114, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116,
		117, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118, 120, 5, 0, 0, 1, 119, 114,
		1, 0, 0, 0, 119, 118, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 122, 1, 0,
		0, 0, 121, 112, 1, 0, 0, 0, 122, 125, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0,
		123, 124, 1, 0, 0, 0, 124, 126, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 126,
		127, 5, 70, 0, 0, 127, 7, 1, 0, 0, 0, 128, 185, 5, 73, 0, 0, 129, 185,
		3, 2, 1, 0, 130, 185, 3, 10, 5, 0, 131, 185, 3, 22, 11, 0, 132, 185, 3,
		12, 6, 0, 133, 142, 5, 18, 0, 0, 134, 139, 3, 30, 15, 0, 135, 136, 5, 74,
		0, 0, 136, 138, 3, 30, 15, 0, 137, 135, 1, 0, 0, 0, 138, 141, 1, 0, 0,
		0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141,
		139, 1, 0, 0, 0, 142, 134, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 185,
		1, 0, 0, 0, 144, 148, 5, 14, 0, 0, 145, 146, 3, 10, 5, 0, 146, 147, 5,
		73, 0, 0, 147, 149, 1, 0, 0, 0, 148, 145, 1, 0, 0, 0, 148, 149, 1, 0, 0,
		0, 149, 150, 1, 0, 0, 0, 150, 151, 3, 30, 15, 0, 151, 154, 3, 6, 3, 0,
		152, 153, 5, 15, 0, 0, 153, 155, 3, 6, 3, 0, 154, 152, 1, 0, 0, 0, 154,
		155, 1, 0, 0, 0, 155, 185, 1, 0, 0, 0, 156, 158, 5, 16, 0, 0, 157, 159,
		3, 18, 9, 0, 158, 157, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 160, 1, 0,
		0, 0, 160, 162, 5, 73, 0, 0, 161, 163, 3, 30, 15, 0, 162, 161, 1, 0, 0,
		0, 162, 163, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 166, 5, 73, 0, 0, 165,
		167, 3, 20, 10, 0, 166, 165, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168,
		1, 0, 0, 0, 168, 185, 3, 6, 3, 0, 169, 170, 5, 16, 0, 0, 170, 171, 3, 16,
		8, 0, 171, 172, 5, 1, 0, 0, 172, 173, 5, 17, 0, 0, 173, 174, 3, 30, 15,
		0, 174, 175, 3, 6, 3, 0, 175, 185, 1, 0, 0, 0, 176, 177, 5, 16, 0, 0, 177,
		178, 3, 30, 15, 0, 178, 179, 3, 6, 3, 0, 179, 185, 1, 0, 0, 0, 180, 185,
		3, 14, 7, 0, 181, 185, 3, 28, 14, 0, 182, 185, 5, 23, 0, 0, 183, 185, 5,
		24, 0, 0, 184, 128, 1, 0, 0, 0, 184, 129, 1, 0, 0, 0, 184, 130, 1, 0, 0,
		0, 184, 131, 1, 0, 0, 0, 184, 132, 1, 0, 0, 0, 184, 133, 1, 0, 0, 0, 184,
		144, 1, 0, 0, 0, 184, 156, 1, 0, 0, 0, 184, 169, 1, 0, 0, 0, 184, 176,
		1, 0, 0, 0, 184, 180, 1, 0, 0, 0, 184, 181, 1, 0, 0, 0, 184, 182, 1, 0,
		0, 0, 184, 183, 1, 0, 0, 0, 185, 9, 1, 0, 0, 0, 186, 191, 3, 46, 23, 0,
		187, 188, 5, 74, 0, 0, 188, 190, 3, 46, 23, 0, 189, 187, 1, 0, 0, 0, 190,
		193, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 194,
		1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 194, 195, 5, 1, 0, 0, 195, 200, 3, 30,
		15, 0, 196, 197, 5, 74, 0, 0, 197, 199, 3, 30, 15, 0, 198, 196, 1, 0, 0,
		0, 199, 202, 1, 0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201,
		11, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 203, 204, 3, 46, 23, 0, 204, 205,
		7, 1, 0, 0, 205, 13, 1, 0, 0, 0, 206, 207, 5, 25, 0, 0, 207, 208, 5, 67,
		0, 0, 208, 209, 3, 30, 15, 0, 209, 210, 5, 68, 0, 0, 210, 284, 1, 0, 0,
		0, 211, 212, 5, 26, 0, 0, 212, 213, 5, 67, 0, 0, 213, 214, 3, 30, 15, 0,
		214, 215, 5, 74, 0, 0, 215, 220, 3, 30, 15, 0, 216, 217, 5, 74, 0, 0, 217,
		219, 3, 30, 15, 0, 218, 216, 1, 0, 0, 0, 219, 222, 1, 0, 0, 0, 220, 218,
		1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 223, 1, 0, 0, 0, 222, 220, 1, 0,
		0, 0, 223, 224, 5, 68, 0, 0, 224, 284, 1, 0, 0, 0, 225, 226, 5, 27, 0,
		0, 226, 227, 5, 67, 0, 0, 227, 228, 3, 30, 15, 0, 228, 229, 5, 74, 0, 0,
		229, 230, 3, 30, 15, 0, 230, 231, 5, 68, 0, 0, 231, 284, 1, 0, 0, 0, 232,
		233, 5, 28, 0, 0, 233, 234, 5, 67, 0, 0, 234, 235, 3, 30, 15, 0, 235, 236,
		5, 68, 0, 0, 236, 284, 1, 0, 0, 0, 237, 238, 5, 29, 0, 0, 238, 239, 5,
		67, 0, 0, 239, 240, 3, 30, 15, 0, 240, 241, 5, 68, 0, 0, 241, 284, 1, 0,
		0, 0, 242, 243, 5, 30, 0, 0, 243, 244, 5, 67, 0, 0, 244, 249, 3, 30, 15,
		0, 245, 246, 5, 74, 0, 0, 246, 248, 3, 30, 15, 0, 247, 245, 1, 0, 0, 0,
		248, 251, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250,
		252, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 252, 253, 5, 68, 0, 0, 253, 284,
		1, 0, 0, 0, 254, 255, 5, 31, 0, 0, 255, 256, 5, 67, 0, 0, 256, 261, 3,
		30, 15, 0, 257, 258, 5, 74, 0, 0, 258, 260, 3, 30, 15, 0, 259, 257, 1,
		0, 0, 0, 260, 263, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261, 262, 1, 0, 0,
		0, 262, 264, 1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 264, 265, 5, 68, 0, 0, 265,
		284, 1, 0, 0, 0, 266, 267, 5, 32, 0, 0, 267, 268, 5, 67, 0, 0, 268, 273,
		3, 30, 15, 0, 269, 270, 5, 74, 0, 0, 270, 272, 3, 30, 15, 0, 271, 269,
		1, 0, 0, 0, 272, 275, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0,
		0, 0, 274, 276, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 276, 277, 5, 68, 0, 0,
		277, 284, 1, 0, 0, 0, 278, 279, 7, 2, 0, 0, 279, 280, 5, 67, 0, 0, 280,
		281, 3, 30, 15, 0, 281, 282, 5, 68, 0, 0, 282, 284, 1, 0, 0, 0, 283, 206,
		1, 0, 0, 0, 283, 211, 1, 0, 0, 0, 283, 225, 1, 0, 0, 0, 283, 232, 1, 0,
		0, 0, 283, 237, 1, 0, 0, 0, 283, 242, 1, 0, 0, 0, 283, 254, 1, 0, 0, 0,
		283, 266, 1, 0, 0, 0, 283, 278, 1, 0, 0, 0, 284, 15, 1, 0, 0, 0, 285, 290,
		5, 82, 0, 0, 286, 287, 5, 82, 0, 0, 287, 288, 5, 74, 0, 0, 288, 290, 5,
		82, 0, 0, 289, 285, 1, 0, 0, 0, 289, 286, 1, 0, 0, 0, 290, 17, 1, 0, 0,
		0, 291, 292, 3, 10, 5, 0, 292, 19, 1, 0, 0, 0, 293, 294, 3, 24, 12, 0,
		294, 21, 1, 0, 0, 0, 295, 296, 3, 46, 23, 0, 296, 297, 3, 26, 13, 0, 297,
		298, 3, 30, 15, 0, 298, 23, 1, 0, 0, 0, 299, 303, 3, 22, 11, 0, 300, 303,
		3, 12, 6, 0, 301, 303, 3, 10, 5, 0, 302, 299, 1, 0, 0, 0, 302, 300, 1,
		0, 0, 0, 302, 301, 1, 0, 0, 0, 303, 25, 1, 0, 0, 0, 304, 305, 7, 3, 0,
		0, 305, 27, 1, 0, 0, 0, 306, 307, 5, 82, 0, 0, 307, 319, 5, 67, 0, 0, 308,
		313, 3, 30, 15, 0, 309, 310, 5, 74, 0, 0, 310, 312, 3, 30, 15, 0, 311,
		309, 1, 0, 0, 0, 312, 315, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 313, 314,
		1, 0, 0, 0, 314, 317, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 316, 318, 5, 74,
		0, 0, 317, 316, 1, 0, 0, 0, 317, 318, 1, 0, 0, 0, 318, 320, 1, 0, 0, 0,
		319, 308, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321,
		322, 5, 68, 0, 0, 322, 29, 1, 0, 0, 0, 323, 324, 3, 32, 16, 0, 324, 31,
		1, 0, 0, 0, 325, 330, 3, 34, 17, 0, 326, 327, 5, 12, 0, 0, 327, 329, 3,
		34, 17, 0, 328, 326, 1, 0, 0, 0, 329, 332, 1, 0, 0, 0, 330, 328, 1, 0,
		0, 0, 330, 331, 1, 0, 0, 0, 331, 33, 1, 0, 0, 0, 332, 330, 1, 0, 0, 0,
		333, 338, 3, 36, 18, 0, 334, 335, 5, 11, 0, 0, 335, 337, 3, 36, 18, 0,
		336, 334, 1, 0, 0, 0, 337, 340, 1, 0, 0, 0, 338, 336, 1, 0, 0, 0, 338,
		339, 1, 0, 0, 0, 339, 35, 1, 0, 0, 0, 340, 338, 1, 0, 0, 0, 341, 345, 3,
		38, 19, 0, 342, 343, 3, 64, 32, 0, 343, 344, 3, 38, 19, 0, 344, 346, 1,
		0, 0, 0, 345, 342, 1, 0, 0, 0, 345, 346, 1, 0, 0, 0, 346, 37, 1, 0, 0,
		0, 347, 353, 3, 40, 20, 0, 348, 349, 3, 66, 33, 0, 349, 350, 3, 40, 20,
		0, 350, 352, 1, 0, 0, 0, 351, 348, 1, 0, 0, 0, 352, 355, 1, 0, 0, 0, 353,
		351, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354, 39, 1, 0, 0, 0, 355, 353, 1,
		0, 0, 0, 356, 362, 3, 42, 21, 0, 357, 358, 3, 68, 34, 0, 358, 359, 3, 42,
		21, 0, 359, 361, 1, 0, 0, 0, 360, 357, 1, 0, 0, 0, 361, 364, 1, 0, 0, 0,
		362, 360, 1, 0, 0, 0, 362, 363, 1, 0, 0, 0, 363, 41, 1, 0, 0, 0, 364, 362,
		1, 0, 0, 0, 365, 371, 3, 44, 22, 0, 366, 367, 3, 70, 35, 0, 367, 368, 3,
		44, 22, 0, 368, 370, 1, 0, 0, 0, 369, 366, 1, 0, 0, 0, 370, 373, 1, 0,
		0, 0, 371, 369, 1, 0, 0, 0, 371, 372, 1, 0, 0, 0, 372, 43, 1, 0, 0, 0,
		373, 371, 1, 0, 0, 0, 374, 375, 5, 50, 0, 0, 375, 397, 3, 44, 22, 0, 376,
		397, 5, 75, 0, 0, 377, 397, 5, 76, 0, 0, 378, 397, 5, 77, 0, 0, 379, 397,
		5, 8, 0, 0, 380, 397, 5, 9, 0, 0, 381, 397, 5, 10, 0, 0, 382, 383, 5, 13,
		0, 0, 383, 397, 3, 30, 15, 0, 384, 397, 3, 58, 29, 0, 385, 397, 3, 14,
		7, 0, 386, 397, 3, 28, 14, 0, 387, 397, 3, 56, 28, 0, 388, 397, 3, 48,
		24, 0, 389, 397, 3, 52, 26, 0, 390, 391, 5, 67, 0, 0, 391, 392, 3, 30,
		15, 0, 392, 393, 5, 68, 0, 0, 393, 397, 1, 0, 0, 0, 394, 395, 5, 51, 0,
		0, 395, 397, 3, 46, 23, 0, 396, 374, 1, 0, 0, 0, 396, 376, 1, 0, 0, 0,
		396, 377, 1, 0, 0, 0, 396, 378, 1, 0, 0, 0, 396, 379, 1, 0, 0, 0, 396,
		380, 1, 0, 0, 0, 396, 381, 1, 0, 0, 0, 396, 382, 1, 0, 0, 0, 396, 384,
		1, 0, 0, 0, 396, 385, 1, 0, 0, 0, 396, 386, 1, 0, 0, 0, 396, 387, 1, 0,
		0, 0, 396, 388, 1, 0, 0, 0, 396, 389, 1, 0, 0, 0, 396, 390, 1, 0, 0, 0,
		396, 394, 1, 0, 0, 0, 397, 45, 1, 0, 0, 0, 398, 402, 3, 58, 29, 0, 399,
		400, 5, 51, 0, 0, 400, 402, 3, 46, 23, 0, 401, 398, 1, 0, 0, 0, 401, 399,
		1, 0, 0, 0, 402, 47, 1, 0, 0, 0, 403, 415, 5, 66, 0, 0, 404, 409, 3, 30,
		15, 0, 405, 406, 5, 74, 0, 0, 406, 408, 3, 30, 15, 0, 407, 405, 1, 0, 0,
		0, 408, 411, 1, 0, 0, 0, 409, 407, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410,
		413, 1, 0, 0, 0, 411, 409, 1, 0, 0, 0, 412, 414, 5, 74, 0, 0, 413, 412,
		1, 0, 0, 0, 413, 414, 1, 0, 0, 0, 414, 416, 1, 0, 0, 0, 415, 404, 1, 0,
		0, 0, 415, 416, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 418, 5, 71, 0, 0,
		418, 49, 1, 0, 0, 0, 419, 421, 3, 30, 15, 0, 420, 419, 1, 0, 0, 0, 420,
		421, 1, 0, 0, 0, 421, 422, 1, 0, 0, 0, 422, 424, 5, 72, 0, 0, 423, 425,
		3, 30, 15, 0, 424, 423, 1, 0, 0, 0, 424, 425, 1, 0, 0, 0, 425, 51, 1, 0,
		0, 0, 426, 438, 5, 69, 0, 0, 427, 432, 3, 54, 27, 0, 428, 429, 5, 74, 0,
		0, 429, 431, 3, 54, 27, 0, 430, 428, 1, 0, 0, 0, 431, 434, 1, 0, 0, 0,
		432, 430, 1, 0, 0, 0, 432, 433, 1, 0, 0, 0, 433, 436, 1, 0, 0, 0, 434,
		432, 1, 0, 0, 0, 435, 437, 5, 74, 0, 0, 436, 435, 1, 0, 0, 0, 436, 437,
		1, 0, 0, 0, 437, 439, 1, 0, 0, 0, 438, 427, 1, 0, 0, 0, 438, 439, 1, 0,
		0, 0, 439, 440, 1, 0, 0, 0, 440, 441, 5, 70, 0, 0, 441, 53, 1, 0, 0, 0,
		442, 443, 7, 4, 0, 0, 443, 444, 5, 72, 0, 0, 444, 450, 3, 30, 15, 0, 445,
		446, 3, 46, 23, 0, 446, 447, 5, 72, 0, 0, 447, 448, 3, 30, 15, 0, 448,
		450, 1, 0, 0, 0, 449, 442, 1, 0, 0, 0, 449, 445, 1, 0, 0, 0, 450, 55, 1,
		0, 0, 0, 451, 452, 5, 22, 0, 0, 452, 453, 5, 82, 0, 0, 453, 469, 5, 69,
		0, 0, 454, 455, 5, 82, 0, 0, 455, 456, 5, 72, 0, 0, 456, 463, 3, 30, 15,
		0, 457, 458, 5, 74, 0, 0, 458, 459, 5, 82, 0, 0, 459, 460, 5, 72, 0, 0,
		460, 462, 3, 30, 15, 0, 461, 457, 1, 0, 0, 0, 462, 465, 1, 0, 0, 0, 463,
		461, 1, 0, 0, 0, 463, 464, 1, 0, 0, 0, 464, 467, 1, 0, 0, 0, 465, 463,
		1, 0, 0, 0, 466, 468, 5, 74, 0, 0, 467, 466, 1, 0, 0, 0, 467, 468, 1, 0,
		0, 0, 468, 470, 1, 0, 0, 0, 469, 454, 1, 0, 0, 0, 469, 470, 1, 0, 0, 0,
		470, 471, 1, 0, 0, 0, 471, 472, 5, 70, 0, 0, 472, 57, 1, 0, 0, 0, 473,
		477, 3, 62, 31, 0, 474, 476, 3, 60, 30, 0, 475, 474, 1, 0, 0, 0, 476, 479,
		1, 0, 0, 0, 477, 475, 1, 0, 0, 0, 477, 478, 1, 0, 0, 0, 478, 59, 1, 0,
		0, 0, 479, 477, 1, 0, 0, 0, 480, 481, 7, 5, 0, 0, 481, 490, 5, 82, 0, 0,
		482, 485, 7, 6, 0, 0, 483, 486, 3, 30, 15, 0, 484, 486, 3, 50, 25, 0, 485,
		483, 1, 0, 0, 0, 485, 484, 1, 0, 0, 0, 486, 487, 1, 0, 0, 0, 487, 488,
		5, 71, 0, 0, 488, 490, 1, 0, 0, 0, 489, 480, 1, 0, 0, 0, 489, 482, 1, 0,
		0, 0, 490, 61, 1, 0, 0, 0, 491, 492, 7, 7, 0, 0, 492, 63, 1, 0, 0, 0, 493,
		494, 7, 8, 0, 0, 494, 65, 1, 0, 0, 0, 495, 496, 7, 9, 0, 0, 496, 67, 1,
		0, 0, 0, 497, 498, 7, 10, 0, 0, 498, 69, 1, 0, 0, 0, 499, 500, 7, 11, 0,
		0, 500, 71, 1, 0, 0, 0, 51, 75, 77, 90, 103, 106, 116, 119, 123, 139, 142,
		148, 154, 158, 162, 166, 184, 191, 200, 220, 249, 261, 273, 283, 289, 302,
		313, 317, 319, 330, 338, 345, 353, 362, 371, 396, 401, 409, 413, 415, 420,
		424, 432, 436, 438, 449, 463, 467, 469, 477, 485, 489,
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
	GsParserENV         = 7
	GsParserTRUE        = 8
	GsParserFALSE       = 9
	GsParserNIL         = 10
	GsParserAND         = 11
	GsParserOR          = 12
	GsParserNOT         = 13
	GsParserIF          = 14
	GsParserELSE        = 15
	GsParserFOR         = 16
	GsParserRANGE       = 17
	GsParserRETURN      = 18
	GsParserFUNC        = 19
	GsParserTYPE        = 20
	GsParserSTRUCT      = 21
	GsParserNEW         = 22
	GsParserBREAK       = 23
	GsParserCONTINUE    = 24
	GsParserLEN         = 25
	GsParserAPPEND      = 26
	GsParserDELETE      = 27
	GsParserCOPY        = 28
	GsParserTOSTRING    = 29
	GsParserPRINT       = 30
	GsParserPRINTF      = 31
	GsParserPRINTLN     = 32
	GsParserUINT        = 33
	GsParserUINT8       = 34
	GsParserUINT16      = 35
	GsParserUINT32      = 36
	GsParserUINT64      = 37
	GsParserINTS        = 38
	GsParserINT8        = 39
	GsParserINT16       = 40
	GsParserINT32       = 41
	GsParserINT64       = 42
	GsParserFLOAT32     = 43
	GsParserFLOAT64     = 44
	GsParserSTRINGS     = 45
	GsParserBOOL        = 46
	GsParserSAFE_DOT    = 47
	GsParserSAFE_LBRACK = 48
	GsParserADD         = 49
	GsParserSUB         = 50
	GsParserMUL         = 51
	GsParserDIV         = 52
	GsParserMOD         = 53
	GsParserEQ          = 54
	GsParserLT          = 55
	GsParserGT          = 56
	GsParserGEQ         = 57
	GsParserLEQ         = 58
	GsParserNEQ         = 59
	GsParserBITAND      = 60
	GsParserBITOR       = 61
	GsParserXOR         = 62
	GsParserINCR        = 63
	GsParserDECR        = 64
	GsParserDOT         = 65
	GsParserLBRACK      = 66
	GsParserLPAREN      = 67
	GsParserRPAREN      = 68
	GsParserLBRACE      = 69
	GsParserRBRACE      = 70
	GsParserRBRACK      = 71
	GsParserCOLON       = 72
	GsParserSEMICOLON   = 73
	GsParserCOMMA       = 74
	GsParserINT         = 75
	GsParserFLOAT       = 76
	GsParserSTRING      = 77
	GsParserWS          = 78
	GsParserNEWLINE     = 79
	GsParserSL_COMMENT  = 80
	GsParserML_COMMENT  = 81
	GsParserID          = 82
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
	GsParserRULE_builtinCall        = 7
	GsParserRULE_iterVar            = 8
	GsParserRULE_forInit            = 9
	GsParserRULE_forUpdate          = 10
	GsParserRULE_selfAssign         = 11
	GsParserRULE_updateItem         = 12
	GsParserRULE_selfAssignOp       = 13
	GsParserRULE_call               = 14
	GsParserRULE_expr               = 15
	GsParserRULE_logicalOrExpr      = 16
	GsParserRULE_logicalAndExpr     = 17
	GsParserRULE_comparisonExpr     = 18
	GsParserRULE_addExpr            = 19
	GsParserRULE_binExpr            = 20
	GsParserRULE_mulExpr            = 21
	GsParserRULE_atom               = 22
	GsParserRULE_lvalue             = 23
	GsParserRULE_arrayLiteral       = 24
	GsParserRULE_sliceExpr          = 25
	GsParserRULE_dictLiteral        = 26
	GsParserRULE_dictEntry          = 27
	GsParserRULE_instance           = 28
	GsParserRULE_qid                = 29
	GsParserRULE_accessor           = 30
	GsParserRULE_primary            = 31
	GsParserRULE_compOp             = 32
	GsParserRULE_addOp              = 33
	GsParserRULE_bitOp              = 34
	GsParserRULE_mulOp              = 35
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
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2392537295569024) != 0) || _la == GsParserSEMICOLON || _la == GsParserID {
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(72)
				p.FunctionDefinition()
			}

		case 2:
			{
				p.SetState(73)
				p.StructDefinition()
			}

		case 3:
			{
				p.SetState(74)
				p.Statement()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(79)
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
		p.SetState(81)
		p.Match(GsParserTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Match(GsParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.Match(GsParserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Match(GsParserLBRACE)
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

	for _la == GsParserCOMMA {
		{
			p.SetState(86)
			p.Match(GsParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.Match(GsParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(93)
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
		p.SetState(95)
		p.Match(GsParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.Match(GsParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)
		p.Match(GsParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GsParserID {
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

		for _la == GsParserCOMMA {
			{
				p.SetState(99)
				p.Match(GsParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(100)
				p.Match(GsParserID)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(108)
		p.Match(GsParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(109)
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
		p.SetState(111)
		p.Match(GsParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2392537295044736) != 0) || _la == GsParserSEMICOLON || _la == GsParserID {
		{
			p.SetState(112)
			p.Statement()
		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			p.SetState(114)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(113)
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

				p.SetState(116)
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
				p.SetState(118)
				p.Match(GsParserEOF)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(126)
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

type BuiltinStmtContext struct {
	StatementContext
}

func NewBuiltinStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BuiltinStmtContext {
	var p = new(BuiltinStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *BuiltinStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinStmtContext) BuiltinCall() IBuiltinCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuiltinCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuiltinCallContext)
}

func (s *BuiltinStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitBuiltinStmt(s)

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

	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEmptyStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
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
			p.SetState(129)
			p.StructDefinition()
		}

	case 3:
		localctx = NewAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(130)
			p.Assign()
		}

	case 4:
		localctx = NewSelfOpAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(131)
			p.SelfAssign()
		}

	case 5:
		localctx = NewIncrDecrStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(132)
			p.IncrDecr()
		}

	case 6:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(133)
			p.Match(GsParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(142)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
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

			for _la == GsParserCOMMA {
				{
					p.SetState(135)
					p.Match(GsParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(136)
					p.Expr()
				}

				p.SetState(141)
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
		localctx = NewIfStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(144)
			p.Match(GsParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(145)
				p.Assign()
			}
			{
				p.SetState(146)
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
			p.SetState(150)
			p.Expr()
		}
		{
			p.SetState(151)
			p.Block()
		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GsParserELSE {
			{
				p.SetState(152)
				p.Match(GsParserELSE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(153)
				p.Block()
			}

		}

	case 8:
		localctx = NewForCStyleStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(156)
			p.Match(GsParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GsParserENV || _la == GsParserMUL || _la == GsParserID {
			{
				p.SetState(157)
				p.ForInit()
			}

		}
		{
			p.SetState(160)
			p.Match(GsParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3518437179533184) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&69131) != 0) {
			{
				p.SetState(161)
				p.Expr()
			}

		}
		{
			p.SetState(164)
			p.Match(GsParserSEMICOLON)
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

		if _la == GsParserENV || _la == GsParserMUL || _la == GsParserID {
			{
				p.SetState(165)
				p.ForUpdate()
			}

		}
		{
			p.SetState(168)
			p.Block()
		}

	case 9:
		localctx = NewForRangeStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(169)
			p.Match(GsParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.IterVar()
		}
		{
			p.SetState(171)
			p.Match(GsParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.Match(GsParserRANGE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Expr()
		}
		{
			p.SetState(174)
			p.Block()
		}

	case 10:
		localctx = NewForCondStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(176)
			p.Match(GsParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Expr()
		}
		{
			p.SetState(178)
			p.Block()
		}

	case 11:
		localctx = NewBuiltinStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(180)
			p.BuiltinCall()
		}

	case 12:
		localctx = NewCallStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(181)
			p.Call()
		}

	case 13:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(182)
			p.Match(GsParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		localctx = NewContinueStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(183)
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
	AllLvalue() []ILvalueContext
	Lvalue(i int) ILvalueContext
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

func (s *AssignContext) AllLvalue() []ILvalueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILvalueContext); ok {
			len++
		}
	}

	tst := make([]ILvalueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILvalueContext); ok {
			tst[i] = t.(ILvalueContext)
			i++
		}
	}

	return tst
}

func (s *AssignContext) Lvalue(i int) ILvalueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
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

	return t.(ILvalueContext)
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
		p.SetState(186)
		p.Lvalue()
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GsParserCOMMA {
		{
			p.SetState(187)
			p.Match(GsParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Lvalue()
		}

		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(194)
		p.Match(GsParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
		p.Expr()
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GsParserCOMMA {
		{
			p.SetState(196)
			p.Match(GsParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Expr()
		}

		p.SetState(202)
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
	Lvalue() ILvalueContext
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

func (s *IncrDecrContext) Lvalue() ILvalueContext {
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
		p.SetState(203)
		p.Lvalue()
	}
	{
		p.SetState(204)
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

// IBuiltinCallContext is an interface to support dynamic dispatch.
type IBuiltinCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBuiltinCallContext differentiates from other interfaces.
	IsBuiltinCallContext()
}

type BuiltinCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltinCallContext() *BuiltinCallContext {
	var p = new(BuiltinCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_builtinCall
	return p
}

func InitEmptyBuiltinCallContext(p *BuiltinCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_builtinCall
}

func (*BuiltinCallContext) IsBuiltinCallContext() {}

func NewBuiltinCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltinCallContext {
	var p = new(BuiltinCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_builtinCall

	return p
}

func (s *BuiltinCallContext) GetParser() antlr.Parser { return s.parser }

func (s *BuiltinCallContext) CopyAll(ctx *BuiltinCallContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BuiltinCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintfCallContext struct {
	BuiltinCallContext
}

func NewPrintfCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintfCallContext {
	var p = new(PrintfCallContext)

	InitEmptyBuiltinCallContext(&p.BuiltinCallContext)
	p.parser = parser
	p.CopyAll(ctx.(*BuiltinCallContext))

	return p
}

func (s *PrintfCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintfCallContext) PRINTF() antlr.TerminalNode {
	return s.GetToken(GsParserPRINTF, 0)
}

func (s *PrintfCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *PrintfCallContext) AllExpr() []IExprContext {
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

func (s *PrintfCallContext) Expr(i int) IExprContext {
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

func (s *PrintfCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *PrintfCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *PrintfCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}

func (s *PrintfCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitPrintfCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type ToStringCallContext struct {
	BuiltinCallContext
}

func NewToStringCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToStringCallContext {
	var p = new(ToStringCallContext)

	InitEmptyBuiltinCallContext(&p.BuiltinCallContext)
	p.parser = parser
	p.CopyAll(ctx.(*BuiltinCallContext))

	return p
}

func (s *ToStringCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToStringCallContext) TOSTRING() antlr.TerminalNode {
	return s.GetToken(GsParserTOSTRING, 0)
}

func (s *ToStringCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *ToStringCallContext) Expr() IExprContext {
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

func (s *ToStringCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *ToStringCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitToStringCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintCallContext struct {
	BuiltinCallContext
}

func NewPrintCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintCallContext {
	var p = new(PrintCallContext)

	InitEmptyBuiltinCallContext(&p.BuiltinCallContext)
	p.parser = parser
	p.CopyAll(ctx.(*BuiltinCallContext))

	return p
}

func (s *PrintCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintCallContext) PRINT() antlr.TerminalNode {
	return s.GetToken(GsParserPRINT, 0)
}

func (s *PrintCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *PrintCallContext) AllExpr() []IExprContext {
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

func (s *PrintCallContext) Expr(i int) IExprContext {
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

func (s *PrintCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *PrintCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *PrintCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}

func (s *PrintCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitPrintCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConvertCallContext struct {
	BuiltinCallContext
}

func NewConvertCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConvertCallContext {
	var p = new(ConvertCallContext)

	InitEmptyBuiltinCallContext(&p.BuiltinCallContext)
	p.parser = parser
	p.CopyAll(ctx.(*BuiltinCallContext))

	return p
}

func (s *ConvertCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConvertCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *ConvertCallContext) Expr() IExprContext {
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

func (s *ConvertCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *ConvertCallContext) UINT() antlr.TerminalNode {
	return s.GetToken(GsParserUINT, 0)
}

func (s *ConvertCallContext) UINT8() antlr.TerminalNode {
	return s.GetToken(GsParserUINT8, 0)
}

func (s *ConvertCallContext) UINT16() antlr.TerminalNode {
	return s.GetToken(GsParserUINT16, 0)
}

func (s *ConvertCallContext) UINT32() antlr.TerminalNode {
	return s.GetToken(GsParserUINT32, 0)
}

func (s *ConvertCallContext) UINT64() antlr.TerminalNode {
	return s.GetToken(GsParserUINT64, 0)
}

func (s *ConvertCallContext) INTS() antlr.TerminalNode {
	return s.GetToken(GsParserINTS, 0)
}

func (s *ConvertCallContext) INT8() antlr.TerminalNode {
	return s.GetToken(GsParserINT8, 0)
}

func (s *ConvertCallContext) INT16() antlr.TerminalNode {
	return s.GetToken(GsParserINT16, 0)
}

func (s *ConvertCallContext) INT32() antlr.TerminalNode {
	return s.GetToken(GsParserINT32, 0)
}

func (s *ConvertCallContext) INT64() antlr.TerminalNode {
	return s.GetToken(GsParserINT64, 0)
}

func (s *ConvertCallContext) FLOAT32() antlr.TerminalNode {
	return s.GetToken(GsParserFLOAT32, 0)
}

func (s *ConvertCallContext) FLOAT64() antlr.TerminalNode {
	return s.GetToken(GsParserFLOAT64, 0)
}

func (s *ConvertCallContext) STRINGS() antlr.TerminalNode {
	return s.GetToken(GsParserSTRINGS, 0)
}

func (s *ConvertCallContext) BOOL() antlr.TerminalNode {
	return s.GetToken(GsParserBOOL, 0)
}

func (s *ConvertCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitConvertCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type LenCallContext struct {
	BuiltinCallContext
}

func NewLenCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LenCallContext {
	var p = new(LenCallContext)

	InitEmptyBuiltinCallContext(&p.BuiltinCallContext)
	p.parser = parser
	p.CopyAll(ctx.(*BuiltinCallContext))

	return p
}

func (s *LenCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LenCallContext) LEN() antlr.TerminalNode {
	return s.GetToken(GsParserLEN, 0)
}

func (s *LenCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *LenCallContext) Expr() IExprContext {
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

func (s *LenCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *LenCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitLenCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type AppendCallContext struct {
	BuiltinCallContext
}

func NewAppendCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppendCallContext {
	var p = new(AppendCallContext)

	InitEmptyBuiltinCallContext(&p.BuiltinCallContext)
	p.parser = parser
	p.CopyAll(ctx.(*BuiltinCallContext))

	return p
}

func (s *AppendCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendCallContext) APPEND() antlr.TerminalNode {
	return s.GetToken(GsParserAPPEND, 0)
}

func (s *AppendCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *AppendCallContext) AllExpr() []IExprContext {
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

func (s *AppendCallContext) Expr(i int) IExprContext {
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

func (s *AppendCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *AppendCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}

func (s *AppendCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *AppendCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitAppendCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeleteCallContext struct {
	BuiltinCallContext
}

func NewDeleteCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeleteCallContext {
	var p = new(DeleteCallContext)

	InitEmptyBuiltinCallContext(&p.BuiltinCallContext)
	p.parser = parser
	p.CopyAll(ctx.(*BuiltinCallContext))

	return p
}

func (s *DeleteCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteCallContext) DELETE() antlr.TerminalNode {
	return s.GetToken(GsParserDELETE, 0)
}

func (s *DeleteCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *DeleteCallContext) AllExpr() []IExprContext {
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

func (s *DeleteCallContext) Expr(i int) IExprContext {
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

func (s *DeleteCallContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, 0)
}

func (s *DeleteCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *DeleteCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitDeleteCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type CopyCallContext struct {
	BuiltinCallContext
}

func NewCopyCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CopyCallContext {
	var p = new(CopyCallContext)

	InitEmptyBuiltinCallContext(&p.BuiltinCallContext)
	p.parser = parser
	p.CopyAll(ctx.(*BuiltinCallContext))

	return p
}

func (s *CopyCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyCallContext) COPY() antlr.TerminalNode {
	return s.GetToken(GsParserCOPY, 0)
}

func (s *CopyCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *CopyCallContext) Expr() IExprContext {
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

func (s *CopyCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *CopyCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitCopyCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintlnCallContext struct {
	BuiltinCallContext
}

func NewPrintlnCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintlnCallContext {
	var p = new(PrintlnCallContext)

	InitEmptyBuiltinCallContext(&p.BuiltinCallContext)
	p.parser = parser
	p.CopyAll(ctx.(*BuiltinCallContext))

	return p
}

func (s *PrintlnCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnCallContext) PRINTLN() antlr.TerminalNode {
	return s.GetToken(GsParserPRINTLN, 0)
}

func (s *PrintlnCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *PrintlnCallContext) AllExpr() []IExprContext {
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

func (s *PrintlnCallContext) Expr(i int) IExprContext {
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

func (s *PrintlnCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *PrintlnCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *PrintlnCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}

func (s *PrintlnCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitPrintlnCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) BuiltinCall() (localctx IBuiltinCallContext) {
	localctx = NewBuiltinCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GsParserRULE_builtinCall)
	var _la int

	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GsParserLEN:
		localctx = NewLenCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(206)
			p.Match(GsParserLEN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.Match(GsParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Expr()
		}
		{
			p.SetState(209)
			p.Match(GsParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GsParserAPPEND:
		localctx = NewAppendCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)
			p.Match(GsParserAPPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(212)
			p.Match(GsParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Expr()
		}
		{
			p.SetState(214)
			p.Match(GsParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.Expr()
		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GsParserCOMMA {
			{
				p.SetState(216)
				p.Match(GsParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(217)
				p.Expr()
			}

			p.SetState(222)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(223)
			p.Match(GsParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GsParserDELETE:
		localctx = NewDeleteCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(225)
			p.Match(GsParserDELETE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.Match(GsParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.Expr()
		}
		{
			p.SetState(228)
			p.Match(GsParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Expr()
		}
		{
			p.SetState(230)
			p.Match(GsParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GsParserCOPY:
		localctx = NewCopyCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(232)
			p.Match(GsParserCOPY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.Match(GsParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Expr()
		}
		{
			p.SetState(235)
			p.Match(GsParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GsParserTOSTRING:
		localctx = NewToStringCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(237)
			p.Match(GsParserTOSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Match(GsParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(239)
			p.Expr()
		}
		{
			p.SetState(240)
			p.Match(GsParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GsParserPRINT:
		localctx = NewPrintCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(242)
			p.Match(GsParserPRINT)
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
		{
			p.SetState(244)
			p.Expr()
		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GsParserCOMMA {
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

			p.SetState(251)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(252)
			p.Match(GsParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GsParserPRINTF:
		localctx = NewPrintfCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(254)
			p.Match(GsParserPRINTF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Match(GsParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)
			p.Expr()
		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GsParserCOMMA {
			{
				p.SetState(257)
				p.Match(GsParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(258)
				p.Expr()
			}

			p.SetState(263)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(264)
			p.Match(GsParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GsParserPRINTLN:
		localctx = NewPrintlnCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(266)
			p.Match(GsParserPRINTLN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.Match(GsParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.Expr()
		}
		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == GsParserCOMMA {
			{
				p.SetState(269)
				p.Match(GsParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(270)
				p.Expr()
			}

			p.SetState(275)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(276)
			p.Match(GsParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GsParserUINT, GsParserUINT8, GsParserUINT16, GsParserUINT32, GsParserUINT64, GsParserINTS, GsParserINT8, GsParserINT16, GsParserINT32, GsParserINT64, GsParserFLOAT32, GsParserFLOAT64, GsParserSTRINGS, GsParserBOOL:
		localctx = NewConvertCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(278)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&140728898420736) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(279)
			p.Match(GsParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)
			p.Expr()
		}
		{
			p.SetState(281)
			p.Match(GsParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
	p.EnterRule(localctx, 16, GsParserRULE_iterVar)
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingleIterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
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
			p.SetState(286)
			p.Match(GsParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(287)
			p.Match(GsParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(288)
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
	p.EnterRule(localctx, 18, GsParserRULE_forInit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
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
	p.EnterRule(localctx, 20, GsParserRULE_forUpdate)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
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
	Lvalue() ILvalueContext
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

func (s *SelfAssignContext) Lvalue() ILvalueContext {
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
	p.EnterRule(localctx, 22, GsParserRULE_selfAssign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.Lvalue()
	}
	{
		p.SetState(296)
		p.SelfAssignOp()
	}
	{
		p.SetState(297)
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
	p.EnterRule(localctx, 24, GsParserRULE_updateItem)
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSelfUpdateContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(299)
			p.SelfAssign()
		}

	case 2:
		localctx = NewIncrDecrUpdateContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(300)
			p.IncrDecr()
		}

	case 3:
		localctx = NewAssignUpdateContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(301)
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
	p.EnterRule(localctx, 26, GsParserRULE_selfAssignOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&124) != 0) {
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
	p.EnterRule(localctx, 28, GsParserRULE_call)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(306)
		p.Match(GsParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(307)
		p.Match(GsParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3518437179533184) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&69131) != 0) {
		{
			p.SetState(308)
			p.Expr()
		}
		p.SetState(313)
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
					p.SetState(309)
					p.Match(GsParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(310)
					p.Expr()
				}

			}
			p.SetState(315)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GsParserCOMMA {
			{
				p.SetState(316)
				p.Match(GsParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(321)
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
	p.EnterRule(localctx, 30, GsParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(323)
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
	p.EnterRule(localctx, 32, GsParserRULE_logicalOrExpr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
		p.LogicalAndExpr()
	}
	p.SetState(330)
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
				p.SetState(326)
				p.Match(GsParserOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(327)
				p.LogicalAndExpr()
			}

		}
		p.SetState(332)
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
	p.EnterRule(localctx, 34, GsParserRULE_logicalAndExpr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(333)
		p.ComparisonExpr()
	}
	p.SetState(338)
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
				p.SetState(334)
				p.Match(GsParserAND)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(335)
				p.ComparisonExpr()
			}

		}
		p.SetState(340)
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
	p.EnterRule(localctx, 36, GsParserRULE_comparisonExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		p.AddExpr()
	}
	p.SetState(345)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(342)
			p.CompOp()
		}
		{
			p.SetState(343)
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
	p.EnterRule(localctx, 38, GsParserRULE_addExpr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.BinExpr()
	}
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(348)
				p.AddOp()
			}
			{
				p.SetState(349)
				p.BinExpr()
			}

		}
		p.SetState(355)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 40, GsParserRULE_binExpr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(356)
		p.MulExpr()
	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(357)
				p.BitOp()
			}
			{
				p.SetState(358)
				p.MulExpr()
			}

		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 42, GsParserRULE_mulExpr)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.Atom()
	}
	p.SetState(371)
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
				p.SetState(366)
				p.MulOp()
			}
			{
				p.SetState(367)
				p.Atom()
			}

		}
		p.SetState(373)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
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

type BuiltinAtomContext struct {
	AtomContext
}

func NewBuiltinAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BuiltinAtomContext {
	var p = new(BuiltinAtomContext)

	InitEmptyAtomContext(&p.AtomContext)
	p.parser = parser
	p.CopyAll(ctx.(*AtomContext))

	return p
}

func (s *BuiltinAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinAtomContext) BuiltinCall() IBuiltinCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuiltinCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuiltinCallContext)
}

func (s *BuiltinAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitBuiltinAtom(s)

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
	p.EnterRule(localctx, 44, GsParserRULE_atom)
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNegAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(374)
			p.Match(GsParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)
			p.Atom()
		}

	case 2:
		localctx = NewIntAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(376)
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
			p.SetState(377)
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
			p.SetState(378)
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
			p.SetState(379)
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
			p.SetState(380)
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
			p.SetState(381)
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
			p.SetState(382)
			p.Match(GsParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(383)
			p.Expr()
		}

	case 9:
		localctx = NewQidAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(384)
			p.Qid()
		}

	case 10:
		localctx = NewBuiltinAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(385)
			p.BuiltinCall()
		}

	case 11:
		localctx = NewCallAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(386)
			p.Call()
		}

	case 12:
		localctx = NewInstanceAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(387)
			p.Instance()
		}

	case 13:
		localctx = NewArrayAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(388)
			p.ArrayLiteral()
		}

	case 14:
		localctx = NewDictAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(389)
			p.DictLiteral()
		}

	case 15:
		localctx = NewParenAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(390)
			p.Match(GsParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)
			p.Expr()
		}
		{
			p.SetState(392)
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
			p.SetState(394)
			p.Match(GsParserMUL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(395)
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
	MUL() antlr.TerminalNode
	Lvalue() ILvalueContext

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
	p.EnterRule(localctx, 46, GsParserRULE_lvalue)
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GsParserENV, GsParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(398)
			p.Qid()
		}

	case GsParserMUL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(399)
			p.Match(GsParserMUL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(400)
			p.Lvalue()
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
	p.EnterRule(localctx, 48, GsParserRULE_arrayLiteral)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(403)
		p.Match(GsParserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(415)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3518437179533184) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&69131) != 0) {
		{
			p.SetState(404)
			p.Expr()
		}
		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(405)
					p.Match(GsParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(406)
					p.Expr()
				}

			}
			p.SetState(411)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(413)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GsParserCOMMA {
			{
				p.SetState(412)
				p.Match(GsParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(417)
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
	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3518437179533184) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&69131) != 0) {
		{
			p.SetState(419)
			p.Expr()
		}

	}
	{
		p.SetState(422)
		p.Match(GsParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3518437179533184) != 0) || ((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&69131) != 0) {
		{
			p.SetState(423)
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
		p.SetState(426)
		p.Match(GsParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2251799813686144) != 0) || ((int64((_la-75)) & ^0x3f) == 0 && ((int64(1)<<(_la-75))&135) != 0) {
		{
			p.SetState(427)
			p.DictEntry()
		}
		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(428)
					p.Match(GsParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(429)
					p.DictEntry()
				}

			}
			p.SetState(434)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(436)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GsParserCOMMA {
			{
				p.SetState(435)
				p.Match(GsParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(440)
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

func (s *IdKeyEntryContext) Lvalue() ILvalueContext {
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

	p.SetState(449)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GsParserTRUE, GsParserFALSE, GsParserINT, GsParserFLOAT, GsParserSTRING:
		localctx = NewConstKeyEntryContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(442)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GsParserTRUE || _la == GsParserFALSE || ((int64((_la-75)) & ^0x3f) == 0 && ((int64(1)<<(_la-75))&7) != 0)) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(443)
			p.Match(GsParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(444)
			p.Expr()
		}

	case GsParserENV, GsParserMUL, GsParserID:
		localctx = NewIdKeyEntryContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(445)
			p.Lvalue()
		}
		{
			p.SetState(446)
			p.Match(GsParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(447)
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
		p.SetState(451)
		p.Match(GsParserNEW)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(452)
		p.Match(GsParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(453)
		p.Match(GsParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(469)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GsParserID {
		{
			p.SetState(454)
			p.Match(GsParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(455)
			p.Match(GsParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(456)
			p.Expr()
		}
		p.SetState(463)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(457)
					p.Match(GsParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(458)
					p.Match(GsParserID)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(459)
					p.Match(GsParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(460)
					p.Expr()
				}

			}
			p.SetState(465)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(467)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == GsParserCOMMA {
			{
				p.SetState(466)
				p.Match(GsParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	}
	{
		p.SetState(471)
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
	AllAccessor() []IAccessorContext
	Accessor(i int) IAccessorContext

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

func (s *QidContext) AllAccessor() []IAccessorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAccessorContext); ok {
			len++
		}
	}

	tst := make([]IAccessorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAccessorContext); ok {
			tst[i] = t.(IAccessorContext)
			i++
		}
	}

	return tst
}

func (s *QidContext) Accessor(i int) IAccessorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessorContext); ok {
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

	return t.(IAccessorContext)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(473)
		p.Primary()
	}
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-47)) & ^0x3f) == 0 && ((int64(1)<<(_la-47))&786435) != 0 {
		{
			p.SetState(474)
			p.Accessor()
		}

		p.SetState(479)
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

// IAccessorContext is an interface to support dynamic dispatch.
type IAccessorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAccessorContext differentiates from other interfaces.
	IsAccessorContext()
}

type AccessorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessorContext() *AccessorContext {
	var p = new(AccessorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_accessor
	return p
}

func InitEmptyAccessorContext(p *AccessorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_accessor
}

func (*AccessorContext) IsAccessorContext() {}

func NewAccessorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessorContext {
	var p = new(AccessorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_accessor

	return p
}

func (s *AccessorContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessorContext) CopyAll(ctx *AccessorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AccessorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IndexAccessContext struct {
	AccessorContext
}

func NewIndexAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexAccessContext {
	var p = new(IndexAccessContext)

	InitEmptyAccessorContext(&p.AccessorContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessorContext))

	return p
}

func (s *IndexAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexAccessContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(GsParserRBRACK, 0)
}

func (s *IndexAccessContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(GsParserLBRACK, 0)
}

func (s *IndexAccessContext) SAFE_LBRACK() antlr.TerminalNode {
	return s.GetToken(GsParserSAFE_LBRACK, 0)
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

func (s *IndexAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitIndexAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

type PropertyAccessContext struct {
	AccessorContext
}

func NewPropertyAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertyAccessContext {
	var p = new(PropertyAccessContext)

	InitEmptyAccessorContext(&p.AccessorContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessorContext))

	return p
}

func (s *PropertyAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyAccessContext) ID() antlr.TerminalNode {
	return s.GetToken(GsParserID, 0)
}

func (s *PropertyAccessContext) DOT() antlr.TerminalNode {
	return s.GetToken(GsParserDOT, 0)
}

func (s *PropertyAccessContext) SAFE_DOT() antlr.TerminalNode {
	return s.GetToken(GsParserSAFE_DOT, 0)
}

func (s *PropertyAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitPropertyAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GsParser) Accessor() (localctx IAccessorContext) {
	localctx = NewAccessorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GsParserRULE_accessor)
	var _la int

	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GsParserSAFE_DOT, GsParserDOT:
		localctx = NewPropertyAccessContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(480)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GsParserSAFE_DOT || _la == GsParserDOT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(481)
			p.Match(GsParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GsParserSAFE_LBRACK, GsParserLBRACK:
		localctx = NewIndexAccessContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(482)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GsParserSAFE_LBRACK || _la == GsParserLBRACK) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(485)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(483)
				p.Expr()
			}

		case 2:
			{
				p.SetState(484)
				p.SliceExpr()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(487)
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
	p.EnterRule(localctx, 62, GsParserRULE_primary)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(491)
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
	p.EnterRule(localctx, 64, GsParserRULE_compOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(493)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1134907106097364992) != 0) {
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
	p.EnterRule(localctx, 66, GsParserRULE_addOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(495)
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
	p.EnterRule(localctx, 68, GsParserRULE_bitOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(497)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8070450532247928832) != 0) {
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
	p.EnterRule(localctx, 70, GsParserRULE_mulOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(499)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15762598695796736) != 0) {
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
