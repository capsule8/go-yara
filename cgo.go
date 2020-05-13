// Copyright © 2015-2019 Hilko Bengen <bengen@hilluzination.de>
// All rights reserved.
//
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

package yara

// #cgo CFLAGS: -D_FILE_OFFSET_BITS=64 -DDOTNET_MODULE -DMAGIC_MODULE -DHASH_MODULE -DMACHO_MODULE -DDEX_MODULE -DHAVE_UNISTD_H
// #cgo LDFLAGS:
import "C"
