// MIT License
//
// Copyright (c) 2024 - Yuuki Takezawa
// Copyright (c) 2022 - present Open Swoole Group
// Copyright (c) 2018 SpiralScout
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package language

import (
	"bytes"
	"strings"
	"unicode"

	"google.golang.org/protobuf/types/descriptorpb"
)

var reservedKeywords = [...]string{
	"__halt_compiler",
	"abstract",
	"and",
	"array",
	"as",
	"break",
	"callable",
	"case",
	"catch",
	"class",
	"clone",
	"const",
	"continue",
	"declare",
	"default",
	"die",
	"do",
	"echo",
	"else",
	"elseif",
	"empty",
	"enddeclare",
	"endfor",
	"endforeach",
	"endif",
	"endswitch",
	"endwhile",
	"eval",
	"exit",
	"extends",
	"final",
	"for",
	"foreach",
	"function",
	"global",
	"goto",
	"if",
	"implements",
	"include",
	"include_once",
	"instanceof",
	"insteadof",
	"interface",
	"isset",
	"list",
	"namespace",
	"new",
	"or",
	"print",
	"private",
	"protected",
	"public",
	"require",
	"require_once",
	"return",
	"static",
	"switch",
	"throw",
	"trait",
	"try",
	"unset",
	"use",
	"var",
	"while",
	"xor",
	"int",
	"float",
	"bool",
	"string",
	"true",
	"false",
	"null",
	"void",
	"iterable",
	"yield",
	"match",
}

type PHP struct {
}

// isReserved check if the name is a reserved keyword
func (_ PHP) isReserved(name string) bool {
	name = strings.ToLower(name)
	for _, k := range reservedKeywords {
		if name == k {
			return true
		}
	}
	return false
}

// Identifier snake_case to CamelCase
func (p PHP) Identifier(name string, suffix string) string {
	name = p.Camelize(name)
	if suffix != "" {
		return name + p.Camelize(suffix)
	}
	return name
}

func (p PHP) resolveReserved(identifier string, pkg string) string {
	if p.isReserved(strings.ToLower(identifier)) {
		if pkg == ".google.protobuf" {
			return "GPB" + identifier
		}
		return "PB" + identifier
	}

	return identifier
}

// Camelize "dino_party" -> "DinoParty"
func (p PHP) Camelize(word string) string {
	words := p.splitCamelCaseWords(word)
	return strings.Join(words, "")
}

func (_ PHP) splitCamelCaseWords(input string) []string {
	words := make([]string, 0)
	wordInRunes := make([]rune, 0)
	for _, inputRune := range input {
		isSpaceCharacter := isSpacerChar(inputRune)
		wordInRunes = buildAndAppendWords(wordInRunes, inputRune, isSpaceCharacter, &words)
	}
	words = append(words, string(wordInRunes))
	return words
}

func buildAndAppendWords(wordInRunes []rune, inputRune rune, isSpaceCharacter bool, words *[]string) []rune {
	if len(wordInRunes) > 0 {
		if unicode.IsUpper(inputRune) || isSpaceCharacter {
			*words = append(*words, string(wordInRunes))
			wordInRunes = make([]rune, 0)
		}
	}
	if !isSpaceCharacter {
		if len(wordInRunes) > 0 {
			wordInRunes = append(wordInRunes, unicode.ToLower(inputRune))
		} else {
			wordInRunes = append(wordInRunes, unicode.ToUpper(inputRune))
		}
	}
	return wordInRunes
}

func isSpacerChar(c rune) bool {
	switch {
	case c == rune("_"[0]):
		return true
	case c == rune(" "[0]):
		return true
	case c == rune(":"[0]):
		return true
	case c == rune("-"[0]):
		return true
	}
	return false
}

func (p PHP) Namespace(pkg *string, sep string) string {
	if pkg == nil {
		return ""
	}
	result := bytes.NewBuffer(nil)
	for _, pk := range strings.Split(*pkg, ".") {
		result.WriteString(p.Identifier(pk, ""))
		result.WriteString(sep)
	}

	return strings.Trim(result.String(), sep)
}

func (p PHP) DetectNamespace(file *descriptorpb.FileDescriptorProto) string {
	ns := p.Namespace(file.Package, "/")
	if file.Options != nil && file.Options.PhpNamespace != nil {
		ns = strings.ReplaceAll(*file.Options.PhpNamespace, `\`, `/`)
	}
	return ns
}
