package torrent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsingSampleTorrent(t *testing.T) {
	torrent := New("sample.torrent")

	assert.Equal(t, "udp://tracker.openbittorrent.com:80", torrent.announce)
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