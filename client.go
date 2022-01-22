package sfexpresscitygo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SFExpressCityClient struct {
	signer *Signer
}

func NewSFExpressCityClient(param *ClientParams) (*SFExpressCityClient, error) {
	return &SFExpressCityClient{
		signer: NewSigner(param.DevId, param.DevKey),
	}, nil
}

// {
//      "error_code": 0,
//      "error_msg": "",
//      "error_data": null,
//      "result": {
//          "delivery_type": 1,//0：预约送达单 1：立即单 3：预约上门单
//          "expect_time": 1546845547,//预计送达（上门）时间
//          "start_time": 1546841347,//预计开始配送的时间
//          "promise_delivery_time": 70,//预计配送时间（单位: 分）
//          "delivery_distance_meter": "6145",//配送距离
//          "charge_price_list": {
//              "shop_pay_price": 1200,//配送费总额（单位:分）
//              "charges_detail": {
//                  "basic_fee": 1100,//常规配送费=起步价+超距离费+超重量费
//                  "basic": 900,//起步价
//                  "over_distance": 100,//超距离费用
//                  "over_weight": 0,//超重量费用
//                  "cancel_excess_fee": 0,//拒收扣费
//                  "special_time_fee": 0,//特殊时段费
//                  "vas_fee": 0,//增值服务费
//                  "vas_fee_detail": {//增值服务费详情
//                      "packing_fee": 0,//包材费
//                      "low_temp_fee": 0,//低温服务费
//                      "take_goods_sms_fee": 0,//取货短信费
//                      "insured": {
//                          "fee": 0,//保价费
//                          "declared_price": 0
//                      },
//                      "big_order": {
//                          "fee": 0,//大额单费
//                          "amount": 0
//                      },
//                      "collection": {
//                          "fee": 0,//代收货款费用
//                          "price": 0
//                      },
//                      "person_direct_fee": 0, //专人直送费用
//                      "vehicle_car_fee": 0 //小轿车配送费用
//                  },
//                  "extra_fee": 0,//附加费
//                  "extra_fee_detail": {
//                      "geography_fee": 0
//                  }
//              },
//          },
//          "gratuity_fee": 100, //订单小费
//          "push_time": 1546841347
//          //以下字段受请求参数中 return_flag 控制：return_flag中未包含的，此字段将不存在，请注意！
//         "total_price": 1300, //配送费总额，当return_flag中包含1时返回，单位分（值为计算出来此单总价）
//         "delivery_distance_meter": 1234, //配送距离，当return_flag中包含2时返回，单位米（值为计算出来实际配送距离）
//         "weight_gram": 1000, //商品重量，当return_flag中包含4时返回，单位克（值为下单传入参数回传）
//         "start_time": 123456789, //起送时间，当return_flag中包含8时返回，时间格式为Unix时间戳，注意转换
//         "expect_time": 123456789, //期望送达时间，当return_flag中包含16时返回，时间格式为Unix时间戳，注意转换
//         "total_pay_money": 1300， //支付费用，当return_flag中包含32时返回，单位分
//         "real_pay_money": 1300 //实际支付金额，当return_flag中包含64时返回，单位分（实际支付金额=总金额-优惠卷总金额）
//         "overflow_fee": 200, //爆单费，单位分
//         "settlement_type": 1 //结算方式，当return_flag中包含256时返回
//      }
//  }
func (c *SFExpressCityClient) PreCreateOrder(request *PreCreateOrderRequest) (int, *Response, error) {
	code, resp, err := c.do(PRE_CREATE_ORDER_FORMAT, request)
	if err != nil {
		return code, resp, err
	}
	// TODO format response
	return code, resp, nil
}

