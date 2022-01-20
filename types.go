package sfexpresscitygo

// pre order
type (
	ClientParams struct {
		DevId  string `json:"dev_id"`
		DevKey string `json:"dev_key"`
	}

	Response struct {
		ErrorCode int                    `json:"error_code"`
		ErrorMsg  string                 `json:"error_msg"`
		ErrorData interface{}            `json:"error_data"`
		Result    map[string]interface{} `json:"result"`
	}

	Shop struct {
		ShopName    string `json:"shop_name"`    //(64)	空	是	店铺姓名
		ShopPhone   string `json:"shop_phone"`   //(64)	空	是	店铺电话
		ShopAddress string `json:"shop_address"` //(255)	空	是	店铺地址
		ShopLng     string `json:"shop_lng"`     //(32)	空	否	店铺经度
		ShopLat     string `json:"shop_lat"`     //(32)	空	否	店铺纬度
	}

	MultiPickupInfo struct {
		PickupShopAddress string `json:"pickup_shop_address"` //(64)	空	是	取货点地址
		PickupLng         string `json:"pickup_lng        "`  //(32)	空	是	取货点经度
		PickupLat         string `json:"pickup_lat"`          //(32)	空	是	取货点纬度
	}

	PreCreateOrderRequest struct {
		DevId            int               `json:"dev_id"`             // dev_id	int(11)	0	是	api开发者ID
		ShopId           string            `json:"shop_id"`            // shop_id	string(64)	0	是	店铺ID
		ShopType         int               `json:"shop_type"`          // shop_type	int(11)	1	否	店铺ID类型	1:顺丰店铺ID 2:接入方店铺ID
		UserLng          string            `json:"user_lng"`           // user_lng	string(32)	空	是	用户地址经度
		UserLat          string            `json:"user_lat"`           // user_lat	string(32)	空	是	用户地址纬度
		UserAddress      string            `json:"user_address"`       // user_address	string(255)	空	是	用户详细地址
		CityName         string            `json:"city_name"`          // city_name	string(32)	空	否	发单城市	用来校验是否跨城；请填写城市的中文名称，如北京市、深圳市
		Weight           int               `json:"weight"`             // weight	int(11)	0	是	物品重量(单位：克)
		ProductType      int               `json:"product_type"`       // product_type	int(11)	0	是	物品类型	枚举值见下面定义
		TotalPrice       int               `json:"total_price"`        // total_price	int(11)	0	否	用户订单总金额(单位：分)
		IsAppoint        int               `json:"is_appoint"`         // is_appoint	int(2)	0	是	是否是预约单	0：非预约单；1：预约单
		AppointType      int               `json:"appoint_type"`       // appoint_type	int(2)	0	否	预约单类型	预约单的时候传入，1：预约单送达单；2：预约单上门单
		ExpectTime       int               `json:"expect_time"`        // expect_time	int(11)	0	否	用户期望送达时间	若传入自此段且时间大于配送时效，则按照预约送达单处理，时间小于配送时效按照立即单处理；appoint_type=1时需必传,秒级时间戳
		//ExpectPickupTime int               `json:"expect_pickup_time"` // expect_pickup_time	int(11)	0	否	用户期望上门时间	appoint_type=2时需必传,秒级时间戳
		LbsType          int               `json:"lbs_type"`           // lbs_type	int(2)	2	否	坐标类型，1：百度坐标，2：高德坐标
		PayType          int               `json:"pay_type"`           // pay_type	int(11)	1	是	用户支付方式：1、已支付 0、货到付款
		ReceiveUserMoney int               `json:"receive_user_money"` // receive_user_money	int(11)	0	否	代收金额	单位：分
		IsInsured        int               `json:"is_insured"`         // is_insured	int(11)	0	是	是否保价，0：非保价；1：保价
		IsPersonDirect   int               `json:"is_person_direct"`   // is_person_direct	int(11)	0	是	是否是专人直送订单，0：否；1：是
		Vehicle          int               `json:"vehicle"`            // vehicle	int(2)	0	否	配送交通工具，0：否；1：电动车；2：小轿车
		DeclaredValue    int               `json:"declared_value"`     // declared_value	int(11)	0	否	保价金额	单位：分
		GratuityFee      int               `json:"gratuity_fee"`       // gratuity_fee	int(11)	0	否	订单小费，不传或者传0为不加小费	单位分，加小费最低不能少于100分
		RiderPickMethod  int               `json:"rider_pick_method"`  // rider_pick_method	int(11)	1	否	物流流向	1：从门店取件送至用户； 2：从用户取件送至门店
		ReturnFlag       int               `json:"return_flag"`        // return_flag	int	1	否	返回字段控制标志位(二进制)	1:商品总价格，2:配送距离，4:物品重量，8:起送时间，16:期望送达时间，32:支付费用，64:实际支持金额，128:优惠卷总金额，256:结算方式 例如全部返回为填入511
		PushTime         int               `json:"push_time"`          // push_time	int(11)	0	是	推单时间	秒级时间戳
		Shop             Shop              `json:"shop"`               // shop	obj	空	否	发货店铺信息；obj，详见shop结构	平台级开发者需要传入
		MultiPickupInfo  []MultiPickupInfo `json:"multi_pickup_info"`  // multi_pickup_info	array	空	否	多点取货信息	详细见multi_pickup_info结构
	}

	// product enums
	// 1:快餐
	// 2:药品
	// 3:百货
	// 4:脏衣服收
	// 5:干净衣服派
	// 6:生鲜
	// 8:高端饮品
	// 9:现场勘验
	// 10:快递
	// 12:文件
	// 13:蛋糕
	// 14:鲜花
	// 15:数码
	// 16:服装
	// 17:汽配
	// 18:珠宝
	// 20:披萨
	// 21:中餐
	// 22:水产
	// 27:专人直送
	// 32:中端饮品
	// 33:便利店
	// 34:面包糕点
	// 35:火锅
	// 36:证照
	// 41:外部落地配
	// 48:成人用品
	// 99:其他

	Receive struct {
		UserName    string `json:"user_name"`    //(64)	空	是	用户姓名
		UserPhone   string `json:"user_phone"`   //(64)	空	是	用户电话
		UserAddress string `json:"user_address"` //(255)	空	是	用户地址
		UserLng     string `json:"user_lng"`     //(32)	空	是	用户经度
		UserLat     string `json:"user_lat"`     //(32)	空	是	用户纬度
		CityName    string `json:"city_name"`    //(32)	空	否	发单城市	用来校验是否跨城；请填写城市的中文名称，如北京市、深圳市
	}

	ProductDetail struct {
		ProductName   string `json:"product_name"`   //(64)	空	是	物品名称
		ProductId     int    `json:"product_id"`     //(11)	0	否	物品ID
		ProductNum    int    `json:"product_num"`    //(11)	0	是	物品数量
		ProductPrice  int    `json:"product_price"`  //(11)	0	否	物品价格
		ProductUnit   string `json:"product_unit"`   //(64)	空	否	物品单位
		ProductRemark string `json:"product_remark"` //(256)	空	否	备注
		ItemDetail    string `json:"item_detail"`    //(1024)	空	否	详情
	}

	OrderDetail struct {
		TotalPrice     int             `json:"total_price"`      //(11)	0	是	用户订单商品总金额(单位：分)	100 表示1元
		ProductType    int             `json:"product_type"`     //(11)	0	是	物品类型	枚举值见下面定义
		UserMoney      int             `json:"user_money"`       //(11)	0	否	用户实付商家金额(单位：分)	100 表示1元
		ShopMoney      int             `json:"shop_money"`       //(11)	0	否	商家实收用户金额(单位：分)	100 表示1元
		WeightGram     int             `json:"weight_gram"`      //(11)	0	是	物品重量(单位：克)	100 表示100g
		VolumeLitre    int             `json:"volume_litre"`     //(11)	0	否	物品体积(单位：升)	1 表示1升
		DeliveryMoney  int             `json:"delivery_money"`   //(11)	0	否	商家收取用户的配送费(单位：分)	100 表示1元
		ProductNum     int             `json:"product_num"`      //(11)	0	是	物品个数
		ProductTypeNum int             `json:"product_type_num"` //(11)	0	是	物品种类个数
		ProductDetail  []ProductDetail `json:"product_detail"`   //空	是	物品详情；数组结构，详见productDetail结构
	}

	CreateOrderRequest struct {
		DevId               int               `json:"dev_id"`                // dev_id	int(11)	0	是	同城开发者ID	可在顺丰同城开放平台上自助申请
		ShopId              string            `json:"shop_id"`               // shop_id	string(64)	0	是	店铺ID
		ShopType            int               `json:"shop_type"`             // shop_type	int(11)	1	否	店铺ID类型	1：顺丰店铺ID ；2：接入方店铺ID
		ShopOrderId         string            `json:"shop_order_id"`         // shop_order_id	string(64)	空	是	商家订单号	不允许重复
		ShopPreparationTime int               `json:"shop_preparation_time"` // shop_preparation_time	int(64)	0	否	商家预计备餐时长	分钟级时间 比如: 10 分钟 则传入 10
		OrderSource         string            `json:"order_source"`          // order_source	string(12)	空	是	订单接入来源	1：美团；2：饿了么；3：百度；4：口碑；其他请直接填写中文字符串值
		OrderSequence       string            `json:"order_sequence"`        // order_sequence	string(12)	空	否	取货序号	与order_source配合使用，如：饿了么10号单，表示如下：order_source=2;order_sequence=10。用于骑士快速寻找配送物
		LbsType             int               `json:"lbs_type"`              // lbs_type	int(2)	2	否	坐标类型	1：百度坐标，2：高德坐标
		PayType             int               `json:"pay_type"`              // pay_type	int(11)	1	是	用户支付方式	1：已付款 0：货到付款
		OrderTime           int               `json:"order_time"`            // order_time	int(11)	0	是	用户下单时间	秒级时间戳
		IsAppoint           int               `json:"is_appoint"`            // is_appoint	int(2)	0	是	是否是预约单	0：非预约单；1：预约单
		AppointType         int               `json:"appoint_type"`          // appoint_type	int(2)	0	否	预约单类型	预约单的时候传入,1：预约单送达单；2：预约单上门单
		ExpectTime          int               `json:"expect_time"`           // expect_time	int(11)	0	否	用户期望送达时间	若传入自此段且时间大于配送时效，则按照预约送达单处理，时间小于配送时效按照立即单处理；appoint_type=1时需必传,秒级时间戳；
		ExpectPickupTime    int               `json:"expect_pickup_time"`    // expect_pickup_time	int(11)	0	否	用户期望上门时间	appoint_type=2时需必传,秒级时间戳
		ShopExpectTime      int               `json:"shop_expect_time"`      // shop_expect_time	int(11)	0	否	商家期望送达时间	只展示给骑士，不参与时效考核；秒级时间戳
		IsInsured           int               `json:"is_insured"`            // is_insured	int(11)	0	是	是否保价，0：非保价；1：保价
		IsPersonDirect      int               `json:"is_person_direct"`      // is_person_direct	int(11)	0	是	是否是专人直送订单，0：否；1：是
		Vehicle             int               `json:"vehicle"`               // vehicle	int(2)	0	否	配送交通工具，0：否；1：电动车；2：小轿车
		DeclaredValue       int               `json:"declared_value"`        // declared_value	int(11)	0	否	保价金额	单位：分
		GratuityFee         int               `json:"gratuity_fee"`          // gratuity_fee	int(11)	0	否	订单小费，不传或者传0为不加小费	单位分，加小费最低不能少于100分
		Remark              string            `json:"remark"`                // remark	string(512)	空	否	订单备注
		RiderPickMethod     int               `json:"rider_pick_method"`     // rider_pick_method	int(11)	1	否	物流流向	1：从门店取件送至用户；2：从用户取件送至门店
		ReturnFlag          int               `json:"return_flag"`           // return_flag	int	1	否	返回字段控制标志位(二进制)	1:商品总价格，2:配送距离，4:物品重量，8:起送时间，16:期望送达时间，32:支付费用，64:实际支持金额，128:优惠卷总金额，256:结算方式，例如全部返回为填入511
		PushTime            int               `json:"push_time"`             // push_time	int(11)	0	是	推单时间	秒级时间戳
		Version             int               `json:"version"`               // version	int(11)	0	是	版本号	参照文档主版本号填写，如：文档版本号1.7,version=17
		Receive             Receive           `json:"receive"`               // receive	OBJ	空	是	收货人信息	Obj，详见receive结构
		Shop                Shop              `json:"shop"`                  // shop	OBJ	空	否	发货店铺信息	Obj，详见shop结构，平台级开发者(如饿了么)需传入，如无特殊说明此字段可忽略
		OrderDetail         OrderDetail       `json:"order_detail"`          // order_detail	OBJ	空	是	订单详情	Obj，详见order_detail结构
		MultiPickupInfo     []MultiPickupInfo `json:"multi_pickup_info"`     // multi_pickup_info	array	空	否	多点取货信息	详细见multi_pickup_info结构
	}

	// cancel
	CancelOrderRequest struct {
		DevId        int    `json:"dev_id"`        //(11)	0	是	api开发者ID
		OrderId      string `json:"order_id"`      //(64)	0	是	订单ID
		OrderType    int    `json:"order_type"`    //(11)	1	是	查询订单ID类型	1、顺丰订单号 2、商家订单号
		ShopId       string `json:"shop_id"`       //(64)	0	否	店铺ID	orderType=2时必传shopId与shopType
		ShopType     int    `json:"shop_type"`     //(11)	1	否	店铺ID类型	1、顺丰店铺ID 2、接入方店铺ID
		CancelCode   int    `json:"cancel_code"`   //(11)	空	否	取消原因代码	不填时默认cancelCode=313,cancelReason=商家发起取消
		CancelReason string `json:"cancel_reason"` //(128)	空	否	其他取消原因
		PushTime     int    `json:"push_time"`     //(11)	0	是	取消时间；秒级时间戳
	}

	PreCancelOrderRequest struct {
		DevId        int    `json:"dev_id"`        //(11)	0	是	api开发者ID
		OrderId      string `json:"order_id"`      //(64)	0	是	订单ID
		OrderType    int    `json:"order_type"`    //(11)	1	是	查询订单ID类型	1、顺丰订单号 2、商家订单号
		ShopId       string `json:"shop_id"`       //(64)	0	是	店铺ID
		ShopType     int    `json:"shop_type"`     //(11)	1	否	店铺ID类型	1、顺丰店铺ID 2、接入方店铺ID
		CancelReason string `json:"cancel_reason"` //(128)	空	否	取消原因
		PushTime     int    `json:"push_time"`     //(11)	0	是	取消时间；秒级时间戳
	}

	GetShopAccountBalanceRequest struct {
		ShopId   string `json:"shop_id"`   //(64)	0	是	店铺ID
		DevId    string `json:"dev_id"`    //	0	是	开发者id
		PushTime int    `json:"push_time"` //(11)	0	是	请求时间 (秒级时间戳)
		ShopType int    `json:"shop_type"` //(11)	1	否	店铺类型	1表示顺丰店铺，2表示第三方店铺
	}

	// 顺丰订单状态：1：订单创建 2：订单取消 10：配送员接单 12：配送员到店 15：配送员配送中(已取货) 17：配送员完成订单
	ListOrderFeedRequest struct {
		DevId     int    `json:"dev_id"`     //(11)	0	是	api开发者ID
		OrderId   string `json:"order_id"`   //(64)	0	是	订单ID
		OrderType int    `json:"order_type"` //(11)	1	是	查询订单ID类型：1、顺丰订单号 2、商家订单号
		ShopId    string `json:"shop_id"`    //(64)	0	否	店铺ID	orderType=2时必传shopId与shopType
		ShopType  int    `json:"shop_type"`  //(11)	1	否	店铺ID类型	1、顺丰店铺ID 2、接入方店铺ID
		PushTime  int    `json:"push_time"`  //(11)	0	是	推送时间；秒级时间戳
	}

	GetOrderStatusRequest struct {
		DevId     int    `json:"dev_id"`     //(11)	0	是	api开发者ID
		OrderId   string `json:"order_id"`   //(64)	0	是	订单ID
		OrderType int    `json:"order_type"` //(11)	1	是	查询订单ID类型：1、顺丰订单号 2、商家订单号
		ShopId    string `json:"shop_id"`    //(64)	0	否	店铺ID	orderType=2时必传shopId与shopType
		ShopType  int    `json:"shop_type"`  //(11)	1	否	店铺ID类型	1、顺丰店铺ID 2、接入方店铺ID
		PushTime  int    `json:"push_time"`  //(11)	0	是	推送时间；秒级时间戳
	}

	ReminderOrderRequest struct {
		DevId     int    `json:"dev_id"`     //(11)	0	是	api 开发者ID
		OrderId   string `json:"order_id"`   //(64)	0	是	订单ID
		OrderType int    `json:"order_type"` //(11)	1	是	订单ID类型	1、顺丰订单号 2、商家订单号
		ShopId    string `json:"shop_id"`    //(64)	0	是	店铺ID
		ShopType  int    `json:"shop_type"`  //(11)	1	是	店铺ID类型	1：顺丰店铺ID 2：接入方店铺ID
		PushTime  int    `json:"push_time"`  //(11)	0	是	取消时间；秒级时间戳
	}

	RiderLatestPositionRequest struct {
		DevId     int    `json:"dev_id"`     //(11)	0	是	api开发者_i_d
		OrderId   string `json:"order_id"`   //(64)	0	是	订单_i_d
		OrderType int    `json:"order_type"` //(11)	1	是	查询订单_i_d类型：1、顺丰订单号 2、商家订单号
		ShopId    int    `json:"shop_id"`    //(11)	0	否	店铺_i_d	order_type=2时必传shop_id与shop_type
		ShopType  int    `json:"shop_type"`  //(11)	1	否	店铺_i_d类型	1、顺丰店铺_i_d 2、接入方店铺_i_d
		PushTime  int    `json:"push_time"`  //(11)	0	是	推送时间
	}

	RiderViewV2Request struct {
		DevId     int    `json:"dev_id"`     //(11)	0	是	开发者ID
		OrderId   string `json:"order_id"`   //(64)	0	是	订单ID
		OrderType int    `json:"order_type"` //(11)	1	是	查询订单ID类型	1、顺丰订单号 2、商家订单号
		ShopId    string `json:"shop_id"`    //(64)	0	否	店铺ID	orderType=2时必传shopId与shopType
		ShopType  int    `json:"shop_type"`  //(11)	1	否	店铺ID类型	1、顺丰店铺ID 2、接入方店铺ID
		PushTime  int    `json:"push_time"`  //(11)	空	是	推送时间
	}

	GetCallbackInfoRequest struct {
		DevId     int    `json:"dev_id"`     //(11)	0	是	api 开发者ID
		OrderId   string `json:"order_id"`   //(64)	0	是	订单ID
		OrderType int    `json:"order_type"` //(11)	1	是	订单ID类型	1、顺丰订单号 2、商家订单号
		ShopId    string `json:"shop_id"`    //(64)	0	是	店铺ID
		ShopType  int    `json:"shop_type"`  //(11)	1	是	店铺ID类型	1：顺丰店铺ID 2：接入方店铺ID
		PushTime  int    `json:"push_time"`  //(11)	0	是	取消时间；秒级时间戳
	}

	NotifyProductReadyRequest struct {
		DevId           int    `json:"dev_id"`            //(11)	0	是	api 开发者ID
		OrderId         string `json:"order_id"`          //(64)	0	是	订单ID
		OrderType       int    `json:"order_type"`        //(11)	1	是	订单ID类型	1、顺丰订单号 2、商家订单号
		ShopId          string `json:"shop_id"`           //(64)	0	是	店铺ID
		ShopType        int    `json:"shop_type"`         //(11)	1	是	店铺ID类型	1：顺丰店铺ID 2：接入方店铺ID
		NoticeReadyTime int    `json:"notice_ready_time"` //(11)	0	是	货物准备好时间，秒级时间戳
		PushTime        int    `json:"push_time"`         //(11)	0	是	推送时间；秒级时间戳
	}
)
