// Code generated from gengine.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // gengine

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 51, 308,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 6, 2, 74, 10, 2, 13, 2, 14, 2,
	75, 3, 3, 3, 3, 3, 3, 5, 3, 81, 10, 3, 3, 3, 5, 3, 84, 10, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	8, 6, 8, 100, 10, 8, 13, 8, 14, 8, 101, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5,
	9, 109, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 116, 10, 10, 12,
	10, 14, 10, 119, 11, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 5, 11, 126,
	10, 11, 3, 11, 3, 11, 5, 11, 130, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5,
	11, 136, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	7, 11, 146, 10, 11, 12, 11, 14, 11, 149, 11, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 5, 12, 157, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 7, 12, 167, 10, 12, 12, 12, 14, 12, 170, 11, 12, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 177, 10, 13, 3, 14, 3, 14, 5, 14,
	181, 10, 14, 3, 14, 3, 14, 3, 14, 5, 14, 186, 10, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 5, 15, 192, 10, 15, 3, 15, 3, 15, 7, 15, 196, 10, 15, 12, 15,
	14, 15, 199, 11, 15, 3, 15, 5, 15, 202, 10, 15, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 5, 16, 209, 10, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 5, 17,
	216, 10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 5, 18, 227, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19,
	235, 10, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 244,
	10, 19, 7, 19, 246, 10, 19, 12, 19, 14, 19, 249, 11, 19, 3, 20, 5, 20,
	252, 10, 20, 3, 20, 3, 20, 3, 21, 5, 21, 257, 10, 21, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 5, 24, 268, 10, 24, 3, 24,
	3, 24, 3, 25, 3, 25, 3, 25, 5, 25, 275, 10, 25, 3, 25, 3, 25, 3, 26, 3,
	26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31,
	3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 298, 10, 33, 3,
	33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 2, 4, 20, 22,
	37, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 2,
	9, 3, 2, 14, 15, 4, 2, 20, 20, 48, 48, 3, 2, 22, 23, 3, 2, 24, 25, 3, 2,
	26, 31, 3, 2, 12, 13, 3, 2, 33, 38, 2, 326, 2, 73, 3, 2, 2, 2, 4, 77, 3,
	2, 2, 2, 6, 89, 3, 2, 2, 2, 8, 91, 3, 2, 2, 2, 10, 93, 3, 2, 2, 2, 12,
	96, 3, 2, 2, 2, 14, 99, 3, 2, 2, 2, 16, 108, 3, 2, 2, 2, 18, 110, 3, 2,
	2, 2, 20, 135, 3, 2, 2, 2, 22, 156, 3, 2, 2, 2, 24, 176, 3, 2, 2, 2, 26,
	180, 3, 2, 2, 2, 28, 187, 3, 2, 2, 2, 30, 203, 3, 2, 2, 2, 32, 212, 3,
	2, 2, 2, 34, 226, 3, 2, 2, 2, 36, 234, 3, 2, 2, 2, 38, 251, 3, 2, 2, 2,
	40, 256, 3, 2, 2, 2, 42, 260, 3, 2, 2, 2, 44, 262, 3, 2, 2, 2, 46, 264,
	3, 2, 2, 2, 48, 271, 3, 2, 2, 2, 50, 278, 3, 2, 2, 2, 52, 280, 3, 2, 2,
	2, 54, 282, 3, 2, 2, 2, 56, 284, 3, 2, 2, 2, 58, 286, 3, 2, 2, 2, 60, 288,
	3, 2, 2, 2, 62, 290, 3, 2, 2, 2, 64, 292, 3, 2, 2, 2, 66, 301, 3, 2, 2,
	2, 68, 303, 3, 2, 2, 2, 70, 305, 3, 2, 2, 2, 72, 74, 5, 4, 3, 2, 73, 72,
	3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2,
	76, 3, 3, 2, 2, 2, 77, 78, 7, 11, 2, 2, 78, 80, 5, 6, 4, 2, 79, 81, 5,
	8, 5, 2, 80, 79, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 83, 3, 2, 2, 2, 82,
	84, 5, 10, 6, 2, 83, 82, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 85, 3, 2,
	2, 2, 85, 86, 7, 18, 2, 2, 86, 87, 5, 12, 7, 2, 87, 88, 7, 19, 2, 2, 88,
	5, 3, 2, 2, 2, 89, 90, 5, 42, 22, 2, 90, 7, 3, 2, 2, 2, 91, 92, 5, 42,
	22, 2, 92, 9, 3, 2, 2, 2, 93, 94, 7, 17, 2, 2, 94, 95, 5, 38, 20, 2, 95,
	11, 3, 2, 2, 2, 96, 97, 5, 14, 8, 2, 97, 13, 3, 2, 2, 2, 98, 100, 5, 16,
	9, 2, 99, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101,
	102, 3, 2, 2, 2, 102, 15, 3, 2, 2, 2, 103, 109, 5, 28, 15, 2, 104, 109,
	5, 48, 25, 2, 105, 109, 5, 46, 24, 2, 106, 109, 5, 26, 14, 2, 107, 109,
	5, 18, 10, 2, 108, 103, 3, 2, 2, 2, 108, 104, 3, 2, 2, 2, 108, 105, 3,
	2, 2, 2, 108, 106, 3, 2, 2, 2, 108, 107, 3, 2, 2, 2, 109, 17, 3, 2, 2,
	2, 110, 111, 7, 3, 2, 2, 111, 117, 7, 42, 2, 2, 112, 116, 5, 48, 25, 2,
	113, 116, 5, 46, 24, 2, 114, 116, 5, 26, 14, 2, 115, 112, 3, 2, 2, 2, 115,
	113, 3, 2, 2, 2, 115, 114, 3, 2, 2, 2, 116, 119, 3, 2, 2, 2, 117, 115,
	3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 120, 3, 2, 2, 2, 119, 117, 3, 2,
	2, 2, 120, 121, 7, 43, 2, 2, 121, 19, 3, 2, 2, 2, 122, 123, 8, 11, 1, 2,
	123, 136, 5, 22, 12, 2, 124, 126, 5, 62, 32, 2, 125, 124, 3, 2, 2, 2, 125,
	126, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 136, 5, 24, 13, 2, 128, 130,
	5, 62, 32, 2, 129, 128, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 131, 3,
	2, 2, 2, 131, 132, 7, 44, 2, 2, 132, 133, 5, 20, 11, 2, 133, 134, 7, 45,
	2, 2, 134, 136, 3, 2, 2, 2, 135, 122, 3, 2, 2, 2, 135, 125, 3, 2, 2, 2,
	135, 129, 3, 2, 2, 2, 136, 147, 3, 2, 2, 2, 137, 138, 12, 6, 2, 2, 138,
	139, 5, 56, 29, 2, 139, 140, 5, 20, 11, 7, 140, 146, 3, 2, 2, 2, 141, 142,
	12, 5, 2, 2, 142, 143, 5, 58, 30, 2, 143, 144, 5, 20, 11, 6, 144, 146,
	3, 2, 2, 2, 145, 137, 3, 2, 2, 2, 145, 141, 3, 2, 2, 2, 146, 149, 3, 2,
	2, 2, 147, 145, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 21, 3, 2, 2, 2,
	149, 147, 3, 2, 2, 2, 150, 151, 8, 12, 1, 2, 151, 157, 5, 24, 13, 2, 152,
	153, 7, 44, 2, 2, 153, 154, 5, 22, 12, 2, 154, 155, 7, 45, 2, 2, 155, 157,
	3, 2, 2, 2, 156, 150, 3, 2, 2, 2, 156, 152, 3, 2, 2, 2, 157, 168, 3, 2,
	2, 2, 158, 159, 12, 6, 2, 2, 159, 160, 5, 54, 28, 2, 160, 161, 5, 22, 12,
	7, 161, 167, 3, 2, 2, 2, 162, 163, 12, 5, 2, 2, 163, 164, 5, 52, 27, 2,
	164, 165, 5, 22, 12, 6, 165, 167, 3, 2, 2, 2, 166, 158, 3, 2, 2, 2, 166,
	162, 3, 2, 2, 2, 167, 170, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 168, 169,
	3, 2, 2, 2, 169, 23, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 171, 177, 5, 48,
	25, 2, 172, 177, 5, 46, 24, 2, 173, 177, 5, 34, 18, 2, 174, 177, 5, 64,
	33, 2, 175, 177, 5, 50, 26, 2, 176, 171, 3, 2, 2, 2, 176, 172, 3, 2, 2,
	2, 176, 173, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 176, 175, 3, 2, 2, 2, 177,
	25, 3, 2, 2, 2, 178, 181, 5, 64, 33, 2, 179, 181, 5, 50, 26, 2, 180, 178,
	3, 2, 2, 2, 180, 179, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 185, 5, 60,
	31, 2, 183, 186, 5, 22, 12, 2, 184, 186, 5, 20, 11, 2, 185, 183, 3, 2,
	2, 2, 185, 184, 3, 2, 2, 2, 186, 27, 3, 2, 2, 2, 187, 188, 7, 4, 2, 2,
	188, 189, 5, 20, 11, 2, 189, 191, 7, 42, 2, 2, 190, 192, 5, 14, 8, 2, 191,
	190, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 197,
	7, 43, 2, 2, 194, 196, 5, 30, 16, 2, 195, 194, 3, 2, 2, 2, 196, 199, 3,
	2, 2, 2, 197, 195, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 201, 3, 2, 2,
	2, 199, 197, 3, 2, 2, 2, 200, 202, 5, 32, 17, 2, 201, 200, 3, 2, 2, 2,
	201, 202, 3, 2, 2, 2, 202, 29, 3, 2, 2, 2, 203, 204, 7, 5, 2, 2, 204, 205,
	7, 4, 2, 2, 205, 206, 5, 20, 11, 2, 206, 208, 7, 42, 2, 2, 207, 209, 5,
	14, 8, 2, 208, 207, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 210, 3, 2, 2,
	2, 210, 211, 7, 43, 2, 2, 211, 31, 3, 2, 2, 2, 212, 213, 7, 5, 2, 2, 213,
	215, 7, 42, 2, 2, 214, 216, 5, 14, 8, 2, 215, 214, 3, 2, 2, 2, 215, 216,
	3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 218, 7, 43, 2, 2, 218, 33, 3, 2,
	2, 2, 219, 227, 5, 44, 23, 2, 220, 227, 5, 38, 20, 2, 221, 227, 5, 40,
	21, 2, 222, 227, 5, 42, 22, 2, 223, 227, 5, 66, 34, 2, 224, 227, 5, 68,
	35, 2, 225, 227, 5, 70, 36, 2, 226, 219, 3, 2, 2, 2, 226, 220, 3, 2, 2,
	2, 226, 221, 3, 2, 2, 2, 226, 222, 3, 2, 2, 2, 226, 223, 3, 2, 2, 2, 226,
	224, 3, 2, 2, 2, 226, 225, 3, 2, 2, 2, 227, 35, 3, 2, 2, 2, 228, 235, 5,
	34, 18, 2, 229, 235, 5, 50, 26, 2, 230, 235, 5, 46, 24, 2, 231, 235, 5,
	48, 25, 2, 232, 235, 5, 64, 33, 2, 233, 235, 5, 20, 11, 2, 234, 228, 3,
	2, 2, 2, 234, 229, 3, 2, 2, 2, 234, 230, 3, 2, 2, 2, 234, 231, 3, 2, 2,
	2, 234, 232, 3, 2, 2, 2, 234, 233, 3, 2, 2, 2, 235, 247, 3, 2, 2, 2, 236,
	243, 7, 6, 2, 2, 237, 244, 5, 34, 18, 2, 238, 244, 5, 50, 26, 2, 239, 244,
	5, 46, 24, 2, 240, 244, 5, 48, 25, 2, 241, 244, 5, 64, 33, 2, 242, 244,
	5, 20, 11, 2, 243, 237, 3, 2, 2, 2, 243, 238, 3, 2, 2, 2, 243, 239, 3,
	2, 2, 2, 243, 240, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2, 243, 242, 3, 2, 2,
	2, 244, 246, 3, 2, 2, 2, 245, 236, 3, 2, 2, 2, 246, 249, 3, 2, 2, 2, 247,
	245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 37, 3, 2, 2, 2, 249, 247, 3,
	2, 2, 2, 250, 252, 7, 23, 2, 2, 251, 250, 3, 2, 2, 2, 251, 252, 3, 2, 2,
	2, 252, 253, 3, 2, 2, 2, 253, 254, 7, 21, 2, 2, 254, 39, 3, 2, 2, 2, 255,
	257, 7, 23, 2, 2, 256, 255, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 258,
	3, 2, 2, 2, 258, 259, 7, 49, 2, 2, 259, 41, 3, 2, 2, 2, 260, 261, 7, 47,
	2, 2, 261, 43, 3, 2, 2, 2, 262, 263, 9, 2, 2, 2, 263, 45, 3, 2, 2, 2, 264,
	265, 7, 20, 2, 2, 265, 267, 7, 44, 2, 2, 266, 268, 5, 36, 19, 2, 267, 266,
	3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 270, 7, 45,
	2, 2, 270, 47, 3, 2, 2, 2, 271, 272, 7, 48, 2, 2, 272, 274, 7, 44, 2, 2,
	273, 275, 5, 36, 19, 2, 274, 273, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275,
	276, 3, 2, 2, 2, 276, 277, 7, 45, 2, 2, 277, 49, 3, 2, 2, 2, 278, 279,
	9, 3, 2, 2, 279, 51, 3, 2, 2, 2, 280, 281, 9, 4, 2, 2, 281, 53, 3, 2, 2,
	2, 282, 283, 9, 5, 2, 2, 283, 55, 3, 2, 2, 2, 284, 285, 9, 6, 2, 2, 285,
	57, 3, 2, 2, 2, 286, 287, 9, 7, 2, 2, 287, 59, 3, 2, 2, 2, 288, 289, 9,
	8, 2, 2, 289, 61, 3, 2, 2, 2, 290, 291, 7, 32, 2, 2, 291, 63, 3, 2, 2,
	2, 292, 293, 5, 50, 26, 2, 293, 297, 7, 39, 2, 2, 294, 298, 5, 38, 20,
	2, 295, 298, 5, 42, 22, 2, 296, 298, 5, 50, 26, 2, 297, 294, 3, 2, 2, 2,
	297, 295, 3, 2, 2, 2, 297, 296, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299,
	300, 7, 40, 2, 2, 300, 65, 3, 2, 2, 2, 301, 302, 7, 7, 2, 2, 302, 67, 3,
	2, 2, 2, 303, 304, 7, 8, 2, 2, 304, 69, 3, 2, 2, 2, 305, 306, 7, 9, 2,
	2, 306, 71, 3, 2, 2, 2, 34, 75, 80, 83, 101, 108, 115, 117, 125, 129, 135,
	145, 147, 156, 166, 168, 176, 180, 185, 191, 197, 201, 208, 215, 226, 234,
	243, 247, 251, 256, 267, 274, 297,
}
var literalNames = []string{
	"", "'conc'", "'if'", "'else'", "','", "'@name'", "'@desc'", "'@id'", "",
	"", "'&&'", "'||'", "", "", "", "", "", "", "", "", "'+'", "'-'", "'/'",
	"'*'", "'=='", "'>'", "'<'", "'>='", "'<='", "'!='", "'!'", "':='", "'='",
	"'+='", "'-='", "'*='", "'/='", "'['", "']'", "';'", "'{'", "'}'", "'('",
	"')'", "'.'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "NIL", "RULE", "AND", "OR", "TRUE", "FALSE",
	"NULL_LITERAL", "SALIENCE", "BEGIN", "END", "SIMPLENAME", "INT", "PLUS",
	"MINUS", "DIV", "MUL", "EQUALS", "GT", "LT", "GTE", "LTE", "NOTEQUALS",
	"NOT", "ASSIGN", "SET", "PLUSEQUAL", "MINUSEQUAL", "MULTIEQUAL", "DIVEQUAL",
	"LSQARE", "RSQARE", "SEMICOLON", "LR_BRACE", "RR_BRACE", "LR_BRACKET",
	"RR_BRACKET", "DOT", "DQUOTA_STRING", "DOTTEDNAME", "REAL_LITERAL", "SL_COMMENT",
	"WS",
}

