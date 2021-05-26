#!/usr/bin/env bash
source_dir="./source"
target_dir="./target"
[[ ! -d "${target_dir}" ]] && mkdir target_dir
arcname="${target_dir}/$(date +'%Y-%m-%d_%H:%M:%S').zip"
make_archive="zip -q -r ${arcname} ${source_dir}"
if ${make_archive}; then
    echo "Successful backup to ${arcname}"
else
    echo "Backup FAILED"
fi
