package fs

import (

)

type FileSystem interface {
	GetRoot() (File, error)
	Flush() error
	Close() error
	IsCaseSensitive() (bool, error)
	IsAvailable() (bool, error)
	//FileObject getBase();
	GetConnectionDescription() (ConnectionDescription,error)
}