// {
// 	"error_code":0,
// 	"error_msg":"",
// 	"error_data":null,
// 	"result":{
// 			"order_type":1,
// 			"promise_delivery_time":75, //承诺配送时长(分)
// 			"delivery_type":1, //订单类型 0：预约送达单，1：立即单，2：预约上门单
// 			"expect_pickup_time":0, //原始期望上门时间
// 			"expect_time":1607006190, //预约时间
// 			"shop_cancel_times":0, //店铺每日取消次数
// 			"free_cancel_times":0, //每日免费取消次数
// 			"is_cancel_charge_price_rule":0, //取消收费规则
// 			"is_over_free_cancel_times":1, //是否超出免费取消次数
// 			"is_deduction_fee":0, //是否有取消收费
// 			"deduction_fee":0, //取消收费
// 			"cancel_charge_price_rule":[ //取消收费规则

// 			],
// 			"could_cancel":true, //能否取消
// 			"push_time":1607003598, //推单时间
// 			"sf_order_id":"3370126808465994769", //顺丰订单号
// 			"shop_order_id":"223044197298" //商家订单号
// 	}
// }
func (c *SFExpressCityClient) PreCancelOrder(request *PreCancelOrderRequest) (int, *Response, error) {
	code, resp, err := c.do(PRE_CANCEL_ORDER_FORMAT, request)
	if err != nil {
		return code, resp, err
	}
	// TODO format response
	return code, resp, nil
}

// {
// 	"error_code": 0,
// 	"error_msg": "",
// 	"error_data": null,//详细报错信息（报错的时候非空）
// 	"result": {
// 			"sf_order_id": "3165848793513984",//顺丰订单号（标准默认为int，可以设置为string）
// 			"sf_bill_id": "509008343346",//顺丰运单号（需要设置）
// 			"shop_order_id": "15104092022333",//商家订单号
// 			"push_time": "1510680568" //推送时间
// 			//以下字段受请求参数中 return_flag 控制：return_flag中未包含的，此字段将不存在，请注意！
// 			"total_price": 1300, //配送费总额，当return_flag中包含1时返回，单位分（值为计算出来此单总价）
// 			"delivery_distance_meter": 1234, //配送距离，当return_flag中包含2时返回，单位米（值为计算出来实际配送距离）
// 			"weight_gram": 1000, //商品重量，当return_flag中包含4时返回，单位克（值为下单传入参数回传）
// 			"start_time": 123456789, //起送时间，当return_flag中包含8时返回，时间格式为Unix时间戳，注意转换
// 			"expect_time": 123456789, //期望送达时间，当return_flag中包含16时返回，时间格式为Unix时间戳，注意转换
// 			"total_pay_money": 1300, //支付费用，当return_flag中包含32时返回，单位分
// 			"real_pay_money": 1300, //实际支付金额，当return_flag中包含64时返回，单位分（实际支付金额=总金额-优惠卷总金额）
// 			"coupons_total_fee": 1300, //优惠券总金额，当return_flag中包含128时返回，单位分
// 			"settlement_type": 1, //结算方式，当return_flag中包含256时返回
// 			"pickup_code ": 1, //取件码。在顺丰同城商户侧配置，配置后有此字段。
// 			"complete_code ": 1, //签收码。在顺丰同城商户侧配置，配置后有此字段。
// 			"overflow_fee": 200, //爆单费，单位分
// 	}
// }
func (c *SFExpressCityClient) CreateOrder(request *CreateOrderRequest) (int, *Response, error) {
	code, resp, err := c.do(CREATE_ORDER_FORMAT, request)
	if err != nil {
		return code, resp, err
	}
	// TODO format response
	return code, resp, nil
}

// {
// 	"error_code": 0,
// 	"error_msg": "",
// 	"error_data": null,
// 	"result": {
// 			"sf_order_id": 3165848793513984,//顺丰订单
// 			"shop_order_id": "15104092022333",//商家订单号
// 			"deduction_detail": {
// 					"deduction_fee": 10,//取消收费金额（单位：分）
// 					"shop_cancel_times": 0,//店铺维度累计的取消次数
// 					"free_cancel_times": 2,//配置的免费取消次数
// 			},
// 			"push_time": 1609165317//接口返回时间
// 	}
// }
func (c *SFExpressCityClient) CancelOrder(request *CancelOrderRequest) (int, *Response, error) {
	code, resp, err := c.do(CANCEL_ORDER_FORMAT, request)
	if err != nil {
		return code, resp, err
	}
	// TODO format response
	return code, resp, nil
}

