package sfexpresscitygo

// Endpoint
const API_ENDPOINT = "https://openic.sf-express.com"

// Interface
const (
	API_PREFIX = API_ENDPOINT + "/open/api/external/"

	API_SUFFIX = "?sign=%s"

	// pre order
	// 并非真正发单；用来验证是否可以发单并在成功时返回时效、计价等信息，也可用来验证地址以及时间是否在顺丰的配送范围内。
	PRE_CREATE_ORDER_FORMAT = API_PREFIX + "precreateorder" + API_SUFFIX
	// 并非真正取消订单，用来验证是否可以取消订单并返回计价等信息。
	PRE_CANCEL_ORDER_FORMAT = API_PREFIX + "precancelorder" + API_SUFFIX

	// order
	// 当对接方接收到订单时，可通过此接口进行物流单的创建。发单时需保证商家订单号在同一家店铺保持唯一。
	// 测试时请额外注意系统自动绑定的测试店铺，因为测试开发者共用，所以商家订单号最好与当前时间相关来避免重复。
	// 注意：除特殊说明外对接方均为非平台型接入即不需传入shop相关信息。
	CREATE_ORDER_FORMAT = API_PREFIX + "createorder" + API_SUFFIX
	// 当商家处发生异常需要取消配送时，可调用此接口对订单进行取消操作，同步返回结果。
	CANCEL_ORDER_FORMAT = API_PREFIX + "cancelorder" + API_SUFFIX

	// balance
	// 获取当前用户储值卡余额，仅支持查询中小商家店铺的储值卡余额
	GET_SHOP_ACCOUNT_BALANCE_FORMAT = API_PREFIX + "getshopaccountbalance" + API_SUFFIX

	// status
	// 此接口可获取到指定订单操作记录；当接收顺丰状态回调失败时，可以主动查询此接口补齐订单操作与状态。
	LIST_ORDER_FEED_FORMAT = API_PREFIX + "listorderfeed" + API_SUFFIX
	// 此接口可获取到指定订单的实时信息，主要包括：订单的实时状态以及骑手的信息。 注：暂时只支持60天内的数据查询。
	GET_ORDER_STATUS_FORMAT = API_PREFIX + "getorderstatus" + API_SUFFIX
	// 当订单为配送状态中，可通过该接口发起催单
	REMINDER_ORDER_FORMAT = API_PREFIX + "reminderorder" + API_SUFFIX

	// progress
	// 订单创建后，商家告知餐品制作完成，骑士侧提醒骑手。
	NOTIFY_PRODUCT_READY_FORMAT = API_PREFIX + "notifyproductready" + API_SUFFIX
	// 此接口用于获取订单配送员的实时经纬度坐标，一般情况下骑士经纬度30s更新一次。
	RIDER_LATEST_POSITION_FORMAT = API_PREFIX + "riderlatestposition" + API_SUFFIX
	// 此接口可获取一个订单的骑士位置H5链接，可进行内嵌或发送给用户（内嵌时无法保证界面的兼容性，如发现兼容性问题可使用获取配送员坐标接口自行开发轨迹H5）。
	RIDER_VIEW_V2_FORMAT = API_PREFIX + "riderviewv2" + API_SUFFIX

	// callback debug
	// 顺丰订单回调详细查看接口。可以订单维度查询所有的回调信息，并在回调信息接收出现问题的时候主动查询此接口进行订单状态同步。
	GET_CALLBACK_INFO_FORMAT = API_PREFIX + "getcallbackinfo" + API_SUFFIX

	// all callbacks
	// 配送状态更改回调
	// 骑士撤单状态回调
	// 订单完成回调
	// 顺丰原因订单取消回调
	// 订单异常回调
	// 授权回调(暂时不知道是干啥用的)
)
