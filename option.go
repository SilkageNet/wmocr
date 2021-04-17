package wmocr

//RetType recognize result type.
type RetType byte

const (
	//RetOnlyCode only return code.
	RetOnlyCode RetType = iota
	//RetCodeAndRange return code and range, such as: S,10,11,12,13|A,1,2,3,4.
	RetCodeAndRange
)

//SegmentationType segmentation type.
type SegmentationType byte

const (
	//SegmentationWhole overall identification.
	SegmentationWhole SegmentationType = iota
	//SegmentationConnected connected partition identification.
	SegmentationConnected
	//SegmentationVertical vertical segmentation identification.
	SegmentationVertical
	//SegmentationHorizontal horizontal segmentation identification.
	SegmentationHorizontal
	//SegmentationVerticalAndHorizontal vertical and horizontal segmentation identification.
	SegmentationVerticalAndHorizontal
)

// RecognizeType recognize type.
type RecognizeType byte

const (
	//RecognizeWord recognize word.
	RecognizeWord RecognizeType = iota
	//RecognizeImage recognize image.
	RecognizeImage
)

type option struct {
	RetType          RetType
	SegmentationType SegmentationType
	RecognizeType    RecognizeType
	AccelerationType byte
	AccelerationRet  byte
	MinSimilarity    byte
	CharSpace        int
}

type OptionFn func(o *option)

func WithRetType(rType RetType) OptionFn {
	return func(o *option) {
		o.RetType = rType
	}
}

func WithSegmentationType(sType SegmentationType) OptionFn {
	return func(o *option) {
		o.SegmentationType = sType
	}
}

func WithRecognizeType(rType RecognizeType) OptionFn {
	return func(o *option) {
		o.RecognizeType = rType
	}
}

func WithAccelerationType() OptionFn {
	return func(o *option) {
		o.RecognizeType = 1
	}
}

func WithAccelerationRet() OptionFn {
	return func(o *option) {
		o.AccelerationRet = 1
	}
}

func WithMinSimilarity(v byte) OptionFn {
	return func(o *option) {
		o.MinSimilarity = v
	}
}

func WithCharSpace(space int) OptionFn {
	return func(o *option) {
		o.CharSpace = space
	}
}
