/**
 * Copyright (c) 2022 ysicing All rights reserved.
 * Use of this source code is governed by AGPL-3.0-or-later
 * license that can be found in the LICENSE file.
 */

export default {

      "setting.title": "Settings",
      "setting.desc": "The configuration suits your style",

      "setting.message.push": "Message Push",
      "setting.message.hook.addr": "Dingding hook addr",
      "setting.message.hook.key": "Dingding secret key",
      "setting.message.smtp": "Email STMP addr",
      "setting.message.smtp.enabled": "Enabled SSL port",
      "setting.message.smtp.port": "SMTP port",
      "setting.message.smtp.user": "SMTP user",
      "setting.message.smtp.password": "SMTP password",
      "setting.message.smtp.test": "Mail test receiver address",
      "setting.message.hook.switch": "Hook push switch",
      "setting.message.mail.switch": "Email push switch",
      "setting.message.action.mail": "Mail test",
      "setting.message.action.hook": "Hook test",

      "setting.adv": "Advanced Settings",
      "setting.adv.env": "Environment",
      "setting.adv.env.tips": "Environment name",
      "setting.adv.env.add": "add env",
      "setting.adv.query.open": "Enable query audit",
      "setting.adv.query.limit": "Query maximum limit",
      "setting.adv.query.expire": "Query time",
      "setting.adv.query.mins": "mins",
      "setting.adv.query.export": "Query Export",
      "setting.adv.query.register": "Register",

      "setting.ldap": "LDAP Settings",
      "setting.ldap.url": "Ldap addr",
      "setting.ldap.url.tips": "[ip/domain]:[port]",
      "setting.ldap.enabled": "Enable SSL",
      "setting.ldap.dn": "LDAP Admin DN",
      "setting.ldap.dn.tips": "Enter the administrator DN",
      "setting.ldap.password": "LDAP Admin Password",
      "setting.ldap.password.tips": "Please enter the administrator password",
      "setting.ldap.filter": "LDAP_Search filter",
      "setting.ldap.filter.tips": "For example (&(objectClass=organizationalPerson)(sAMAccountName=%s)), %s must be a placeholder",
      "setting.ldap.sc": "LDAP_SCBASE",
      "setting.ldap.test": "LDAP test",
      "setting.ldap.alert": "1.LDAP login user name must be globally unique. For ldap configuration, see Grafana. <br><br>2. Different mail service providers have different filtering mechanisms for junk mails. As a result, mails may fail to be received. So please test whether it is stable before using. For STMP connections that use SSL, select the enable SSL port check box <br><br>3. Message push is enabled only after the corresponding message push switch is enabled",

      "setting.data.clear": "Data Clear",
      "setting.data.clear.tips": "Are you sure to delete the work order?",
      "setting.data.clear.order": "Order before",
      "setting.data.clear.query": "Query before",
      "setting.data.clear.alert": "1. After the maximum Limit is set, the results of all query statements cannot exceed the Limit. <br><br>2. After the query audit switch is enabled, all queries must be approved by the administrator. Off can be independent query"

}
