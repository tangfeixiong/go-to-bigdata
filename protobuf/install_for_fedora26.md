## CI/CD

source
```
[vagrant@kubedev-172-17-4-59 ~]$ cd /Users/fanhongling/go/src/github.com/google/protobuf/
```

```
[vagrant@kubedev-172-17-4-59 protobuf]$ git log -1
commit 4b2977b39684f762b9d2a9746a0029af4f24442b (HEAD -> master, origin/master, origin/HEAD)
Merge: b5f09c1e ae49cfd1
Author: Feng Xiao <xfxyjwf@gmail.com>
Date:   Wed Dec 20 17:45:57 2017 -0800

    Merge pull request #4082 from matt-kwong/kokoro_jobs
    
    Migrate Jenkins jobs to Kokoro
```

tool chain
```
[vagrant@kubedev-172-17-4-59 protobuf]$ sudo dnf install --disablerepo=kubernetes autoconf automake libtool curl make gcc-c++ unzip
Failed to set locale, defaulting to C
Last metadata expiration check: 3:19:30 ago on Thu Dec 21 05:03:10 2017.
Package curl-7.53.1-7.fc26.x86_64 is already installed, skipping.
Package make-1:4.2.1-2.fc26.x86_64 is already installed, skipping.
Dependencies resolved.
================================================================================================================================================
 Package                                 Arch                         Version                               Repository                     Size
================================================================================================================================================
Installing:
 autoconf                                noarch                       2.69-24.fc26                          fedora                        709 k
 automake                                noarch                       1.15-9.fc26                           fedora                        695 k
 gcc-c++                                 x86_64                       7.2.1-2.fc26                          updates                        11 M
 libtool                                 x86_64                       2.4.6-17.fc26                         fedora                        706 k
 unzip                                   x86_64                       6.0-34.fc26                           updates                       186 k
Upgrading:
 libstdc++                               x86_64                       7.2.1-2.fc26                          updates                       462 k
Installing dependencies:
 libstdc++-devel                         x86_64                       7.2.1-2.fc26                          updates                       1.9 M
 m4                                      x86_64                       1.4.18-3.fc26                         fedora                        217 k
 perl-Thread-Queue                       noarch                       3.12-1.fc26                           fedora                         22 k

Transaction Summary
================================================================================================================================================
Install  8 Packages
Upgrade  1 Package

Total download size: 16 M
Is this ok [y/N]: y
Downloading Packages:
(1/9): automake-1.15-9.fc26.noarch.rpm                                                                          832 kB/s | 695 kB     00:00    
(2/9): libtool-2.4.6-17.fc26.x86_64.rpm                                                                         830 kB/s | 706 kB     00:00    
(3/9): perl-Thread-Queue-3.12-1.fc26.noarch.rpm                                                                 297 kB/s |  22 kB     00:00    
(4/9): autoconf-2.69-24.fc26.noarch.rpm                                                                         663 kB/s | 709 kB     00:01    
(5/9): m4-1.4.18-3.fc26.x86_64.rpm                                                                              661 kB/s | 217 kB     00:00    
(6/9): unzip-6.0-34.fc26.x86_64.rpm                                                                             952 kB/s | 186 kB     00:00    
(7/9): libstdc++-7.2.1-2.fc26.x86_64.rpm                                                                        904 kB/s | 462 kB     00:00    
(8/9): libstdc++-devel-7.2.1-2.fc26.x86_64.rpm                                                                  527 kB/s | 1.9 MB     00:03    
(9/9): gcc-c++-7.2.1-2.fc26.x86_64.rpm                                                                          1.7 MB/s |  11 MB     00:06    
------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                           1.5 MB/s |  16 MB     00:10     
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                                                                                        1/1 
  Upgrading        : libstdc++-7.2.1-2.fc26.x86_64                                                                                         1/10 
  Running scriptlet: libstdc++-7.2.1-2.fc26.x86_64                                                                                         1/10 
  Installing       : libstdc++-devel-7.2.1-2.fc26.x86_64                                                                                   2/10 
  Installing       : perl-Thread-Queue-3.12-1.fc26.noarch                                                                                  3/10 
  Installing       : m4-1.4.18-3.fc26.x86_64                                                                                               4/10 
  Running scriptlet: m4-1.4.18-3.fc26.x86_64                                                                                               4/10 
  Installing       : autoconf-2.69-24.fc26.noarch                                                                                          5/10 
  Running scriptlet: autoconf-2.69-24.fc26.noarch                                                                                          5/10 
  Installing       : automake-1.15-9.fc26.noarch                                                                                           6/10 
  Running scriptlet: automake-1.15-9.fc26.noarch                                                                                           6/10 
  Installing       : libtool-2.4.6-17.fc26.x86_64                                                                                          7/10 
  Running scriptlet: libtool-2.4.6-17.fc26.x86_64                                                                                          7/10 
  Installing       : gcc-c++-7.2.1-2.fc26.x86_64                                                                                           8/10 
  Installing       : unzip-6.0-34.fc26.x86_64                                                                                              9/10 
  Cleanup          : libstdc++-7.1.1-3.fc26.x86_64                                                                                        10/10 
  Running scriptlet: libstdc++-7.1.1-3.fc26.x86_64                                                                                        10/10 
  Verifying        : autoconf-2.69-24.fc26.noarch                                                                                          1/10 
  Verifying        : automake-1.15-9.fc26.noarch                                                                                           2/10 
  Verifying        : libtool-2.4.6-17.fc26.x86_64                                                                                          3/10 
  Verifying        : m4-1.4.18-3.fc26.x86_64                                                                                               4/10 
  Verifying        : perl-Thread-Queue-3.12-1.fc26.noarch                                                                                  5/10 
  Verifying        : gcc-c++-7.2.1-2.fc26.x86_64                                                                                           6/10 
  Verifying        : unzip-6.0-34.fc26.x86_64                                                                                              7/10 
  Verifying        : libstdc++-devel-7.2.1-2.fc26.x86_64                                                                                   8/10 
  Verifying        : libstdc++-7.2.1-2.fc26.x86_64                                                                                         9/10 
  Verifying        : libstdc++-7.1.1-3.fc26.x86_64                                                                                        10/10 

Installed:
  autoconf.noarch 2.69-24.fc26    automake.noarch 1.15-9.fc26            gcc-c++.x86_64 7.2.1-2.fc26    libtool.x86_64 2.4.6-17.fc26           
  unzip.x86_64 6.0-34.fc26        libstdc++-devel.x86_64 7.2.1-2.fc26    m4.x86_64 1.4.18-3.fc26        perl-Thread-Queue.noarch 3.12-1.fc26   

Upgraded:
  libstdc++.x86_64 7.2.1-2.fc26                                                                                                                 

Complete!
```

autogen
```
[vagrant@kubedev-172-17-4-59 protobuf]$ ./autogen.sh 
+ autoreconf -f -i -Wall,no-obsolete
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
libtoolize: putting auxiliary files in AC_CONFIG_AUX_DIR, 'build-aux'.
libtoolize: copying file 'build-aux/ltmain.sh'
libtoolize: putting macros in AC_CONFIG_MACRO_DIRS, 'm4'.
libtoolize: copying file 'm4/libtool.m4'
libtoolize: copying file 'm4/ltoptions.m4'
libtoolize: copying file 'm4/ltsugar.m4'
libtoolize: copying file 'm4/ltversion.m4'
libtoolize: copying file 'm4/lt~obsolete.m4'
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
configure.ac:27: installing 'build-aux/compile'
configure.ac:24: installing 'build-aux/missing'
Makefile.am: installing 'build-aux/depcomp'
libtoolize: putting auxiliary files in AC_CONFIG_AUX_DIR, 'build-aux'.
libtoolize: copying file 'build-aux/ltmain.sh'
libtoolize: Consider adding 'AC_CONFIG_MACRO_DIRS([m4])' to configure.ac,
libtoolize: and rerunning libtoolize and aclocal.
libtoolize: Consider adding '-I m4' to ACLOCAL_AMFLAGS in Makefile.am.
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
configure.ac:22: installing 'build-aux/compile'
configure.ac:19: installing 'build-aux/missing'
Makefile.am: installing 'build-aux/depcomp'
libtoolize: putting auxiliary files in '.'.
libtoolize: copying file './ltmain.sh'
libtoolize: putting macros in AC_CONFIG_MACRO_DIRS, 'm4'.
libtoolize: copying file 'm4/libtool.m4'
libtoolize: copying file 'm4/ltoptions.m4'
libtoolize: copying file 'm4/ltsugar.m4'
libtoolize: copying file 'm4/ltversion.m4'
libtoolize: copying file 'm4/lt~obsolete.m4'
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
perl: warning: Setting locale failed.
perl: warning: Please check that your locale settings:
	LANGUAGE = (unset),
	LC_ALL = (unset),
	LANG = "zh_CN.UTF-8"
    are supported and installed on your system.
perl: warning: Falling back to the standard locale ("C").
configure.ac:71: installing './compile'
configure.ac:48: installing './missing'
benchmarks/Makefile.am: installing './depcomp'
+ rm -rf autom4te.cache config.h.in~
+ exit 0
```

locale
```
[vagrant@kubedev-172-17-4-59 protobuf]$ sudo localectl list-locales
C.utf8
en_US
en_US.iso88591
en_US.iso885915
en_US.utf8
```


```
[vagrant@kubedev-172-17-4-59 protobuf]$ sudo localectl status
   System Locale: LANG=en_US.UTF-8
       VC Keymap: us
      X11 Layout: n/a
```

```
[vagrant@kubedev-172-17-4-59 protobuf]$ locale
locale: Cannot set LC_CTYPE to default locale: No such file or directory
locale: Cannot set LC_MESSAGES to default locale: No such file or directory
locale: Cannot set LC_ALL to default locale: No such file or directory
LANG=zh_CN.UTF-8
LC_CTYPE="zh_CN.UTF-8"
LC_NUMERIC="zh_CN.UTF-8"
LC_TIME="zh_CN.UTF-8"
LC_COLLATE="zh_CN.UTF-8"
LC_MONETARY="zh_CN.UTF-8"
LC_MESSAGES="zh_CN.UTF-8"
LC_PAPER="zh_CN.UTF-8"
LC_NAME="zh_CN.UTF-8"
LC_ADDRESS="zh_CN.UTF-8"
LC_TELEPHONE="zh_CN.UTF-8"
LC_MEASUREMENT="zh_CN.UTF-8"
LC_IDENTIFICATION="zh_CN.UTF-8"
LC_ALL=
```

