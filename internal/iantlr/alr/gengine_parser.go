// Code generated from /Users/renyunyi/go/src/gengine/internal/iantlr/gengine.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 45, 301, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23, 
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4, 
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34, 
	9, 34, 4, 35, 9, 35, 3, 2, 6, 2, 72, 10, 2, 13, 2, 14, 2, 73, 3, 3, 3, 
	3, 3, 3, 5, 3, 79, 10, 3, 3, 3, 5, 3, 82, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 
	3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 6, 8, 98, 10, 
	8, 13, 8, 14, 8, 99, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 107, 10, 9, 3, 
	10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 114, 10, 10, 12, 10, 14, 10, 117, 
	11, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 5, 11, 124, 10, 11, 3, 11, 3, 
	11, 5, 11, 128, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 134, 10, 11, 
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 144, 10, 
	11, 12, 11, 14, 11, 147, 11, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 
	12, 5, 12, 155, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 
	3, 12, 7, 12, 165, 10, 12, 12, 12, 14, 12, 168, 11, 12, 3, 13, 3, 13, 3, 
	13, 3, 13, 3, 13, 5, 13, 175, 10, 13, 3, 14, 3, 14, 5, 14, 179, 10, 14, 
	3, 14, 3, 14, 5, 14, 183, 10, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 
	15, 5, 15, 191, 10, 15, 3, 15, 3, 15, 7, 15, 195, 10, 15, 12, 15, 14, 15, 
	198, 11, 15, 3, 15, 5, 15, 201, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 
	16, 5, 16, 208, 10, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 5, 17, 215, 
	10, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 224, 10, 
	18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 231, 10, 19, 3, 19, 3, 19, 
	3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 239, 10, 19, 7, 19, 241, 10, 19, 12, 
	19, 14, 19, 244, 11, 19, 3, 20, 5, 20, 247, 10, 20, 3, 20, 3, 20, 3, 21, 
	5, 21, 252, 10, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 
	24, 3, 24, 5, 24, 263, 10, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 5, 25, 
	270, 10, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 
	29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 
	3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 295, 10, 34, 3, 34, 3, 34, 3, 35, 3, 
	35, 3, 35, 2, 4, 20, 22, 36, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 
	62, 64, 66, 68, 2, 8, 3, 2, 12, 13, 4, 2, 18, 18, 42, 42, 3, 2, 20, 21, 
	3, 2, 22, 23, 3, 2, 24, 29, 3, 2, 10, 11, 2, 316, 2, 71, 3, 2, 2, 2, 4, 
	75, 3, 2, 2, 2, 6, 87, 3, 2, 2, 2, 8, 89, 3, 2, 2, 2, 10, 91, 3, 2, 2, 
	2, 12, 94, 3, 2, 2, 2, 14, 97, 3, 2, 2, 2, 16, 106, 3, 2, 2, 2, 18, 108, 
	3, 2, 2, 2, 20, 133, 3, 2, 2, 2, 22, 154, 3, 2, 2, 2, 24, 174, 3, 2, 2, 
	2, 26, 178, 3, 2, 2, 2, 28, 186, 3, 2, 2, 2, 30, 202, 3, 2, 2, 2, 32, 211, 
	3, 2, 2, 2, 34, 223, 3, 2, 2, 2, 36, 230, 3, 2, 2, 2, 38, 246, 3, 2, 2, 
	2, 40, 251, 3, 2, 2, 2, 42, 255, 3, 2, 2, 2, 44, 257, 3, 2, 2, 2, 46, 259, 
	3, 2, 2, 2, 48, 266, 3, 2, 2, 2, 50, 273, 3, 2, 2, 2, 52, 275, 3, 2, 2, 
	2, 54, 277, 3, 2, 2, 2, 56, 279, 3, 2, 2, 2, 58, 281, 3, 2, 2, 2, 60, 283, 
	3, 2, 2, 2, 62, 285, 3, 2, 2, 2, 64, 287, 3, 2, 2, 2, 66, 289, 3, 2, 2, 
	2, 68, 298, 3, 2, 2, 2, 70, 72, 5, 4, 3, 2, 71, 70, 3, 2, 2, 2, 72, 73, 
	3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 3, 3, 2, 2, 2, 
	75, 76, 7, 9, 2, 2, 76, 78, 5, 6, 4, 2, 77, 79, 5, 8, 5, 2, 78, 77, 3, 
	2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 81, 3, 2, 2, 2, 80, 82, 5, 10, 6, 2, 81, 
	80, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 84, 7, 16, 
	2, 2, 84, 85, 5, 12, 7, 2, 85, 86, 7, 17, 2, 2, 86, 5, 3, 2, 2, 2, 87, 
	88, 5, 42, 22, 2, 88, 7, 3, 2, 2, 2, 89, 90, 5, 42, 22, 2, 90, 9, 3, 2, 
	2, 2, 91, 92, 7, 15, 2, 2, 92, 93, 5, 38, 20, 2, 93, 11, 3, 2, 2, 2, 94, 
	95, 5, 14, 8, 2, 95, 13, 3, 2, 2, 2, 96, 98, 5, 16, 9, 2, 97, 96, 3, 2, 
	2, 2, 98, 99, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 
	15, 3, 2, 2, 2, 101, 107, 5, 28, 15, 2, 102, 107, 5, 48, 25, 2, 103, 107, 
	5, 46, 24, 2, 104, 107, 5, 26, 14, 2, 105, 107, 5, 18, 10, 2, 106, 101, 
	3, 2, 2, 2, 106, 102, 3, 2, 2, 2, 106, 103, 3, 2, 2, 2, 106, 104, 3, 2, 
	2, 2, 106, 105, 3, 2, 2, 2, 107, 17, 3, 2, 2, 2, 108, 109, 7, 3, 2, 2, 
	109, 115, 7, 36, 2, 2, 110, 114, 5, 48, 25, 2, 111, 114, 5, 46, 24, 2, 
	112, 114, 5, 26, 14, 2, 113, 110, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 113, 
	112, 3, 2, 2, 2, 114, 117, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115, 116, 
	3, 2, 2, 2, 116, 118, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 118, 119, 7, 37, 
	2, 2, 119, 19, 3, 2, 2, 2, 120, 121, 8, 11, 1, 2, 121, 134, 5, 22, 12, 
	2, 122, 124, 5, 64, 33, 2, 123, 122, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 
	124, 125, 3, 2, 2, 2, 125, 134, 5, 24, 13, 2, 126, 128, 5, 64, 33, 2, 127, 
	126, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129, 130, 
	7, 38, 2, 2, 130, 131, 5, 20, 11, 2, 131, 132, 7, 39, 2, 2, 132, 134, 3, 
	2, 2, 2, 133, 120, 3, 2, 2, 2, 133, 123, 3, 2, 2, 2, 133, 127, 3, 2, 2, 
	2, 134, 145, 3, 2, 2, 2, 135, 136, 12, 6, 2, 2, 136, 137, 5, 56, 29, 2, 
	137, 138, 5, 20, 11, 7, 138, 144, 3, 2, 2, 2, 139, 140, 12, 5, 2, 2, 140, 
	141, 5, 58, 30, 2, 141, 142, 5, 20, 11, 6, 142, 144, 3, 2, 2, 2, 143, 135, 
	3, 2, 2, 2, 143, 139, 3, 2, 2, 2, 144, 147, 3, 2, 2, 2, 145, 143, 3, 2, 
	2, 2, 145, 146, 3, 2, 2, 2, 146, 21, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 
	148, 149, 8, 12, 1, 2, 149, 155, 5, 24, 13, 2, 150, 151, 7, 38, 2, 2, 151, 
	152, 5, 22, 12, 2, 152, 153, 7, 39, 2, 2, 153, 155, 3, 2, 2, 2, 154, 148, 
	3, 2, 2, 2, 154, 150, 3, 2, 2, 2, 155, 166, 3, 2, 2, 2, 156, 157, 12, 6, 
	2, 2, 157, 158, 5, 54, 28, 2, 158, 159, 5, 22, 12, 7, 159, 165, 3, 2, 2, 
	2, 160, 161, 12, 5, 2, 2, 161, 162, 5, 52, 27, 2, 162, 163, 5, 22, 12, 
	6, 163, 165, 3, 2, 2, 2, 164, 156, 3, 2, 2, 2, 164, 160, 3, 2, 2, 2, 165, 
	168, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 23, 3, 
	2, 2, 2, 168, 166, 3, 2, 2, 2, 169, 175, 5, 48, 25, 2, 170, 175, 5, 46, 
	24, 2, 171, 175, 5, 34, 18, 2, 172, 175, 5, 66, 34, 2, 173, 175, 5, 50, 
	26, 2, 174, 169, 3, 2, 2, 2, 174, 170, 3, 2, 2, 2, 174, 171, 3, 2, 2, 2, 
	174, 172, 3, 2, 2, 2, 174, 173, 3, 2, 2, 2, 175, 25, 3, 2, 2, 2, 176, 179, 
	5, 66, 34, 2, 177, 179, 5, 50, 26, 2, 178, 176, 3, 2, 2, 2, 178, 177, 3, 
	2, 2, 2, 179, 182, 3, 2, 2, 2, 180, 183, 5, 60, 31, 2, 181, 183, 5, 62, 
	32, 2, 182, 180, 3, 2, 2, 2, 182, 181, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 
	184, 185, 5, 22, 12, 2, 185, 27, 3, 2, 2, 2, 186, 187, 7, 4, 2, 2, 187, 
	188, 5, 20, 11, 2, 188, 190, 7, 36, 2, 2, 189, 191, 5, 14, 8, 2, 190, 189, 
	3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 196, 7, 37, 
	2, 2, 193, 195, 5, 30, 16, 2, 194, 193, 3, 2, 2, 2, 195, 198, 3, 2, 2, 
	2, 196, 194, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 200, 3, 2, 2, 2, 198, 
	196, 3, 2, 2, 2, 199, 201, 5, 32, 17, 2, 200, 199, 3, 2, 2, 2, 200, 201, 
	3, 2, 2, 2, 201, 29, 3, 2, 2, 2, 202, 203, 7, 5, 2, 2, 203, 204, 7, 4, 
	2, 2, 204, 205, 5, 20, 11, 2, 205, 207, 7, 36, 2, 2, 206, 208, 5, 14, 8, 
	2, 207, 206, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 
	210, 7, 37, 2, 2, 210, 31, 3, 2, 2, 2, 211, 212, 7, 5, 2, 2, 212, 214, 
	7, 36, 2, 2, 213, 215, 5, 14, 8, 2, 214, 213, 3, 2, 2, 2, 214, 215, 3, 
	2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 217, 7, 37, 2, 2, 217, 33, 3, 2, 2, 
	2, 218, 224, 5, 44, 23, 2, 219, 224, 5, 38, 20, 2, 220, 224, 5, 40, 21, 
	2, 221, 224, 5, 42, 22, 2, 222, 224, 5, 68, 35, 2, 223, 218, 3, 2, 2, 2, 
	223, 219, 3, 2, 2, 2, 223, 220, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 223, 
	222, 3, 2, 2, 2, 224, 35, 3, 2, 2, 2, 225, 231, 5, 34, 18, 2, 226, 231, 
	5, 50, 26, 2, 227, 231, 5, 46, 24, 2, 228, 231, 5, 48, 25, 2, 229, 231, 
	5, 66, 34, 2, 230, 225, 3, 2, 2, 2, 230, 226, 3, 2, 2, 2, 230, 227, 3, 
	2, 2, 2, 230, 228, 3, 2, 2, 2, 230, 229, 3, 2, 2, 2, 231, 242, 3, 2, 2, 
	2, 232, 238, 7, 6, 2, 2, 233, 239, 5, 34, 18, 2, 234, 239, 5, 50, 26, 2, 
	235, 239, 5, 46, 24, 2, 236, 239, 5, 48, 25, 2, 237, 239, 5, 66, 34, 2, 
	238, 233, 3, 2, 2, 2, 238, 234, 3, 2, 2, 2, 238, 235, 3, 2, 2, 2, 238, 
	236, 3, 2, 2, 2, 238, 237, 3, 2, 2, 2, 239, 241, 3, 2, 2, 2, 240, 232, 
	3, 2, 2, 2, 241, 244, 3, 2, 2, 2, 242, 240, 3, 2, 2, 2, 242, 243, 3, 2, 
	2, 2, 243, 37, 3, 2, 2, 2, 244, 242, 3, 2, 2, 2, 245, 247, 7, 21, 2, 2, 
	246, 245, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 
	249, 7, 19, 2, 2, 249, 39, 3, 2, 2, 2, 250, 252, 7, 21, 2, 2, 251, 250, 
	3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 254, 7, 43, 
	2, 2, 254, 41, 3, 2, 2, 2, 255, 256, 7, 41, 2, 2, 256, 43, 3, 2, 2, 2, 
	257, 258, 9, 2, 2, 2, 258, 45, 3, 2, 2, 2, 259, 260, 7, 18, 2, 2, 260, 
	262, 7, 38, 2, 2, 261, 263, 5, 36, 19, 2, 262, 261, 3, 2, 2, 2, 262, 263, 
	3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 265, 7, 39, 2, 2, 265, 47, 3, 2, 
	2, 2, 266, 267, 7, 42, 2, 2, 267, 269, 7, 38, 2, 2, 268, 270, 5, 36, 19, 
	2, 269, 268, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270, 271, 3, 2, 2, 2, 271, 
	272, 7, 39, 2, 2, 272, 49, 3, 2, 2, 2, 273, 274, 9, 3, 2, 2, 274, 51, 3, 
	2, 2, 2, 275, 276, 9, 4, 2, 2, 276, 53, 3, 2, 2, 2, 277, 278, 9, 5, 2, 
	2, 278, 55, 3, 2, 2, 2, 279, 280, 9, 6, 2, 2, 280, 57, 3, 2, 2, 2, 281, 
	282, 9, 7, 2, 2, 282, 59, 3, 2, 2, 2, 283, 284, 7, 31, 2, 2, 284, 61, 3, 
	2, 2, 2, 285, 286, 7, 32, 2, 2, 286, 63, 3, 2, 2, 2, 287, 288, 7, 30, 2, 
	2, 288, 65, 3, 2, 2, 2, 289, 290, 5, 50, 26, 2, 290, 294, 7, 33, 2, 2, 
	291, 295, 5, 38, 20, 2, 292, 295, 5, 42, 22, 2, 293, 295, 5, 50, 26, 2, 
	294, 291, 3, 2, 2, 2, 294, 292, 3, 2, 2, 2, 294, 293, 3, 2, 2, 2, 295, 
	296, 3, 2, 2, 2, 296, 297, 7, 34, 2, 2, 297, 67, 3, 2, 2, 2, 298, 299, 
	7, 7, 2, 2, 299, 69, 3, 2, 2, 2, 34, 73, 78, 81, 99, 106, 113, 115, 123, 
	127, 133, 143, 145, 154, 164, 166, 174, 178, 182, 190, 196, 200, 207, 214, 
	223, 230, 238, 242, 246, 251, 262, 269, 294,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'conc'", "'if'", "'else'", "','", "'@name'", "", "", "'&&'", "'||'", 
	"", "", "", "", "", "", "", "", "'+'", "'-'", "'/'", "'*'", "'=='", "'>'", 
	"'<'", "'>='", "'<='", "'!='", "'!'", "':='", "'='", "'['", "']'", "';'", 
	"'{'", "'}'", "'('", "')'", "'.'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "NIL", "RULE", "AND", "OR", "TRUE", "FALSE", "NULL_LITERAL", 
	"SALIENCE", "BEGIN", "END", "SIMPLENAME", "INT", "PLUS", "MINUS", "DIV", 
	"MUL", "EQUALS", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "NOT", "ASSIGN", 
	"SET", "LSQARE", "RSQARE", "SEMICOLON", "LR_BRACE", "RR_BRACE", "LR_BRACKET", 
	"RR_BRACKET", "DOT", "DQUOTA_STRING", "DOTTEDNAME", "REAL_LITERAL", "SL_COMMENT", 
	"WS",
}

