# xlsx2go
Build xlsx files into go files

### 将xlsx文件构建成go文件
##### 说明
1. 在build.json中分别配置xlsx目录位置\<input\> | 构建导出目录\<output\> | 构建文件包名\<package\>
2. 在xlsx目录文件创建xlsx 其中sheet的名字对应到go结构体的名字
3. 命名规则
    |字段1|字段2|字段3|
    |:-:|:-:|:-:|
    |int|string|bool|
4. 跑程序在导出目录生成go文件
5. 程序持续监听xlsx变动去重新构建go文件


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