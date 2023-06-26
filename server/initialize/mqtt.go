package initialize

import (
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/panjf2000/ants/v2"
	"go.uber.org/zap"
	"log"
	"os"
	"strings"
	"time"
)

const (
	Qos0 = byte(0)
	Qos1 = byte(1)
	Qos2 = byte(2)
)

const (
	Topic_Inhand_Data  = "topic/data/+/#"     //映瀚通网关数据通道
	Topic_Beilai_Data  = "datatopic/device/#" //钡铼网关数据通道
	Topic_Huachen_Data = "v2/testTopic/pub"   //华辰网关数据通道
)

var TopicList = map[string]byte{
	Topic_Inhand_Data:  Qos0,
	Topic_Beilai_Data:  Qos0,
	Topic_Huachen_Data: Qos0,
}

func InitMqtt() {
	opts := mqtt.NewClientOptions()
	broker := fmt.Sprintf("%s:%d", global.GVA_CONFIG.Mqtt.Host, global.GVA_CONFIG.Mqtt.Port)
	opts.AddBroker(broker)
	opts.SetUsername(global.GVA_CONFIG.Mqtt.Username)
	opts.SetPassword(global.GVA_CONFIG.Mqtt.Password)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	p1, _ := ants.NewPool(500)
	//pOther, _ := ants.NewPool(500)
	var messageHandler mqtt.MessageHandler = func(c mqtt.Client, d mqtt.Message) {
		topic := d.Topic()
		if strings.HasPrefix(topic, "datatopic/device/") {
			// 匹配以 "datatopic/device/" 开头的主题
			_ = p1.Submit(func() {
				global.GVA_LOG.Info("收到钡铼网关数据", zap.String("topic", d.Topic()), zap.String("payload", string(d.Payload())))
			})
		} else if strings.HasPrefix(topic, "topic/data/") {
			// 匹配 Topic_Inhand_Data 主题
			_ = p1.Submit(func() {
				global.GVA_LOG.Info("收到映瀚通网关数据", zap.String("topic", d.Topic()), zap.String("payload", string(d.Payload())))
				ProInHandData(topic, d.Payload())
			})
		} else if topic == Topic_Huachen_Data {
			// 匹配 Topic_Huachen_Data 主题
			_ = p1.Submit(func() {
				global.GVA_LOG.Info("收到华辰网关数据", zap.String("topic", d.Topic()), zap.String("payload", string(d.Payload())))
			})
		} else {
			global.GVA_LOG.Info("收到未知topic数据", zap.String("topic", d.Topic()), zap.String("payload", string(d.Payload())))
			fmt.Println("undefine topic")
		}
	}
	// 批量订阅
	if token := client.SubscribeMultiple(TopicList, messageHandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	global.GVA_MQTT = &client
}
func Publish(topic string, payload []byte) {
	token := (*global.GVA_MQTT).Publish(topic, Qos0, false, payload)
	token.Wait()
}
func ProInHandData(topic string, payload []byte) {
	segments := strings.Split(topic, "/")
	// 检查分割后的片段数量是否符合预期
	if len(segments) == 4 && segments[0] == "topic" && segments[1] == "data" {
		gateid := segments[2]
		deviceid := segments[3]
		// 获取 "Device" 和 "timestamp" 的值
		value_dict := make(map[string]interface{})
		err := json.Unmarshal(payload, &value_dict)
		if err != nil {
			fmt.Println("Error parsing payload:", err)
			return
		}
		device := value_dict["Device"]
		timestamp := value_dict["timestamp"]

		// 获取 "Data" 字段的值
		data := make(map[string]interface{})
		if dataObj, ok := value_dict["Data"].(map[string]interface{}); ok {
			data = dataObj
		}

		// 打印解析结果
		global.GVA_LOG.Info("Received InHand data",
			zap.String("gateid", gateid),
			zap.String("deviceid", deviceid),
			zap.Any("Device", device),
			zap.Any("timestamp", timestamp),
			zap.Any("Data", data),
		)
		// 将 data 和 timestamp 转换为 JSON 字符串
		dataJSON, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error marshaling data:", err)
			return
		}
		timestampJSON, err := json.Marshal(timestamp)
		if err != nil {
			fmt.Println("Error marshaling timestamp:", err)
			return
		}
		// 存储到 Redis 的哈希值
		err = global.GVA_REDIS.HMSet(context.Background(), gateid+":"+deviceid, "data", dataJSON, "timestamp", timestampJSON).Err()
		if err != nil {
			fmt.Println("Error storing data in Redis:", err)
			return
		}
	} else {
		global.GVA_LOG.Info("映瀚通网关数据格式错误", zap.String("topic", topic), zap.String("payload", string(payload)))
	}

}
