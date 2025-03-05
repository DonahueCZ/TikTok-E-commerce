namespace go user

service UserService {
    // 注册
    RegisterResp Register(1: RegisterReq req) (api.post = "/register")
    // 登录
    LoginResp Login(1: LoginReq req) (api.post = "/login")
    // 退出登录
    LogoutResp Logout(1: LogoutReq req) (api.post = "/logout")
    // 删除用户
    DeleteUserResp DeleteUser(1: DeleteUserReq req) (api.post = "/deleteuser")
    // 更新用户信息
    UpdateUserResp UpdateUser(1: UpdateUserReq req) (api.post = "/updateuser")
    // 获取用户信息
    GetUserResp GetUser(1: GetUserReq req) (api.get = "/getuser")
    // 检查权限
    CheckPermissionMiddlewareResp CheckPermissionMiddleware(1: CheckPermissionMiddlewareReq req)
}


struct ResponseStatus {
    1: bool status
    2: string message
}

struct RegisterReq {
    1: string email
    2: string user_name
    3: string password
    4: string confirm_password
    5: i32 user_permissions //0是用户，1是商家
}

struct RegisterResp {
    1: i64 user_id
    2: ResponseStatus response_status
}

struct LoginReq {
    1: string email
    2: string password
}

struct LoginResp {
    1: ResponseStatus response_status
}

struct LogoutReq {
    1: i64 user_id
}

struct LogoutResp {
    1: ResponseStatus response_status
}

struct DeleteUserReq {
    1: i64 user_id
}

struct DeleteUserResp {
    1: ResponseStatus response_status
}

struct UpdateUserReq {
    1: i64 user_id // 用户ID
    2: string new_email // 新的电子邮件地址
    3: string new_user_name
    4: string current_password // 当前密码（用于验证）
    5: string new_password // 新密码（如果需要更改）
}

struct UpdateUserResp {
    1: ResponseStatus response_status
}

struct GetUserReq {
    1: i64 user_id (api.query = "user_id")
}

struct GetUserResp {
    1: i64 user_id
    2: string user_name
    3: string email
    4: string created_at
    5: string updated_at
    6: i32 user_permissions
    7: ResponseStatus response_status
}

struct CheckPermissionMiddlewareReq{
    1 : i64 user_id
}

struct CheckPermissionMiddlewareResp{
    1 : i32 user_permissions
    2 : ResponseStatus response_status
}