``
[vagrant@kubedev-172-17-4-59 protobuf]$ sudo localectl set-locale LC_CTYPE=C.utf8 LC_MESSAGES=en_US.utf8 LANG=en_US.utf8
```

```
[vagrant@kubedev-172-17-4-59 protobuf]$ cat /etc/locale.conf 
LANG=en_US.utf8
LC_CTYPE=C.utf8
```

configure
```
[vagrant@kubedev-172-17-4-59 protobuf]$ ./configure 
checking whether to enable maintainer-specific portions of Makefiles... yes
checking build system type... x86_64-pc-linux-gnu
checking host system type... x86_64-pc-linux-gnu
checking target system type... x86_64-pc-linux-gnu
checking for a BSD-compatible install... /usr/bin/install -c
checking whether build environment is sane... yes
checking for a thread-safe mkdir -p... /usr/bin/mkdir -p
checking for gawk... gawk
checking whether make sets $(MAKE)... yes
checking whether make supports nested variables... yes
checking whether UID '1000' is supported by ustar format... yes
checking whether GID '1000' is supported by ustar format... yes
checking how to create a ustar tar archive... gnutar
checking for gcc... gcc
checking whether the C compiler works... yes
checking for C compiler default output file name... a.out
checking for suffix of executables... 
checking whether we are cross compiling... no
checking for suffix of object files... o
checking whether we are using the GNU C compiler... yes
checking whether gcc accepts -g... yes
checking for gcc option to accept ISO C89... none needed
checking whether gcc understands -c and -o together... yes
checking for style of include used by make... GNU
checking dependency style of gcc... gcc3
checking for g++... g++
checking whether we are using the GNU C++ compiler... yes
checking whether g++ accepts -g... yes
checking dependency style of g++... gcc3
checking how to run the C preprocessor... gcc -E
checking for gcc... gcc
checking whether we are using the GNU C compiler... (cached) yes
checking whether gcc accepts -g... yes
checking for gcc option to accept ISO C89... (cached) none needed
checking whether gcc understands -c and -o together... (cached) yes
checking dependency style of gcc... (cached) gcc3
checking how to run the C preprocessor... gcc -E
checking how to run the C++ preprocessor... g++ -E
checking for g++... g++
checking whether we are using the GNU C++ compiler... (cached) yes
checking whether g++ accepts -g... yes
checking dependency style of g++... (cached) gcc3
checking how to run the C++ preprocessor... g++ -E
checking for grep that handles long lines and -e... /usr/bin/grep
checking for egrep... /usr/bin/grep -E
checking for ANSI C header files... yes
checking for sys/types.h... yes
checking for sys/stat.h... yes
checking for stdlib.h... yes
checking for string.h... yes
checking for memory.h... yes
checking for strings.h... yes
checking for inttypes.h... yes
checking for stdint.h... yes
checking for unistd.h... yes
checking minix/config.h usability... no
checking minix/config.h presence... no
checking for minix/config.h... no
checking whether it is safe to define __EXTENSIONS__... yes
checking for ar... ar
checking the archiver (ar) interface... ar
checking for gcc... gcc
checking whether we are using the GNU Objective C compiler... no
checking whether gcc accepts -g... no
checking dependency style of gcc... gcc3
checking C++ compiler flags...... use default: -O2  -g -DNDEBUG
checking whether __SUNPRO_CC is declared... no
checking how to print strings... printf
checking for a sed that does not truncate output... /usr/bin/sed
checking for fgrep... /usr/bin/grep -F
checking for ld used by gcc... /usr/bin/ld
checking if the linker (/usr/bin/ld) is GNU ld... yes
checking for BSD- or MS-compatible name lister (nm)... /usr/bin/nm -B
checking the name lister (/usr/bin/nm -B) interface... BSD nm
checking whether ln -s works... yes
checking the maximum length of command line arguments... 1572864
checking how to convert x86_64-pc-linux-gnu file names to x86_64-pc-linux-gnu format... func_convert_file_noop
checking how to convert x86_64-pc-linux-gnu file names to toolchain format... func_convert_file_noop
checking for /usr/bin/ld option to reload object files... -r
checking for objdump... objdump
checking how to recognize dependent libraries... pass_all
checking for dlltool... no
checking how to associate runtime and link libraries... printf %s\n
checking for archiver @FILE support... @
checking for strip... strip
checking for ranlib... ranlib
checking command to parse /usr/bin/nm -B output from gcc object... ok
checking for sysroot... no
checking for a working dd... /usr/bin/dd
checking how to truncate binary pipes... /usr/bin/dd bs=4096 count=1
checking for mt... no
checking if : is a manifest tool... no
checking for dlfcn.h... yes
checking for objdir... .libs
checking if gcc supports -fno-rtti -fno-exceptions... no
checking for gcc option to produce PIC... -fPIC -DPIC
checking if gcc PIC flag -fPIC -DPIC works... yes
checking if gcc static flag -static works... no
checking if gcc supports -c -o file.o... yes
checking if gcc supports -c -o file.o... (cached) yes
checking whether the gcc linker (/usr/bin/ld -m elf_x86_64) supports shared libraries... yes
checking whether -lc should be explicitly linked in... no
checking dynamic linker characteristics... GNU/Linux ld.so
checking how to hardcode library paths into programs... immediate
checking whether stripping libraries is possible... yes
checking if libtool supports shared libraries... yes
checking whether to build shared libraries... yes
checking whether to build static libraries... yes
checking how to run the C++ preprocessor... g++ -E
checking for ld used by g++... /usr/bin/ld -m elf_x86_64
checking if the linker (/usr/bin/ld -m elf_x86_64) is GNU ld... yes
checking whether the g++ linker (/usr/bin/ld -m elf_x86_64) supports shared libraries... yes
checking for g++ option to produce PIC... -fPIC -DPIC
checking if g++ PIC flag -fPIC -DPIC works... yes
checking if g++ static flag -static works... no
checking if g++ supports -c -o file.o... yes
checking if g++ supports -c -o file.o... (cached) yes
checking whether the g++ linker (/usr/bin/ld -m elf_x86_64) supports shared libraries... yes
checking dynamic linker characteristics... (cached) GNU/Linux ld.so
checking how to hardcode library paths into programs... immediate
checking whether the linker supports version scripts... yes
checking for ANSI C header files... (cached) yes
checking fcntl.h usability... yes
checking fcntl.h presence... yes
checking for fcntl.h... yes
checking for inttypes.h... (cached) yes
checking limits.h usability... yes
checking limits.h presence... yes
checking for limits.h... yes
checking for stdlib.h... (cached) yes
checking for unistd.h... (cached) yes
checking for working memcmp... yes
checking for working strtod... yes
checking for ftruncate... yes
checking for memset... yes
checking for mkdir... yes
checking for strchr... yes
checking for strerror... yes
checking for strtol... yes
checking zlib version... ok (1.2.0.4 or later)
checking for library containing zlibVersion... -lz
checking for the pthreads library -lpthreads... no
checking whether pthreads work without any flags... no
checking whether pthreads work with -Kthread... no
checking whether pthreads work with -kthread... no
checking for the pthreads library -llthread... no
checking whether pthreads work with -pthread... yes
checking for joinable pthread attribute... PTHREAD_CREATE_JOINABLE
checking if more special flags are required for pthreads... no
checking whether to check for GCC pthread/shared inconsistencies... yes
checking whether -pthread is sufficient with -shared... yes
checking whether what we have so far is sufficient with -nostdlib... no
checking whether -lpthread saves the day... yes
checking the location of hash_map... <unordered_map>
checking for library containing sched_yield... none required
checking whether g++ supports C++11 features by default... yes
checking that generated files are newer than configure... done
configure: creating ./config.status
config.status: creating Makefile
config.status: creating src/Makefile
config.status: creating benchmarks/Makefile
config.status: creating conformance/Makefile
config.status: creating protobuf.pc
config.status: creating protobuf-lite.pc
config.status: creating config.h
config.status: executing depfiles commands
config.status: executing libtool commands
=== configuring in gmock (/Users/fanhongling/go/src/github.com/google/protobuf/gmock)
configure: running /bin/sh ./configure --disable-option-checking '--prefix=/usr/local'  --cache-file=/dev/null --srcdir=.
checking for a BSD-compatible install... /usr/bin/install -c
checking whether build environment is sane... yes
checking for a thread-safe mkdir -p... /usr/bin/mkdir -p
checking for gawk... gawk
checking whether make sets $(MAKE)... yes
checking whether make supports nested variables... yes
checking for gcc... gcc
checking whether the C compiler works... yes
checking for C compiler default output file name... a.out
checking for suffix of executables... 
checking whether we are cross compiling... no
checking for suffix of object files... o
checking whether we are using the GNU C compiler... yes
checking whether gcc accepts -g... yes
checking for gcc option to accept ISO C89... none needed
checking whether gcc understands -c and -o together... yes
checking for style of include used by make... GNU
checking dependency style of gcc... gcc3
checking for g++... g++
checking whether we are using the GNU C++ compiler... yes
checking whether g++ accepts -g... yes
checking dependency style of g++... gcc3
checking build system type... x86_64-pc-linux-gnu
checking host system type... x86_64-pc-linux-gnu
checking how to print strings... printf
checking for a sed that does not truncate output... /usr/bin/sed
checking for grep that handles long lines and -e... /usr/bin/grep
checking for egrep... /usr/bin/grep -E
checking for fgrep... /usr/bin/grep -F
checking for ld used by gcc... /usr/bin/ld
checking if the linker (/usr/bin/ld) is GNU ld... yes
checking for BSD- or MS-compatible name lister (nm)... /usr/bin/nm -B
checking the name lister (/usr/bin/nm -B) interface... BSD nm
checking whether ln -s works... yes
checking the maximum length of command line arguments... 1572864
checking how to convert x86_64-pc-linux-gnu file names to x86_64-pc-linux-gnu format... func_convert_file_noop
checking how to convert x86_64-pc-linux-gnu file names to toolchain format... func_convert_file_noop
checking for /usr/bin/ld option to reload object files... -r
checking for objdump... objdump
checking how to recognize dependent libraries... pass_all
checking for dlltool... no
checking how to associate runtime and link libraries... printf %s\n
checking for ar... ar
checking for archiver @FILE support... @
checking for strip... strip
checking for ranlib... ranlib
checking command to parse /usr/bin/nm -B output from gcc object... ok
checking for sysroot... no
checking for a working dd... /usr/bin/dd
checking how to truncate binary pipes... /usr/bin/dd bs=4096 count=1
checking for mt... no
checking if : is a manifest tool... no
checking how to run the C preprocessor... gcc -E
checking for ANSI C header files... yes
checking for sys/types.h... yes
checking for sys/stat.h... yes
checking for stdlib.h... yes
checking for string.h... yes
checking for memory.h... yes
checking for strings.h... yes
checking for inttypes.h... yes
checking for stdint.h... yes
checking for unistd.h... yes
checking for dlfcn.h... yes
checking for objdir... .libs
checking if gcc supports -fno-rtti -fno-exceptions... no
checking for gcc option to produce PIC... -fPIC -DPIC
checking if gcc PIC flag -fPIC -DPIC works... yes
checking if gcc static flag -static works... no
checking if gcc supports -c -o file.o... yes
checking if gcc supports -c -o file.o... (cached) yes
checking whether the gcc linker (/usr/bin/ld -m elf_x86_64) supports shared libraries... yes
checking whether -lc should be explicitly linked in... no
checking dynamic linker characteristics... GNU/Linux ld.so
checking how to hardcode library paths into programs... immediate
checking whether stripping libraries is possible... yes
checking if libtool supports shared libraries... yes
checking whether to build shared libraries... yes
checking whether to build static libraries... yes
checking how to run the C++ preprocessor... g++ -E
checking for ld used by g++... /usr/bin/ld -m elf_x86_64
checking if the linker (/usr/bin/ld -m elf_x86_64) is GNU ld... yes
checking whether the g++ linker (/usr/bin/ld -m elf_x86_64) supports shared libraries... yes
checking for g++ option to produce PIC... -fPIC -DPIC
checking if g++ PIC flag -fPIC -DPIC works... yes
checking if g++ static flag -static works... no
checking if g++ supports -c -o file.o... yes
checking if g++ supports -c -o file.o... (cached) yes
checking whether the g++ linker (/usr/bin/ld -m elf_x86_64) supports shared libraries... yes
checking dynamic linker characteristics... (cached) GNU/Linux ld.so
checking how to hardcode library paths into programs... immediate
checking for python... :
checking for the pthreads library -lpthreads... no
checking whether pthreads work without any flags... no
checking whether pthreads work with -Kthread... no
checking whether pthreads work with -kthread... no
checking for the pthreads library -llthread... no
checking whether pthreads work with -pthread... yes
checking for joinable pthread attribute... PTHREAD_CREATE_JOINABLE
checking if more special flags are required for pthreads... no
checking whether to check for GCC pthread/shared inconsistencies... yes
checking whether -pthread is sufficient with -shared... yes
checking for gtest-config... no
checking that generated files are newer than configure... done
configure: creating ./config.status
config.status: creating Makefile
config.status: creating scripts/gmock-config
config.status: creating build-aux/config.h
config.status: build-aux/config.h is unchanged
config.status: executing depfiles commands
config.status: executing libtool commands
=== configuring in gtest (/Users/fanhongling/go/src/github.com/google/protobuf/gmock/gtest)
configure: running /bin/sh ./configure --disable-option-checking '--prefix=/usr/local'  'CFLAGS=' 'CXXFLAGS= -g -DNDEBUG' --cache-file=/dev/null --srcdir=.
checking for a BSD-compatible install... /usr/bin/install -c
checking whether build environment is sane... yes
checking for a thread-safe mkdir -p... /usr/bin/mkdir -p
checking for gawk... gawk
checking whether make sets $(MAKE)... yes
checking whether make supports nested variables... yes
checking for gcc... gcc
checking whether the C compiler works... yes
checking for C compiler default output file name... a.out
checking for suffix of executables... 
checking whether we are cross compiling... no
checking for suffix of object files... o
checking whether we are using the GNU C compiler... yes
checking whether gcc accepts -g... yes
checking for gcc option to accept ISO C89... none needed
checking whether gcc understands -c and -o together... yes
checking for style of include used by make... GNU
checking dependency style of gcc... gcc3
checking for g++... g++
checking whether we are using the GNU C++ compiler... yes
checking whether g++ accepts -g... yes
checking dependency style of g++... gcc3
checking build system type... x86_64-pc-linux-gnu
checking host system type... x86_64-pc-linux-gnu
checking how to print strings... printf
checking for a sed that does not truncate output... /usr/bin/sed
checking for grep that handles long lines and -e... /usr/bin/grep
checking for egrep... /usr/bin/grep -E
checking for fgrep... /usr/bin/grep -F
checking for ld used by gcc... /usr/bin/ld
checking if the linker (/usr/bin/ld) is GNU ld... yes
checking for BSD- or MS-compatible name lister (nm)... /usr/bin/nm -B
checking the name lister (/usr/bin/nm -B) interface... BSD nm
checking whether ln -s works... yes
checking the maximum length of command line arguments... 1572864
checking how to convert x86_64-pc-linux-gnu file names to x86_64-pc-linux-gnu format... func_convert_file_noop
checking how to convert x86_64-pc-linux-gnu file names to toolchain format... func_convert_file_noop
checking for /usr/bin/ld option to reload object files... -r
checking for objdump... objdump
checking how to recognize dependent libraries... pass_all
checking for dlltool... no
checking how to associate runtime and link libraries... printf %s\n
checking for ar... ar
checking for archiver @FILE support... @
checking for strip... strip
checking for ranlib... ranlib
checking command to parse /usr/bin/nm -B output from gcc object... ok
checking for sysroot... no
checking for a working dd... /usr/bin/dd
checking how to truncate binary pipes... /usr/bin/dd bs=4096 count=1
checking for mt... no
checking if : is a manifest tool... no
checking how to run the C preprocessor... gcc -E
checking for ANSI C header files... yes
checking for sys/types.h... yes
checking for sys/stat.h... yes
checking for stdlib.h... yes
checking for string.h... yes
checking for memory.h... yes
checking for strings.h... yes
checking for inttypes.h... yes
checking for stdint.h... yes
checking for unistd.h... yes
checking for dlfcn.h... yes
checking for objdir... .libs
checking if gcc supports -fno-rtti -fno-exceptions... no
checking for gcc option to produce PIC... -fPIC -DPIC
checking if gcc PIC flag -fPIC -DPIC works... yes
checking if gcc static flag -static works... no
checking if gcc supports -c -o file.o... yes
checking if gcc supports -c -o file.o... (cached) yes
checking whether the gcc linker (/usr/bin/ld -m elf_x86_64) supports shared libraries... yes
checking whether -lc should be explicitly linked in... no
checking dynamic linker characteristics... GNU/Linux ld.so
checking how to hardcode library paths into programs... immediate
checking whether stripping libraries is possible... yes
checking if libtool supports shared libraries... yes
checking whether to build shared libraries... yes
checking whether to build static libraries... yes
checking how to run the C++ preprocessor... g++ -E
checking for ld used by g++... /usr/bin/ld -m elf_x86_64
checking if the linker (/usr/bin/ld -m elf_x86_64) is GNU ld... yes
checking whether the g++ linker (/usr/bin/ld -m elf_x86_64) supports shared libraries... yes
checking for g++ option to produce PIC... -fPIC -DPIC
checking if g++ PIC flag -fPIC -DPIC works... yes
checking if g++ static flag -static works... no
checking if g++ supports -c -o file.o... yes
checking if g++ supports -c -o file.o... (cached) yes
checking whether the g++ linker (/usr/bin/ld -m elf_x86_64) supports shared libraries... yes
checking dynamic linker characteristics... (cached) GNU/Linux ld.so
checking how to hardcode library paths into programs... immediate
checking for python... :
checking for the pthreads library -lpthreads... no
checking whether pthreads work without any flags... no
checking whether pthreads work with -Kthread... no
checking whether pthreads work with -kthread... no
checking for the pthreads library -llthread... no
checking whether pthreads work with -pthread... yes
checking for joinable pthread attribute... PTHREAD_CREATE_JOINABLE
checking if more special flags are required for pthreads... no
checking whether to check for GCC pthread/shared inconsistencies... yes
checking whether -pthread is sufficient with -shared... yes
checking that generated files are newer than configure... done
configure: creating ./config.status
config.status: creating Makefile
config.status: creating scripts/gtest-config
config.status: creating build-aux/config.h
config.status: build-aux/config.h is unchanged
config.status: executing depfiles commands
config.status: executing libtool commands
```

make
```
[vagrant@kubedev-172-17-4-59 protobuf]$ make
make  all-recursive
make[1]: Entering directory '/Users/fanhongling/go/src/github.com/google/protobuf'
Making all in .
make[2]: Entering directory '/Users/fanhongling/go/src/github.com/google/protobuf'
make[2]: Nothing to be done for 'all-am'.
make[2]: Leaving directory '/Users/fanhongling/go/src/github.com/google/protobuf'
Making all in src
make[2]: Entering directory '/Users/fanhongling/go/src/github.com/google/protobuf/src'
depbase=`echo google/protobuf/compiler/main.o | sed 's|[^/]*$|.deps/&|;s|\.o$||'`;\
g++ -DHAVE_CONFIG_H -I. -I..    -pthread -DHAVE_PTHREAD=1 -DHAVE_ZLIB=1 -Wall -Wno-sign-compare -O2 -g -DNDEBUG -MT google/protobuf/compiler/main.o -MD -MP -MF $depbase.Tpo -c -o google/protobuf/compiler/main.o google/protobuf/compiler/main.cc &&\
mv -f $depbase.Tpo $depbase.Po
depbase=`echo google/protobuf/stubs/atomicops_internals_x86_gcc.lo | sed 's|[^/]*$|.deps/&|;s|\.lo$||'`;\
/bin/sh ../libtool  --tag=CXX   --mode=compile g++ -DHAVE_CONFIG_H -I. -I..    -pthread -DHAVE_PTHREAD=1 -DHAVE_ZLIB=1 -Wall -Wno-sign-compare -O2 -g -DNDEBUG -MT google/protobuf/stubs/atomicops_internals_x86_gcc.lo -MD -MP -MF $depbase.Tpo -c -o google/protobuf/stubs/atomicops_internals_x86_gcc.lo google/protobuf/stubs/atomicops_internals_x86_gcc.cc &&\
mv -f $depbase.Tpo $depbase.Plo
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I.. -pthread -DHAVE_PTHREAD=1 -DHAVE_ZLIB=1 -Wall -Wno-sign-compare -O2 -g -DNDEBUG -MT google/protobuf/stubs/atomicops_internals_x86_gcc.lo -MD -MP -MF google/protobuf/stubs/.deps/atomicops_internals_x86_gcc.Tpo -c google/protobuf/stubs/atomicops_internals_x86_gcc.cc  -fPIC -DPIC -o google/protobuf/stubs/.libs/atomicops_internals_x86_gcc.o
### snippets ###
libtool: compile:  g++ -DHAVE_CONFIG_H -I. -I.. -pthread -DHAVE_PTHREAD=1 -DHAVE_ZLIB=1 -Wall -Wno-sign-compare -O2 -g -DNDEBUG -MT google/protobuf/compiler/csharp/csharp_wrapper_field.lo -MD -MP -MF google/protobuf/compiler/csharp/.deps/csharp_wrapper_field.Tpo -c google/protobuf/compiler/csharp/csharp_wrapper_field.cc -o google/protobuf/compiler/csharp/csharp_wrapper_field.o >/dev/null 2>&1
/bin/sh ../libtool  --tag=CXX   --mode=link g++ -pthread -DHAVE_PTHREAD=1 -DHAVE_ZLIB=1 -Wall -Wno-sign-compare -O2 -g -DNDEBUG -version-info 15:0:0 -export-dynamic -no-undefined -Wl,--version-script=./libprotoc.map  -o libprotoc.la -rpath /usr/local/lib google/protobuf/compiler/code_generator.lo google/protobuf/compiler/command_line_interface.lo google/protobuf/compiler/plugin.lo google/protobuf/compiler/plugin.pb.lo google/protobuf/compiler/subprocess.lo google/protobuf/compiler/zip_writer.lo google/protobuf/compiler/cpp/cpp_enum.lo google/protobuf/compiler/cpp/cpp_enum_field.lo google/protobuf/compiler/cpp/cpp_extension.lo google/protobuf/compiler/cpp/cpp_field.lo google/protobuf/compiler/cpp/cpp_file.lo google/protobuf/compiler/cpp/cpp_generator.lo google/protobuf/compiler/cpp/cpp_helpers.lo google/protobuf/compiler/cpp/cpp_map_field.lo google/protobuf/compiler/cpp/cpp_message.lo google/protobuf/compiler/cpp/cpp_message_field.lo google/protobuf/compiler/cpp/cpp_padding_optimizer.lo google/protobuf/compiler/cpp/cpp_primitive_field.lo google/protobuf/compiler/cpp/cpp_service.lo google/protobuf/compiler/cpp/cpp_string_field.lo google/protobuf/compiler/java/java_context.lo google/protobuf/compiler/java/java_enum.lo google/protobuf/compiler/java/java_enum_lite.lo google/protobuf/compiler/java/java_enum_field.lo google/protobuf/compiler/java/java_enum_field_lite.lo google/protobuf/compiler/java/java_extension.lo google/protobuf/compiler/java/java_extension_lite.lo google/protobuf/compiler/java/java_field.lo google/protobuf/compiler/java/java_file.lo google/protobuf/compiler/java/java_generator.lo google/protobuf/compiler/java/java_generator_factory.lo google/protobuf/compiler/java/java_helpers.lo google/protobuf/compiler/java/java_lazy_message_field.lo google/protobuf/compiler/java/java_lazy_message_field_lite.lo google/protobuf/compiler/java/java_map_field.lo google/protobuf/compiler/java/java_map_field_lite.lo google/protobuf/compiler/java/java_message.lo google/protobuf/compiler/java/java_message_lite.lo google/protobuf/compiler/java/java_message_builder.lo google/protobuf/compiler/java/java_message_builder_lite.lo google/protobuf/compiler/java/java_message_field.lo google/protobuf/compiler/java/java_message_field_lite.lo google/protobuf/compiler/java/java_name_resolver.lo google/protobuf/compiler/java/java_primitive_field.lo google/protobuf/compiler/java/java_primitive_field_lite.lo google/protobuf/compiler/java/java_shared_code_generator.lo google/protobuf/compiler/java/java_service.lo google/protobuf/compiler/java/java_string_field.lo google/protobuf/compiler/java/java_string_field_lite.lo google/protobuf/compiler/java/java_doc_comment.lo google/protobuf/compiler/js/js_generator.lo google/protobuf/compiler/js/well_known_types_embed.lo google/protobuf/compiler/javanano/javanano_enum.lo google/protobuf/compiler/javanano/javanano_enum_field.lo google/protobuf/compiler/javanano/javanano_extension.lo google/protobuf/compiler/javanano/javanano_field.lo google/protobuf/compiler/javanano/javanano_file.lo google/protobuf/compiler/javanano/javanano_generator.lo google/protobuf/compiler/javanano/javanano_helpers.lo google/protobuf/compiler/javanano/javanano_map_field.lo google/protobuf/compiler/javanano/javanano_message.lo google/protobuf/compiler/javanano/javanano_message_field.lo google/protobuf/compiler/javanano/javanano_primitive_field.lo google/protobuf/compiler/objectivec/objectivec_enum.lo google/protobuf/compiler/objectivec/objectivec_enum_field.lo google/protobuf/compiler/objectivec/objectivec_extension.lo google/protobuf/compiler/objectivec/objectivec_field.lo google/protobuf/compiler/objectivec/objectivec_file.lo google/protobuf/compiler/objectivec/objectivec_generator.lo google/protobuf/compiler/objectivec/objectivec_helpers.lo google/protobuf/compiler/objectivec/objectivec_map_field.lo google/protobuf/compiler/objectivec/objectivec_message.lo google/protobuf/compiler/objectivec/objectivec_message_field.lo google/protobuf/compiler/objectivec/objectivec_oneof.lo google/protobuf/compiler/objectivec/objectivec_primitive_field.lo google/protobuf/compiler/php/php_generator.lo google/protobuf/compiler/python/python_generator.lo google/protobuf/compiler/ruby/ruby_generator.lo google/protobuf/compiler/csharp/csharp_doc_comment.lo google/protobuf/compiler/csharp/csharp_enum.lo google/protobuf/compiler/csharp/csharp_enum_field.lo google/protobuf/compiler/csharp/csharp_field_base.lo google/protobuf/compiler/csharp/csharp_generator.lo google/protobuf/compiler/csharp/csharp_helpers.lo google/protobuf/compiler/csharp/csharp_map_field.lo google/protobuf/compiler/csharp/csharp_message.lo google/protobuf/compiler/csharp/csharp_message_field.lo google/protobuf/compiler/csharp/csharp_primitive_field.lo google/protobuf/compiler/csharp/csharp_reflection_class.lo google/protobuf/compiler/csharp/csharp_repeated_enum_field.lo google/protobuf/compiler/csharp/csharp_repeated_message_field.lo google/protobuf/compiler/csharp/csharp_repeated_primitive_field.lo google/protobuf/compiler/csharp/csharp_source_generator_base.lo google/protobuf/compiler/csharp/csharp_wrapper_field.lo -lpthread libprotobuf.la -lz 
libtool: link: rm -fr  .libs/libprotoc.a .libs/libprotoc.la .libs/libprotoc.lai .libs/libprotoc.so .libs/libprotoc.so.10 .libs/libprotoc.so.10.0.0 .libs/libprotoc.so.10.0.0T
libtool: link: g++  -fPIC -DPIC -shared -nostdlib /usr/lib/gcc/x86_64-redhat-linux/7/../../../../lib64/crti.o /usr/lib/gcc/x86_64-redhat-linux/7/crtbeginS.o  google/protobuf/compiler/.libs/code_generator.o google/protobuf/compiler/.libs/command_line_interface.o google/protobuf/compiler/.libs/plugin.o google/protobuf/compiler/.libs/plugin.pb.o google/protobuf/compiler/.libs/subprocess.o google/protobuf/compiler/.libs/zip_writer.o google/protobuf/compiler/cpp/.libs/cpp_enum.o google/protobuf/compiler/cpp/.libs/cpp_enum_field.o google/protobuf/compiler/cpp/.libs/cpp_extension.o google/protobuf/compiler/cpp/.libs/cpp_field.o google/protobuf/compiler/cpp/.libs/cpp_file.o google/protobuf/compiler/cpp/.libs/cpp_generator.o google/protobuf/compiler/cpp/.libs/cpp_helpers.o google/protobuf/compiler/cpp/.libs/cpp_map_field.o google/protobuf/compiler/cpp/.libs/cpp_message.o google/protobuf/compiler/cpp/.libs/cpp_message_field.o google/protobuf/compiler/cpp/.libs/cpp_padding_optimizer.o google/protobuf/compiler/cpp/.libs/cpp_primitive_field.o google/protobuf/compiler/cpp/.libs/cpp_service.o google/protobuf/compiler/cpp/.libs/cpp_string_field.o google/protobuf/compiler/java/.libs/java_context.o google/protobuf/compiler/java/.libs/java_enum.o google/protobuf/compiler/java/.libs/java_enum_lite.o google/protobuf/compiler/java/.libs/java_enum_field.o google/protobuf/compiler/java/.libs/java_enum_field_lite.o google/protobuf/compiler/java/.libs/java_extension.o google/protobuf/compiler/java/.libs/java_extension_lite.o google/protobuf/compiler/java/.libs/java_field.o google/protobuf/compiler/java/.libs/java_file.o google/protobuf/compiler/java/.libs/java_generator.o google/protobuf/compiler/java/.libs/java_generator_factory.o google/protobuf/compiler/java/.libs/java_helpers.o google/protobuf/compiler/java/.libs/java_lazy_message_field.o google/protobuf/compiler/java/.libs/java_lazy_message_field_lite.o google/protobuf/compiler/java/.libs/java_map_field.o google/protobuf/compiler/java/.libs/java_map_field_lite.o google/protobuf/compiler/java/.libs/java_message.o google/protobuf/compiler/java/.libs/java_message_lite.o google/protobuf/compiler/java/.libs/java_message_builder.o google/protobuf/compiler/java/.libs/java_message_builder_lite.o google/protobuf/compiler/java/.libs/java_message_field.o google/protobuf/compiler/java/.libs/java_message_field_lite.o google/protobuf/compiler/java/.libs/java_name_resolver.o google/protobuf/compiler/java/.libs/java_primitive_field.o google/protobuf/compiler/java/.libs/java_primitive_field_lite.o google/protobuf/compiler/java/.libs/java_shared_code_generator.o google/protobuf/compiler/java/.libs/java_service.o google/protobuf/compiler/java/.libs/java_string_field.o google/protobuf/compiler/java/.libs/java_string_field_lite.o google/protobuf/compiler/java/.libs/java_doc_comment.o google/protobuf/compiler/js/.libs/js_generator.o google/protobuf/compiler/js/.libs/well_known_types_embed.o google/protobuf/compiler/javanano/.libs/javanano_enum.o google/protobuf/compiler/javanano/.libs/javanano_enum_field.o google/protobuf/compiler/javanano/.libs/javanano_extension.o google/protobuf/compiler/javanano/.libs/javanano_field.o google/protobuf/compiler/javanano/.libs/javanano_file.o google/protobuf/compiler/javanano/.libs/javanano_generator.o google/protobuf/compiler/javanano/.libs/javanano_helpers.o google/protobuf/compiler/javanano/.libs/javanano_map_field.o google/protobuf/compiler/javanano/.libs/javanano_message.o google/protobuf/compiler/javanano/.libs/javanano_message_field.o google/protobuf/compiler/javanano/.libs/javanano_primitive_field.o google/protobuf/compiler/objectivec/.libs/objectivec_enum.o google/protobuf/compiler/objectivec/.libs/objectivec_enum_field.o google/protobuf/compiler/objectivec/.libs/objectivec_extension.o google/protobuf/compiler/objectivec/.libs/objectivec_field.o google/protobuf/compiler/objectivec/.libs/objectivec_file.o google/protobuf/compiler/objectivec/.libs/objectivec_generator.o google/protobuf/compiler/objectivec/.libs/objectivec_helpers.o google/protobuf/compiler/objectivec/.libs/objectivec_map_field.o google/protobuf/compiler/objectivec/.libs/objectivec_message.o google/protobuf/compiler/objectivec/.libs/objectivec_message_field.o google/protobuf/compiler/objectivec/.libs/objectivec_oneof.o google/protobuf/compiler/objectivec/.libs/objectivec_primitive_field.o google/protobuf/compiler/php/.libs/php_generator.o google/protobuf/compiler/python/.libs/python_generator.o google/protobuf/compiler/ruby/.libs/ruby_generator.o google/protobuf/compiler/csharp/.libs/csharp_doc_comment.o google/protobuf/compiler/csharp/.libs/csharp_enum.o google/protobuf/compiler/csharp/.libs/csharp_enum_field.o google/protobuf/compiler/csharp/.libs/csharp_field_base.o google/protobuf/compiler/csharp/.libs/csharp_generator.o google/protobuf/compiler/csharp/.libs/csharp_helpers.o google/protobuf/compiler/csharp/.libs/csharp_map_field.o google/protobuf/compiler/csharp/.libs/csharp_message.o google/protobuf/compiler/csharp/.libs/csharp_message_field.o google/protobuf/compiler/csharp/.libs/csharp_primitive_field.o google/protobuf/compiler/csharp/.libs/csharp_reflection_class.o google/protobuf/compiler/csharp/.libs/csharp_repeated_enum_field.o google/protobuf/compiler/csharp/.libs/csharp_repeated_message_field.o google/protobuf/compiler/csharp/.libs/csharp_repeated_primitive_field.o google/protobuf/compiler/csharp/.libs/csharp_source_generator_base.o google/protobuf/compiler/csharp/.libs/csharp_wrapper_field.o   -Wl,-rpath -Wl,/Users/fanhongling/go/src/github.com/google/protobuf/src/.libs -Wl,-rpath -Wl,/usr/local/lib ./.libs/libprotobuf.so -lpthread -lz -L/usr/lib/gcc/x86_64-redhat-linux/7 -L/usr/lib/gcc/x86_64-redhat-linux/7/../../../../lib64 -L/lib/../lib64 -L/usr/lib/../lib64 -L/usr/lib/gcc/x86_64-redhat-linux/7/../../.. -lstdc++ -lm -lc -lgcc_s /usr/lib/gcc/x86_64-redhat-linux/7/crtendS.o /usr/lib/gcc/x86_64-redhat-linux/7/../../../../lib64/crtn.o  -pthread -O2 -g -Wl,--version-script=./libprotoc.map   -pthread -Wl,-soname -Wl,libprotoc.so.15 -o .libs/libprotoc.so.15.0.0
libtool: link: (cd ".libs" && rm -f "libprotoc.so.15" && ln -s "libprotoc.so.15.0.0" "libprotoc.so.15")
libtool: link: (cd ".libs" && rm -f "libprotoc.so" && ln -s "libprotoc.so.15.0.0" "libprotoc.so")
libtool: link: ar cru .libs/libprotoc.a  google/protobuf/compiler/code_generator.o google/protobuf/compiler/command_line_interface.o google/protobuf/compiler/plugin.o google/protobuf/compiler/plugin.pb.o google/protobuf/compiler/subprocess.o google/protobuf/compiler/zip_writer.o google/protobuf/compiler/cpp/cpp_enum.o google/protobuf/compiler/cpp/cpp_enum_field.o google/protobuf/compiler/cpp/cpp_extension.o google/protobuf/compiler/cpp/cpp_field.o google/protobuf/compiler/cpp/cpp_file.o google/protobuf/compiler/cpp/cpp_generator.o google/protobuf/compiler/cpp/cpp_helpers.o google/protobuf/compiler/cpp/cpp_map_field.o google/protobuf/compiler/cpp/cpp_message.o google/protobuf/compiler/cpp/cpp_message_field.o google/protobuf/compiler/cpp/cpp_padding_optimizer.o google/protobuf/compiler/cpp/cpp_primitive_field.o google/protobuf/compiler/cpp/cpp_service.o google/protobuf/compiler/cpp/cpp_string_field.o google/protobuf/compiler/java/java_context.o google/protobuf/compiler/java/java_enum.o google/protobuf/compiler/java/java_enum_lite.o google/protobuf/compiler/java/java_enum_field.o google/protobuf/compiler/java/java_enum_field_lite.o google/protobuf/compiler/java/java_extension.o google/protobuf/compiler/java/java_extension_lite.o google/protobuf/compiler/java/java_field.o google/protobuf/compiler/java/java_file.o google/protobuf/compiler/java/java_generator.o google/protobuf/compiler/java/java_generator_factory.o google/protobuf/compiler/java/java_helpers.o google/protobuf/compiler/java/java_lazy_message_field.o google/protobuf/compiler/java/java_lazy_message_field_lite.o google/protobuf/compiler/java/java_map_field.o google/protobuf/compiler/java/java_map_field_lite.o google/protobuf/compiler/java/java_message.o google/protobuf/compiler/java/java_message_lite.o google/protobuf/compiler/java/java_message_builder.o google/protobuf/compiler/java/java_message_builder_lite.o google/protobuf/compiler/java/java_message_field.o google/protobuf/compiler/java/java_message_field_lite.o google/protobuf/compiler/java/java_name_resolver.o google/protobuf/compiler/java/java_primitive_field.o google/protobuf/compiler/java/java_primitive_field_lite.o google/protobuf/compiler/java/java_shared_code_generator.o google/protobuf/compiler/java/java_service.o google/protobuf/compiler/java/java_string_field.o google/protobuf/compiler/java/java_string_field_lite.o google/protobuf/compiler/java/java_doc_comment.o google/protobuf/compiler/js/js_generator.o google/protobuf/compiler/js/well_known_types_embed.o google/protobuf/compiler/javanano/javanano_enum.o google/protobuf/compiler/javanano/javanano_enum_field.o google/protobuf/compiler/javanano/javanano_extension.o google/protobuf/compiler/javanano/javanano_field.o google/protobuf/compiler/javanano/javanano_file.o google/protobuf/compiler/javanano/javanano_generator.o google/protobuf/compiler/javanano/javanano_helpers.o google/protobuf/compiler/javanano/javanano_map_field.o google/protobuf/compiler/javanano/javanano_message.o google/protobuf/compiler/javanano/javanano_message_field.o google/protobuf/compiler/javanano/javanano_primitive_field.o google/protobuf/compiler/objectivec/objectivec_enum.o google/protobuf/compiler/objectivec/objectivec_enum_field.o google/protobuf/compiler/objectivec/objectivec_extension.o google/protobuf/compiler/objectivec/objectivec_field.o google/protobuf/compiler/objectivec/objectivec_file.o google/protobuf/compiler/objectivec/objectivec_generator.o google/protobuf/compiler/objectivec/objectivec_helpers.o google/protobuf/compiler/objectivec/objectivec_map_field.o google/protobuf/compiler/objectivec/objectivec_message.o google/protobuf/compiler/objectivec/objectivec_message_field.o google/protobuf/compiler/objectivec/objectivec_oneof.o google/protobuf/compiler/objectivec/objectivec_primitive_field.o google/protobuf/compiler/php/php_generator.o google/protobuf/compiler/python/python_generator.o google/protobuf/compiler/ruby/ruby_generator.o google/protobuf/compiler/csharp/csharp_doc_comment.o google/protobuf/compiler/csharp/csharp_enum.o google/protobuf/compiler/csharp/csharp_enum_field.o google/protobuf/compiler/csharp/csharp_field_base.o google/protobuf/compiler/csharp/csharp_generator.o google/protobuf/compiler/csharp/csharp_helpers.o google/protobuf/compiler/csharp/csharp_map_field.o google/protobuf/compiler/csharp/csharp_message.o google/protobuf/compiler/csharp/csharp_message_field.o google/protobuf/compiler/csharp/csharp_primitive_field.o google/protobuf/compiler/csharp/csharp_reflection_class.o google/protobuf/compiler/csharp/csharp_repeated_enum_field.o google/protobuf/compiler/csharp/csharp_repeated_message_field.o google/protobuf/compiler/csharp/csharp_repeated_primitive_field.o google/protobuf/compiler/csharp/csharp_source_generator_base.o google/protobuf/compiler/csharp/csharp_wrapper_field.o
libtool: link: ranlib .libs/libprotoc.a
libtool: link: ( cd ".libs" && rm -f "libprotoc.la" && ln -s "../libprotoc.la" "libprotoc.la" )
/bin/sh ../libtool  --tag=CXX   --mode=link g++ -pthread -DHAVE_PTHREAD=1 -DHAVE_ZLIB=1 -Wall -Wno-sign-compare -O2 -g -DNDEBUG -pthread  -o protoc google/protobuf/compiler/main.o -lpthread libprotobuf.la libprotoc.la -lz 
libtool: link: g++ -pthread -DHAVE_PTHREAD=1 -DHAVE_ZLIB=1 -Wall -Wno-sign-compare -O2 -g -DNDEBUG -pthread -o .libs/protoc google/protobuf/compiler/main.o  ./.libs/libprotobuf.so ./.libs/libprotoc.so /Users/fanhongling/go/src/github.com/google/protobuf/src/.libs/libprotobuf.so -lpthread -lz -pthread -Wl,-rpath -Wl,/usr/local/lib
oldpwd=`pwd` && ( cd . && $oldpwd/protoc -I. --cpp_out=$oldpwd google/protobuf/any_test.proto google/protobuf/compiler/cpp/cpp_test_bad_identifiers.proto google/protobuf/map_lite_unittest.proto google/protobuf/map_proto2_unittest.proto google/protobuf/map_unittest.proto google/protobuf/unittest_arena.proto google/protobuf/unittest_custom_options.proto google/protobuf/unittest_drop_unknown_fields.proto google/protobuf/unittest_embed_optimize_for.proto google/protobuf/unittest_empty.proto google/protobuf/unittest_enormous_descriptor.proto google/protobuf/unittest_import_lite.proto google/protobuf/unittest_import.proto google/protobuf/unittest_import_public_lite.proto google/protobuf/unittest_import_public.proto google/protobuf/unittest_lazy_dependencies.proto google/protobuf/unittest_lazy_dependencies_custom_option.proto google/protobuf/unittest_lazy_dependencies_enum.proto google/protobuf/unittest_lite_imports_nonlite.proto google/protobuf/unittest_lite.proto google/protobuf/unittest_mset.proto google/protobuf/unittest_mset_wire_format.proto google/protobuf/unittest_no_arena_lite.proto google/protobuf/unittest_no_arena_import.proto google/protobuf/unittest_no_arena.proto google/protobuf/unittest_no_field_presence.proto google/protobuf/unittest_no_generic_services.proto google/protobuf/unittest_optimize_for.proto google/protobuf/unittest_preserve_unknown_enum2.proto google/protobuf/unittest_preserve_unknown_enum.proto google/protobuf/unittest.proto google/protobuf/unittest_proto3_arena.proto google/protobuf/unittest_proto3_arena_lite.proto google/protobuf/unittest_proto3_lite.proto google/protobuf/unittest_well_known_types.proto google/protobuf/util/internal/testdata/anys.proto google/protobuf/util/internal/testdata/books.proto google/protobuf/util/internal/testdata/default_value.proto google/protobuf/util/internal/testdata/default_value_test.proto google/protobuf/util/internal/testdata/field_mask.proto google/protobuf/util/internal/testdata/maps.proto google/protobuf/util/internal/testdata/oneofs.proto google/protobuf/util/internal/testdata/proto3.proto google/protobuf/util/internal/testdata/struct.proto google/protobuf/util/internal/testdata/timestamp_duration.proto google/protobuf/util/internal/testdata/wrappers.proto google/protobuf/util/json_format_proto3.proto google/protobuf/util/message_differencer_unittest.proto google/protobuf/compiler/cpp/cpp_test_large_enum_value.proto )
touch unittest_proto_middleman
make  all-am
make[3]: Entering directory '/Users/fanhongling/go/src/github.com/google/protobuf/src'
/bin/sh ../libtool  --tag=CXX   --mode=link g++ -pthread -DHAVE_PTHREAD=1 -DHAVE_ZLIB=1 -Wall -Wno-sign-compare -O2 -g -DNDEBUG -version-info 15:0:0 -export-dynamic -no-undefined -Wl,--version-script=./libprotobuf-lite.map  -o libprotobuf-lite.la -rpath /usr/local/lib google/protobuf/stubs/atomicops_internals_x86_gcc.lo google/protobuf/stubs/atomicops_internals_x86_msvc.lo google/protobuf/stubs/bytestream.lo google/protobuf/stubs/common.lo google/protobuf/stubs/int128.lo google/protobuf/stubs/io_win32.lo google/protobuf/stubs/once.lo google/protobuf/stubs/status.lo google/protobuf/stubs/statusor.lo google/protobuf/stubs/stringpiece.lo google/protobuf/stubs/stringprintf.lo google/protobuf/stubs/structurally_valid.lo google/protobuf/stubs/strutil.lo google/protobuf/stubs/time.lo google/protobuf/arena.lo google/protobuf/arenastring.lo google/protobuf/extension_set.lo google/protobuf/generated_message_util.lo google/protobuf/generated_message_table_driven_lite.lo google/protobuf/implicit_weak_message.lo google/protobuf/message_lite.lo google/protobuf/repeated_field.lo google/protobuf/wire_format_lite.lo google/protobuf/io/coded_stream.lo google/protobuf/io/zero_copy_stream.lo google/protobuf/io/zero_copy_stream_impl_lite.lo -lpthread -lz 
libtool: link: rm -fr  .libs/libprotobuf-lite.a .libs/libprotobuf-lite.la .libs/libprotobuf-lite.lai .libs/libprotobuf-lite.so .libs/libprotobuf-lite.so.10 .libs/libprotobuf-lite.so.10.0.0
libtool: link: g++  -fPIC -DPIC -shared -nostdlib /usr/lib/gcc/x86_64-redhat-linux/7/../../../../lib64/crti.o /usr/lib/gcc/x86_64-redhat-linux/7/crtbeginS.o  google/protobuf/stubs/.libs/atomicops_internals_x86_gcc.o google/protobuf/stubs/.libs/atomicops_internals_x86_msvc.o google/protobuf/stubs/.libs/bytestream.o google/protobuf/stubs/.libs/common.o google/protobuf/stubs/.libs/int128.o google/protobuf/stubs/.libs/io_win32.o google/protobuf/stubs/.libs/once.o google/protobuf/stubs/.libs/status.o google/protobuf/stubs/.libs/statusor.o google/protobuf/stubs/.libs/stringpiece.o google/protobuf/stubs/.libs/stringprintf.o google/protobuf/stubs/.libs/structurally_valid.o google/protobuf/stubs/.libs/strutil.o google/protobuf/stubs/.libs/time.o google/protobuf/.libs/arena.o google/protobuf/.libs/arenastring.o google/protobuf/.libs/extension_set.o google/protobuf/.libs/generated_message_util.o google/protobuf/.libs/generated_message_table_driven_lite.o google/protobuf/.libs/implicit_weak_message.o google/protobuf/.libs/message_lite.o google/protobuf/.libs/repeated_field.o google/protobuf/.libs/wire_format_lite.o google/protobuf/io/.libs/coded_stream.o google/protobuf/io/.libs/zero_copy_stream.o google/protobuf/io/.libs/zero_copy_stream_impl_lite.o   -lpthread -lz -L/usr/lib/gcc/x86_64-redhat-linux/7 -L/usr/lib/gcc/x86_64-redhat-linux/7/../../../../lib64 -L/lib/../lib64 -L/usr/lib/../lib64 -L/usr/lib/gcc/x86_64-redhat-linux/7/../../.. -lstdc++ -lm -lc -lgcc_s /usr/lib/gcc/x86_64-redhat-linux/7/crtendS.o /usr/lib/gcc/x86_64-redhat-linux/7/../../../../lib64/crtn.o  -pthread -O2 -g -Wl,--version-script=./libprotobuf-lite.map   -pthread -Wl,-soname -Wl,libprotobuf-lite.so.15 -o .libs/libprotobuf-lite.so.15.0.0
libtool: link: (cd ".libs" && rm -f "libprotobuf-lite.so.15" && ln -s "libprotobuf-lite.so.15.0.0" "libprotobuf-lite.so.15")
libtool: link: (cd ".libs" && rm -f "libprotobuf-lite.so" && ln -s "libprotobuf-lite.so.15.0.0" "libprotobuf-lite.so")
libtool: link: ar cru .libs/libprotobuf-lite.a  google/protobuf/stubs/atomicops_internals_x86_gcc.o google/protobuf/stubs/atomicops_internals_x86_msvc.o google/protobuf/stubs/bytestream.o google/protobuf/stubs/common.o google/protobuf/stubs/int128.o google/protobuf/stubs/io_win32.o google/protobuf/stubs/once.o google/protobuf/stubs/status.o google/protobuf/stubs/statusor.o google/protobuf/stubs/stringpiece.o google/protobuf/stubs/stringprintf.o google/protobuf/stubs/structurally_valid.o google/protobuf/stubs/strutil.o google/protobuf/stubs/time.o google/protobuf/arena.o google/protobuf/arenastring.o google/protobuf/extension_set.o google/protobuf/generated_message_util.o google/protobuf/generated_message_table_driven_lite.o google/protobuf/implicit_weak_message.o google/protobuf/message_lite.o google/protobuf/repeated_field.o google/protobuf/wire_format_lite.o google/protobuf/io/coded_stream.o google/protobuf/io/zero_copy_stream.o google/protobuf/io/zero_copy_stream_impl_lite.o
libtool: link: ranlib .libs/libprotobuf-lite.a
libtool: link: ( cd ".libs" && rm -f "libprotobuf-lite.la" && ln -s "../libprotobuf-lite.la" "libprotobuf-lite.la" )
make[3]: Leaving directory '/Users/fanhongling/go/src/github.com/google/protobuf/src'
make[2]: Leaving directory '/Users/fanhongling/go/src/github.com/google/protobuf/src'
make[1]: Leaving directory '/Users/fanhongling/go/src/github.com/google/protobuf'
```
install
```
[vagrant@kubedev-172-17-4-59 protobuf]$ sudo make install
Making install in .
make[1]: Entering directory '/Users/fanhongling/go/src/github.com/google/protobuf'
make[2]: Entering directory '/Users/fanhongling/go/src/github.com/google/protobuf'
make[2]: Nothing to be done for 'install-exec-am'.
 /usr/bin/mkdir -p '/usr/local/lib/pkgconfig'
 /usr/bin/install -c -m 644 protobuf.pc protobuf-lite.pc '/usr/local/lib/pkgconfig'
