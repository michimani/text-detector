package types

// DetectedText is a structure of detected text.
type DetectedText struct {
	// The detected string.
	Text string
	// The value of the coordinates expressed as a rate of the image size.
	Position Position
}

// Position is a structure representing the coordinates of
// each vertex of the rectangular area containing each detected string.
//    LT (0,0)       RT (1,0)
//     +-------------+
//     |             |
//     |  text area  |
//     |             |
//     +-------------+
//    LB (0,1)       RB (1,1)
type Position struct {
	// Left Top
	LT Coordinate
	// Right Top
	RT Coordinate
	// Right Bottom
	RB Coordinate
	// Left Bottom
	LB Coordinate
}

// Coordinate is a structure that represents the X and Y coordinates of each vertex.
type Coordinate struct {
	X float32
	Y float32
}

// IsContainedIn determines whether the detected string is contained within the specified Position.
// example: true
//        LT (0,0)                       RT (10,0)
//  target +------------------------------+
//         |                              |
//         |     LT (1,1)      RT (9,1)   |
//         |   dt +-------------+         |
//         |      |             |         |
//         |      |  text area  |         |
//         |      |             |         |
//         |      +-------------+         |
//         |     LB (1,9)      RB (9,9)   |
//         |                              |
//         +------------------------------+
//        LB (0,10)                      RB (10,10)
func (dt *DetectedText) IsContainedIn(target Position) bool {
	if dt == nil {
		return false
	}

	return dt.Position.LT.X > target.LT.X && dt.Position.LT.Y > target.LT.Y &&
		dt.Position.RT.X < target.RT.X && dt.Position.RT.Y > target.RT.Y &&
		dt.Position.RB.X < target.RB.X && dt.Position.RB.Y < target.RB.Y &&
		dt.Position.LB.X > target.LB.X && dt.Position.LB.Y < target.LB.Y
}

// DetectedTextList is a structure representing a slice of detected string information.
type DetectedTextList []*DetectedText