var ruleNames = []string{
	"primary", "ruleEntity", "ruleName", "ruleDescription", "salience", "ruleContent", 
	"statements", "statement", "concStatement", "expression", "mathExpression", 
	"expressionAtom", "assignment", "ifStmt", "elseIfStmt", "elseStmt", "constant", 
	"functionArgs", "integer", "realLiteral", "stringLiteral", "booleanLiteral", 
	"functionCall", "methodCall", "variable", "mathPmOperator", "mathMdOperator", 
	"comparisonOperator", "logicalOperator", "assignOperator", "setOperator", 
	"notOperator", "mapVar", "atName",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type gengineParser struct {
	*antlr.BaseParser
}

func NewgengineParser(input antlr.TokenStream) *gengineParser {
	this := new(gengineParser)

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
	gengineParserEOF = antlr.TokenEOF
	gengineParserT__0 = 1
	gengineParserT__1 = 2
	gengineParserT__2 = 3
	gengineParserT__3 = 4
	gengineParserT__4 = 5
	gengineParserNIL = 6
	gengineParserRULE = 7
	gengineParserAND = 8
	gengineParserOR = 9
	gengineParserTRUE = 10
	gengineParserFALSE = 11
	gengineParserNULL_LITERAL = 12
	gengineParserSALIENCE = 13
	gengineParserBEGIN = 14
	gengineParserEND = 15
	gengineParserSIMPLENAME = 16
	gengineParserINT = 17
	gengineParserPLUS = 18
	gengineParserMINUS = 19
	gengineParserDIV = 20
	gengineParserMUL = 21
	gengineParserEQUALS = 22
	gengineParserGT = 23
	gengineParserLT = 24
	gengineParserGTE = 25
	gengineParserLTE = 26
	gengineParserNOTEQUALS = 27
	gengineParserNOT = 28
	gengineParserASSIGN = 29
	gengineParserSET = 30
	gengineParserLSQARE = 31
	gengineParserRSQARE = 32
	gengineParserSEMICOLON = 33
	gengineParserLR_BRACE = 34
	gengineParserRR_BRACE = 35
	gengineParserLR_BRACKET = 36
	gengineParserRR_BRACKET = 37
	gengineParserDOT = 38
	gengineParserDQUOTA_STRING = 39
	gengineParserDOTTEDNAME = 40
	gengineParserREAL_LITERAL = 41
	gengineParserSL_COMMENT = 42
	gengineParserWS = 43
)

// gengineParser rules.
const (
	gengineParserRULE_primary = 0
	gengineParserRULE_ruleEntity = 1
	gengineParserRULE_ruleName = 2
	gengineParserRULE_ruleDescription = 3
	gengineParserRULE_salience = 4
	gengineParserRULE_ruleContent = 5
	gengineParserRULE_statements = 6
	gengineParserRULE_statement = 7
	gengineParserRULE_concStatement = 8
	gengineParserRULE_expression = 9
	gengineParserRULE_mathExpression = 10
	gengineParserRULE_expressionAtom = 11
	gengineParserRULE_assignment = 12
	gengineParserRULE_ifStmt = 13
	gengineParserRULE_elseIfStmt = 14
	gengineParserRULE_elseStmt = 15
	gengineParserRULE_constant = 16
	gengineParserRULE_functionArgs = 17
	gengineParserRULE_integer = 18
	gengineParserRULE_realLiteral = 19
	gengineParserRULE_stringLiteral = 20
	gengineParserRULE_booleanLiteral = 21
	gengineParserRULE_functionCall = 22
	gengineParserRULE_methodCall = 23
	gengineParserRULE_variable = 24
	gengineParserRULE_mathPmOperator = 25
	gengineParserRULE_mathMdOperator = 26
	gengineParserRULE_comparisonOperator = 27
	gengineParserRULE_logicalOperator = 28
	gengineParserRULE_assignOperator = 29
	gengineParserRULE_setOperator = 30
	gengineParserRULE_notOperator = 31
	gengineParserRULE_mapVar = 32
	gengineParserRULE_atName = 33
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
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = _la == gengineParserRULE {
		{
			p.SetState(68)
			p.RuleEntity()
		}


		p.SetState(71)
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
		p.SetState(73)
		p.Match(gengineParserRULE)
	}
	{
		p.SetState(74)
		p.RuleName()
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == gengineParserDQUOTA_STRING {
		{
			p.SetState(75)
			p.RuleDescription()
		}

	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == gengineParserSALIENCE {
		{
			p.SetState(78)
			p.Salience()
		}

	}
	{
		p.SetState(81)
		p.Match(gengineParserBEGIN)
	}
	{
		p.SetState(82)
		p.RuleContent()
	}
	{
		p.SetState(83)
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
		p.SetState(85)
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
		p.SetState(87)
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
		p.SetState(89)
		p.Match(gengineParserSALIENCE)
	}
	{
		p.SetState(90)
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
		p.SetState(92)
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
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for ok := true; ok; ok = (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << gengineParserT__0) | (1 << gengineParserT__1) | (1 << gengineParserSIMPLENAME))) != 0) || _la == gengineParserDOTTEDNAME {
		{
			p.SetState(94)
			p.Statement()
		}


		p.SetState(97)
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

	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.IfStmt()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
			p.MethodCall()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(101)
			p.FunctionCall()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(102)
			p.Assignment()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(103)
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
		p.SetState(106)
		p.Match(gengineParserT__0)
	}
	{
		p.SetState(107)
		p.Match(gengineParserLR_BRACE)
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == gengineParserSIMPLENAME || _la == gengineParserDOTTEDNAME {
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(108)
				p.MethodCall()
			}


		case 2:
			{
				p.SetState(109)
				p.FunctionCall()
			}


		case 3:
			{
				p.SetState(110)
				p.Assignment()
			}

		}

		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(116)
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
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(119)
			p.mathExpression(0)
		}


	case 2:
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == gengineParserNOT {
			{
				p.SetState(120)
				p.NotOperator()
			}

		}
		{
			p.SetState(123)
			p.ExpressionAtom()
		}


	case 3:
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == gengineParserNOT {
			{
				p.SetState(124)
				p.NotOperator()
			}

		}
		{
			p.SetState(127)
			p.Match(gengineParserLR_BRACKET)
		}
		{
			p.SetState(128)
			p.expression(0)
		}
		{
			p.SetState(129)
			p.Match(gengineParserRR_BRACKET)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(141)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_expression)
				p.SetState(133)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(134)
					p.ComparisonOperator()
				}
				{
					p.SetState(135)
					p.expression(5)
				}


			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_expression)
				p.SetState(137)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(138)
					p.LogicalOperator()
				}
				{
					p.SetState(139)
					p.expression(4)
				}

			}

		}
		p.SetState(145)
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
	p.SetState(152)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gengineParserT__4, gengineParserTRUE, gengineParserFALSE, gengineParserSIMPLENAME, gengineParserINT, gengineParserMINUS, gengineParserDQUOTA_STRING, gengineParserDOTTEDNAME, gengineParserREAL_LITERAL:
		{
			p.SetState(147)
			p.ExpressionAtom()
		}


	case gengineParserLR_BRACKET:
		{
			p.SetState(148)
			p.Match(gengineParserLR_BRACKET)
		}
		{
			p.SetState(149)
			p.mathExpression(0)
		}
		{
			p.SetState(150)
			p.Match(gengineParserRR_BRACKET)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(162)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_mathExpression)
				p.SetState(154)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(155)
					p.MathMdOperator()
				}
				{
					p.SetState(156)
					p.mathExpression(5)
				}


			case 2:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_mathExpression)
				p.SetState(158)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(159)
					p.MathPmOperator()
				}
				{
					p.SetState(160)
					p.mathExpression(4)
				}

			}

		}
		p.SetState(166)
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

	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.MethodCall()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(168)
			p.FunctionCall()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(169)
			p.Constant()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(170)
			p.MapVar()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(171)
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

