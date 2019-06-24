
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
		code: 状态值， {200: 成功，500: 异常，403: 权限异常},
		data: {
			data_list: [],
			total:0
		},
		message: 返回消息
	}

删除电影： delMovie		=>	/movie 			=> movie.MovieHandler
	请求： delete
	参数：{
		id: 电影ID，
	}
	返回：{
		code: 状态值， {200: 成功，500: 异常，403: 权限异常},
		data: {},
		message: 返回消息
	}
	
电影详情： getMovieDetail 	=> /movie/detail 	=> movie_detail.MovieDetailHandler
	请求： get
	参数： {
		title: 电影标题
	}
	返回：{
		code: 状态值， {200: 成功，500: 异常，403: 权限异常},
		data: {
			movie: {},	// 电影基本信息
			images: [], // 电影截图
			links:[],   // 电影链接
		},
		message: 返回消息
	}

演员管理：getPerformer 	=> /performer 		=> performer.PerformerHandler
	请求： get
	参数：{
		page: 当前页,
		page_size: 每页大小
	}
	返回：返回：{
		code: 状态值， {200: 成功，500: 异常，403: 权限异常},
		data: {
			data_list: [],
			total:0
		},
		message: 返回消息
	}

删除演员： delPerformer	=>	/performer 		=> performer.PerformerHandler
	请求： delete
	参数：{
		name : ""
	}
	返回：返回：{
		code: 状态值， {200: 成功，500: 异常，403: 权限异常},
		data: {},
		message: 返回消息
	}
	}
	
演员详情： getPerformerDetail 	=> /performer/detail 	=> performer_detail.PerformerDetailHandler
	请求： get
	参数： {
		name: 演员名称
		page: 1
		page_size: 20
	}
	返回：{
		code: 状态值， {200: 成功，500: 异常，403: 权限异常},
		data: {
			performer: {},	// 演员基本信息
			movies: [], // 演员电影信息
			total:0
		},
		message: 返回消息
	}
	
注册：register     => /register    =>  register.RegisterHandler
    请求： post
    参数： {
        username: "",
        password: "",
        nickname: ""
    }
    返回：{
        code: 状态值， {200: 成功，500: 异常，403: 权限异常},
        message: 返回消息
    }

登录：login        => /login       => login.LoginHandler
    请求： post
    参数： {
        username: "",
        password: "",
    }
    返回：{
        code: 状态值， {200: 成功，500: 异常，403: 权限异常},
        message: 返回消息
    }
    
收藏：
查看收藏： getMoviecol   => /moviecol    => moviecol.MoviecolHandler
    请求： get
    参数： {
        user_id: 用户ID, 默认为空
        page: 1
        page_size: 20
    } 
    返回：{
        code: 状态值， {200: 成功，500: 异常，403: 权限异常},
        data: {
            movies: [], // 演员电影信息
            total:0
        },
        message: 返回消息
    }
添加收藏： addMoviecol   => /moviecol    => moviecol.MoviecolHandler
    请求： post
    参数： {
        user_id: 用户ID,
        title: 电影标题
    }
    返回：{
        code: 状态值， {200: 成功，500: 异常，403: 权限异常},
        message: 返回消息
    }
删除收藏： deleteMoviecol   => /moviecol    => moviecol.MoviecolHandler
    请求： delete
    参数： {
        user_id: 用户ID,
        title: 电影标题
    }
    返回：{
        code: 状态值， {200: 成功，500: 异常，403: 权限异常},
        message: 返回消息
    }

关注： 
查看关注： getFollow   => /follow      => follow.FollowHandler
    请求： get
    参数： {
        user_id: 用户ID, 默认为空
        page: 1
        page_size: 20
    } 
    返回：{
        code: 状态值， {200: 成功，500: 异常，403: 权限异常},
        data: {
            performer: [], // 演员电影信息
            total:0
        },
        message: 返回消息
    }
添加关注： addFollow   => /follow    => moviecol.FollowHandler
    请求： post
    参数： {
        user_id: 用户ID,
        performer: 演员名称
    }
    返回：{
        code: 状态值， {200: 成功，500: 异常，403: 权限异常},
        message: 返回消息
    }
删除关注： deleteFollow   => /follow    => moviecol.FollowHandler
    请求： delete
    参数： {
        user_id: 用户ID,
        performer: 演员名称
    }
    返回：{
        code: 状态值， {200: 成功，500: 异常，403: 权限异常},
        message: 返回消息
    }