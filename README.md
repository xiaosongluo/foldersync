# foldersync
[![Build Status](https://travis-ci.org/xiaosongluo/foldersync.svg?branch=master)](https://travis-ci.org/xiaosongluo/foldersync)
[![Coverage Status](https://coveralls.io/repos/github/xiaosongluo/foldersync/badge.svg?branch=master)](https://coveralls.io/github/xiaosongluo/foldersync?branch=master)

The folder sync allows you to keep folder synced between local and remote via ftp.

The folder sync define a flag file, when a folder's content has changed, the flag file in the folder must be change too. So, the folder sync can use the flag file to judge whether the remote folder changes. It's important to avoid useless sync.

## Basic Usage

```Go
package main

import (
	"github.com/xiaosongluo/foldersync"
)

func main(){
	ftp := &foldersync.FTP{"localhost",21,"user","password"}
	ftp.FolderSync("dataset","/000_Business/001_Engine/004_Face","metadata.xml")
}

```    