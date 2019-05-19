Controller提供 webapi 给其他终端，例如前端是一个端


Controller必须要负责将传入的参数，例如dto 组装成业务模型，也负责对dto的校验


普通的调用顺序应是

Web页面 -> Controller -> Service -> Repository

涉及到有些只有查询的接口，我的习惯是有时直接使用 Repository 

例如修改用户密码，

        func changePwd(userDto UserDto){
            error:=userDto.validate();
            
            if(error!=nul){
                //返回统一的异常 dto
                return &BizExceptionDto{msg:error.msg(),status:200}
            }
            
            //初始化 UserService
            userService:=UserService
            
            err:=userService.changePwd();
            
            if(err!=nil){
                //返回统一的异常 dto
                return &BizExceptionDto{msg:err.msg(),status:200}
            }
            
            //返回成功的 dto
            return CommonDto;
            
        }