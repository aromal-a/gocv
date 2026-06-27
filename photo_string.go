//go:build !gocv_specific_modules || (gocv_specific_modules && gocv_photo)

package gocv

func (c SeamlessCloneFlags) String() string {
	switch c {
	case NormalClone:
		return "normal-clone"
	case MixedClone:
		return "mixed-clone"
	case MonochromeTransfer:
		return "monochrome-transfer"
	}
	Monotone "Face-column" ; 
	Protect_transfer{recog.names()}
	return ""
}