var ruleNames = []string{
	"primary", "ruleEntity", "ruleName", "ruleDescription", "salience", "ruleContent",
	"statements", "statement", "concStatement", "expression", "mathExpression",
	"expressionAtom", "assignment", "ifStmt", "elseIfStmt", "elseStmt", "constant",
	"functionArgs", "integer", "realLiteral", "stringLiteral", "booleanLiteral",
	"functionCall", "methodCall", "variable", "mathPmOperator", "mathMdOperator",
	"comparisonOperator", "logicalOperator", "assignOperator", "notOperator",
	"mapVar", "atName", "atDesc", "atId",
}

type gengineParser struct {
	*antlr.BaseParser
}

// NewgengineParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *gengineParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewgengineParser(input antlr.TokenStream) *gengineParser {
	this := new(gengineParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "gengine.g4"

	return this
}

// gengineParser tokens.
const (
	gengineParserEOF           = antlr.TokenEOF
	gengineParserT__0          = 1
	gengineParserT__1          = 2
	gengineParserT__2          = 3
	gengineParserT__3          = 4
	gengineParserT__4          = 5
	gengineParserT__5          = 6
	gengineParserT__6          = 7
	gengineParserNIL           = 8
	gengineParserRULE          = 9
	gengineParserAND           = 10
	gengineParserOR            = 11
	gengineParserTRUE          = 12
	gengineParserFALSE         = 13
	gengineParserNULL_LITERAL  = 14
	gengineParserSALIENCE      = 15
	gengineParserBEGIN         = 16
	gengineParserEND           = 17
	gengineParserSIMPLENAME    = 18
	gengineParserINT           = 19
	gengineParserPLUS          = 20
	gengineParserMINUS         = 21
	gengineParserDIV           = 22
	gengineParserMUL           = 23
	gengineParserEQUALS        = 24
	gengineParserGT            = 25
	gengineParserLT            = 26
	gengineParserGTE           = 27
	gengineParserLTE           = 28
	gengineParserNOTEQUALS     = 29
	gengineParserNOT           = 30
	gengineParserASSIGN        = 31
	gengineParserSET           = 32
	gengineParserPLUSEQUAL     = 33
	gengineParserMINUSEQUAL    = 34
	gengineParserMULTIEQUAL    = 35
	gengineParserDIVEQUAL      = 36
	gengineParserLSQARE        = 37
	gengineParserRSQARE        = 38
	gengineParserSEMICOLON     = 39
	gengineParserLR_BRACE      = 40
	gengineParserRR_BRACE      = 41
	gengineParserLR_BRACKET    = 42
	gengineParserRR_BRACKET    = 43
	gengineParserDOT           = 44
	gengineParserDQUOTA_STRING = 45
	gengineParserDOTTEDNAME    = 46
	gengineParserREAL_LITERAL  = 47
	gengineParserSL_COMMENT    = 48
	gengineParserWS            = 49
)

