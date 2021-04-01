
# 常用命令

- 拷贝文件
kubectl cp【pod名称】:文件名称 /tmp/文件名  必须要绝对路径

- 查看支持的资源名称
kubectl api-resources --namespaced=true 

- 查看资源的具体清单
kubectl explain deployment

