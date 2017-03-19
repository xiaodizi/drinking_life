package utils

const Is_used = 1     //后台管理用户启用
const Is_not_used = 0 //后台管理用户禁用

const Is_frame_up = 1   //上架
const Is_frame_down = 0 //下架

const Is_use = 1     //使用
const Is_not_use = 0 //不使用

const Is_display = 1     //广告允许显示
const Is_not_display = 0 //广告禁止显示

const (
	KC_RAND_KIND_NUM   = 0 //纯数字
	KC_RAND_KIND_LOWER = 1 //小写字母
	KC_RAND_KIND_UPPER = 2 //大写字母
	KC_RAND_KIND_ALL   = 3 //数字、大小写与字母
)

const PAY = 1    //支付
const CHARGE = 2 //充值

const ACCOUNT_TYPE = 1 //主账户

<<<<<<< HEAD
const Regular = "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\\\d{8}$"
=======
const ALERD_PAYID = 11000 //已支付
const UNPAID = 11001      //未支付

const ALERD_RECEIVE = 21000 //已领取
const UPRECEIVE = 21001     //未领取
>>>>>>> 27507c37804c1a4a1c8c65557a42a8282fd8463f
