package detector_test

import (
	"context"
	"testing"

	"github.com/michimani/text-detector/detector"
	"github.com/michimani/text-detector/types"
	"github.com/stretchr/testify/assert"
)

type mockClient struct{}

func (c mockClient) DetectText(ctx context.Context, b []byte) (types.DetectedTextList, error) {
	return nil, nil
}

func Test_New(t *testing.T) {
	cases := []struct {
		name    string
		in      *detector.NewInput
		setEnv  bool
		wantErr bool
	}{
		{
			name: "ok: Amazon Rekognition",
			in: &detector.NewInput{
				ClientType: detector.ClientTypeRekognition,
			},
			setEnv:  true,
			wantErr: false,
		},
		{
			name: "ok: Custom Client",
			in: &detector.NewInput{
				ClientType:   detector.ClientTypeCustom,
				CustomClient: mockClient{},
			},
			wantErr: false,
		},
		{
			name: "ng: Amazon Rekognition",
			in: &detector.NewInput{
				ClientType: detector.ClientTypeRekognition,
			},
			setEnv:  false,
			wantErr: true,
		},
		{
			name: "ng: Google Cloud Vision",
			in: &detector.NewInput{
				ClientType: detector.ClientTypeCloudVision,
			},
			setEnv:  false,
			wantErr: true,
		},
		{
			name: "ng: Custom Client",
			in: &detector.NewInput{
				ClientType: detector.ClientTypeCustom,
			},
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)

			if c.setEnv {
				t.Setenv("AWS_DEFAULT_REGION", "ap-northeast-1")
			} else {
				t.Setenv("AWS_DEFAULT_REGION", "")
			}

			td, err := detector.New(context.Background(), c.in)
			if c.wantErr {
				asst.Error(err)
				asst.Nil(td)
				return
			}

			asst.NoError(err)
			asst.NotNil(td)
		})
	}
}