func (s *AssignmentContext) MathExpression() IMathExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
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

func (s *AssignmentContext) AssignOperator() IAssignOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignOperatorContext)
}

func (s *AssignmentContext) SetOperator() ISetOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetOperatorContext)
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
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(174)
			p.MapVar()
		}


	case 2:
		{
			p.SetState(175)
			p.Variable()
		}

	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gengineParserASSIGN:
		{
			p.SetState(178)
			p.AssignOperator()
		}


	case gengineParserSET:
		{
			p.SetState(179)
			p.SetOperator()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(182)
		p.mathExpression(0)
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
		p.SetState(184)
		p.Match(gengineParserT__1)
	}
	{
		p.SetState(185)
		p.expression(0)
	}
	{
		p.SetState(186)
		p.Match(gengineParserLR_BRACE)
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << gengineParserT__0) | (1 << gengineParserT__1) | (1 << gengineParserSIMPLENAME))) != 0) || _la == gengineParserDOTTEDNAME {
		{
			p.SetState(187)
			p.Statements()
		}

	}
	{
		p.SetState(190)
		p.Match(gengineParserRR_BRACE)
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(191)
				p.ElseIfStmt()
			}


		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == gengineParserT__2 {
		{
			p.SetState(197)
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
		p.SetState(200)
		p.Match(gengineParserT__2)
	}
	{
		p.SetState(201)
		p.Match(gengineParserT__1)
	}
	{
		p.SetState(202)
		p.expression(0)
	}
	{
		p.SetState(203)
		p.Match(gengineParserLR_BRACE)
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << gengineParserT__0) | (1 << gengineParserT__1) | (1 << gengineParserSIMPLENAME))) != 0) || _la == gengineParserDOTTEDNAME {
		{
			p.SetState(204)
			p.Statements()
		}

	}
	{
		p.SetState(207)
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
		p.SetState(209)
		p.Match(gengineParserT__2)
	}
	{
		p.SetState(210)
		p.Match(gengineParserLR_BRACE)
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << gengineParserT__0) | (1 << gengineParserT__1) | (1 << gengineParserSIMPLENAME))) != 0) || _la == gengineParserDOTTEDNAME {
		{
			p.SetState(211)
			p.Statements()
		}

	}
	{
		p.SetState(214)
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

	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(216)
			p.BooleanLiteral()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Integer()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(218)
			p.RealLiteral()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(219)
			p.StringLiteral()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(220)
			p.AtName()
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
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(223)
			p.Constant()
		}


	case 2:
		{
			p.SetState(224)
			p.Variable()
		}


	case 3:
		{
			p.SetState(225)
			p.FunctionCall()
		}


	case 4:
		{
			p.SetState(226)
			p.MethodCall()
		}


	case 5:
		{
			p.SetState(227)
			p.MapVar()
		}

	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == gengineParserT__3 {
		{
			p.SetState(230)
			p.Match(gengineParserT__3)
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(231)
				p.Constant()
			}


		case 2:
			{
				p.SetState(232)
				p.Variable()
			}


		case 3:
			{
				p.SetState(233)
				p.FunctionCall()
			}


		case 4:
			{
				p.SetState(234)
				p.MethodCall()
			}


		case 5:
			{
				p.SetState(235)
				p.MapVar()
			}

		}


		p.SetState(242)
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
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == gengineParserMINUS {
		{
			p.SetState(243)
			p.Match(gengineParserMINUS)
		}

	}
	{
		p.SetState(246)
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
		p.SetState(253)
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
		p.SetState(255)
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
		p.SetState(257)
		p.Match(gengineParserSIMPLENAME)
	}
	{
		p.SetState(258)
		p.Match(gengineParserLR_BRACKET)
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << gengineParserT__4) | (1 << gengineParserTRUE) | (1 << gengineParserFALSE) | (1 << gengineParserSIMPLENAME) | (1 << gengineParserINT) | (1 << gengineParserMINUS))) != 0) || ((((_la - 39)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 39))) & ((1 << (gengineParserDQUOTA_STRING - 39)) | (1 << (gengineParserDOTTEDNAME - 39)) | (1 << (gengineParserREAL_LITERAL - 39)))) != 0) {
		{
			p.SetState(259)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(262)
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
		p.SetState(264)
		p.Match(gengineParserDOTTEDNAME)
	}
	{
		p.SetState(265)
		p.Match(gengineParserLR_BRACKET)
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << gengineParserT__4) | (1 << gengineParserTRUE) | (1 << gengineParserFALSE) | (1 << gengineParserSIMPLENAME) | (1 << gengineParserINT) | (1 << gengineParserMINUS))) != 0) || ((((_la - 39)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 39))) & ((1 << (gengineParserDQUOTA_STRING - 39)) | (1 << (gengineParserDOTTEDNAME - 39)) | (1 << (gengineParserREAL_LITERAL - 39)))) != 0) {
		{
			p.SetState(266)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(269)
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
		p.SetState(271)
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
		p.SetState(273)
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
		p.SetState(275)
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
		p.SetState(277)
		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << gengineParserEQUALS) | (1 << gengineParserGT) | (1 << gengineParserLT) | (1 << gengineParserGTE) | (1 << gengineParserLTE) | (1 << gengineParserNOTEQUALS))) != 0)) {
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
		p.SetState(279)
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
		p.SetState(281)
		p.Match(gengineParserASSIGN)
	}



	return localctx
}


