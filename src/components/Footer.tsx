import { For, type Accessor } from "solid-js";
import { dim, TextAttributes } from "@opentui/core";
import { colors } from "../lib/theme";

type KeyDef = {
  bind: string
  desc: string
};

const defaultKeys: KeyDef[] = [
  { bind: "j|k", desc: "scroll" },
  { bind: "q", desc: "quit" },
];

function KeyDesc(props: {
  key: KeyDef
}) {

  return (
    <box
      flexDirection="row"
      columnGap={1}
    >
      <text fg={colors.dark.highlight} attributes={TextAttributes.BOLD}>{props.key.bind}</text>
      <text fg={colors.dark.text}>{props.key.desc}</text>
    </box>
  );
}

export default function Footer(props: {
  width: number,
  keys: Accessor<KeyDef[]>
}) {
  return (
    <box flexDirection="column" alignItems="center" width={props.width}>
      <text fg={colors.dark.border}>{"─".repeat(props.width)}</text>
      <box flexDirection="row" justifyContent="center" columnGap={2}>
        <For each={props.keys()}>
          {(key) => (
            <KeyDesc key={key} />
          )}
        </For>
        <For each={defaultKeys}>
          {(key) => (
            <KeyDesc key={key} />
          )}
        </For>
      </box>
    </box>
  );
}
