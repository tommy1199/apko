// Copyright 2023 Chainguard, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package build

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	apkfs "chainguard.dev/apko/pkg/apk/impl/fs"
)

var (
	busyboxList = `
/bin/fdflush
/bin/sync
/bin/ln
/bin/chmod
/bin/dd
/bin/stat
/bin/ls
/bin/ipcalc
/bin/chattr
/bin/fatattr
/bin/sed
/bin/stty
/bin/login
/bin/bbconfig
/bin/link
/bin/rev
/bin/iostat
/bin/ionice
/bin/arch
/bin/reformime
/bin/zcat
/bin/gunzip
/bin/mpstat
/bin/rm
/bin/usleep
/bin/touch
/bin/pipe_progress
/bin/getopt
/bin/ed
/bin/cp
/bin/netstat
/bin/ash
/bin/printenv
/bin/nice
/bin/false
/bin/df
/bin/pidof
/bin/dumpkmap
/bin/fsync
/bin/ping6
/bin/chown
/bin/base64
/bin/kbd_mode
/bin/umount
/bin/ping
/bin/dmesg
/bin/watch
/bin/date
/bin/ps
/bin/setserial
/bin/tar
/bin/sleep
/bin/mknod
/bin/linux32
/bin/mount
/bin/egrep
/bin/run-parts
/bin/makemime
/bin/cat
/bin/lsattr
/bin/mountpoint
/bin/fgrep
/bin/uname
/bin/lzop
/bin/mkdir
/bin/setpriv
/bin/hostname
/bin/gzip
/bin/echo
/bin/dnsdomainname
/bin/mv
/bin/rmdir
/bin/mktemp
/bin/linux64
/bin/more
/bin/grep
/bin/true
/bin/sh
/bin/chgrp
/bin/su
/bin/pwd
/bin/kill
/usr/bin/vi
/usr/bin/sha3sum
/usr/bin/xzcat
/usr/bin/crontab
/usr/bin/bunzip2
/usr/bin/logger
/usr/bin/cksum
/usr/bin/comm
/usr/bin/uniq
/usr/bin/bzip2
/usr/bin/resize
/usr/bin/wc
/usr/bin/free
/usr/bin/install
/usr/bin/unix2dos
/usr/bin/nl
/usr/bin/fallocate
/usr/bin/unxz
/usr/bin/pstree
/usr/bin/sum
/usr/bin/reset
/usr/bin/unlink
/usr/bin/realpath
/usr/bin/sha1sum
/usr/bin/pgrep
/usr/bin/traceroute6
/usr/bin/lzopcat
/usr/bin/cal
/usr/bin/factor
/usr/bin/flock
/usr/bin/volname
/usr/bin/diff
/usr/bin/shred
/usr/bin/unexpand
/usr/bin/lsusb
/usr/bin/nmeter
/usr/bin/sort
/usr/bin/passwd
/usr/bin/nsenter
/usr/bin/nc
/usr/bin/fuser
/usr/bin/tty
/usr/bin/env
/usr/bin/openvt
/usr/bin/od
/usr/bin/expr
/usr/bin/renice
/usr/bin/readlink
/usr/bin/pkill
/usr/bin/cut
/usr/bin/uudecode
/usr/bin/printf
/usr/bin/beep
/usr/bin/xxd
/usr/bin/microcom
/usr/bin/md5sum
/usr/bin/hd
/usr/bin/du
/usr/bin/cpio
/usr/bin/cmp
/usr/bin/hexdump
/usr/bin/cryptpw
/usr/bin/lzcat
/usr/bin/paste
/usr/bin/which
/usr/bin/setkeycodes
/usr/bin/nslookup
/usr/bin/sha256sum
/usr/bin/head
/usr/bin/time
/usr/bin/whoami
/usr/bin/vlock
/usr/bin/mkpasswd
/usr/bin/dirname
/usr/bin/bc
/usr/bin/awk
/usr/bin/bzcat
/usr/bin/uuencode
/usr/bin/unlzop
/usr/bin/pmap
/usr/bin/strings
/usr/bin/expand
/usr/bin/tac
/usr/bin/tee
/usr/bin/eject
/usr/bin/sha512sum
/usr/bin/ipcs
/usr/bin/who
/usr/bin/unlzma
/usr/bin/showkey
/usr/bin/clear
/usr/bin/ipcrm
/usr/bin/nohup
/usr/bin/last
/usr/bin/ttysize
/usr/bin/deallocvt
/usr/bin/unzip
/usr/bin/truncate
/usr/bin/blkdiscard
/usr/bin/seq
/usr/bin/lsof
/usr/bin/pscan
/usr/bin/whois
/usr/bin/killall
/usr/bin/mesg
/usr/bin/lzma
/usr/bin/[[
/usr/bin/id
/usr/bin/nproc
/usr/bin/hostid
/usr/bin/[
/usr/bin/chvt
/usr/bin/dos2unix
/usr/bin/less
/usr/bin/mkfifo
/usr/bin/uptime
/usr/bin/pwdx
/usr/bin/unshare
/usr/bin/shuf
/usr/bin/setsid
/usr/bin/udhcpc6
/usr/bin/tr
/usr/bin/traceroute
/usr/bin/timeout
/usr/bin/groups
/usr/bin/tail
/usr/bin/wget
/usr/bin/split
/usr/bin/top
/usr/bin/test
/usr/bin/yes
/usr/bin/fold
/usr/bin/basename
/usr/bin/dc
/usr/sbin/readahead
/usr/sbin/addgroup
/usr/sbin/setfont
/usr/sbin/rfkill
/usr/sbin/delgroup
/usr/sbin/partprobe
/usr/sbin/remove-shell
/usr/sbin/add-shell
/usr/sbin/crond
/usr/sbin/adduser
/usr/sbin/deluser
/usr/sbin/nanddump
/usr/sbin/sendmail
/usr/sbin/chroot
/usr/sbin/killall5
/usr/sbin/setlogcons
/usr/sbin/nbd-client
/usr/sbin/nandwrite
/usr/sbin/rdate
/usr/sbin/rdev
/usr/sbin/loadfont
/usr/sbin/brctl
/usr/sbin/fbset
/usr/sbin/chpasswd
/usr/sbin/ntpd
/usr/sbin/ether-wake
/usr/sbin/arping
/sbin/pivot_root
/sbin/mkdosfs
/sbin/logread
/sbin/ifconfig
/sbin/slattach
/sbin/inotifyd
/sbin/syslogd
/sbin/halt
/sbin/ifenslave
/sbin/fsck
/sbin/poweroff
/sbin/udhcpc
/sbin/tunctl
/sbin/ifup
/sbin/arp
/sbin/modprobe
/sbin/adjtimex
/sbin/setconsole
/sbin/iplink
/sbin/watchdog
/sbin/route
/sbin/getty
/sbin/ipneigh
/sbin/nologin
/sbin/ipaddr
/sbin/modinfo
/sbin/switch_root
/sbin/blockdev
/sbin/acpid
/sbin/iproute
/sbin/loadkmap
/sbin/mdev
/sbin/init
/sbin/iprule
/sbin/depmod
/sbin/insmod
/sbin/raidautorun
/sbin/nameif
/sbin/blkid
/sbin/findfs
/sbin/ifdown
/sbin/fdisk
/sbin/losetup
/sbin/fbsplash
/sbin/lsmod
/sbin/fstrim
/sbin/iptunnel
/sbin/hwclock
/sbin/mkswap
/sbin/vconfig
/sbin/swapon
/sbin/klogd
/sbin/sysctl
/sbin/rmmod
/sbin/swapoff
/sbin/ip
/sbin/reboot
/sbin/mkfs.vfat
`

	busyboxLinks = strings.Fields(busyboxList)
)

func (di *defaultBuildImplementation) InstallBusyboxLinks(fsys apkfs.FullFS) error {
	// does busybox exist? if not, do not bother with symlinks
	if _, err := fsys.Stat("/bin/busybox"); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
		return nil
	}
	for _, link := range busyboxLinks {
		dir := filepath.Dir(link)
		if err := fsys.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("creating directory %s: %w", dir, err)
		}
		if err := fsys.Symlink("/bin/busybox", link); err != nil {
			return fmt.Errorf("creating busybox link %s: %w", link, err)
		}
	}
	return nil
}