// gengineParser rules.
const (
	gengineParserRULE_primary            = 0
	gengineParserRULE_ruleEntity         = 1
	gengineParserRULE_ruleName           = 2
	gengineParserRULE_ruleDescription    = 3
	gengineParserRULE_salience           = 4
	gengineParserRULE_ruleContent        = 5
	gengineParserRULE_statements         = 6
	gengineParserRULE_statement          = 7
	gengineParserRULE_concStatement      = 8
	gengineParserRULE_expression         = 9
	gengineParserRULE_mathExpression     = 10
	gengineParserRULE_expressionAtom     = 11
	gengineParserRULE_assignment         = 12
	gengineParserRULE_ifStmt             = 13
	gengineParserRULE_elseIfStmt         = 14
	gengineParserRULE_elseStmt           = 15
	gengineParserRULE_constant           = 16
	gengineParserRULE_functionArgs       = 17
	gengineParserRULE_integer            = 18
	gengineParserRULE_realLiteral        = 19
	gengineParserRULE_stringLiteral      = 20
	gengineParserRULE_booleanLiteral     = 21
	gengineParserRULE_functionCall       = 22
	gengineParserRULE_methodCall         = 23
	gengineParserRULE_variable           = 24
	gengineParserRULE_mathPmOperator     = 25
	gengineParserRULE_mathMdOperator     = 26
	gengineParserRULE_comparisonOperator = 27
	gengineParserRULE_logicalOperator    = 28
	gengineParserRULE_assignOperator     = 29
	gengineParserRULE_notOperator        = 30
	gengineParserRULE_mapVar             = 31
	gengineParserRULE_atName             = 32
	gengineParserRULE_atDesc             = 33
	gengineParserRULE_atId               = 34
)

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) AllRuleEntity() []IRuleEntityContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRuleEntityContext)(nil)).Elem())
	var tst = make([]IRuleEntityContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRuleEntityContext)
		}
	}

	return tst
}

