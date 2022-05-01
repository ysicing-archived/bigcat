import dayjs from 'dayjs';
import configs from "../../package.json";

const FeatureList = [
      "TODO",
]

const Announce = [
      "大猫平台采用 AGPL 3.0 许可证, 以下为许可中相关的义务与责任:",
      "  1.未经原作者授权不得用于任何商业用途, 如需商业使用请联系原作者获得授权;",
      "  2.本平台所有条款符合相应开源许可, 请严格按照相关许可使用及开发;",
      "  3.内部使用不在上述条款之内.",
      "免责声明:",
      "  由平台所产生的一切后果, 与平台作者本人无关! 请在进行安全评估及测试体验后使用.",
]

const Copyright = `BigCat(v${configs.version }) © ${dayjs().format("YYYY")} ysicing`

export { FeatureList, Copyright, Announce }