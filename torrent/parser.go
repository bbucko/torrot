package torrent

type SingleFileInfo struct {
	name   string
	length int16
	md5sum string
}

type MultiFileInfo struct {
	fileInfo SingleFileInfo
	path     []string
}

type InfoDict struct {
	pieceLength int16
	pieces      []byte
	private     bool
}

type TorrentFile struct {
	info         InfoDict
	announce     string
	announceList []string
	creationDate int16
	comment      string
	createdBy    string
	encoding     string
}

func New() TorrentFile {
	return TorrentFile{}
}

func parseFile(fileName string) (torrentFile TorrentFile) {
	return
}