func (s *PrimaryContext) RuleEntity(i int) IRuleEntityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleEntityContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRuleEntityContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, gengineParserRULE_primary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == gengineParserRULE {
		{
			p.SetState(70)
			p.RuleEntity()
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRuleEntityContext is an interface to support dynamic dispatch.
type IRuleEntityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleEntityContext differentiates from other interfaces.
	IsRuleEntityContext()
}

type RuleEntityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleEntityContext() *RuleEntityContext {
	var p = new(RuleEntityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_ruleEntity
	return p
}

func (*RuleEntityContext) IsRuleEntityContext() {}

func NewRuleEntityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleEntityContext {
	var p = new(RuleEntityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ruleEntity

	return p
}

func (s *RuleEntityContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleEntityContext) RULE() antlr.TerminalNode {
	return s.GetToken(gengineParserRULE, 0)
}

func (s *RuleEntityContext) RuleName() IRuleNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleNameContext)
}

func (s *RuleEntityContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(gengineParserBEGIN, 0)
}

func (s *RuleEntityContext) RuleContent() IRuleContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleContentContext)
}

func (s *RuleEntityContext) END() antlr.TerminalNode {
	return s.GetToken(gengineParserEND, 0)
}

func (s *RuleEntityContext) RuleDescription() IRuleDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleDescriptionContext)
}

func (s *RuleEntityContext) Salience() ISalienceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISalienceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISalienceContext)
}

func (s *RuleEntityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleEntityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleEntityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRuleEntity(s)
	}
}

func (s *RuleEntityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRuleEntity(s)
	}
}

func (s *RuleEntityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRuleEntity(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) RuleEntity() (localctx IRuleEntityContext) {
	localctx = NewRuleEntityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, gengineParserRULE_ruleEntity)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(gengineParserRULE)
	}
	{
		p.SetState(76)
		p.RuleName()
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == gengineParserDQUOTA_STRING {
		{
			p.SetState(77)
			p.RuleDescription()
		}

	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == gengineParserSALIENCE {
		{
			p.SetState(80)
			p.Salience()
		}

	}
	{
		p.SetState(83)
		p.Match(gengineParserBEGIN)
	}
	{
		p.SetState(84)
		p.RuleContent()
	}
	{
		p.SetState(85)
		p.Match(gengineParserEND)
	}

	return localctx
}

// IRuleNameContext is an interface to support dynamic dispatch.
type IRuleNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleNameContext differentiates from other interfaces.
	IsRuleNameContext()
}

type RuleNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleNameContext() *RuleNameContext {
	var p = new(RuleNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_ruleName
	return p
}

func (*RuleNameContext) IsRuleNameContext() {}

func NewRuleNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleNameContext {
	var p = new(RuleNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ruleName

	return p
}

func (s *RuleNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleNameContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *RuleNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRuleName(s)
	}
}

func (s *RuleNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRuleName(s)
	}
}

func (s *RuleNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRuleName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) RuleName() (localctx IRuleNameContext) {
	localctx = NewRuleNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, gengineParserRULE_ruleName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.StringLiteral()
	}

	return localctx
}

// IRuleDescriptionContext is an interface to support dynamic dispatch.
type IRuleDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleDescriptionContext differentiates from other interfaces.
	IsRuleDescriptionContext()
}

type RuleDescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleDescriptionContext() *RuleDescriptionContext {
	var p = new(RuleDescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_ruleDescription
	return p
}

func (*RuleDescriptionContext) IsRuleDescriptionContext() {}

func NewRuleDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleDescriptionContext {
	var p = new(RuleDescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ruleDescription

	return p
}

func (s *RuleDescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleDescriptionContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *RuleDescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleDescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleDescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRuleDescription(s)
	}
}

func (s *RuleDescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRuleDescription(s)
	}
}

func (s *RuleDescriptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRuleDescription(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) RuleDescription() (localctx IRuleDescriptionContext) {
	localctx = NewRuleDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, gengineParserRULE_ruleDescription)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.StringLiteral()
	}

	return localctx
}

// ISalienceContext is an interface to support dynamic dispatch.
type ISalienceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSalienceContext differentiates from other interfaces.
	IsSalienceContext()
}

type SalienceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySalienceContext() *SalienceContext {
	var p = new(SalienceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_salience
	return p
}

func (*SalienceContext) IsSalienceContext() {}

func NewSalienceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SalienceContext {
	var p = new(SalienceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_salience

	return p
}

func (s *SalienceContext) GetParser() antlr.Parser { return s.parser }

func (s *SalienceContext) SALIENCE() antlr.TerminalNode {
	return s.GetToken(gengineParserSALIENCE, 0)
}

func (s *SalienceContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *SalienceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SalienceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SalienceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterSalience(s)
	}
}

func (s *SalienceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitSalience(s)
	}
}

func (s *SalienceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitSalience(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Salience() (localctx ISalienceContext) {
	localctx = NewSalienceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, gengineParserRULE_salience)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(gengineParserSALIENCE)
	}
	{
		p.SetState(92)
		p.Integer()
	}

	return localctx
}

// IRuleContentContext is an interface to support dynamic dispatch.
type IRuleContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleContentContext differentiates from other interfaces.
	IsRuleContentContext()
}

type RuleContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleContentContext() *RuleContentContext {
	var p = new(RuleContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_ruleContent
	return p
}

func (*RuleContentContext) IsRuleContentContext() {}

func NewRuleContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleContentContext {
	var p = new(RuleContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ruleContent

	return p
}

func (s *RuleContentContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleContentContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *RuleContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRuleContent(s)
	}
}

func (s *RuleContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRuleContent(s)
	}
}

func (s *RuleContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRuleContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) RuleContent() (localctx IRuleContentContext) {
	localctx = NewRuleContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, gengineParserRULE_ruleContent)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Statements()
	}

	return localctx
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, gengineParserRULE_statements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gengineParserT__0)|(1<<gengineParserT__1)|(1<<gengineParserSIMPLENAME))) != 0) || _la == gengineParserDOTTEDNAME {
		{
			p.SetState(96)
			p.Statement()
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) IfStmt() IIfStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StatementContext) MethodCall() IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *StatementContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) ConcStatement() IConcStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConcStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConcStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, gengineParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.IfStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.MethodCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(103)
			p.FunctionCall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(104)
			p.Assignment()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(105)
			p.ConcStatement()
		}

	}

	return localctx
}

// IConcStatementContext is an interface to support dynamic dispatch.
type IConcStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConcStatementContext differentiates from other interfaces.
	IsConcStatementContext()
}

type ConcStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcStatementContext() *ConcStatementContext {
	var p = new(ConcStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_concStatement
	return p
}

func (*ConcStatementContext) IsConcStatementContext() {}

func NewConcStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConcStatementContext {
	var p = new(ConcStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_concStatement

	return p
}

func (s *ConcStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ConcStatementContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *ConcStatementContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *ConcStatementContext) AllMethodCall() []IMethodCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodCallContext)(nil)).Elem())
	var tst = make([]IMethodCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodCallContext)
		}
	}

	return tst
}

func (s *ConcStatementContext) MethodCall(i int) IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *ConcStatementContext) AllFunctionCall() []IFunctionCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem())
	var tst = make([]IFunctionCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionCallContext)
		}
	}

	return tst
}

func (s *ConcStatementContext) FunctionCall(i int) IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ConcStatementContext) AllAssignment() []IAssignmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignmentContext)(nil)).Elem())
	var tst = make([]IAssignmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignmentContext)
		}
	}

	return tst
}

func (s *ConcStatementContext) Assignment(i int) IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ConcStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConcStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterConcStatement(s)
	}
}

