//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package jwt

import (
	jg "github.com/dgrijalva/jwt-go"
	driver "github.com/gigaphoton/go-driver"
)

const (
	issArangod = "arangodb"
)

// CreateArangodJwtAuthorizationHeader calculates a JWT authorization header, for authorization
// of a request to an arangod server, based on the given secret.
// If the secret is empty, nothing is done.
// Use the result of this function as input for driver.RawAuthentication.
func CreateArangodJwtAuthorizationHeader(jwtSecret, serverID string) (string, error) {
	if jwtSecret == "" || serverID == "" {
		return "", nil
	}
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jg.NewWithClaims(jg.SigningMethodHS256, jg.MapClaims{
		"iss":       issArangod,
		"server_id": serverID,
	})

	// Sign and get the complete encoded token as a string using the secret
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", driver.WithStack(err)
	}

	return "bearer " + signedToken, nil
}
