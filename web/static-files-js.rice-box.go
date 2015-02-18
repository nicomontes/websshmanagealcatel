package web

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	filem := &embedded.EmbeddedFile{
		Filename:    `WebManageAlcatel.js`,
		FileModTime: time.Unix(1423728646, 0),
		Content:     string([]byte{0x2f, 0x2f, 0x20, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x20, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0xa, 0x76, 0x61, 0x72, 0x20, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x3b, 0xa, 0xa, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x73, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x28, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x69, 0x64, 0x20, 0x3d, 0x20, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x65, 0x74, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x28, 0x27, 0x69, 0x64, 0x27, 0x29, 0x3b, 0xa, 0x9, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x28, 0x22, 0x2f, 0x44, 0x53, 0x4c, 0x41, 0x4d, 0x3f, 0x69, 0x64, 0x3d, 0x22, 0x20, 0x2b, 0x20, 0x69, 0x64, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x29, 0x3b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x75, 0x72, 0x6c, 0x20, 0x3d, 0x20, 0x22, 0x2f, 0x44, 0x53, 0x4c, 0x41, 0x4d, 0x3f, 0x69, 0x64, 0x3d, 0x22, 0x20, 0x2b, 0x20, 0x69, 0x64, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x3d, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x58, 0x4d, 0x4c, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x28, 0x29, 0x3b, 0xa, 0x9, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x28, 0x22, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x22, 0x2c, 0x20, 0x75, 0x72, 0x6c, 0x2c, 0x20, 0x74, 0x72, 0x75, 0x65, 0x29, 0x3b, 0xa, 0x9, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x28, 0x29, 0x3b, 0xa, 0x9, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x3d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x29, 0xa, 0x9, 0x7b, 0xa, 0x9, 0x9, 0x69, 0x66, 0x20, 0x28, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x3d, 0x3d, 0x34, 0x20, 0x26, 0x26, 0x20, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3d, 0x3d, 0x32, 0x30, 0x30, 0x29, 0xa, 0x9, 0x9, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3b, 0xa, 0x9, 0x9, 0x7d, 0xa, 0x9, 0x7d, 0xa, 0x7d, 0xa, 0xa, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x6f, 0x44, 0x73, 0x6c, 0x61, 0x6d, 0x28, 0x69, 0x64, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x3d, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x58, 0x4d, 0x4c, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x28, 0x29, 0x3b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x75, 0x72, 0x6c, 0x20, 0x3d, 0x20, 0x22, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3f, 0x69, 0x64, 0x3d, 0x22, 0x20, 0x2b, 0x20, 0x69, 0x64, 0x3b, 0xa, 0x9, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x28, 0x22, 0x47, 0x45, 0x54, 0x22, 0x2c, 0x20, 0x75, 0x72, 0x6c, 0x2c, 0x20, 0x74, 0x72, 0x75, 0x65, 0x29, 0x3b, 0xa, 0x9, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x28, 0x29, 0x3b, 0xa, 0x9, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x3d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x29, 0xa, 0x9, 0x7b, 0xa, 0x9, 0x9, 0x69, 0x66, 0x20, 0x28, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x3d, 0x3d, 0x34, 0x20, 0x26, 0x26, 0x20, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3d, 0x3d, 0x32, 0x30, 0x30, 0x29, 0xa, 0x9, 0x9, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x76, 0x61, 0x72, 0x20, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x20, 0x3d, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x28, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x65, 0x78, 0x74, 0x29, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x3b, 0xa, 0x9, 0x9, 0x9, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x64, 0x28, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x29, 0x3b, 0xa, 0x9, 0x9, 0x7d, 0xa, 0x9, 0x7d, 0xa, 0x7d, 0xa, 0xa, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x64, 0x28, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x3d, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x58, 0x4d, 0x4c, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x28, 0x29, 0x3b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x75, 0x72, 0x6c, 0x20, 0x3d, 0x20, 0x22, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x3f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x3d, 0x22, 0x20, 0x2b, 0x20, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x20, 0x2b, 0x20, 0x22, 0x26, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x3d, 0x73, 0x68, 0x6f, 0x77, 0x20, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x73, 0x6c, 0x6f, 0x74, 0x22, 0x3b, 0xa, 0x9, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x28, 0x22, 0x47, 0x45, 0x54, 0x22, 0x2c, 0x20, 0x75, 0x72, 0x6c, 0x2c, 0x20, 0x74, 0x72, 0x75, 0x65, 0x29, 0x3b, 0xa, 0x9, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x28, 0x29, 0x3b, 0xa, 0x9, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x3d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x29, 0xa, 0x9, 0x7b, 0xa, 0x9, 0x9, 0x69, 0x66, 0x20, 0x28, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x3d, 0x3d, 0x34, 0x20, 0x26, 0x26, 0x20, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3d, 0x3d, 0x32, 0x30, 0x30, 0x29, 0xa, 0x9, 0x9, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x76, 0x61, 0x72, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x75, 0x74, 0x20, 0x3d, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x28, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x65, 0x78, 0x74, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x75, 0x74, 0x3b, 0xa, 0x9, 0x9, 0x9, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x28, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4f, 0x75, 0x74, 0x29, 0x3b, 0xa, 0x9, 0x9, 0x7d, 0xa, 0x9, 0x7d, 0xa, 0x7d, 0xa, 0xa, 0xa, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x67, 0x65, 0x74, 0x44, 0x73, 0x6c, 0x61, 0x6d, 0x28, 0x69, 0x64, 0x29, 0x20, 0x7b, 0xa, 0x9, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x3d, 0x20, 0x6e, 0x65, 0x77, 0x20, 0x58, 0x4d, 0x4c, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x28, 0x29, 0x3b, 0xa, 0x9, 0x76, 0x61, 0x72, 0x20, 0x75, 0x72, 0x6c, 0x20, 0x3d, 0x20, 0x22, 0x2f, 0x67, 0x65, 0x74, 0x44, 0x73, 0x6c, 0x61, 0x6d, 0x3f, 0x69, 0x64, 0x3d, 0x22, 0x20, 0x2b, 0x20, 0x69, 0x64, 0x3b, 0xa, 0x9, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x28, 0x22, 0x47, 0x45, 0x54, 0x22, 0x2c, 0x20, 0x75, 0x72, 0x6c, 0x2c, 0x20, 0x74, 0x72, 0x75, 0x65, 0x29, 0x3b, 0xa, 0x9, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x28, 0x29, 0x3b, 0xa, 0x9, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6f, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x79, 0x73, 0x74, 0x61, 0x74, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x3d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x28, 0x29, 0xa, 0x9, 0x7b, 0xa, 0x9, 0x9, 0x69, 0x66, 0x20, 0x28, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x3d, 0x3d, 0x34, 0x20, 0x26, 0x26, 0x20, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3d, 0x3d, 0x32, 0x30, 0x30, 0x29, 0xa, 0x9, 0x9, 0x7b, 0xa, 0x9, 0x9, 0x9, 0x76, 0x61, 0x72, 0x20, 0x64, 0x73, 0x6c, 0x61, 0x6d, 0x20, 0x3d, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x28, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x65, 0x78, 0x74, 0x29, 0x2e, 0x64, 0x73, 0x6c, 0x61, 0x6d, 0x3b, 0xa, 0x9, 0x9, 0x9, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x20, 0x3d, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x2e, 0x70, 0x61, 0x72, 0x73, 0x65, 0x28, 0x78, 0x6d, 0x6c, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x65, 0x78, 0x74, 0x29, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x3b, 0xa, 0x9, 0x9, 0x9, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x28, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x29, 0x3b, 0xa, 0x9, 0x9, 0x7d, 0xa, 0x9, 0x7d, 0xa, 0x7d, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}

	// define dirs
	dirl := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1420995157, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			filem, // WebManageAlcatel.js

		},
	}

	// link ChildDirs
	dirl.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`static-files/js`, &embedded.EmbeddedBox{
		Name: `static-files/js`,
		Time: time.Unix(1423730269, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dirl,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"WebManageAlcatel.js": filem,
		},
	})
}
