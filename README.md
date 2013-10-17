regex
=====
a regex of go

go实现的NFA为底层的regex

	支持匹配边界；
	支持捕获分组、模式；
	支持匹配编号分组的匹配结果；
	支持匹配编号的模式；
	支持匹配自身；
	支持固化分组；
	支持前向、后向的预查；
	不支持分组的命名；
	不支持unicode预定义字符组；
	对中文支持不好，建议只用于匹配utf-8格式的序列；
	分组的再匹配（结果/表达式）与perl、posix不同，参见下述；

自定义字符集

  	[]	：'[]'允许内部的'[]'嵌套，需要注意；
	^	：位于'['后表示取互补的字符集；否则表示'^'
	-	：如两侧都是指代单个字符，表示取范围；否则表示'-'

预设字符集

	\s	：[\r\n\t\f\v ]
	\S	：[^\r\n\t\f\v ]
	\w	：[a-zA-Z0-9_]
	\W	：[^a-zA-Z0-9_]
	\c	：[a-zA-Z]
	\C	：[^a-zA-Z]
	\d	：[0-9]
	\D	：[^0-9]
	.  	：[^\x00]

转义字符

	\r		：回车符
	\n		：换行符
	\t		：水平制表符
	\v		：垂直制表符
	\f		：换页符
	\0		：NULL
	\ooo	：oct数字表示的字符
	\xhh	：hex数字表示的字符

边界

	\a	：文本开始
	\A	：非文本开始
	\b	：词首或词尾
	\B	：非词首且非词尾
	\z	：文本结尾
	\Z	：非文本结尾
	^	：行首或文本开始
	$	：行尾或文本结尾

子匹配

	@{dd}	：根据编号匹配某个分组的匹配结果，无效组会被忽略
	#{dd}	：根据编号匹配某个模式，0为自身，无效组会被忽略

分组

	(		：编号并捕获分组
	(?:		：编号并捕获模式
	(?>		：固化分组，移动
	(?=		：向后确准，不移动
	(?!		：向后排除，不移动
	(?<=	：向前确准，不移动
	(?<!	：向前排除，不移动
	) 		：……

一元后置算符

	?	：重复0或1次，次数多优先
	??	：重复0或1次，次数少优先
	*	：重复0或多次，次数多优先
	*?	：重复0或多次，次数少优先
	+	：重复1或多次，次数多优先
	+?	：重复1或多次，次数少优先

二元左优先算符

	|	：选择
	&	：连续（隐藏算符）

匹配过程
	
	使用堆栈的深度搜索 + 回溯算法