make[2]: Leaving directory '/Users/fanhongling/go/src/github.com/google/protobuf'
make[1]: Leaving directory '/Users/fanhongling/go/src/github.com/google/protobuf'
Making install in src
make[1]: Entering directory '/Users/fanhongling/go/src/github.com/google/protobuf/src'
make  install-am
make[2]: Entering directory '/Users/fanhongling/go/src/github.com/google/protobuf/src'
make[3]: Entering directory '/Users/fanhongling/go/src/github.com/google/protobuf/src'
 /usr/bin/mkdir -p '/usr/local/lib'
 /bin/sh ../libtool   --mode=install /usr/bin/install -c   libprotobuf-lite.la libprotobuf.la libprotoc.la '/usr/local/lib'
libtool: install: /usr/bin/install -c .libs/libprotobuf-lite.so.15.0.0 /usr/local/lib/libprotobuf-lite.so.15.0.0
libtool: install: (cd /usr/local/lib && { ln -s -f libprotobuf-lite.so.15.0.0 libprotobuf-lite.so.15 || { rm -f libprotobuf-lite.so.15 && ln -s libprotobuf-lite.so.15.0.0 libprotobuf-lite.so.15; }; })
libtool: install: (cd /usr/local/lib && { ln -s -f libprotobuf-lite.so.15.0.0 libprotobuf-lite.so || { rm -f libprotobuf-lite.so && ln -s libprotobuf-lite.so.15.0.0 libprotobuf-lite.so; }; })
libtool: install: /usr/bin/install -c .libs/libprotobuf-lite.lai /usr/local/lib/libprotobuf-lite.la
libtool: install: /usr/bin/install -c .libs/libprotobuf.so.15.0.0 /usr/local/lib/libprotobuf.so.15.0.0
libtool: install: (cd /usr/local/lib && { ln -s -f libprotobuf.so.15.0.0 libprotobuf.so.15 || { rm -f libprotobuf.so.15 && ln -s libprotobuf.so.15.0.0 libprotobuf.so.15; }; })
libtool: install: (cd /usr/local/lib && { ln -s -f libprotobuf.so.15.0.0 libprotobuf.so || { rm -f libprotobuf.so && ln -s libprotobuf.so.15.0.0 libprotobuf.so; }; })
libtool: install: /usr/bin/install -c .libs/libprotobuf.lai /usr/local/lib/libprotobuf.la
libtool: warning: relinking 'libprotoc.la'
libtool: install: (cd /Users/fanhongling/go/src/github.com/google/protobuf/src; /bin/sh "/Users/fanhongling/go/src/github.com/google/protobuf/libtool"  --tag CXX --mode=relink g++ -pthread -DHAVE_PTHREAD=1 -DHAVE_ZLIB=1 -Wall -Wno-sign-compare -O2 -g -DNDEBUG -version-info 15:0:0 -export-dynamic -no-undefined -Wl,--version-script=./libprotoc.map -o libprotoc.la -rpath /usr/local/lib google/protobuf/compiler/code_generator.lo google/protobuf/compiler/command_line_interface.lo google/protobuf/compiler/plugin.lo google/protobuf/compiler/plugin.pb.lo google/protobuf/compiler/subprocess.lo google/protobuf/compiler/zip_writer.lo google/protobuf/compiler/cpp/cpp_enum.lo google/protobuf/compiler/cpp/cpp_enum_field.lo google/protobuf/compiler/cpp/cpp_extension.lo google/protobuf/compiler/cpp/cpp_field.lo google/protobuf/compiler/cpp/cpp_file.lo google/protobuf/compiler/cpp/cpp_generator.lo google/protobuf/compiler/cpp/cpp_helpers.lo google/protobuf/compiler/cpp/cpp_map_field.lo google/protobuf/compiler/cpp/cpp_message.lo google/protobuf/compiler/cpp/cpp_message_field.lo google/protobuf/compiler/cpp/cpp_padding_optimizer.lo google/protobuf/compiler/cpp/cpp_primitive_field.lo google/protobuf/compiler/cpp/cpp_service.lo google/protobuf/compiler/cpp/cpp_string_field.lo google/protobuf/compiler/java/java_context.lo google/protobuf/compiler/java/java_enum.lo google/protobuf/compiler/java/java_enum_lite.lo google/protobuf/compiler/java/java_enum_field.lo google/protobuf/compiler/java/java_enum_field_lite.lo google/protobuf/compiler/java/java_extension.lo google/protobuf/compiler/java/java_extension_lite.lo google/protobuf/compiler/java/java_field.lo google/protobuf/compiler/java/java_file.lo google/protobuf/compiler/java/java_generator.lo google/protobuf/compiler/java/java_generator_factory.lo google/protobuf/compiler/java/java_helpers.lo google/protobuf/compiler/java/java_lazy_message_field.lo google/protobuf/compiler/java/java_lazy_message_field_lite.lo google/protobuf/compiler/java/java_map_field.lo google/protobuf/compiler/java/java_map_field_lite.lo google/protobuf/compiler/java/java_message.lo google/protobuf/compiler/java/java_message_lite.lo google/protobuf/compiler/java/java_message_builder.lo google/protobuf/compiler/java/java_message_builder_lite.lo google/protobuf/compiler/java/java_message_field.lo google/protobuf/compiler/java/java_message_field_lite.lo google/protobuf/compiler/java/java_name_resolver.lo google/protobuf/compiler/java/java_primitive_field.lo google/protobuf/compiler/java/java_primitive_field_lite.lo google/protobuf/compiler/java/java_shared_code_generator.lo google/protobuf/compiler/java/java_service.lo google/protobuf/compiler/java/java_string_field.lo google/protobuf/compiler/java/java_string_field_lite.lo google/protobuf/compiler/java/java_doc_comment.lo google/protobuf/compiler/js/js_generator.lo google/protobuf/compiler/js/well_known_types_embed.lo google/protobuf/compiler/javanano/javanano_enum.lo google/protobuf/compiler/javanano/javanano_enum_field.lo google/protobuf/compiler/javanano/javanano_extension.lo google/protobuf/compiler/javanano/javanano_field.lo google/protobuf/compiler/javanano/javanano_file.lo google/protobuf/compiler/javanano/javanano_generator.lo google/protobuf/compiler/javanano/javanano_helpers.lo google/protobuf/compiler/javanano/javanano_map_field.lo google/protobuf/compiler/javanano/javanano_message.lo google/protobuf/compiler/javanano/javanano_message_field.lo google/protobuf/compiler/javanano/javanano_primitive_field.lo google/protobuf/compiler/objectivec/objectivec_enum.lo google/protobuf/compiler/objectivec/objectivec_enum_field.lo google/protobuf/compiler/objectivec/objectivec_extension.lo google/protobuf/compiler/objectivec/objectivec_field.lo google/protobuf/compiler/objectivec/objectivec_file.lo google/protobuf/compiler/objectivec/objectivec_generator.lo google/protobuf/compiler/objectivec/objectivec_helpers.lo google/protobuf/compiler/objectivec/objectivec_map_field.lo google/protobuf/compiler/objectivec/objectivec_message.lo google/protobuf/compiler/objectivec/objectivec_message_field.lo google/protobuf/compiler/objectivec/objectivec_oneof.lo google/protobuf/compiler/objectivec/objectivec_primitive_field.lo google/protobuf/compiler/php/php_generator.lo google/protobuf/compiler/python/python_generator.lo google/protobuf/compiler/ruby/ruby_generator.lo google/protobuf/compiler/csharp/csharp_doc_comment.lo google/protobuf/compiler/csharp/csharp_enum.lo google/protobuf/compiler/csharp/csharp_enum_field.lo google/protobuf/compiler/csharp/csharp_field_base.lo google/protobuf/compiler/csharp/csharp_generator.lo google/protobuf/compiler/csharp/csharp_helpers.lo google/protobuf/compiler/csharp/csharp_map_field.lo google/protobuf/compiler/csharp/csharp_message.lo google/protobuf/compiler/csharp/csharp_message_field.lo google/protobuf/compiler/csharp/csharp_primitive_field.lo google/protobuf/compiler/csharp/csharp_reflection_class.lo google/protobuf/compiler/csharp/csharp_repeated_enum_field.lo google/protobuf/compiler/csharp/csharp_repeated_message_field.lo google/protobuf/compiler/csharp/csharp_repeated_primitive_field.lo google/protobuf/compiler/csharp/csharp_source_generator_base.lo google/protobuf/compiler/csharp/csharp_wrapper_field.lo -lpthread libprotobuf.la -lz )
libtool: relink: g++  -fPIC -DPIC -shared -nostdlib /usr/lib/gcc/x86_64-redhat-linux/7/../../../../lib64/crti.o /usr/lib/gcc/x86_64-redhat-linux/7/crtbeginS.o  google/protobuf/compiler/.libs/code_generator.o google/protobuf/compiler/.libs/command_line_interface.o google/protobuf/compiler/.libs/plugin.o google/protobuf/compiler/.libs/plugin.pb.o google/protobuf/compiler/.libs/subprocess.o google/protobuf/compiler/.libs/zip_writer.o google/protobuf/compiler/cpp/.libs/cpp_enum.o google/protobuf/compiler/cpp/.libs/cpp_enum_field.o google/protobuf/compiler/cpp/.libs/cpp_extension.o google/protobuf/compiler/cpp/.libs/cpp_field.o google/protobuf/compiler/cpp/.libs/cpp_file.o google/protobuf/compiler/cpp/.libs/cpp_generator.o google/protobuf/compiler/cpp/.libs/cpp_helpers.o google/protobuf/compiler/cpp/.libs/cpp_map_field.o google/protobuf/compiler/cpp/.libs/cpp_message.o google/protobuf/compiler/cpp/.libs/cpp_message_field.o google/protobuf/compiler/cpp/.libs/cpp_padding_optimizer.o google/protobuf/compiler/cpp/.libs/cpp_primitive_field.o google/protobuf/compiler/cpp/.libs/cpp_service.o google/protobuf/compiler/cpp/.libs/cpp_string_field.o google/protobuf/compiler/java/.libs/java_context.o google/protobuf/compiler/java/.libs/java_enum.o google/protobuf/compiler/java/.libs/java_enum_lite.o google/protobuf/compiler/java/.libs/java_enum_field.o google/protobuf/compiler/java/.libs/java_enum_field_lite.o google/protobuf/compiler/java/.libs/java_extension.o google/protobuf/compiler/java/.libs/java_extension_lite.o google/protobuf/compiler/java/.libs/java_field.o google/protobuf/compiler/java/.libs/java_file.o google/protobuf/compiler/java/.libs/java_generator.o google/protobuf/compiler/java/.libs/java_generator_factory.o google/protobuf/compiler/java/.libs/java_helpers.o google/protobuf/compiler/java/.libs/java_lazy_message_field.o google/protobuf/compiler/java/.libs/java_lazy_message_field_lite.o google/protobuf/compiler/java/.libs/java_map_field.o google/protobuf/compiler/java/.libs/java_map_field_lite.o google/protobuf/compiler/java/.libs/java_message.o google/protobuf/compiler/java/.libs/java_message_lite.o google/protobuf/compiler/java/.libs/java_message_builder.o google/protobuf/compiler/java/.libs/java_message_builder_lite.o google/protobuf/compiler/java/.libs/java_message_field.o google/protobuf/compiler/java/.libs/java_message_field_lite.o google/protobuf/compiler/java/.libs/java_name_resolver.o google/protobuf/compiler/java/.libs/java_primitive_field.o google/protobuf/compiler/java/.libs/java_primitive_field_lite.o google/protobuf/compiler/java/.libs/java_shared_code_generator.o google/protobuf/compiler/java/.libs/java_service.o google/protobuf/compiler/java/.libs/java_string_field.o google/protobuf/compiler/java/.libs/java_string_field_lite.o google/protobuf/compiler/java/.libs/java_doc_comment.o google/protobuf/compiler/js/.libs/js_generator.o google/protobuf/compiler/js/.libs/well_known_types_embed.o google/protobuf/compiler/javanano/.libs/javanano_enum.o google/protobuf/compiler/javanano/.libs/javanano_enum_field.o google/protobuf/compiler/javanano/.libs/javanano_extension.o google/protobuf/compiler/javanano/.libs/javanano_field.o google/protobuf/compiler/javanano/.libs/javanano_file.o google/protobuf/compiler/javanano/.libs/javanano_generator.o google/protobuf/compiler/javanano/.libs/javanano_helpers.o google/protobuf/compiler/javanano/.libs/javanano_map_field.o google/protobuf/compiler/javanano/.libs/javanano_message.o google/protobuf/compiler/javanano/.libs/javanano_message_field.o google/protobuf/compiler/javanano/.libs/javanano_primitive_field.o google/protobuf/compiler/objectivec/.libs/objectivec_enum.o google/protobuf/compiler/objectivec/.libs/objectivec_enum_field.o google/protobuf/compiler/objectivec/.libs/objectivec_extension.o google/protobuf/compiler/objectivec/.libs/objectivec_field.o google/protobuf/compiler/objectivec/.libs/objectivec_file.o google/protobuf/compiler/objectivec/.libs/objectivec_generator.o google/protobuf/compiler/objectivec/.libs/objectivec_helpers.o google/protobuf/compiler/objectivec/.libs/objectivec_map_field.o google/protobuf/compiler/objectivec/.libs/objectivec_message.o google/protobuf/compiler/objectivec/.libs/objectivec_message_field.o google/protobuf/compiler/objectivec/.libs/objectivec_oneof.o google/protobuf/compiler/objectivec/.libs/objectivec_primitive_field.o google/protobuf/compiler/php/.libs/php_generator.o google/protobuf/compiler/python/.libs/python_generator.o google/protobuf/compiler/ruby/.libs/ruby_generator.o google/protobuf/compiler/csharp/.libs/csharp_doc_comment.o google/protobuf/compiler/csharp/.libs/csharp_enum.o google/protobuf/compiler/csharp/.libs/csharp_enum_field.o google/protobuf/compiler/csharp/.libs/csharp_field_base.o google/protobuf/compiler/csharp/.libs/csharp_generator.o google/protobuf/compiler/csharp/.libs/csharp_helpers.o google/protobuf/compiler/csharp/.libs/csharp_map_field.o google/protobuf/compiler/csharp/.libs/csharp_message.o google/protobuf/compiler/csharp/.libs/csharp_message_field.o google/protobuf/compiler/csharp/.libs/csharp_primitive_field.o google/protobuf/compiler/csharp/.libs/csharp_reflection_class.o google/protobuf/compiler/csharp/.libs/csharp_repeated_enum_field.o google/protobuf/compiler/csharp/.libs/csharp_repeated_message_field.o google/protobuf/compiler/csharp/.libs/csharp_repeated_primitive_field.o google/protobuf/compiler/csharp/.libs/csharp_source_generator_base.o google/protobuf/compiler/csharp/.libs/csharp_wrapper_field.o   -Wl,-rpath -Wl,/usr/local/lib -L/usr/local/lib -lprotobuf -lpthread -lz -L/usr/lib/gcc/x86_64-redhat-linux/7 -L/usr/lib/gcc/x86_64-redhat-linux/7/../../../../lib64 -L/lib/../lib64 -L/usr/lib/../lib64 -L/usr/lib/gcc/x86_64-redhat-linux/7/../../.. -lstdc++ -lm -lc -lgcc_s /usr/lib/gcc/x86_64-redhat-linux/7/crtendS.o /usr/lib/gcc/x86_64-redhat-linux/7/../../../../lib64/crtn.o  -pthread -O2 -g -Wl,--version-script=./libprotoc.map   -pthread -Wl,-soname -Wl,libprotoc.so.15 -o .libs/libprotoc.so.15.0.0
libtool: install: /usr/bin/install -c .libs/libprotoc.so.15.0.0T /usr/local/lib/libprotoc.so.15.0.0
libtool: install: (cd /usr/local/lib && { ln -s -f libprotoc.so.15.0.0 libprotoc.so.15 || { rm -f libprotoc.so.15 && ln -s libprotoc.so.15.0.0 libprotoc.so.15; }; })
libtool: install: (cd /usr/local/lib && { ln -s -f libprotoc.so.15.0.0 libprotoc.so || { rm -f libprotoc.so && ln -s libprotoc.so.15.0.0 libprotoc.so; }; })
libtool: install: /usr/bin/install -c .libs/libprotoc.lai /usr/local/lib/libprotoc.la
libtool: install: /usr/bin/install -c .libs/libprotobuf-lite.a /usr/local/lib/libprotobuf-lite.a
libtool: install: chmod 644 /usr/local/lib/libprotobuf-lite.a
libtool: install: ranlib /usr/local/lib/libprotobuf-lite.a
libtool: install: /usr/bin/install -c .libs/libprotobuf.a /usr/local/lib/libprotobuf.a
libtool: install: chmod 644 /usr/local/lib/libprotobuf.a
libtool: install: ranlib /usr/local/lib/libprotobuf.a
libtool: install: /usr/bin/install -c .libs/libprotoc.a /usr/local/lib/libprotoc.a
libtool: install: chmod 644 /usr/local/lib/libprotoc.a
libtool: install: ranlib /usr/local/lib/libprotoc.a
libtool: finish: PATH="/sbin:/bin:/usr/sbin:/usr/bin:/sbin" ldconfig -n /usr/local/lib
----------------------------------------------------------------------
Libraries have been installed in:
   /usr/local/lib

