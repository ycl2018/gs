// Code generated from github.com/ycl2018/gs/Gs.g4 by ANTLR 4.13.2. DO NOT EDIT.

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
    "'continue'", "'global'", "'go'", "'len'", "'append'", "'delete'", "'copy'", 
    "'toString'", "'print'", "'printf'", "'sprintf'", "'println'", "'uint8'", 
    "'uint16'", "'uint32'", "'uint64'", "'uint'", "'int8'", "'int16'", "'int32'", 
    "'int64'", "'int'", "'float32'", "'float64'", "'string'", "'bool'", 
    "'...'", "'<<'", "'>>'", "'++'", "'--'", "'>='", "'<='", "'!='", "'=='", 
    "'+'", "'-'", "'*'", "'/'", "'%'", "'<'", "'>'", "'&'", "'|'", "'^'", 
    "'.'", "'['", "'('", "')'", "'{'", "'}'", "']'", "':'", "';'", "','",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "ENV", "TRUE", "FALSE", "NIL", "AND", "OR", 
    "NOT", "IF", "ELSE", "FOR", "RANGE", "RETURN", "FUNC", "TYPE", "STRUCT", 
    "NEW", "BREAK", "CONTINUE", "GLOBAL", "GO", "LEN", "APPEND", "DELETE", 
    "COPY", "TOSTRING", "PRINT", "PRINTF", "SPRINTF", "PRINTLN", "UINT8", 
    "UINT16", "UINT32", "UINT64", "UINT", "INT8", "INT16", "INT32", "INT64", 
    "INTS", "FLOAT32", "FLOAT64", "STRINGS", "BOOL", "EXPAND", "LSHIFT", 
    "RSHIFT", "INCR", "DECR", "GEQ", "LEQ", "NEQ", "EQ", "ADD", "SUB", "MUL", 
    "DIV", "MOD", "LT", "GT", "BITAND", "BITOR", "XOR", "DOT", "LBRACK", 
    "LPAREN", "RPAREN", "LBRACE", "RBRACE", "RBRACK", "COLON", "SEMICOLON", 
    "COMMA", "INT", "FLOAT", "STRING", "WS", "NEWLINE", "SL_COMMENT", "ML_COMMENT", 
    "ID",
  }
  staticData.RuleNames = []string{
    "program", "structDefinition", "functionDefinition", "block", "statement", 
    "assign", "incrDecr", "builtinCall", "iterVar", "forInit", "forUpdate", 
    "selfAssign", "updateItem", "selfAssignOp", "call", "expr", "atom", 
    "lvalue", "arrayLiteral", "sliceExpr", "dictLiteral", "dictEntry", "instance", 
    "qid", "accessor", "primary", "compOp", "addOp", "mulOp",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 86, 488, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2, 
	21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26, 
	7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 1, 0, 1, 0, 4, 0, 62, 8, 0, 11, 
	0, 12, 0, 63, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 
	1, 75, 8, 1, 10, 1, 12, 1, 78, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 
	1, 2, 1, 2, 5, 2, 88, 8, 2, 10, 2, 12, 2, 91, 9, 2, 3, 2, 93, 8, 2, 1, 
	2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 4, 3, 101, 8, 3, 11, 3, 12, 3, 102, 1, 
	3, 3, 3, 106, 8, 3, 5, 3, 108, 8, 3, 10, 3, 12, 3, 111, 9, 3, 1, 3, 1, 
	3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 124, 8, 
	4, 10, 4, 12, 4, 127, 9, 4, 3, 4, 129, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 
	4, 135, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 141, 8, 4, 1, 4, 1, 4, 3, 4, 
	145, 8, 4, 1, 4, 1, 4, 3, 4, 149, 8, 4, 1, 4, 1, 4, 3, 4, 153, 8, 4, 1, 
	4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 
	4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 175, 8, 4, 10, 4, 12, 
	4, 178, 9, 4, 3, 4, 180, 8, 4, 1, 5, 1, 5, 1, 5, 5, 5, 185, 8, 5, 10, 5, 
	12, 5, 188, 9, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 194, 8, 5, 10, 5, 12, 5, 
	197, 9, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 
	1, 7, 1, 7, 1, 7, 5, 7, 212, 8, 7, 10, 7, 12, 7, 215, 9, 7, 1, 7, 1, 7, 
	1, 7, 1, 7, 3, 7, 221, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 
	1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 
	5, 7, 242, 8, 7, 10, 7, 12, 7, 245, 9, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 
	1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 259, 8, 7, 1, 8, 1, 8, 
	1, 8, 1, 8, 3, 8, 265, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 
	11, 1, 11, 1, 12, 1, 12, 1, 12, 3, 12, 278, 8, 12, 1, 13, 1, 13, 1, 14, 
	1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 287, 8, 14, 10, 14, 12, 14, 290, 9, 
	14, 1, 14, 3, 14, 293, 8, 14, 3, 14, 295, 8, 14, 1, 14, 1, 14, 1, 14, 4, 
	14, 300, 8, 14, 11, 14, 12, 14, 301, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 
	308, 8, 14, 10, 14, 12, 14, 311, 9, 14, 1, 14, 3, 14, 314, 8, 14, 3, 14, 
	316, 8, 14, 1, 14, 1, 14, 3, 14, 320, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 
	1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 
	15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 343, 8, 15, 10, 15, 
	12, 15, 346, 9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 
	16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 
	1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 370, 8, 16, 1, 17, 1, 17, 1, 17, 3, 
	17, 375, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 381, 8, 18, 10, 18, 
	12, 18, 384, 9, 18, 1, 18, 3, 18, 387, 8, 18, 3, 18, 389, 8, 18, 1, 18, 
	1, 18, 1, 19, 3, 19, 394, 8, 19, 1, 19, 1, 19, 3, 19, 398, 8, 19, 1, 20, 
	1, 20, 1, 20, 1, 20, 5, 20, 404, 8, 20, 10, 20, 12, 20, 407, 9, 20, 1, 
	20, 3, 20, 410, 8, 20, 3, 20, 412, 8, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 
	21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 423, 8, 21, 1, 22, 1, 22, 1, 22, 
	1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 435, 8, 22, 10, 
	22, 12, 22, 438, 9, 22, 1, 22, 3, 22, 441, 8, 22, 3, 22, 443, 8, 22, 1, 
	22, 1, 22, 1, 23, 1, 23, 5, 23, 449, 8, 23, 10, 23, 12, 23, 452, 9, 23, 
	1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 459, 8, 24, 1, 24, 1, 24, 1, 
	24, 1, 24, 1, 24, 1, 24, 5, 24, 467, 8, 24, 10, 24, 12, 24, 470, 9, 24, 
	1, 24, 3, 24, 473, 8, 24, 3, 24, 475, 8, 24, 1, 24, 3, 24, 478, 8, 24, 
	1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 0, 1, 30, 
	29, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 0, 10, 2, 0, 77, 77, 83, 83, 
	1, 0, 53, 54, 1, 0, 32, 35, 1, 0, 36, 49, 1, 0, 2, 6, 2, 0, 8, 9, 79, 81, 
	2, 0, 7, 7, 86, 86, 2, 0, 55, 58, 64, 65, 2, 0, 59, 60, 67, 68, 3, 0, 51, 
	52, 61, 63, 66, 66, 552, 0, 61, 1, 0, 0, 0, 2, 67, 1, 0, 0, 0, 4, 81, 1, 
	0, 0, 0, 6, 97, 1, 0, 0, 0, 8, 179, 1, 0, 0, 0, 10, 181, 1, 0, 0, 0, 12, 
	198, 1, 0, 0, 0, 14, 258, 1, 0, 0, 0, 16, 264, 1, 0, 0, 0, 18, 266, 1, 
	0, 0, 0, 20, 268, 1, 0, 0, 0, 22, 270, 1, 0, 0, 0, 24, 277, 1, 0, 0, 0, 
	26, 279, 1, 0, 0, 0, 28, 319, 1, 0, 0, 0, 30, 321, 1, 0, 0, 0, 32, 369, 
	1, 0, 0, 0, 34, 374, 1, 0, 0, 0, 36, 376, 1, 0, 0, 0, 38, 393, 1, 0, 0, 
	0, 40, 399, 1, 0, 0, 0, 42, 422, 1, 0, 0, 0, 44, 424, 1, 0, 0, 0, 46, 446, 
	1, 0, 0, 0, 48, 477, 1, 0, 0, 0, 50, 479, 1, 0, 0, 0, 52, 481, 1, 0, 0, 
	0, 54, 483, 1, 0, 0, 0, 56, 485, 1, 0, 0, 0, 58, 62, 3, 4, 2, 0, 59, 62, 
	3, 2, 1, 0, 60, 62, 3, 8, 4, 0, 61, 58, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 
	61, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 64, 1, 
	0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 66, 5, 0, 0, 1, 66, 1, 1, 0, 0, 0, 67, 
	68, 5, 20, 0, 0, 68, 69, 5, 86, 0, 0, 69, 70, 5, 21, 0, 0, 70, 71, 5, 73, 
	0, 0, 71, 76, 5, 86, 0, 0, 72, 73, 5, 78, 0, 0, 73, 75, 5, 86, 0, 0, 74, 
	72, 1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0, 0, 76, 77, 1, 0, 0, 
	0, 77, 79, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 80, 5, 74, 0, 0, 80, 3, 
	1, 0, 0, 0, 81, 82, 5, 19, 0, 0, 82, 83, 5, 86, 0, 0, 83, 92, 5, 71, 0, 
	0, 84, 89, 5, 86, 0, 0, 85, 86, 5, 78, 0, 0, 86, 88, 5, 86, 0, 0, 87, 85, 
	1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 
	90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 84, 1, 0, 0, 0, 92, 93, 1, 
	0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95, 5, 72, 0, 0, 95, 96, 3, 6, 3, 0, 96, 
	5, 1, 0, 0, 0, 97, 109, 5, 73, 0, 0, 98, 105, 3, 8, 4, 0, 99, 101, 7, 0, 
	0, 0, 100, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 
	102, 103, 1, 0, 0, 0, 103, 106, 1, 0, 0, 0, 104, 106, 5, 0, 0, 1, 105, 
	100, 1, 0, 0, 0, 105, 104, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 108, 
	1, 0, 0, 0, 107, 98, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0, 
	0, 0, 109, 110, 1, 0, 0, 0, 110, 112, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 
	112, 113, 5, 74, 0, 0, 113, 7, 1, 0, 0, 0, 114, 180, 5, 77, 0, 0, 115, 
	180, 3, 2, 1, 0, 116, 180, 3, 10, 5, 0, 117, 180, 3, 22, 11, 0, 118, 180, 
	3, 12, 6, 0, 119, 128, 5, 18, 0, 0, 120, 125, 3, 30, 15, 0, 121, 122, 5, 
	78, 0, 0, 122, 124, 3, 30, 15, 0, 123, 121, 1, 0, 0, 0, 124, 127, 1, 0, 
	0, 0, 125, 123, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 129, 1, 0, 0, 0, 
	127, 125, 1, 0, 0, 0, 128, 120, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 
	180, 1, 0, 0, 0, 130, 134, 5, 14, 0, 0, 131, 132, 3, 10, 5, 0, 132, 133, 
	5, 77, 0, 0, 133, 135, 1, 0, 0, 0, 134, 131, 1, 0, 0, 0, 134, 135, 1, 0, 
	0, 0, 135, 136, 1, 0, 0, 0, 136, 137, 3, 30, 15, 0, 137, 140, 3, 6, 3, 
	0, 138, 139, 5, 15, 0, 0, 139, 141, 3, 6, 3, 0, 140, 138, 1, 0, 0, 0, 140, 
	141, 1, 0, 0, 0, 141, 180, 1, 0, 0, 0, 142, 144, 5, 16, 0, 0, 143, 145, 
	3, 18, 9, 0, 144, 143, 1, 0, 0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 1, 0, 
	0, 0, 146, 148, 5, 77, 0, 0, 147, 149, 3, 30, 15, 0, 148, 147, 1, 0, 0, 
	0, 148, 149, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 152, 5, 77, 0, 0, 151, 
	153, 3, 20, 10, 0, 152, 151, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 
	1, 0, 0, 0, 154, 180, 3, 6, 3, 0, 155, 156, 5, 16, 0, 0, 156, 157, 3, 16, 
	8, 0, 157, 158, 5, 1, 0, 0, 158, 159, 5, 17, 0, 0, 159, 160, 3, 30, 15, 
	0, 160, 161, 3, 6, 3, 0, 161, 180, 1, 0, 0, 0, 162, 163, 5, 16, 0, 0, 163, 
	164, 3, 30, 15, 0, 164, 165, 3, 6, 3, 0, 165, 180, 1, 0, 0, 0, 166, 180, 
	3, 14, 7, 0, 167, 180, 3, 28, 14, 0, 168, 180, 5, 23, 0, 0, 169, 180, 5, 
	24, 0, 0, 170, 171, 5, 25, 0, 0, 171, 176, 5, 86, 0, 0, 172, 173, 5, 78, 
	0, 0, 173, 175, 5, 86, 0, 0, 174, 172, 1, 0, 0, 0, 175, 178, 1, 0, 0, 0, 
	176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 180, 1, 0, 0, 0, 178, 
	176, 1, 0, 0, 0, 179, 114, 1, 0, 0, 0, 179, 115, 1, 0, 0, 0, 179, 116, 
	1, 0, 0, 0, 179, 117, 1, 0, 0, 0, 179, 118, 1, 0, 0, 0, 179, 119, 1, 0, 
	0, 0, 179, 130, 1, 0, 0, 0, 179, 142, 1, 0, 0, 0, 179, 155, 1, 0, 0, 0, 
	179, 162, 1, 0, 0, 0, 179, 166, 1, 0, 0, 0, 179, 167, 1, 0, 0, 0, 179, 
	168, 1, 0, 0, 0, 179, 169, 1, 0, 0, 0, 179, 170, 1, 0, 0, 0, 180, 9, 1, 
	0, 0, 0, 181, 186, 3, 34, 17, 0, 182, 183, 5, 78, 0, 0, 183, 185, 3, 34, 
	17, 0, 184, 182, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 
	186, 187, 1, 0, 0, 0, 187, 189, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 189, 
	190, 5, 1, 0, 0, 190, 195, 3, 30, 15, 0, 191, 192, 5, 78, 0, 0, 192, 194, 
	3, 30, 15, 0, 193, 191, 1, 0, 0, 0, 194, 197, 1, 0, 0, 0, 195, 193, 1, 
	0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 11, 1, 0, 0, 0, 197, 195, 1, 0, 0, 
	0, 198, 199, 3, 34, 17, 0, 199, 200, 7, 1, 0, 0, 200, 13, 1, 0, 0, 0, 201, 
	202, 5, 27, 0, 0, 202, 203, 5, 71, 0, 0, 203, 204, 3, 30, 15, 0, 204, 205, 
	5, 72, 0, 0, 205, 259, 1, 0, 0, 0, 206, 207, 5, 28, 0, 0, 207, 208, 5, 
	71, 0, 0, 208, 220, 3, 30, 15, 0, 209, 210, 5, 78, 0, 0, 210, 212, 3, 30, 
	15, 0, 211, 209, 1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 
	213, 214, 1, 0, 0, 0, 214, 221, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 216, 
	217, 5, 78, 0, 0, 217, 218, 3, 30, 15, 0, 218, 219, 5, 50, 0, 0, 219, 221, 
	1, 0, 0, 0, 220, 213, 1, 0, 0, 0, 220, 216, 1, 0, 0, 0, 221, 222, 1, 0, 
	0, 0, 222, 223, 5, 72, 0, 0, 223, 259, 1, 0, 0, 0, 224, 225, 5, 29, 0, 
	0, 225, 226, 5, 71, 0, 0, 226, 227, 3, 30, 15, 0, 227, 228, 5, 78, 0, 0, 
	228, 229, 3, 30, 15, 0, 229, 230, 5, 72, 0, 0, 230, 259, 1, 0, 0, 0, 231, 
	232, 5, 31, 0, 0, 232, 233, 5, 71, 0, 0, 233, 234, 3, 30, 15, 0, 234, 235, 
	5, 72, 0, 0, 235, 259, 1, 0, 0, 0, 236, 237, 7, 2, 0, 0, 237, 238, 5, 71, 
	0, 0, 238, 243, 3, 30, 15, 0, 239, 240, 5, 78, 0, 0, 240, 242, 3, 30, 15, 
	0, 241, 239, 1, 0, 0, 0, 242, 245, 1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 243, 
	244, 1, 0, 0, 0, 244, 246, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 246, 247, 
	5, 72, 0, 0, 247, 259, 1, 0, 0, 0, 248, 249, 7, 3, 0, 0, 249, 250, 5, 71, 
	0, 0, 250, 251, 3, 30, 15, 0, 251, 252, 5, 72, 0, 0, 252, 259, 1, 0, 0, 
	0, 253, 254, 5, 26, 0, 0, 254, 255, 5, 71, 0, 0, 255, 256, 3, 30, 15, 0, 
	256, 257, 5, 72, 0, 0, 257, 259, 1, 0, 0, 0, 258, 201, 1, 0, 0, 0, 258, 
	206, 1, 0, 0, 0, 258, 224, 1, 0, 0, 0, 258, 231, 1, 0, 0, 0, 258, 236, 
	1, 0, 0, 0, 258, 248, 1, 0, 0, 0, 258, 253, 1, 0, 0, 0, 259, 15, 1, 0, 
	0, 0, 260, 265, 5, 86, 0, 0, 261, 262, 5, 86, 0, 0, 262, 263, 5, 78, 0, 
	0, 263, 265, 5, 86, 0, 0, 264, 260, 1, 0, 0, 0, 264, 261, 1, 0, 0, 0, 265, 
	17, 1, 0, 0, 0, 266, 267, 3, 10, 5, 0, 267, 19, 1, 0, 0, 0, 268, 269, 3, 
	24, 12, 0, 269, 21, 1, 0, 0, 0, 270, 271, 3, 34, 17, 0, 271, 272, 3, 26, 
	13, 0, 272, 273, 3, 30, 15, 0, 273, 23, 1, 0, 0, 0, 274, 278, 3, 22, 11, 
	0, 275, 278, 3, 12, 6, 0, 276, 278, 3, 10, 5, 0, 277, 274, 1, 0, 0, 0, 
	277, 275, 1, 0, 0, 0, 277, 276, 1, 0, 0, 0, 278, 25, 1, 0, 0, 0, 279, 280, 
	7, 4, 0, 0, 280, 27, 1, 0, 0, 0, 281, 282, 5, 86, 0, 0, 282, 294, 5, 71, 
	0, 0, 283, 288, 3, 30, 15, 0, 284, 285, 5, 78, 0, 0, 285, 287, 3, 30, 15, 
	0, 286, 284, 1, 0, 0, 0, 287, 290, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 288, 
	289, 1, 0, 0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0, 291, 293, 
	5, 78, 0, 0, 292, 291, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 295, 1, 0, 
	0, 0, 294, 283, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 
	296, 320, 5, 72, 0, 0, 297, 299, 3, 50, 25, 0, 298, 300, 3, 48, 24, 0, 
	299, 298, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 299, 1, 0, 0, 0, 301, 
	302, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 315, 5, 71, 0, 0, 304, 309, 
	3, 30, 15, 0, 305, 306, 5, 78, 0, 0, 306, 308, 3, 30, 15, 0, 307, 305, 
	1, 0, 0, 0, 308, 311, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 309, 310, 1, 0, 
	0, 0, 310, 313, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 312, 314, 5, 78, 0, 0, 
	313, 312, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 316, 1, 0, 0, 0, 315, 
	304, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 318, 
	5, 72, 0, 0, 318, 320, 1, 0, 0, 0, 319, 281, 1, 0, 0, 0, 319, 297, 1, 0, 
	0, 0, 320, 29, 1, 0, 0, 0, 321, 322, 6, 15, -1, 0, 322, 323, 3, 32, 16, 
	0, 323, 344, 1, 0, 0, 0, 324, 325, 10, 5, 0, 0, 325, 326, 3, 56, 28, 0, 
	326, 327, 3, 30, 15, 6, 327, 343, 1, 0, 0, 0, 328, 329, 10, 4, 0, 0, 329, 
	330, 3, 54, 27, 0, 330, 331, 3, 30, 15, 5, 331, 343, 1, 0, 0, 0, 332, 333, 
	10, 3, 0, 0, 333, 334, 3, 52, 26, 0, 334, 335, 3, 30, 15, 4, 335, 343, 
	1, 0, 0, 0, 336, 337, 10, 2, 0, 0, 337, 338, 5, 11, 0, 0, 338, 343, 3, 
	30, 15, 3, 339, 340, 10, 1, 0, 0, 340, 341, 5, 12, 0, 0, 341, 343, 3, 30, 
	15, 2, 342, 324, 1, 0, 0, 0, 342, 328, 1, 0, 0, 0, 342, 332, 1, 0, 0, 0, 
	342, 336, 1, 0, 0, 0, 342, 339, 1, 0, 0, 0, 343, 346, 1, 0, 0, 0, 344, 
	342, 1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 31, 1, 0, 0, 0, 346, 344, 1, 
	0, 0, 0, 347, 348, 5, 60, 0, 0, 348, 370, 3, 32, 16, 0, 349, 350, 5, 13, 
	0, 0, 350, 370, 3, 32, 16, 0, 351, 352, 5, 61, 0, 0, 352, 370, 3, 34, 17, 
	0, 353, 370, 5, 79, 0, 0, 354, 370, 5, 80, 0, 0, 355, 370, 5, 81, 0, 0, 
	356, 370, 5, 8, 0, 0, 357, 370, 5, 9, 0, 0, 358, 370, 5, 10, 0, 0, 359, 
	370, 3, 14, 7, 0, 360, 370, 3, 28, 14, 0, 361, 370, 3, 44, 22, 0, 362, 
	370, 3, 36, 18, 0, 363, 370, 3, 40, 20, 0, 364, 365, 5, 71, 0, 0, 365, 
	366, 3, 30, 15, 0, 366, 367, 5, 72, 0, 0, 367, 370, 1, 0, 0, 0, 368, 370, 
	3, 46, 23, 0, 369, 347, 1, 0, 0, 0, 369, 349, 1, 0, 0, 0, 369, 351, 1, 
	0, 0, 0, 369, 353, 1, 0, 0, 0, 369, 354, 1, 0, 0, 0, 369, 355, 1, 0, 0, 
	0, 369, 356, 1, 0, 0, 0, 369, 357, 1, 0, 0, 0, 369, 358, 1, 0, 0, 0, 369, 
	359, 1, 0, 0, 0, 369, 360, 1, 0, 0, 0, 369, 361, 1, 0, 0, 0, 369, 362, 
	1, 0, 0, 0, 369, 363, 1, 0, 0, 0, 369, 364, 1, 0, 0, 0, 369, 368, 1, 0, 
	0, 0, 370, 33, 1, 0, 0, 0, 371, 375, 3, 46, 23, 0, 372, 373, 5, 61, 0, 
	0, 373, 375, 3, 34, 17, 0, 374, 371, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0, 
	375, 35, 1, 0, 0, 0, 376, 388, 5, 70, 0, 0, 377, 382, 3, 30, 15, 0, 378, 
	379, 5, 78, 0, 0, 379, 381, 3, 30, 15, 0, 380, 378, 1, 0, 0, 0, 381, 384, 
	1, 0, 0, 0, 382, 380, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 386, 1, 0, 
	0, 0, 384, 382, 1, 0, 0, 0, 385, 387, 5, 78, 0, 0, 386, 385, 1, 0, 0, 0, 
	386, 387, 1, 0, 0, 0, 387, 389, 1, 0, 0, 0, 388, 377, 1, 0, 0, 0, 388, 
	389, 1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390, 391, 5, 75, 0, 0, 391, 37, 
	1, 0, 0, 0, 392, 394, 3, 30, 15, 0, 393, 392, 1, 0, 0, 0, 393, 394, 1, 
	0, 0, 0, 394, 395, 1, 0, 0, 0, 395, 397, 5, 76, 0, 0, 396, 398, 3, 30, 
	15, 0, 397, 396, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 39, 1, 0, 0, 0, 
	399, 411, 5, 73, 0, 0, 400, 405, 3, 42, 21, 0, 401, 402, 5, 78, 0, 0, 402, 
	404, 3, 42, 21, 0, 403, 401, 1, 0, 0, 0, 404, 407, 1, 0, 0, 0, 405, 403, 
	1, 0, 0, 0, 405, 406, 1, 0, 0, 0, 406, 409, 1, 0, 0, 0, 407, 405, 1, 0, 
	0, 0, 408, 410, 5, 78, 0, 0, 409, 408, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 
	410, 412, 1, 0, 0, 0, 411, 400, 1, 0, 0, 0, 411, 412, 1, 0, 0, 0, 412, 
	413, 1, 0, 0, 0, 413, 414, 5, 74, 0, 0, 414, 41, 1, 0, 0, 0, 415, 416, 
	7, 5, 0, 0, 416, 417, 5, 76, 0, 0, 417, 423, 3, 30, 15, 0, 418, 419, 3, 
	34, 17, 0, 419, 420, 5, 76, 0, 0, 420, 421, 3, 30, 15, 0, 421, 423, 1, 
	0, 0, 0, 422, 415, 1, 0, 0, 0, 422, 418, 1, 0, 0, 0, 423, 43, 1, 0, 0, 
	0, 424, 425, 5, 22, 0, 0, 425, 426, 5, 86, 0, 0, 426, 442, 5, 73, 0, 0, 
	427, 428, 5, 86, 0, 0, 428, 429, 5, 76, 0, 0, 429, 436, 3, 30, 15, 0, 430, 
	431, 5, 78, 0, 0, 431, 432, 5, 86, 0, 0, 432, 433, 5, 76, 0, 0, 433, 435, 
	3, 30, 15, 0, 434, 430, 1, 0, 0, 0, 435, 438, 1, 0, 0, 0, 436, 434, 1, 
	0, 0, 0, 436, 437, 1, 0, 0, 0, 437, 440, 1, 0, 0, 0, 438, 436, 1, 0, 0, 
	0, 439, 441, 5, 78, 0, 0, 440, 439, 1, 0, 0, 0, 440, 441, 1, 0, 0, 0, 441, 
	443, 1, 0, 0, 0, 442, 427, 1, 0, 0, 0, 442, 443, 1, 0, 0, 0, 443, 444, 
	1, 0, 0, 0, 444, 445, 5, 74, 0, 0, 445, 45, 1, 0, 0, 0, 446, 450, 3, 50, 
	25, 0, 447, 449, 3, 48, 24, 0, 448, 447, 1, 0, 0, 0, 449, 452, 1, 0, 0, 
	0, 450, 448, 1, 0, 0, 0, 450, 451, 1, 0, 0, 0, 451, 47, 1, 0, 0, 0, 452, 
	450, 1, 0, 0, 0, 453, 454, 5, 69, 0, 0, 454, 478, 5, 86, 0, 0, 455, 458, 
	5, 70, 0, 0, 456, 459, 3, 30, 15, 0, 457, 459, 3, 38, 19, 0, 458, 456, 
	1, 0, 0, 0, 458, 457, 1, 0, 0, 0, 459, 460, 1, 0, 0, 0, 460, 461, 5, 75, 
	0, 0, 461, 478, 1, 0, 0, 0, 462, 474, 5, 71, 0, 0, 463, 468, 3, 30, 15, 
	0, 464, 465, 5, 78, 0, 0, 465, 467, 3, 30, 15, 0, 466, 464, 1, 0, 0, 0, 
	467, 470, 1, 0, 0, 0, 468, 466, 1, 0, 0, 0, 468, 469, 1, 0, 0, 0, 469, 
	472, 1, 0, 0, 0, 470, 468, 1, 0, 0, 0, 471, 473, 5, 78, 0, 0, 472, 471, 
	1, 0, 0, 0, 472, 473, 1, 0, 0, 0, 473, 475, 1, 0, 0, 0, 474, 463, 1, 0, 
	0, 0, 474, 475, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476, 478, 5, 72, 0, 0, 
	477, 453, 1, 0, 0, 0, 477, 455, 1, 0, 0, 0, 477, 462, 1, 0, 0, 0, 478, 
	49, 1, 0, 0, 0, 479, 480, 7, 6, 0, 0, 480, 51, 1, 0, 0, 0, 481, 482, 7, 
	7, 0, 0, 482, 53, 1, 0, 0, 0, 483, 484, 7, 8, 0, 0, 484, 55, 1, 0, 0, 0, 
	485, 486, 7, 9, 0, 0, 486, 57, 1, 0, 0, 0, 55, 61, 63, 76, 89, 92, 102, 
	105, 109, 125, 128, 134, 140, 144, 148, 152, 176, 179, 186, 195, 213, 220, 
	243, 258, 264, 277, 288, 292, 294, 301, 309, 313, 315, 319, 342, 344, 369, 
	374, 382, 386, 388, 393, 397, 405, 409, 411, 422, 436, 440, 442, 450, 458, 
	468, 472, 474, 477,
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
	GsParserEOF = antlr.TokenEOF
	GsParserT__0 = 1
	GsParserT__1 = 2
	GsParserT__2 = 3
	GsParserT__3 = 4
	GsParserT__4 = 5
	GsParserT__5 = 6
	GsParserENV = 7
	GsParserTRUE = 8
	GsParserFALSE = 9
	GsParserNIL = 10
	GsParserAND = 11
	GsParserOR = 12
	GsParserNOT = 13
	GsParserIF = 14
	GsParserELSE = 15
	GsParserFOR = 16
	GsParserRANGE = 17
	GsParserRETURN = 18
	GsParserFUNC = 19
	GsParserTYPE = 20
	GsParserSTRUCT = 21
	GsParserNEW = 22
	GsParserBREAK = 23
	GsParserCONTINUE = 24
	GsParserGLOBAL = 25
	GsParserGO = 26
	GsParserLEN = 27
	GsParserAPPEND = 28
	GsParserDELETE = 29
	GsParserCOPY = 30
	GsParserTOSTRING = 31
	GsParserPRINT = 32
	GsParserPRINTF = 33
	GsParserSPRINTF = 34
	GsParserPRINTLN = 35
	GsParserUINT8 = 36
	GsParserUINT16 = 37
	GsParserUINT32 = 38
	GsParserUINT64 = 39
	GsParserUINT = 40
	GsParserINT8 = 41
	GsParserINT16 = 42
	GsParserINT32 = 43
	GsParserINT64 = 44
	GsParserINTS = 45
	GsParserFLOAT32 = 46
	GsParserFLOAT64 = 47
	GsParserSTRINGS = 48
	GsParserBOOL = 49
	GsParserEXPAND = 50
	GsParserLSHIFT = 51
	GsParserRSHIFT = 52
	GsParserINCR = 53
	GsParserDECR = 54
	GsParserGEQ = 55
	GsParserLEQ = 56
	GsParserNEQ = 57
	GsParserEQ = 58
	GsParserADD = 59
	GsParserSUB = 60
	GsParserMUL = 61
	GsParserDIV = 62
	GsParserMOD = 63
	GsParserLT = 64
	GsParserGT = 65
	GsParserBITAND = 66
	GsParserBITOR = 67
	GsParserXOR = 68
	GsParserDOT = 69
	GsParserLBRACK = 70
	GsParserLPAREN = 71
	GsParserRPAREN = 72
	GsParserLBRACE = 73
	GsParserRBRACE = 74
	GsParserRBRACK = 75
	GsParserCOLON = 76
	GsParserSEMICOLON = 77
	GsParserCOMMA = 78
	GsParserINT = 79
	GsParserFLOAT = 80
	GsParserSTRING = 81
	GsParserWS = 82
	GsParserNEWLINE = 83
	GsParserSL_COMMENT = 84
	GsParserML_COMMENT = 85
	GsParserID = 86
)

