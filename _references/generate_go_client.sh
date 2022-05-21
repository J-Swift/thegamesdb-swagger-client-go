#!/usr/bin/env sh

set -euo pipefail

### Config

readonly yaml_file='thegamesdb.yaml'
readonly git_user='J-Swift'
readonly git_repo='thegamesdb-swagger-client-go'
readonly go_package_name='gamesdb'

readonly output_dir_openapi='out_openapi'
readonly output_dir_goswagger='out_goswagger'
readonly refs_dir_openapi="${output_dir_openapi}/_references"
readonly refs_dir_goswagger="${output_dir_goswagger}/_references"

readonly cur_file_name="$( basename $0 )"

### DO WORK SON

function main_openapi_generator() {
  local -r output_dir="${output_dir_openapi}"
  local -r refs_dir="${refs_dir_openapi}"

  rm -rf "${output_dir}"

  docker run \
    --rm \
    --volume $(pwd):/local \
    openapitools/openapi-generator-cli \
      generate \
      --input-spec /local/"${yaml_file}" \
      --generator-name go \
      --output /local/"${output_dir}" \
      --package-name "${go_package_name}" \
      --git-user-id "${git_user}" \
      --git-repo-id "${git_repo}"

  rm -rf \
    "${output_dir}"/model_*_all_of.go \
    "${output_dir}"/model_base_api_response.go \
    "${output_dir}"/model_paginated_api_response.go \
    "${output_dir}"/docs \
    && true

  gofmt -w "${output_dir}"

  mkdir "${refs_dir}"
  cp "${yaml_file}" "${cur_file_name}" client_test.go "${refs_dir}/"

  cp client_test.go "${output_dir}"
  sed -i "s/__GIT_USERNAME__/${git_user}/g" "${output_dir}/client_test.go"
  sed -i "s/__GIT_REPONAME__/${git_repo}/g" "${output_dir}/client_test.go"

  chmod +x "${output_dir}/git_push.sh"
}


function main_goswagger() {
  local -r output_dir="${output_dir_goswagger}"
  local -r refs_dir="${refs_dir_goswagger}"

  rm -rf "${output_dir}"
  mkdir -p "${output_dir}/src"

  docker run \
    --rm \
    --interactive \
    --tty \
    --env GOPATH=/local/${output_dir} \
    --volume $(pwd):/local \
    --workdir /local \
    quay.io/goswagger/swagger \
      generate client \
      --spec=/local/"${yaml_file}" \
      --target="${output_dir}/src"

  mkdir "${refs_dir}"
  cp "${yaml_file}" "${cur_file_name}" "${refs_dir}/"
}

function main() {
  local -r generator="${1:-openapi}"

  case "${generator}" in
    goswagger)
      echo '> Generating with goswagger...'
      echo
      main_goswagger
      ;;
    openapi)
      echo '> Generating with openapi...'
      echo
      main_openapi_generator
      ;;
    *)
      echo "[ERROR] unrecognized generator [${generator}]"
      exit 1
  esac

  echo
  echo '> Done'
}

main $@