// ISetOperatorContext is an interface to support dynamic dispatch.
type ISetOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetOperatorContext differentiates from other interfaces.
	IsSetOperatorContext()
}

type SetOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetOperatorContext() *SetOperatorContext {
	var p = new(SetOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_setOperator
	return p
}

func (*SetOperatorContext) IsSetOperatorContext() {}

func NewSetOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetOperatorContext {
	var p = new(SetOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_setOperator

	return p
}

func (s *SetOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *SetOperatorContext) SET() antlr.TerminalNode {
	return s.GetToken(gengineParserSET, 0)
}

func (s *SetOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SetOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterSetOperator(s)
	}
}

func (s *SetOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitSetOperator(s)
	}
}

func (s *SetOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitSetOperator(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *gengineParser) SetOperator() (localctx ISetOperatorContext) {
	localctx = NewSetOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, gengineParserRULE_setOperator)

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
		p.SetState(283)
		p.Match(gengineParserSET)
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
	p.EnterRule(localctx, 62, gengineParserRULE_notOperator)

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
		p.SetState(285)
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
	p.EnterRule(localctx, 64, gengineParserRULE_mapVar)

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
		p.SetState(287)
		p.Variable()
	}
	{
		p.SetState(288)
		p.Match(gengineParserLSQARE)
	}
	p.SetState(292)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gengineParserINT, gengineParserMINUS:
		{
			p.SetState(289)
			p.Integer()
		}


	case gengineParserDQUOTA_STRING:
		{
			p.SetState(290)
			p.StringLiteral()
		}


	case gengineParserSIMPLENAME, gengineParserDOTTEDNAME:
		{
			p.SetState(291)
			p.Variable()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(294)
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
	p.EnterRule(localctx, 66, gengineParserRULE_atName)

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
		p.SetState(296)
		p.Match(gengineParserT__4)
	}



	return localctx
}


func (p *gengineParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
			var t *ExpressionContext = nil
			if localctx != nil { t = localctx.(*ExpressionContext) }
			return p.Expression_Sempred(t, predIndex)

	case 10:
			var t *MathExpressionContext = nil
			if localctx != nil { t = localctx.(*MathExpressionContext) }
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