func (s *ConcStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitConcStatement(s)
	}
}

func (s *ConcStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitConcStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ConcStatement() (localctx IConcStatementContext) {
	localctx = NewConcStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, gengineParserRULE_concStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(108)
		p.Match(gengineParserT__0)
	}
	{
		p.SetState(109)
		p.Match(gengineParserLR_BRACE)
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == gengineParserSIMPLENAME || _la == gengineParserDOTTEDNAME {
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(110)
				p.MethodCall()
			}

		case 2:
			{
				p.SetState(111)
				p.FunctionCall()
			}

		case 3:
			{
				p.SetState(112)
				p.Assignment()
			}

		}

		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(118)
		p.Match(gengineParserRR_BRACE)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) MathExpression() IMathExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
}

func (s *ExpressionContext) ExpressionAtom() IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *ExpressionContext) NotOperator() INotOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotOperatorContext)
}

func (s *ExpressionContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *ExpressionContext) ComparisonOperator() IComparisonOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *ExpressionContext) LogicalOperator() ILogicalOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicalOperatorContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *gengineParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, gengineParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(121)
			p.mathExpression(0)
		}

	case 2:
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == gengineParserNOT {
			{
				p.SetState(122)
				p.NotOperator()
			}

		}
		{
			p.SetState(125)
			p.ExpressionAtom()
		}

	case 3:
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == gengineParserNOT {
			{
				p.SetState(126)
				p.NotOperator()
			}

		}
		{
			p.SetState(129)
			p.Match(gengineParserLR_BRACKET)
		}
		{
			p.SetState(130)
			p.expression(0)
		}
		{
			p.SetState(131)
			p.Match(gengineParserRR_BRACKET)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(143)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_expression)
				p.SetState(135)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(136)
					p.ComparisonOperator()
				}
				{
					p.SetState(137)
					p.expression(5)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_expression)
				p.SetState(139)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(140)
					p.LogicalOperator()
				}
				{
					p.SetState(141)
					p.expression(4)
				}

			}

		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IMathExpressionContext is an interface to support dynamic dispatch.
type IMathExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathExpressionContext differentiates from other interfaces.
	IsMathExpressionContext()
}

type MathExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathExpressionContext() *MathExpressionContext {
	var p = new(MathExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_mathExpression
	return p
}

func (*MathExpressionContext) IsMathExpressionContext() {}

func NewMathExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathExpressionContext {
	var p = new(MathExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_mathExpression

	return p
}

func (s *MathExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MathExpressionContext) ExpressionAtom() IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *MathExpressionContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *MathExpressionContext) AllMathExpression() []IMathExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMathExpressionContext)(nil)).Elem())
	var tst = make([]IMathExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMathExpressionContext)
		}
	}

	return tst
}

func (s *MathExpressionContext) MathExpression(i int) IMathExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
}

func (s *MathExpressionContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *MathExpressionContext) MathMdOperator() IMathMdOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathMdOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathMdOperatorContext)
}

func (s *MathExpressionContext) MathPmOperator() IMathPmOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathPmOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathPmOperatorContext)
}

func (s *MathExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMathExpression(s)
	}
}

func (s *MathExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMathExpression(s)
	}
}

func (s *MathExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMathExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) MathExpression() (localctx IMathExpressionContext) {
	return p.mathExpression(0)
}

func (p *gengineParser) mathExpression(_p int) (localctx IMathExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMathExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMathExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, gengineParserRULE_mathExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(154)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gengineParserT__4, gengineParserT__5, gengineParserT__6, gengineParserTRUE, gengineParserFALSE, gengineParserSIMPLENAME, gengineParserINT, gengineParserMINUS, gengineParserDQUOTA_STRING, gengineParserDOTTEDNAME, gengineParserREAL_LITERAL:
		{
			p.SetState(149)
			p.ExpressionAtom()
		}

	case gengineParserLR_BRACKET:
		{
			p.SetState(150)
			p.Match(gengineParserLR_BRACKET)
		}
		{
			p.SetState(151)
			p.mathExpression(0)
		}
		{
			p.SetState(152)
			p.Match(gengineParserRR_BRACKET)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(164)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_mathExpression)
				p.SetState(156)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(157)
					p.MathMdOperator()
				}
				{
					p.SetState(158)
					p.mathExpression(5)
				}

			case 2:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_mathExpression)
				p.SetState(160)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(161)
					p.MathPmOperator()
				}
				{
					p.SetState(162)
					p.mathExpression(4)
				}

			}

		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())
	}

	return localctx
}

// IExpressionAtomContext is an interface to support dynamic dispatch.
type IExpressionAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionAtomContext differentiates from other interfaces.
	IsExpressionAtomContext()
}

type ExpressionAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionAtomContext() *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_expressionAtom
	return p
}

func (*ExpressionAtomContext) IsExpressionAtomContext() {}

func NewExpressionAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_expressionAtom

	return p
}

func (s *ExpressionAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionAtomContext) MethodCall() IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *ExpressionAtomContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionAtomContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ExpressionAtomContext) MapVar() IMapVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapVarContext)
}

func (s *ExpressionAtomContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterExpressionAtom(s)
	}
}

func (s *ExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitExpressionAtom(s)
	}
}

func (s *ExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ExpressionAtom() (localctx IExpressionAtomContext) {
	localctx = NewExpressionAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, gengineParserRULE_expressionAtom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(169)
			p.MethodCall()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(170)
			p.FunctionCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(171)
			p.Constant()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(172)
			p.MapVar()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(173)
			p.Variable()
		}

	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) AssignOperator() IAssignOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignOperatorContext)
}

func (s *AssignmentContext) MapVar() IMapVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapVarContext)
}

func (s *AssignmentContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AssignmentContext) MathExpression() IMathExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, gengineParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(176)
			p.MapVar()
		}

	case 2:
		{
			p.SetState(177)
			p.Variable()
		}

	}
	{
		p.SetState(180)
		p.AssignOperator()
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(181)
			p.mathExpression(0)
		}

	case 2:
		{
			p.SetState(182)
			p.expression(0)
		}

	}

	return localctx
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_ifStmt
	return p
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *IfStmtContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *IfStmtContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *IfStmtContext) AllElseIfStmt() []IElseIfStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElseIfStmtContext)(nil)).Elem())
	var tst = make([]IElseIfStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElseIfStmtContext)
		}
	}

	return tst
}

func (s *IfStmtContext) ElseIfStmt(i int) IElseIfStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseIfStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElseIfStmtContext)
}

func (s *IfStmtContext) ElseStmt() IElseStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseStmtContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, gengineParserRULE_ifStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(gengineParserT__1)
	}
	{
		p.SetState(186)
		p.expression(0)
	}
	{
		p.SetState(187)
		p.Match(gengineParserLR_BRACE)
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gengineParserT__0)|(1<<gengineParserT__1)|(1<<gengineParserSIMPLENAME))) != 0) || _la == gengineParserDOTTEDNAME {
		{
			p.SetState(188)
			p.Statements()
		}

	}
	{
		p.SetState(191)
		p.Match(gengineParserRR_BRACE)
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(192)
				p.ElseIfStmt()
			}

		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == gengineParserT__2 {
		{
			p.SetState(198)
			p.ElseStmt()
		}

	}

	return localctx
}

