// Code generated from /Users/renyunyi/go_project/gengine/internal/iantlr/gengine.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 53, 474, 
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9, 
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75, 
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 3, 
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 
	7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 
	13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 
	3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 
	23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 
	3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 
	34, 3, 34, 5, 34, 241, 10, 34, 3, 34, 6, 34, 244, 10, 34, 13, 34, 14, 34, 
	245, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 
	37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 
	3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 
	42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 
	3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 
	46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 
	3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 6, 49, 319, 
	10, 49, 13, 49, 14, 49, 320, 3, 49, 7, 49, 324, 10, 49, 12, 49, 14, 49, 
	327, 11, 49, 3, 50, 6, 50, 330, 10, 50, 13, 50, 14, 50, 331, 3, 51, 3, 
	51, 3, 52, 3, 52, 3, 53, 3, 53, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 56, 
	3, 56, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 60, 3, 
	60, 3, 60, 3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 3, 63, 3, 63, 3, 64, 3, 64, 
	3, 64, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3, 66, 3, 67, 3, 67, 3, 67, 3, 
	68, 3, 68, 3, 69, 3, 69, 3, 70, 3, 70, 3, 71, 3, 71, 3, 72, 3, 72, 3, 73, 
	3, 73, 3, 74, 3, 74, 3, 75, 3, 75, 3, 76, 3, 76, 3, 76, 3, 76, 3, 76, 3, 
	76, 7, 76, 399, 10, 76, 12, 76, 14, 76, 402, 11, 76, 3, 76, 3, 76, 3, 77, 
	3, 77, 3, 77, 3, 77, 3, 78, 6, 78, 411, 10, 78, 13, 78, 14, 78, 412, 5, 
	78, 415, 10, 78, 3, 78, 3, 78, 6, 78, 419, 10, 78, 13, 78, 14, 78, 420, 
	3, 78, 6, 78, 424, 10, 78, 13, 78, 14, 78, 425, 3, 78, 3, 78, 3, 78, 3, 
	78, 6, 78, 432, 10, 78, 13, 78, 14, 78, 433, 5, 78, 436, 10, 78, 3, 78, 
	3, 78, 6, 78, 440, 10, 78, 13, 78, 14, 78, 441, 3, 78, 3, 78, 3, 78, 6, 
	78, 447, 10, 78, 13, 78, 14, 78, 448, 3, 78, 3, 78, 5, 78, 453, 10, 78, 
	3, 79, 3, 79, 3, 79, 3, 79, 7, 79, 459, 10, 79, 12, 79, 14, 79, 462, 11, 
	79, 3, 79, 3, 79, 3, 79, 3, 79, 3, 80, 6, 80, 469, 10, 80, 13, 80, 14, 
	80, 470, 3, 80, 3, 80, 3, 460, 2, 81, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 
	2, 15, 2, 17, 2, 19, 2, 21, 2, 23, 2, 25, 2, 27, 2, 29, 2, 31, 2, 33, 2, 
	35, 2, 37, 2, 39, 2, 41, 2, 43, 2, 45, 2, 47, 2, 49, 2, 51, 2, 53, 2, 55, 
	2, 57, 2, 59, 2, 61, 2, 63, 2, 65, 2, 67, 2, 69, 8, 71, 9, 73, 10, 75, 
	11, 77, 12, 79, 13, 81, 14, 83, 15, 85, 16, 87, 17, 89, 18, 91, 19, 93, 
	20, 95, 21, 97, 22, 99, 23, 101, 24, 103, 25, 105, 26, 107, 27, 109, 28, 
	111, 29, 113, 30, 115, 31, 117, 32, 119, 33, 121, 34, 123, 35, 125, 36, 
	127, 37, 129, 38, 131, 39, 133, 40, 135, 41, 137, 42, 139, 43, 141, 44, 
	143, 45, 145, 46, 147, 47, 149, 48, 151, 49, 153, 50, 155, 51, 157, 52, 
	159, 53, 3, 2, 33, 3, 2, 50, 59, 4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100, 
	100, 4, 2, 69, 69, 101, 101, 4, 2, 70, 70, 102, 102, 4, 2, 71, 71, 103, 
	103, 4, 2, 72, 72, 104, 104, 4, 2, 73, 73, 105, 105, 4, 2, 74, 74, 106, 
	106, 4, 2, 75, 75, 107, 107, 4, 2, 76, 76, 108, 108, 4, 2, 77, 77, 109, 
	109, 4, 2, 78, 78, 110, 110, 4, 2, 79, 79, 111, 111, 4, 2, 80, 80, 112, 
	112, 4, 2, 81, 81, 113, 113, 4, 2, 82, 82, 114, 114, 4, 2, 83, 83, 115, 
	115, 4, 2, 84, 84, 116, 116, 4, 2, 85, 85, 117, 117, 4, 2, 86, 86, 118, 
	118, 4, 2, 87, 87, 119, 119, 4, 2, 88, 88, 120, 120, 4, 2, 89, 89, 121, 
	121, 4, 2, 90, 90, 122, 122, 4, 2, 91, 91, 123, 123, 4, 2, 92, 92, 124, 
	124, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 
	4, 2, 36, 36, 94, 94, 5, 2, 11, 12, 15, 15, 34, 34, 2, 466, 2, 3, 3, 2, 
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 
	2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 
	2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 
	3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 
	91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 
	2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 
	2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 
	3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 
	2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127, 3, 
	2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3, 2, 2, 2, 2, 
	135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2, 2, 2, 2, 141, 3, 2, 
	2, 2, 2, 143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 3, 2, 2, 2, 2, 149, 
	3, 2, 2, 2, 2, 151, 3, 2, 2, 2, 2, 153, 3, 2, 2, 2, 2, 155, 3, 2, 2, 2, 
	2, 157, 3, 2, 2, 2, 2, 159, 3, 2, 2, 2, 3, 161, 3, 2, 2, 2, 5, 163, 3, 
	2, 2, 2, 7, 169, 3, 2, 2, 2, 9, 173, 3, 2, 2, 2, 11, 179, 3, 2, 2, 2, 13, 
	184, 3, 2, 2, 2, 15, 186, 3, 2, 2, 2, 17, 188, 3, 2, 2, 2, 19, 190, 3, 
	2, 2, 2, 21, 192, 3, 2, 2, 2, 23, 194, 3, 2, 2, 2, 25, 196, 3, 2, 2, 2, 
	27, 198, 3, 2, 2, 2, 29, 200, 3, 2, 2, 2, 31, 202, 3, 2, 2, 2, 33, 204, 
	3, 2, 2, 2, 35, 206, 3, 2, 2, 2, 37, 208, 3, 2, 2, 2, 39, 210, 3, 2, 2, 
	2, 41, 212, 3, 2, 2, 2, 43, 214, 3, 2, 2, 2, 45, 216, 3, 2, 2, 2, 47, 218, 
	3, 2, 2, 2, 49, 220, 3, 2, 2, 2, 51, 222, 3, 2, 2, 2, 53, 224, 3, 2, 2, 
	2, 55, 226, 3, 2, 2, 2, 57, 228, 3, 2, 2, 2, 59, 230, 3, 2, 2, 2, 61, 232, 
	3, 2, 2, 2, 63, 234, 3, 2, 2, 2, 65, 236, 3, 2, 2, 2, 67, 238, 3, 2, 2, 
	2, 69, 247, 3, 2, 2, 2, 71, 251, 3, 2, 2, 2, 73, 256, 3, 2, 2, 2, 75, 259, 
	3, 2, 2, 2, 77, 262, 3, 2, 2, 2, 79, 267, 3, 2, 2, 2, 81, 270, 3, 2, 2, 
	2, 83, 275, 3, 2, 2, 2, 85, 282, 3, 2, 2, 2, 87, 287, 3, 2, 2, 2, 89, 293, 
	3, 2, 2, 2, 91, 298, 3, 2, 2, 2, 93, 307, 3, 2, 2, 2, 95, 313, 3, 2, 2, 
	2, 97, 318, 3, 2, 2, 2, 99, 329, 3, 2, 2, 2, 101, 333, 3, 2, 2, 2, 103, 
	335, 3, 2, 2, 2, 105, 337, 3, 2, 2, 2, 107, 339, 3, 2, 2, 2, 109, 341, 
	3, 2, 2, 2, 111, 344, 3, 2, 2, 2, 113, 346, 3, 2, 2, 2, 115, 348, 3, 2, 
	2, 2, 117, 351, 3, 2, 2, 2, 119, 354, 3, 2, 2, 2, 121, 357, 3, 2, 2, 2, 
	123, 359, 3, 2, 2, 2, 125, 362, 3, 2, 2, 2, 127, 364, 3, 2, 2, 2, 129, 
	367, 3, 2, 2, 2, 131, 370, 3, 2, 2, 2, 133, 373, 3, 2, 2, 2, 135, 376, 
	3, 2, 2, 2, 137, 378, 3, 2, 2, 2, 139, 380, 3, 2, 2, 2, 141, 382, 3, 2, 
	2, 2, 143, 384, 3, 2, 2, 2, 145, 386, 3, 2, 2, 2, 147, 388, 3, 2, 2, 2, 
	149, 390, 3, 2, 2, 2, 151, 392, 3, 2, 2, 2, 153, 405, 3, 2, 2, 2, 155, 
	452, 3, 2, 2, 2, 157, 454, 3, 2, 2, 2, 159, 468, 3, 2, 2, 2, 161, 162, 
	7, 46, 2, 2, 162, 4, 3, 2, 2, 2, 163, 164, 7, 66, 2, 2, 164, 165, 7, 112, 
	2, 2, 165, 166, 7, 99, 2, 2, 166, 167, 7, 111, 2, 2, 167, 168, 7, 103, 
	2, 2, 168, 6, 3, 2, 2, 2, 169, 170, 7, 66, 2, 2, 170, 171, 7, 107, 2, 2, 
	171, 172, 7, 102, 2, 2, 172, 8, 3, 2, 2, 2, 173, 174, 7, 66, 2, 2, 174, 
	175, 7, 102, 2, 2, 175, 176, 7, 103, 2, 2, 176, 177, 7, 117, 2, 2, 177, 
	178, 7, 101, 2, 2, 178, 10, 3, 2, 2, 2, 179, 180, 7, 66, 2, 2, 180, 181, 
	7, 117, 2, 2, 181, 182, 7, 99, 2, 2, 182, 183, 7, 110, 2, 2, 183, 12, 3, 
	2, 2, 2, 184, 185, 9, 2, 2, 2, 185, 14, 3, 2, 2, 2, 186, 187, 9, 3, 2, 
	2, 187, 16, 3, 2, 2, 2, 188, 189, 9, 4, 2, 2, 189, 18, 3, 2, 2, 2, 190, 
	191, 9, 5, 2, 2, 191, 20, 3, 2, 2, 2, 192, 193, 9, 6, 2, 2, 193, 22, 3, 
	2, 2, 2, 194, 195, 9, 7, 2, 2, 195, 24, 3, 2, 2, 2, 196, 197, 9, 8, 2, 
	2, 197, 26, 3, 2, 2, 2, 198, 199, 9, 9, 2, 2, 199, 28, 3, 2, 2, 2, 200, 
	201, 9, 10, 2, 2, 201, 30, 3, 2, 2, 2, 202, 203, 9, 11, 2, 2, 203, 32, 
	3, 2, 2, 2, 204, 205, 9, 12, 2, 2, 205, 34, 3, 2, 2, 2, 206, 207, 9, 13, 
	2, 2, 207, 36, 3, 2, 2, 2, 208, 209, 9, 14, 2, 2, 209, 38, 3, 2, 2, 2, 
	210, 211, 9, 15, 2, 2, 211, 40, 3, 2, 2, 2, 212, 213, 9, 16, 2, 2, 213, 
	42, 3, 2, 2, 2, 214, 215, 9, 17, 2, 2, 215, 44, 3, 2, 2, 2, 216, 217, 9, 
	18, 2, 2, 217, 46, 3, 2, 2, 2, 218, 219, 9, 19, 2, 2, 219, 48, 3, 2, 2, 
	2, 220, 221, 9, 20, 2, 2, 221, 50, 3, 2, 2, 2, 222, 223, 9, 21, 2, 2, 223, 
	52, 3, 2, 2, 2, 224, 225, 9, 22, 2, 2, 225, 54, 3, 2, 2, 2, 226, 227, 9, 
	23, 2, 2, 227, 56, 3, 2, 2, 2, 228, 229, 9, 24, 2, 2, 229, 58, 3, 2, 2, 
	2, 230, 231, 9, 25, 2, 2, 231, 60, 3, 2, 2, 2, 232, 233, 9, 26, 2, 2, 233, 
	62, 3, 2, 2, 2, 234, 235, 9, 27, 2, 2, 235, 64, 3, 2, 2, 2, 236, 237, 9, 
	28, 2, 2, 237, 66, 3, 2, 2, 2, 238, 240, 9, 7, 2, 2, 239, 241, 7, 47, 2, 
	2, 240, 239, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 243, 3, 2, 2, 2, 242, 
	244, 5, 13, 7, 2, 243, 242, 3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245, 243, 
	3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 68, 3, 2, 2, 2, 247, 248, 5, 41, 
	21, 2, 248, 249, 5, 31, 16, 2, 249, 250, 5, 37, 19, 2, 250, 70, 3, 2, 2, 
	2, 251, 252, 5, 49, 25, 2, 252, 253, 5, 55, 28, 2, 253, 254, 5, 37, 19, 
	2, 254, 255, 5, 23, 12, 2, 255, 72, 3, 2, 2, 2, 256, 257, 7, 40, 2, 2, 
	257, 258, 7, 40, 2, 2, 258, 74, 3, 2, 2, 2, 259, 260, 7, 126, 2, 2, 260, 
	261, 7, 126, 2, 2, 261, 76, 3, 2, 2, 2, 262, 263, 5, 19, 10, 2, 263, 264, 
	5, 43, 22, 2, 264, 265, 5, 41, 21, 2, 265, 266, 5, 19, 10, 2, 266, 78, 
	3, 2, 2, 2, 267, 268, 5, 31, 16, 2, 268, 269, 5, 25, 13, 2, 269, 80, 3, 
	2, 2, 2, 270, 271, 5, 23, 12, 2, 271, 272, 5, 37, 19, 2, 272, 273, 5, 51, 
	26, 2, 273, 274, 5, 23, 12, 2, 274, 82, 3, 2, 2, 2, 275, 276, 5, 49, 25, 
	2, 276, 277, 5, 23, 12, 2, 277, 278, 5, 53, 27, 2, 278, 279, 5, 55, 28, 
	2, 279, 280, 5, 49, 25, 2, 280, 281, 5, 41, 21, 2, 281, 84, 3, 2, 2, 2, 
	282, 283, 5, 53, 27, 2, 283, 284, 5, 49, 25, 2, 284, 285, 5, 55, 28, 2, 
	285, 286, 5, 23, 12, 2, 286, 86, 3, 2, 2, 2, 287, 288, 5, 25, 13, 2, 288, 
	289, 5, 15, 8, 2, 289, 290, 5, 37, 19, 2, 290, 291, 5, 51, 26, 2, 291, 
	292, 5, 23, 12, 2, 292, 88, 3, 2, 2, 2, 293, 294, 5, 41, 21, 2, 294, 295, 
	5, 55, 28, 2, 295, 296, 5, 37, 19, 2, 296, 297, 5, 37, 19, 2, 297, 90, 
	3, 2, 2, 2, 298, 299, 5, 51, 26, 2, 299, 300, 5, 15, 8, 2, 300, 301, 5, 
	37, 19, 2, 301, 302, 5, 31, 16, 2, 302, 303, 5, 23, 12, 2, 303, 304, 5, 
	41, 21, 2, 304, 305, 5, 19, 10, 2, 305, 306, 5, 23, 12, 2, 306, 92, 3, 
	2, 2, 2, 307, 308, 5, 17, 9, 2, 308, 309, 5, 23, 12, 2, 309, 310, 5, 27, 
	14, 2, 310, 311, 5, 31, 16, 2, 311, 312, 5, 41, 21, 2, 312, 94, 3, 2, 2, 
	2, 313, 314, 5, 23, 12, 2, 314, 315, 5, 41, 21, 2, 315, 316, 5, 21, 11, 
	2, 316, 96, 3, 2, 2, 2, 317, 319, 9, 29, 2, 2, 318, 317, 3, 2, 2, 2, 319, 
	320, 3, 2, 2, 2, 320, 318, 3, 2, 2, 2, 320, 321, 3, 2, 2, 2, 321, 325, 
	3, 2, 2, 2, 322, 324, 9, 30, 2, 2, 323, 322, 3, 2, 2, 2, 324, 327, 3, 2, 
	2, 2, 325, 323, 3, 2, 2, 2, 325, 326, 3, 2, 2, 2, 326, 98, 3, 2, 2, 2, 
	327, 325, 3, 2, 2, 2, 328, 330, 4, 50, 59, 2, 329, 328, 3, 2, 2, 2, 330, 
	331, 3, 2, 2, 2, 331, 329, 3, 2, 2, 2, 331, 332, 3, 2, 2, 2, 332, 100, 
	3, 2, 2, 2, 333, 334, 7, 45, 2, 2, 334, 102, 3, 2, 2, 2, 335, 336, 7, 47, 
	2, 2, 336, 104, 3, 2, 2, 2, 337, 338, 7, 49, 2, 2, 338, 106, 3, 2, 2, 2, 
	339, 340, 7, 44, 2, 2, 340, 108, 3, 2, 2, 2, 341, 342, 7, 63, 2, 2, 342, 
	343, 7, 63, 2, 2, 343, 110, 3, 2, 2, 2, 344, 345, 7, 64, 2, 2, 345, 112, 
	3, 2, 2, 2, 346, 347, 7, 62, 2, 2, 347, 114, 3, 2, 2, 2, 348, 349, 7, 64, 
	2, 2, 349, 350, 7, 63, 2, 2, 350, 116, 3, 2, 2, 2, 351, 352, 7, 62, 2, 
	2, 352, 353, 7, 63, 2, 2, 353, 118, 3, 2, 2, 2, 354, 355, 7, 35, 2, 2, 
	355, 356, 7, 63, 2, 2, 356, 120, 3, 2, 2, 2, 357, 358, 7, 35, 2, 2, 358, 
	122, 3, 2, 2, 2, 359, 360, 7, 60, 2, 2, 360, 361, 7, 63, 2, 2, 361, 124, 
	3, 2, 2, 2, 362, 363, 7, 63, 2, 2, 363, 126, 3, 2, 2, 2, 364, 365, 7, 45, 
	2, 2, 365, 366, 7, 63, 2, 2, 366, 128, 3, 2, 2, 2, 367, 368, 7, 47, 2, 
	2, 368, 369, 7, 63, 2, 2, 369, 130, 3, 2, 2, 2, 370, 371, 7, 44, 2, 2, 
	371, 372, 7, 63, 2, 2, 372, 132, 3, 2, 2, 2, 373, 374, 7, 49, 2, 2, 374, 
	375, 7, 63, 2, 2, 375, 134, 3, 2, 2, 2, 376, 377, 7, 93, 2, 2, 377, 136, 
	3, 2, 2, 2, 378, 379, 7, 95, 2, 2, 379, 138, 3, 2, 2, 2, 380, 381, 7, 61, 
	2, 2, 381, 140, 3, 2, 2, 2, 382, 383, 7, 125, 2, 2, 383, 142, 3, 2, 2, 
	2, 384, 385, 7, 127, 2, 2, 385, 144, 3, 2, 2, 2, 386, 387, 7, 42, 2, 2, 
	387, 146, 3, 2, 2, 2, 388, 389, 7, 43, 2, 2, 389, 148, 3, 2, 2, 2, 390, 
	391, 7, 48, 2, 2, 391, 150, 3, 2, 2, 2, 392, 400, 7, 36, 2, 2, 393, 394, 
	7, 94, 2, 2, 394, 399, 11, 2, 2, 2, 395, 396, 7, 36, 2, 2, 396, 399, 7, 
	36, 2, 2, 397, 399, 10, 31, 2, 2, 398, 393, 3, 2, 2, 2, 398, 395, 3, 2, 
	2, 2, 398, 397, 3, 2, 2, 2, 399, 402, 3, 2, 2, 2, 400, 398, 3, 2, 2, 2, 
	400, 401, 3, 2, 2, 2, 401, 403, 3, 2, 2, 2, 402, 400, 3, 2, 2, 2, 403, 
	404, 7, 36, 2, 2, 404, 152, 3, 2, 2, 2, 405, 406, 5, 97, 49, 2, 406, 407, 
	5, 149, 75, 2, 407, 408, 5, 97, 49, 2, 408, 154, 3, 2, 2, 2, 409, 411, 
	5, 13, 7, 2, 410, 409, 3, 2, 2, 2, 411, 412, 3, 2, 2, 2, 412, 410, 3, 2, 
	2, 2, 412, 413, 3, 2, 2, 2, 413, 415, 3, 2, 2, 2, 414, 410, 3, 2, 2, 2, 
	414, 415, 3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 418, 7, 48, 2, 2, 417, 
	419, 5, 13, 7, 2, 418, 417, 3, 2, 2, 2, 419, 420, 3, 2, 2, 2, 420, 418, 
	3, 2, 2, 2, 420, 421, 3, 2, 2, 2, 421, 453, 3, 2, 2, 2, 422, 424, 5, 13, 
	7, 2, 423, 422, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2, 425, 423, 3, 2, 2, 2, 
	425, 426, 3, 2, 2, 2, 426, 427, 3, 2, 2, 2, 427, 428, 7, 48, 2, 2, 428, 
	429, 5, 67, 34, 2, 429, 453, 3, 2, 2, 2, 430, 432, 5, 13, 7, 2, 431, 430, 
	3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 431, 3, 2, 2, 2, 433, 434, 3, 2, 
	2, 2, 434, 436, 3, 2, 2, 2, 435, 431, 3, 2, 2, 2, 435, 436, 3, 2, 2, 2, 
	436, 437, 3, 2, 2, 2, 437, 439, 7, 48, 2, 2, 438, 440, 5, 13, 7, 2, 439, 
	438, 3, 2, 2, 2, 440, 441, 3, 2, 2, 2, 441, 439, 3, 2, 2, 2, 441, 442, 
	3, 2, 2, 2, 442, 443, 3, 2, 2, 2, 443, 444, 5, 67, 34, 2, 444, 453, 3, 
	2, 2, 2, 445, 447, 5, 13, 7, 2, 446, 445, 3, 2, 2, 2, 447, 448, 3, 2, 2, 
	2, 448, 446, 3, 2, 2, 2, 448, 449, 3, 2, 2, 2, 449, 450, 3, 2, 2, 2, 450, 
	451, 5, 67, 34, 2, 451, 453, 3, 2, 2, 2, 452, 414, 3, 2, 2, 2, 452, 423, 
	3, 2, 2, 2, 452, 435, 3, 2, 2, 2, 452, 446, 3, 2, 2, 2, 453, 156, 3, 2, 
	2, 2, 454, 455, 7, 49, 2, 2, 455, 456, 7, 49, 2, 2, 456, 460, 3, 2, 2, 
	2, 457, 459, 11, 2, 2, 2, 458, 457, 3, 2, 2, 2, 459, 462, 3, 2, 2, 2, 460, 
	461, 3, 2, 2, 2, 460, 458, 3, 2, 2, 2, 461, 463, 3, 2, 2, 2, 462, 460, 
	3, 2, 2, 2, 463, 464, 7, 12, 2, 2, 464, 465, 3, 2, 2, 2, 465, 466, 8, 79, 
	2, 2, 466, 158, 3, 2, 2, 2, 467, 469, 9, 32, 2, 2, 468, 467, 3, 2, 2, 2, 
	469, 470, 3, 2, 2, 2, 470, 468, 3, 2, 2, 2, 470, 471, 3, 2, 2, 2, 471, 
	472, 3, 2, 2, 2, 472, 473, 8, 80, 2, 2, 473, 160, 3, 2, 2, 2, 22, 2, 240, 
	245, 320, 323, 325, 331, 398, 400, 412, 414, 420, 425, 433, 435, 441, 448, 
	452, 460, 470, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "'@name'", "'@id'", "'@desc'", "'@sal'", "", "", "'&&'", "'||'", 
	"", "", "", "", "", "", "", "", "", "", "", "", "'+'", "'-'", "'/'", "'*'", 
	"'=='", "'>'", "'<'", "'>='", "'<='", "'!='", "'!'", "':='", "'='", "'+='", 
	"'-='", "'*='", "'/='", "'['", "']'", "';'", "'{'", "'}'", "'('", "')'", 
	"'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "NIL", "RULE", "AND", "OR", "CONC", "IF", "ELSE", 
	"RETURN", "TRUE", "FALSE", "NULL_LITERAL", "SALIENCE", "BEGIN", "END", 
	"SIMPLENAME", "INT", "PLUS", "MINUS", "DIV", "MUL", "EQUALS", "GT", "LT", 
	"GTE", "LTE", "NOTEQUALS", "NOT", "ASSIGN", "SET", "PLUSEQUAL", "MINUSEQUAL", 
	"MULTIEQUAL", "DIVEQUAL", "LSQARE", "RSQARE", "SEMICOLON", "LR_BRACE", 
	"RR_BRACE", "LR_BRACKET", "RR_BRACKET", "DOT", "DQUOTA_STRING", "DOTTEDNAME", 
	"REAL_LITERAL", "SL_COMMENT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "DEC_DIGIT", "A", "B", "C", "D", 
	"E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", 
	"T", "U", "V", "W", "X", "Y", "Z", "EXPONENT_NUM_PART", "NIL", "RULE", 
	"AND", "OR", "CONC", "IF", "ELSE", "RETURN", "TRUE", "FALSE", "NULL_LITERAL", 
	"SALIENCE", "BEGIN", "END", "SIMPLENAME", "INT", "PLUS", "MINUS", "DIV", 
	"MUL", "EQUALS", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "NOT", "ASSIGN", 
	"SET", "PLUSEQUAL", "MINUSEQUAL", "MULTIEQUAL", "DIVEQUAL", "LSQARE", "RSQARE", 
	"SEMICOLON", "LR_BRACE", "RR_BRACE", "LR_BRACKET", "RR_BRACKET", "DOT", 
	"DQUOTA_STRING", "DOTTEDNAME", "REAL_LITERAL", "SL_COMMENT", "WS",
}
type gengineLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

// NewgengineLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *gengineLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewgengineLexer(input antlr.CharStream) *gengineLexer {
	l := new(gengineLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "gengine.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// gengineLexer tokens.
const (
	gengineLexerT__0 = 1
	gengineLexerT__1 = 2
	gengineLexerT__2 = 3
	gengineLexerT__3 = 4
	gengineLexerT__4 = 5
	gengineLexerNIL = 6
	gengineLexerRULE = 7
	gengineLexerAND = 8
	gengineLexerOR = 9
	gengineLexerCONC = 10
	gengineLexerIF = 11
	gengineLexerELSE = 12
	gengineLexerRETURN = 13
	gengineLexerTRUE = 14
	gengineLexerFALSE = 15
	gengineLexerNULL_LITERAL = 16
	gengineLexerSALIENCE = 17
	gengineLexerBEGIN = 18
	gengineLexerEND = 19
	gengineLexerSIMPLENAME = 20
	gengineLexerINT = 21
	gengineLexerPLUS = 22
	gengineLexerMINUS = 23
	gengineLexerDIV = 24
	gengineLexerMUL = 25
	gengineLexerEQUALS = 26
	gengineLexerGT = 27
	gengineLexerLT = 28
	gengineLexerGTE = 29
	gengineLexerLTE = 30
	gengineLexerNOTEQUALS = 31
	gengineLexerNOT = 32
	gengineLexerASSIGN = 33
	gengineLexerSET = 34
	gengineLexerPLUSEQUAL = 35
	gengineLexerMINUSEQUAL = 36
	gengineLexerMULTIEQUAL = 37
	gengineLexerDIVEQUAL = 38
	gengineLexerLSQARE = 39
	gengineLexerRSQARE = 40
	gengineLexerSEMICOLON = 41
	gengineLexerLR_BRACE = 42
	gengineLexerRR_BRACE = 43
	gengineLexerLR_BRACKET = 44
	gengineLexerRR_BRACKET = 45
	gengineLexerDOT = 46
	gengineLexerDQUOTA_STRING = 47
	gengineLexerDOTTEDNAME = 48
	gengineLexerREAL_LITERAL = 49
	gengineLexerSL_COMMENT = 50
	gengineLexerWS = 51
)

