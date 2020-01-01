# cmiot-xmind-parser
## 这是做什么的？
cmiot-xind-parser可以读取一个或多个xmind文件，并统计树形图顶部子节点的数量（即测试点/用例的数量），便于测试时间估算及测试报告数据统计。该应用的WEB APP及更多扩展功能，敬请期待。

## 使用示例
注：如果一个.xmind文件中没有任何子节点，则会跳过该文件（日志将被打印到控制台但程序不会中止）

通配符统计多个文件(推荐)

`./xmindparser examples/*.xmind`

统计一个文件

`./xmindparser examples/map1.xmind`

统计多个文件

`./xmindparser examples/map1.xmind examples/map2.xmind`