If you ever happen to want to link against installed libraries
in a given directory, LIBDIR, you must either use libtool, and
specify the full pathname of the library, or use the '-LLIBDIR'
flag during linking and do at least one of the following:
   - add LIBDIR to the 'LD_LIBRARY_PATH' environment variable
     during execution
   - add LIBDIR to the 'LD_RUN_PATH' environment variable
     during linking
   - use the '-Wl,-rpath -Wl,LIBDIR' linker flag
   - have your system administrator add LIBDIR to '/etc/ld.so.conf'

See any operating system documentation about shared libraries for
more information, such as the ld(1) and ld.so(8) manual pages.
----------------------------------------------------------------------
 /usr/bin/mkdir -p '/usr/local/bin'
  /bin/sh ../libtool   --mode=install /usr/bin/install -c protoc '/usr/local/bin'
libtool: install: /usr/bin/install -c .libs/protoc /usr/local/bin/protoc
 /usr/bin/mkdir -p '/usr/local/include'
 /usr/bin/mkdir -p '/usr/local/include/google/protobuf'
 /usr/bin/install -c -m 644  google/protobuf/descriptor.proto google/protobuf/any.proto google/protobuf/api.proto google/protobuf/duration.proto google/protobuf/empty.proto google/protobuf/field_mask.proto google/protobuf/source_context.proto google/protobuf/struct.proto google/protobuf/timestamp.proto google/protobuf/type.proto google/protobuf/wrappers.proto '/usr/local/include/google/protobuf'
 /usr/bin/mkdir -p '/usr/local/include/google/protobuf/compiler'
 /usr/bin/install -c -m 644  google/protobuf/compiler/plugin.proto '/usr/local/include/google/protobuf/compiler'
 /usr/bin/mkdir -p '/usr/local/include'
 /usr/bin/mkdir -p '/usr/local/include/google/protobuf'
 /usr/bin/install -c -m 644  google/protobuf/any.pb.h google/protobuf/api.pb.h google/protobuf/any.h google/protobuf/arena.h google/protobuf/arena_impl.h google/protobuf/arenastring.h google/protobuf/descriptor_database.h google/protobuf/descriptor.h google/protobuf/descriptor.pb.h google/protobuf/duration.pb.h google/protobuf/dynamic_message.h google/protobuf/empty.pb.h google/protobuf/extension_set.h google/protobuf/field_mask.pb.h google/protobuf/generated_enum_reflection.h google/protobuf/generated_enum_util.h google/protobuf/generated_message_reflection.h google/protobuf/generated_message_table_driven.h google/protobuf/generated_message_util.h google/protobuf/has_bits.h google/protobuf/implicit_weak_message.h google/protobuf/map_entry.h google/protobuf/map_entry_lite.h google/protobuf/map_field.h google/protobuf/map_field_inl.h google/protobuf/map_field_lite.h google/protobuf/map.h google/protobuf/map_type_handler.h google/protobuf/message.h google/protobuf/message_lite.h google/protobuf/metadata.h google/protobuf/metadata_lite.h google/protobuf/reflection.h google/protobuf/reflection_ops.h google/protobuf/repeated_field.h google/protobuf/service.h google/protobuf/source_context.pb.h google/protobuf/struct.pb.h google/protobuf/text_format.h google/protobuf/timestamp.pb.h '/usr/local/include/google/protobuf'
 /usr/bin/mkdir -p '/usr/local/include/google/protobuf/compiler/javanano'
 /usr/bin/install -c -m 644  google/protobuf/compiler/javanano/javanano_generator.h '/usr/local/include/google/protobuf/compiler/javanano'
 /usr/bin/mkdir -p '/usr/local/include/google/protobuf/compiler/java'
 /usr/bin/install -c -m 644  google/protobuf/compiler/java/java_generator.h google/protobuf/compiler/java/java_names.h '/usr/local/include/google/protobuf/compiler/java'
 /usr/bin/mkdir -p '/usr/local/include/google/protobuf/compiler/cpp'
 /usr/bin/install -c -m 644  google/protobuf/compiler/cpp/cpp_generator.h '/usr/local/include/google/protobuf/compiler/cpp'
 /usr/bin/mkdir -p '/usr/local/include/google/protobuf/compiler/python'
 /usr/bin/install -c -m 644  google/protobuf/compiler/python/python_generator.h '/usr/local/include/google/protobuf/compiler/python'
 /usr/bin/mkdir -p '/usr/local/include/google/protobuf/compiler/js'
 /usr/bin/install -c -m 644  google/protobuf/compiler/js/js_generator.h google/protobuf/compiler/js/well_known_types_embed.h '/usr/local/include/google/protobuf/compiler/js'
 /usr/bin/mkdir -p '/usr/local/include/google/protobuf'
 /usr/bin/install -c -m 644  google/protobuf/type.pb.h google/protobuf/unknown_field_set.h google/protobuf/wire_format.h google/protobuf/wire_format_lite.h google/protobuf/wire_format_lite_inl.h google/protobuf/wrappers.pb.h '/usr/local/include/google/protobuf'
 /usr/bin/mkdir -p '/usr/local/include/google/protobuf/stubs'
 /usr/bin/install -c -m 644  google/protobuf/stubs/atomic_sequence_num.h google/protobuf/stubs/atomicops.h google/protobuf/stubs/atomicops_internals_power.h google/protobuf/stubs/atomicops_internals_ppc_gcc.h google/protobuf/stubs/atomicops_internals_arm64_gcc.h google/protobuf/stubs/atomicops_internals_arm_gcc.h google/protobuf/stubs/atomicops_internals_arm_qnx.h google/protobuf/stubs/atomicops_internals_generic_c11_atomic.h google/protobuf/stubs/atomicops_internals_generic_gcc.h google/protobuf/stubs/atomicops_internals_mips_gcc.h google/protobuf/stubs/atomicops_internals_solaris.h google/protobuf/stubs/atomicops_internals_tsan.h google/protobuf/stubs/atomicops_internals_x86_gcc.h google/protobuf/stubs/atomicops_internals_x86_msvc.h google/protobuf/stubs/callback.h google/protobuf/stubs/bytestream.h google/protobuf/stubs/casts.h google/protobuf/stubs/common.h google/protobuf/stubs/fastmem.h google/protobuf/stubs/hash.h google/protobuf/stubs/logging.h google/protobuf/stubs/macros.h google/protobuf/stubs/mutex.h google/protobuf/stubs/once.h google/protobuf/stubs/platform_macros.h google/protobuf/stubs/port.h google/protobuf/stubs/scoped_ptr.h google/protobuf/stubs/shared_ptr.h google/protobuf/stubs/singleton.h google/protobuf/stubs/status.h google/protobuf/stubs/stl_util.h google/protobuf/stubs/stringpiece.h google/protobuf/stubs/template_util.h google/protobuf/stubs/type_traits.h '/usr/local/include/google/protobuf/stubs'
 /usr/bin/mkdir -p '/usr/local/include/google/protobuf/util'
 /usr/bin/install -c -m 644  google/protobuf/util/type_resolver.h google/protobuf/util/delimited_message_util.h google/protobuf/util/field_comparator.h google/protobuf/util/field_mask_util.h google/protobuf/util/json_util.h google/protobuf/util/time_util.h google/protobuf/util/type_resolver_util.h google/protobuf/util/message_differencer.h '/usr/local/include/google/protobuf/util'
 /usr/bin/mkdir -p '/usr/local/include/google/protobuf/compiler/php'
 /usr/bin/install -c -m 644  google/protobuf/compiler/php/php_generator.h '/usr/local/include/google/protobuf/compiler/php'
 /usr/bin/mkdir -p '/usr/local/include/google/protobuf/compiler'
 /usr/bin/install -c -m 644  google/protobuf/compiler/code_generator.h google/protobuf/compiler/command_line_interface.h google/protobuf/compiler/importer.h google/protobuf/compiler/parser.h google/protobuf/compiler/plugin.h google/protobuf/compiler/plugin.pb.h '/usr/local/include/google/protobuf/compiler'
 /usr/bin/mkdir -p '/usr/local/include/google/protobuf/compiler/ruby'
 /usr/bin/install -c -m 644  google/protobuf/compiler/ruby/ruby_generator.h '/usr/local/include/google/protobuf/compiler/ruby'
 /usr/bin/mkdir -p '/usr/local/include/google/protobuf/io'
 /usr/bin/install -c -m 644  google/protobuf/io/coded_stream.h google/protobuf/io/gzip_stream.h google/protobuf/io/printer.h google/protobuf/io/strtod.h google/protobuf/io/tokenizer.h google/protobuf/io/zero_copy_stream.h google/protobuf/io/zero_copy_stream_impl.h google/protobuf/io/zero_copy_stream_impl_lite.h '/usr/local/include/google/protobuf/io'
 /usr/bin/mkdir -p '/usr/local/include/google/protobuf/compiler/csharp'
 /usr/bin/install -c -m 644  google/protobuf/compiler/csharp/csharp_generator.h google/protobuf/compiler/csharp/csharp_names.h '/usr/local/include/google/protobuf/compiler/csharp'
 /usr/bin/mkdir -p '/usr/local/include/google/protobuf/compiler/objectivec'
 /usr/bin/install -c -m 644  google/protobuf/compiler/objectivec/objectivec_generator.h google/protobuf/compiler/objectivec/objectivec_helpers.h '/usr/local/include/google/protobuf/compiler/objectivec'
make[3]: Leaving directory '/Users/fanhongling/go/src/github.com/google/protobuf/src'
make[2]: Leaving directory '/Users/fanhongling/go/src/github.com/google/protobuf/src'
make[1]: Leaving directory '/Users/fanhongling/go/src/github.com/google/protobuf/src'
```

shared libraries
```
[vagrant@kubedev-172-17-4-59 protobuf]$ sudo ldconfig
```

bin
```
[vagrant@kubedev-172-17-4-59 protobuf]$ protoc --version
libprotoc 3.5.0
```