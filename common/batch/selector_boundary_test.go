package batch

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductionCodeDoesNotCallCompactCommitParser(t *testing.T) {
	packageDir, err := os.Getwd()
	require.NoError(t, err)

	entries, err := os.ReadDir(packageDir)
	require.NoError(t, err)

	files := token.NewFileSet()
	var references []string
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") {
			continue
		}

		path := filepath.Join(packageDir, name)
		file, err := parser.ParseFile(files, path, nil, 0)
		require.NoError(t, err, "parse production source %s", name)

		definitions := make(map[token.Pos]struct{})
		for _, declaration := range file.Decls {
			function, ok := declaration.(*ast.FuncDecl)
			if ok && function.Name.Name == "parseCommitBatchTxData" {
				definitions[function.Name.Pos()] = struct{}{}
			}
		}

		ast.Inspect(file, func(node ast.Node) bool {
			identifier, ok := node.(*ast.Ident)
			if !ok || identifier.Name != "parseCommitBatchTxData" {
				return true
			}
			if _, isDefinition := definitions[identifier.Pos()]; isDefinition {
				return true
			}

			position := files.Position(identifier.Pos())
			references = append(references, fmt.Sprintf("%s:%d", filepath.Base(position.Filename), position.Line))
			return true
		})
	}

	require.Empty(t, references,
		"production batch code must not reference the test-only compact commit parser")
}
