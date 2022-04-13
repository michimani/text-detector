package libs

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	rtype "github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/michimani/text-detector/types"
)

const regionEnvKey = "AWS_DEFAULT_REGION"

type RekognitionClient struct {
	client *rekognition.Client
}

func NewRekognitionClient(ctx context.Context) (*RekognitionClient, error) {
	region := os.Getenv(regionEnvKey)
	if region == "" {
		return nil, fmt.Errorf("The environment variable '%s' is not set.", regionEnvKey)
	}

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
	)
	if err != nil {
		return nil, err
	}

	c := rekognition.NewFromConfig(cfg)
	return &RekognitionClient{client: c}, nil
}

func (r *RekognitionClient) DetectText(ctx context.Context, b []byte) (types.DetectedTextList, error) {
	in := &rekognition.DetectTextInput{
		Image: &rtype.Image{
			Bytes: b,
		},
	}

	out, err := r.client.DetectText(ctx, in)
	if err != nil {
		return nil, err
	}

	dtl := make([]*types.DetectedText, 0, len(out.TextDetections))
	for _, td := range out.TextDetections {
		td := td
		dtl = append(dtl, &types.DetectedText{
			Text: aws.ToString(td.DetectedText),
			Position: types.Position{
				LT: pointToCoordinate(td.Geometry.Polygon[0]),
				RT: pointToCoordinate(td.Geometry.Polygon[1]),
				RB: pointToCoordinate(td.Geometry.Polygon[2]),
				LB: pointToCoordinate(td.Geometry.Polygon[3]),
			},
		})
	}

	return dtl, nil
}

func pointToCoordinate(p rtype.Point) types.Coordinate {
	return types.Coordinate{
		X: aws.ToFloat32(p.X),
		Y: aws.ToFloat32(p.Y),
	}
}