// {
// 	"error_code": 0,
// 	"error_msg": "success",
// 	"error_data": {
// 			"desc":"", //错误信息。和error_msg字段相同
// 	},
// 	"result": {
// 			"balance":3000, //账户总余额，单位：分
// 			"reserved_amount":0, // 预留余额，单位：分
// 			"withdraw_amount":3000 //实际可提现月，单位：分
// 	}
// }
func (c *SFExpressCityClient) GetShopAccountBalance(request *GetShopAccountBalanceRequest) (int, *Response, error) {
	code, resp, err := c.do(GET_SHOP_ACCOUNT_BALANCE_FORMAT, request)
	if err != nil {
		return code, resp, err
	}
	// TODO format response
	return code, resp, nil
}

// {
// 	"error_code": 0,
// 	"error_msg": "",
// 	"result": {
// 			"sf_order_id": 3171874687750684700,
// 			"shop_order_id": "1000007894570",
// 			"create_time": 1512467712,//订单创建时间
// 			"push_time": 1512467712,//当前推送时间
// 			"feed": [
// 					{
// 							"sf_order_id": 3171874687750684700,
// 							"shop_order_id": "1000007894570",
// 							"order_status": 1,
// 							"operator": "【门店生成订单】门店ID3241827460869",
// 							"operator_name": "",
// 							"operator_phone":"",
// 							"status_desc": "订单提交成功",
// 							"content": "订单号：3171874687750684673",
// 							"create_time": "2017-12-05 17:55:39"
// 					},
// 					{
// 							"sf_order_id": 3171874687750684700,
// 							"shop_order_id": "1000007894570",
// 							"order_status": 10,
// 							"operator": "lize",
// 							"operator_name": "lize_test",
// 							"operator_phone": 18811588750,
// 							"status_desc": "配送员已接单",
// 							"content": "配送员接单：lize，电话：18811588750",
// 							"create_time": "2017-12-05 18:04:53"
// 					},
// 					{
// 							"sf_order_id": 3171874687750684700,
// 							"shop_order_id": "1000007894570",
// 							"order_status": 10,
// 							"operator": "lize",
// 							"operator_name": "lize_test",
// 							"operator_phone": 18811588750,
// 							"status_desc": "配送员已接单",
// 							"content": "配送员接单：lize，电话：18811588750",
// 							"create_time": "2017-12-05 18:05:17"
// 					},
// 					{
// 							"sf_order_id": 3171874687750684700,
// 							"shop_order_id": "1000007894570",
// 							"order_status": 12,
// 							"operator": "lize",
// 							"operator_name": "lize_test",
// 							"operator_phone": 18811588750,
// 							"status_desc": "配送员已到店",
// 							"content": "配送员已到店：配送员：lize，配送员电话：18811588750",
// 							"create_time": "2017-12-05 18:05:50"
// 					},
// 					{
// 							"sf_order_id": 3171874687750684700,
// 							"shop_order_id": "1000007894570",
// 							"order_status": 14,
// 							"operator": "lize",
// 							"operator_name": "lize_test",
// 							"operator_phone": 18811588750,
// 							"status_desc": "配送员已取货",
// 							"content": "正在配送中：配送员：lize，配送员电话：18811588750",
// 							"create_time": "2017-12-05 18:06:17"
// 					},
// 					{
// 							"sf_order_id": 3171874687750684700,
// 							"shop_order_id": "1000007894570",
// 							"order_status": 17,
// 							"operator": "lize",
// 							"operator_name": "lize_test",
// 							"operator_phone": 18811588750,
// 							"status_desc": "配送员完成",
// 							"content": "配送员点击完成，配送员：，配送员电话：",
// 							"create_time": "2017-12-05 18:07:09"
// 					}
// 			]
// 	}
// }
func (c *SFExpressCityClient) ListOrderFeed(request *ListOrderFeedRequest) (int, *Response, error) {
	code, resp, err := c.do(LIST_ORDER_FEED_FORMAT, request)
	if err != nil {
		return code, resp, err
	}
	// TODO format response
	return code, resp, nil
}

