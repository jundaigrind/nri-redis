#!/bin/bash

GH_REPO=newrelic/nri-redis
GH_TAG=v1.6.0
INTEGRATION=nri-redis
PKG_VERSION=1.6.0

#RPM_ARCHS=(386 1.x86_64 arm arm64)
#RPM='${PKG_NAME}-${PKG_VERSION}-${ARCH}.rpm'
#
#DEB_ARCHS=(386 amd64 arm arm64)
#DEB='${PKG_NAME}_${PKG_VERSION}-1_${ARCH}.deb'

download_pkg () {
  PKG_NAME=$1
  URL="https://github.com/${GH_REPO}/releases/download/${GH_TAG}/${PKG_NAME}"
  printf "downloading ${URL} ... "
  set +e && curl -sS -L --fail -o ./assets/${PKG_NAME} "${URL}"
  test $? -eq 0 && echo "OK!"
}

download () {
  PKG_SCHEMA=$1
  shift
  ARCHS=("$@")

  for arch in "${ARCHS[@]}"; do
    download_pkg "$(echo $PKG_SCHEMA | sed "s/ARCH/${arch}/g")"
  done
}

WIN_ARCHS=(386 amd64)
MSI="${INTEGRATION}-ARCH.${PKG_VERSION}.msi"
download $MSI "${WIN_ARCHS[@]}"

ZIP="${INTEGRATION}-ARCH.${PKG_VERSION}.zip"
download $ZIP "${WIN_ARCHS[@]}"

TAR_ARCHS=(386 amd64 arm arm64)
TAR="${INTEGRATION}_linux_${PKG_VERSION}_ARCH.tar.gz"
download $TAR "${TAR_ARCHS[@]}"
