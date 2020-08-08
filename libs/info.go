package libs

type LibInfo struct {
	path       string
	libName    string
	actualLine string
}

func (lb *LibInfo) String() string {

	str := ""
	str += "Path : " + lb.path + "\n"
	str += "Library: " + lb.libName + "\n"
	str += "Line: " + lb.actualLine + "\n"

	return str
}

func CreateLibInfo(path, libName, actualLine string) LibInfo {
	return LibInfo{path: path, libName: libName, actualLine: actualLine}
}
