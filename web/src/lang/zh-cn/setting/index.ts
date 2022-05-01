/**
 * Copyright (c) 2022 ysicing All rights reserved.
 * Use of this source code is governed by AGPL-3.0-or-later
 * license that can be found in the LICENSE file.
 */

export default {
      "setting.title": "平台设置",
      "setting.desc": "",

      "setting.message.push": "消息推送",
      "setting.message.hook.addr": "webhook地址",
      "setting.message.hook.key": "webhook秘钥",
      "setting.message.smtp": "邮件SMTP服务地址",
      "setting.message.smtp.enabled": "启用ssl端口",
      "setting.message.smtp.port": "SMTP服务端口",
      "setting.message.smtp.user": "邮件推送人用户名",
      "setting.message.smtp.password": "邮件推送人密码",
      "setting.message.smtp.test": "邮件测试收件地址",
      "setting.message.hook.switch": "webhook推送开关",
      "setting.message.mail.switch": "email推送开关",
      "setting.message.action.mail": "邮件测试",
      "setting.message.action.hook": "hook测试",

      "setting.adv": "进阶设置",
      "setting.adv.env": "自定义环境",
      "setting.adv.env.tips": "环境名称",
      "setting.adv.env.add": "添加环境",
      "setting.adv.query.open": "开启查询审核",
      "setting.adv.query.limit": "查询最大Limit限制",
      "setting.adv.query.expire": "查询时限",
      "setting.adv.query.mins": "分",
      "setting.adv.query.export": "查询导出",
      "setting.adv.query.register": "开启用户注册",

      "setting.ldap": "LDAP设置",
      "setting.ldap.enabled": "Ldap地址",
      "setting.ldap.ssl": "启用ldaps",
      "setting.ldap.url": "Ldap 地址",
      "setting.ldap.url.tips": "[ip地址或者域名]:[端口号]",
      "setting.ldap.dn": "LDAP管理员DN",
      "setting.ldap.dn.tips": "请填写管理员DN",
      "setting.ldap.password": "LDAP管理员密码",
      "setting.ldap.password.tips": "请填写管理员密码",
      "setting.ldap.filter": "LDAP_Search filter",
      "setting.ldap.filter.tips": "例如:(&(objectClass=organizationalPerson)(sAMAccountName=%s))，%s为占位符必须存在",
      "setting.ldap.sc": "LDAP_SCBASE",
      "setting.ldap.test": "ldap测试",
      "setting.ldap.alert": "1.LDAP登录用户名，必须全局唯一。ldap配置请参考Grafana。<br><br>2.由于各个邮件服务提供商对于垃圾邮件过滤的机制各不相同，可能会造成邮件无法接收的情况。所以使用前请测试是否稳定。对于使用了ssl安全协议的stmp连接需勾选启动ssl端口复选框<br><br>3.只有开启相应的消息推送开关后，消息推送才会开启。",


      "setting.data.clear": "数据清除",
      "setting.data.clear.tips": "你确定要删除工单吗?",
      "setting.data.clear.order": "指定时间内工单",
      "setting.data.clear.query": "指定时间内查询工单",
      "setting.data.clear.alert": "1.设置最大Limit数后，所有的查询语句的查询结果都不会超过这个数值。<br><br>2.查询审核开关开启后，所有的查询都必须通过管理员同意才能进行。关闭则可自主查询"

}