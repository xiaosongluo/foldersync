package fs

import (
	"container/list"
)

type File interface {
	GetName() (string, error)
	GetPath() (string, error)
	GetParent() (File, error)
	IsDirectory() (bool, error)
	IsFile() (bool, error)
	Exists() (bool, error)
	IsBuffered() (bool, error)
	GetUnbuffered() (File, error)
	RefreshBuffer() error
	WriteFileAttributes() error
	GetLastModified() (int, error)
	SetLastModified(lastModified int) error
	GetSize() (int, error)
	SetSize(size int) error
	GetChildren() (list.List, error)
	GetChild(name string) (File, error)
	CreateChild(name string, directory bool) (File, error)
	Refresh() error
	MakeDirectory() (bool, error)
	//public InputStream getInputStream() throws IOException;
	//public OutputStream getOutputStream() throws IOException;
	Delete() (bool, error)
}