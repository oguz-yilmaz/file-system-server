package fsmod

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	Conf "github.com/oguz-yilmaz/file-system-server/pkg/config"
	"github.com/oguz-yilmaz/file-system-server/pkg/protocol"
)

type MockTransportChannel struct {
	WrittenResponses []any
}

func (m *MockTransportChannel) Read(req *protocol.Request) (*jsoniter.Decoder, error) {
	return nil, nil
}

func (m *MockTransportChannel) Write(res any) error {
	m.WrittenResponses = append(m.WrittenResponses, res)
	return nil
}

func (m *MockTransportChannel) Close() error {
	return nil
}

func TestHandleCreateFile(t *testing.T) {
	tempDir := t.TempDir()

	jsonParams := createFileParamsJsonString(map[string]any{
		"name":      "testfile.txt",
		"content":   "This is an example file.",
		"file-type": "txt",
		"overwrite": true,
	}, tempDir)

	req := protocol.Request{
		ID:     1,
		RPC:    "2.0",
		Method: protocol.METHOD_CREATE_FILE,
		Params: jsoniter.RawMessage(jsonParams),
	}

	mockChannel := &MockTransportChannel{}

	HandleCreateFile(req, mockChannel, Conf.NewDefaultConfig())

	if len(mockChannel.WrittenResponses) != 1 {
		t.Fatalf("expected 1 response, got %d", len(mockChannel.WrittenResponses))
	}

	response, ok := mockChannel.WrittenResponses[0].(*CreateFileSuccessResponse)
	if !ok {
		t.Fatalf("expected CreateFileSuccessResponse, got %T", mockChannel.WrittenResponses[0])
	}

	if response.Result.Name != "testfile.txt" {
		t.Errorf("expected file name 'testfile.txt', got '%s'", response.Result.Name)
	}

	if string(response.Result.Content) != "This is an example file." {
		t.Errorf("expected file content 'This is a test file.', got '%s'", response.Result.Content)
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
