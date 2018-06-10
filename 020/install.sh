#!/bin/bash
. ./goenv.sh $@

PACKAGES=`cd src && find * -maxdepth 0 -type d`
IGNORE_PATTERNS="github.com gopkg.in golang.org"
VERSION=1.0.0

for PACKAGE in ${PACKAGES}
do
	IGNORE=0
	for IGNORE_PATTERN in ${IGNORE_PATTERNS}
	do
		if [ "${PACKAGE}" = "${IGNORE_PATTERN}" ]
		then
			IGNORE=1
			break
		fi
	done
	if [ ${IGNORE} -ne 0 ]
	then
		continue
	fi
	echo install ${PACKAGE}
	go fmt ${PACKAGE}
	go install -ldflags "-s -w -X main.version=${VERSION} -X main.build=debug" ${PACKAGE}
done
