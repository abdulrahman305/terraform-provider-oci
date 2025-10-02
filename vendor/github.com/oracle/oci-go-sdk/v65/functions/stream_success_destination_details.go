// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Functions Service API
//
// API for the Functions service.
//

package functions

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StreamSuccessDestinationDetails The destination stream in the Streaming service to which to send the response of the successful detached function invocation.
// Example: `{"kind": "STREAM", "streamId": "stream_OCID"}`
type StreamSuccessDestinationDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stream.
	StreamId *string `mandatory:"true" json:"streamId"`
}

func (m StreamSuccessDestinationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StreamSuccessDestinationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m StreamSuccessDestinationDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeStreamSuccessDestinationDetails StreamSuccessDestinationDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeStreamSuccessDestinationDetails
	}{
		"STREAM",
		(MarshalTypeStreamSuccessDestinationDetails)(m),
	}

	return json.Marshal(&s)
}
