package main

import (
	"errors"
	"os"
	"testing"
)

type mockFileReader struct {
	content []byte
	err     error
}

func (m *mockFileReader) ReadFile(filename string) ([]byte, error) {
	return m.content, m.err
}

type mockFileWriter struct {
	err error
}

func (m *mockFileWriter) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return m.err
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name           string
		inputContent   string
		expectedOutput string
		readerErr      error
		writerErr      error
		expectedErr    bool
	}{
		{
			name:           "Successful reversal",
			inputContent:   "hello",
			expectedOutput: "olleh",
		},
		{
			name:        "Reader error",
			readerErr:   errors.New("read error"),
			expectedErr: true,
		},
		{
			name:         "Writer error",
			inputContent: "hello",
			writerErr:    errors.New("write error"),
			expectedErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := &mockFileReader{
				content: []byte(tt.inputContent),
				err:     tt.readerErr,
			}
			writer := &mockFileWriter{
				err: tt.writerErr,
			}

			err := reverse("input.txt", "output.txt", reader, writer)
			if tt.expectedErr {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})
	}
}
