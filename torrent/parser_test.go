package torrent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonExistingFile(t *testing.T) {
	torrent := New("nonexisting.torrent")
	assert.Nil(t, torrent)
}

func TestBrokenFile(t *testing.T) {
	torrent := New("sample-broken.torrent")
	assert.Nil(t, torrent)
}


func TestParsingSampleTorrent(t *testing.T) {
	torrent := New("sample.torrent")

	assert.Equal(t, "udp://tracker.openbittorrent.com:80", torrent.announce)
	assert.Equal(t, [][]string{}, torrent.announceList)
	assert.Equal(t, "", torrent.comment)
	assert.Equal(t, "", torrent.createdBy)
	assert.Equal(t, int64(0), torrent.creationDate)
	assert.Equal(t, false, torrent.private)

	assert.Equal(t, "sample.txt", torrent.info.name)
	assert.Equal(t, int64(65536), torrent.info.pieceLength)

	assert.Equal(t, 1, len(torrent.info.files))
	assert.Equal(t, int64(20), torrent.info.files[0].length)
	assert.Equal(t, []string{"sample.txt"}, torrent.info.files[0].path)
	assert.Equal(t, "", torrent.info.files[0].md5sum)
}

func TestParsingSample2Torrent(t *testing.T) {
	torrent := New("sample2.torrent")

	assert.Equal(t, "udp://tracker.publicbt.com:80", torrent.announce)
	assert.Equal(t, "", torrent.comment)
	assert.Equal(t, "", torrent.createdBy)
	assert.Equal(t, int64(0), torrent.creationDate)
	assert.Equal(t, false, torrent.private)

	assert.Equal(t, "urlteam", torrent.info.name)
	assert.Equal(t, int64(4194304), torrent.info.pieceLength)

	assert.Equal(t, 3859, len(torrent.info.files))
	assert.Equal(t, int64(29780), torrent.info.files[0].length)
	assert.Equal(t, []string{"4url.cc.txt.xz"}, torrent.info.files[0].path)
	assert.Equal(t, "", torrent.info.files[0].md5sum)
}

func TestParsingSample3Torrent(t *testing.T) {
	torrent := New("sample3.torrent")

	assert.Equal(t, "udp://open.demonii.com:1337/announce", torrent.announce)
	assert.Equal(t, [][]string{
		[]string{"udp://open.demonii.com:1337/announce"},
		[]string{"udp://tracker.publicbt.com:80/announce"},
		[]string{"udp://tracker.openbittorrent.com:80/announce"},
		[]string{"udp://tracker.istole.it:80/announce"},
		[]string{"http://torrent.gresille.org/announce"},
		[]string{"udp://eddie4.nl:6969/announce"},
		[]string{"udp://coppersurfer.tk:6969/announce"},
		[]string{"http://tracker.aletorrenty.pl:2710/announce"},
		[]string{"udp://glotorrents.pw:6969/announce"},
		[]string{"udp://9.rarbg.com:2710/announce"},
	}, torrent.announceList)
	assert.Equal(t, "Torrent downloaded from torrent cache at http://torcache.net/", torrent.comment)
	assert.Equal(t, "", torrent.createdBy)
	assert.Equal(t, int64(0), torrent.creationDate)
	assert.Equal(t, false, torrent.private)

	assert.Equal(t, "The.Longest.Ride.2015.HDRip.XViD.AC3-ETRG", torrent.info.name)
	assert.Equal(t, int64(262144), torrent.info.pieceLength)

	assert.Equal(t, 6, len(torrent.info.files))
	assert.Equal(t, int64(1515390), torrent.info.files[0].length)
	assert.Equal(t, []string{"ETRG.mp4"}, torrent.info.files[0].path)
	assert.Equal(t, "", torrent.info.files[0].md5sum)
}
