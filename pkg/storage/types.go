package storage

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"path"
	"time"
)

var (
	ErrNotFound = errors.New("file not found")
)

type IStorage interface {
	WithPrefix(prefix string) (IStorage, error)

	Get(ctx context.Context, filePath string) (IObject, IObjectInfo, error)
	Stat(ctx context.Context, filePath string) (IObjectInfo, error)
	Put(ctx context.Context, filePath string, reader io.Reader, size int64, contentType string) (string, error)
	Delete(ctx context.Context, filePath string) error

	List(ctx context.Context, filePath string, offset int, pageSize int) ([]*FileInfo, error)
}

type IObject interface {
	Read(p []byte) (n int, err error)
	ReadAt(p []byte, off int64) (n int, err error)
	Seek(offset int64, whence int) (int64, error)
	Close() error
}

type IObjectInfo interface {
	GetExtension() string
	GetContentType() string
	GetSize() int64
	GetExpiration() time.Time
}

type ObjectInfo struct {
	extension   string
	contentType string
	size        int64
	expiration  time.Time
}

func (o *ObjectInfo) GetExtension() string {
	return o.extension
}

func (o *ObjectInfo) GetContentType() string {
	return o.contentType
}

func (o *ObjectInfo) GetSize() int64 {
	return o.size
}

func (o *ObjectInfo) GetExpiration() time.Time {
	return o.expiration
}

type FileInfo struct {
	Name         string
	LastModified time.Time
	Size         int64
	ContentType  string
}

func GetFilename(uid string, fileName string, fileExtension string) string {
	hasher := sha1.New()
	hasher.Write([]byte(fileName))
	fileHash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	directory := fileHash[0:2]

	return fmt.Sprintf("%s/%s-%s.%s", directory, uid, fileHash, fileExtension)
}

func FileNameSplitter(fileName string) string {
	if len(fileName) < 3 {
		fileName = "00" + fileName
	}

	chars := fileName
	part1 := chars[0:2]
	return path.Join(part1, fileName)
}
