package bigfix

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// shouldIgnoreErrors returns an ErrorPredicate for BigFix API calls
// Following the same pattern as AWS plugin
// This function combines default "not found" errors with any additional
// error messages specified in the configuration's ignore_error_messages parameter
func shouldIgnoreErrors(notFoundErrors []string) plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		if err == nil {
			return false
		}

		errorStr := strings.ToLower(err.Error())

		// Get configuration to check for additional ignore error messages
		config := GetConfig(d.Connection)

		// Combine default "not found" errors with configured ignore error messages
		allIgnorePatterns := append([]string{}, notFoundErrors...)
		if config.IgnoreErrorMessages != nil {
			allIgnorePatterns = append(allIgnorePatterns, config.IgnoreErrorMessages...)
		}

		for _, pattern := range allIgnorePatterns {
			if strings.Contains(errorStr, strings.ToLower(pattern)) {
				return true
			}
		}

		return false
	}
}
