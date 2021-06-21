// 该程序将打印著名的"Hello, world",然后终止程序


argv	IS		$1  			//命令行参数向量
		LOC		#100
Main	LDOU	$255,argv,0 	//$255 <- 程序名的地址
		TRAP	0,Fputs,StdOut  //打印程序名
		GETA	$255,String 	//$255 <- ", world"的地址
		TRAP	0,Fputs,StdOut  //打印", world"
		TRAP	0,Halt,0 		//终止程序
String	BYTE	", world",#a,0  //带有换行符和结束符的字符串


