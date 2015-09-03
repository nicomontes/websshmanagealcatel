package web

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file5 := &embedded.EmbeddedFile{
		Filename:    `index.tmpl`,
		FileModTime: time.Unix(1437675603, 0),
		Content:     string([]byte{0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x67, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x44, 0x53, 0x4c, 0x41, 0x4d, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x44, 0x53, 0x4c, 0x41, 0x4d, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x2d, 0x31, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x2d, 0x6d, 0x64, 0x2d, 0x31, 0x2d, 0x34, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x44, 0x53, 0x4c, 0x41, 0x4d, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x76, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x67, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7b, 0x7b, 0x2e, 0x44, 0x53, 0x4c, 0x41, 0x4d, 0x4c, 0x69, 0x73, 0x74, 0x7d, 0x7d, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x2d, 0x31, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x2d, 0x6d, 0x64, 0x2d, 0x31, 0x2d, 0x34, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x76, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x67, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7b, 0x7b, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x7d, 0x7d, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x50, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x50, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x2d, 0x31, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x2d, 0x6d, 0x64, 0x2d, 0x31, 0x2d, 0x34, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x50, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x76, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x67, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7b, 0x7b, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x7d, 0x7d, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x2d, 0x31, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x2d, 0x6d, 0x64, 0x2d, 0x31, 0x2d, 0x34, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x69, 0x64, 0x3d, 0x20, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x6d, 0x22, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x2d, 0x74, 0x6f, 0x70, 0x3a, 0x33, 0x30, 0x70, 0x78, 0x3b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x3a, 0x6e, 0x6f, 0x6e, 0x65, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x22, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3d, 0x22, 0x2f, 0x53, 0x49, 0x54, 0x45, 0x41, 0x50, 0x49, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x20, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x3d, 0x22, 0x50, 0x55, 0x54, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x67, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x2d, 0x31, 0x2d, 0x32, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x5c, 0x22, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5c, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3d, 0x22, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x67, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x2d, 0x31, 0x2d, 0x32, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x68, 0x34, 0x3e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x3c, 0x2f, 0x68, 0x34, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x20, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x22, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x22, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x20, 0x66, 0x6f, 0x72, 0x3d, 0x22, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2d, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x3e, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2d, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x22, 0x3e, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x67, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x2d, 0x31, 0x2d, 0x32, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x68, 0x34, 0x3e, 0x54, 0xc3, 0xa9, 0x6c, 0xc3, 0xa9, 0x70, 0x68, 0x6f, 0x6e, 0x69, 0x65, 0x3c, 0x2f, 0x68, 0x34, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x20, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x22, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x76, 0x6f, 0x69, 0x70, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x22, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x76, 0x6f, 0x69, 0x70, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x20, 0x66, 0x6f, 0x72, 0x3d, 0x22, 0x76, 0x6f, 0x69, 0x70, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2d, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x3e, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2d, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x22, 0x3e, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x67, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x2d, 0x31, 0x2d, 0x32, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x68, 0x34, 0x3e, 0x56, 0x69, 0x64, 0xc3, 0xa9, 0x6f, 0x3c, 0x2f, 0x68, 0x34, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x20, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x22, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x69, 0x70, 0x74, 0x76, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x22, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x69, 0x70, 0x74, 0x76, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x20, 0x66, 0x6f, 0x72, 0x3d, 0x22, 0x69, 0x70, 0x74, 0x76, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2d, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x3e, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6f, 0x6e, 0x6f, 0x66, 0x66, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2d, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x22, 0x3e, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x67, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x22, 0x20, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x3d, 0x22, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x39, 0x36, 0x70, 0x78, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x22, 0x3e, 0x53, 0x61, 0x75, 0x76, 0x65, 0x72, 0x3c, 0x2f, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3d, 0x22, 0x22, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x64, 0x73, 0x6c, 0x61, 0x6d, 0x49, 0x44, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3d, 0x22, 0x22, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x73, 0x6c, 0x6f, 0x74, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3d, 0x22, 0x22, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    `options.tmpl`,
		FileModTime: time.Unix(1420828420, 0),
		Content:     string([]byte{0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x67, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x44, 0x53, 0x4c, 0x41, 0x4d, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x44, 0x53, 0x4c, 0x41, 0x4d, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x2d, 0x31, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x2d, 0x6d, 0x64, 0x2d, 0x31, 0x2d, 0x34, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x67, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7b, 0x7b, 0x2e, 0x44, 0x53, 0x4c, 0x41, 0x4d, 0x4c, 0x69, 0x73, 0x74, 0x7d, 0x7d, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3f, 0x61, 0x64, 0x64, 0x22, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x2d, 0x31, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x66, 0x61, 0x20, 0x66, 0x61, 0x2d, 0x70, 0x6c, 0x75, 0x73, 0x20, 0x66, 0x61, 0x2d, 0x32, 0x78, 0x22, 0x3e, 0x3c, 0x2f, 0x69, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x70, 0x61, 0x6e, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x44, 0x53, 0x4c, 0x41, 0x4d, 0x22, 0x3e, 0x41, 0x6a, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x3c, 0x2f, 0x73, 0x70, 0x61, 0x6e, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x61, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x44, 0x53, 0x4c, 0x41, 0x4d, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x2d, 0x31, 0x20, 0x70, 0x75, 0x72, 0x65, 0x2d, 0x75, 0x2d, 0x6d, 0x64, 0x2d, 0x33, 0x2d, 0x34, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7b, 0x7b, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x7d, 0x7d, 0xa, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}

	// define dirs
	dir4 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1420662006, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file5, // index.tmpl
			file6, // options.tmpl

		},
	}

	// link ChildDirs
	dir4.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`views`, &embedded.EmbeddedBox{
		Name: `views`,
		Time: time.Unix(1441308902, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir4,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"index.tmpl":   file5,
			"options.tmpl": file6,
		},
	})
}
