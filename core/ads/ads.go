package ads

const (
	// TypeAudio is Audio ad
	TypeAudio = "AD_AUDIO"

	// TypeImage is Image ad
	TypeImage = "AD_IMAGE"

	// TypeVideo is Video ad
	TypeVideo = "AD_VIDEO"
)

// Ads stores an ads info
type Ads struct {
	Type string
}
