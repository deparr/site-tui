import { TextAttributes } from "@opentui/core";

export default function Contact() {
  return (<box paddingLeft={2}>
    <text>lne 1</text>
    <box paddingLeft={2}>
      <text>line 2 should be indented more</text>
    </box>
    <text>line 3 should no same as line 1</text>
    <text attributes={TextAttributes.STRIKETHROUGH}>STRIKE</text>
    <text><span style={{ attribute: TextAttributes.STRIKETHROUGH }}>SPAN STRIKE</span></text>
  </box>);
}