// GsParser rules.
const (
	GsParserRULE_program = 0
	GsParserRULE_structDefinition = 1
	GsParserRULE_functionDefinition = 2
	GsParserRULE_block = 3
	GsParserRULE_statement = 4
	GsParserRULE_assign = 5
	GsParserRULE_incrDecr = 6
	GsParserRULE_builtinCall = 7
	GsParserRULE_iterVar = 8
	GsParserRULE_forInit = 9
	GsParserRULE_forUpdate = 10
	GsParserRULE_selfAssign = 11
	GsParserRULE_updateItem = 12
	GsParserRULE_selfAssignOp = 13
	GsParserRULE_call = 14
	GsParserRULE_expr = 15
	GsParserRULE_atom = 16
	GsParserRULE_lvalue = 17
	GsParserRULE_arrayLiteral = 18
	GsParserRULE_sliceExpr = 19
	GsParserRULE_dictLiteral = 20
	GsParserRULE_dictEntry = 21
	GsParserRULE_instance = 22
	GsParserRULE_qid = 23
	GsParserRULE_accessor = 24
	GsParserRULE_primary = 25
	GsParserRULE_compOp = 26
	GsParserRULE_addOp = 27
	GsParserRULE_mulOp = 28
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

func InitEmptyProgramContext(p *ProgramContext)  {
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 2306968908040323200) != 0) || _la == GsParserSEMICOLON || _la == GsParserID {
		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(58)
				p.FunctionDefinition()
			}


		case 2:
			{
				p.SetState(59)
				p.StructDefinition()
			}


		case 3:
			{
				p.SetState(60)
				p.Statement()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(65)
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

func InitEmptyStructDefinitionContext(p *StructDefinitionContext)  {
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
		p.SetState(67)
		p.Match(GsParserTYPE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(68)
		p.Match(GsParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(69)
		p.Match(GsParserSTRUCT)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(70)
		p.Match(GsParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(71)
		p.Match(GsParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == GsParserCOMMA {
		{
			p.SetState(72)
			p.Match(GsParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(73)
			p.Match(GsParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(79)
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

func InitEmptyFunctionDefinitionContext(p *FunctionDefinitionContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext);
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
		p.SetState(81)
		p.Match(GsParserFUNC)
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
		p.Match(GsParserLPAREN)
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


	if _la == GsParserID {
		{
			p.SetState(84)
			p.Match(GsParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == GsParserCOMMA {
			{
				p.SetState(85)
				p.Match(GsParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(86)
				p.Match(GsParserID)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(91)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(94)
		p.Match(GsParserRPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(95)
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

func InitEmptyBlockContext(p *BlockContext)  {
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
		p.SetState(97)
		p.Match(GsParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 2306968908039798912) != 0) || _la == GsParserSEMICOLON || _la == GsParserID {
		{
			p.SetState(98)
			p.Statement()
		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
						{
							p.SetState(99)
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

				p.SetState(102)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

			} else if p.HasError() { // JIM
				goto errorExit} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(104)
				p.Match(GsParserEOF)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}


		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(112)
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

func InitEmptyStatementContext(p *StatementContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForCondStmtContext) Block() IBlockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDefinitionContext); ok {
			t = ctx.(antlr.RuleContext);
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


type GlobalStmtContext struct {
	StatementContext
}

func NewGlobalStmtContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GlobalStmtContext {
	var p = new(GlobalStmtContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *GlobalStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalStmtContext) GLOBAL() antlr.TerminalNode {
	return s.GetToken(GsParserGLOBAL, 0)
}

func (s *GlobalStmtContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GsParserID)
}

func (s *GlobalStmtContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GsParserID, i)
}

func (s *GlobalStmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *GlobalStmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}


func (s *GlobalStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitGlobalStmt(s)

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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncrDecrContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForCStyleStmtContext) ForInit() IForInitContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForInitContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForInitContext)
}

func (s *ForCStyleStmtContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForCStyleStmtContext) ForUpdate() IForUpdateContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForUpdateContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuiltinCallContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelfAssignContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIterVarContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ForRangeStmtContext) Block() IBlockContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext);
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

	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEmptyStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
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
			p.SetState(115)
			p.StructDefinition()
		}


	case 3:
		localctx = NewAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(116)
			p.Assign()
		}


	case 4:
		localctx = NewSelfOpAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(117)
			p.SelfAssign()
		}


	case 5:
		localctx = NewIncrDecrStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(118)
			p.IncrDecr()
		}


	case 6:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(119)
			p.Match(GsParserRETURN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(120)
				p.expr(0)
			}
			p.SetState(125)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)


			for _la == GsParserCOMMA {
				{
					p.SetState(121)
					p.Match(GsParserCOMMA)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(122)
					p.expr(0)
				}


				p.SetState(127)
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
			p.SetState(130)
			p.Match(GsParserIF)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(131)
				p.Assign()
			}
			{
				p.SetState(132)
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
			p.SetState(136)
			p.expr(0)
		}
		{
			p.SetState(137)
			p.Block()
		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == GsParserELSE {
			{
				p.SetState(138)
				p.Match(GsParserELSE)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(139)
				p.Block()
			}

		}


	case 8:
		localctx = NewForCStyleStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(142)
			p.Match(GsParserFOR)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == GsParserENV || _la == GsParserMUL || _la == GsParserID {
			{
				p.SetState(143)
				p.ForInit()
			}

		}
		{
			p.SetState(146)
			p.Match(GsParserSEMICOLON)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 3459890412590737280) != 0) || ((int64((_la - 70)) & ^0x3f) == 0 && ((int64(1) << (_la - 70)) & 69131) != 0) {
			{
				p.SetState(147)
				p.expr(0)
			}

		}
		{
			p.SetState(150)
			p.Match(GsParserSEMICOLON)
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


		if _la == GsParserENV || _la == GsParserMUL || _la == GsParserID {
			{
				p.SetState(151)
				p.ForUpdate()
			}

		}
		{
			p.SetState(154)
			p.Block()
		}


	case 9:
		localctx = NewForRangeStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(155)
			p.Match(GsParserFOR)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(156)
			p.IterVar()
		}
		{
			p.SetState(157)
			p.Match(GsParserT__0)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Match(GsParserRANGE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(159)
			p.expr(0)
		}
		{
			p.SetState(160)
			p.Block()
		}


	case 10:
		localctx = NewForCondStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(162)
			p.Match(GsParserFOR)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(163)
			p.expr(0)
		}
		{
			p.SetState(164)
			p.Block()
		}


	case 11:
		localctx = NewBuiltinStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(166)
			p.BuiltinCall()
		}


	case 12:
		localctx = NewCallStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(167)
			p.Call()
		}


	case 13:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(168)
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
			p.SetState(169)
			p.Match(GsParserCONTINUE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 15:
		localctx = NewGlobalStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(170)
			p.Match(GsParserGLOBAL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Match(GsParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == GsParserCOMMA {
			{
				p.SetState(172)
				p.Match(GsParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(173)
				p.Match(GsParserID)
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

func InitEmptyAssignContext(p *AssignContext)  {
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
		p.SetState(181)
		p.Lvalue()
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == GsParserCOMMA {
		{
			p.SetState(182)
			p.Match(GsParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(183)
			p.Lvalue()
		}


		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(189)
		p.Match(GsParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(190)
		p.expr(0)
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == GsParserCOMMA {
		{
			p.SetState(191)
			p.Match(GsParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(192)
			p.expr(0)
		}


		p.SetState(197)
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

func InitEmptyIncrDecrContext(p *IncrDecrContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext);
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
		p.SetState(198)
		p.Lvalue()
	}
	{
		p.SetState(199)
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

func InitEmptyBuiltinCallContext(p *BuiltinCallContext)  {
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




type PrintXCallContext struct {
	BuiltinCallContext
}

func NewPrintXCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintXCallContext {
	var p = new(PrintXCallContext)

	InitEmptyBuiltinCallContext(&p.BuiltinCallContext)
	p.parser = parser
	p.CopyAll(ctx.(*BuiltinCallContext))

	return p
}

func (s *PrintXCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintXCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *PrintXCallContext) AllExpr() []IExprContext {
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

func (s *PrintXCallContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *PrintXCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *PrintXCallContext) PRINT() antlr.TerminalNode {
	return s.GetToken(GsParserPRINT, 0)
}

func (s *PrintXCallContext) PRINTLN() antlr.TerminalNode {
	return s.GetToken(GsParserPRINTLN, 0)
}

func (s *PrintXCallContext) PRINTF() antlr.TerminalNode {
	return s.GetToken(GsParserPRINTF, 0)
}

func (s *PrintXCallContext) SPRINTF() antlr.TerminalNode {
	return s.GetToken(GsParserSPRINTF, 0)
}

func (s *PrintXCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *PrintXCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}


func (s *PrintXCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitPrintXCall(s)

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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
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


type GoCallContext struct {
	BuiltinCallContext
}

func NewGoCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GoCallContext {
	var p = new(GoCallContext)

	InitEmptyBuiltinCallContext(&p.BuiltinCallContext)
	p.parser = parser
	p.CopyAll(ctx.(*BuiltinCallContext))

	return p
}

func (s *GoCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GoCallContext) GO() antlr.TerminalNode {
	return s.GetToken(GsParserGO, 0)
}

func (s *GoCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *GoCallContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GoCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}


func (s *GoCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitGoCall(s)

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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *AppendCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *AppendCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *AppendCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}

func (s *AppendCallContext) EXPAND() antlr.TerminalNode {
	return s.GetToken(GsParserEXPAND, 0)
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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



func (p *GsParser) BuiltinCall() (localctx IBuiltinCallContext) {
	localctx = NewBuiltinCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GsParserRULE_builtinCall)
	var _la int

	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GsParserLEN:
		localctx = NewLenCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(201)
			p.Match(GsParserLEN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(202)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(203)
			p.expr(0)
		}
		{
			p.SetState(204)
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
			p.SetState(206)
			p.Match(GsParserAPPEND)
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
			p.expr(0)
		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
		case 1:
			p.SetState(213)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)


			for _la == GsParserCOMMA {
				{
					p.SetState(209)
					p.Match(GsParserCOMMA)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(210)
					p.expr(0)
				}


				p.SetState(215)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
			    	goto errorExit
			    }
				_la = p.GetTokenStream().LA(1)
			}


		case 2:
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
				p.expr(0)
			}
			{
				p.SetState(218)
				p.Match(GsParserEXPAND)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(222)
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
			p.SetState(224)
			p.Match(GsParserDELETE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(225)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(226)
			p.expr(0)
		}
		{
			p.SetState(227)
			p.Match(GsParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(228)
			p.expr(0)
		}
		{
			p.SetState(229)
			p.Match(GsParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case GsParserTOSTRING:
		localctx = NewToStringCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(231)
			p.Match(GsParserTOSTRING)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(232)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(233)
			p.expr(0)
		}
		{
			p.SetState(234)
			p.Match(GsParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case GsParserPRINT, GsParserPRINTF, GsParserSPRINTF, GsParserPRINTLN:
		localctx = NewPrintXCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(236)
			_la = p.GetTokenStream().LA(1)

			if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 64424509440) != 0)) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(237)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(238)
			p.expr(0)
		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == GsParserCOMMA {
			{
				p.SetState(239)
				p.Match(GsParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(240)
				p.expr(0)
			}


			p.SetState(245)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(246)
			p.Match(GsParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case GsParserUINT8, GsParserUINT16, GsParserUINT32, GsParserUINT64, GsParserUINT, GsParserINT8, GsParserINT16, GsParserINT32, GsParserINT64, GsParserINTS, GsParserFLOAT32, GsParserFLOAT64, GsParserSTRINGS, GsParserBOOL:
		localctx = NewConvertCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(248)
			_la = p.GetTokenStream().LA(1)

			if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 1125831187365888) != 0)) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(249)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(250)
			p.expr(0)
		}
		{
			p.SetState(251)
			p.Match(GsParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case GsParserGO:
		localctx = NewGoCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(253)
			p.Match(GsParserGO)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(254)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(255)
			p.expr(0)
		}
		{
			p.SetState(256)
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

func InitEmptyIterVarContext(p *IterVarContext)  {
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
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingleIterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(260)
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
			p.SetState(261)
			p.Match(GsParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(262)
			p.Match(GsParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(263)
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

func InitEmptyForInitContext(p *ForInitContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext);
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
		p.SetState(266)
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

func InitEmptyForUpdateContext(p *ForUpdateContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdateItemContext); ok {
			t = ctx.(antlr.RuleContext);
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
		p.SetState(268)
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

func InitEmptySelfAssignContext(p *SelfAssignContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *SelfAssignContext) SelfAssignOp() ISelfAssignOpContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelfAssignOpContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelfAssignOpContext)
}

func (s *SelfAssignContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
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
		p.SetState(270)
		p.Lvalue()
	}
	{
		p.SetState(271)
		p.SelfAssignOp()
	}
	{
		p.SetState(272)
		p.expr(0)
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

func InitEmptyUpdateItemContext(p *UpdateItemContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelfAssignContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncrDecrContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignContext); ok {
			t = ctx.(antlr.RuleContext);
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
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSelfUpdateContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)
			p.SelfAssign()
		}


	case 2:
		localctx = NewIncrDecrUpdateContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(275)
			p.IncrDecr()
		}


	case 3:
		localctx = NewAssignUpdateContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(276)
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

func InitEmptySelfAssignOpContext(p *SelfAssignOpContext)  {
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
		p.SetState(279)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 124) != 0)) {
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

func InitEmptyCallContext(p *CallContext)  {
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

func (s *CallContext) CopyAll(ctx *CallContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type OuterCallContext struct {
	CallContext
}

func NewOuterCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OuterCallContext {
	var p = new(OuterCallContext)

	InitEmptyCallContext(&p.CallContext)
	p.parser = parser
	p.CopyAll(ctx.(*CallContext))

	return p
}

func (s *OuterCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OuterCallContext) Primary() IPrimaryContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *OuterCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *OuterCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *OuterCallContext) AllAccessor() []IAccessorContext {
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

func (s *OuterCallContext) Accessor(i int) IAccessorContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *OuterCallContext) AllExpr() []IExprContext {
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

func (s *OuterCallContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *OuterCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *OuterCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}


func (s *OuterCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitOuterCall(s)

	default:
		return t.VisitChildren(s)
	}
}


type InnerCallContext struct {
	CallContext
}

func NewInnerCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InnerCallContext {
	var p = new(InnerCallContext)

	InitEmptyCallContext(&p.CallContext)
	p.parser = parser
	p.CopyAll(ctx.(*CallContext))

	return p
}

func (s *InnerCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerCallContext) ID() antlr.TerminalNode {
	return s.GetToken(GsParserID, 0)
}

func (s *InnerCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *InnerCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *InnerCallContext) AllExpr() []IExprContext {
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

func (s *InnerCallContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *InnerCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *InnerCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}


func (s *InnerCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitInnerCall(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *GsParser) Call() (localctx ICallContext) {
	localctx = NewCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GsParserRULE_call)
	var _la int

	var _alt int

	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewInnerCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)
			p.Match(GsParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(282)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 3459890412590737280) != 0) || ((int64((_la - 70)) & ^0x3f) == 0 && ((int64(1) << (_la - 70)) & 69131) != 0) {
			{
				p.SetState(283)
				p.expr(0)
			}
			p.SetState(288)
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
						p.SetState(284)
						p.Match(GsParserCOMMA)
						if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
						}
					}
					{
						p.SetState(285)
						p.expr(0)
					}


				}
				p.SetState(290)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
			    	goto errorExit
			    }
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}
			p.SetState(292)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)


			if _la == GsParserCOMMA {
				{
					p.SetState(291)
					p.Match(GsParserCOMMA)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}

			}

		}
		{
			p.SetState(296)
			p.Match(GsParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 2:
		localctx = NewOuterCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(297)
			p.Primary()
		}
		p.SetState(299)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
					{
						p.SetState(298)
						p.Accessor()
					}




			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(301)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(303)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 3459890412590737280) != 0) || ((int64((_la - 70)) & ^0x3f) == 0 && ((int64(1) << (_la - 70)) & 69131) != 0) {
			{
				p.SetState(304)
				p.expr(0)
			}
			p.SetState(309)
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
						p.SetState(305)
						p.Match(GsParserCOMMA)
						if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
						}
					}
					{
						p.SetState(306)
						p.expr(0)
					}


				}
				p.SetState(311)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
			    	goto errorExit
			    }
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}
			p.SetState(313)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)


			if _la == GsParserCOMMA {
				{
					p.SetState(312)
					p.Match(GsParserCOMMA)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}

			}

		}
		{
			p.SetState(317)
			p.Match(GsParserRPAREN)
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


// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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

func InitEmptyExprContext(p *ExprContext)  {
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

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type LogicalOrExprContext struct {
	ExprContext
}

func NewLogicalOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalOrExprContext {
	var p = new(LogicalOrExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LogicalOrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExprContext) AllExpr() []IExprContext {
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

func (s *LogicalOrExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *LogicalOrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(GsParserOR, 0)
}


func (s *LogicalOrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitLogicalOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type LogicalAndExprContext struct {
	ExprContext
}

func NewLogicalAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalAndExprContext {
	var p = new(LogicalAndExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *LogicalAndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExprContext) AllExpr() []IExprContext {
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

func (s *LogicalAndExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *LogicalAndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(GsParserAND, 0)
}


func (s *LogicalAndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitLogicalAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type AddExprContext struct {
	ExprContext
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllExpr() []IExprContext {
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

func (s *AddExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *AddExprContext) AddOp() IAddOpContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddOpContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddOpContext)
}


func (s *AddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitAddExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type MulExprContext struct {
	ExprContext
}

func NewMulExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExprContext {
	var p = new(MulExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExprContext) AllExpr() []IExprContext {
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

func (s *MulExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *MulExprContext) MulOp() IMulOpContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulOpContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMulOpContext)
}


func (s *MulExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitMulExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type AtomExprContext struct {
	ExprContext
}

func NewAtomExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomExprContext {
	var p = new(AtomExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AtomExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomExprContext) Atom() IAtomContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}


func (s *AtomExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitAtomExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type ComparisonExprContext struct {
	ExprContext
}

func NewComparisonExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonExprContext {
	var p = new(ComparisonExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ComparisonExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonExprContext) AllExpr() []IExprContext {
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

func (s *ComparisonExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *ComparisonExprContext) CompOp() ICompOpContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompOpContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompOpContext)
}


func (s *ComparisonExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitComparisonExpr(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *GsParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *GsParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, GsParserRULE_expr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewAtomExprContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(322)
		p.Atom()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(342)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GsParserRULE_expr)
				p.SetState(324)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(325)
					p.MulOp()
				}
				{
					p.SetState(326)
					p.expr(6)
				}


			case 2:
				localctx = NewAddExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GsParserRULE_expr)
				p.SetState(328)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(329)
					p.AddOp()
				}
				{
					p.SetState(330)
					p.expr(5)
				}


			case 3:
				localctx = NewComparisonExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GsParserRULE_expr)
				p.SetState(332)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(333)
					p.CompOp()
				}
				{
					p.SetState(334)
					p.expr(4)
				}


			case 4:
				localctx = NewLogicalAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GsParserRULE_expr)
				p.SetState(336)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(337)
					p.Match(GsParserAND)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(338)
					p.expr(3)
				}


			case 5:
				localctx = NewLogicalOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GsParserRULE_expr)
				p.SetState(339)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(340)
					p.Match(GsParserOR)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(341)
					p.expr(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
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
	p.UnrollRecursionContexts(_parentctx)
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

func InitEmptyAtomContext(p *AtomContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuiltinCallContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQidContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictLiteralContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstanceContext); ok {
			t = ctx.(antlr.RuleContext);
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

func (s *NotAtomContext) Atom() IAtomContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
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
	p.EnterRule(localctx, 32, GsParserRULE_atom)
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNegAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(347)
			p.Match(GsParserSUB)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(348)
			p.Atom()
		}


	case 2:
		localctx = NewNotAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(349)
			p.Match(GsParserNOT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(350)
			p.Atom()
		}


	case 3:
		localctx = NewDerefAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(351)
			p.Match(GsParserMUL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(352)
			p.Lvalue()
		}


	case 4:
		localctx = NewIntAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(353)
			p.Match(GsParserINT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 5:
		localctx = NewFloatAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(354)
			p.Match(GsParserFLOAT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 6:
		localctx = NewStringAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(355)
			p.Match(GsParserSTRING)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 7:
		localctx = NewTrueAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(356)
			p.Match(GsParserTRUE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 8:
		localctx = NewFalseAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(357)
			p.Match(GsParserFALSE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 9:
		localctx = NewNilAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(358)
			p.Match(GsParserNIL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 10:
		localctx = NewBuiltinAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(359)
			p.BuiltinCall()
		}


	case 11:
		localctx = NewCallAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(360)
			p.Call()
		}


	case 12:
		localctx = NewInstanceAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(361)
			p.Instance()
		}


	case 13:
		localctx = NewArrayAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(362)
			p.ArrayLiteral()
		}


	case 14:
		localctx = NewDictAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(363)
			p.DictLiteral()
		}


	case 15:
		localctx = NewParenAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(364)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(365)
			p.expr(0)
		}
		{
			p.SetState(366)
			p.Match(GsParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case 16:
		localctx = NewQidAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(368)
			p.Qid()
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

func InitEmptyLvalueContext(p *LvalueContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQidContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext);
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
	p.EnterRule(localctx, 34, GsParserRULE_lvalue)
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GsParserENV, GsParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(371)
			p.Qid()
		}


	case GsParserMUL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(372)
			p.Match(GsParserMUL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(373)
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

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext)  {
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	p.EnterRule(localctx, 36, GsParserRULE_arrayLiteral)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(GsParserLBRACK)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 3459890412590737280) != 0) || ((int64((_la - 70)) & ^0x3f) == 0 && ((int64(1) << (_la - 70)) & 69131) != 0) {
		{
			p.SetState(377)
			p.expr(0)
		}
		p.SetState(382)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(378)
					p.Match(GsParserCOMMA)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(379)
					p.expr(0)
				}


			}
			p.SetState(384)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == GsParserCOMMA {
			{
				p.SetState(385)
				p.Match(GsParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}

	}
	{
		p.SetState(390)
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

func InitEmptySliceExprContext(p *SliceExprContext)  {
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	p.EnterRule(localctx, 38, GsParserRULE_sliceExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 3459890412590737280) != 0) || ((int64((_la - 70)) & ^0x3f) == 0 && ((int64(1) << (_la - 70)) & 69131) != 0) {
		{
			p.SetState(392)
			p.expr(0)
		}

	}
	{
		p.SetState(395)
		p.Match(GsParserCOLON)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 3459890412590737280) != 0) || ((int64((_la - 70)) & ^0x3f) == 0 && ((int64(1) << (_la - 70)) & 69131) != 0) {
		{
			p.SetState(396)
			p.expr(0)
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

func InitEmptyDictLiteralContext(p *DictLiteralContext)  {
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDictEntryContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	p.EnterRule(localctx, 40, GsParserRULE_dictLiteral)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		p.Match(GsParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 2305843009213694848) != 0) || ((int64((_la - 79)) & ^0x3f) == 0 && ((int64(1) << (_la - 79)) & 135) != 0) {
		{
			p.SetState(400)
			p.DictEntry()
		}
		p.SetState(405)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(401)
					p.Match(GsParserCOMMA)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(402)
					p.DictEntry()
				}


			}
			p.SetState(407)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == GsParserCOMMA {
			{
				p.SetState(408)
				p.Match(GsParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}

	}
	{
		p.SetState(413)
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

func InitEmptyDictEntryContext(p *DictEntryContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
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
	p.EnterRule(localctx, 42, GsParserRULE_dictEntry)
	var _la int

	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GsParserTRUE, GsParserFALSE, GsParserINT, GsParserFLOAT, GsParserSTRING:
		localctx = NewConstKeyEntryContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(415)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GsParserTRUE || _la == GsParserFALSE || ((int64((_la - 79)) & ^0x3f) == 0 && ((int64(1) << (_la - 79)) & 7) != 0)) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(416)
			p.Match(GsParserCOLON)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(417)
			p.expr(0)
		}


	case GsParserENV, GsParserMUL, GsParserID:
		localctx = NewIdKeyEntryContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(418)
			p.Lvalue()
		}
		{
			p.SetState(419)
			p.Match(GsParserCOLON)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(420)
			p.expr(0)
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

func InitEmptyInstanceContext(p *InstanceContext)  {
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	p.EnterRule(localctx, 44, GsParserRULE_instance)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(424)
		p.Match(GsParserNEW)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(425)
		p.Match(GsParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(426)
		p.Match(GsParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(442)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == GsParserID {
		{
			p.SetState(427)
			p.Match(GsParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(428)
			p.Match(GsParserCOLON)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(429)
			p.expr(0)
		}
		p.SetState(436)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(430)
					p.Match(GsParserCOMMA)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(431)
					p.Match(GsParserID)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(432)
					p.Match(GsParserCOLON)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(433)
					p.expr(0)
				}


			}
			p.SetState(438)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(440)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == GsParserCOMMA {
			{
				p.SetState(439)
				p.Match(GsParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}

	}
	{
		p.SetState(444)
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

func InitEmptyQidContext(p *QidContext)  {
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
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryContext); ok {
			t = ctx.(antlr.RuleContext);
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
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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
	p.EnterRule(localctx, 46, GsParserRULE_qid)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(446)
		p.Primary()
	}
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(447)
				p.Accessor()
			}


		}
		p.SetState(452)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext())
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

func InitEmptyAccessorContext(p *AccessorContext)  {
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

func (s *IndexAccessContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(GsParserLBRACK, 0)
}

func (s *IndexAccessContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(GsParserRBRACK, 0)
}

func (s *IndexAccessContext) Expr() IExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IndexAccessContext) SliceExpr() ISliceExprContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceExprContext); ok {
			t = ctx.(antlr.RuleContext);
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


type MethodCallAccessContext struct {
	AccessorContext
}

func NewMethodCallAccessContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MethodCallAccessContext {
	var p = new(MethodCallAccessContext)

	InitEmptyAccessorContext(&p.AccessorContext)
	p.parser = parser
	p.CopyAll(ctx.(*AccessorContext))

	return p
}

func (s *MethodCallAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallAccessContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *MethodCallAccessContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *MethodCallAccessContext) AllExpr() []IExprContext {
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

func (s *MethodCallAccessContext) Expr(i int) IExprContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
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

func (s *MethodCallAccessContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GsParserCOMMA)
}

func (s *MethodCallAccessContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, i)
}


func (s *MethodCallAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitMethodCallAccess(s)

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

func (s *PropertyAccessContext) DOT() antlr.TerminalNode {
	return s.GetToken(GsParserDOT, 0)
}

func (s *PropertyAccessContext) ID() antlr.TerminalNode {
	return s.GetToken(GsParserID, 0)
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
	p.EnterRule(localctx, 48, GsParserRULE_accessor)
	var _la int

	var _alt int

	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GsParserDOT:
		localctx = NewPropertyAccessContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(453)
			p.Match(GsParserDOT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(454)
			p.Match(GsParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case GsParserLBRACK:
		localctx = NewIndexAccessContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(455)
			p.Match(GsParserLBRACK)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(458)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(456)
				p.expr(0)
			}


		case 2:
			{
				p.SetState(457)
				p.SliceExpr()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(460)
			p.Match(GsParserRBRACK)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case GsParserLPAREN:
		localctx = NewMethodCallAccessContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(462)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(474)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 3459890412590737280) != 0) || ((int64((_la - 70)) & ^0x3f) == 0 && ((int64(1) << (_la - 70)) & 69131) != 0) {
			{
				p.SetState(463)
				p.expr(0)
			}
			p.SetState(468)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(464)
						p.Match(GsParserCOMMA)
						if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
						}
					}
					{
						p.SetState(465)
						p.expr(0)
					}


				}
				p.SetState(470)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
			    	goto errorExit
			    }
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 51, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}
			p.SetState(472)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)


			if _la == GsParserCOMMA {
				{
					p.SetState(471)
					p.Match(GsParserCOMMA)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}

			}

		}
		{
			p.SetState(476)
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

func InitEmptyPrimaryContext(p *PrimaryContext)  {
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
	p.EnterRule(localctx, 50, GsParserRULE_primary)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(479)
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

func InitEmptyCompOpContext(p *CompOpContext)  {
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
	p.EnterRule(localctx, 52, GsParserRULE_compOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(481)
		_la = p.GetTokenStream().LA(1)

		if !(((int64((_la - 55)) & ^0x3f) == 0 && ((int64(1) << (_la - 55)) & 1551) != 0)) {
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
	BITOR() antlr.TerminalNode
	XOR() antlr.TerminalNode

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

func InitEmptyAddOpContext(p *AddOpContext)  {
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

func (s *AddOpContext) BITOR() antlr.TerminalNode {
	return s.GetToken(GsParserBITOR, 0)
}

func (s *AddOpContext) XOR() antlr.TerminalNode {
	return s.GetToken(GsParserXOR, 0)
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
	p.EnterRule(localctx, 54, GsParserRULE_addOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(483)
		_la = p.GetTokenStream().LA(1)

		if !(((int64((_la - 59)) & ^0x3f) == 0 && ((int64(1) << (_la - 59)) & 771) != 0)) {
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
	LSHIFT() antlr.TerminalNode
	RSHIFT() antlr.TerminalNode
	BITAND() antlr.TerminalNode

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

func InitEmptyMulOpContext(p *MulOpContext)  {
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

func (s *MulOpContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(GsParserLSHIFT, 0)
}

func (s *MulOpContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(GsParserRSHIFT, 0)
}

func (s *MulOpContext) BITAND() antlr.TerminalNode {
	return s.GetToken(GsParserBITAND, 0)
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
	p.EnterRule(localctx, 56, GsParserRULE_mulOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(485)
		_la = p.GetTokenStream().LA(1)

		if !(((int64((_la - 51)) & ^0x3f) == 0 && ((int64(1) << (_la - 51)) & 39939) != 0)) {
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


func (p *GsParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 15:
			var t *ExprContext = nil
			if localctx != nil { t = localctx.(*ExprContext) }
			return p.Expr_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GsParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
			return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
			return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
			return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
			return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

