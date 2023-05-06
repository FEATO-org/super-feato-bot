package utility

import (
	"context"
	"net/http"

	"github.com/FEATO-org/support-feato-system/config"
	"google.golang.org/api/option"
	"google.golang.org/api/script/v1"
)

// FIXME: infrastructureに移動させる
func ExecuteGASApi(ctx context.Context, client *http.Client, gas *config.GAS, devMode bool, function string, parameters []interface{}) (*script.Operation, error) {
	scriptService, err := script.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	call := scriptService.Scripts.Run(gas.ScriptID, &script.ExecutionRequest{
		DevMode:         devMode,
		Function:        function,
		Parameters:      parameters,
		SessionState:    "",
		ForceSendFields: []string{},
		NullFields:      []string{},
	})
	call.Context(ctx)
	operation, err := call.Do()
	if err != nil {
		return nil, err
	}

	return operation, nil
}

// 文字配列をユニークにして返す
func StringArrayUnique(arr []string) []string {
	array := make(map[string]struct{})
	unique := []string{}

	for _, ele := range arr {
		array[ele] = struct{}{}
	}
	for arr := range array {
		unique = append(unique, arr)
	}

	return unique
}

const NO_MATCH_RESULT = "sql: no rows in result set"
