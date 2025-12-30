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
    "'toString'", "'print'", "'printf'", "'sprintf'", "'println'", "'initRef'", 
    "'newFromType'", "'uint8'", "'uint16'", "'uint32'", "'uint64'", "'uint'", 
    "'int8'", "'int16'", "'int32'", "'int64'", "'int'", "'float32'", "'float64'", 
    "'string'", "'bool'", "'...'", "'<<'", "'>>'", "'++'", "'--'", "'>='", 
    "'<='", "'!='", "'=='", "'+'", "'-'", "'*'", "'/'", "'%'", "'<'", "'>'", 
    "'&'", "'|'", "'^'", "'.'", "'['", "'('", "')'", "'{'", "'}'", "']'", 
    "':'", "';'", "','",
  }
  staticData.SymbolicNames = []string{
    "", "", "", "", "", "", "", "ENV", "TRUE", "FALSE", "NIL", "AND", "OR", 
    "NOT", "IF", "ELSE", "FOR", "RANGE", "RETURN", "FUNC", "TYPE", "STRUCT", 
    "NEW", "BREAK", "CONTINUE", "GLOBAL", "GO", "LEN", "APPEND", "DELETE", 
    "COPY", "TOSTRING", "PRINT", "PRINTF", "SPRINTF", "PRINTLN", "INITREF", 
    "NEWFROMTYPE", "UINT8", "UINT16", "UINT32", "UINT64", "UINT", "INT8", 
    "INT16", "INT32", "INT64", "INTS", "FLOAT32", "FLOAT64", "STRINGS", 
    "BOOL", "EXPAND", "LSHIFT", "RSHIFT", "INCR", "DECR", "GEQ", "LEQ", 
    "NEQ", "EQ", "ADD", "SUB", "MUL", "DIV", "MOD", "LT", "GT", "BITAND", 
    "BITOR", "XOR", "DOT", "LBRACK", "LPAREN", "RPAREN", "LBRACE", "RBRACE", 
    "RBRACK", "COLON", "SEMICOLON", "COMMA", "INT", "FLOAT", "STRING", "WS", 
    "NEWLINE", "SL_COMMENT", "ML_COMMENT", "ID",
  }
  staticData.RuleNames = []string{
    "program", "structDefinition", "functionDefinition", "block", "statement", 
    "ifStatement", "simpleStmt", "assign", "incrDecr", "builtinCall", "iterVar", 
    "forInit", "forUpdate", "selfAssign", "updateItem", "selfAssignOp", 
    "call", "expr", "atom", "lvalue", "arrayLiteral", "sliceExpr", "dictLiteral", 
    "dictEntry", "instance", "qid", "accessor", "primary", "compOp", "addOp", 
    "mulOp",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 88, 509, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 
	10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15, 
	2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2, 
	21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26, 
	7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 1, 0, 1, 
	0, 1, 0, 4, 0, 66, 8, 0, 11, 0, 12, 0, 67, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 
	1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 79, 8, 1, 10, 1, 12, 1, 82, 9, 1, 1, 1, 1, 
	1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 92, 8, 2, 10, 2, 12, 2, 95, 
	9, 2, 3, 2, 97, 8, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 4, 3, 105, 8, 
	3, 11, 3, 12, 3, 106, 1, 3, 3, 3, 110, 8, 3, 5, 3, 112, 8, 3, 10, 3, 12, 
	3, 115, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 
	4, 1, 4, 5, 4, 128, 8, 4, 10, 4, 12, 4, 131, 9, 4, 3, 4, 133, 8, 4, 1, 
	4, 1, 4, 1, 4, 3, 4, 138, 8, 4, 1, 4, 1, 4, 3, 4, 142, 8, 4, 1, 4, 1, 4, 
	3, 4, 146, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 
	1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 
	168, 8, 4, 10, 4, 12, 4, 171, 9, 4, 3, 4, 173, 8, 4, 1, 5, 1, 5, 1, 5, 
	1, 5, 3, 5, 179, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 186, 8, 5, 3, 
	5, 188, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 195, 8, 6, 1, 7, 1, 7, 
	1, 7, 5, 7, 200, 8, 7, 10, 7, 12, 7, 203, 9, 7, 1, 7, 1, 7, 1, 7, 1, 7, 
	5, 7, 209, 8, 7, 10, 7, 12, 7, 212, 9, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 
	1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 227, 8, 9, 1, 9, 
	1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 
	241, 8, 9, 10, 9, 12, 9, 244, 9, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 250, 
	8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 
	1, 9, 1, 9, 1, 9, 5, 9, 266, 8, 9, 10, 9, 12, 9, 269, 9, 9, 1, 9, 1, 9, 
	1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 280, 8, 9, 1, 10, 1, 10, 
	1, 10, 1, 10, 3, 10, 286, 8, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 
	13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 3, 14, 299, 8, 14, 1, 15, 1, 15, 
	1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 308, 8, 16, 10, 16, 12, 16, 311, 
	9, 16, 1, 16, 3, 16, 314, 8, 16, 3, 16, 316, 8, 16, 1, 16, 1, 16, 1, 16, 
	4, 16, 321, 8, 16, 11, 16, 12, 16, 322, 1, 16, 1, 16, 1, 16, 1, 16, 5, 
	16, 329, 8, 16, 10, 16, 12, 16, 332, 9, 16, 1, 16, 3, 16, 335, 8, 16, 3, 
	16, 337, 8, 16, 1, 16, 1, 16, 3, 16, 341, 8, 16, 1, 17, 1, 17, 1, 17, 1, 
	17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 
	1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 364, 8, 17, 10, 
	17, 12, 17, 367, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 
	1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 
	18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 391, 8, 18, 1, 19, 1, 19, 1, 19, 
	3, 19, 396, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 402, 8, 20, 10, 20, 
	12, 20, 405, 9, 20, 1, 20, 3, 20, 408, 8, 20, 3, 20, 410, 8, 20, 1, 20, 
	1, 20, 1, 21, 3, 21, 415, 8, 21, 1, 21, 1, 21, 3, 21, 419, 8, 21, 1, 22, 
	1, 22, 1, 22, 1, 22, 5, 22, 425, 8, 22, 10, 22, 12, 22, 428, 9, 22, 1, 
	22, 3, 22, 431, 8, 22, 3, 22, 433, 8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 
	23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 444, 8, 23, 1, 24, 1, 24, 1, 24, 
	1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 456, 8, 24, 10, 
	24, 12, 24, 459, 9, 24, 1, 24, 3, 24, 462, 8, 24, 3, 24, 464, 8, 24, 1, 
	24, 1, 24, 1, 25, 1, 25, 5, 25, 470, 8, 25, 10, 25, 12, 25, 473, 9, 25, 
	1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 480, 8, 26, 1, 26, 1, 26, 1, 
	26, 1, 26, 1, 26, 1, 26, 5, 26, 488, 8, 26, 10, 26, 12, 26, 491, 9, 26, 
	1, 26, 3, 26, 494, 8, 26, 3, 26, 496, 8, 26, 1, 26, 3, 26, 499, 8, 26, 
	1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 0, 1, 34, 
	31, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 0, 10, 2, 0, 79, 79, 
	85, 85, 1, 0, 55, 56, 1, 0, 32, 35, 1, 0, 38, 51, 1, 0, 2, 6, 2, 0, 8, 
	9, 81, 83, 2, 0, 7, 7, 88, 88, 2, 0, 57, 60, 66, 67, 2, 0, 61, 62, 69, 
	70, 3, 0, 53, 54, 63, 65, 68, 68, 578, 0, 65, 1, 0, 0, 0, 2, 71, 1, 0, 
	0, 0, 4, 85, 1, 0, 0, 0, 6, 101, 1, 0, 0, 0, 8, 172, 1, 0, 0, 0, 10, 174, 
	1, 0, 0, 0, 12, 194, 1, 0, 0, 0, 14, 196, 1, 0, 0, 0, 16, 213, 1, 0, 0, 
	0, 18, 279, 1, 0, 0, 0, 20, 285, 1, 0, 0, 0, 22, 287, 1, 0, 0, 0, 24, 289, 
	1, 0, 0, 0, 26, 291, 1, 0, 0, 0, 28, 298, 1, 0, 0, 0, 30, 300, 1, 0, 0, 
	0, 32, 340, 1, 0, 0, 0, 34, 342, 1, 0, 0, 0, 36, 390, 1, 0, 0, 0, 38, 395, 
	1, 0, 0, 0, 40, 397, 1, 0, 0, 0, 42, 414, 1, 0, 0, 0, 44, 420, 1, 0, 0, 
	0, 46, 443, 1, 0, 0, 0, 48, 445, 1, 0, 0, 0, 50, 467, 1, 0, 0, 0, 52, 498, 
	1, 0, 0, 0, 54, 500, 1, 0, 0, 0, 56, 502, 1, 0, 0, 0, 58, 504, 1, 0, 0, 
	0, 60, 506, 1, 0, 0, 0, 62, 66, 3, 4, 2, 0, 63, 66, 3, 2, 1, 0, 64, 66, 
	3, 8, 4, 0, 65, 62, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 64, 1, 0, 0, 0, 
	66, 67, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 69, 1, 
	0, 0, 0, 69, 70, 5, 0, 0, 1, 70, 1, 1, 0, 0, 0, 71, 72, 5, 20, 0, 0, 72, 
	73, 5, 88, 0, 0, 73, 74, 5, 21, 0, 0, 74, 75, 5, 75, 0, 0, 75, 80, 5, 88, 
	0, 0, 76, 77, 5, 80, 0, 0, 77, 79, 5, 88, 0, 0, 78, 76, 1, 0, 0, 0, 79, 
	82, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 83, 1, 0, 0, 
	0, 82, 80, 1, 0, 0, 0, 83, 84, 5, 76, 0, 0, 84, 3, 1, 0, 0, 0, 85, 86, 
	5, 19, 0, 0, 86, 87, 5, 88, 0, 0, 87, 96, 5, 73, 0, 0, 88, 93, 5, 88, 0, 
	0, 89, 90, 5, 80, 0, 0, 90, 92, 5, 88, 0, 0, 91, 89, 1, 0, 0, 0, 92, 95, 
	1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0, 
	95, 93, 1, 0, 0, 0, 96, 88, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 98, 1, 
	0, 0, 0, 98, 99, 5, 74, 0, 0, 99, 100, 3, 6, 3, 0, 100, 5, 1, 0, 0, 0, 
	101, 113, 5, 75, 0, 0, 102, 109, 3, 8, 4, 0, 103, 105, 7, 0, 0, 0, 104, 
	103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 107, 
	1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 110, 5, 0, 0, 1, 109, 104, 1, 0, 
	0, 0, 109, 108, 1, 0, 0, 0, 109, 110, 1, 0, 0, 0, 110, 112, 1, 0, 0, 0, 
	111, 102, 1, 0, 0, 0, 112, 115, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 113, 
	114, 1, 0, 0, 0, 114, 116, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116, 117, 
	5, 76, 0, 0, 117, 7, 1, 0, 0, 0, 118, 173, 5, 79, 0, 0, 119, 173, 3, 2, 
	1, 0, 120, 173, 3, 14, 7, 0, 121, 173, 3, 26, 13, 0, 122, 173, 3, 16, 8, 
	0, 123, 132, 5, 18, 0, 0, 124, 129, 3, 34, 17, 0, 125, 126, 5, 80, 0, 0, 
	126, 128, 3, 34, 17, 0, 127, 125, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 
	127, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 133, 1, 0, 0, 0, 131, 129, 
	1, 0, 0, 0, 132, 124, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 173, 1, 0, 
	0, 0, 134, 173, 3, 10, 5, 0, 135, 137, 5, 16, 0, 0, 136, 138, 3, 22, 11, 
	0, 137, 136, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 
	141, 5, 79, 0, 0, 140, 142, 3, 34, 17, 0, 141, 140, 1, 0, 0, 0, 141, 142, 
	1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 145, 5, 79, 0, 0, 144, 146, 3, 24, 
	12, 0, 145, 144, 1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 
	147, 173, 3, 6, 3, 0, 148, 149, 5, 16, 0, 0, 149, 150, 3, 20, 10, 0, 150, 
	151, 5, 1, 0, 0, 151, 152, 5, 17, 0, 0, 152, 153, 3, 34, 17, 0, 153, 154, 
	3, 6, 3, 0, 154, 173, 1, 0, 0, 0, 155, 156, 5, 16, 0, 0, 156, 157, 3, 34, 
	17, 0, 157, 158, 3, 6, 3, 0, 158, 173, 1, 0, 0, 0, 159, 173, 3, 18, 9, 
	0, 160, 173, 3, 32, 16, 0, 161, 173, 5, 23, 0, 0, 162, 173, 5, 24, 0, 0, 
	163, 164, 5, 25, 0, 0, 164, 169, 5, 88, 0, 0, 165, 166, 5, 80, 0, 0, 166, 
	168, 5, 88, 0, 0, 167, 165, 1, 0, 0, 0, 168, 171, 1, 0, 0, 0, 169, 167, 
	1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 169, 1, 0, 
	0, 0, 172, 118, 1, 0, 0, 0, 172, 119, 1, 0, 0, 0, 172, 120, 1, 0, 0, 0, 
	172, 121, 1, 0, 0, 0, 172, 122, 1, 0, 0, 0, 172, 123, 1, 0, 0, 0, 172, 
	134, 1, 0, 0, 0, 172, 135, 1, 0, 0, 0, 172, 148, 1, 0, 0, 0, 172, 155, 
	1, 0, 0, 0, 172, 159, 1, 0, 0, 0, 172, 160, 1, 0, 0, 0, 172, 161, 1, 0, 
	0, 0, 172, 162, 1, 0, 0, 0, 172, 163, 1, 0, 0, 0, 173, 9, 1, 0, 0, 0, 174, 
	178, 5, 14, 0, 0, 175, 176, 3, 12, 6, 0, 176, 177, 5, 79, 0, 0, 177, 179, 
	1, 0, 0, 0, 178, 175, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 1, 0, 
	0, 0, 180, 181, 3, 34, 17, 0, 181, 187, 3, 6, 3, 0, 182, 185, 5, 15, 0, 
	0, 183, 186, 3, 6, 3, 0, 184, 186, 3, 10, 5, 0, 185, 183, 1, 0, 0, 0, 185, 
	184, 1, 0, 0, 0, 186, 188, 1, 0, 0, 0, 187, 182, 1, 0, 0, 0, 187, 188, 
	1, 0, 0, 0, 188, 11, 1, 0, 0, 0, 189, 195, 3, 14, 7, 0, 190, 195, 3, 26, 
	13, 0, 191, 195, 3, 16, 8, 0, 192, 195, 3, 18, 9, 0, 193, 195, 3, 32, 16, 
	0, 194, 189, 1, 0, 0, 0, 194, 190, 1, 0, 0, 0, 194, 191, 1, 0, 0, 0, 194, 
	192, 1, 0, 0, 0, 194, 193, 1, 0, 0, 0, 195, 13, 1, 0, 0, 0, 196, 201, 3, 
	38, 19, 0, 197, 198, 5, 80, 0, 0, 198, 200, 3, 38, 19, 0, 199, 197, 1, 
	0, 0, 0, 200, 203, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 
	0, 202, 204, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 204, 205, 5, 1, 0, 0, 205, 
	210, 3, 34, 17, 0, 206, 207, 5, 80, 0, 0, 207, 209, 3, 34, 17, 0, 208, 
	206, 1, 0, 0, 0, 209, 212, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210, 211, 
	1, 0, 0, 0, 211, 15, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 213, 214, 3, 38, 
	19, 0, 214, 215, 7, 1, 0, 0, 215, 17, 1, 0, 0, 0, 216, 217, 5, 27, 0, 0, 
	217, 218, 5, 73, 0, 0, 218, 219, 3, 34, 17, 0, 219, 220, 5, 74, 0, 0, 220, 
	280, 1, 0, 0, 0, 221, 222, 5, 36, 0, 0, 222, 223, 5, 73, 0, 0, 223, 226, 
	3, 50, 25, 0, 224, 225, 5, 80, 0, 0, 225, 227, 3, 34, 17, 0, 226, 224, 
	1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 229, 5, 74, 
	0, 0, 229, 280, 1, 0, 0, 0, 230, 231, 5, 37, 0, 0, 231, 232, 5, 73, 0, 
	0, 232, 233, 3, 34, 17, 0, 233, 234, 5, 74, 0, 0, 234, 280, 1, 0, 0, 0, 
	235, 236, 5, 28, 0, 0, 236, 237, 5, 73, 0, 0, 237, 249, 3, 34, 17, 0, 238, 
	239, 5, 80, 0, 0, 239, 241, 3, 34, 17, 0, 240, 238, 1, 0, 0, 0, 241, 244, 
	1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 250, 1, 0, 
	0, 0, 244, 242, 1, 0, 0, 0, 245, 246, 5, 80, 0, 0, 246, 247, 3, 34, 17, 
	0, 247, 248, 5, 52, 0, 0, 248, 250, 1, 0, 0, 0, 249, 242, 1, 0, 0, 0, 249, 
	245, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 252, 5, 74, 0, 0, 252, 280, 
	1, 0, 0, 0, 253, 254, 5, 29, 0, 0, 254, 255, 5, 73, 0, 0, 255, 256, 3, 
	34, 17, 0, 256, 257, 5, 80, 0, 0, 257, 258, 3, 34, 17, 0, 258, 259, 5, 
	74, 0, 0, 259, 280, 1, 0, 0, 0, 260, 261, 7, 2, 0, 0, 261, 262, 5, 73, 
	0, 0, 262, 267, 3, 34, 17, 0, 263, 264, 5, 80, 0, 0, 264, 266, 3, 34, 17, 
	0, 265, 263, 1, 0, 0, 0, 266, 269, 1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 267, 
	268, 1, 0, 0, 0, 268, 270, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 270, 271, 
	5, 74, 0, 0, 271, 280, 1, 0, 0, 0, 272, 273, 7, 3, 0, 0, 273, 274, 5, 73, 
	0, 0, 274, 275, 3, 34, 17, 0, 275, 276, 5, 74, 0, 0, 276, 280, 1, 0, 0, 
	0, 277, 278, 5, 26, 0, 0, 278, 280, 3, 32, 16, 0, 279, 216, 1, 0, 0, 0, 
	279, 221, 1, 0, 0, 0, 279, 230, 1, 0, 0, 0, 279, 235, 1, 0, 0, 0, 279, 
	253, 1, 0, 0, 0, 279, 260, 1, 0, 0, 0, 279, 272, 1, 0, 0, 0, 279, 277, 
	1, 0, 0, 0, 280, 19, 1, 0, 0, 0, 281, 286, 5, 88, 0, 0, 282, 283, 5, 88, 
	0, 0, 283, 284, 5, 80, 0, 0, 284, 286, 5, 88, 0, 0, 285, 281, 1, 0, 0, 
	0, 285, 282, 1, 0, 0, 0, 286, 21, 1, 0, 0, 0, 287, 288, 3, 14, 7, 0, 288, 
	23, 1, 0, 0, 0, 289, 290, 3, 28, 14, 0, 290, 25, 1, 0, 0, 0, 291, 292, 
	3, 38, 19, 0, 292, 293, 3, 30, 15, 0, 293, 294, 3, 34, 17, 0, 294, 27, 
	1, 0, 0, 0, 295, 299, 3, 26, 13, 0, 296, 299, 3, 16, 8, 0, 297, 299, 3, 
	14, 7, 0, 298, 295, 1, 0, 0, 0, 298, 296, 1, 0, 0, 0, 298, 297, 1, 0, 0, 
	0, 299, 29, 1, 0, 0, 0, 300, 301, 7, 4, 0, 0, 301, 31, 1, 0, 0, 0, 302, 
	303, 5, 88, 0, 0, 303, 315, 5, 73, 0, 0, 304, 309, 3, 34, 17, 0, 305, 306, 
	5, 80, 0, 0, 306, 308, 3, 34, 17, 0, 307, 305, 1, 0, 0, 0, 308, 311, 1, 
	0, 0, 0, 309, 307, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 313, 1, 0, 0, 
	0, 311, 309, 1, 0, 0, 0, 312, 314, 5, 80, 0, 0, 313, 312, 1, 0, 0, 0, 313, 
	314, 1, 0, 0, 0, 314, 316, 1, 0, 0, 0, 315, 304, 1, 0, 0, 0, 315, 316, 
	1, 0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 341, 5, 74, 0, 0, 318, 320, 3, 54, 
	27, 0, 319, 321, 3, 52, 26, 0, 320, 319, 1, 0, 0, 0, 321, 322, 1, 0, 0, 
	0, 322, 320, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 
	336, 5, 73, 0, 0, 325, 330, 3, 34, 17, 0, 326, 327, 5, 80, 0, 0, 327, 329, 
	3, 34, 17, 0, 328, 326, 1, 0, 0, 0, 329, 332, 1, 0, 0, 0, 330, 328, 1, 
	0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 334, 1, 0, 0, 0, 332, 330, 1, 0, 0, 
	0, 333, 335, 5, 80, 0, 0, 334, 333, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 
	337, 1, 0, 0, 0, 336, 325, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 338, 
	1, 0, 0, 0, 338, 339, 5, 74, 0, 0, 339, 341, 1, 0, 0, 0, 340, 302, 1, 0, 
	0, 0, 340, 318, 1, 0, 0, 0, 341, 33, 1, 0, 0, 0, 342, 343, 6, 17, -1, 0, 
	343, 344, 3, 36, 18, 0, 344, 365, 1, 0, 0, 0, 345, 346, 10, 5, 0, 0, 346, 
	347, 3, 60, 30, 0, 347, 348, 3, 34, 17, 6, 348, 364, 1, 0, 0, 0, 349, 350, 
	10, 4, 0, 0, 350, 351, 3, 58, 29, 0, 351, 352, 3, 34, 17, 5, 352, 364, 
	1, 0, 0, 0, 353, 354, 10, 3, 0, 0, 354, 355, 3, 56, 28, 0, 355, 356, 3, 
	34, 17, 4, 356, 364, 1, 0, 0, 0, 357, 358, 10, 2, 0, 0, 358, 359, 5, 11, 
	0, 0, 359, 364, 3, 34, 17, 2, 360, 361, 10, 1, 0, 0, 361, 362, 5, 12, 0, 
	0, 362, 364, 3, 34, 17, 1, 363, 345, 1, 0, 0, 0, 363, 349, 1, 0, 0, 0, 
	363, 353, 1, 0, 0, 0, 363, 357, 1, 0, 0, 0, 363, 360, 1, 0, 0, 0, 364, 
	367, 1, 0, 0, 0, 365, 363, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 35, 1, 
	0, 0, 0, 367, 365, 1, 0, 0, 0, 368, 369, 5, 62, 0, 0, 369, 391, 3, 36, 
	18, 0, 370, 371, 5, 13, 0, 0, 371, 391, 3, 36, 18, 0, 372, 373, 5, 63, 
	0, 0, 373, 391, 3, 38, 19, 0, 374, 391, 5, 81, 0, 0, 375, 391, 5, 82, 0, 
	0, 376, 391, 5, 83, 0, 0, 377, 391, 5, 8, 0, 0, 378, 391, 5, 9, 0, 0, 379, 
	391, 5, 10, 0, 0, 380, 391, 3, 18, 9, 0, 381, 391, 3, 32, 16, 0, 382, 391, 
	3, 48, 24, 0, 383, 391, 3, 40, 20, 0, 384, 391, 3, 44, 22, 0, 385, 386, 
	5, 73, 0, 0, 386, 387, 3, 34, 17, 0, 387, 388, 5, 74, 0, 0, 388, 391, 1, 
	0, 0, 0, 389, 391, 3, 50, 25, 0, 390, 368, 1, 0, 0, 0, 390, 370, 1, 0, 
	0, 0, 390, 372, 1, 0, 0, 0, 390, 374, 1, 0, 0, 0, 390, 375, 1, 0, 0, 0, 
	390, 376, 1, 0, 0, 0, 390, 377, 1, 0, 0, 0, 390, 378, 1, 0, 0, 0, 390, 
	379, 1, 0, 0, 0, 390, 380, 1, 0, 0, 0, 390, 381, 1, 0, 0, 0, 390, 382, 
	1, 0, 0, 0, 390, 383, 1, 0, 0, 0, 390, 384, 1, 0, 0, 0, 390, 385, 1, 0, 
	0, 0, 390, 389, 1, 0, 0, 0, 391, 37, 1, 0, 0, 0, 392, 396, 3, 50, 25, 0, 
	393, 394, 5, 63, 0, 0, 394, 396, 3, 38, 19, 0, 395, 392, 1, 0, 0, 0, 395, 
	393, 1, 0, 0, 0, 396, 39, 1, 0, 0, 0, 397, 409, 5, 72, 0, 0, 398, 403, 
	3, 34, 17, 0, 399, 400, 5, 80, 0, 0, 400, 402, 3, 34, 17, 0, 401, 399, 
	1, 0, 0, 0, 402, 405, 1, 0, 0, 0, 403, 401, 1, 0, 0, 0, 403, 404, 1, 0, 
	0, 0, 404, 407, 1, 0, 0, 0, 405, 403, 1, 0, 0, 0, 406, 408, 5, 80, 0, 0, 
	407, 406, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408, 410, 1, 0, 0, 0, 409, 
	398, 1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 412, 
	5, 77, 0, 0, 412, 41, 1, 0, 0, 0, 413, 415, 3, 34, 17, 0, 414, 413, 1, 
	0, 0, 0, 414, 415, 1, 0, 0, 0, 415, 416, 1, 0, 0, 0, 416, 418, 5, 78, 0, 
	0, 417, 419, 3, 34, 17, 0, 418, 417, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 
	419, 43, 1, 0, 0, 0, 420, 432, 5, 75, 0, 0, 421, 426, 3, 46, 23, 0, 422, 
	423, 5, 80, 0, 0, 423, 425, 3, 46, 23, 0, 424, 422, 1, 0, 0, 0, 425, 428, 
	1, 0, 0, 0, 426, 424, 1, 0, 0, 0, 426, 427, 1, 0, 0, 0, 427, 430, 1, 0, 
	0, 0, 428, 426, 1, 0, 0, 0, 429, 431, 5, 80, 0, 0, 430, 429, 1, 0, 0, 0, 
	430, 431, 1, 0, 0, 0, 431, 433, 1, 0, 0, 0, 432, 421, 1, 0, 0, 0, 432, 
	433, 1, 0, 0, 0, 433, 434, 1, 0, 0, 0, 434, 435, 5, 76, 0, 0, 435, 45, 
	1, 0, 0, 0, 436, 437, 7, 5, 0, 0, 437, 438, 5, 78, 0, 0, 438, 444, 3, 34, 
	17, 0, 439, 440, 3, 38, 19, 0, 440, 441, 5, 78, 0, 0, 441, 442, 3, 34, 
	17, 0, 442, 444, 1, 0, 0, 0, 443, 436, 1, 0, 0, 0, 443, 439, 1, 0, 0, 0, 
	444, 47, 1, 0, 0, 0, 445, 446, 5, 22, 0, 0, 446, 447, 5, 88, 0, 0, 447, 
	463, 5, 75, 0, 0, 448, 449, 5, 88, 0, 0, 449, 450, 5, 78, 0, 0, 450, 457, 
	3, 34, 17, 0, 451, 452, 5, 80, 0, 0, 452, 453, 5, 88, 0, 0, 453, 454, 5, 
	78, 0, 0, 454, 456, 3, 34, 17, 0, 455, 451, 1, 0, 0, 0, 456, 459, 1, 0, 
	0, 0, 457, 455, 1, 0, 0, 0, 457, 458, 1, 0, 0, 0, 458, 461, 1, 0, 0, 0, 
	459, 457, 1, 0, 0, 0, 460, 462, 5, 80, 0, 0, 461, 460, 1, 0, 0, 0, 461, 
	462, 1, 0, 0, 0, 462, 464, 1, 0, 0, 0, 463, 448, 1, 0, 0, 0, 463, 464, 
	1, 0, 0, 0, 464, 465, 1, 0, 0, 0, 465, 466, 5, 76, 0, 0, 466, 49, 1, 0, 
	0, 0, 467, 471, 3, 54, 27, 0, 468, 470, 3, 52, 26, 0, 469, 468, 1, 0, 0, 
	0, 470, 473, 1, 0, 0, 0, 471, 469, 1, 0, 0, 0, 471, 472, 1, 0, 0, 0, 472, 
	51, 1, 0, 0, 0, 473, 471, 1, 0, 0, 0, 474, 475, 5, 71, 0, 0, 475, 499, 
	5, 88, 0, 0, 476, 479, 5, 72, 0, 0, 477, 480, 3, 34, 17, 0, 478, 480, 3, 
	42, 21, 0, 479, 477, 1, 0, 0, 0, 479, 478, 1, 0, 0, 0, 480, 481, 1, 0, 
	0, 0, 481, 482, 5, 77, 0, 0, 482, 499, 1, 0, 0, 0, 483, 495, 5, 73, 0, 
	0, 484, 489, 3, 34, 17, 0, 485, 486, 5, 80, 0, 0, 486, 488, 3, 34, 17, 
	0, 487, 485, 1, 0, 0, 0, 488, 491, 1, 0, 0, 0, 489, 487, 1, 0, 0, 0, 489, 
	490, 1, 0, 0, 0, 490, 493, 1, 0, 0, 0, 491, 489, 1, 0, 0, 0, 492, 494, 
	5, 80, 0, 0, 493, 492, 1, 0, 0, 0, 493, 494, 1, 0, 0, 0, 494, 496, 1, 0, 
	0, 0, 495, 484, 1, 0, 0, 0, 495, 496, 1, 0, 0, 0, 496, 497, 1, 0, 0, 0, 
	497, 499, 5, 74, 0, 0, 498, 474, 1, 0, 0, 0, 498, 476, 1, 0, 0, 0, 498, 
	483, 1, 0, 0, 0, 499, 53, 1, 0, 0, 0, 500, 501, 7, 6, 0, 0, 501, 55, 1, 
	0, 0, 0, 502, 503, 7, 7, 0, 0, 503, 57, 1, 0, 0, 0, 504, 505, 7, 8, 0, 
	0, 505, 59, 1, 0, 0, 0, 506, 507, 7, 9, 0, 0, 507, 61, 1, 0, 0, 0, 58, 
	65, 67, 80, 93, 96, 106, 109, 113, 129, 132, 137, 141, 145, 169, 172, 178, 
	185, 187, 194, 201, 210, 226, 242, 249, 267, 279, 285, 298, 309, 313, 315, 
	322, 330, 334, 336, 340, 363, 365, 390, 395, 403, 407, 409, 414, 418, 426, 
	430, 432, 443, 457, 461, 463, 471, 479, 489, 493, 495, 498,
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
	GsParserINITREF = 36
	GsParserNEWFROMTYPE = 37
	GsParserUINT8 = 38
	GsParserUINT16 = 39
	GsParserUINT32 = 40
	GsParserUINT64 = 41
	GsParserUINT = 42
	GsParserINT8 = 43
	GsParserINT16 = 44
	GsParserINT32 = 45
	GsParserINT64 = 46
	GsParserINTS = 47
	GsParserFLOAT32 = 48
	GsParserFLOAT64 = 49
	GsParserSTRINGS = 50
	GsParserBOOL = 51
	GsParserEXPAND = 52
	GsParserLSHIFT = 53
	GsParserRSHIFT = 54
	GsParserINCR = 55
	GsParserDECR = 56
	GsParserGEQ = 57
	GsParserLEQ = 58
	GsParserNEQ = 59
	GsParserEQ = 60
	GsParserADD = 61
	GsParserSUB = 62
	GsParserMUL = 63
	GsParserDIV = 64
	GsParserMOD = 65
	GsParserLT = 66
	GsParserGT = 67
	GsParserBITAND = 68
	GsParserBITOR = 69
	GsParserXOR = 70
	GsParserDOT = 71
	GsParserLBRACK = 72
	GsParserLPAREN = 73
	GsParserRPAREN = 74
	GsParserLBRACE = 75
	GsParserRBRACE = 76
	GsParserRBRACK = 77
	GsParserCOLON = 78
	GsParserSEMICOLON = 79
	GsParserCOMMA = 80
	GsParserINT = 81
	GsParserFLOAT = 82
	GsParserSTRING = 83
	GsParserWS = 84
	GsParserNEWLINE = 85
	GsParserSL_COMMENT = 86
	GsParserML_COMMENT = 87
	GsParserID = 88
)

