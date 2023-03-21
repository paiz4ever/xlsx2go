# xlsx2go
Build xlsx files into go files

### 将xlsx文件构建成go文件
##### 功能
1. 依据xlsx的sheet和cell生成目标结构体
2. 为结构体添加初始方法读取方法
3. xlsx读取目录和go文件导出目录自定义
4. 监听xlsx文件实时构建go文件

##### 后续待完善
1. xlsx文件删除的处理
2. 实现版本化 避免重复构建
3. 支持更多的字段类型（暂int, string, bool）
4. 异常的处理
5. 更多的自定义配置