package auth

import (
 "testing"
 "errors"
 "net/http"
)

func TestGetAPIKey(t *testing.T){
  tests := map[string]struct {
     input http.Header
     wantedString string
     wantedErr error
  }{
    "No auth header": {input: http.Header{}, wantedString: "", wantedErr: errors.New("no authorization header included")},
    "Empty auth header": {input: http.Header{"Authorization": []string{"anotherthing 12345"}}, wantedString: "", wantedErr: errors.New("malformed authorization header")},
    "Correct auth header": {input: http.Header{"Authorization": []string{"ApiKey 12345"}}, wantedString: "12345", wantedErr: nil},
  }

  for name, tc := range tests {
    t.Run(name, func(t *testing.T) {
      got, gotErr := GetAPIKey(tc.input)
      if got != tc.wantedString || (gotErr != nil && gotErr.Error() != tc.wantedErr.Error() ) {
      t.Fatalf("%s: expected string: %v, got: %v, expected err: %v, gotErr: %v", name, tc.wantedString, got, tc.wantedErr, gotErr)
    }
    })
  }
}