// IElseIfStmtContext is an interface to support dynamic dispatch.
type IElseIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseIfStmtContext differentiates from other interfaces.
	IsElseIfStmtContext()
}

type ElseIfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfStmtContext() *ElseIfStmtContext {
	var p = new(ElseIfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_elseIfStmt
	return p
}

func (*ElseIfStmtContext) IsElseIfStmtContext() {}

func NewElseIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfStmtContext {
	var p = new(ElseIfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_elseIfStmt

	return p
}

func (s *ElseIfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElseIfStmtContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *ElseIfStmtContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *ElseIfStmtContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ElseIfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterElseIfStmt(s)
	}
}

func (s *ElseIfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitElseIfStmt(s)
	}
}

func (s *ElseIfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitElseIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ElseIfStmt() (localctx IElseIfStmtContext) {
	localctx = NewElseIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, gengineParserRULE_elseIfStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(gengineParserT__2)
	}
	{
		p.SetState(202)
		p.Match(gengineParserT__1)
	}
	{
		p.SetState(203)
		p.expression(0)
	}
	{
		p.SetState(204)
		p.Match(gengineParserLR_BRACE)
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gengineParserT__0)|(1<<gengineParserT__1)|(1<<gengineParserSIMPLENAME))) != 0) || _la == gengineParserDOTTEDNAME {
		{
			p.SetState(205)
			p.Statements()
		}

	}
	{
		p.SetState(208)
		p.Match(gengineParserRR_BRACE)
	}

	return localctx
}

// IElseStmtContext is an interface to support dynamic dispatch.
type IElseStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseStmtContext differentiates from other interfaces.
	IsElseStmtContext()
}

type ElseStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStmtContext() *ElseStmtContext {
	var p = new(ElseStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_elseStmt
	return p
}

func (*ElseStmtContext) IsElseStmtContext() {}

func NewElseStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStmtContext {
	var p = new(ElseStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_elseStmt

	return p
}

func (s *ElseStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStmtContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *ElseStmtContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *ElseStmtContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ElseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterElseStmt(s)
	}
}

func (s *ElseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitElseStmt(s)
	}
}

func (s *ElseStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitElseStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ElseStmt() (localctx IElseStmtContext) {
	localctx = NewElseStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, gengineParserRULE_elseStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(gengineParserT__2)
	}
	{
		p.SetState(211)
		p.Match(gengineParserLR_BRACE)
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gengineParserT__0)|(1<<gengineParserT__1)|(1<<gengineParserSIMPLENAME))) != 0) || _la == gengineParserDOTTEDNAME {
		{
			p.SetState(212)
			p.Statements()
		}

	}
	{
		p.SetState(215)
		p.Match(gengineParserRR_BRACE)
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *ConstantContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *ConstantContext) RealLiteral() IRealLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealLiteralContext)
}

func (s *ConstantContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *ConstantContext) AtName() IAtNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtNameContext)
}

func (s *ConstantContext) AtDesc() IAtDescContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtDescContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtDescContext)
}

func (s *ConstantContext) AtId() IAtIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtIdContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, gengineParserRULE_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(217)
			p.BooleanLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(218)
			p.Integer()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(219)
			p.RealLiteral()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(220)
			p.StringLiteral()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(221)
			p.AtName()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(222)
			p.AtDesc()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(223)
			p.AtId()
		}

	}

	return localctx
}

// IFunctionArgsContext is an interface to support dynamic dispatch.
type IFunctionArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionArgsContext differentiates from other interfaces.
	IsFunctionArgsContext()
}

type FunctionArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgsContext() *FunctionArgsContext {
	var p = new(FunctionArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_functionArgs
	return p
}

func (*FunctionArgsContext) IsFunctionArgsContext() {}

func NewFunctionArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgsContext {
	var p = new(FunctionArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_functionArgs

	return p
}

func (s *FunctionArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgsContext) AllConstant() []IConstantContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstantContext)(nil)).Elem())
	var tst = make([]IConstantContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstantContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) Constant(i int) IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FunctionArgsContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *FunctionArgsContext) AllFunctionCall() []IFunctionCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem())
	var tst = make([]IFunctionCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionCallContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) FunctionCall(i int) IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionArgsContext) AllMethodCall() []IMethodCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodCallContext)(nil)).Elem())
	var tst = make([]IMethodCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodCallContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) MethodCall(i int) IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *FunctionArgsContext) AllMapVar() []IMapVarContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMapVarContext)(nil)).Elem())
	var tst = make([]IMapVarContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMapVarContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) MapVar(i int) IMapVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapVarContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMapVarContext)
}

func (s *FunctionArgsContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterFunctionArgs(s)
	}
}

func (s *FunctionArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitFunctionArgs(s)
	}
}

func (s *FunctionArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitFunctionArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) FunctionArgs() (localctx IFunctionArgsContext) {
	localctx = NewFunctionArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, gengineParserRULE_functionArgs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(226)
			p.Constant()
		}

	case 2:
		{
			p.SetState(227)
			p.Variable()
		}

	case 3:
		{
			p.SetState(228)
			p.FunctionCall()
		}

	case 4:
		{
			p.SetState(229)
			p.MethodCall()
		}

	case 5:
		{
			p.SetState(230)
			p.MapVar()
		}

	case 6:
		{
			p.SetState(231)
			p.expression(0)
		}

	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == gengineParserT__3 {
		{
			p.SetState(234)
			p.Match(gengineParserT__3)
		}
		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(235)
				p.Constant()
			}

		case 2:
			{
				p.SetState(236)
				p.Variable()
			}

		case 3:
			{
				p.SetState(237)
				p.FunctionCall()
			}

		case 4:
			{
				p.SetState(238)
				p.MethodCall()
			}

		case 5:
			{
				p.SetState(239)
				p.MapVar()
			}

		case 6:
			{
				p.SetState(240)
				p.expression(0)
			}

		}

		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) INT() antlr.TerminalNode {
	return s.GetToken(gengineParserINT, 0)
}

func (s *IntegerContext) MINUS() antlr.TerminalNode {
	return s.GetToken(gengineParserMINUS, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, gengineParserRULE_integer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == gengineParserMINUS {
		{
			p.SetState(248)
			p.Match(gengineParserMINUS)
		}

	}
	{
		p.SetState(251)
		p.Match(gengineParserINT)
	}

	return localctx
}

// IRealLiteralContext is an interface to support dynamic dispatch.
type IRealLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRealLiteralContext differentiates from other interfaces.
	IsRealLiteralContext()
}

type RealLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealLiteralContext() *RealLiteralContext {
	var p = new(RealLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_realLiteral
	return p
}

func (*RealLiteralContext) IsRealLiteralContext() {}

func NewRealLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealLiteralContext {
	var p = new(RealLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_realLiteral

	return p
}

func (s *RealLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *RealLiteralContext) REAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(gengineParserREAL_LITERAL, 0)
}

func (s *RealLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(gengineParserMINUS, 0)
}

func (s *RealLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RealLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRealLiteral(s)
	}
}

func (s *RealLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRealLiteral(s)
	}
}

