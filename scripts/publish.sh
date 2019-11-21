#! /bin/sh

# 目录不存在则创建目录
if [ ! -d "./dist" ]; then
  mkdir -p -m 755 ./dist
fi

# 追加配置及资源文件
ZIP_FILES="conf/cert.pem conf/key.pem conf/model.conf conf/config.yml"

# 打包压缩文件
tar -czf dist/"${1}_${2}".tar.gz "${ZIP_FILES}"