// {
// 	"error_code": 0,
// 	"error_msg": "",
// 	"error_data": null,
// 	"result": {
// 			"order_id": 3368153620852828000,//顺丰同城订单ID
// 			"shop_id": 3256895871000,//顺丰店铺ID
// 			"out_order_id": "28930139715870",//商家订单ID
// 			"order_status": 2,//当前状态
// 			"status_desc": "订单已取消",//当前状态描述
// 			"rider_name": "lize",//骑士名称
// 			"rider_phone": "188115000",//骑手电话
// 			"push_time": 1606060800,//订单推送生成的时间
// 	}
// }
func (c *SFExpressCityClient) GetOrderStatus(request *GetOrderStatusRequest) (int, *Response, error) {
	code, resp, err := c.do(GET_ORDER_STATUS_FORMAT, request)
	if err != nil {
		return code, resp, err
	}
	// TODO format response
	return code, resp, nil
}

// {
// 	"error_code": 0,//只要是成功errno=0，非0 代表催单失败；
// 	"error_msg": "", //催单失败返回相应的文案
// 	"result": {
// 			"sf_order_id": "3165848793513984",//顺丰订单号
// 			"shop_order_id": "15104092022333",//商家订单号
// 			"push_time": "1510680568" //推送时间
// 	}
// }
func (c *SFExpressCityClient) ReminderOrder(request *ReminderOrderRequest) (int, *Response, error) {
	code, resp, err := c.do(REMINDER_ORDER_FORMAT, request)
	if err != nil {
		return code, resp, err
	}
	// TODO format response
	return code, resp, nil
}

// {
// 	"error_code": 0,
// 	"error_msg": "",
// 	"error_data":null,
// 	"result":{
// 			"sf_order_id":"3387878149299501057", //顺丰订单号
// 			"shop_order_id":"15104092022333", //商家订单号
// 			"push_time":"1513333333",//推送时间
// 	}
// }
func (c *SFExpressCityClient) NotifyProductReady(request *NotifyProductReadyRequest) (int, *Response, error) {
	code, resp, err := c.do(NOTIFY_PRODUCT_READY_FORMAT, request)
	if err != nil {
		return code, resp, err
	}
	// TODO format response
	return code, resp, nil
}

// {
// 	"error_code": 0,
// 	"error_msg": "",
// 	"result": {
// 			"sf_order_id": "3165848793513984",//顺丰订单号
// 			"shop_order_id": "14570817752333",//商家订单号
// 			"rider_name": "1509935206",//配送员姓名
// 			"rider_phone": "18811588750",//配送员联系方式
// 			"rider_lng":"116.359337",//配送员经度
// 			"rider_lat":"40.020761",//配送员纬度
// 			"upload_time":"1510733606",//坐标上传时间（秒级时间戳）
// 	}
// }
func (c *SFExpressCityClient) RiderLatestPositionRequest(request *RiderLatestPositionRequest) (int, *Response, error) {
	code, resp, err := c.do(RIDER_LATEST_POSITION_FORMAT, request)
	if err != nil {
		return code, resp, err
	}
	// TODO format response
	return code, resp, nil
}

// {
// 	"error_code": 0,
// 	"error_msg": "success",
// 	"result":{
// 			'url':"https://shopic.sf-express.com/clover?order=kRn28C12UoCPWCKgM5FIONLdUGTjuawmVXNFYSVlPQ4",//H5页面URL
// 			"sf_order_id": "3165848793513984",//顺丰订单号
// 			"shop_order_id": "15104092022333",//商家订单号
// 	}
// }
func (c *SFExpressCityClient) RiderViewV2(request *RiderViewV2Request) (int, *Response, error) {
	code, resp, err := c.do(RIDER_VIEW_V2_FORMAT, request)
	if err != nil {
		return code, resp, err
	}
	// TODO format response
	return code, resp, nil
}

