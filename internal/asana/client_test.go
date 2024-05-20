package asana

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"

	"github.com/devzcraft/assignment/internal/config"
)

func TestClient_User(t *testing.T) {
	t.Parallel()

	type args struct {
		userGID string
	}

	config := &config.Config{
		Asana: config.Asana{
			Token:        "token",
			WorkspaceGID: "13241234",
			BaseURL:      "https://asana.com",
		},
		RateLimit: "150",
	}

	tests := []struct {
		name    string
		client  *Client
		args    args
		want    *resty.Response
		wantErr bool
	}{
		// TODO: add additional test cases
		{
			name:    "Successfully fetch user data",
			client:  NewClientMock(config, "test response"),
			args:    args{},
			want:    WantResponse("test response"),
			wantErr: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := tc.client.User(tc.args.userGID)
			if (err != nil) != tc.wantErr {
				t.Errorf("Client.User() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got.String() != tc.want.String() {
				t.Errorf("Client.User() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestClient_Project(t *testing.T) {
	t.Parallel()

	config := &config.Config{
		Asana: config.Asana{
			Token:        "token",
			WorkspaceGID: "13241234",
			BaseURL:      "https://asana.com",
		},
		RateLimit: "150",
	}
	type args struct {
		projectGID string
	}
	tests := []struct {
		name    string
		client  *Client
		args    args
		want    *resty.Response
		wantErr bool
	}{
		{
			name:    "Successfully fetch project data",
			client:  NewClientMock(config, "test response"),
			args:    args{},
			want:    WantResponse("test response"),
			wantErr: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := tc.client.Project(tc.args.projectGID)
			if (err != nil) != tc.wantErr {
				t.Errorf("Client.Project() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got.String() != tc.want.String() {
				t.Errorf("Client.Project() = %v, want %v", got, tc.want)
			}
		})
	}
}

func NewClientMock(config *config.Config, response string) *Client {
	client := NewClient(config)
	httpmock.ActivateNonDefault(client.HTTP().GetClient())
	httpmock.RegisterResponder("GET", `=~^https://asana.com?.+\z`,
		httpmock.NewStringResponder(200, response))

	return client
}

func WantResponse(response string) *resty.Response {
	resp := &resty.Response{}
	resp.SetBody([]byte(response))

	return resp
}
