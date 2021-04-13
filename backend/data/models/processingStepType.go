package data

type ProcessingStepType int

const (
	Resize          = 10
	Fit             = 11
	Fill            = 12
	Crop            = 13
	CropAnchor      = 14
	CropCenter      = 15
	Thumbnail       = 16
	Blur            = 20
	Sharpen         = 21
	Gamma           = 30
	Contrast        = 31
	ContrastSigmoid = 32
	Brightness      = 33
	Saturation      = 34
	Hue             = 35
	Grayscale       = 36
	Invert          = 37
	Convolve3x3     = 40
	Convolve5x5     = 41
	FlipH           = 50
	FlipV           = 51
	Rotate          = 52
	Rotate90        = 53
	Rotate180       = 54
	Rotate270       = 55
	Transpose       = 56
	Transverse      = 57
	Overlay         = 60
	OverlayCenter   = 61
	Paste           = 62
	PasteCenter     = 63
)
