package detector

import (
	"context"
	"errors"

	"github.com/michimani/text-detector/libs"
	"github.com/michimani/text-detector/types"
)

// IDetectTextClient is an interface that represents a client that actually detects strings from images.
// This library generates clients that use Amazon Rekogntion or Google Cloud Vision.
type IDetectTextClient interface {
	DetectText(ctx context.Context, b []byte) (types.DetectedTextList, error)
}

// TextDetector is a structure representing a client for detecting strings in images.
type TextDetector struct {
	client IDetectTextClient
}

// ClientType is a string representing a client type you want to use.
type ClientType string

var (
	ClientTypeRekognition ClientType = "rekognition"
	ClientTypeCloudVision ClientType = "cloudvision"
	ClientTypeCustom      ClientType = "custom"
)

// NewInput is a structure for generating new TextDetector.
type NewInput struct {
	// Type of client that detects text. (required)
	ClientType ClientType
	// Custom client. If ClientType is ClientTypeCustom, this field is required.
	CustomClient IDetectTextClient
}

// New generate new TextDetector.
func New(ctx context.Context, in *NewInput) (*TextDetector, error) {
	var client IDetectTextClient
	switch in.ClientType {
	case ClientTypeRekognition:
		c, err := libs.NewRekognitionClient(context.Background())
		if err != nil {
			return nil, err
		}
		client = c
	case ClientTypeCloudVision:
		c, err := libs.NewCloudVisionClient(context.Background())
		if err != nil {
			return nil, err
		}
		client = c
	case ClientTypeCustom:
		if in.CustomClient == nil {
			return nil, errors.New("If you want to use custom client, NewInput.Client is required.")
		}
		client = in.CustomClient
	}

	return &TextDetector{
		client: client,
	}, nil
}

// DetectText detects a string from a sequence of image bytes given as an argument
// and returns the result in a pointer slice of types.DetectedText.
func (td *TextDetector) DetectText(ctx context.Context, b []byte) (types.DetectedTextList, error) {
	return td.client.DetectText(ctx, b)
}
