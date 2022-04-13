package libs

import (
	"bytes"
	"context"
	"image"

	visionapi "cloud.google.com/go/vision/apiv1"
	"github.com/michimani/text-detector/types"
	vision "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

type CloudVisionClient struct {
	client *visionapi.ImageAnnotatorClient
}

func NewCloudVisionClient(ctx context.Context) (*CloudVisionClient, error) {
	c, err := visionapi.NewImageAnnotatorClient(ctx)
	if err != nil {
		return nil, err
	}

	return &CloudVisionClient{client: c}, nil
}

func (c *CloudVisionClient) DetectText(ctx context.Context, b []byte) (types.DetectedTextList, error) {
	img, err := visionapi.NewImageFromReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	imageinfo, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	width := imageinfo.Bounds().Dx()
	height := imageinfo.Bounds().Dy()

	ictx := &vision.ImageContext{
		LanguageHints: []string{"ja"},
	}
	annotations, err := c.client.DetectTexts(ctx, img, ictx, 10)
	if err != nil {
		return nil, err
	}

	dtl := make([]*types.DetectedText, 0, len(annotations))
	for _, ann := range annotations {
		ann := ann
		dtl = append(dtl, &types.DetectedText{
			Text: ann.Description,
			Position: types.Position{
				LT: vertexToCoordinate(ann.BoundingPoly.Vertices[0], width, height),
				RT: vertexToCoordinate(ann.BoundingPoly.Vertices[1], width, height),
				RB: vertexToCoordinate(ann.BoundingPoly.Vertices[2], width, height),
				LB: vertexToCoordinate(ann.BoundingPoly.Vertices[3], width, height),
			},
		})
	}

	return dtl, nil
}

func vertexToCoordinate(vertex *vision.Vertex, baseWidth, baseHeight int) types.Coordinate {
	return types.Coordinate{
		X: float32(vertex.X) / float32(baseWidth),
		Y: float32(vertex.Y) / float32(baseHeight),
	}
}
