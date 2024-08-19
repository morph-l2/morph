package localpool

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/morph-l2/go-ethereum/core/types"
)

// The current `contentfrom` method of `blobpool` returns empty,
// so we need to locally record transactions sent to the `memepool`
// to prevent losing track of the rollup progress after the `submitter` restarts.

// case
//

type Journal struct {
	path string
	mu   sync.Mutex
}

func New(path string) *Journal {
	return &Journal{path: path}
}
func (j *Journal) Init() error {
	// create file if file not exist
	_, err := os.Stat(j.path)
	if err != nil {
		if os.IsNotExist(err) {
			// create file
			_, err := os.Create(j.path)
			if err != nil {
				return fmt.Errorf("failed to create journal file: %w", err)
			}
			return nil
		} else {
			return fmt.Errorf("failed to stat journal file: %w", err)
		}

	}

	return nil
}

func (j *Journal) AppendTx(tx *types.Transaction) error {
	encTx, err := EncodeTx(tx)
	if err != nil {
		return fmt.Errorf("add tx to journal filed:%v", err)
	}
	return j.AddToFileEnd(encTx)
}

func (j *Journal) AddToFileStart(str string) error {
	j.mu.Lock()
	defer j.mu.Unlock()
	return addToFileStart(j.path, str)
}

func (j *Journal) AddToFileEnd(str string) error {
	j.mu.Lock()
	defer j.mu.Unlock()
	return addToFileEnd(j.path, str)

}

func (j *Journal) ParseAllTxsAndCleanJournal() ([]*types.Transaction, error) {
	txs, err := j.ParseAllTxs()
	if err != nil {
		return nil, fmt.Errorf("failed to parse txs: %w", err)
	}
	err = j.clean()
	if err != nil {
		return nil, fmt.Errorf("failed to clean journal file: %w", err)
	}
	return txs, nil
}

func (j *Journal) ParseAllTxs() ([]*types.Transaction, error) {
	content, err := readFileContent(j.path)

	if err != nil {
		return nil, fmt.Errorf("failed to parse txs: %w", err)
	}

	if len(content) == 0 {
		return nil, nil
	}

	var ans []*types.Transaction
	lines := getLines(content)
	for _, line := range lines {
		// parse tx
		tx, err := ParseTx(line)
		if err != nil {
			return nil, fmt.Errorf("failed to parse tx: %w", err)
		}
		ans = append(ans, tx)
	}

	return ans, nil

}

func addToFileStart(path string, str string) error {
	content, err := readFileContent(path)
	if err != nil {
		return fmt.Errorf("failed to add line to file start: %w", err)
	}
	var newContent string
	if len(content) > 0 {
		newContent = str + "\n" + content
	} else {
		newContent = str
	}

	err = os.WriteFile(path, []byte(newContent), 0600)
	if err != nil {
		return fmt.Errorf("failed to write journal file: %w", err)
	}
	return nil
}

func (j *Journal) rm() error {
	// file if exist

	// clean journal file
	err := os.Remove(j.path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			return fmt.Errorf("failed to remove journal file: %w", err)
		}
	}
	return nil
}

func (j *Journal) Clean() error {
	j.mu.Lock()
	defer j.mu.Unlock()
	return j.clean()
}

func (j *Journal) clean() error {
	// clean content in journal
	err := os.Truncate(j.path, 0)
	if err != nil {
		return fmt.Errorf("failed to clean journal file: %w", err)
	}
	return nil
}

func addToFileEnd(path string, str string) error {
	// add a line to file end
	content, err := readFileContent(path)
	if err != nil {
		return fmt.Errorf("failed to add line to file end: %w", err)
	}
	var newContent string
	if len(content) > 0 {
		newContent = content + "\n" + str
	} else {
		newContent = str
	}
	err = os.WriteFile(path, []byte(newContent), 0600)
	if err != nil {
		return fmt.Errorf("failed to add line to file end: %w", err)
	}
	return nil
}

func (j *Journal) GetFirstLine() (string, error) {
	j.mu.Lock()
	defer j.mu.Unlock()
	return getFirstLine(j.path)
}

func (j *Journal) GetLastLine() (string, error) {
	j.mu.Lock()
	defer j.mu.Unlock()
	return getLastLine(j.path)
}

func readFileContent(path string) (string, error) {
	content, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return "", fmt.Errorf("failed to read journal file: %w", err)
	}
	return string(content), nil
}
func getLines(str string) []string {
	return strings.Split(str, "\n")
}
func getFirstLine(path string) (string, error) {

	content, err := readFileContent(path)
	if err != nil {
		return "", fmt.Errorf("failed to get first line: %w", err)
	}

	// split content by line
	lines := getLines(content)

	// first line
	if len(lines) > 0 {
		return lines[0], nil
	}

	return "", nil

}

func getLastLine(path string) (string, error) {
	content, err := readFileContent(path)
	if err != nil {
		return "", fmt.Errorf("failed to get last line: %w", err)
	}
	lines := getLines(content)
	if len(lines) > 0 {
		return lines[len(lines)-1], nil
	}

	return "", nil
}
