// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// findings contains example snippets for working with findings
// and their parent resource "sources".
package findings

// [START list_all_findings]
import (
	"context"
	"fmt"
	"io"

	securitycenter "cloud.google.com/go/securitycenter/apiv1"
	"google.golang.org/api/iterator"
	securitycenterpb "google.golang.org/genproto/googleapis/cloud/securitycenter/v1"
)

// listFindings prints all findings in orgID to w and returns the count of
// findings found.  orgID is the numeric identifier of the organization.
func listFindings(w io.Writer, orgID string) (int, error) {
	// orgID := "12321311"
	// Instantiate a context and a security service client to make API calls.
	ctx := context.Background()
	client, err := securitycenter.NewClient(ctx)
	if err != nil {
		return -1, fmt.Errorf("Error instantiating client %v\n", err)
	}
	defer client.Close() // Closing the client safely cleans up background resources.

	req := &securitycenterpb.ListFindingsRequest{
		// List findings across all sources.
		Parent: fmt.Sprintf("organizations/%s/source/-", orgID),
	}
	findingsFound := 0
	it := client.ListFindings(ctx, req)
	for {
		finding_result, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return -1, fmt.Errorf("Error listing sources: %v", err)
		}
		finding := finding_result.Finding
		fmt.Fprintf(w, "Finding Name: %s, ", finding.Name)
		fmt.Fprintf(w, "Resource Name %s, ", finding.ResourceName)
		fmt.Fprintf(w, "Category: %s\n", finding.Category)
		findingsFound++
	}
	return findingsFound, nil
}

// [END list_all_findings]