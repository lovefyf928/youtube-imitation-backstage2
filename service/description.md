业务逻辑写在这里

例如 

用户修改密码

        changePwd(){
            
            //初始化 UserRepository
            user:=UserRepository
            
            //判断用户密码是否有效
            //判断用户密码是否和原来是一致的
            
            //修改用户密码
            UserRepository.changeUserPwd(userId,newPwd); 
        }