func (s *RealLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRealLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) RealLiteral() (localctx IRealLiteralContext) {
	localctx = NewRealLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, gengineParserRULE_realLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == gengineParserMINUS {
		{
			p.SetState(253)
			p.Match(gengineParserMINUS)
		}

	}
	{
		p.SetState(256)
		p.Match(gengineParserREAL_LITERAL)
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) DQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(gengineParserDQUOTA_STRING, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, gengineParserRULE_stringLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Match(gengineParserDQUOTA_STRING)
	}

	return localctx
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(gengineParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(gengineParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, gengineParserRULE_booleanLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gengineParserTRUE || _la == gengineParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(gengineParserSIMPLENAME, 0)
}

func (s *FunctionCallContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *FunctionCallContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *FunctionCallContext) FunctionArgs() IFunctionArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, gengineParserRULE_functionCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		p.Match(gengineParserSIMPLENAME)
	}
	{
		p.SetState(263)
		p.Match(gengineParserLR_BRACKET)
	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gengineParserT__4)|(1<<gengineParserT__5)|(1<<gengineParserT__6)|(1<<gengineParserTRUE)|(1<<gengineParserFALSE)|(1<<gengineParserSIMPLENAME)|(1<<gengineParserINT)|(1<<gengineParserMINUS)|(1<<gengineParserNOT))) != 0) || (((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(gengineParserLR_BRACKET-42))|(1<<(gengineParserDQUOTA_STRING-42))|(1<<(gengineParserDOTTEDNAME-42))|(1<<(gengineParserREAL_LITERAL-42)))) != 0) {
		{
			p.SetState(264)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(267)
		p.Match(gengineParserRR_BRACKET)
	}

	return localctx
}

// IMethodCallContext is an interface to support dynamic dispatch.
type IMethodCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodCallContext differentiates from other interfaces.
	IsMethodCallContext()
}

type MethodCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallContext() *MethodCallContext {
	var p = new(MethodCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_methodCall
	return p
}

func (*MethodCallContext) IsMethodCallContext() {}

func NewMethodCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallContext {
	var p = new(MethodCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_methodCall

	return p
}

func (s *MethodCallContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallContext) DOTTEDNAME() antlr.TerminalNode {
	return s.GetToken(gengineParserDOTTEDNAME, 0)
}

func (s *MethodCallContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *MethodCallContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *MethodCallContext) FunctionArgs() IFunctionArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *MethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMethodCall(s)
	}
}

func (s *MethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMethodCall(s)
	}
}

func (s *MethodCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMethodCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) MethodCall() (localctx IMethodCallContext) {
	localctx = NewMethodCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, gengineParserRULE_methodCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.Match(gengineParserDOTTEDNAME)
	}
	{
		p.SetState(270)
		p.Match(gengineParserLR_BRACKET)
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gengineParserT__4)|(1<<gengineParserT__5)|(1<<gengineParserT__6)|(1<<gengineParserTRUE)|(1<<gengineParserFALSE)|(1<<gengineParserSIMPLENAME)|(1<<gengineParserINT)|(1<<gengineParserMINUS)|(1<<gengineParserNOT))) != 0) || (((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(gengineParserLR_BRACKET-42))|(1<<(gengineParserDQUOTA_STRING-42))|(1<<(gengineParserDOTTEDNAME-42))|(1<<(gengineParserREAL_LITERAL-42)))) != 0) {
		{
			p.SetState(271)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(274)
		p.Match(gengineParserRR_BRACKET)
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(gengineParserSIMPLENAME, 0)
}

func (s *VariableContext) DOTTEDNAME() antlr.TerminalNode {
	return s.GetToken(gengineParserDOTTEDNAME, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, gengineParserRULE_variable)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gengineParserSIMPLENAME || _la == gengineParserDOTTEDNAME) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMathPmOperatorContext is an interface to support dynamic dispatch.
type IMathPmOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathPmOperatorContext differentiates from other interfaces.
	IsMathPmOperatorContext()
}

type MathPmOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathPmOperatorContext() *MathPmOperatorContext {
	var p = new(MathPmOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_mathPmOperator
	return p
}

func (*MathPmOperatorContext) IsMathPmOperatorContext() {}

func NewMathPmOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathPmOperatorContext {
	var p = new(MathPmOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_mathPmOperator

	return p
}

func (s *MathPmOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MathPmOperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(gengineParserPLUS, 0)
}

func (s *MathPmOperatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(gengineParserMINUS, 0)
}

func (s *MathPmOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathPmOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathPmOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMathPmOperator(s)
	}
}

func (s *MathPmOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMathPmOperator(s)
	}
}

func (s *MathPmOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMathPmOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) MathPmOperator() (localctx IMathPmOperatorContext) {
	localctx = NewMathPmOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, gengineParserRULE_mathPmOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gengineParserPLUS || _la == gengineParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMathMdOperatorContext is an interface to support dynamic dispatch.
type IMathMdOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathMdOperatorContext differentiates from other interfaces.
	IsMathMdOperatorContext()
}

type MathMdOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathMdOperatorContext() *MathMdOperatorContext {
	var p = new(MathMdOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_mathMdOperator
	return p
}

func (*MathMdOperatorContext) IsMathMdOperatorContext() {}

func NewMathMdOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathMdOperatorContext {
	var p = new(MathMdOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_mathMdOperator

	return p
}

func (s *MathMdOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MathMdOperatorContext) MUL() antlr.TerminalNode {
	return s.GetToken(gengineParserMUL, 0)
}

func (s *MathMdOperatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(gengineParserDIV, 0)
}

func (s *MathMdOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathMdOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathMdOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMathMdOperator(s)
	}
}

func (s *MathMdOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMathMdOperator(s)
	}
}

func (s *MathMdOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMathMdOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) MathMdOperator() (localctx IMathMdOperatorContext) {
	localctx = NewMathMdOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, gengineParserRULE_mathMdOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gengineParserDIV || _la == gengineParserMUL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(gengineParserGT, 0)
}

func (s *ComparisonOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(gengineParserLT, 0)
}

func (s *ComparisonOperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(gengineParserGTE, 0)
}

func (s *ComparisonOperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(gengineParserLTE, 0)
}

func (s *ComparisonOperatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(gengineParserEQUALS, 0)
}

func (s *ComparisonOperatorContext) NOTEQUALS() antlr.TerminalNode {
	return s.GetToken(gengineParserNOTEQUALS, 0)
}

func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitComparisonOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, gengineParserRULE_comparisonOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gengineParserEQUALS)|(1<<gengineParserGT)|(1<<gengineParserLT)|(1<<gengineParserGTE)|(1<<gengineParserLTE)|(1<<gengineParserNOTEQUALS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILogicalOperatorContext is an interface to support dynamic dispatch.
type ILogicalOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicalOperatorContext differentiates from other interfaces.
	IsLogicalOperatorContext()
}

type LogicalOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOperatorContext() *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_logicalOperator
	return p
}

func (*LogicalOperatorContext) IsLogicalOperatorContext() {}

func NewLogicalOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_logicalOperator

	return p
}

func (s *LogicalOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOperatorContext) AND() antlr.TerminalNode {
	return s.GetToken(gengineParserAND, 0)
}

func (s *LogicalOperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(gengineParserOR, 0)
}