// {
// 	"error_code": 0,
// 	"error_msg": "",
// 	"error_data": null,
// 	"result": [
// 			{
// 					"shop_id": 3242477352000,
// 					"sf_order_id": 3233276321394156552,
// 					"shop_order_id": "100000162004032",
// 					"url_index": "rider_status",
// 					"operator_name": "测试",
// 					"operator_phone": "18270000000",
// 					"rider_lng": 113.90560399999999674,
// 					"rider_lat": 22.558814000000001698,
// 					"order_status": 10,
// 					"status_desc": "配送员已指派",
// 					"push_time": 1541746298
// 			},
// 			{
// 					"shop_id": 3242477352000,
// 					"sf_order_id": 3233276321394156552,
// 					"shop_order_id": "100000162004032",
// 					"url_index": "rider_status",
// 					"operator_name": "测试",
// 					"operator_phone": "18270000000",
// 					"rider_lng": 113.90560700000000338,
// 					"rider_lat": 22.558820000000000761,
// 					"order_status": 12,
// 					"status_desc": "配送员已到店",
// 					"push_time": 1541746299
// 			},
// 			{
// 					"shop_id": 3242477352000,
// 					"sf_order_id": 3233276321394156552,
// 					"shop_order_id": "100000162004032",
// 					"url_index": "rider_status",
// 					"operator_name": "测试",
// 					"operator_phone": "18270000000",
// 					"rider_lng": 113.90561300000000244,
// 					"rider_lat": 22.558828999999999354,
// 					"order_status": 15,
// 					"status_desc": "配送员已取货",
// 					"push_time": 1541746306
// 			},
// 			{
// 					"shop_id": 3242477352000,
// 					"sf_order_id": 3233276321394156552,
// 					"shop_order_id": "100000162004032",
// 					"url_index": "order_complete",
// 					"operator_name": "测试",
// 					"operator_phone": "18270000000",
// 					"rider_lng": 113.90560700000000338,
// 					"rider_lat": 22.558821999999999264,
// 					"order_status": 17,
// 					"status_desc": "配送员完成",
// 					"push_time": 1541746312
// 			}
// 	]
// }
func (c *SFExpressCityClient) GetCallbackInfo(request *GetCallbackInfoRequest) (int, *Response, error) {
	code, resp, err := c.do(GET_CALLBACK_INFO_FORMAT, request)
	if err != nil {
		return code, resp, err
	}
	// TODO format response
	return code, resp, nil
}

// {
// 	"error_code": 0,
// 	"error_msg": "",
// 	"error_data": null,
// 	"result": {
// 			"sf_order_id": "3165848793513984",//顺丰订单号
// 			"shop_order_id": "15104092022333",//商家订单号
// 			"gratuity_fee": 10, //本次加小费金额
// 			"total_gratuity_fee":20,//该订单总的加小费金额
// 			"push_time": "1510680568" //推送时间
// 	}
// }
func (c *SFExpressCityClient) AddOrderGratuityFee(request *AddOrderGratuityFeeRequest) (int, *Response, error) {
	code, resp, err := c.do(ADD_ORDER_GRATUITY_FEE, request)
	if err != nil {
		return code, resp, err
	}
	// TODO format response
	return code, resp, nil
}

func (c *SFExpressCityClient) do(apiFormat string, request interface{}) (int, *Response, error) {

	payload, err := json.Marshal(request)
	if err != nil {
		return 400, nil, err
	}

	signature := c.signer.Sign(string(payload))
	responseBody := bytes.NewBuffer([]byte(payload))

	url := fmt.Sprintf(apiFormat, signature)

	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(url, "application/json", responseBody)
	//Handle Error
	if err != nil {
		return 400, nil, err
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 400, nil, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 500, nil, err
	}

	return resp.StatusCode, &response, nil
}
