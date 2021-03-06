package vpx

func (fmt ImageFormat) String() string {
	switch fmt {
	case ImageFormatNone:
		return "NONE"
	case ImageFormatRgb24:
		return "RGB24"
	case ImageFormatRgb32:
		return "RGB32"
	case ImageFormatRgb565:
		return "RGB565"
	case ImageFormatRgb555:
		return "RGB555"
	case ImageFormatUyvy:
		return "UYVY"
	case ImageFormatYuy2:
		return "YUY2"
	case ImageFormatYvyu:
		return "YVYU"
	case ImageFormatBgr24:
		return "BGR24"
	case ImageFormatRgb32Le:
		return "RGB32_LE"
	case ImageFormatArgb:
		return "ARGB"
	case ImageFormatArgbLe:
		return "ARGB_LE"
	case ImageFormatRgb565Le:
		return "RGB565_LE"
	case ImageFormatRgb555Le:
		return "RGB555_LE"
	case ImageFormatYv12:
		return "YV12"
	case ImageFormatI420:
		return "I420"
	case ImageFormatVpxyv12:
		return "VPXYV12"
	case ImageFormatVpxi420:
		return "VPXI420"
	case ImageFormatI422:
		return "I422"
	case ImageFormatI444:
		return "I444"
	case ImageFormatI440:
		return "I440"
	case ImageFormat444a:
		return "444A"
	case ImageFormatI42016:
		return "I42016"
	case ImageFormatI42216:
		return "I42216"
	case ImageFormatI44416:
		return "I44416"
	case ImageFormatI44016:
		return "I44016"
	}
	return ""
}