// GsParser rules.
const (
	GsParserRULE_program = 0
	GsParserRULE_structDefinition = 1
	GsParserRULE_functionDefinition = 2
	GsParserRULE_block = 3
	GsParserRULE_statement = 4
	GsParserRULE_ifStatement = 5
	GsParserRULE_simpleStmt = 6
	GsParserRULE_assign = 7
	GsParserRULE_incrDecr = 8
	GsParserRULE_builtinCall = 9
	GsParserRULE_iterVar = 10
	GsParserRULE_forInit = 11
	GsParserRULE_forUpdate = 12
	GsParserRULE_selfAssign = 13
	GsParserRULE_updateItem = 14
	GsParserRULE_selfAssignOp = 15
	GsParserRULE_call = 16
	GsParserRULE_expr = 17
	GsParserRULE_atom = 18
	GsParserRULE_lvalue = 19
	GsParserRULE_arrayLiteral = 20
	GsParserRULE_sliceExpr = 21
	GsParserRULE_dictLiteral = 22
	GsParserRULE_dictEntry = 23
	GsParserRULE_instance = 24
	GsParserRULE_qid = 25
	GsParserRULE_accessor = 26
	GsParserRULE_primary = 27
	GsParserRULE_compOp = 28
	GsParserRULE_addOp = 29
	GsParserRULE_mulOp = 30
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
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & -9218868440455102336) != 0) || _la == GsParserSEMICOLON || _la == GsParserID {
		p.SetState(65)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(62)
				p.FunctionDefinition()
			}


		case 2:
			{
				p.SetState(63)
				p.StructDefinition()
			}


		case 3:
			{
				p.SetState(64)
				p.Statement()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(69)
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
		p.SetState(71)
		p.Match(GsParserTYPE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(72)
		p.Match(GsParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Match(GsParserSTRUCT)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(74)
		p.Match(GsParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(75)
		p.Match(GsParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == GsParserCOMMA {
		{
			p.SetState(76)
			p.Match(GsParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(GsParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(83)
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
		p.SetState(85)
		p.Match(GsParserFUNC)
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
	{
		p.SetState(87)
		p.Match(GsParserLPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == GsParserID {
		{
			p.SetState(88)
			p.Match(GsParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == GsParserCOMMA {
			{
				p.SetState(89)
				p.Match(GsParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(90)
				p.Match(GsParserID)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(95)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(98)
		p.Match(GsParserRPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(99)
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
		p.SetState(101)
		p.Match(GsParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & -9218868440455626624) != 0) || _la == GsParserSEMICOLON || _la == GsParserID {
		{
			p.SetState(102)
			p.Statement()
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			p.SetState(104)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
						{
							p.SetState(103)
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

				p.SetState(106)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}

			} else if p.HasError() { // JIM
				goto errorExit} else if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 2 {
			{
				p.SetState(108)
				p.Match(GsParserEOF)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

			} else if p.HasError() { // JIM
				goto errorExit
		}


		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(116)
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

func (s *IfStmtContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
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

	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewEmptyStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(118)
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
			p.SetState(119)
			p.StructDefinition()
		}


	case 3:
		localctx = NewAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(120)
			p.Assign()
		}


	case 4:
		localctx = NewSelfOpAssignStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(121)
			p.SelfAssign()
		}


	case 5:
		localctx = NewIncrDecrStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(122)
			p.IncrDecr()
		}


	case 6:
		localctx = NewReturnStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(123)
			p.Match(GsParserRETURN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(124)
				p.expr(0)
			}
			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)


			for _la == GsParserCOMMA {
				{
					p.SetState(125)
					p.Match(GsParserCOMMA)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(126)
					p.expr(0)
				}


				p.SetState(131)
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
			p.SetState(134)
			p.IfStatement()
		}


	case 8:
		localctx = NewForCStyleStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(135)
			p.Match(GsParserFOR)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == GsParserENV || _la == GsParserMUL || _la == GsParserID {
			{
				p.SetState(136)
				p.ForInit()
			}

		}
		{
			p.SetState(139)
			p.Match(GsParserSEMICOLON)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & -4607182422084147328) != 0) || ((int64((_la - 72)) & ^0x3f) == 0 && ((int64(1) << (_la - 72)) & 69131) != 0) {
			{
				p.SetState(140)
				p.expr(0)
			}

		}
		{
			p.SetState(143)
			p.Match(GsParserSEMICOLON)
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


		if _la == GsParserENV || _la == GsParserMUL || _la == GsParserID {
			{
				p.SetState(144)
				p.ForUpdate()
			}

		}
		{
			p.SetState(147)
			p.Block()
		}


	case 9:
		localctx = NewForRangeStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(148)
			p.Match(GsParserFOR)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(149)
			p.IterVar()
		}
		{
			p.SetState(150)
			p.Match(GsParserT__0)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Match(GsParserRANGE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(152)
			p.expr(0)
		}
		{
			p.SetState(153)
			p.Block()
		}


	case 10:
		localctx = NewForCondStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
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
			p.expr(0)
		}
		{
			p.SetState(157)
			p.Block()
		}


	case 11:
		localctx = NewBuiltinStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(159)
			p.BuiltinCall()
		}


	case 12:
		localctx = NewCallStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(160)
			p.Call()
		}


	case 13:
		localctx = NewBreakStmtContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(161)
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
			p.SetState(162)
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
			p.SetState(163)
			p.Match(GsParserGLOBAL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(164)
			p.Match(GsParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == GsParserCOMMA {
			{
				p.SetState(165)
				p.Match(GsParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(166)
				p.Match(GsParserID)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}


			p.SetState(171)
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


// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expr() IExprContext
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	SimpleStmt() ISimpleStmtContext
	SEMICOLON() antlr.TerminalNode
	ELSE() antlr.TerminalNode
	IfStatement() IIfStatementContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(GsParserIF, 0)
}

func (s *IfStatementContext) Expr() IExprContext {
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

func (s *IfStatementContext) AllBlock() []IBlockContext {
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

func (s *IfStatementContext) Block(i int) IBlockContext {
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

func (s *IfStatementContext) SimpleStmt() ISimpleStmtContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStmtContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *IfStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(GsParserSEMICOLON, 0)
}

func (s *IfStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GsParserELSE, 0)
}

func (s *IfStatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *GsParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GsParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(GsParserIF)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)


	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(175)
			p.SimpleStmt()
		}
		{
			p.SetState(176)
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
		p.SetState(180)
		p.expr(0)
	}
	{
		p.SetState(181)
		p.Block()
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == GsParserELSE {
		{
			p.SetState(182)
			p.Match(GsParserELSE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case GsParserLBRACE:
			{
				p.SetState(183)
				p.Block()
			}


		case GsParserIF:
			{
				p.SetState(184)
				p.IfStatement()
			}



		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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


// ISimpleStmtContext is an interface to support dynamic dispatch.
type ISimpleStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assign() IAssignContext
	SelfAssign() ISelfAssignContext
	IncrDecr() IIncrDecrContext
	BuiltinCall() IBuiltinCallContext
	Call() ICallContext

	// IsSimpleStmtContext differentiates from other interfaces.
	IsSimpleStmtContext()
}

type SimpleStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStmtContext() *SimpleStmtContext {
	var p = new(SimpleStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_simpleStmt
	return p
}

func InitEmptySimpleStmtContext(p *SimpleStmtContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GsParserRULE_simpleStmt
}

func (*SimpleStmtContext) IsSimpleStmtContext() {}

func NewSimpleStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStmtContext {
	var p = new(SimpleStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GsParserRULE_simpleStmt

	return p
}

func (s *SimpleStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStmtContext) Assign() IAssignContext {
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

func (s *SimpleStmtContext) SelfAssign() ISelfAssignContext {
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

func (s *SimpleStmtContext) IncrDecr() IIncrDecrContext {
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

func (s *SimpleStmtContext) BuiltinCall() IBuiltinCallContext {
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

func (s *SimpleStmtContext) Call() ICallContext {
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

func (s *SimpleStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SimpleStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitSimpleStmt(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *GsParser) SimpleStmt() (localctx ISimpleStmtContext) {
	localctx = NewSimpleStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GsParserRULE_simpleStmt)
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Assign()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.SelfAssign()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(191)
			p.IncrDecr()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(192)
			p.BuiltinCall()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(193)
			p.Call()
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
	p.EnterRule(localctx, 14, GsParserRULE_assign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Lvalue()
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == GsParserCOMMA {
		{
			p.SetState(197)
			p.Match(GsParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Lvalue()
		}


		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(204)
		p.Match(GsParserT__0)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(205)
		p.expr(0)
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == GsParserCOMMA {
		{
			p.SetState(206)
			p.Match(GsParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(207)
			p.expr(0)
		}


		p.SetState(212)
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
	p.EnterRule(localctx, 16, GsParserRULE_incrDecr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		p.Lvalue()
	}
	{
		p.SetState(214)
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

func (s *GoCallContext) Call() ICallContext {
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


func (s *GoCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitGoCall(s)

	default:
		return t.VisitChildren(s)
	}
}


type InitRefCallContext struct {
	BuiltinCallContext
}

func NewInitRefCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InitRefCallContext {
	var p = new(InitRefCallContext)

	InitEmptyBuiltinCallContext(&p.BuiltinCallContext)
	p.parser = parser
	p.CopyAll(ctx.(*BuiltinCallContext))

	return p
}

func (s *InitRefCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitRefCallContext) INITREF() antlr.TerminalNode {
	return s.GetToken(GsParserINITREF, 0)
}

func (s *InitRefCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *InitRefCallContext) Qid() IQidContext {
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

func (s *InitRefCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}

func (s *InitRefCallContext) COMMA() antlr.TerminalNode {
	return s.GetToken(GsParserCOMMA, 0)
}

func (s *InitRefCallContext) Expr() IExprContext {
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


func (s *InitRefCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitInitRefCall(s)

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


type NewFromTypeCallContext struct {
	BuiltinCallContext
}

func NewNewFromTypeCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewFromTypeCallContext {
	var p = new(NewFromTypeCallContext)

	InitEmptyBuiltinCallContext(&p.BuiltinCallContext)
	p.parser = parser
	p.CopyAll(ctx.(*BuiltinCallContext))

	return p
}

func (s *NewFromTypeCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewFromTypeCallContext) NEWFROMTYPE() antlr.TerminalNode {
	return s.GetToken(GsParserNEWFROMTYPE, 0)
}

func (s *NewFromTypeCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserLPAREN, 0)
}

func (s *NewFromTypeCallContext) Expr() IExprContext {
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

func (s *NewFromTypeCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(GsParserRPAREN, 0)
}


func (s *NewFromTypeCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GsVisitor:
		return t.VisitNewFromTypeCall(s)

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
	p.EnterRule(localctx, 18, GsParserRULE_builtinCall)
	var _la int

	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GsParserLEN:
		localctx = NewLenCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(216)
			p.Match(GsParserLEN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(217)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(218)
			p.expr(0)
		}
		{
			p.SetState(219)
			p.Match(GsParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case GsParserINITREF:
		localctx = NewInitRefCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(221)
			p.Match(GsParserINITREF)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(222)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Qid()
		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == GsParserCOMMA {
			{
				p.SetState(224)
				p.Match(GsParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(225)
				p.expr(0)
			}

		}
		{
			p.SetState(228)
			p.Match(GsParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case GsParserNEWFROMTYPE:
		localctx = NewNewFromTypeCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(230)
			p.Match(GsParserNEWFROMTYPE)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(232)
			p.expr(0)
		}
		{
			p.SetState(233)
			p.Match(GsParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case GsParserAPPEND:
		localctx = NewAppendCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(235)
			p.Match(GsParserAPPEND)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(236)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(237)
			p.expr(0)
		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
		case 1:
			p.SetState(242)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)


			for _la == GsParserCOMMA {
				{
					p.SetState(238)
					p.Match(GsParserCOMMA)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(239)
					p.expr(0)
				}


				p.SetState(244)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
			    	goto errorExit
			    }
				_la = p.GetTokenStream().LA(1)
			}


		case 2:
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
				p.expr(0)
			}
			{
				p.SetState(247)
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
			p.SetState(251)
			p.Match(GsParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case GsParserDELETE:
		localctx = NewDeleteCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(253)
			p.Match(GsParserDELETE)
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
			p.Match(GsParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(257)
			p.expr(0)
		}
		{
			p.SetState(258)
			p.Match(GsParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case GsParserPRINT, GsParserPRINTF, GsParserSPRINTF, GsParserPRINTLN:
		localctx = NewPrintXCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(260)
			_la = p.GetTokenStream().LA(1)

			if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 64424509440) != 0)) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(261)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(262)
			p.expr(0)
		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		for _la == GsParserCOMMA {
			{
				p.SetState(263)
				p.Match(GsParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}
			{
				p.SetState(264)
				p.expr(0)
			}


			p.SetState(269)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(270)
			p.Match(GsParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case GsParserUINT8, GsParserUINT16, GsParserUINT32, GsParserUINT64, GsParserUINT, GsParserINT8, GsParserINT16, GsParserINT32, GsParserINT64, GsParserINTS, GsParserFLOAT32, GsParserFLOAT64, GsParserSTRINGS, GsParserBOOL:
		localctx = NewConvertCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(272)
			_la = p.GetTokenStream().LA(1)

			if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 4503324749463552) != 0)) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(273)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(274)
			p.expr(0)
		}
		{
			p.SetState(275)
			p.Match(GsParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}


	case GsParserGO:
		localctx = NewGoCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(277)
			p.Match(GsParserGO)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(278)
			p.Call()
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
	p.EnterRule(localctx, 20, GsParserRULE_iterVar)
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSingleIterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)
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
			p.SetState(282)
			p.Match(GsParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(283)
			p.Match(GsParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(284)
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
	p.EnterRule(localctx, 22, GsParserRULE_forInit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
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
	p.EnterRule(localctx, 24, GsParserRULE_forUpdate)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
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
	p.EnterRule(localctx, 26, GsParserRULE_selfAssign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Lvalue()
	}
	{
		p.SetState(292)
		p.SelfAssignOp()
	}
	{
		p.SetState(293)
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
	p.EnterRule(localctx, 28, GsParserRULE_updateItem)
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSelfUpdateContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(295)
			p.SelfAssign()
		}


	case 2:
		localctx = NewIncrDecrUpdateContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(296)
			p.IncrDecr()
		}


	case 3:
		localctx = NewAssignUpdateContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(297)
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
	p.EnterRule(localctx, 30, GsParserRULE_selfAssignOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
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
	p.EnterRule(localctx, 32, GsParserRULE_call)
	var _la int

	var _alt int

	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		localctx = NewInnerCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(302)
			p.Match(GsParserID)
			if p.HasError() {
					// Recognition error - abort rule
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


		if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & -4607182422084147328) != 0) || ((int64((_la - 72)) & ^0x3f) == 0 && ((int64(1) << (_la - 72)) & 69131) != 0) {
			{
				p.SetState(304)
				p.expr(0)
			}
			p.SetState(309)
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
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
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


	case 2:
		localctx = NewOuterCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(318)
			p.Primary()
		}
		p.SetState(320)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
					{
						p.SetState(319)
						p.Accessor()
					}




			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(322)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(324)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & -4607182422084147328) != 0) || ((int64((_la - 72)) & ^0x3f) == 0 && ((int64(1) << (_la - 72)) & 69131) != 0) {
			{
				p.SetState(325)
				p.expr(0)
			}
			p.SetState(330)
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
						p.SetState(326)
						p.Match(GsParserCOMMA)
						if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
						}
					}
					{
						p.SetState(327)
						p.expr(0)
					}


				}
				p.SetState(332)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
			    	goto errorExit
			    }
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}
			p.SetState(334)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)


			if _la == GsParserCOMMA {
				{
					p.SetState(333)
					p.Match(GsParserCOMMA)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}

			}

		}
		{
			p.SetState(338)
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
	_startState := 34
	p.EnterRecursionRule(localctx, 34, GsParserRULE_expr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewAtomExprContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(343)
		p.Atom()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(365)
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
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(363)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GsParserRULE_expr)
				p.SetState(345)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(346)
					p.MulOp()
				}
				{
					p.SetState(347)
					p.expr(6)
				}


			case 2:
				localctx = NewAddExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GsParserRULE_expr)
				p.SetState(349)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(350)
					p.AddOp()
				}
				{
					p.SetState(351)
					p.expr(5)
				}


			case 3:
				localctx = NewComparisonExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GsParserRULE_expr)
				p.SetState(353)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(354)
					p.CompOp()
				}
				{
					p.SetState(355)
					p.expr(4)
				}


			case 4:
				localctx = NewLogicalAndExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GsParserRULE_expr)
				p.SetState(357)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(358)
					p.Match(GsParserAND)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(359)
					p.expr(2)
				}


			case 5:
				localctx = NewLogicalOrExprContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GsParserRULE_expr)
				p.SetState(360)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(361)
					p.Match(GsParserOR)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(362)
					p.expr(1)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(367)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 36, GsParserRULE_atom)
	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNegAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(368)
			p.Match(GsParserSUB)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(369)
			p.Atom()
		}


	case 2:
		localctx = NewNotAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(370)
			p.Match(GsParserNOT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(371)
			p.Atom()
		}


	case 3:
		localctx = NewDerefAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
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


	case 4:
		localctx = NewIntAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(374)
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
			p.SetState(375)
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
			p.SetState(376)
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
			p.SetState(377)
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
			p.SetState(378)
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
			p.SetState(379)
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
			p.SetState(380)
			p.BuiltinCall()
		}


	case 11:
		localctx = NewCallAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(381)
			p.Call()
		}


	case 12:
		localctx = NewInstanceAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(382)
			p.Instance()
		}


	case 13:
		localctx = NewArrayAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(383)
			p.ArrayLiteral()
		}


	case 14:
		localctx = NewDictAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(384)
			p.DictLiteral()
		}


	case 15:
		localctx = NewParenAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(385)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(386)
			p.expr(0)
		}
		{
			p.SetState(387)
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
			p.SetState(389)
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
	p.EnterRule(localctx, 38, GsParserRULE_lvalue)
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GsParserENV, GsParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(392)
			p.Qid()
		}


	case GsParserMUL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(393)
			p.Match(GsParserMUL)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(394)
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
	p.EnterRule(localctx, 40, GsParserRULE_arrayLiteral)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		p.Match(GsParserLBRACK)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & -4607182422084147328) != 0) || ((int64((_la - 72)) & ^0x3f) == 0 && ((int64(1) << (_la - 72)) & 69131) != 0) {
		{
			p.SetState(398)
			p.expr(0)
		}
		p.SetState(403)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(399)
					p.Match(GsParserCOMMA)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(400)
					p.expr(0)
				}


			}
			p.SetState(405)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(407)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == GsParserCOMMA {
			{
				p.SetState(406)
				p.Match(GsParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}

	}
	{
		p.SetState(411)
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
	p.EnterRule(localctx, 42, GsParserRULE_sliceExpr)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & -4607182422084147328) != 0) || ((int64((_la - 72)) & ^0x3f) == 0 && ((int64(1) << (_la - 72)) & 69131) != 0) {
		{
			p.SetState(413)
			p.expr(0)
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
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & -4607182422084147328) != 0) || ((int64((_la - 72)) & ^0x3f) == 0 && ((int64(1) << (_la - 72)) & 69131) != 0) {
		{
			p.SetState(417)
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
	p.EnterRule(localctx, 44, GsParserRULE_dictLiteral)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(420)
		p.Match(GsParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(432)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & -9223372036854774912) != 0) || ((int64((_la - 81)) & ^0x3f) == 0 && ((int64(1) << (_la - 81)) & 135) != 0) {
		{
			p.SetState(421)
			p.DictEntry()
		}
		p.SetState(426)
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
					p.SetState(422)
					p.Match(GsParserCOMMA)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(423)
					p.DictEntry()
				}


			}
			p.SetState(428)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == GsParserCOMMA {
			{
				p.SetState(429)
				p.Match(GsParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}

	}
	{
		p.SetState(434)
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
	p.EnterRule(localctx, 46, GsParserRULE_dictEntry)
	var _la int

	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GsParserTRUE, GsParserFALSE, GsParserINT, GsParserFLOAT, GsParserSTRING:
		localctx = NewConstKeyEntryContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(436)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GsParserTRUE || _la == GsParserFALSE || ((int64((_la - 81)) & ^0x3f) == 0 && ((int64(1) << (_la - 81)) & 7) != 0)) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(437)
			p.Match(GsParserCOLON)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(438)
			p.expr(0)
		}


	case GsParserENV, GsParserMUL, GsParserID:
		localctx = NewIdKeyEntryContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(439)
			p.Lvalue()
		}
		{
			p.SetState(440)
			p.Match(GsParserCOLON)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(441)
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
	p.EnterRule(localctx, 48, GsParserRULE_instance)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(445)
		p.Match(GsParserNEW)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(446)
		p.Match(GsParserID)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(447)
		p.Match(GsParserLBRACE)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if _la == GsParserID {
		{
			p.SetState(448)
			p.Match(GsParserID)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(449)
			p.Match(GsParserCOLON)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(450)
			p.expr(0)
		}
		p.SetState(457)
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
					p.SetState(451)
					p.Match(GsParserCOMMA)
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
					p.Match(GsParserCOLON)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(454)
					p.expr(0)
				}


			}
			p.SetState(459)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
		    	goto errorExit
		    }
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 49, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(461)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if _la == GsParserCOMMA {
			{
				p.SetState(460)
				p.Match(GsParserCOMMA)
				if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
				}
			}

		}

	}
	{
		p.SetState(465)
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
	p.EnterRule(localctx, 50, GsParserRULE_qid)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(467)
		p.Primary()
	}
	p.SetState(471)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(468)
				p.Accessor()
			}


		}
		p.SetState(473)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 52, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 52, GsParserRULE_accessor)
	var _la int

	var _alt int

	p.SetState(498)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GsParserDOT:
		localctx = NewPropertyAccessContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(474)
			p.Match(GsParserDOT)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(475)
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
			p.SetState(476)
			p.Match(GsParserLBRACK)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(479)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 53, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(477)
				p.expr(0)
			}


		case 2:
			{
				p.SetState(478)
				p.SliceExpr()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(481)
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
			p.SetState(483)
			p.Match(GsParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		p.SetState(495)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)


		if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & -4607182422084147328) != 0) || ((int64((_la - 72)) & ^0x3f) == 0 && ((int64(1) << (_la - 72)) & 69131) != 0) {
			{
				p.SetState(484)
				p.expr(0)
			}
			p.SetState(489)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(485)
						p.Match(GsParserCOMMA)
						if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
						}
					}
					{
						p.SetState(486)
						p.expr(0)
					}


				}
				p.SetState(491)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
			    	goto errorExit
			    }
				_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 54, p.GetParserRuleContext())
				if p.HasError() {
					goto errorExit
				}
			}
			p.SetState(493)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)


			if _la == GsParserCOMMA {
				{
					p.SetState(492)
					p.Match(GsParserCOMMA)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}

			}

		}
		{
			p.SetState(497)
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
	p.EnterRule(localctx, 54, GsParserRULE_primary)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(500)
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
	p.EnterRule(localctx, 56, GsParserRULE_compOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(502)
		_la = p.GetTokenStream().LA(1)

		if !(((int64((_la - 57)) & ^0x3f) == 0 && ((int64(1) << (_la - 57)) & 1551) != 0)) {
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
	p.EnterRule(localctx, 58, GsParserRULE_addOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(504)
		_la = p.GetTokenStream().LA(1)

		if !(((int64((_la - 61)) & ^0x3f) == 0 && ((int64(1) << (_la - 61)) & 771) != 0)) {
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
	p.EnterRule(localctx, 60, GsParserRULE_mulOp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(506)
		_la = p.GetTokenStream().LA(1)

		if !(((int64((_la - 53)) & ^0x3f) == 0 && ((int64(1) << (_la - 53)) & 39939) != 0)) {
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
	case 17:
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

