// Copyright © 2015-2020 Hilko Bengen <bengen@hilluzination.de>
// All rights reserved.
//
// Use of this source code is governed by the license that can be
// found in the LICENSE file.

package yara

// #cgo CFLAGS: -D_FILE_OFFSET_BITS=64 -DDOTNET_MODULE -DMAGIC_MODULE -DHASH_MODULE -DMACHO_MODULE -DDEX_MODULE -DHAVE_UNISTD_H -DHAVE_STDBOOL_H
// #cgo !yara_no_pkg_config,!yara_static  pkg-config: yara
// #cgo !yara_no_pkg_config,yara_static   pkg-config: --static yara
// #cgo yara_no_pkg_config,!yara_static   LDFLAGS:    -lyara
// #cgo yara_no_pkg_config,yara_static    LDFLAGS:
/*
#include <yara.h>
#if YR_MAJOR_VERSION != 4
#error YARA version 4 required
#endif
*/
import "C"
