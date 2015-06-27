package torrent

import (
	"github.com/marksamman/bencode"
	"log"
	"os"
)

type fileDict struct {
	length int64
	path   []string
	md5sum string
}

type infoDict struct {
	pieceLength int64
	pieces      string
	name        string

	files []fileDict
}

//MetainfoFile parsed .torrent file
type MetainfoFile struct {
	info         infoDict
	announce     string
	announceList [][]string
	creationDate int64
	comment      string
	createdBy    string
	encoding     string
	private      bool
}

//New creates new torrent
func New(fileName string) MetainfoFile {

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Opening file", err)
	}
	defer file.Close()

	encodedTorrent, err := bencode.Decode(file)
	if err != nil {
		log.Fatal("Decoding torrent", err)
	}

	encodedInfo := encodedTorrent["info"].(map[string]interface{})

	//Parse File element
	files := []fileDict{}
	if length, ok := encodedInfo["length"].(int64); ok {
		//Single File Mode
		files = append(files, fileDict{
			length: length,
			path:   []string{encodedInfo["name"].(string)},
			md5sum: getFromMapWithDefault(encodedTorrent, "md5sum", "").(string),
		})
	} else {
		//Multi File Mode
		for _, encodedFileElement := range encodedInfo["files"].([]interface{}) {
			encodedFile := encodedFileElement.(map[string]interface{})

			encodedPath := []string{}
			for _, encodedPathElement := range encodedFile["path"].([]interface{}) {
				encodedPath = append(encodedPath, encodedPathElement.(string))
			}
			files = append(files, fileDict{
				length: encodedFile["length"].(int64),
				path:   encodedPath,
				md5sum: getFromMapWithDefault(encodedTorrent, "md5sum", "").(string),
			})
		}
	}

	//Parse Announce List (arrays of arrays of strings
	announceList := [][]string{}
	if encodedAnnounceList, ok := encodedTorrent["announce-list"]; ok {
		for _, encodedAnnounceListElement := range encodedAnnounceList.([]interface{}) {
			announceListElement := []string{}
			for _, encodedAnnounceListElementElement := range encodedAnnounceListElement.([]interface{}) {
				announceListElement = append(announceListElement, encodedAnnounceListElementElement.(string))
			}
			announceList = append(announceList, announceListElement)
		}
	}

	return MetainfoFile{
		info: infoDict{
			name:        encodedInfo["name"].(string),
			pieceLength: encodedInfo["piece length"].(int64),
			pieces:      encodedInfo["pieces"].(string),

			files: files,
		},
		announce:     encodedTorrent["announce"].(string),
		announceList: announceList,
		comment:      getFromMapWithDefault(encodedTorrent, "comment", "").(string),
		private:      getFromMapWithDefault(encodedTorrent, "private", false).(bool),
	}
}

func getFromMapWithDefault(aMap map[string]interface{}, key string, defaultValue interface{}) interface{} {
	if element, ok := aMap[key]; ok {
		return element
	} else {
		return defaultValue
	}
}
