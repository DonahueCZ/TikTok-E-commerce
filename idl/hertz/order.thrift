namespace go order

service OrderService {
    BaseResponse CreateOrder(1: CreateOrderRequest req) (api.post = "/order");
    OrderResponse GetOrder(1: OrderRequest req) (api.get = "/order");
    OrderListResponse GetUserOrders(1: UserOrderRequest req) (api.get = "/user/orders");
    BaseResponse UpdateOrder(1: UpdateOrderRequest req) (api.put = "/order");
    BaseResponse UpdateOrderStatus(1: OrderStatusRequest req) (api.post = "/order/status");
    BaseResponse UpdateOrderAddress(1: OrderAddressRequest req) (api.post = "/order/address");
}

struct OrderRequest {
    1: i64 orderId (api.query = "id")
}

struct UserOrderRequest {
    1: i64 userId (api.query = "userId")
    2: i32 page (api.query = "page")
    3: i32 size (api.query = "size")
}

struct CreateOrderRequest {
    1: i64     userId
    2: i64     goodsId
    3: i32     goods_count
    4: i32     cost
    5: Address address
}

struct UpdateOrderRequest {
    1: i64     orderId
    2: i32     status
    3: i32     goods_count
    4: i32     cost
    5: Address address
}

struct OrderStatusRequest {
    1: i64 orderId
    2: i32 status
}

struct OrderAddressRequest {
    1: i64     orderId
    2: Address address
}

struct Address {
    1: string name
    2: string phone
    3: string address
}

struct Order {
    1: i64     id
    2: i64     userId
    3: i64     goodsId
    4: i32     status
    5: i64     createTime
    6: Address address
    7: i32     goods_count
    8: i32     cost
}

struct OrderResponse {
    1: Order order
}

struct OrderListResponse {
    1: list<Order> orders
}

struct BaseResponse {
    1: i32    code
    2: string message
}