func (s *LogicalOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitLogicalOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) LogicalOperator() (localctx ILogicalOperatorContext) {
	localctx = NewLogicalOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, gengineParserRULE_logicalOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gengineParserAND || _la == gengineParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAssignOperatorContext is an interface to support dynamic dispatch.
type IAssignOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignOperatorContext differentiates from other interfaces.
	IsAssignOperatorContext()
}

type AssignOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignOperatorContext() *AssignOperatorContext {
	var p = new(AssignOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_assignOperator
	return p
}

func (*AssignOperatorContext) IsAssignOperatorContext() {}

func NewAssignOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignOperatorContext {
	var p = new(AssignOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_assignOperator

	return p
}

func (s *AssignOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignOperatorContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(gengineParserASSIGN, 0)
}

func (s *AssignOperatorContext) SET() antlr.TerminalNode {
	return s.GetToken(gengineParserSET, 0)
}

func (s *AssignOperatorContext) PLUSEQUAL() antlr.TerminalNode {
	return s.GetToken(gengineParserPLUSEQUAL, 0)
}

func (s *AssignOperatorContext) MINUSEQUAL() antlr.TerminalNode {
	return s.GetToken(gengineParserMINUSEQUAL, 0)
}

func (s *AssignOperatorContext) MULTIEQUAL() antlr.TerminalNode {
	return s.GetToken(gengineParserMULTIEQUAL, 0)
}

func (s *AssignOperatorContext) DIVEQUAL() antlr.TerminalNode {
	return s.GetToken(gengineParserDIVEQUAL, 0)
}

func (s *AssignOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAssignOperator(s)
	}
}

func (s *AssignOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAssignOperator(s)
	}
}

func (s *AssignOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAssignOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) AssignOperator() (localctx IAssignOperatorContext) {
	localctx = NewAssignOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, gengineParserRULE_assignOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(gengineParserASSIGN-31))|(1<<(gengineParserSET-31))|(1<<(gengineParserPLUSEQUAL-31))|(1<<(gengineParserMINUSEQUAL-31))|(1<<(gengineParserMULTIEQUAL-31))|(1<<(gengineParserDIVEQUAL-31)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INotOperatorContext is an interface to support dynamic dispatch.
type INotOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotOperatorContext differentiates from other interfaces.
	IsNotOperatorContext()
}

type NotOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotOperatorContext() *NotOperatorContext {
	var p = new(NotOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_notOperator
	return p
}

func (*NotOperatorContext) IsNotOperatorContext() {}

func NewNotOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotOperatorContext {
	var p = new(NotOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_notOperator

	return p
}

func (s *NotOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *NotOperatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(gengineParserNOT, 0)
}

func (s *NotOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterNotOperator(s)
	}
}

func (s *NotOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitNotOperator(s)
	}
}

func (s *NotOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitNotOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) NotOperator() (localctx INotOperatorContext) {
	localctx = NewNotOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, gengineParserRULE_notOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		p.Match(gengineParserNOT)
	}

	return localctx
}

// IMapVarContext is an interface to support dynamic dispatch.
type IMapVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapVarContext differentiates from other interfaces.
	IsMapVarContext()
}

type MapVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapVarContext() *MapVarContext {
	var p = new(MapVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_mapVar
	return p
}

func (*MapVarContext) IsMapVarContext() {}

func NewMapVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapVarContext {
	var p = new(MapVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_mapVar

	return p
}

func (s *MapVarContext) GetParser() antlr.Parser { return s.parser }

func (s *MapVarContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *MapVarContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *MapVarContext) LSQARE() antlr.TerminalNode {
	return s.GetToken(gengineParserLSQARE, 0)
}

func (s *MapVarContext) RSQARE() antlr.TerminalNode {
	return s.GetToken(gengineParserRSQARE, 0)
}

func (s *MapVarContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *MapVarContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *MapVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMapVar(s)
	}
}

func (s *MapVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMapVar(s)
	}
}

func (s *MapVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMapVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) MapVar() (localctx IMapVarContext) {
	localctx = NewMapVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, gengineParserRULE_mapVar)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Variable()
	}
	{
		p.SetState(291)
		p.Match(gengineParserLSQARE)
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gengineParserINT, gengineParserMINUS:
		{
			p.SetState(292)
			p.Integer()
		}

	case gengineParserDQUOTA_STRING:
		{
			p.SetState(293)
			p.StringLiteral()
		}

	case gengineParserSIMPLENAME, gengineParserDOTTEDNAME:
		{
			p.SetState(294)
			p.Variable()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(297)
		p.Match(gengineParserRSQARE)
	}

	return localctx
}

// IAtNameContext is an interface to support dynamic dispatch.
type IAtNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtNameContext differentiates from other interfaces.
	IsAtNameContext()
}

type AtNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtNameContext() *AtNameContext {
	var p = new(AtNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_atName
	return p
}

func (*AtNameContext) IsAtNameContext() {}

func NewAtNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtNameContext {
	var p = new(AtNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_atName

	return p
}

func (s *AtNameContext) GetParser() antlr.Parser { return s.parser }
func (s *AtNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAtName(s)
	}
}

func (s *AtNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAtName(s)
	}
}

func (s *AtNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAtName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) AtName() (localctx IAtNameContext) {
	localctx = NewAtNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, gengineParserRULE_atName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Match(gengineParserT__4)
	}

	return localctx
}

// IAtDescContext is an interface to support dynamic dispatch.
type IAtDescContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtDescContext differentiates from other interfaces.
	IsAtDescContext()
}

type AtDescContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtDescContext() *AtDescContext {
	var p = new(AtDescContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_atDesc
	return p
}

func (*AtDescContext) IsAtDescContext() {}

func NewAtDescContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtDescContext {
	var p = new(AtDescContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_atDesc

	return p
}

func (s *AtDescContext) GetParser() antlr.Parser { return s.parser }
func (s *AtDescContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtDescContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtDescContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAtDesc(s)
	}
}

func (s *AtDescContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAtDesc(s)
	}
}

func (s *AtDescContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAtDesc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) AtDesc() (localctx IAtDescContext) {
	localctx = NewAtDescContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, gengineParserRULE_atDesc)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.Match(gengineParserT__5)
	}

	return localctx
}

// IAtIdContext is an interface to support dynamic dispatch.
type IAtIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtIdContext differentiates from other interfaces.
	IsAtIdContext()
}

type AtIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtIdContext() *AtIdContext {
	var p = new(AtIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_atId
	return p
}

func (*AtIdContext) IsAtIdContext() {}

func NewAtIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtIdContext {
	var p = new(AtIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_atId

	return p
}

func (s *AtIdContext) GetParser() antlr.Parser { return s.parser }
func (s *AtIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAtId(s)
	}
}

func (s *AtIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAtId(s)
	}
}

func (s *AtIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAtId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) AtId() (localctx IAtIdContext) {
	localctx = NewAtIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, gengineParserRULE_atId)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.Match(gengineParserT__6)
	}

	return localctx
}

func (p *gengineParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 10:
		var t *MathExpressionContext = nil
		if localctx != nil {
			t = localctx.(*MathExpressionContext)
		}
		return p.MathExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *gengineParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *gengineParser) MathExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
