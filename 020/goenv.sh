#!/bin/bash
export GOPATH=$(pwd)

case $1 in
	aarch64)
		export PATH=$PATH:/opt/toolchain/golang/gcc-linaro-7.2.1-2017.11-x86_64_aarch64-linux-gnu/bin
		export GOOS=linux
		export GOARCH=arm64
		export GOARM=
		export CGO_ENABLED=1
		export CC=aarch64-linux-gnu-gcc
		export CXX=aarch64-linux-gnu-g++
		;;
	arm-eabi)
		export PATH=$PATH:/opt/toolchain/golang/gcc-linaro-7.2.1-2017.11-x86_64_arm-eabi/bin
		export GOOS=linux
		export GOARCH=arm
		export GOARM=6
		export CGO_ENABLED=1
		export CC=arm-eabi-gcc
		export CXX=arm-eabi-g++
		;;
	arm-linux-eabi)
		export PATH=$PATH:/opt/toolchain/golang/gcc-linaro-7.2.1-2017.11-x86_64_arm-linux-gnueabi/bin
		export GOOS=linux
		export GOARCH=arm
		export GOARM=6
		export CGO_ENABLED=1
		export CC=arm-linux-gnueabi-gcc
		export CXX=arm-linux-gnueabi-g++
		;;
	arm-linux-eabihf)
		export PATH=$PATH:/opt/toolchain/golang/gcc-linaro-7.2.1-2017.11-x86_64_arm-linux-gnueabihf/bin
		export GOOS=linux
		export GOARCH=arm
		export GOARM=6
		export CGO_ENABLED=1
		export CC=arm-linux-gnueabihf-gcc
		export CXX=arm-linux-gnueabihf-g++
		;;
	linux)
		export GOOS=linux
		export GOARCH=amd64
		export GOARM=
		export CGO_ENABLED=1
		export CC=gcc
		export CXX=g++
		;;
	win32)
		export PATH=$PATH:/opt/toolchain/golang/mingw-w32-bin_x86_64-linux_20131227/bin

		export GOOS=windows
		export GOARCH=386
		export GOARM=
		export CGO_ENABLED=1
		export CC=i686-w64-mingw32-gcc
		export CXX=i686-w64-mingw32-g++
		;;
	win64)
		export PATH=$PATH:/opt/toolchain/golang/mingw-w64-bin_x86_64-linux_20131228/bin

		export GOOS=windows
		export GOARCH=amd64
		export GOARM=
		export CGO_ENABLED=1
		export CC=x86_64-w64-mingw32-gcc
		export CXX=x86_64-w64-mingw32-g++
		;;
esac

go env
