# goseeker
a small search engine implemented by go

#### 这里记录下遇到的问题
2021 Nov 19
倒排链表现在在小文本量的情况下还能用varchar(xx)，一篇文章就有点顶不住了，后续该怎么处理？
用什么样的结构/方式储存？

- 现在能想到的办法就是DB里用文件指针替换要存的数据，数据单独存到文件里面
- TODO: 把倒排文件的名字后缀用uuid形式替换现在的中文