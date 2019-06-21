电影管理：
获取电影：getMovie		=> 	/movie 			=> movie.MovieHandler
	请求：get
	参数：{
		keyword:  电影名, 
		performer: 角色名，默认为空
		genre: 类型， 默认为空
		series: 系列名，默认为空
		page: 1, 页面数
		page_size: 50，每页大小
	}
	返回：{
		code: 状态值， {200: 成功，500: 异常，400: 权限异常},
		data: {},
		message: 返回消息
	}

删除电影： delMovie		=>	/movie 			=> movie.MovieHandler
	请求： delete
	参数：{
		id: 电影ID，

	}